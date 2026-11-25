package gate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type TaskState string

const (
	TaskNew         TaskState = "NEW"
	TaskSpecRed     TaskState = "SPEC_RED"
	TaskImplemented TaskState = "IMPLEMENTED"
	TaskVerified    TaskState = "VERIFIED"
	TaskBlocked     TaskState = "BLOCKED"
	TaskClosed      TaskState = "CLOSED"
)

type TaskSource struct {
	Type string `json:"type"`
	ID   string `json:"id,omitempty"`
}

type TaskEvidence struct {
	Regression   string `json:"regression,omitempty"`
	ProbeReport  string `json:"probeReport,omitempty"`
	GateReport   string `json:"gateReport,omitempty"`
	VerifiedTree string `json:"verifiedTree,omitempty"`
}

type Task struct {
	Version       int          `json:"version"`
	ID            string       `json:"id"`
	Title         string       `json:"title"`
	Component     string       `json:"component"`
	State         TaskState    `json:"state"`
	Source        TaskSource   `json:"source"`
	Requirements  []string     `json:"requirements"`
	Specs         []string     `json:"specs"`
	Probes        []string     `json:"probes"`
	Evidence      TaskEvidence `json:"evidence"`
	BlockedReason string       `json:"blockedReason,omitempty"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
	ClosedAt      *time.Time   `json:"closedAt,omitempty"`
}

type Feedback struct {
	Version   int      `json:"version"`
	ID        string   `json:"id"`
	Component string   `json:"component"`
	Severity  string   `json:"severity"`
	Steps     []string `json:"steps"`
	Expected  string   `json:"expected"`
	Actual    string   `json:"actual"`
	Reviewer  string   `json:"reviewer"`
	Evidence  []string `json:"evidence,omitempty"`
}

type Regression struct {
	Version   int    `json:"version"`
	TaskID    string `json:"taskId"`
	CaseName  string `json:"caseName"`
	Component string `json:"component"`
	Expected  string `json:"expected"`
	Actual    string `json:"actual"`
	Status    string `json:"status"`
}

type LedgerEntry struct {
	TaskID    string    `json:"taskId"`
	Component string    `json:"component"`
	State     TaskState `json:"state"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var idPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9_-]*$`)

func ValidateID(id string) error {
	if !idPattern.MatchString(id) {
		return fmt.Errorf("ID %q 非法：仅允许字母、数字、-、_，且不能以符号开头", id)
	}
	return nil
}

func TaskDir(root, id string) string {
	return filepath.Join(root, "wg", "tasks", id)
}

func TaskPath(root, id string) string {
	return filepath.Join(TaskDir(root, id), "task.json")
}

func componentSlug(component string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(component) {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			b.WriteRune(r)
		} else if b.Len() > 0 {
			b.WriteByte('-')
		}
	}
	return strings.Trim(b.String(), "-")
}

func defaultSpec(root, component string) []string {
	candidate := filepath.Join("wg", "examples", componentSlug(component), "spec.json")
	if _, err := os.Stat(filepath.Join(root, candidate)); err == nil {
		return []string{filepath.ToSlash(candidate)}
	}
	return nil
}

func NewTask(root, id, component, title string, requirements []string) (Task, error) {
	if err := ValidateID(id); err != nil {
		return Task{}, err
	}
	if component == "" || title == "" {
		return Task{}, fmt.Errorf("component 与 title 不能为空")
	}
	dir := TaskDir(root, id)
	if _, err := os.Stat(dir); err == nil {
		return Task{}, fmt.Errorf("任务目录已存在: %s", dir)
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return Task{}, err
	}
	now := time.Now()
	task := Task{
		Version: 1, ID: id, Title: title, Component: component, State: TaskNew,
		Source: TaskSource{Type: "plan"}, Requirements: requirements,
		Specs: defaultSpec(root, component), CreatedAt: now, UpdatedAt: now,
	}
	if err := SaveTask(root, task); err != nil {
		return Task{}, err
	}
	if err := UpdateLedger(root, task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func ImportFeedback(root, taskID, file string) (Task, error) {
	if err := ValidateID(taskID); err != nil {
		return Task{}, err
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return Task{}, fmt.Errorf("读取反馈失败: %w", err)
	}
	var feedback Feedback
	if err := json.Unmarshal(data, &feedback); err != nil {
		return Task{}, fmt.Errorf("解析反馈失败: %w", err)
	}
	if feedback.Version != 1 || feedback.ID == "" || feedback.Component == "" || feedback.Expected == "" || feedback.Actual == "" || len(feedback.Steps) == 0 {
		return Task{}, fmt.Errorf("反馈字段不完整：需要 version=1、id、component、steps、expected、actual")
	}
	task, err := NewTask(root, taskID, feedback.Component, feedback.ID, []string{feedback.Expected})
	if err != nil {
		return Task{}, err
	}
	task.Source = TaskSource{Type: "human-feedback", ID: feedback.ID}
	task.State = TaskSpecRed
	regression := Regression{
		Version: 1, TaskID: taskID, CaseName: "反馈回归：" + feedback.ID,
		Component: feedback.Component, Expected: feedback.Expected, Actual: feedback.Actual,
		Status: "RED_REQUIRED",
	}
	task.Evidence.Regression = filepath.ToSlash(filepath.Join("wg", "tasks", taskID, "regression.json"))
	if err := writeJSON(filepath.Join(TaskDir(root, taskID), "feedback.json"), feedback); err != nil {
		return Task{}, err
	}
	if err := writeJSON(filepath.Join(root, task.Evidence.Regression), regression); err != nil {
		return Task{}, err
	}
	if err := SaveTask(root, task); err != nil {
		return Task{}, err
	}
	if err := UpdateLedger(root, task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func LoadTask(root, id string) (Task, error) {
	if err := ValidateID(id); err != nil {
		return Task{}, err
	}
	data, err := os.ReadFile(TaskPath(root, id))
	if err != nil {
		return Task{}, fmt.Errorf("读取任务失败: %w", err)
	}
	var task Task
	if err := json.Unmarshal(data, &task); err != nil {
		return Task{}, err
	}
	if task.Version != 1 || task.ID != id {
		return Task{}, fmt.Errorf("任务清单版本或 ID 不匹配")
	}
	return task, nil
}

func SaveTask(root string, task Task) error {
	task.UpdatedAt = time.Now()
	if err := writeJSON(TaskPath(root, task.ID), task); err != nil {
		return err
	}
	return UpdateLedger(root, task)
}

func UpdateLedger(root string, task Task) error {
	path := filepath.Join(root, "wg", "tasks", "ledger.json")
	entries := []LedgerEntry{}
	if data, err := os.ReadFile(path); err == nil {
		if err := json.Unmarshal(data, &entries); err != nil {
			return err
		}
	}
	entry := LedgerEntry{TaskID: task.ID, Component: task.Component, State: task.State, UpdatedAt: time.Now()}
	found := false
	for i := range entries {
		if entries[i].TaskID == task.ID {
			entries[i] = entry
			found = true
			break
		}
	}
	if !found {
		entries = append(entries, entry)
	}
	return writeJSON(path, entries)
}

func writeJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(path)
		return os.Rename(tmp, path)
	}
	return nil
}

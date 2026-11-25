package gate

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

func gitOutput(root string, args ...string) (string, error) {
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, out)
	}
	return strings.TrimSpace(string(out)), nil
}

func RepoRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("定位 Git 仓库根失败: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func EnsureStaged(root string) error {
	unstaged, err := gitOutput(root, "diff", "--name-only")
	if err != nil {
		return err
	}
	if unstaged != "" {
		return fmt.Errorf("存在未暂存改动，请先 git add 本轮文件：\n%s", unstaged)
	}
	untracked, err := gitOutput(root, "ls-files", "--others", "--exclude-standard", "--", "wg", ".githooks")
	if err != nil {
		return err
	}
	if untracked != "" {
		return fmt.Errorf("存在未暂存门禁源码，请先 git add：\n%s", untracked)
	}
	return nil
}

func StagedTree(root string) (string, error) {
	return gitOutput(root, "write-tree")
}

func HeadTree(root string) (string, error) {
	out, err := gitOutput(root, "rev-parse", "HEAD")
	if err != nil {
		return "", nil
	}
	return out, nil
}

func ChangedFiles(root, tree string) ([]string, error) {
	args := []string{"diff", "--name-only"}
	if head, err := HeadTree(root); err == nil && head != "" {
		args = append(args, head, tree)
	} else {
		args = []string{"ls-tree", "-r", "--name-only", tree}
	}
	out, err := gitOutput(root, args...)
	if err != nil {
		return nil, err
	}
	if out == "" {
		return []string{}, nil
	}
	files := strings.Split(out, "\n")
	for i := range files {
		files[i] = filepath.ToSlash(files[i])
	}
	sort.Strings(files)
	return files, nil
}

func ChangedFilesBetween(root, from, to string) ([]string, error) {
	out, err := gitOutput(root, "diff", "--name-only", from, to)
	if err != nil {
		return nil, err
	}
	if out == "" {
		return []string{}, nil
	}
	files := strings.Split(out, "\n")
	for i := range files {
		files[i] = filepath.ToSlash(files[i])
	}
	sort.Strings(files)
	return files, nil
}

func OnlyTaskChanges(root, from, to string) (bool, error) {
	changes, err := ChangedFilesBetween(root, from, to)
	if err != nil {
		return false, err
	}
	for _, file := range changes {
		if !strings.HasPrefix(file, "wg/tasks/") {
			return false, nil
		}
	}
	return true, nil
}

func WithoutTaskFiles(files []string) []string {
	out := make([]string, 0, len(files))
	for _, file := range files {
		if !strings.HasPrefix(file, "wg/tasks/") {
			out = append(out, file)
		}
	}
	return out
}

func FeatureWindows(root string) ([]string, error) {
	dir := filepath.Join(root, "wg", "examples")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var features []string
	for _, entry := range entries {
		if !entry.IsDir() || entry.Name() == "doc" || entry.Name() == "internal" {
			continue
		}
		if _, err := os.Stat(filepath.Join(dir, entry.Name(), "main.go")); err == nil {
			features = append(features, entry.Name())
		}
	}
	sort.Strings(features)
	if len(features) < 60 {
		return nil, fmt.Errorf("功能窗数量异常：%d（应 ≥60）", len(features))
	}
	return features, nil
}

func exeName(name string) string {
	if runtime.GOOS == "windows" {
		return name + ".exe"
	}
	return name
}

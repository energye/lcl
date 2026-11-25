package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/energye/lcl/wg/internal/gate"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "run":
		err = run()
	case "check":
		err = check()
	case "install-hook":
		err = installHook()
	case "task":
		err = taskCommand(os.Args[2:])
	case "feedback":
		err = feedbackCommand(os.Args[2:])
	case "probe":
		err = probeCommand(os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "[wg-gate]", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "用法:")
	fmt.Fprintln(os.Stderr, "  wg-gate run [--tree <staged-tree>]")
	fmt.Fprintln(os.Stderr, "  wg-gate check --tree <staged-tree>")
	fmt.Fprintln(os.Stderr, "  wg-gate install-hook")
	fmt.Fprintln(os.Stderr, "  wg-gate task new|status|verify|close|reopen ...")
	fmt.Fprintln(os.Stderr, "  wg-gate feedback import ...")
	fmt.Fprintln(os.Stderr, "  wg-gate probe new ...")
}

func root() (string, error) { return gate.RepoRoot() }

func run() error {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	tree := fs.String("tree", "", "staged tree 哈希；默认取 git write-tree")
	reportDir := fs.String("report-dir", ".shots/gates", "报告目录")
	probeFile := fs.String("probe-file", "wg/gate.probes.json", "性能探针清单")
	timeout := fs.Duration("timeout", 10*time.Minute, "单步骤默认超时")
	auditTimeout := fs.Duration("audit-timeout", 2*time.Minute, "单窗审计超时")
	_ = fs.Parse(os.Args[2:])
	repo, err := root()
	if err != nil {
		return err
	}
	report, err := gate.Run(context.Background(), repo, *tree, gate.Options{
		ReportDir: *reportDir, ProbeFile: *probeFile, DefaultTimeout: *timeout, AuditTimeout: *auditTimeout,
		Stdout: os.Stdout, Stderr: os.Stderr,
	})
	if err != nil {
		return err
	}
	fmt.Printf("[wg-gate] %s tree=%s audits=%d/%d report=%s\n",
		report.Status, report.TreeHash, report.Audits.Passed, report.Audits.Expected,
		gate.ReportPath(repo, *reportDir, report.TreeHash))
	if report.Status != gate.StatusPass {
		return fmt.Errorf("完整门禁未通过")
	}
	return nil
}

func check() error {
	fs := flag.NewFlagSet("check", flag.ExitOnError)
	tree := fs.String("tree", "", "staged tree 哈希（必填）")
	reportDir := fs.String("report-dir", ".shots/gates", "报告目录")
	probeFile := fs.String("probe-file", "wg/gate.probes.json", "性能探针清单")
	_ = fs.Parse(os.Args[2:])
	if *tree == "" {
		return fmt.Errorf("check 必须提供 --tree")
	}
	repo, err := root()
	if err != nil {
		return err
	}
	if err := gate.Check(repo, *tree, gate.Options{ReportDir: *reportDir, ProbeFile: *probeFile}); err != nil {
		return err
	}
	fmt.Printf("[wg-gate] PASS tree=%s\n", *tree)
	return nil
}

func installHook() error {
	repo, err := root()
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "config", "core.hooksPath", ".githooks")
	cmd.Dir = repo
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("安装 hook 失败: %w\n%s", err, out)
	}
	fmt.Println("已启用 .githooks/pre-commit；后续 commit 将校验完整门禁报告。")
	return nil
}

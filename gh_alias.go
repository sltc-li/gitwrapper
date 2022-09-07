package gitwrapper

import (
	"os"
	"os/exec"
	"strings"
)

type GHPRAlias struct{}

func runCmd(sCmd string) error {
	logger.Println(green(sCmd))
	cmd := exec.Command("sh", "-c", sCmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (a GHPRAlias) View(pr string) error {
	if pr != "" {
		return runCmd("gh pr view " + pr)
	}
	return runCmd("gh pr view")
}

func (a GHPRAlias) List() error {
	return runCmd("gh pr list")
}

func (a GHPRAlias) Checkout(issue string) error {
	return runCmd("gh pr checkout " + strings.TrimPrefix(issue, "#"))
}

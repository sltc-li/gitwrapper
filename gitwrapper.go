package gitwrapper

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	logger = log.New(os.Stdout, "[gw] ", 0)
)

func green(s string) string {
	return "\033[0;32m" + s + "\033[0m"
}

func joinArgs(args []string) string {
	var cc []string
	for _, a := range args {
		if strings.Contains(a, " ") {
			a = `"` + strings.Replace(a, `"`, `\"`, -1) + `"`
		}
		cc = append(cc, a)
	}
	return strings.Join(cc, " ")
}

func runGitCmd(trace bool, args ...string) (string, error) {
	if trace {
		logger.Println(green("git " + joinArgs(args)))
	}

	cmd := exec.Command("git", args...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	// skip git error if std output exists
	if err != nil && stdout.Len() == 0 {
		return "", err
	}

	if trace {
		buffCopy := stdout
		_, _ = io.Copy(os.Stdout, &buffCopy)
	}

	return stdout.String(), nil
}

type Branch struct {
	Name                string
	IsCurrent, IsRemote bool
}

func getAllBranches() ([]Branch, error) {
	o, err := runGitCmd(false, "branch", "-a")
	if err != nil {
		return nil, err
	}
	var bb []Branch
	for _, row := range strings.Split(o, "\n") {
		if len(row) == 0 {
			continue
		}

		b := Branch{Name: row[2:]}

		if strings.HasPrefix(row, "* ") {
			b.IsCurrent = true
		}

		if strings.HasPrefix(b.Name, "remotes/origin/") {
			b.Name = b.Name[15:]
			b.IsRemote = true
		}

		bb = append(bb, b)
	}
	return bb, nil
}

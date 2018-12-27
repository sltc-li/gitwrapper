package gitwrapper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (rc RepoConfig) Commit(add bool, msg string) error {
	// step 1
	_, err := runGitCmd(true, "status")
	if err != nil {
		return err
	}

	// step 2
	if add {
		fmt.Print("press enter to continue...")
		reader := bufio.NewReader(os.Stdin)
		_, err = reader.ReadString('\n')
		if err != nil {
			return err
		}

		_, err = runGitCmd(true, "add", "-A")
		if err != nil {
			return err
		}
	}

	// step 3
	if len(msg) == 0 {
		msg = rc.getDefaultCommitMessage()
	}
	_, err = runGitCmd(true, "commit", "-m", msg)
	if err != nil {
		return err
	}

	return nil
}

func (rc RepoConfig) getDefaultCommitMessage() string {
	var m string
	ss := strings.SplitN(rc.CurrentBranch, "/", 2)
	if len(ss) > 1 {
		m = "[" + strings.Title(ss[0]) + "] "
		ss = ss[1:]
	}
	if len(ss[0]) == 0 {
		return m
	}
	m += strings.Replace(strings.Title(ss[0][0:1])+ss[0][1:], "-", " ", -1)
	return strings.Replace(m, "_", " ", -1)
}

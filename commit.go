package gitwrapper

import (
	"strings"
)

func (rc RepoConfig) Commit(add bool, msg string) error {
	// step 1
	_, err := runGitCmd(true, "git status")
	if err != nil {
		return err
	}

	// step 2
	if add {
		_, err = runGitCmd(true, "git add -A")
		if err != nil {
			return err
		}
	}

	// step 3
	if len(msg) == 0 {
		msg = rc.getDefaultCommitMessage()
	}
	_, err = runGitCmd(true, "git commit -am '"+msg+"'")
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

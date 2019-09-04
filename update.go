package gitwrapper

import "strings"

func (rc RepoConfig) Update() error {
	o, err := runGitCmd(true, "git status")
	if err != nil {
		return err
	}

	var clean bool
	if strings.Contains(o, "nothing to commit, working tree clean") {
		clean = true
	}

	// step 1
	if !clean {
		_, err = runGitCmd(true, "git stash")
		if err != nil {
			return err
		}
	}

	// step 2
	_, err = runGitCmd(true, "git fetch -p")
	if err != nil {
		return err
	}

	// step 3
	_, err = runGitCmd(true, "git rebase origin/"+rc.CurrentBranch)
	if err != nil {
		return err
	}

	// step 4
	if clean {
		_, err = runGitCmd(true, "git status")
		if err != nil {
			return err
		}
	} else {
		_, err = runGitCmd(true, "git stash pop")
		if err != nil {
			return err
		}
	}

	return nil
}

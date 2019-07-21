package gitwrapper

import "strings"

func (rc RepoConfig) CreateBranch(branch string) error {
	o, err := runGitCmd(true, "git status")
	if branch == rc.CurrentBranch {
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
	_, err = runGitCmd(true, "git checkout "+rc.DefaultBranch)
	if err != nil {
		return err
	}

	// step 3
	_, err = runGitCmd(true, "git fetch -p")
	if err != nil {
		return err
	}

	// step 4
	_, err = runGitCmd(true, "git rebase origin/"+rc.DefaultBranch)
	if err != nil {
		return err
	}

	// step 5
	bb, err := getAllBranches()
	if err != nil {
		return err
	}
	for _, b := range bb {
		if b.IsCurrent || b.IsRemote {
			continue
		}
		_, err = runGitCmd(true, "git branch -D "+b.Name)
		if err != nil {
			return err
		}
	}

	// step 6
	var exists bool
	for _, b := range rc.Branches {
		if b == branch {
			exists = true
			break
		}
	}
	if exists {
		_, err = runGitCmd(true, "git checkout "+branch)
	} else {
		_, err = runGitCmd(true, "git checkout -b "+branch)
	}
	if err != nil {
		return err
	}

	// step 7
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

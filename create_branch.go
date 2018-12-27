package gitwrapper

import "strings"

func (rc RepoConfig) CreateBranch(branch string) error {
	o, err := runGitCmd(true, "status")
	if branch == rc.CurrentBranch {
		return err
	}

	var clean bool
	if strings.Contains(o, "nothing to commit, working tree clean") {
		clean = true
	}

	// step 1
	if !clean {
		_, err = runGitCmd(true, "stash")
		if err != nil {
			return err
		}
	}

	// step 2
	_, err = runGitCmd(true, "checkout", rc.DefaultBranch)
	if err != nil {
		return err
	}

	// step 3
	_, err = runGitCmd(true, "fetch", "-p")
	if err != nil {
		return err
	}

	// step 4
	_, err = runGitCmd(true, "rebase", "origin/"+rc.DefaultBranch)
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
		_, err = runGitCmd(true, "branch", "-D", b.Name)
		if err != nil {
			return err
		}
	}

	// step 6
	_, err = runGitCmd(true, "checkout", "-b", branch)
	if err != nil {
		return err
	}

	// step 7
	if clean {
		_, err = runGitCmd(true, "status")
		if err != nil {
			return err
		}
	} else {
		_, err = runGitCmd(true, "stash", "pop")
		if err != nil {
			return err
		}
	}

	return nil
}

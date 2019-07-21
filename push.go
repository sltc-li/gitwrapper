package gitwrapper

func (rc RepoConfig) Push(force bool) error {
	// step 1
	branch := rc.CurrentBranch
	if force {
		branch = "+" + rc.CurrentBranch
	}
	_, err := runGitCmd(true, "git push -u origin "+branch)
	if err != nil {
		return err
	}

	return nil
}

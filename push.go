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

func (rc RepoConfig) PushTags() error {
	// step 1
	_, err := runGitCmd(true, "git push origin --tags")
	if err != nil {
		return err
	}

	return nil
}

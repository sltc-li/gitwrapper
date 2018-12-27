package gitwrapper

func (rc RepoConfig) Push() error {
	// step 1
	_, err := runGitCmd(true, "push", "-u", "origin", rc.CurrentBranch)
	if err != nil {
		return err
	}

	return nil
}

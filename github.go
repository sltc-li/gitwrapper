package gitwrapper

import (
	"fmt"
	"os/exec"
)

func (rc RepoConfig) GetLatestCommitURL() (string, error) {
	commitID, err := runGitCmd(false, "git rev-parse HEAD")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://github.com/%s/%s/commit/%s", rc.Organization, rc.Repository, commitID), nil
}

func (rc RepoConfig) GetCompareURL() string {
	return fmt.Sprintf("https://github.com/%s/%s/compare/%s?expand=1", rc.Organization, rc.Repository, rc.CurrentBranch)
}

func (rc RepoConfig) GetRepositoryURL() string {
	return fmt.Sprintf("https://github.com/%s/%s/tree/%s", rc.Organization, rc.Repository, rc.CurrentBranch)
}

func (rc RepoConfig) OpenCompare() error {
	return open(rc.GetCompareURL())
}

func (rc RepoConfig) OpenRepo() error {
	return open(rc.GetRepositoryURL())
}

func open(url string) error {
	logger.Println("opening", url, "...")
	return exec.Command("open", url).Start()
}

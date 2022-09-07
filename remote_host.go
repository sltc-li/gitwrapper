package gitwrapper

import (
	"fmt"
	"os/exec"
)

func (rc RepoConfig) GetShortHash() (string, error) {
	commitID, err := runGitCmd(false, "git rev-parse --short HEAD")
	if err != nil {
		return "", err
	}
	return commitID, nil
}

func (rc RepoConfig) GetLatestCommitURL() (string, error) {
	commitID, err := runGitCmd(false, "git rev-parse HEAD")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s/%s/%s/commit/%s", rc.RemoteHost, rc.Organization, rc.Repository, commitID), nil
}

func (rc RepoConfig) GetCompareURL() string {
	return fmt.Sprintf("https://%s/%s/%s/compare/%s?expand=1", rc.RemoteHost, rc.Organization, rc.Repository, rc.CurrentBranch)
}

func (rc RepoConfig) GetRepositoryURL() string {
	return fmt.Sprintf("https://%s/%s/%s/tree/%s", rc.RemoteHost, rc.Organization, rc.Repository, rc.CommitHash)
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

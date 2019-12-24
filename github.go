package gitwrapper

import (
	"fmt"
	"os/exec"
)

func (rc RepoConfig) OpenCompare() error {
	url := fmt.Sprintf("https://github.com/%s/%s/compare/%s?expand=1",
		rc.Organization, rc.Repository, rc.CurrentBranch)
	return open(url)
}

func (rc RepoConfig) OpenRepo() error {
	url := fmt.Sprintf("https://github.com/%s/%s/tree/%s", rc.Organization, rc.Repository, rc.CurrentBranch)
	return open(url)
}

func open(url string) error {
	logger.Println("opening", url, "...")
	return exec.Command("open", url).Start()
}

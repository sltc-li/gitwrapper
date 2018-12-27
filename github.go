package gitwrapper

import (
	"fmt"
	"os/exec"
)

func (rc RepoConfig) OpenCompareURL() error {
	compareURL := fmt.Sprintf("https://github.com/%s/%s/compare/%s?expand=1",
		rc.Organization, rc.Repository, rc.CurrentBranch)
	logger.Println("opening", compareURL, "...")
	return exec.Command("open", compareURL).Start()
}

package gitwrapper

import (
	"strings"
	"time"
)

func (rc RepoConfig) AddReleaseTag() error {
	// step 1
	tag, err := genReleaseTag()
	if err != nil {
		return err
	}

	// step 2
	_, err = runGitCmd(true, "git tag "+tag)
	if err != nil {
		return err
	}

	return nil
}

func genReleaseTag() (string, error) {
	out, err := runGitCmd(false, "git tag")
	if err != nil {
		return "", err
	}
	tags := strings.Split(out, "\n")
	for {
		tag := _genReleaseTag()
		for _, t := range tags {
			if tag != t {
				return tag, nil
			}
		}
	}
}

func _genReleaseTag() string {
	return "release-" + time.Now().Format("200601021504")
}

package gitwrapper

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

var (
	ErrNoRemoteRepo    = errors.New("no remote repository")
	ErrNoDefaultBranch = errors.New("no default branch")
)

type RepoConfig struct {
	Organization  string
	Repository    string
	Branches      []string
	CurrentBranch string
	DefaultBranch string
}

func NewRepoConfig(dir string) (*RepoConfig, error) {
	if err := os.Chdir(dir); err != nil {
		return nil, err
	}

	var repoConfig = new(RepoConfig)
	var err error

	if repoConfig.Organization, repoConfig.Repository, err = getRemoteInfo(); err != nil {
		return nil, err
	}

	if repoConfig.Branches, repoConfig.CurrentBranch, repoConfig.DefaultBranch, err = getBranchInfo(); err != nil {
		return nil, err
	}

	return repoConfig, nil
}

func getRemoteInfo() (string, string, error) {
	o, err := runGitCmd(false, "git remote -v")
	if err != nil {
		return "", "", err
	}
	r, _ := regexp.Compile(`(?m)^origin\s+(?:ssh://)?(?:git@)?github\.com(?::|/)(.*)/(.*|.*\.git)\s+\(fetch\)$`)
	matches := r.FindStringSubmatch(o)
	if len(matches) == 0 {
		return "", "", ErrNoRemoteRepo
	}
	return matches[1], strings.Split(matches[2], ".git")[0], nil
}

func getBranchInfo() (branches []string, curBranch string, defBranch string, err error) {
	bb, err := getAllBranches()
	if err != nil {
		return
	}

	for _, b := range bb {
		branches = append(branches, b.Name)
		if b.IsCurrent {
			curBranch = b.Name
		}

		if b.Name == "develop" && b.IsRemote {
			defBranch = b.Name
			break
		}
		if b.Name == "master" && b.IsRemote {
			defBranch = b.Name
		}
	}
	if len(defBranch) == 0 {
		err = ErrNoDefaultBranch
		return
	}

	return
}

package gitwrapper

import (
	"errors"
	"os"
	"regexp"
)

var (
	ErrNoRemoteRepo    = errors.New("no remote repository")
	ErrNoCurrentBranch = errors.New("no current branch")
	ErrNoDefaultBranch = errors.New("no default branch")
)

type RepoConfig struct {
	Organization  string
	Repository    string
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

	if repoConfig.CurrentBranch, repoConfig.DefaultBranch, err = getBranchInfo(); err != nil {
		return nil, err
	}

	return repoConfig, nil
}

func getRemoteInfo() (string, string, error) {
	o, err := runGitCmd(false, "remote", "-v")
	if err != nil {
		return "", "", err
	}
	r, _ := regexp.Compile(`git@github\.com:(.*)\/(.*)\.git`)
	matches := r.FindStringSubmatch(o)
	if len(matches) == 0 {
		return "", "", ErrNoRemoteRepo
	}
	return matches[1], matches[2], nil
}

func getBranchInfo() (string, string, error) {
	bb, err := getAllBranches()
	if err != nil {
		return "", "", err
	}

	var cb, db string
	for _, b := range bb {
		if b.IsCurrent {
			cb = b.Name
		}

		if b.Name == "develop" && b.IsRemote {
			db = b.Name
			break
		}
		if b.Name == "master" && b.IsRemote {
			db = b.Name
		}
	}
	if len(db) == 0 {
		return "", "", ErrNoDefaultBranch
	}

	return cb, db, nil
}

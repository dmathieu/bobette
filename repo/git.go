package repo

import (
	"fmt"
	"net/url"
	"os/exec"
)

func pullGit(dir, u, user, password string) error {
	uri, err := url.Parse(u)
	if err != nil {
		return err
	}

	if user != "" && password != "" {
		uri.User = url.UserPassword(user, password)
	}

	gitArgs := []string{"clone"}
	branch := uri.Fragment
	possibleSHA := len(branch) == 40

	if !possibleSHA {
		gitArgs = append(gitArgs, "--depth=1")
		gitArgs = append(gitArgs, "--single-branch")
		if len(branch) != 0 {
			gitArgs = append(gitArgs, fmt.Sprintf("--branch=%s", branch))
		}
	}

	uri.Fragment = ""
	gitArgs = append(gitArgs, uri.String())
	gitArgs = append(gitArgs, ".")

	c := exec.Command("git", gitArgs...)
	c.Dir = dir
	err = c.Run()
	if err != nil {
		return err
	}

	// Checkout the specified branch, which may be a SHA.
	if possibleSHA {
		c = exec.Command("git", "checkout", branch)
		c.Dir = dir
		err := c.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

package repo

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/bgentry/go-netrc/netrc"
)

func pullGit(dir, u, user, password string) error {
	uri, err := url.Parse(u)
	if err != nil {
		return err
	}

	if user != "" && password != "" {
		netrcPath := path.Join(os.Getenv("HOME"), ".netrc")
		n, err := netrc.ParseFile(netrcPath)
		if err != nil {
			return err
		}

		name := strings.Split(uri.Host, ":")[0]
		m := n.FindMachine(name)
		if m == nil {
			n.NewMachine(name, user, password, "")

			t, err := n.MarshalText()
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(netrcPath, t, 0644)
			if err != nil {
				return err
			}
		}
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

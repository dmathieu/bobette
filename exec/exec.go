package exec

import (
	"io"
	"os/exec"
)

// Execute executes the specified commands into the folder
func Execute(dir string, commands []string, out, err io.Writer) error {
	for _, cmd := range commands {
		err := executeSingle(dir, cmd, out, err)
		if err != nil {
			return err
		}
	}
	return nil
}

func executeSingle(dir, command string, out, err io.Writer) error {
	c := exec.Command("sh", "-c", command)
	c.Stdout = out
	c.Stderr = err
	c.Dir = dir
	return c.Run()
}

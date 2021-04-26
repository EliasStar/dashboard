package command

import (
	"bytes"
	"os/exec"
)

type CmdStore map[int]*exec.Cmd

func (c *CmdStore) Start(exe string, args ...string) (int, error) {
	cmd := exec.Command(exe, args...)

	err := cmd.Start()
	if err != nil {
		return 0, err
	}

	(*c)[cmd.Process.Pid] = cmd
	return cmd.Process.Pid, nil
}

func (c *CmdStore) Kill(pid int) error {
	cmd := (*c)[pid]

	if err := cmd.Process.Kill(); err != nil {
		return err
	}

	delete(*c, pid)

	return cmd.Wait()
}

func (c *CmdStore) IsDone(pid int) bool {
	return (*c)[pid].ProcessState.Exited()
}

func (c *CmdStore) Output(pid int) (string, error) {
	cmd := (*c)[pid]
	buf := new(bytes.Buffer)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	delete(*c, pid)

	buf.ReadFrom(stdout)
	buf.ReadFrom(stderr)

	return buf.String(), cmd.Wait()
}

func (c *CmdStore) Find(exe string) (cmds []*exec.Cmd) {
	for _, v := range *c {
		if v.Args[0] == exe {
			cmds = append(cmds, v)
		}
	}

	return
}

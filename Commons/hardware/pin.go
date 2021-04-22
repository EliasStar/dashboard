package hardware

import (
	"fmt"
	"os/exec"
	"time"
)

type Pin uint

func (p Pin) Mode(out bool) error {
	val := "in"
	if out {
		val = "out"
	}

	return exec.Command("gpio", "-g", "mode", fmt.Sprint(p), val).Run()
}

func (p Pin) Write(value bool) error {
	val := "0"
	if value {
		val = "1"
	}

	return exec.Command("gpio", "-g", "write", fmt.Sprint(p), val).Run()
}

func (p Pin) Read() (value bool, err error) {
	out, err := exec.Command("gpio", "-g", "read", fmt.Sprint(p)).Output()
	value = string(out[0]) == "1"
	return
}

func (p Pin) Toggle(delay time.Duration) error {
	val, err := p.Read()
	if err != nil {
		return err
	}

	if err := p.Write(!val); err != nil {
		return err
	}

	time.Sleep(delay)

	return p.Write(val)
}

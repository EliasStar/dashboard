package screen

import (
	"fmt"
	"os/exec"
	"time"
)

type Button uint

const (
	ButtonPower  Button = 17
	ButtonMenu   Button = 27
	ButtonPlus   Button = 22
	ButtonMinus  Button = 23
	ButtonSource Button = 24
)

func Buttons() []Button {
	return []Button{
		ButtonPower,
		ButtonMenu,
		ButtonPlus,
		ButtonMinus,
		ButtonSource,
	}
}

func (b Button) IsValid() bool {
	for _, v := range Buttons() {
		if b == v {
			return true
		}
	}

	return false
}

func (b Button) SetInput() error {
	return exec.Command("gpio", "-g", "mode", fmt.Sprint(b), "in").Run()
}

func (b Button) SetOutput() error {
	return exec.Command("gpio", "-g", "mode", fmt.Sprint(b), "out").Run()
}

func (b Button) Read() (value bool, err error) {
	out, err := exec.Command("gpio", "-g", "read", fmt.Sprint(b)).Output()
	value = string(out[0]) == "1"
	return
}

func (b Button) Write(value bool) error {
	val := "0"
	if value {
		val = "1"
	}

	return exec.Command("gpio", "-g", "write", fmt.Sprint(b), val).Run()
}

func (b Button) Toggle(delay time.Duration) error {
	val, err := b.Read()
	if err != nil {
		return err
	}

	if err := b.Write(!val); err != nil {
		return err
	}

	time.Sleep(delay)

	return b.Write(val)
}

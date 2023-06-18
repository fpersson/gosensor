package syscmd

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/fpersson/gosensor/webservice/model"
)

func ReadLog() (data *model.AllMessages, err error) {
	retval := model.AllMessages{}
	prg := "journalctl"
	arg := "-u"
	arg2 := "tempsensor.service"

	cmd := exec.Command(prg, arg, arg2)

	stdout, err := cmd.Output()

	if err != nil {
		return &retval, err
	}

	v := strings.Split(string(stdout), "\n")
	retval.LogMessages = &v

	return &retval, nil
}

func ReadLogFile(logfile string) (data *model.AllMessages, err error) {
	retval := model.AllMessages{}
	readFile, err := os.Open(logfile)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	buffer := bufio.NewScanner(readFile)
	buffer.Split(bufio.ScanLines)
	var d []string
	for buffer.Scan() {
		d = append(d, buffer.Text())
	}

	retval.LogMessages = &d

	return &retval, nil
}

func ReadStatus() (status *model.SystemdStatus, err error) {
	retval := model.SystemdStatus{}

	prg := "systemctl"
	arg := "status"
	arg2 := "tempsensor.service"

	cmd := exec.Command(prg, arg, arg2)
	stdout, err := cmd.Output()
	if err != nil {
		retval.Active = false
		return &retval, err
	}

	v := strings.Split(string(stdout), "\n")
	if strings.Contains(v[2], "Active: active") {
		retval.Active = true
	} else {
		retval.Active = false
	}
	retval.Message = &v

	return &retval, nil
}

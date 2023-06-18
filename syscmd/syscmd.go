package syscmd

import "os/exec"

const Webservice string = "webserver.service"
const Tempsensorservice string = "tempsensor.service"

func CmdRestart(service string) (err error) {
	prg := "systemctl"
	action := "restart"
	command := exec.Command(prg, action, service)
	stdout, err := command.Output()
	_ = stdout
	return err
}

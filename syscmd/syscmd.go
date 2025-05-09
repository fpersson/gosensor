package syscmd

import "os/exec"

// Webservice represents the name of the webserver systemd service.
const Webservice string = "webserver.service"

// Tempsensorservice represents the name of the temperature sensor systemd service.
const Tempsensorservice string = "tempsensor.service"

// CmdRestart restarts a specified systemd service.
//
// Parameters:
//   - service: The name of the systemd service to restart.
//
// Returns:
//   - error: An error object if the restart command fails.
func CmdRestart(service string) (err error) {
	prg := "systemctl"
	action := "restart"
	command := exec.Command(prg, action, service)
	stdout, err := command.Output()
	_ = stdout
	return err
}

package device

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/ernoaapa/can/pkg/model"
	"github.com/matishsiao/goInfo"
)

// GetInfo resolves information about the device
// Note: Darwin (OSX) implementation is just for development purpose
// For example, BootID get generated every time when process restarts
func (r *Resolver) GetInfo() *model.DeviceInfo {
	osInfo := goInfo.GetInfo()
	ioregOutput := runCommandOrFail("ioreg", "-rd1", "-c", "IOPlatformExpertDevice")

	return &model.DeviceInfo{
		Labels:   r.labels,
		Platform: osInfo.Platform,
		OS:       osInfo.GoOS,
		Kernel:   osInfo.Kernel,
		Core:     osInfo.Core,
		Hostname: osInfo.Hostname,
		CPUs:     osInfo.CPUs,

		MachineID: parseFieldFromIoregOutput(ioregOutput, "IOPlatformSerialNumber"),

		SystemUUID: parseFieldFromIoregOutput(ioregOutput, "IOPlatformUUID"),

		BootID: runCommandOrFail("/usr/bin/uuidgen"),
	}
}

func runCommandOrFail(name string, arg ...string) string {
	bytes, err := exec.Command(name, arg...).Output()
	if err != nil {
		log.Fatalf("Failed to resolve device info: %s", err)
	}
	return strings.TrimSpace(string(bytes))
}

func parseFieldFromIoregOutput(output, field string) string {
	exp := regexp.MustCompile(fmt.Sprintf(".*\"%s\".*\"(.*)\"", field))
	return exp.FindStringSubmatch(output)[1]
}
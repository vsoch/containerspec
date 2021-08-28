package spec

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/archspec/archspec-go/archspec/cpu"
	"github.com/vsoch/containerspec/utils"
)

// Detect the host architecture (this maps to the detect command)
func Detect() cpu.Microarchitecture {

	// Currently just support parsing Linux
	if runtime.GOOS == "linux" {
		//info := getInfoLinux()
		// TODO: we want to map something from info here to cpu.TARGETS

	} else {
		log.Fatal("Currently only Linux is supporting, because Macs and Windows are terrible.")
	}
	return cpu.TARGETS["x86"]
}

// Returns a raw info dictionary by parsing the first entry of /proc/cpuinfo
func getInfoLinux() map[string]string {

	info := make(map[string]string)
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Read through each line and split by :
	for scanner.Scan() {

		// key, separator, value
		values := strings.Split(scanner.Text(), ":")

		// If there's no separator and info was already populated
		// according to what's written here:
		//
		// http://www.linfo.org/proc_cpuinfo.html
		//
		// we are on a blank line separating two cpus. Exit early as
		// we want to read just the first entry in /proc/cpuinfo
		if len(values) == 0 && len(info) > 0 {
			break
		}
		// The key is the first in the list of values
		key := strings.TrimSpace(values[0])

		// If it's empty, just mark as empty?
		var value string
		if len(values) == 1 {
			value = "empty"
		} else {
			value = strings.TrimSpace(values[1])
		}
		info[key] = value

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return info
}

func checkOutput(args []string, env []string) string {
	output := utils.RunCommand(args, env)
	return output
}

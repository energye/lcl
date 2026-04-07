//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux

package version

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseOSRelease() {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		println("unable to read system information")
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		ini := strings.SplitN(line, "=", 2)
		if len(ini) != 2 {
			continue
		}
		name := ini[0]
		value := strings.Trim(ini[1], "\"")
		switch name {
		case "NAME":
			OSVersion.Name = value
		case "VERSION_ID":
			v := strings.Split(value, ".")
			OSVersion.Major, _ = strconv.Atoi(v[0])
			OSVersion.Minor, _ = strconv.Atoi(v[1])
		}
	}
}

func initOSVersion() {
	OSVersion.Platform = PfLinux
	parseOSRelease()
	OSVersion.fmtVerString = fmt.Sprintf("%s (Version %d.%d)", OSVersion.Name, OSVersion.Major, OSVersion.Minor)
}

package main

import (
	"github.com/godbus/dbus"
	log "github.com/sirupsen/logrus"
)

func main() {
	bus, err := dbus.SystemBus()
	if err != nil {
		log.Errorf("could not establish connection to system bus: %v", err)
		return
	}

	obj := bus.Object("org.freedesktop.PackageKit", "/org/freedesktop/PackageKit")
	call := obj.Call(
		"org.freedesktop.PackageKit.StateHasChanged",
		dbus.FlagNoReplyExpected,
		"posttrans")
	if call.Err != nil {
		log.Errorf("could not drop PackageKit cache: %v", call.Err)
		return
	}
}

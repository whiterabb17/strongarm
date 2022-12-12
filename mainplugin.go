package main

import "log"

type Symbol any

type Plugin struct {
	// contains filtered or unexported fields
}

var (
	RestoreTask               bool
	PathToUsernameList        string
	UsernameListRandomization bool
	PathToPasswordList        string
	PasswordListRandomization bool
	Protocol                  string
	IpList                    bool
	Target                    string
	TargetList                []string
	WorkersNumber             int
	TaskStateObj              taskState
)

func PluginService() {
	log.Printf("Path to User: %s\nPath to Pass: %s\nUserRandomization: %v\nPassRandomization: %v\nProtocal: %s\nTarget: %s\nWorkers: %c", PathToUsernameList, PathToPasswordList, UsernameListRandomization, PasswordListRandomization, Protocol, Target, WorkersNumber)
	//beginBrute(PathToUsernameList, PathToPasswordList, UsernameListRandomization, PasswordListRandomization, Protocol, Target, WorkersNumber)
}

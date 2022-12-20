package main

import "flag"

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
	//
	restoreTask               bool
	pathToUsernameList        string
	usernameListRandomization bool
	pathToPasswordList        string
	passwordListRandomization bool
	protocol                  string
	ipList                    bool
	target                    string
	targetList                []string
	workersNumber             int
	taskStateObj              taskState
)

func init() {
	flag.BoolVar(&restoreTask, "restore", false, "Restore task")
	flag.StringVar(&pathToUsernameList, "ul", "usernames.txt", "Path to usernames list")
	flag.StringVar(&pathToPasswordList, "pl", "passwords.txt", "Path to passwords list")
	flag.BoolVar(&usernameListRandomization, "ru", false, "Randomize users list")
	flag.BoolVar(&passwordListRandomization, "rp", false, "Randomize passwords list")
	flag.StringVar(&protocol, "p", "ftp", "Protocol (ftp,ssh,httpbasic,httpdigest,rdp,winldap)")
	flag.StringVar(&target, "t", "10.0.0.1:21", "Target")
	flag.IntVar(&workersNumber, "w", 5, "Number of Workers")
	flag.Parse()
}

func main() {
	flag.Parse()
	BeginBrute(PathToUsernameList, PathToPasswordList, UsernameListRandomization, PasswordListRandomization, Protocol, Target, WorkersNumber)
}

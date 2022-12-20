package main

import (
	"flag"
	"log"

	strongarm "github.com/whiterabb17/strongarm/library"
)

var (
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
	flag.StringVar(&protocol, "p", "ssh", "Protocol (ftp,ssh,httpbasic,httpdigest,rdp,winldap)")
	flag.StringVar(&target, "t", "192.168.96.132:22", "Target")
	flag.IntVar(&workersNumber, "w", 5, "Number of Workers")
	flag.Parse()
}

func main() {
	flag.Parse()
	log.Println(pathToUsernameList)
	log.Println(pathToPasswordList)
	log.Println(protocol)
	log.Println(target)
	strongarm.BeginBrute(pathToUsernameList, pathToPasswordList, usernameListRandomization, passwordListRandomization, protocol, target, workersNumber)
}

package strongarm

import (
	// "os"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
	// "FtpSpray"
)

// PLUGIN FUNCTION
// func PluginService() {
// 	//log.Printf("Path to User: %s\nPath to Pass: %s\nUserRandomization: %v\nPassRandomization: %v\nProtocal: %s\nTarget: %s\nWorkers: %c", PathToUsernameList, PathToPasswordList, UsernameListRandomization, PasswordListRandomization, Protocol, Target, WorkersNumber)
// 	BeginBrute(PathToUsernameList, PathToPasswordList, UsernameListRandomization, PasswordListRandomization, Protocol, Target, WorkersNumber)
// }

func getIPs(path string) []string {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

func BeginBrute(userlist string, passlist string, rndUlist bool, rndPlist bool, proto string, targets string, wrkr int) {
	pathToUsernameList = userlist
	pathToPasswordList = passlist
	usernameListRandomization = rndUlist
	passwordListRandomization = rndPlist
	protocol = proto
	if strings.Contains(targets, ":") {
		ipList = false
		target = targets
	}
	/*
		else {
			_, err := os.Stat(targets)
			if err != nil {
				log.Println("IPList does not exist")
			} else {
				ipList = true
				targetList = getIPs(targets)
			}
		}
	*/
	startTaskService()
}

func printSuccessfulLogin(c chan string) {
	for {
		credentials := <-c
		fmt.Println("\nSuccess: " + credentials)
	}

}

var (
	restoreTask               bool
	pathToUsernameList        string
	usernameListRandomization bool
	pathToPasswordList        string
	passwordListRandomization bool
	protocol                  string
	ipList                    bool
	target                    string
	//targetList                []string
	workersNumber             int
	taskStateObj              taskState
)

type runningTask struct {
	RandomSeed             int64
	UsersList              string
	PasswordsList          string
	ProtocolToSpray        string
	Target                 string
	WorkersCount           int
	WorkersStates          []workerState
	UsernamesRandomization bool
	PasswordsRandomization bool
}

var currentTask runningTask

func startTaskService() { // main() {
	if restoreTask {
		err := readGob("./progress.gob", &currentTask)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		if usernameListRandomization || passwordListRandomization {
			currentTime := time.Now().UnixNano()
			currentTask.RandomSeed = currentTime
			if usernameListRandomization {
				currentTask.UsernamesRandomization = true
			}
			if passwordListRandomization {
				currentTask.PasswordsRandomization = true
			}
		} else {
			currentTask.RandomSeed = 0
			currentTask.PasswordsRandomization = false
			currentTask.UsernamesRandomization = false
		}

		currentTask.UsersList = pathToUsernameList
		currentTask.PasswordsList = pathToPasswordList
		currentTask.ProtocolToSpray = protocol
		currentTask.Target = target
		currentTask.WorkersCount = workersNumber

		for i := 1; i <= workersNumber; i++ {
			currentTask.WorkersStates = append(currentTask.WorkersStates, workerState{
				WorkerId:       i,
				WorkerProgress: 0,
			})
		}

		saveProgress()
	}

	usernames := loadList(currentTask.UsersList)
	passwords := loadList(currentTask.PasswordsList)

	if currentTask.UsernamesRandomization {
		rand.Seed(currentTask.RandomSeed)
		rand.Shuffle(len(usernames), func(i, j int) { usernames[i], usernames[j] = usernames[j], usernames[i] })
	}
	if currentTask.PasswordsRandomization {
		rand.Seed(currentTask.RandomSeed)
		rand.Shuffle(len(passwords), func(i, j int) { passwords[i], passwords[j] = passwords[j], passwords[i] })
	}

	targetToSpray := parseTarget(currentTask.Target)
	wholeTask := task{target: targetToSpray, usernames: usernames, passwords: passwords, numberOfWorkers: currentTask.WorkersCount}
	tasks := dispatchTask(wholeTask)

	var wg sync.WaitGroup
	channelForWorker := make(chan string)
	go printSuccessfulLogin(channelForWorker)
	iter := 0
	if currentTask.ProtocolToSpray == "ftp" {
		for _, task := range tasks {
			wg.Add(1)
			go ftpSpray(&wg, channelForWorker, task, &currentTask.WorkersStates[iter].WorkerProgress)
			iter++
		}
	} else if currentTask.ProtocolToSpray == "ssh" {
		for _, task := range tasks {
			wg.Add(1)
			go sshSpray(&wg, channelForWorker, task, &currentTask.WorkersStates[iter].WorkerProgress)
			iter++
		}
	} else if currentTask.ProtocolToSpray == "httpbasic" {
		for _, task := range tasks {
			wg.Add(1)
			go basicSpray(&wg, channelForWorker, task, &currentTask.WorkersStates[iter].WorkerProgress)
			iter++
		}
	} else if currentTask.ProtocolToSpray == "httpdigest" {
		for _, task := range tasks {
			wg.Add(1)
			go digestSpray(&wg, channelForWorker, task, &currentTask.WorkersStates[iter].WorkerProgress)
			iter++
		}
	} else if currentTask.ProtocolToSpray == "rdp" {

		for _, task := range tasks {
			wg.Add(1)
			go rdpSpray(&wg, channelForWorker, task, &currentTask.WorkersStates[iter].WorkerProgress)
			iter++
		}
	} else if currentTask.ProtocolToSpray == "winldap" {

		for _, task := range tasks {
			wg.Add(1)
			go ldapSpray(&wg, channelForWorker, task, &currentTask.WorkersStates[iter].WorkerProgress)
			iter++
		}
	}
	go monitorCurrentTask()
	wg.Wait()
	close(channelForWorker)
}

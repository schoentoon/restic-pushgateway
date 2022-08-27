package main

import (
	"flag"
	"io"
	"os"

	"github.com/schoentoon/restic-pushgateway/pkg/stats"
	"github.com/sirupsen/logrus"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		logrus.Fatal(err)
	}
	pushGatewayAddress := flag.String("prometheus_push_gateway", "", "Address of the push gateway")
	job := flag.String("job", "restic", "Prometheus job label")
	instance := flag.String("instance", "restic@"+hostname, "Prometheus instance label. Should be as descriptive as possible")
	tee := flag.Bool("tee", false, "Print the json output from restic to stdout")
	flag.Parse()

	if *pushGatewayAddress == "" {
		logrus.Fatal("No pushgateway address specified")
	}

	var in io.Reader = os.Stdin
	if *tee {
		in = io.TeeReader(in, os.Stdout)
	}

	resticStats, err := stats.ParseOutput(in)
	if err != nil {
		logrus.Fatal(err)
	}

	err = resticStats.Push(*pushGatewayAddress, *job, *instance)
	if err != nil {
		logrus.Fatal(err)
	}
}

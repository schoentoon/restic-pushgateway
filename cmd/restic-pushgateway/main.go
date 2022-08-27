package main

import (
	"flag"
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
	flag.Parse()

	if *pushGatewayAddress == "" {
		logrus.Fatal("No pushgateway address specified")
	}

	resticStats, err := stats.ParseOutput(os.Stdin)
	if err != nil {
		logrus.Fatal(err)
	}

	err = resticStats.Push(*pushGatewayAddress, *job, *instance)
	if err != nil {
		logrus.Fatal(err)
	}
}

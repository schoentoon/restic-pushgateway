package stats

import (
	"fmt"
	"reflect"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func (s *Stats) Push(address, job, instance string) error {
	pusher := push.New(address, job).Grouping("instance", instance)

	lastBackup := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: job + "_last_backup",
		Help: "The timestamp of the last backup",
	})
	lastBackup.SetToCurrentTime()
	pusher.Collector(lastBackup)

	value := reflect.ValueOf(*s)
	fields := reflect.VisibleFields(value.Type())

	for _, field := range fields {
		if label, ok := field.Tag.Lookup("push"); ok {
			collector := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: fmt.Sprintf("%s_%s", job, label),
				Help: fmt.Sprintf("Automatically generated from label %s", label),
			})

			property := value.FieldByName(field.Name)
			switch property.Kind() {
			case reflect.Int:
				collector.Set(float64(property.Int()))
			case reflect.Float64:
				collector.Set(property.Float())
			}

			pusher.Collector(collector)
		}
	}

	return pusher.Push()
}

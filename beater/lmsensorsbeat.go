package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/eskibars/lmsensorsbeat/config"
	"github.com/eskibars/gosensors"
)

type Lmsensorsbeat struct {
	beatConfig *config.Config
	done       chan struct{}
	sensors    []gosensors.Feature
	period     time.Duration
}

// Creates beater
func New() *Lmsensorsbeat {
	return &Lmsensorsbeat{
		done: make(chan struct{}),
	}
}

/// *** Beater interface methods ***///

func (bt *Lmsensorsbeat) Config(b *beat.Beat) error {

	// Load beater beatConfig
	err := cfgfile.Read(&bt.beatConfig, "")
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	return nil
}

func (bt *Lmsensorsbeat) Setup(b *beat.Beat) error {

	// Setting default period if not set
	if bt.beatConfig.Lmsensorsbeat.Period == "" {
		bt.beatConfig.Lmsensorsbeat.Period = "1s"
	}

	var err error
	bt.period, err = time.ParseDuration(bt.beatConfig.Lmsensorsbeat.Period)
	if err != nil {
		return err
	}

	gosensors.Init()
	chips := gosensors.GetDetectedChips()
	bt.sensors = nil

	for i := 0; i < len(chips); i++ {
		features := chips[i].GetFeatures()
		for j := 0; j < len(features); j++ {
			bt.sensors = append(bt.sensors, features[j])
		}
	}

	return nil
}

func (bt *Lmsensorsbeat) Run(b *beat.Beat) error {
	logp.Info("lmsensorsbeat is running! Hit CTRL-C to stop it.")

	ticker := time.NewTicker(bt.period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		var sensorMap common.MapStr
		for i := 0; i < len(bt.sensors); i++ {
			sensorMap = common.MapStrUnion(sensorMap, common.MapStr { bt.sensors[i].Name +"("+ bt.sensors[i].GetLabel() +")": bt.sensors[i].GetValue() })
		}
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"sensors":    sensorMap,
		}
		b.Events.PublishEvent(event)
		logp.Info("Event sent")
	}
}

func (bt *Lmsensorsbeat) Cleanup(b *beat.Beat) error {
	gosensors.Cleanup()
	return nil
}

func (bt *Lmsensorsbeat) Stop() {
	close(bt.done)
}

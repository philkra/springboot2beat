package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/philkra/springboot2beat/config"
)

// Springboot2beat configuration.
type Springboot2beat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of springboot2beat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Springboot2beat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts springboot2beat.
func (bt *Springboot2beat) Run(b *beat.Beat) error {
	logp.Info("springboot2beat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
        bt.ProcessMetricsActuator(b)
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
	}
}

// Stop stops springboot2beat.
func (bt *Springboot2beat) Stop() {
	bt.client.Close()
	close(bt.done)
}

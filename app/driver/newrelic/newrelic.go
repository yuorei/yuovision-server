package newrelic

import (
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelic() *newrelic.Application {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("yuovision-server"),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	return app
}

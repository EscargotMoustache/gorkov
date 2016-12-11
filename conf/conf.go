package conf

import (
	"io/ioutil"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Server      string
	Channel     string
	BotName     string
	TLS         bool
	InsecureTLS bool
}

var C = new(Configuration)

func Load(cp string) error {
	conf, err := ioutil.ReadFile(cp)

	if err != nil {
		return fmt.Errorf("Conf: Could not read configuration: %v", err)
	}

	if err = yaml.Unmarshal(conf, &C); err != nil {
		return fmt.Errorf("Conf: Error while parsing yaml: %v", err)
	}

	return nil
}

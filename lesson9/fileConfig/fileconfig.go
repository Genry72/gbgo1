package fileConfig

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
)

type AppConf struct {
	Port        int    ` yaml:"port"`
	DbUrl       myurl  `yaml:"db_url"`
	JaegerUrl   myurl  `yaml:"jaeger_url"`
	SentryUrl   myurl  `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppID   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}
type myurl string

func (ac *myurl) UnmarshalYAML(value *yaml.Node) error {
	u, err := url.Parse(value.Value)
	if err != nil || u.Scheme == "" || u.Host == "" {
		err = fmt.Errorf("не удалось распарсить url: %s", value.Value)
		return err
	}
	*ac = myurl(value.Value)
	return nil
}

func New(confName string) (*AppConf, error) {
	var conf = &AppConf{}
	file, err := os.ReadFile(confName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

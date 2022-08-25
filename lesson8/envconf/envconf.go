package envconf

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"net/url"
)

type AppConf struct {
	Port        int
	DbUrl       myurl  `envconfig:"DB_URL"`
	JaegerUrl   myurl  `envconfig:"JAEGER_URL"`
	SentryUrl   myurl  `envconfig:"SENTRY_URL"`
	KafkaBroker string `envconfig:"KAFKA_BROKER"`
	SomeAppID   string `envconfig:"SOME_APP_ID"`
	SomeAppKey  string `envconfig:"SOME_APP_KEY"`
}
type myurl string

func (ac *myurl) Decode(val string) error {
	u, err := url.Parse(val)
	if err != nil || u.Scheme == "" || u.Host == "" {
		err = fmt.Errorf("не удалось распарсить url: %s", val)
		return err
	}
	return nil
}

func New() (*AppConf, error) {
	var conf = &AppConf{}
	err := envconfig.Process("", conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

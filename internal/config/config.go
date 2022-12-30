package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
	"strings"
	"time"
)

const (
	envPrefix = "translator"
	Default   = "config.defaults.yaml"
)

var C *Config

type Config struct {
	HTTPServer HTTPServer `yaml:"http_server"`
	NatsServer NatsServer `yaml:"nats_server"`
	Mysql      Mysql      `yaml:"mysql"`
	Logger     Logger     `yaml:"logger"`
}

type HTTPServer struct {
	Listen            string        `yaml:"listen"`
	ReadTimeout       time.Duration `yaml:"read_Timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
}

type NatsServer struct {
	Address            string      `yaml:"address"`
	StreamName         string      `yaml:"streamName"`
	StreamSubjects     string      `yaml:"streamSubjects"`
	SubjectNameMessage string      `yaml:"subjectNameMessage"`
	PublicTopic        PublicTopic `yaml:"publicTopic"`
}

type PublicTopic struct {
	MessagePrefix string `yaml:"messageTopicPrefix"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type Logger struct {
	Level string `yaml:"level"`
}

func Init() *Config {
	c := new(Config)
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	v.SetConfigName(Default)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error loading default configs: %s", err.Error()))
	}

	err := v.Unmarshal(c, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		config.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			TimeLocationDecodeHook(),
		)
	})
	if err != nil {
		panic(fmt.Errorf("failed on config unmarshal: %w", err))
	}

	C = c

	return c
}

func TimeLocationDecodeHook() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		var timeLocation *time.Location
		if t != reflect.TypeOf(timeLocation) {
			return data, nil
		}

		return time.LoadLocation(data.(string))
	}
}

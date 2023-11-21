package opt

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type pg struct {
	Host   string `mapstructure:"host"`
	User   string `mapstructure:"user"`
	Passwd string `mapstructure:"passwd"`
	DBName string `mapstructure:"db_name"`
	Port   string `mapstructure:"port"`
}

type access struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type meta struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type data struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type config struct {
	Pg       pg     `mapstructure:"pg"`
	Access   access `mapstructure:"access"`
	Meta     meta   `mapstructure:"meta"`
	Data     data   `mapstructure:"data"`
	LogLevel string `mapstructure:"log_level"`
}

var (
	configFile string
	Cfg        = new(config)
)

func init() {
	flag.Parse()
	flag.StringVar(&configFile, "c", "./etc/config.yaml", "set config file path")
}

func MustInitConfig() {
	flag.Parse()

	var (
		err error
	)

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	if err = viper.ReadInConfig(); err != nil {
		logrus.Panicf("read in config file err: %v", err)
	}

	if err = viper.Unmarshal(Cfg); err != nil {
		logrus.Panicf("unmarshal config file err: %v", err)
	}

	switch strings.ToLower(Cfg.LogLevel) {
	case "err", "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
			CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
				file := path.Base(frame.File)
				return "", " " + file + ":" + strconv.Itoa(frame.Line)
			},
		})
	}

	logrus.Infof("read in config detail: %+v", *Cfg)
}

package service

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type KafkaConsumer struct {
	Brokers  string `yaml:"brokers"`
	Group    string `yaml:"group"`
	Version  string `yaml:"version"`
	Topics   string `yaml:"topics"`
	Assignor string `yaml:"assignor"`
	Oldest   bool   `yaml:"oldest"`
	Verbose  bool   `yaml:"verbose"`
}

type Mysql struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Databases   string `yaml:"databases"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Charset     string `yaml:"charset"`
	MaxOpen     int    `yaml:"maxOpen"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxLifetime int    `yaml:"maxLifetime"`

}

type Option struct {
	KafkaConsumer KafkaConsumer `yaml:"kafkaConsumer"`
	Mysql          Mysql        `yaml:"mysql"`
	RuntimeDir     string		`yaml:"runtime_dir"`
	LoggerFileName string		`yaml:"logger_file_name"`
}

func (o *Option) GetOption() *Option {
	yamlFile, err := ioutil.ReadFile(YamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
		os.Exit(0)
	}

	err = yaml.Unmarshal(yamlFile, o)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
		os.Exit(0)
	}
	return o
}
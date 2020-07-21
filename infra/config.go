package infra

import (
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"os"
)

type config struct {
	Name   string `yaml:"name"`
	Server Server `yaml:"server"`
}

//type Modules struct {
//	Seaweed Seaweed `yaml:"seaweed"`
//	Um      Um      `yaml:"um"`
//}
//
//type Seaweed struct {
//	MasterUrl string `yaml:"masterUrl"`
//}
//
//type Um struct {
//	GrpcAddr string `yaml:"GRPCAddr"`
//}

type Server struct {
	Swagger bool `yaml:"swagger"`
	Port    Port
}

type Port struct {
	HTTP int `yaml:"http"`
	GRPC int `yaml:"grpc"`
}

var (
	Config     config
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config.path", "config.yml", "path of config")
}

func parse() {
	glog.Infof("config path: %s", configPath)

	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	//反序列化到Config
	if err := yaml.NewDecoder(file).Decode(&Config); err != nil {
		panic(err)
	}

	glog.Warning("==============================CONFIG==============================")
	bytes, err := json.MarshalIndent(Config, "", "    ")
	if err != nil {
		panic(err)
	}
	glog.Warningf("config: \n%s", bytes)
	glog.Warning("==============================CONFIG==============================")
}

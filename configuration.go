package treesiplibs

import (
    "io/ioutil"
    "fmt"
    "gopkg.in/yaml.v2"
)

type Conf struct {
    TargetSync   float64  `yaml:"target"`
    Nodes        int    `yaml:"nodes"`
    Timeout      int    `yaml:"timeout"`
    RootNode     int    `yaml:"rootnode"`
    Port         int    `yaml:"port"`
}

func (c *Conf) GetConf( filename string ) *Conf {

    yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
	fmt.Errorf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
	fmt.Errorf("Unmarshal: %v", err)
    }

    return c
}
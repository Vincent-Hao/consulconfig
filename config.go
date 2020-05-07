package main

import (
    "errors"
    "fmt"
    "github.com/sirupsen/logrus"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "path/filepath"
)

type Config interface {
    LoadLocalConfig(path string,filename string) error
    //Validation(key string,value []byte,config interface{}) error
    LoadCenterConfig(key string,src []byte) error
}



type LogConfig struct{
    LogLevel     string `yaml:"level"`              //日志等级
    MaxFileSize  int64  `yaml:"filesize"`           //最大日志文件大小（M）
    BackendName  string `yaml:"backendname"`       //后端名(rpc)
    ServerName  string `yaml:"servername"`         //服务名(service)
    LogField    string `yaml:"logfield"`           //日志打印域控制
}

func (lg *LogConfig) LoadLocalConfig(path string,filename string) error{
    fp := filepath.Join(path,filename)
    if !filepath.IsAbs(fp){
        err := errors.New(fmt.Sprintf("path %s, filename %s is invalid",path,filename))
        logrus.Error(err)
        return err
    }
    rd, err := ioutil.ReadFile(fp)
    if err != nil{
        logrus.Error("read config file %s error: %s",fp,err.Error())
        return err
    }
    
    err = yaml.Unmarshal(rd,lg)
    if err != nil{
        logrus.Error(err)
        return err
    }
    return nil
}

func (lg *LogConfig) LoadCenterConfig(key string,src []byte) error {
    err := yaml.Unmarshal(src,lg)
    if err != nil{
        logrus.Error(err)
        return err
    }
    return nil
}
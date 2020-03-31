package center

import (
    "center/config"
)

var ConfigMap = make(map[string]*Entry)

type Entry struct{
    Key string
    Value []byte
    Config config.Config
    Validation func(key string,value []byte,config interface{}) error
    FileType string
    Change  chan struct{}
}
//每次新建一个entry，建立一个change监控，entry保存到map中
func NewEntry(key string,value []byte,config config.Config,filetype string,change chan struct{}) *Entry{
    return &Entry{key,value,config,nil,
        filetype,make(chan struct{})}
}

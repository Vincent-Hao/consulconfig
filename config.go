package config

type Config interface {
    LoadLocalConfig(path string,filename string) error
    //Validation(key string,value []byte,config interface{}) error
    LoadCenterConfig(key string,src []byte) error
}

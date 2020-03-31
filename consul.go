package center

import (
    "github.com/docker/libkv/store"
    "github.com/docker/libkv/store/consul"
)
func init(){
    consul.Register()
}
type Consul struct{
    consul store.Store
}

func NewConsul(addrs []string, options *store.Config) (*Consul,error){
    con := new(Consul)
    store,err :=consul.New(addrs,options)
    if err != nil{
        return nil,err
    }
    con.consul = store
    return con,nil
}

func (con *Consul) Put(key string, value []byte, options *store.WriteOptions) error{
    return con.consul.Put(key,value,options)
}

func (con *Consul) Get(key string) (*store.KVPair, error){
    return con.consul.Get(key)
}

func (con *Consul) Delete(key string) error{
    return con.consul.Delete(key)
}

func (con *Consul)Exists(key string) (bool, error){
    return con.consul.Exists(key)
}

func (con *Consul) Watch(key string, stopCh <-chan struct{}) (<-chan *store.KVPair, error){
    return con.consul.Watch(key,stopCh)
}

func (con *Consul) Close(){
    con.consul.Close()
}


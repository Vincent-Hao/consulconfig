package center

import (
    "github.com/docker/libkv/store"
    "github.com/docker/libkv/store/zookeeper"
)
func init(){
    zookeeper.Register()
}
type Zookeeper struct{
    zookeeper store.Store
}
func NewZookeeper(addrs []string, options *store.Config) (*Zookeeper,error){
    et := new(Zookeeper)
    store,err :=zookeeper.New(addrs,options)
    if err != nil{
        return nil,err
    }
    et.zookeeper = store
    return et,nil
}
func (zk *Zookeeper) Put(key string, value []byte, options *store.WriteOptions) error{
    return zk.zookeeper.Put(key,value,options)
}

func (zk *Zookeeper) Get(key string) (*store.KVPair, error){
    return zk.zookeeper.Get(key)
}

func (zk *Zookeeper) Delete(key string) error{
    return zk.zookeeper.Delete(key)
}

func (zk *Zookeeper)Exists(key string) (bool, error){
    return zk.zookeeper.Exists(key)
}

func (zk *Zookeeper) Watch(key string, stopCh <-chan struct{}) (<-chan *store.KVPair, error){
    return zk.zookeeper.Watch(key,stopCh)
}

func (zk *Zookeeper) Close(){
    zk.zookeeper.Close()
}
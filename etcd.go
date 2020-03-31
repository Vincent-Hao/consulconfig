package center

import (
    "github.com/docker/libkv/store"
    "github.com/docker/libkv/store/etcd"
)
func init(){
    etcd.Register()
}
type Etcd struct{
    etcd store.Store
}

func NewEtcd(addrs []string, options *store.Config) (*Etcd,error){
    et := new(Etcd)
    store,err :=etcd.New(addrs,options)
    if err != nil{
        return nil,err
    }
    et.etcd = store
    return et,nil
}

func (con *Etcd) Put(key string, value []byte, options *store.WriteOptions) error{
    return con.etcd.Put(key,value,options)
}

func (con *Etcd) Get(key string) (*store.KVPair, error){
    return con.etcd.Get(key)
}

func (con *Etcd) Delete(key string) error{
    return con.etcd.Delete(key)
}

func (con *Etcd)Exists(key string) (bool, error){
    return con.etcd.Exists(key)
}

func (con *Etcd) Watch(key string, stopCh <-chan struct{}) (<-chan *store.KVPair, error){
    return con.etcd.Watch(key,stopCh)
}

func (con *Etcd) Close(){
    con.etcd.Close()
}
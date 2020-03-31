package center

import (
	"center/config"
	"context"
	"fmt"
	"github.com/docker/libkv/store"
	"github.com/sirupsen/logrus"
)

type Center interface {
	// Put a value at the specified key
	Put(key string, value []byte, options *store.WriteOptions) error

	// Get a value given its key
	Get(key string) (*store.KVPair, error)

	// Delete the value at the specified key
	Delete(key string) error

	// Verify if a Key exists in the store
	Exists(key string) (bool, error)

	// Watch for changes on a key
	Watch(key string, stopCh <-chan struct{}) (<-chan *store.KVPair, error)

	// Close the store connection
	Close()
}

func RegisterCenterMonitor(center Center, key string, config config.Config, ctx context.Context, cancel context.CancelFunc,notifyCh chan config.Config) error {
	stopch := make(chan struct{})
	change := make(chan struct{})
	watchchan := make(<-chan *store.KVPair)
	watchchan, err := center.Watch(key, stopch)
	defer close(stopch)
	defer center.Close()
	if err != nil {
		logrus.Error("")
		return err
	}
	for {
		select {
		case kvp := <-watchchan:
			err = config.LoadCenterConfig(kvp.Key, kvp.Value)
			if err != nil {
				logrus.Error("config value is invalid")
				continue
			}
			entry := NewEntry(kvp.Key, kvp.Value, config, "", change)
			if _, ok := ConfigMap[entry.Key]; ok {
				logrus.Info("config is exist")
			}
			ConfigMap[entry.Key] = entry
			notifyCh <- config
			fmt.Println(string(kvp.Value))
		case <-ctx.Done():
			cancel()
		default:
		}
	}

	return nil
}

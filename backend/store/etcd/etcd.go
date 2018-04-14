package etcd

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"

	"github.com/zumak/zumo/backend/store"
	"github.com/zumak/zumo/datatypes"
	"github.com/zumak/zumo/utils/log"
)

func init() {
	store.Register("etcd", New)
}

// New is
func New(emmiter *store.EventEmitter, endpoint string, opt map[string]string) (store.Store, error) {
	d, err := time.ParseDuration(opt["time-out"])
	if err != nil {
		d = 5 * time.Second
	}
	// TODO opt for time out
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(endpoint, ","),
		DialTimeout: d,
	})
	if err != nil {
		// handle error!
		return nil, err
	}

	go watchInit(cli, emmiter)

	//defer cli.Close()

	return &Store{cli}, nil
}

// Store is
type Store struct {
	clientv3.KV
}

func watchInit(cli clientv3.Watcher, emmiter *store.EventEmitter) {
	channelCh := cli.Watch(
		context.Background(),
		"/channels",
		clientv3.WithPrefix(),
	)
	msgCh := cli.Watch(
		context.Background(),
		"/messages",
		clientv3.WithPrefix(),
	)
	log.Debug("watch start")
	for {
		select {
		case res := <-channelCh:
			for _, ev := range res.Events {
				log.Debug("channel changed %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				if ev.Type == clientv3.EventTypePut {
					channel := &datatypes.Channel{}
					err := json.Unmarshal(ev.Kv.Value, channel)
					if err != nil {
						log.Warn("Unmarshal error: %s", err.Error())
						continue
					}
					emmiter.PutChannel <- *channel
				}
				//if ev.Type == clientv3.EventTypeDelete
			}

		case res := <-msgCh:
			for _, ev := range res.Events {
				log.Debug("message appended %s %q", ev.Type, ev.Kv.Key)
				if ev.Type == clientv3.EventTypePut {
					msg := &datatypes.Message{}
					err := json.Unmarshal(ev.Kv.Value, msg)
					if err != nil {
						log.Warn("Unmarshal error: %s", err.Error())
						continue
					}
					arr := strings.SplitN(string(ev.Kv.Key), "/", 4)
					if len(arr) < 4 {
						log.Warn("parse error: connot find channelID from key")
						continue
					}
					channelID := arr[2]
					emmiter.PutMessage <- store.PutMessageEvent{
						ChannelID: channelID,
						Message:   *msg,
					}
				}
			}
		}
	}
}

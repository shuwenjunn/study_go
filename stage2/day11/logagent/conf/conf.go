package conf

import "time"

type AppConfig struct {
	KafkaConfig `ini:"kafka"`
	//TailConfig  `ini:"tail"`
	EtcdConfig `ini:"etcd"`
}

type KafkaConfig struct {
	Addr  []string `ini:"addr"`
}

type EtcdConfig struct {
	Addr    []string      `ini:"addr"`
	Timeout time.Duration `ini:"timeout"`
	EntryKey string        `ini:"entrykey"`
}

//type TailConfig struct {
//	Filename string `ini:"filename"`
//}

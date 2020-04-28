package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"study_go/day11/logagent/conf"
	"study_go/day11/logagent/etcd"
	"study_go/day11/logagent/kafka"
	"study_go/day11/logagent/tail"
	"sync"
	"time"
)

var wg sync.WaitGroup

var appConf conf.AppConfig

func emptyForFunc() {
	defer wg.Done()
	for {
		select {}
	}
}

func main() {
	//0. 加载配置文件
	err := ini.MapTo(&appConf, "./conf/conf.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("配置文件加载成功")
	//1. 初始化kafka
	err = kafka.Init(appConf.KafkaConfig.Addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("kafka初始化成功")

	fmt.Printf("%T \n", appConf.EtcdConfig.Timeout)

	//初始化etcd
	err = etcd.Init(appConf.EtcdConfig.Addr, time.Duration(appConf.EtcdConfig.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	fmt.Println("etcd初始化成功")

	//获取etcd存储的配置
	entryConfig, err := etcd.GetConf(appConf.EtcdConfig.EntryKey)
	if err != nil {
		fmt.Printf("err:%v \n", err)
	}

	//2. 初始化tailf
	err = tail.Init(entryConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("3.tail初始化成功")
	confChan := tail.ExposeNewConfChan()
	wg.Add(1)
	go etcd.WatchNewConf(appConf.EtcdConfig.EntryKey, confChan)

	wg.Wait()
}

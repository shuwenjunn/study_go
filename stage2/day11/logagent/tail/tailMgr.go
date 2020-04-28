package tail

import (
	"fmt"
	"study_go/day11/logagent/etcd"
	"time"
)

var taskMgr *tailMgr

type tailMgr struct {
	logEntry    []*etcd.LogEntry
	newConfChan chan []*etcd.LogEntry
	taskMap     map[string]*tailObj
}

func Init(logEntry []*etcd.LogEntry) (err error) {
	taskMgr = &tailMgr{
		logEntry:    logEntry,
		taskMap:     make(map[string]*tailObj, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}
	for _, logEntryItem := range logEntry {
		var tailObj *tailObj
		tailObj, err = NewTailTask(logEntryItem.Path, logEntryItem.Topic)
		mk := fmt.Sprintf("%s_%s", logEntryItem.Path, logEntryItem.Topic)
		taskMgr.taskMap[mk] = tailObj
		if err != nil {
			return
		}
	}
	go taskMgr.run()
	return
}

func (t *tailMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 新增的
					tailObj, _ := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[mk] = tailObj
				}
			}
			// 找出原来t.logEntry有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry { // 从原来的配置中依次拿出配置项
				isDelete := true
				for _, c2 := range newConf { // 去新的配置中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					//t.tskMap[mk] ==> tailObj
					t.taskMap[mk].cancelFunc()
				}
			}
			// 2. 配置删除
			fmt.Println("新的配置来了！", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

//暴露newConfChan
func ExposeNewConfChan() (confChanchan chan []*etcd.LogEntry) {
	return taskMgr.newConfChan
}

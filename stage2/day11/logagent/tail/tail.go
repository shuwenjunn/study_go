package tail

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"study_go/day11/logagent/kafka"
)

type tailObj struct {
	path     string
	topic    string
	instance *tail.Tail

	// 为了能实现退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// Init 初始化tail
func NewTailTask(path string, topic string) (tailTask *tailObj, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	tailTask = &tailObj{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	err = tailTask.init()
	return
}

func (t *tailObj) init() (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	tailClient, err := tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	t.instance = tailClient
	go t.run()
	return
}

func (t *tailObj) run() {
	for {
		select {
		case line := <-t.instance.Lines:
			fmt.Printf("%#v \n", line)
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}

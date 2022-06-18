package tail

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/hpcloud/tail"
	"log-agent/internal/util"
)

type TailHandle struct {
	Filename string `json:"filename"`
	Date     string `json:"date"`
	Tail     *tail.Tail
	config   *Config
}

func NewTailHandle(conf *Config) (*TailHandle, error) {
	handle := &TailHandle{config: conf}
	if err := handle.init(); err != nil {
		return nil, err
	}

	return handle, nil
}

func (t *TailHandle) init() error {
	t.Filename = t.toFilename()
	t.Date = util.Date()

	// 配置信息
	tailConfig := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}

	handle, err := tail.TailFile(t.Filename, tailConfig)
	if err != nil {
		return err
	}

	t.Tail = handle
	return nil
}

func (t *TailHandle) Run(ctx context.Context, group *sync.WaitGroup) {
	go func() {
		defer group.Done()

		// 1. 读取日志
		for {
			select {
			case <-ctx.Done():
				log.Printf("Tail Handle Close: %s\n", t.Filename)
				return
			case line := <-t.Tail.Lines:
				if err := t.write(line); err != nil {
					log.Printf("Tail Handle Error: %s\n", err)
				}
			default:
				time.Sleep(time.Second)
				if util.Date() != t.Date {
					if err := t.Tail.Stop(); err != nil {
						log.Printf("Tail Stop Error: %s\n", err)
						continue
					}

					if err := t.init(); err != nil {
						log.Printf("Tail Handle Error: %s\n", err)
					}
				}
			}
		}
	}()
}

func (t *TailHandle) write(line *tail.Line) error {
	var (
		data interface{}
		err  error
	)

	// 查找格式化方法
	if format, ok := formatHandles[t.config.Type]; ok {
		// 获取到数据
		data, err = format(line.Text, t.config)
		if err != nil {
			return err
		}
	} else {
		data = line.Text
	}

	// 写入数据
	if driver, ok := logDriver[t.config.Driver]; ok {
		return driver(data, t.config)
	} else {
		// 写入数据
		fmt.Printf("data: %v\n", data)
		return nil
	}
}

func (t *TailHandle) formatJson(text string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(text), &data)
	return data, err
}

func (t *TailHandle) toFilename() string {
	filename := t.config.Filename
	if t.config.Rule {
		filename = time.Now().Format(t.config.Filename)
	}

	return fmt.Sprintf("%s/%s", strings.TrimRight(t.config.Path, "/"), strings.TrimLeft(filename, "/"))
}

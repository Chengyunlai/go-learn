package repository

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"sync"
	"top.chengyunlai/go-learn/2-day/4-topic/model"
)

var (
	// key 为 int；value 为 结构体类型
	topicIndexMap map[int64]*model.Topic

	topicDao *TopicDao
	// 适合高并发，只执行一次的场景，主要是用于返回一个实例，即单例模式
	topicOnce sync.Once
)

// TODO：换成数据库查询
func InitTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "test.json")
	if err != nil {
		log.Println("文件打开错误")
		return err
	}
	// 关闭文件流是一个好习惯
	defer open.Close()
	res, err := io.ReadAll(open)
	log.Println(res)
	if err != nil {
		log.Println("读取文件内容错误")
		return err
	}
	var topic model.Topic
	if err := json.Unmarshal(res, &topic); err != nil {
		log.Println("转换JSON时发生错误")
		return err
	}
	topicTmpMap := make(map[int64]*model.Topic)
	topicTmpMap[topic.Id] = &topic
	topicIndexMap = topicTmpMap
	return nil
}

type TopicDao struct {
}

func NewTopicDapInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) (*model.Topic, error) {
	res := topicIndexMap[id]
	if res != nil {
		return res, nil
	}
	return nil, errors.New("没有找到对应的内容")
}

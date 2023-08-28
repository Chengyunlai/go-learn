package repository

import (
	"fmt"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	errTopic := InitTopicIndexMap("../")
	if errTopic != nil {
		log.Fatal(errTopic)
	}
	errPost := InitPostIndexMap("../")
	if errPost != nil {
		log.Fatal(errPost)
	}
	m.Run()
	// 测试后资源释放等收尾工作
}

func TestInitTopicIndexMap(t *testing.T) {
	err := InitTopicIndexMap("../")
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range topicIndexMap {
		fmt.Println(key, *value)
	}
}

func TestInitPostIndexMap(t *testing.T) {
	err := InitPostIndexMap("../")
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range postIndexMap {
		fmt.Println(key, *value)
	}
}

func TestTopicDaoQueryTopicById(t *testing.T) {
	//dao := TopicDao{}
	// 使用单例模式
	dao := NewTopicDapInstance()
	topic, _ := dao.QueryTopicById(1)
	log.Println(topic)
}

func TestPostDaoQueryTopicById(t *testing.T) {
	//dao := TopicDao{}
	// 使用单例模式
	dao := NewPostInstance()
	post, _ := dao.QueryPostById(1)
	log.Println(post)
}

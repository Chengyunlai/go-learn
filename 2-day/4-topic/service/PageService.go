package service

import (
	"errors"
	"log"
	"sync"
	"top.chengyunlai/go-learn/2-day/4-topic/model"
	"top.chengyunlai/go-learn/2-day/4-topic/repository"
)

type PageInfo struct {
	Topic    *model.Topic
	PostList *model.ListPosts
}

type QueryPageInfoFlow struct {
	// 话题的id
	topicId int64
	// 该话题对应的主题内容：Topic以及评论
	pageInfo *PageInfo
}

// 执行do操作返回PageInfo
func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	// 初始化数据
	errTopic := repository.InitTopicIndexMap("../")
	if errTopic != nil {
		log.Fatal(errTopic)
	}
	errPost := repository.InitPostIndexMap("../")
	if errPost != nil {
		log.Fatal(errPost)
	}

	// 检查查询的话题id
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	// 检查pageInfo的封装是否有问题，没有问题即封装PageInfo
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	// 如果检查都没问题，将PageInfo封装并返回
	return f.pageInfo, nil
}

// 检查查询的话题id是否有问题
func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

// 检查pageInfo中的Topic信息以及Post信息是否，并且封装
func (f *QueryPageInfoFlow) prepareInfo() error {

	var wg sync.WaitGroup
	wg.Add(2)
	var topicErr, postErr error
	//获取topic信息
	go func() {
		defer wg.Done()
		topic, err := repository.NewTopicDapInstance().QueryTopicById(f.topicId)
		if err != nil {
			topicErr = err
			return
		}
		f.pageInfo.Topic = topic
	}()
	//获取post列表
	go func() {
		defer wg.Done()
		posts, err := repository.NewPostInstance().QueryPostById(f.topicId)
		if err != nil {
			postErr = err
			return
		}
		f.pageInfo.PostList = posts
	}()
	wg.Wait()
	if topicErr != nil {
		return topicErr
	}
	if postErr != nil {
		return postErr
	}
	return nil
}

func QueryPageInfo(id int64) (*PageInfo, error) {
	flow := QueryPageInfoFlow{topicId: id, pageInfo: &PageInfo{&model.Topic{0, "", "", ""}, &model.ListPosts{0, nil}}}
	do, err := flow.Do()
	if err != nil {
		return nil, errors.New("error")
	}
	return do, nil
}

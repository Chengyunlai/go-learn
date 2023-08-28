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
	postIndexMap map[int64]*model.ListPosts
	postDao      *PostDao
	postOnce     sync.Once
)

// TODO：换成数据库查询
func InitPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "Post.json")
	if err != nil {
		log.Println("文件打开错误")
		return err
	}
	// 关闭文件流是一个好习惯
	defer open.Close()
	res, err := io.ReadAll(open)
	if err != nil {
		log.Println("读取文件内容错误")
		return err
	}
	var post model.ListPosts
	if err := json.Unmarshal(res, &post); err != nil {
		log.Println("转换JSON时发生错误")
		return err
	}
	log.Println("封装post", post)
	postTmpMap := make(map[int64]*model.ListPosts)
	postTmpMap[post.TopicId] = &post
	postIndexMap = postTmpMap
	return nil
}

type PostDao struct {
}

func (*PostDao) QueryPostById(id int64) (*model.ListPosts, error) {
	res := postIndexMap[id]
	if res != nil {
		return res, nil
	}
	return nil, errors.New("没有找到对应的内容")
}

func NewPostInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

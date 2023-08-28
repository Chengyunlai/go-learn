package res

import "top.chengyunlai/go-learn/2-day/4-topic/service"

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Error() *R {
	return &R{Code: 201, Msg: "error", Data: nil}
}

func Success(info *service.PageInfo) *R {
	return &R{Code: 200, Msg: "success", Data: info}
}

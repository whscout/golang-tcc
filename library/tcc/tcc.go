package tcc

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"tcc_example/library/tcc/runtime"
	"tcc_example/library/utils"
	"time"
)

type tcc interface{
	// params：请求参数
	// api：根据当前请求，从配置文件中获取的Try的URL信息
	// 返回值：1、尝试过程中，成功的步骤  2、错误信息
	Try(params []byte, api *RuntimeApi) ([]*runtime.SuccessStep, error)
	// params：请求参数
	// api：根据当前请求，从配置文件中获取的Confirm的URL信息
	// 返回值：1、错误信息
	Confirm(params []byte, api *RuntimeApi, ) error
	// params：请求参数
	// api：根据当前请求，从配置文件中获取的Cancel的URL信息
	// nodes：Try时可能成功的步骤，即需要回滚的步骤（根据Try返回值封装生成）
	// 返回值：1、执行取消时，失败步骤的ID编号集合  2、错误信息
	Cancel(params []byte, api *RuntimeApi) error
}

type DefaultTCC struct {
	httpCli *http.Client
	l *zap.Logger
}

func NewDefaultTCC() *DefaultTCC {
	httpCli := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second*30,
	}
	return &DefaultTCC{httpCli: httpCli}
}

func (dc *DefaultTCC) Try(params []byte, api *RuntimeApi) ([]*runtime.SuccessStep, error) {
	// 获取tcc urls
	success := make([]*runtime.SuccessStep, 0)
	for _, node := range api.Nodes {
		req, err := http.NewRequest(node.Try.Method, node.Try.Url, bytes.NewBuffer(params))
		if err != nil {
			dc.l.Error("new http request err", zap.Error(err))
			return nil, err
		}
		resp, err := dc.httpReq(req)
		if err != nil {
			dc.l.Error("http request err", zap.Error(err), zap.Any("data", api.Nodes))
			return nil, err
		}
		res := utils.Response{}
		if err := json.Unmarshal(resp, &res); err != nil {
			dc.l.Error("unmarshal err", zap.Error(err))
			return nil, err
		}
		// TODO 请求结果判断

		// 不管成功与否（主要为了防止：当服务方接收并处理成功，但返回时失败），将结果保存起来，以备使用
		// 如果插入失败，则直接返回，并在后续回滚之前的步骤
		ss := &runtime.SuccessStep{
			RequestId: api.RequestInfo.Id,
			Index:     node.Index,
			Url:       node.Try.Url,
			Method:    node.Try.Method,
			Param:     string(api.RequestInfo.Param),
			Result:    string(resp),
			Status:    0,
		}
		success = append(success, ss)
	}
	return success, nil
}

func (dc *DefaultTCC) Confirm(params []byte, api *RuntimeApi) error {

	return nil
}

func (dc *DefaultTCC) Cancel(params []byte, api *RuntimeApi) error {
	return nil
}

func (dc *DefaultTCC) httpReq(req *http.Request) ([]byte, error) {
	resp, err := dc.httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}






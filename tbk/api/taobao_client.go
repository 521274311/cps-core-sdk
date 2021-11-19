package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sdk-core/tbk/api/r"
	"sdk-core/utils"
	"sort"
	"strings"
	"time"
)

// TaobaoClient 淘宝联盟客户端
type TaobaoClient interface {
	// Execute 接口调用
	Execute(req interface{}, resp interface{}) error
}

// NewDefaultTaobaoClient 创建默认淘宝联盟客户端
func NewDefaultTaobaoClient(appKey, secret string, timeout time.Duration) TaobaoClient {
	return NewDefaultTaobaoClientSession(appKey, secret, "", timeout)
}

// NewDefaultTaobaoClientSession 创建默认淘宝联盟客户端(含session)
// todo session功能暂未实现
func NewDefaultTaobaoClientSession(appKey, secret, session string, timeout time.Duration) TaobaoClient {
	return &defaultTaobaoClient{
		url:     "http://gw.api.taobao.com/router/rest",
		appKey:  appKey,
		secret:  secret,
		session: session,
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        2000,
				MaxIdleConnsPerHost: 2000,
				IdleConnTimeout:     90 * time.Second,
			},
			Timeout: timeout,
		},
	}
}

type defaultTaobaoClient struct {
	url     string
	appKey  string
	secret  string
	session string
	client  *http.Client
}

func (c *defaultTaobaoClient) Execute(p interface{}, res interface{}) error {
	// 对象转map统一处理
	m, err := json.Marshal(p)
	if err != nil {
		return err
	}
	mp := make(map[string]interface{})
	err = json.Unmarshal(m, &mp)
	if err != nil {
		return err
	}
	// 绑定方法
	mp["method"] = getMethodName(p)
	// 注入公共参数
	c.injectCommonReq(mp)
	// 注入加密参数
	c.injectMd5Sign(mp)
	// 发起http请求
	body, err := c.invoke(mp)
	if err != nil {
		return err
	}
	// 结果注入
	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}
	return nil
}

func (c *defaultTaobaoClient) invoke(params map[string]interface{}) ([]byte, error) {
	reqUrl := c.url + "?"
	for k, v := range params {
		reqUrl += k + "=" + url.QueryEscape(utils.ToString(v)) + "&"
	}
	req, err := http.NewRequest("POST", reqUrl, nil)
	// 设置长连接
	req.Header.Set("Connection", "keep-alive")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *defaultTaobaoClient) injectCommonReq(p map[string]interface{}) {
	p["app_key"] = c.appKey
	p["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	p["v"] = "2.0"
	p["format"] = "json"
	p["simplify"] = false
}

func (c *defaultTaobaoClient) injectMd5Sign(p map[string]interface{}) {
	p["sign_method"] = "md5"
	keys := make([]string, len(p))
	idx := 0
	for k := range p {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	pf := c.secret
	for _, v := range keys {
		pf += v + utils.ToString(p[v])
	}
	pf += c.secret
	p["sign"] = strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(pf))))
}

func getMethodName(req interface{}) string {
	switch req.(type) {
	case r.TbkDgMaterialOptionalRequest:
		return "taobao.tbk.dg.material.optional"
	}
	return ""
}

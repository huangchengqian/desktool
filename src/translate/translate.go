package translate

import (
	"crypto/md5"
	"desktool/src/config"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
)

type cons struct {
	TransResult []struct {
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

func Translate(word string, from string, to string) string {
	res := get(word, from, to)
	var dstCons cons
	// 解析结果集
	err := json.Unmarshal(res, &dstCons)
	if err != nil {
		panic("解析失败")
	}
	return dstCons.TransResult[0].Dst
}

func get(word string, from string, to string) []byte {
	config := config.GetConf()
	parseURL, err := url.Parse(config.Trans.TransApi)
	if err != nil {
		panic("wrong url")
	}

	params := url.Values{}
	params.Set("q", word)
	params.Set("from", from)
	params.Set("to", to)
	params.Set("appid", config.Trans.AppId)
	salt := generateSalt()
	params.Set("salt", salt)
	params.Set("sign", generateSign(word, salt))
	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	resp, err := http.Get(urlPathWithParams)
	if err != nil {
		panic(err)
	}
    defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
    return body
}

func generateSalt() string {
	salt := rand.Int63()
	return fmt.Sprint(salt)
} 

func generateSign(ques string, salt string) string {
	config := config.GetConf()
	aimStr := fmt.Sprintf("%s%s%s%s", config.Trans.AppId, ques, salt, config.Trans.Secret)
	md5Stream := md5.Sum([]byte(aimStr))
	return fmt.Sprintf("%x", md5Stream)
}
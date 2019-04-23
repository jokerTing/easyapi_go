// Copyright 2019 Haochen Ding. All rights reserved.

package easyapi

import (
	"fmt"
	//"os"
	
	"strconv"
	"strings"
	"bytes"
	
	"crypto/hmac"
	"crypto/sha1"
	"crypto/md5"
	
	"encoding/hex"
	"encoding/json"
	
	"net/http"
	"net/url"
	"io/ioutil"
	
	"time"
	"sort"
	//"reflect"
)

// EasyApi specifies a tool using EasyOps CMDB api
type Easyapi struct {
    cmdbAddr string
    ak string
    sk string
    header map[string]string
}

// NewEasyapi create a new EasyApi
func NewEasyapi(cmdbAddr, ak, sk string) *Easyapi {
	fmt.Println("Create a new EasyApi", cmdbAddr, ak, sk)
	header := map[string]string {"Host":"openapi.easyops-only.com", "Content-Type":"application/json;charset=UTF-8"}
	return &Easyapi{cmdbAddr, ak, sk, header}
}

// SendRequest to EasyOps OpenApi
func (ez *Easyapi) SendRequest(reqUrl string, method string, params map[string]interface{}) (string, bool) {
	var isSuccess = false
	var ret = ""
	// timestamp
	var nowTS = strconv.FormatInt(time.Now().Unix(), 10)
	//nowTS = "1546957559"
	method = strings.ToUpper(method)
	var sign = ez.genSignature(reqUrl, method, params, nowTS)
	
	var requestUrl = "http://" + ez.cmdbAddr + reqUrl
	name := url.Values{"accesskey": {ez.ak}, "expires": {nowTS}, "signature": {sign}}
	param := name.Encode()
	requestUrl = fmt.Sprintf("%s?%s", requestUrl, param)
	//requestUrl = requestUrl + "?accesskey=" + ez.ak + "&expires=" + nowTS + "&signature=" + sign
	client := http.Client{}
	if method == "GET" || method == "DELETE"{
		if params != nil{
			for k, v := range params {
                requestUrl = requestUrl + "&" + k + "=" + TransInterfaceToString(v)
			}
		}
		req, err := http.NewRequest(method, requestUrl, nil)
		if err != nil {fmt.Println("[Fatal error] ", err.Error());return ret, isSuccess}
		for k, v := range ez.header {
    		req.Header.Set(k, v)
		}
		//  the Host header is promoted to the Request.Host field and removed from the Header map.
        req.Host = "openapi.easyops-only.com"
		fmt.Println("[Request] ", req)
		response, err := client.Do(req)
		if err != nil {fmt.Println("[Fatal error] ", err.Error());return ret, isSuccess}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
    	if err != nil {fmt.Println("[Fatal error] ", err.Error());return ret, isSuccess}
    	//fmt.Println(reflect.TypeOf(body))
		isSuccess = true
		ret = string(body)
		
	} else if method == "POST" || method == "PUT"{
		bytesData, err := json.Marshal(params)
		if err != nil {fmt.Println("[Fatal error] ", err.Error());return ret, isSuccess}
		reader := bytes.NewReader(bytesData)
		
		req, err := http.NewRequest(method, requestUrl, reader)
		if err != nil {fmt.Println("[Fatal error] ", err.Error());return ret, isSuccess}
		for k, v := range ez.header {
    		req.Header.Set(k, v)
		}
		//  the Host header is promoted to the Request.Host field and removed from the Header map.
        req.Host = "openapi.easyops-only.com"
		
		fmt.Println("[Request] ", req)
		response, err := client.Do(req)
		if err != nil {fmt.Println("[Fatal error] ", err.Error());return ret, isSuccess}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
    	if err != nil {fmt.Println("[Fatal error] ", err.Error());return ret, isSuccess}
    	
    	//fmt.Println(reflect.TypeOf(body))
		isSuccess = true
		ret = string(body)
	} else {
		fmt.Println(method)
		panic("Request method not known")
	}
	
	return ret, isSuccess
}

// hmacSHA1Encrypt encrypt the encryptText use encryptKey
func (ez *Easyapi) hmacSHA1Encrypt(encryptText, encryptKey string) string {
	key := []byte(encryptKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(encryptText))
	var str string = hex.EncodeToString(mac.Sum(nil))
	//fmt.Printf("[encrypt result] %v\n", str)
	return str
}

// genSignature generate the Signature using in easyosp api request
/*
Signature = HMAC-SHA1('SecretKey', UTF-8-Encoding-Of( StringToSign ) ) );
StringToSign = HTTP-Verb + "\n" +
               URL + "\n" +
               Parameters + "\n" +
               Content-Type + "\n" +
               Content-MD5 + "\n" +
               Date + "\n" +
               AccessKey;
*/
func (ez *Easyapi) genSignature(url string, method string, params map[string]interface{}, nowTS string) string {
	
	//fmt.Println(reflect.ValueOf(nowTS))
	var urlParams string = ""
	var bodyContent string = ""
	if params != nil{
		// method is GET or DELETE , params build  the url_params
		method = strings.ToUpper(method)
		if method == "GET" || method == "DELETE"{
			// sort the params
			keys := SortMapByKey(params)
            for _, k := range keys {
            	urlParams = urlParams + k + TransInterfaceToString(params[k])
            }
        // method is POST or PUT, params build the bodyContent
		}else if method == "POST" || method == "PUT"{
			jsonStr, err := json.Marshal(params)
			//fmt.Println(string(jsonStr))
			//fmt.Println(reflect.TypeOf(jsonStr))
			if err != nil {panic(err)}
			md5Ctx := md5.New()
            md5Ctx.Write(jsonStr)
            cipherStr := md5Ctx.Sum(nil)
            bodyContent = hex.EncodeToString(cipherStr)
		} else {
			fmt.Println(method)
			panic("Request method not known")
		}
	}
	
	// HTTP-Verb + "\n" +URL + "\n" +Parameters + "\n" +Content-Type + "\n" +Content-MD5 + "\n" +Date + "\n" +AccessKey;
	var str_sign = method+"\n"+url+"\n"+urlParams+"\n"+ez.header["Content-Type"]+"\n"+bodyContent+"\n"+nowTS+"\n"+ez.ak
	/*
	fmt.Println("-------------------------------")
	fmt.Println("before encrypt:\n"+str_sign)
	fmt.Println("-------------------------------")*/
	var sign = ez.hmacSHA1Encrypt(str_sign, ez.sk)
	return sign
}

// SortMapByKey sort the key of map and return the sorted key slice
func SortMapByKey(OriMap map[string]interface{}) []string {
    keys := make([]string, len(OriMap))
    i := 0
    for k, _ := range OriMap {
        keys[i] = k
        i++
    }
    sort.Strings(keys)
    return keys
}

// TransInterfaceToString translate the interface{} to string
func TransInterfaceToString(i interface{}) string {
	var ret string
	switch i.(type) {
	case string:
		ret = i.(string)
		break
	case int:
    	ret = strconv.FormatInt(int64(i.(int)), 10)	
    	break
	default:
    	fmt.Println("params type not supported", i)
        panic("params type not supported")
	}
	return ret
}

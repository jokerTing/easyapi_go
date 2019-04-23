# easyapi_go
EasyOps OpenAPI SDK by Golang

1.结构说明
easyapi包为核心包，main中为使用示例


2.SDK介绍

函数名	|输入	|输出	|说明
--|--|---|--
NewEasyapi()|cmdbAddr string：cmdb地址<br>ak sk string：API用户的key值|*Easyapi 返回一个EasyApi对象实例|公有方法，初始化SDK实例
(ez *Easyapi) GetAllInstance|objectId string, CMDB模型ID<br>params map[string]interface{}, 请求参数（参考POST /object/@object_id/instance/_search的参数）<br>pagesize int， 一页实例数(小于2000)|[]map[string]interface{}, bool<br>第一个：所有实例的列表<br>第二个：成功标志位|公有方法，获取指定模型的所有实例，使用实例搜索方法接口
(ez *Easyapi) ChangeListToMap	|srcList []map[string]interface{} GetAllInstance方法返回的实例列表<br>keys []string 用作唯一键的属性名，空的话会选择isntanceId，多于两个会用&#124拼接|map[string]interface{}, bool<br>第一个：以keys值为键的字典<br>第二个：成功标志位|公有方法，把字典列表变成字典（选定键值）
(ez *Easyapi) SendRequest|reqUrl string, 请求URL路径<br>method string, 请求方法GET DELETE POST PUT<br>params map[string]interface{} http请求参数|string, bool<br>第一个：http请求返回结果<br>第二个：成功标志位|公有方法，通用的http请求函数

3.使用说明
例子1：

POST请求

接口 /cmdb_resource/object/HOST/instance/_search

参数 

{

    “fields”:{"ip": 1, "hostname": 1, "APP.name": 1},

    "page_size":30,

     "page":1

}

var easyApi = easyapi.NewEasyapi("10.2.148.180", "118a03c99886a3d51af8f5c6", "171cd4a4b5d5d2051e9de97bdf15e908d22f1d55fa6656f91f48ffe4a8d4f5eb")

 // POST requestfmt.Println("-------POST EXAMPLE-------")

fieldInPostData := map[string]int{"ip": 1, "hostname": 1, "APP.name": 1}

postData :=  map[string]interface{} {"page_size":30, "page":1}

postData["fields"] = fieldInPostData

fmt.Println("[DATA]: ",postData)

ret ,isSuccess := easyApi.SendRequest("/cmdb_resource/object/HOST/instance/_search", "POST", postData)

fmt.Println("[Result]", ret, isSuccess)

例子2:

GET请求

接口/cmdb_resource/object/APP/instance/5c1064fd42f06

var easyApi = easyapi.NewEasyapi("10.2.148.180", "118a03c99886a3d51af8f5c6", "171cd4a4b5d5d2051e9de97bdf15e908d22f1d55fa6656f91f48ffe4a8d4f5eb")

fmt.Println("-------GET EXAMPLE-------")

getData :=  map[string]interface{} {}

fmt.Println("[DATA]: ",getData)

ret ,isSuccess = easyApi.SendRequest("/cmdb_resource/object/APP/instance/5c1064fd42f06", "GET", getData)

fmt.Println("[Result]", ret, isSuccess)



例子3:

获取主机的所有实例，实例返回instanceId、ip、hostname属性

并把列表构件成字典，键值是ip和hostname值的组合

var easyApi = easyapi.NewEasyapi("10.2.148.180", "118a03c99886a3d51af8f5c6", "171cd4a4b5d5d2051e9de97bdf15e908d22f1d55fa6656f91f48ffe4a8d4f5eb")

fmt.Println("-------GetAllInstance-------")

fieldInPostData := map[string]int{"instanceId": 1, "ip": 1, "hostname": 1}

postData := map[string]interface{}{"page_size": 3000, "page": 1}

postData["fields"] = fieldInPostData

ret, isSuccess := easyApi.GetAllInstance("HOST", postData, 300)

fmt.Println("[Result]", ret, isSuccess)

retMap, isSuccess := easyApi.ChangeListToMap(ret, []string{"ip", "hostname"})

fmt.Println("[Result Map]", retMap, isSuccess)


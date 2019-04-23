package main

import (
    "fmt"
    "../easyapi"
)

func main(){
	var easyApi = easyapi.NewEasyapi("here is ip", "here is ak", "here is sk")
    // POST request
    fmt.Println("-------POST EXAMPLE-------")
    fieldInPostData := map[string]int{"ip": 1, "hostname": 1, "owner.name": 1, "VMHOSTHARDWAREBASIC.name": 1,
        "_deviceList_CLUSTER.clusterId": 1, "_deviceList_CLUSTER.name": 1}
    postData :=  map[string]interface{} {"page_size":1, "page":1}
    postData["fields"] = fieldInPostData
    fmt.Println("[DATA]: ",postData)
    ret ,isSuccess := easyApi.SendRequest("/cmdb/object/HOST/instance/_search", "POST", postData)
    fmt.Println("[Result]", ret, isSuccess)

    fmt.Println("-------POST EXAMPLE-------")
    fieldInPostData2 := map[string]int{"name": 1, "owner.name": 1, "developer.name": 1, "tester.name": 1,
        "businesses.name": 1, "clusters.deviceList.hostname": 1}
    postData2 :=  map[string]interface{} {"page_size":2000, "page":1}
    postData2["fields"] = fieldInPostData2
    fmt.Println("[DATA]: ",postData2)
    ret2 ,isSuccess2 := easyApi.SendRequest("/cmdb/object/APP/instance/_search", "POST", postData2)
    fmt.Println("[Result]", ret2, isSuccess2)

    fmt.Println("-------POST EXAMPLE-------")
    fieldInPostData3 := map[string]int{"name": 1, "op_admin.name": 1, "dev_admin.name": 1}
    postData3 :=  map[string]interface{} {"page_size":2000, "page":1}
    postData3["fields"] = fieldInPostData3
    fmt.Println("[DATA]: ",postData3)
    ret3 ,isSuccess3 := easyApi.SendRequest("/cmdb/object/BUSINESS/instance/_search", "POST", postData3)
    fmt.Println("[Result]", ret3, isSuccess3)

    fmt.Println("-------POST EXAMPLE-------")
    fieldInPostData4 := map[string]int{"_deviceList_CLUSTER.name": 1, "owner.name": 1, "instanceId": 1}
    postData4 :=  map[string]interface{}{"page_size" : 2000, "page" : 1}
    var array4 = [1]string{"10.2.148.181"}  
    fieldInPostData41 := map[string][1]string{"$in": array4}
    fieldInPostData42 := map[string]map[string][1]string{"ip": fieldInPostData41}
    postData4["fields"] = fieldInPostData4
    postData4["query"] = fieldInPostData42
    fmt.Println("[DATA]: ",postData4)
    ret4 ,isSuccess4 := easyApi.SendRequest("/cmdb/object/HOST/instance/_search", "POST", postData4)
    fmt.Println("[Result]", ret4, isSuccess4) 

    fmt.Println("-------POST EXAMPLE-------")
    fieldInPostData5 := map[string]int{"op_admin.name": 1, "dev_admin.name": 1, "instanceId": 1}
    var array5 = [1]string{"test"}  
    fieldInPostData51 := map[string][1]string{"$in": array5}
    fieldInPostData52 := map[string]map[string][1]string{"name": fieldInPostData51}
    postData5 :=  map[string]interface{} {"page_size":2000, "page":1}
    postData5["fields"] = fieldInPostData5
    postData5["query"] = fieldInPostData52
    fmt.Println("[DATA]: ",postData5)
    ret5 ,isSuccess5 := easyApi.SendRequest("/cmdb/object/BUSINESS/instance/_search", "POST", postData5)
    fmt.Println("[Result]", ret5, isSuccess5) 

    fmt.Println("-------POST EXAMPLE-------")
    fieldInPostData6 := map[string]int{"developer.name": 1, "owner.name": 1, "instanceId": 1}
    var array6 = [1]string{"123_123"}  
    fieldInPostData61 := map[string][1]string{"$in": array6}
    fieldInPostData62 := map[string]map[string][1]string{"name": fieldInPostData61}
    postData6 :=  map[string]interface{} {"page_size":2000, "page":1}
    postData6["fields"] = fieldInPostData6
    postData6["query"] = fieldInPostData62
    fmt.Println("[DATA]: ",postData6)
    ret6 ,isSuccess6 := easyApi.SendRequest("/cmdb/object/APP/instance/_search", "POST", postData6)
    fmt.Println("[Result]", ret6, isSuccess6) 


    
}

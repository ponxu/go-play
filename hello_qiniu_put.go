package main

import (
    "qiniupkg.com/api.v7/kodo"
    "qiniupkg.com/api.v7/conf"
    "qiniupkg.com/api.v7/kodocli"
    "fmt"
)

var (
//设置上传到的空间
    bucket = "myfile"
)

//构造返回值字段
type PutRet struct {
    Hash string `json:"hash"`
    Key  string `json:"key"`
}

func main() {
    //初始化AK，SK
    conf.ACCESS_KEY = "ORVCeAIb5D5uBE_xhNQ1eerZxvzgwR9DmTnf46bA"
    conf.SECRET_KEY = "PQUY-B8zksdS5ehDox_OXMnpVZ73J1FDNcoMbwHm"

    //创建一个Client
    c := kodo.New(0, nil)

    //设置上传的策略
    policy := &kodo.PutPolicy{
        Scope:   bucket,
        //设置Token过期时间
        Expires: 3600,
    }
    //生成一个上传token
    token := c.MakeUptoken(policy)

    //构建一个uploader
    zone := 0
    uploader := kodocli.NewUploader(zone, nil)

    var ret PutRet
    //设置上传文件的路径
    filepath := "/home/xwz/test.txt"

    //调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
    //res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)

    err := uploader.PutFile(nil, &ret, token, "test.txt", filepath, nil)

    //打印返回的信息
    fmt.Println(ret)
    //打印出错信息
    if err != nil {
        fmt.Println("io.Put failed:", err)
        return
    }

}



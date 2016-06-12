package main

import (
    "qiniupkg.com/api.v7/kodo"
    "qiniupkg.com/api.v7/conf"
)

var (
//指定私有空间的域名
    domain = "7xv4yh.com2.z0.glb.clouddn.com"
//指定文件的key
    key = "test.txt"
)

//调用封装好的downloadUrl方法生成一个下载链接
func downloadUrl(domain, key string) string {
    //调用MakeBaseUrl()方法将domain,key处理成http://domain/key的形式
    baseUrl := kodo.MakeBaseUrl(domain, key)
    policy := kodo.GetPolicy{Expires:20}
    //生成一个client对象
    c := kodo.New(0, nil)
    //调用MakePrivateUrl方法返回url
    return c.MakePrivateUrl(baseUrl, &policy)
}

func main() {
    //初始化AK，SK
    conf.ACCESS_KEY = "ORVCeAIb5D5uBE_xhNQ1eerZxvzgwR9DmTnf46bA"
    conf.SECRET_KEY = "PQUY-B8zksdS5ehDox_OXMnpVZ73J1FDNcoMbwHm"
    //打印出下载链接
    println(downloadUrl(domain, key))
}
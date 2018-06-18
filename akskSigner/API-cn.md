# Package signer
    import "github.com/huaweicloudsdk/signer"

**[概述](#概述)**  

**[目录](#目录)**  

**[开始](#开始)**  

## 概述
signer包实现了华为云AK/SK认证客户端所必要的功能

示例代码，使用AK/SK对http request附加签名。
```
client := &http.Client{
		Timeout: time.Duration(3 * time.Second),
	}
req, err := http.NewRequest("POST", "https://30030113-3657-4fb6-a7ef-90764239b038.apigw.cn-north-1.huaweicloud.com/app1?name=value", bytes.NewBuffer([]byte("demo")))

signOptions := signer.SignOptions{
		AccessKey: "------------",
		SecretKey: "------------",
	}

signer.Sign(req, signOptions)
resp, err := client.Do(req)
```

## 目录
**[func Sign(req *http.Request, signOptions SignOptions)](#user-content-sign)**  
**[func HmacSha256(data string,key []byte) []byte](#user-content-hmacsha256)**  
**[func HashSha256(msg []byte) []byte](#user-content-hashsha256)**  
**[SignOptions](#user-content-signoptions)**

## 开始
## Sign
    func Sign(req *http.Request, signOptions SignOptions) 

对http.Request实例添加AK/SK签名，如果请求的body比较大，建议先行使用signer.HashSha256(body)方法计算hashValue，并添加到request header上（key为x-sdk-content-sha256）

## HmacSha256
    func HmacSha256(data string,key []byte) []byte  

根据密钥和信息，使用hmac算法生成消息摘要

## HashSha256
    func HashSha256(msg []byte) []byte

使用sha256算法，生成消息摘要

## SignOptions

- AccessKey `用户访问密钥,登陆华为云控制台账户管理处获取`
- SecretKey `用户加密密钥,登陆华为云控制台账户管理处获取`
- SignAlgorithm `指定签名算法，保持它为空会使用默认值SDK-HMAC-SHA256，目前仅支持该算法`
- TimeOffsetInseconds `时间offset，用于需要对请求的时间戳做微调的场景`
- EnableCacheSignKey `启用或关闭Signkey的cache，默认为关闭状态，亦即是每次重新计算signKey`
- RegionName `Region Name`
- ServiceName `Service Name`
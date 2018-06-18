# Package signer
    import "github.com/huaweicloudsdk/signer"

**[Overview](#overview)**  

**[Index](#index)**  

**[Start](#start)**  

## Overview
signer implemented the requied functions for Huawei AK/SK authentication

Sample code, use AK/SK sign http.Request instance
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

## Index
**[func Sign(req *http.Request, signOptions SignOptions)](#user-content-sign)**  
**[func HmacSha256(data string,key []byte) []byte](#user-content-hmacsha256)**  
**[func HashSha256(msg []byte) []byte](#user-content-hashsha256)**  
**[SignOptions](#user-content-signoptions)**

## Start
## Sign
    func Sign(req *http.Request, signOptions SignOptions) 

Sign http.Request instance with AK/SK requirement, please compete body hash value though signer.HashSha256(body) and add to request header上（key为x-sdk-content-sha256 once the body is large.
 
## HmacSha256
    func HmacSha256(data string,key []byte) []byte  

Compute digest value use hmac algorithm

## HashSha256
    func HashSha256(msg []byte) []byte

Compute digest value use sha256 algorithm

## SignOptions

- AccessKey `AccessKey, get it in developement console in official site`
- SecretKey `SecretKey, get it in developement console in official site`
- SignAlgorithm `Sigh algorithm，The default value is SDK-HMAC-SHA256，and can be only this value in currently`
- TimeOffsetInseconds `Timeoffset, used when you want to adjust the request time`
- EnableCacheSignKey `Enable or disable signKey cache，it is disabled by default，it means recompute signKey every time`
- RegionName `Region Name`
- ServiceName `Service Name`
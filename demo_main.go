package main

import (
	"fmt"

	captcha "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func demoMain() {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		"AKIDGI5cuvfBS********OwQnvqX0FF0V",
		"LQG1Y6RZUBd************HdmhRwEI",
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "captcha.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := captcha.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := captcha.NewDescribeCaptchaResultRequest()

	request.CaptchaType = common.Uint64Ptr(9)
	request.Ticket = common.StringPtr("t03YD7nLiPa7MC48_w0n_UnidHHs191XPoVIuQ23PteFhFJxhRqv6B1UuyjMAjK88BFGIYtmK3i8nMchqfR-R46UnAbdFA69vA3dYejKEExQjwljHanq5vMyUEiNTEOEC4vKkpW12ELPaE*")
	request.UserIp = common.StringPtr("127.0.0.1")
	request.BusinessId = common.Uint64Ptr(1)
	request.SceneId = common.Uint64Ptr(1)
	request.MacAddress = common.StringPtr("")
	request.Imei = common.StringPtr("")
	request.Randstr = common.StringPtr("@clI")
	request.CaptchaAppId = common.Uint64Ptr(1120)
	request.AppSecretKey = common.StringPtr("He941s*********o6r6yPhNYJE")
	request.NeedGetCaptchaTime = common.Int64Ptr(1)

	// 返回的resp是一个DescribeCaptchaResultResponse的实例，与请求对象对应
	response, err := client.DescribeCaptchaResult(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

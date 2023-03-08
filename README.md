# go语言腾讯云图像验证 （检验是否是机器人）

**此实验仅仅是测试该功能起到的作用，并没有用于完整的web开发系统，简单实现了调用腾讯云相关接口来实现图像验证的功能**

**实现步骤：在腾讯云新建图形验证，获取验证密钥，放入配置文件中，另外连接腾讯云的密钥也放入配置文件中；从腾讯云文档--验证码 中下载前端html页面，从API explore中生成相关go语言代码，做点简单的调整即可。**

***本该开源代码中依旧将密钥进行了打码***

**该实验和我的另一个开源实验“ocr_healthCode_travelCard”非常类似，因此不再做过多说明，另外腾讯云还有很多免费供大家实验的接口，实现方法也基本类似，阅读官方的文档即可快速实现。其他的腾讯云接口项目就不打算再做了，感兴趣的同学可以自己尝试。**

实验截图请看resultImage文件夹，对于后台校验部分仅有两个有效字段（rand_str与ticket），其他字段并没有过多的研究。

### 核心功能概述

```go
// captcha/captcha.go
// 以下部分简单将API explore中生成相关go语言代码进行了拆分
func (c *myCaptcha) getCredential() *common.Credential {
	credential := common.NewCredential(
		config.Secret.GetString("TencentCloudSecret.SecretId"),
		config.Secret.GetString("TencentCloudSecret.SecretKey"),
	)
	return credential
}

func (c *myCaptcha) getClient() *captcha.Client {
	credential := c.getCredential()
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.Conf.GetString("CaptchaConf.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := captcha.NewClient(credential, "", cpf)
	return client
}

func (c *myCaptcha) getRequest(req *CaptchaRequest) *captcha.DescribeCaptchaResultRequest {
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := captcha.NewDescribeCaptchaResultRequest()

	request.CaptchaType = common.Uint64Ptr(CAPTCHA_TYPE)
	request.Ticket = common.StringPtr(req.Ticket)
	request.UserIp = common.StringPtr(req.UserIp)
	request.BusinessId = common.Uint64Ptr(req.BusinessId)
	request.SceneId = common.Uint64Ptr(req.SceneId)
	request.Randstr = common.StringPtr(req.Randstr)
	captchaIdStr := config.Secret.GetString("Captcha.CapthaAppId")
	captchaId, _ := strconv.ParseUint(captchaIdStr, 10, 64)
	request.CaptchaAppId = common.Uint64Ptr(captchaId)
	request.AppSecretKey = common.StringPtr(config.Secret.GetString("Captcha.AppSecretKey"))
	request.NeedGetCaptchaTime = common.Int64Ptr(req.NeedGetCaptchaTime)
	request.Imei = common.StringPtr(req.Imei)
	return request
}

func (c *myCaptcha) DescribeCaptchaResult(req *CaptchaRequest) (string, error) {
	client := c.getClient()
	request := c.getRequest(req)
	res, err := client.DescribeCaptchaResult(request)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.ToJsonString(), err
}
```

```go
// web/captcha.go
// 接收前端数据，将数据传入上面的DescribeCaptchaResult进行校验

func CaptchaHandler(w http.ResponseWriter, req *http.Request) {
	c := captcha.NewCaptcha()
	request := &captcha.CaptchaRequest{}
	request.NeedGetCaptchaTime = 0
	request.Randstr = req.FormValue("rand_str")
	request.SceneId, _ = strconv.ParseUint(req.FormValue("scene_id"), 10, 64)
	request.BusinessId, _ = strconv.ParseUint(req.FormValue("business_id"), 10, 64)
	userIP := ClientPublicIP(req)
	if userIP != "" {
		request.UserIp = userIP
	} else {
		request.UserIp = "127.0.0.1"
	}
	request.Ticket = req.FormValue("ticket")
	request.Imei = req.FormValue("imei")
	request.MacAddress = req.FormValue("mac_address")
	fmt.Printf("#{request}")
	res, err := c.DescribeCaptchaResult(request)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, res)

}

```


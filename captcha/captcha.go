package captcha

import (
	"go_verification_code/config"
	"log"
	"strconv"

	captcha "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const (
	CAPTCHA_TYPE = 9
)

type CaptchaRequest struct {
	Ticket             string
	UserIp             string
	BusinessId         uint64
	SceneId            uint64
	Randstr            string
	MacAddress         string
	Imei               string
	NeedGetCaptchaTime int64
}

type myCaptcha struct {
}

func NewCaptcha() *myCaptcha {
	return &myCaptcha{}
}

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

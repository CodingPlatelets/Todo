package helper

import (
	"Todo/config"
	"Todo/constants"
	"Todo/db_server"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net/http"
	"time"
)

func SendMail(EmailAddress string) (ReturnType, error) {

	mailConfig := config.GetMailConfig()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	VerifyCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	keyValue := "VerifyCode" + EmailAddress
	// save to redis
	err := db_server.DeleteFromRedis(keyValue)
	if err != nil {
		return ReturnType{Status: constants.CodeError, Msg: err.Error(), Data: ""}, err
	}
	err = db_server.PutToRedis(keyValue, VerifyCode, 1000*60*15)
	if err != nil {
		return ReturnType{Status: constants.CodeError, Msg: err.Error(), Data: ""}, err
	}

	SendTime := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	html := fmt.Sprintf(`<div>
		<div>
			尊敬的%s, 您好!
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p> 您于 %s 提交了邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为15分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解和使用。 </p>
		</div>
		<div>
			Thanks,
			夜莺科技
		</div>
		<div>
			<p> 此邮箱为系统邮箱，请勿回复。</p>
			<p> 🐼 </p>
		</div>
	<div>`, EmailAddress, SendTime, VerifyCode)

	message := gomail.NewMessage()
	message.SetAddressHeader("From", mailConfig["from"].(string), mailConfig["from_name"].(string))
	message.SetHeader("To", EmailAddress)
	message.SetHeader("Subject", "[我的验证码]邮箱验证")
	message.SetBody("text/html", html)

	dia := gomail.NewDialer(mailConfig["host"].(string), mailConfig["port"].(int), mailConfig["username"].(string), mailConfig["password"].(string))

	if err := dia.DialAndSend(message); err != nil {
		return ReturnType{Status: http.StatusBadRequest, Msg: "邮件发送失败", Data: err.Error()}, err
	}
	return ReturnType{Status: http.StatusOK, Msg: "邮件发送成功，请注意查收", Data: ""}, nil
}

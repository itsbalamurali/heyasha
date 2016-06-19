package platforms

import (
	"github.com/gin-gonic/gin"
	//"github.com/hectane/hectane/queue"
	//"github.com/hectane/hectane/email"
	"net/smtp"
	"log"
	"fmt"
	"encoding/base64"
	"strings"
	"net/mail"
)

func encodeRFC2047(String string) string{
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func EmailBot(c *gin.Context) {
	sender,_ := c.GetPostForm("sender")
	from,_ := c.GetPostForm("from")
	subject,_ := c.GetPostForm("subject")

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"postmaster@heyasha.com",
		"915dc442862acfda08672c116d117be2",
		"smtp.mailgun.org",
	)

	header := make(map[string]string)
	header["From"] = from
	header["To"] = sender
	header["Subject"] = encodeRFC2047("RE: "+subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	body := "<html><body> Hi "+ from +", <br> Thanks for mailing me. <br> Currently i am still learning to read and replying to your emails <br> Regards, Yours Asha ;) </body></html>"
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))


	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	merr := smtp.SendMail(
		"smtp.mailgun.org:25",
		auth,
		"asha@heyasha.com",
		[]string{sender},
		[]byte(message),)
	if merr != nil {
		log.Fatal(merr)
	}

	c.JSON(200,gin.H{"message":"Mail Received thanks :)"})
}

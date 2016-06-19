package platforms

import (
	"github.com/gin-gonic/gin"
	//"github.com/hectane/hectane/queue"
	//"github.com/hectane/hectane/email"
	"net/smtp"
	"log"
)

func EmailBot(c *gin.Context) {
	sender,_ := c.GetPostForm("sender")
	from,_ := c.GetPostForm("from")

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"postmaster@heyasha.com",
		"915dc442862acfda08672c116d117be2",
		"smtp.mailgun.org",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	merr := smtp.SendMail(
		"smtp.mailgun.org:25",
		auth,
		"asha@heyasha.com",
		[]string{sender},
		[]byte("<html><body> Hi "+ from +", <br> Thanks for mailing me. <br> Currently i am still learning to read and replying to your emails <br> Regards, Yours Asha ;) </body></html>"),
	)
	if merr != nil {
		log.Fatal(merr)
	}

	c.JSON(200,gin.H{"message":"Mail Received thanks :)"})
}

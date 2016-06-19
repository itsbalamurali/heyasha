//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:30 PM
package config

const (
	AppSecret = "ILoveAsha<3"
)

var Config = struct {
	ENVIRONMENT string `default:"production"`

	//Databases
	MYSQL_URI string
	MONGO_URI string

	SMTP struct {
		HOSTNAME  string `required:"true"`
		USERNAME string `required:"true"`
		PASSWORD string `required:"true"`
		PORT string `required:"true"`
	}


	SMS []struct{
		HOSTNAME string
		API_KEY string
	}

	//GoogleCloudStorage
	GCS struct{
		BucketName string `required:"true"`

	}
}{}


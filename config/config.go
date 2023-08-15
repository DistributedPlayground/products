package config

type vars struct {
	AWS_ACCT              string `required:"false"`
	AWS_ACCESS_KEY_ID     string `required:"false"`
	AWS_SECRET_ACCESS_KEY string `required:"false"`
	AWS_REGION            string `required:"false"`
	AWS_KMS_KEY_ID        string `required:"false"`
	DEBUG_MODE            string `required:"false"`
	BASE_URL              string `required:"true"`
	ENV                   string `required:"true"`
	PORT                  string `required:"true"`
	DB_NAME               string `required:"true"`
	DB_USERNAME           string `required:"true"`
	DB_PASSWORD           string `required:"true"`
	DB_HOST               string `required:"true"`
	DB_PORT               string `required:"true"`
}

var Var vars

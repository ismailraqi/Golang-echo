package tronics

//ConfigDatabase read a configuration file and environment variables in a single function call.
type ConfigDatabase struct {
	AppName  string `env:"APP_NAME" env-default:"TRONICS"`
	AppEnv   string `env:"APP_ENV" env-default:"DEV"`
	Port     string `env:"MY_APP_PORT" env-default:"3030"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG8LEVEL" env-default:"ERROR"`
}

var cfg ConfigDatabase

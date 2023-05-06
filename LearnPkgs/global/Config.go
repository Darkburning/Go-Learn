package global

type Config struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}
type Redis struct {
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	PoolSize int    `yaml:"poolSize"`
}

type Mysql struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Db            string `yaml:"db"`
	User          string `yaml:"user"`
	PassWord      string `yaml:"password"`
	LogLevel      string `yaml:"log_level"`
	Configuration string `yaml:"configuration"`
}

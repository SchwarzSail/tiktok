package config

var Config Conf

type Conf struct {
	System   *System `mapstructure:"system"`
	Mysql    `mapstructure:"mysql"`
	Redis    `mapstructure:"redis"`
	Rabbitmq `mapstructure:"rabbitmq"`
	Etcd     `mapstructure:"etcd"`
}

type System struct {
	Domain string `mapstructure:"domain"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
}

type Mysql struct {
	UserName      string `mapstructure:"userName"`
	MysqlPassword string `mapstructure:"mysqlPassword"`
	DbName        string `mapstructure:"dbName"`
	DbHost        string `mapstructure:"dbHost"`
	DbPort        string `mapstructure:"dbPort"`
}

type Redis struct {
	RedisHost     string `mapstructure:"redisHost"`
	RedisPort     string `mapstructure:"redisPort"`
	RedisPassword string `mapstructure:"redisPassword"`
	RedisDbName   int    `mapstructure:"redisDbName"`
}

type Rabbitmq struct {
	RabbitmqUser     string `mapstructure:"rabbitmqUserName"`
	RabbitmqPassword string `mapstructure:"rabbitmqPassword"`
	RabbitmqHost     string `mapstructure:"rabbitmqHost"`
	RabbitmqPort     string `mapstructure:"rabbitmqPort"`
}

type Etcd struct {
	EtcdHost string `mapstruct:"etcdHost"`
	EtcdPort string `mapstruct:"etcdPort"`
}

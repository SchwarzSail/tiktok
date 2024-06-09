package config

var Config Conf

type Conf struct {
	System   *System `mapstructure:"system"`
	Rabbitmq `mapstructure:"rabbitmq"`
	Etcd     `mapstructure:"etcd"`
}

type System struct {
	Domain string `mapstructure:"domain"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
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

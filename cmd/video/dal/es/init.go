package es

import (
	"github.com/olivere/elastic/v7"
	"tiktok/config"
)

var EsClient *elastic.Client

func LinkEs() {
	conf := config.Config.Es
	url := "http://" + conf.Host + ":" + conf.Port
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false), //是否开启集群嗅探
		elastic.SetBasicAuth(conf.UserName, conf.Password),
	)
	if err != nil {
		panic(err)
	}
	EsClient = client
}

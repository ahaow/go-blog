package initialize

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"go-blog/global"
	"go.uber.org/zap"
	"os"
)

func ConnectES() *elasticsearch.TypedClient {
	esConf := global.Config.ES
	conf := elasticsearch.Config{
		Addresses: []string{esConf.URL},
		Username:  esConf.Username,
		Password:  esConf.Password,
	}

	if esConf.IsConsolePrint {
		conf.Logger = &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}

	client, err := elasticsearch.NewTypedClient(conf)
	if err != nil {
		global.Log.Error("Failed to connect to Elasticsearch:", zap.Error(err))
		os.Exit(1)
	}
	return client
}

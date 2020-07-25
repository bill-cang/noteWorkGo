package dbcomm

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"net"
	"net/http"
	"noteWork/bestLib/terror"
	"time"
)

var (
	esClient *elasticsearch.Client
)

func InitElasticsearch(username, passwd string, addresses ...string) (err error) {
	dialer := &net.Dialer{
		Timeout: 1 * time.Second,
	}
	esClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  passwd,
		//Logger: &estransport.TextLogger{
		//	Output:             os.Stdout,
		//	EnableRequestBody:  true,
		//	EnableResponseBody: true,
		//},
		Transport: &http.Transport{
			MaxIdleConns:    80,
			IdleConnTimeout: 500 * time.Second,
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				return dialer.DialContext(ctx, network, addr)
			},
		},
	})

	if err != nil {
		return
	}

	response, err := esClient.Ping()
	if err != nil {
		return
	}

	if response.IsError() {
		err = terror.New(response.StatusCode, response.Status())
		return
	}
	return
}

func GetEsClient() *elasticsearch.Client {
	return esClient
}

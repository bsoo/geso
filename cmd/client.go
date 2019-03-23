package cmd

import (
	"io/ioutil"
	"log"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/spf13/viper"
)

type Client struct {
	*elasticsearch.Client
}

func newClient() (*Client, error) {
	viper.BindPFlag("url", rootCmd.Flags().Lookup("url"))
	endpointURL := viper.GetString("url")
	conf := elasticsearch.Config{
		Addresses: []string{endpointURL},
	}

	es, err := elasticsearch.NewClient(conf)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
		return nil, err
	}
	client := &Client{
		Client: es,
	}
	return client, nil
}

func (c *Client) decodeBody(resp *esapi.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

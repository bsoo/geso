package cmd

import (
	"context"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "search",
		Short:         "search document",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          runSearchCmd,
	}
	return cmd
}

func runSearchCmd(cmd *cobra.Command, args []string) error {
	client, err := newClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := getSearchOption(cmd)

	res, err := client.Search(ctx, opt)
	if err != nil {
		return errors.Wrapf(err, "res = %+v", res)
	}
	cmd.Println(string(res))
	return nil
}

func getSearchOption(cmd *cobra.Command) *SearchOption {
	opt := SearchOption{}
	viper.IsSet("")
	return &opt
}

func (c *Client) Search(ctx context.Context, opt *SearchOption) ([]byte, error) {
	req := esapi.SearchRequest{
		Index: opt.Index,
		Query: opt.Query,
	}
	resp, err := req.Do(ctx, c.Client)
	if err != nil {
		return nil, err
	}
	body, err := c.decodeBody(resp)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type SearchOption struct {
	Index []string
	Query string
}

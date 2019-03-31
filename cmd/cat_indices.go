package cmd

import (
	"context"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newCatIndicesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "indices",
		Short:         "show indices status",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          runCatIndicesCmd,
	}
	return cmd
}

func runCatIndicesCmd(cmd *cobra.Command, args []string) error {
	client, err := newClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.CatIndices(ctx)
	if err != nil {
		return errors.Wrapf(err, "res = %+v", res)
	}
	cmd.Println(string(res))
	return nil
}

func (c *Client) CatIndices(ctx context.Context) ([]byte, error) {
	req := esapi.CatIndicesRequest{
		V: boolPtr(true),
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

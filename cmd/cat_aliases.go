package cmd

import (
	"context"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newCatAliasesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "aliases",
		Short:         "show aliases status",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          runCatAliasesCmd,
	}
	setupFlags(cmd)
	return cmd
}

func setupFlags(cmd *cobra.Command) {
	cmd.Flags().StringSlice("name", []string{}, "alias name")
	cmd.Flags().String("format", "", "format output")
	cmd.Flags().StringSlice("h", []string{}, "header")
	cmd.Flags().Bool("local", false, "")
	cmd.Flags().Duration("timeout", 0, "set timeout")
	cmd.Flags().StringSliceP("sort", "s", []string{}, "sort")
	cmd.Flags().BoolP("verbose", "v", false, "verbose output")
	cmd.Flags().Bool("pretty", false, "pretty output")
	cmd.Flags().Bool("human", false, "")
	cmd.Flags().Bool("error-trace", false, "")
	cmd.Flags().StringSlice("filter-path", []string{}, "")
	viper.BindPFlag("name", cmd.Flags().Lookup("name"))
	viper.BindPFlag("format", cmd.Flags().Lookup("format"))
	viper.BindPFlag("h", cmd.Flags().Lookup("h"))
	viper.BindPFlag("local", cmd.Flags().Lookup("local"))
	viper.BindPFlag("timeout", cmd.Flags().Lookup("timeout"))
	viper.BindPFlag("sort", cmd.Flags().Lookup("sort"))
	viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))
	viper.BindPFlag("pretty", cmd.Flags().Lookup("pretty"))
	viper.BindPFlag("human", cmd.Flags().Lookup("human"))
	viper.BindPFlag("error-trace", cmd.Flags().Lookup("error-trace"))
	viper.BindPFlag("filter-path", cmd.Flags().Lookup("filter-path"))
}

func runCatAliasesCmd(cmd *cobra.Command, args []string) error {
	opt := getCatAliasesOptions()
	return catAliases(cmd, args, opt)
}

func catAliases(cmd *cobra.Command, args []string, opt *catAliasesOptions) error {
	client, err := newClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.CatAliases(ctx, opt)
	if err != nil {
		return errors.Wrapf(err, "res = %+v", res)
	}
	cmd.Println(string(res))
	return nil
}

func getCatAliasesOptions() *catAliasesOptions {
	opt := catAliasesOptions{}
	if viper.IsSet("name") {
		opt.Name = viper.GetStringSlice("name")
	}
	opt.Format = viper.GetString("format")
	if viper.IsSet("h") {
		opt.H = viper.GetStringSlice("h")
	}
	if viper.IsSet("local") {
		local := viper.GetBool("local")
		opt.Local = &local
	}
	opt.MasterTimeout = viper.GetDuration("timeout")
	opt.S = viper.GetStringSlice("sort")
	if viper.IsSet("verbose") {
		verbose := viper.GetBool("verbose")
		opt.V = &verbose
	}
	opt.Pretty = viper.GetBool("pretty")
	opt.Human = viper.GetBool("human")
	opt.ErrorTrace = viper.GetBool("error-trace")
	if viper.IsSet("filter-path") {
		opt.FilterPath = viper.GetStringSlice("filter-path")
	}
	return &opt
}

func (c *Client) CatAliases(ctx context.Context, opt *catAliasesOptions) ([]byte, error) {
	req := opt.ToRequest()
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

type catAliasesOptions struct {
	Name          []string
	Format        string
	H             []string
	Help          *bool
	Local         *bool
	MasterTimeout time.Duration
	S             []string
	V             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string
}

func (o *catAliasesOptions) ToRequest() esapi.CatAliasesRequest {
	req := esapi.CatAliasesRequest{
		Name:          o.Name,
		Format:        o.Format,
		H:             o.H,
		Help:          o.Help,
		Local:         o.Local,
		MasterTimeout: o.MasterTimeout,
		S:             o.S,
		V:             o.V,
		Pretty:        o.Pretty,
		Human:         o.Human,
		ErrorTrace:    o.ErrorTrace,
		FilterPath:    o.FilterPath,
	}
	return req
}

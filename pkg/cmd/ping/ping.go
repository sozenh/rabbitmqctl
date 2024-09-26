package ping

import (
	"context"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/rocketmqctl/pkg/rocketmq"
	"github.com/rocketmqctl/pkg/utils"
)

type Options struct {
	Hosts []string

	AccessKey string
	SecretKey string
}

func NewCmdPing() *cobra.Command {
	o := &Options{}

	cmd := &cobra.Command{
		Use:   "ping",
		Short: "ping",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(o.Run(context.TODO()))
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(o.CheckFlags())
		},
		PostRun: func(cmd *cobra.Command, args []string) {

		},
	}

	o.Parse(cmd)
	return cmd
}

func (ping *Options) CheckFlags() error {
	return nil
}

func (ping *Options) Parse(c *cobra.Command) {
	c.Flags().StringSliceVarP(&ping.Hosts, "hosts", "x", nil, "rocketmq hosts")
	c.Flags().StringVarP(&ping.AccessKey, "accesskey", "a", "", "rocketmq accessKey")
	c.Flags().StringVarP(&ping.SecretKey, "secretkey", "s", "", "rocketmq secretKey")
}

func (ping *Options) Run(ctx context.Context) error {
	mqAdmin, err := rocketmq.
		NewMqAdmin(ping.Hosts, ping.AccessKey, ping.SecretKey)
	if err != nil {
		return fmt.Errorf("connect to rocketmq cluster failed: %v", err)
	}

	clusterInfo, err := mqAdmin.FetchClusterInfo(ctx)
	if err != nil {
		return fmt.Errorf("fetch rocketmq cluster info failed: %v", err)
	}

	for name, addrs := range clusterInfo.BrokerAddrTable {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{name, "BrokerName", "ID", "Role", "Addr"})
		for id, addr := range addrs.BrokerAddrs {
			table.Append([]string{addrs.Cluster, addrs.BrokerName, id, utils.BrokerRole(id), addr})
		}
		table.Render()
	}

	return nil
}

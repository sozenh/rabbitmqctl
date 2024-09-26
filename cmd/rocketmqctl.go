package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/rocketmqctl/pkg/cmd/consume"
	"github.com/rocketmqctl/pkg/cmd/ping"
	"github.com/rocketmqctl/pkg/cmd/product"
	"github.com/rocketmqctl/pkg/utils"
)

func main() {
	command := NewCommand()

	err := command.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rocketmqctl",
		Short: "rocketmqctl use to manager rocketmq cluster",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(cmd.Help())
		},
	}

	cmd.AddCommand(ping.NewCmdPing())
	cmd.AddCommand(consume.NewCmdConsume())
	cmd.AddCommand(product.NewCmdProduct())

	return cmd
}

package consume

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"

	"github.com/rocketmqctl/pkg/cmd/ping"
	"github.com/rocketmqctl/pkg/utils"
)

type Options struct {
	ping.Options

	Topic string
	Group string
}

func NewCmdConsume() *cobra.Command {
	o := &Options{}

	cmd := &cobra.Command{
		Use:   "consume",
		Short: "consume",
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

func (consume *Options) CheckFlags() error {
	return nil
}

func (consume *Options) Parse(c *cobra.Command) {
	consume.Options.Parse(c)

	c.Flags().StringVarP(&consume.Topic, "topic", "t", "default", "rocketmq topic")
	c.Flags().StringVarP(&consume.Group, "group", "g", "default", "rocketmq group")
}

func (consume *Options) Run(ctx context.Context) error {

	c, _ := consumer.NewPushConsumer(
		consumer.WithGroupName(consume.Group),
		consumer.WithNameServer(consume.Hosts),
		consumer.WithCredentials(primitive.Credentials{AccessKey: consume.AccessKey, SecretKey: consume.SecretKey}),
	)

	err := c.Subscribe(
		consume.Topic,
		consumer.MessageSelector{},
		func(
			ctx context.Context,
			msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				fmt.Printf("subscribe callback: %s\n", msg.String())
			}
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		return fmt.Errorf("subscribe error: %v", err)
	}

	err = c.Start()
	if err != nil {
		return fmt.Errorf("start producer error: %v", err)
	}

	return nil
}

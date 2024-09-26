package product

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"

	"github.com/rocketmqctl/pkg/cmd/ping"
	"github.com/rocketmqctl/pkg/utils"
)

type Options struct {
	ping.Options

	Topic        string
	Group        string
	MessageCount int
}

func NewCmdProduct() *cobra.Command {
	o := &Options{}

	backupCmd := &cobra.Command{
		Use:   "product",
		Short: "product",
		Run: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(o.Run(context.TODO()))
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			utils.CheckErr(o.CheckFlags())
		},
		PostRun: func(cmd *cobra.Command, args []string) {

		},
	}

	o.Parse(backupCmd)
	return backupCmd
}

func (product *Options) CheckFlags() error {
	return nil
}

func (product *Options) Parse(c *cobra.Command) {
	product.Options.Parse(c)

	c.Flags().StringVarP(&product.Topic, "topic", "t", "default", "rocketmq topic")
	c.Flags().StringVarP(&product.Group, "group", "g", "default", "rocketmq group")
	c.Flags().IntVarP(&product.MessageCount, "message", "m", 100, "how many message do you want to product")
}

func (product *Options) Run(ctx context.Context) error {

	p, _ := producer.NewDefaultProducer(
		producer.WithRetry(2),
		producer.WithGroupName(product.Group),
		producer.WithNameServer(product.Hosts),
		producer.WithCredentials(primitive.Credentials{AccessKey: product.AccessKey, SecretKey: product.SecretKey}),
	)

	err := p.Start()
	if err != nil {
		return fmt.Errorf("start producer error: %v", err)
	}

	for i := 0; i < product.MessageCount; i++ {
		res, err := p.SendSync(
			context.Background(),
			&primitive.Message{
				Topic: product.Topic,
				Body:  []byte(fmt.Sprintf("Hello RocketMQ %d!", i)),
			},
		)
		if err != nil {
			fmt.Printf("send message error: %v\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}

	err = p.Shutdown()
	if err != nil {
		return fmt.Errorf("shutdown producer error: %v", err)
	}

	return nil
}

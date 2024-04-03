package producer

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	host  string
	port  string
	topic string
)

// producerCmd represents the producer command
var ProducerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Command to produce broker messages",
	Long:  ``,
	Run:   handleProducer,
}

func init() {
	ProducerCmd.Flags().StringVarP(&host, "host", "u", "localhost", "Broker url to connect")
	ProducerCmd.Flags().StringVarP(&port, "port", "p", "8080", "Broker port to connect")
	ProducerCmd.Flags().StringVarP(&topic, "topic", "t", "", "Broker topic to produce messages")

	if err := ProducerCmd.MarkFlagRequired("topic"); err != nil {
		fmt.Println(err)
	}
}

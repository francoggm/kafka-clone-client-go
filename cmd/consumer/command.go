package consumer

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	host string
	port string
)

// consumerCmd represents the consumer command
var ConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Command to consume broker messages",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consumer called")
	},
}

func init() {
	ConsumerCmd.Flags().StringVarP(&host, "host", "u", "localhost", "Broker url to connect")
	ConsumerCmd.Flags().StringVarP(&port, "port", "p", "8080", "Broker port to connect")
}

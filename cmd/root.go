package cmd

import (
	"kafka-client/cmd/consumer"
	"kafka-client/cmd/producer"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafka-client",
	Short: "A client to connect in Kafka broker",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommands() {
	rootCmd.AddCommand(producer.ProducerCmd)
	rootCmd.AddCommand(consumer.ConsumerCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommands()
}

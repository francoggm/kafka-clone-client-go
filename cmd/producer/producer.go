package producer

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func handleProducer(cmd *cobra.Command, args []string) {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		cmd.PrintErrf("Error: Failed to connect in %s:%s\n%v\n", host, port, err)
		return
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(fmt.Sprintf("0 %s", topic))); err != nil {
		cmd.PrintErrf("Error: Failed to send config values in %s:%s\n%v\n", host, port, err)
		return
	}

	for {
		var message string

		fmt.Print(">")
		if _, err := fmt.Scanf("%s\n", &message); err != nil {
			cmd.PrintErrf("Error: Invalid input %v\n", err)
			continue
		}

		if _, err := conn.Write([]byte(message)); err != nil {
			cmd.PrintErrf("Error: Failed to connect in %s:%s\n%v\n", host, port, err)
			break
		}
	}
}

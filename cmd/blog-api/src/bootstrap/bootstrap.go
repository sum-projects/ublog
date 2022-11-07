package bootstrap

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sum-project/ublog/cmd/blog-api/src/server"
	"os"
)

const (
	flagAddress = "addr"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "Start",
		Run: func(cmd *cobra.Command, args []string) {
			addr, _ := cmd.Flags().GetString(flagAddress)

			srv, err := server.NewApiServer(addr)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			srv.Start()
		},
	}

	rootCmd.Flags().String(flagAddress, "Server addr", ":8080")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing CLI: '%s'", err)
		os.Exit(1)
	}
}

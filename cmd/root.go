package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dodiicahya/trade/usecase"
	"github.com/spf13/cobra"
)

// Start handler registering service command
func Start() {

	rootCmd := &cobra.Command{}
	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	cmd := []*cobra.Command{

		{
			Use:   "max-trade",
			Short: "max profit",
			Run: func(cmd *cobra.Command, args []string) {
				usecase.Max(ctx)
			},
		},
		{
			Use:   "unique-string",
			Short: "unique character ",
			Run: func(cmd *cobra.Command, args []string) {
				usecase.UniqueString(ctx)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

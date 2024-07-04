package cmd

import (
	"context"
	"os/signal"

	"github.com/spf13/cobra"

	garmWs "github.com/cloudbase/garm/websocket"
)

var eventsFilters string

var logCmd = &cobra.Command{
	Use:          "debug-log",
	SilenceUsage: true,
	Short:        "Stream garm log",
	Long:         `Stream all garm logging to the terminal.`,
	RunE: func(_ *cobra.Command, _ []string) error {
		ctx, stop := signal.NotifyContext(context.Background(), signals...)
		defer stop()

		reader, err := garmWs.NewReader(ctx, mgr.BaseURL, "/api/v1/ws", mgr.Token)
		if err != nil {
			return err
		}

		if err := reader.Start(); err != nil {
			return err
		}

		<-reader.Done()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

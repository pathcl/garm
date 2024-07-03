package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"

	"github.com/cloudbase/garm-provider-common/util"
	apiParams "github.com/cloudbase/garm/apiserver/params"
	garmWs "github.com/cloudbase/garm/websocket"
)

var eventsFilters string

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

func getWebsocketConnection(pth string) (*websocket.Conn, error) {
	parsedURL, err := url.Parse(mgr.BaseURL)
	if err != nil {
		return nil, err
	}

	wsScheme := "ws"
	if parsedURL.Scheme == "https" {
		wsScheme = "wss"
	}
	u := url.URL{Scheme: wsScheme, Host: parsedURL.Host, Path: pth}
	slog.Debug("connecting", "url", u.String())

	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", mgr.Token))

	c, response, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		var resp apiParams.APIErrorResponse
		var msg string
		var status string
		if response != nil {
			if response.Body != nil {
				if err := json.NewDecoder(response.Body).Decode(&resp); err == nil {
					msg = resp.Details
				}
			}
			status = response.Status
		}
		return nil, fmt.Errorf("failed to stream logs: %q %s (%s)", err, msg, status)
	}
	return c, nil
}

var logCmd = &cobra.Command{
	Use:          "debug-log",
	SilenceUsage: true,
	Short:        "Stream garm log",
	Long:         `Stream all garm logging to the terminal.`,
	RunE: func(_ *cobra.Command, _ []string) error {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		conn, err := getWebsocketConnection("/api/v1/ws")
		if err != nil {
			return err
		}
		defer conn.Close()

		done := make(chan struct{})

		go func() {
			defer close(done)
			conn.SetReadDeadline(time.Now().Add(pongWait))
			conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					if garmWs.IsErrorOfInterest(err) {
						slog.With(slog.Any("error", err)).Error("reading log message")
					}
					return
				}
				fmt.Println(util.SanitizeLogEntry(string(message)))
			}
		}()

		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return nil
			case <-ticker.C:
				conn.SetWriteDeadline(time.Now().Add(writeWait))
				err := conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					return err
				}
			case <-interrupt:
				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				conn.SetWriteDeadline(time.Now().Add(writeWait))
				err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					return err
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
				return nil
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

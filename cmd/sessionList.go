package cmd

import (
	"context"
	api "github.com/32leaves/ruruku/pkg/server/api/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"io"
	"time"
)

// sessionListCmd represents the sessionList command
var sessionListCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints a table of the available sessions and their status",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(remoteCmdValues.server, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := api.NewSessionServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(remoteCmdValues.timeout)*time.Second)
		defer cancel()

		stream, err := client.List(ctx, &api.ListSessionsRequest{})
		if err != nil {
			log.WithError(err).Fatal()
		}

		resp := make([]*api.ListSessionsResponse, 0)
		mxidlen := 0
		for {
			session, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.WithError(err).Fatal()
			}
			resp = append(resp, session)
			if len(session.Id) > mxidlen {
				mxidlen = len(session.Id)
			}
		}

		tpl := `ID	IS OPEN	NAME
{{- range . }}
{{ .Id }}	{{ .IsOpen }}	{{ .Name -}}
{{ end }}
`
		ctnt := remoteCmdValues.GetOutputFormat(resp, tpl)
		if err := ctnt.Print(); err != nil {
			log.WithError(err).Fatal()
		}
	},
}

func init() {
	sessionCmd.AddCommand(sessionListCmd)
}

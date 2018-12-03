package cmd

import (
	"github.com/32leaves/ruruku/pkg/server"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
)

var sessionName string

// serveCmd represents the start command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a ruruku API server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := GetConfigFromViper()
		if err != nil {
			log.Fatalf("Error while loading the configuration: %v", err)
		}

		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			log.Fatalf("Error while opening database: %v", err)
		}
        store := server.NewGormBackedSessionStore(db)
        return
        // store := server.NewMemoryBackedSessionStore()

		srvcfg := cfg.Server
		if err := server.Start(&srvcfg, store); err != nil {
			log.Fatalf("Error while starting the ruruku server: %v", err)
		}

		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt)
		<-signalChannel
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().Int("ui-port", 8080, "Port to run UI the server on")
	viper.BindPFlag("server.ui.port", serveCmd.Flags().Lookup("ui-port"))
}

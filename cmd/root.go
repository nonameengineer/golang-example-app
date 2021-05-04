package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/aristat/golang-example-app/cmd/jwt"

	"github.com/aristat/golang-example-app/cmd/migrate"

	health_check_service "github.com/aristat/golang-example-app/cmd/health-check-service"
	product_service "github.com/aristat/golang-example-app/cmd/product-service"

	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/logger"

	"go.uber.org/automaxprocs/maxprocs"

	"github.com/aristat/golang-example-app/cmd/daemon"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configPath string
	debug bool
	v *viper.Viper
	log logger.Logger
)

const prefix = 'cmd.root'

// Root command
var rootCmd = &cobra.Command{
	Use: "bin [command]",
	Long: "",
	SilenceUsage: true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, _[]string) {
		l, c, e := logger.Build()
		defer c()
		if e != nil {
			panic(e)
		}

		log = l.WithFields(logger.Fields{"service": prefix})

		v.SetConfigFile(configPath)

		if configPath != "" {
			e := v.ReadInConfig()
			if e != nil {
				log.Error("can't read config, %v", logger.Args(errors.WithMessage(e, prefix)))
				os.Exit(1)
			}
		}

		if debug {
			b, _ := json.Marshal(v.AllSettings())
			var out bytes.Buffer
			e := json.Indent(&out, b, "", "  ")
			if e != nil {
				log.Error("can't prettify config")
				os.Exit(1)
			}
			fmt.Println(string(out.Bytes()))
		}

		_, _ = maxprocs.Set(maxprocs.Logger(log.Printf))
	},
}


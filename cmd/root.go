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
)
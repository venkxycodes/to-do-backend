package commands

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
	"to-do/config"
)

var globalConfig = &config.Config{}

func SetupCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "to-do",
		Short:             "lb",
		PersistentPreRunE: commandSetup(),
		SilenceUsage:      true,
	}
	cmd.AddCommand(apiServerCommand())
	return cmd
}

func commandSetup() func(*cobra.Command, []string) error {
	return func(command *cobra.Command, strings []string) (err error) {
		if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
			return err
		}
		if config.GetConfig().ENV != "dev" && config.GetConfig().JWTSecret == "" {
			return fmt.Errorf("JWT_SECRET must be set outside development")
		}
		return
	}
}

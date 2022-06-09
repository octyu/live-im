package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"live-im/internal/app/routers"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "live-im",
	Short: "A simple server used in live",
	Long: "A simple im server used in live",
}


func Execute()  {
	router := routers.InitRouter()
	err := router.Run()
	if err != nil {
		return
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
package cmd

import (
	"github.com/gookit/color"
	"github.com/mamau/starter/config"
	"github.com/mamau/starter/entity"
	"github.com/mamau/starter/libs"
	"github.com/mamau/starter/services"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			color.Red.Printf("You should pass args")
			return
		}
		serviceName := args[0]
		color.Cyan.Printf("Start %s\n", serviceName)

		config := config.GetConfig()
		s := config.GetService(serviceName)

		ser := entity.NewService(s, args[1:])
		collector := services.NewCollector(ser)
		libs.RunCommandAtPTY(Docker(collector))
	},
}
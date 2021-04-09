package cmd

import (
	"fmt"
	manager "github.com/richard2259/dependencies-manager"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

//getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Return sorted docker-compose services as slice",
	Long: "Return of sorted docker-compose services as a slice. " +
		"The first argument MUST be the path to the docker-compose file, the second argument is optional. " +
		"If the second argument exists will return the element by this index else return the whole slice. " +
		"You can pass many indexes and the program return all services for these indexes.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Arguments is empty!")
			os.Exit(-1)
		}

		sortedServices := manager.ReturnAllSortedServices(args[0])

		if len(args) > 1 {
			var services []string
			for i := 1; i < len(args); i++ {
				services = append(services, sortedServices[i])
			}
			fmt.Println(strings.Join(services[:], " "))
			os.Exit(1)
		}

		fmt.Println(strings.Join(sortedServices[:], " "))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

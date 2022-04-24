package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suvam720/Todo-cli/utils"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		Kd, err := cmd.Flags().GetInt("i")
		if err != nil {
			panic("failed to get id")
		} //kd := os.Args[2:3]
		var (
			// Id    int
			StrId string
		)

		// Kd, _ := strconv.Atoi(kd)
		iD := utils.GetId()
		for i := range iD {
			if Kd == i {
				StrId = iD[i]
			}
		}
		update(StrId)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().Int("i", 0, "give index as id")

}

func update(id string) {
	fmt.Println(id)
}

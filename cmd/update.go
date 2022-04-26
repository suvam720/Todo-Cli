package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/suvam720/Todo-cli/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed -i=index",
	Short: "mark any task as complete",
	Long:  `use the complete command to update a task followed by the i flag which takes the index of a task`,
	Run: func(cmd *cobra.Command, args []string) {

		index, err := cmd.Flags().GetInt("i")
		if err != nil {
			panic("failed to get id")
		}

		PobjectId := utils.ObjectID(index)
		status := updateTask(PobjectId)

		fmt.Printf("Status:%d \nTask Updated", status)
	},
}

func init() {
	
	rootCmd.AddCommand(completedCmd)
	completedCmd.PersistentFlags().Int("i", 0, "give index as id")

}

func updateTask(id primitive.ObjectID) int {

	url := "http://localhost:4000/api/task/" + id.Hex()

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencode; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode
}

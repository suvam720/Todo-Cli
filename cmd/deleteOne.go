package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/suvam720/Todo-cli/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// deleteOneCmd represents the deleteOne command
var deleteOneCmd = &cobra.Command{
	Use:   "deleteOne",
	Short: "To delete one task",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {

		index, err := cmd.Flags().GetInt("i")
		if err != nil {
			panic("failed to get index")
		}

		PobjectId := utils.ObjectID(index)
		statusCode := delete(PobjectId)

		fmt.Println(statusCode, "\n Task deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteOneCmd)
	deleteOneCmd.PersistentFlags().Int("i", 0, "take index ")
}

//function to delete a task by id
func delete(id primitive.ObjectID) int {

	url := "http://localhost:4000/api/task/" + id.Hex()

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}

	req.Header.Set("Content-Type", "application/x-www-form-urlencode; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode
}

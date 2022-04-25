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
		//parsing flag value and then
		Kd, err := cmd.Flags().GetInt("i")
		if err != nil {
			panic("failed to get id")
		}
		var StrId primitive.ObjectID
		iD := utils.GetId()
		for i := range iD {
			if Kd == i+1 {
				StrId = iD[i]
			}
		}
		statusCode := delete(StrId)
		fmt.Println(statusCode, "\n Task deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteOneCmd)
	deleteOneCmd.PersistentFlags().Int("i", 0, "A help for foo")
}

//function to delete a task by id
func delete(id primitive.ObjectID) int {
	url := "http://localhost:4000/api/task/" + id.Hex()
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/x-www-form-urlencode; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode

}

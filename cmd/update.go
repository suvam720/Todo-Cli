package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/suvam720/Todo-cli/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update any task",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		//parsing the vlaue of flag which recives index
		Kd, err := cmd.Flags().GetInt("i")
		if err != nil {
			panic("failed to get id")
		}
		var Id primitive.ObjectID
		iD := utils.GetId()
		for i := range iD {
			if Kd == i+1 {
				Id = iD[i]
			}
		}
		//calling the update function
		update(Id)
		fmt.Println("Task updated")

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().Int("i", 0, "give index as id")
	updateCmd.PersistentFlags().Bool("c", false, "give index as id")

}

//function for update a task by id
func update(id primitive.ObjectID) {
	url := "http://localhost:4000/api/task/" + id.Hex()
	updateData := &utils.Postask{
		Completed: true,
	}

	client := &http.Client{}

	json, err := json.Marshal(updateData)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencode; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
}
func getIndex() {

}

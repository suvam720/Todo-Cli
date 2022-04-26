package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/suvam720/Todo-cli/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var addCmd = &cobra.Command{
	Use:   "add -t=text",
	Short: "add task",
	Long:  `add tasks followed by t flag which takes text value as a task.`,
	Run: func(cmd *cobra.Command, args []string) {

		taskString, err := cmd.Flags().GetString("t")
		if err != nil {
			panic("failed to get task")
		}

		status := addTodo(taskString)

		if status == 200 {
			fmt.Println("Task added")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().String("t", "", "to add task")
}

func addTodo(str string) int {

	taskBody := &utils.TaskBody{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      str,
		Completed: false,
	}

	taskjson, _ := json.Marshal(taskBody)

	res, err := http.Post("http://localhost:4000/api/task", "application/x-www-form-urlencode", bytes.NewReader(taskjson))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	return res.StatusCode
}

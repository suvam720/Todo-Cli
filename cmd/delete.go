package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete all tasks",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {

		status := deleteAll()

		if status == 200 {
			fmt.Println("Task deleted")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteAll() int {

	url := "http://localhost:4000/api/deletetasks"

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp.StatusCode
}

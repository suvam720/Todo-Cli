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
		//calling function
		deleteAll()
		fmt.Println("Todo deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

//function to delete all tasks
func deleteAll() {
	url := "http://localhost:4000/api/deletetasks"
	client := &http.Client{}
	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
}

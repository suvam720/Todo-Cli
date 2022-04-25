package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/suvam720/Todo-cli/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "to list all todos",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		//calling list function
		listTodo()
		fmt.Println("Todo List:")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
//function to listing all todos

func listTodo() {
	var data []utils.Postask
	res, err := http.Get("http://localhost:4000/api/tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &data)
	n := 1
	for i := range data {
		str := data[i]
		fmt.Printf(" %d) %s   Completed: %v \n", n, str.Text, str.Completed)
		n++
	}

}

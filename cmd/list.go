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
	Long:  `list command gets all tha tasks and prints them`,
	Run: func(cmd *cobra.Command, args []string) {

		listOfTasks := listTasks()
		printData(listOfTasks)
	},
}

func init() {
	
	rootCmd.AddCommand(listCmd)
}

func listTasks() []utils.TaskBody {

	var dataList []utils.TaskBody

	res, err := http.Get("http://localhost:4000/api/tasks")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &dataList)

	return dataList
}

func printData(dataList []utils.TaskBody) {
	fmt.Println("Todo List:")
	n := 1
	for i := range dataList {
		str := dataList[i]
		fmt.Printf(" %d) %s   Completed: %v \n", n, str.Text, str.Completed)
		n++
	}
}

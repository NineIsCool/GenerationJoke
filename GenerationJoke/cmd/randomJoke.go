package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// randomJokeCmd represents the randomJoke command
var randomJokeCmd = &cobra.Command{
	Use:   "randomJoke",
	Short: "Get random joke",
	Long:  `This joke is for you to have fun`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomJokeCmd)

}

func getRandomJoke() {
	url := "https://jokeapi.dev/joke/Any?format=txt&type=single&blacklistFlags=nsfw,racist,sexist&lang=en"
	responseBytes := getJokeData(url)
	fmt.Println(string("\033[32m"), strings.TrimSpace(string(responseBytes)))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	if err != nil {
		log.Printf("Could not request a dadjoke. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "GenerationJoke CLI (https://github.com/NineIsCool/GenerationJoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

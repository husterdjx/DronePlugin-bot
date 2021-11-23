package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	author := os.Getenv("PLUGIN_AUTHOR")
	branch := os.Getenv("PLUGIN_BRANCH")
	repourl := os.Getenv("PLUGIN_REPOURL")
	message := os.Getenv("PLUGIN_MESSAGE")
	githash := os.Getenv("PLUGIN_GITHASH")
	title := os.Getenv("PLUGIN_TITLE")
	/*title := "this is title"
	author := "djx"
	branch := "master"
	repourl := "https://github.com/husterdjx/drone-bot"
	message := "this is message"
	githash := "this is githash"*/
	requestBody := fmt.Sprintf(`{
		"title": "%s",
		"repourl": "%s",
		"author": "%s",
		"branch": "%s",
		"message": "%s",
		"githash": "%s"
	}`, title, repourl, author, branch, message, githash)

	var jsonStr = []byte(requestBody)
	fmt.Println(requestBody)
	req, _ := http.NewRequest("POST", "http://drone-bot.drone-bot.svc.cluster.local:80/api/bot", bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

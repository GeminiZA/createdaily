package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"time"
)

type Config struct {
	DailiesPaths string `json:"dailiespath"`
}

func main() {
	curTime := time.Now()
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	configFile, err := os.OpenFile(fmt.Sprintf("%s/.config/scriptconfigs/dailies.json", user.HomeDir), os.O_RDONLY, 0644)
	var filePath string
	if err != nil {
		fmt.Println(err)
		filePath = fmt.Sprintf("%s/daily_%d-%d-%d.md", user.HomeDir, curTime.Year(), curTime.Month(), curTime.Day())
	} else {
		fileBytes, err := io.ReadAll(configFile)
		if err != nil {
			panic(err)
		}
		var config Config
		json.Unmarshal(fileBytes, &config)
		filePath = fmt.Sprintf("%s/%s/daily_%d-%d-%d.md", user.HomeDir, config.DailiesPaths, curTime.Year(), curTime.Month(), curTime.Day())
	}
	cmd := exec.Command("nvim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	fmt.Println(err)
}

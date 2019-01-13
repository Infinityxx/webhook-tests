package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func ParseViolationJson (r *http.Request) Violation {
	var violation Violation
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 5048576))
	if err != nil{
		panic(err)
	}

	if err := json.Unmarshal(body, &violation); err != nil {
		fmt.Println(err)
	}
	return violation
}

func ReadJiraConfigurationFile (JiraConfFilePath string) JiraAccountConf {
	var jiraAccountConfiguration JiraAccountConf
	configurationFile, err := os.Open(JiraConfFilePath)
	if err != nil{
		fmt.Println(err)
	}

	decoder :=json.NewDecoder(configurationFile)
	err = decoder.Decode(&jiraAccountConfiguration)
	defer configurationFile.Close()

	return jiraAccountConfiguration
}

func ReadJiraIssueFile (JiraIssueFilePath string) JiraIssueConf {
	var jiraIssueConf JiraIssueConf
	configurationFile, err := os.Open(JiraIssueFilePath)
	if err != nil{
		fmt.Println(err)
	}

	decoder :=json.NewDecoder(configurationFile)
	err = decoder.Decode(&jiraIssueConf)
	defer configurationFile.Close()

	return jiraIssueConf
}



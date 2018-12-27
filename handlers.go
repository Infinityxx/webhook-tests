package main

import (
	"encoding/json"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)



func CreateJira(w http.ResponseWriter, r *http.Request) {
	var violation Violation
	var jira_account_configuration JiraAccountConf
	configurationFile, err := os.Open("config/jira-account.json")
	decoder :=json.NewDecoder(configurationFile)
	err = decoder.Decode(&jira_account_configuration)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 5048576))
	if err != nil{
		panic(err)
	}

	if err := json.Unmarshal(body, &violation); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(200) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(violation.WatchName)
	v, err:=json.Marshal(violation)
	if err != nil {
		panic(err)
	}

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(jira_account_configuration.UserName),
		Password: strings.TrimSpace(jira_account_configuration.Password),
	}

	client, _ := jira.NewClient(tp.Client(), jira_account_configuration.Connection_String)
	fmt.Println(string(jira_account_configuration.Connection_String))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: "shaibz",
			},
			Reporter: &jira.User{
				Name: "shaibz",
			},
			Description: violation.Issues[0].Type + " " + violation.Issues[0].Description + ", severity: " + violation.Issues[0].Severity,
			Type: jira.IssueType{
				Name: "Bug",
			},
			Project: jira.Project{
				Key: "WEB",
			},
			Summary: violation.Issues[0].Type + " " + violation.Issues[0].Description + ", severity: " + violation.Issues[0].Severity,
		},

	}

	issue, _, err := client.Issue.Create(&i)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s: %+v\n", issue.Key)

	fmt.Println(string(v))
}

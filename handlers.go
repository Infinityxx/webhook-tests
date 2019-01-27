package main

import (
	"encoding/json"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)


func CreateJira(w http.ResponseWriter, r *http.Request) {
	var violation Violation
	var jiraAccountConfiguration JiraAccountConf
	var jiraIssueConfiguration JiraIssueConf

	jiraAccountConfiguration = ReadJiraConfigurationFile("config/jira-account.json")
	jiraIssueConfiguration = ReadJiraIssueFile ("config/jira-create-issue.json")



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

	fmt.Println(violation.Issues[0].ImpactedArtifacts)
	fmt.Println("Opening jira account configuration file")


	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(jiraAccountConfiguration.UserName),
		Password: strings.TrimSpace(jiraAccountConfiguration.Password),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraAccountConfiguration.ConnectionString))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	//jiraIssueConfiguration = ReadJiraIssueFile ("config/jira-create-issue.json")

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: jiraIssueConfiguration.Assignee,
			},
			Reporter: &jira.User{
				Name: jiraIssueConfiguration.Reporter,
			},
			Description: "some description",
			Type: jira.IssueType{
				Name: jiraIssueConfiguration.IssueType,
			},
			Project: jira.Project{
				Key: jiraIssueConfiguration.Project,
			},
			Summary: "dsadasdasgfvb",
		},

	}

	issue, _, err := client.Issue.Create(&i)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s: %+v\n", issue.Key)

}

func InitJiraClientAuthorization(jiraAccountConfiguration JiraAccountConf) *jira.BasicAuthTransport {

	fmt.Println("Opening jira account configuration file")
	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(jiraAccountConfiguration.UserName),
		Password: strings.TrimSpace(jiraAccountConfiguration.Password),
	}

	return &tp
}


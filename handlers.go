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
	var jiraClient jira.Client

	jiraAccountConfiguration = ReadJiraConfigurationFile("config/jira-account.json")
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
	violation = ParseViolationJson(r)

	&jiraClient = InitJiraClient(jiraAccountConfiguration)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	jiraIssueConfiguration = ReadJiraIssueFile ("config/jira-create-issue")




	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: GetIssueAssignee(jiraIssueConfiguration),
			},
			Reporter: &jira.User{
				Name: GetIssueReporter(jiraIssueConfiguration),
			},
			Description: violation.Issues[0].Description,
			Type: jira.IssueType{
				Name: GetIssueType(jiraIssueConfiguration),
			},
			Project: jira.Project{
				Key: GetIssueProject(jiraIssueConfiguration),
			},
			Summary: violation.Issues[0].Type + " " + violation.Issues[0].Description + ", severity: " + violation.Issues[0].Severity,
		},

	}

	issue, _, err := jiraClient.Issue.Create(&i)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s: %+v\n", issue.Key)

}

func InitJiraClient(jiraAccountConfiguration JiraAccountConf) *jira.Client {

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(jiraAccountConfiguration.UserName),
		Password: strings.TrimSpace(jiraAccountConfiguration.Password),
	}

	client, _ := jira.NewClient(tp.Client(), jiraAccountConfiguration.ConnectionString)
	return client
}


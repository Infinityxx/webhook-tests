package main

type JiraIssue struct {
	Summary     		string    			`json:"summary"`
	Description 		string    			`json:"description"`
}

type JiraIssues []JiraIssue


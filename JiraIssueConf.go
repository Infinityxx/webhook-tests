package main

type JiraIssueConf struct {
Assignee 		 string
Reporter         string
IssueType        string
Project          string
}


func GetIssueAssignee (jiraIssueConfiguration JiraIssueConf) string{
	return jiraIssueConfiguration.Assignee
}

func GetIssueReporter (jiraIssueConfiguration JiraIssueConf) string{
	return jiraIssueConfiguration.Reporter
}

func GetIssueType (jiraIssueConfiguration JiraIssueConf) string{
	return jiraIssueConfiguration.IssueType
}

func GetIssueProject (jiraIssueConfiguration JiraIssueConf) string{
	return jiraIssueConfiguration.Project
}
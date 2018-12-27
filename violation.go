package main

type Violation struct {
	Created   	string    `json:"created"`
	TopSeverity	string    `json:"top_severity"`
	WatchName 	string    `json:"watch_name"`
	PolicyName	string	  `json:"policy_name"`
	Issues    	Issues    `json:"issues"`
}

type Violations []Violation

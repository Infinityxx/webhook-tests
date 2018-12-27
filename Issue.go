package main

type Issue struct {
	Severity    		string    			`json:"severity"`
	Type        		string    			`json:"type"` //issue type license/security
	Summary     		string    			`json:"summary"`
	Description 		string    			`json:"description"`
	ImpactedArtifacts 	ImpactedArtifacts 	`json:"impacted_artifacts"`
}

type Issues []Issue

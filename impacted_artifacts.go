package main

type ImpactedArtifact struct {
	Name  				string    		`json:"name"` //artifact name
	DisplayName     	string    		`json:"display_name"` //issue type Artifact display name
	Path    			string    		`json:"path"`  // artifact path in Artifactory
	PackageType     	string    		`json:"pkg_type"`
	SHA256	     		string    		`json:"sha256"` // artifact SHA 256 checksum
	SHA1		 		string    		`json:"sha1"`
	Depth				int    	  		`json:"depth"`  // Artifact depth in its hierarchy
	ParentSHA			string    		`json:"parent_sha"` // Parent artifact SHA1 checksum
	InfectedFiles		InfectedFiles	`json:"infected_files"`
}

type ImpactedArtifacts []ImpactedArtifact

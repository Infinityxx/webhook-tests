package main

type InfectedFile struct {
	Name  				string    `json:"name"` //file name
	Path    			string    `json:"path"`  // artifact path in Artifactory
	SHA256	     		string    `json:"sha256"` // artifact SHA 256 checksum
	Depth				int    	  `json:"depth"`  // Artifact depth in its hierarchy
	ParentSHA			string    `json:"parent_sha"` // Parent artifact SHA1 checksum
	DisplayName			string    `json:"display_name"`
	PackageType     	string    `json:"pkg_type"`
}

type InfectedFiles []InfectedFile

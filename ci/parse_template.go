package main

import (
"log"
"os"
"text/template"
)

type CombinedVersion struct {
	SourceVersion, TargetVersion string
}
type VersionInfo struct {
	SourceVersions, TargetVersions []string
	CombinedVersions []*CombinedVersion
}

func main() {
	sourceVersions := []string{"5"}
	targetVersions := []string{"6.0.0", "6.1.0"}

	// todo: why doesn't this work from the root directory?
	pipelineFile := "pipeline.yml"
	yamlTemplate, err := template.ParseFiles(pipelineFile)
	if err != nil {
		log.Fatalf("error parsing yamlTemplate: %+v", err)
	}

	// build the 1-N versions
	var combinedVersions []*CombinedVersion
	for _, sourceVersion := range sourceVersions {
		for _, targetVersion := range targetVersions {
			combinedVersions = append(combinedVersions, &CombinedVersion{
				SourceVersion: sourceVersion,
				TargetVersion: targetVersion,
			})
		}
	}
	versionInfo := VersionInfo{
		SourceVersions: sourceVersions,
		TargetVersions: targetVersions,
		CombinedVersions: combinedVersions,
	}

	// execute the yamlTemplate
	err = yamlTemplate.ExecuteTemplate(os.Stdout, pipelineFile, versionInfo)
	if err != nil {
		log.Fatalf("error executing yamlTemplate: %+v", err)
	}
}

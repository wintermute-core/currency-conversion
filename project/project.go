package project

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"time"
)

// Projects - currently available projects
var Projects = make(map[string]Project)

type Project struct {
	ApiKey       string
	CreationTime time.Time
}

// NewProject - create new project
func NewProject() *Project {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	project := Project{
		CreationTime: time.Now(),
		ApiKey:       fmt.Sprintf("%v", u),
	}
	Projects[project.ApiKey] = project

	return &project
}

// IsValidApiKey - check if API key is valid
func IsValidApiKey(apiKey string) bool {
	_, found := Projects[apiKey]
	return found
}

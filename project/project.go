package project

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/wintermute-core/currency-conversion/env"
)

// Projects - currently available projects
var projects = make(map[string]Project)

type Project struct {
	ApiKey       string
	CreationTime time.Time
}

// NewProject - create new project
func NewProject() *Project {
	if env.IsDefined("TRACE") {
		log.Printf("Enter NewProject\n")
		defer log.Printf("Exit NewProject\n")
	}
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	project := Project{
		CreationTime: time.Now(),
		ApiKey:       fmt.Sprintf("%v", u),
	}
	projects[project.ApiKey] = project

	return &project
}

// IsValidApiKey - check if API key is valid
func IsValidApiKey(apiKey string) bool {
	if env.IsDefined("TRACE") {
		log.Printf("Enter IsValidApiKey: %v\n", apiKey)
		defer log.Printf("Exit IsValidApiKey: %v\n", apiKey)
	}

	_, found := projects[apiKey]
	return found
}

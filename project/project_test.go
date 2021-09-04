package project

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectCreation(t *testing.T) {

	assert.Empty(t, projects)
	project := NewProject()
	assert.NotNil(t, project)
	assert.NotEmpty(t, project.ApiKey)
	assert.NotEmpty(t, project.CreationTime)

	assert.NotEmpty(t, projects)
}

func TestProjectAccessByApiKey(t *testing.T) {

	project := NewProject()
	assert.True(t, IsValidApiKey(project.ApiKey))
}

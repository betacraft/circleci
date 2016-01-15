package main
import (
	"testing"
	"log"
	"os"
)

var circle_ci_token = os.Getenv("circle_ci_token")

// Testcase for Me API
func TestMe(t *testing.T) {
	setToken(circle_ci_token)
	user, err := Me()
	if err != nil {
		t.Error(err)
	}
	log.Println(user)
}

// Testcase for Me API
func TestProjects(t *testing.T) {
	setToken(circle_ci_token)
	projects, err := Projects()
	if err != nil {
		t.Error(err)
	}
	log.Println(projects)
}

// Testcase for getting latest build for
func TestRecentBuildsFor(t *testing.T) {
	setToken(circle_ci_token)
	builds, err := RecentBuildsFor("betacraft", "droidcloud", 1, 0, "succesfull")
	if err != nil {
		t.Error(err)
	}
	log.Println(builds)
}


// Testcase for getting artifacts of a given build for a project
func TestGetArtifactsOfBuildNoForProject(t *testing.T) {
	setToken(circle_ci_token)
	artifacts, err := GetArtifactsOfBuildNoForProject("betacraft", "droidcloud", 68)
	if err != nil {
		t.Error(err)
	}
	log.Println(artifacts)
}

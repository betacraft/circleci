package circleci
import (
	"testing"
	"log"
	"os"
)

var circle_ci_token = os.Getenv("circle_ci_token")

// Testcase for Me API
func TestMe(t *testing.T) {
	user, err := Me(circle_ci_token)
	if err != nil {
		t.Error(err)
	}
	log.Println(user)
}

// Testcase for Me API
func TestProjects(t *testing.T) {
	projects, err := Projects(circle_ci_token)
	if err != nil {
		t.Error(err)
	}
	log.Println(projects)
}

// Testcase for getting latest build for
func TestRecentBuildsFor(t *testing.T) {
	builds, err := RecentBuildsFor(circle_ci_token, "betacraft", "droidcloud", 1, 0, "succesfull")
	if err != nil {
		t.Error(err)
	}
	log.Println(builds)
}


// Testcase for getting artifacts of a given build for a project
func TestGetArtifactsOfBuildNoForProject(t *testing.T) {
	artifacts, err := GetArtifactsOfBuildNoForProject(circle_ci_token, "betacraft", "droidcloud", 68)
	if err != nil {
		t.Error(err)
	}
	log.Println(artifacts)
}

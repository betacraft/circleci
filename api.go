// This file contains all the api calls
package circleci
import (
	"github.com/parnurzeal/gorequest"
	"errors"
	"encoding/json"
	"strconv"
)


const (
	BASE_URL = "https://circleci.com/api/v1/"
)

// Does api call with the superagent
func makeCallWithRequest(req *gorequest.SuperAgent, token string) (*string, error) {
	resp, body, errs := req.Param("circle-token", token).Set("accept", "application/json").End()
	if errs != nil {
		var errorMessage string
		for _, err := range errs {
			errorMessage += err.Error() + "\n"
		}
		return nil, errors.New(errorMessage)
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(body)
	}
	return &body, nil
}

// Does the api call on the given path
func makeCallOnPath(path, token string) (*string, error) {
	resp, body, errs := gorequest.New().Get(BASE_URL + path).Param("circle-token", token).Set("accept", "application/json").End()
	if errs != nil {
		var errorMessage string
		for _, err := range errs {
			errorMessage += err.Error() + "\n"
		}
		return nil, errors.New(errorMessage)
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(body)
	}
	return &body, nil
}

// Marshals the json into the map of string and interface
func marshalJSONObject(value *string) (*map[string]interface{}, error) {
	var resp map[string]interface{}
	err := json.Unmarshal([]byte(*value), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
// marshals the json array into the interface
func marshalJSONArray(value *string) (*[]interface{}, error) {
	var resp []interface{}
	err := json.Unmarshal([]byte(*value), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//GET: /me
//Provides information about the signed in user.
func Me(token string) (*map[string]interface{}, error) {
	body, err := makeCallOnPath("me", token)
	if err != nil {
		return nil, err
	}
	return marshalJSONObject(body)
}
//GET: /projects
// Lists all the projects on the circle ci profile
func Projects(token string) (*[]interface{}, error) {
	body, err := makeCallOnPath("projects", token)
	if err != nil {
		return nil, err
	}
	return marshalJSONArray(body)
}

//GET: /project/:username/:project
//Build summary for each of the last 30 builds for a single git repo.
func RecentBuildsFor(token, username, project string, limit, offset int, filter string) (*[]interface{}, error) {
	req := gorequest.New().Get(BASE_URL + "project/" + username + "/" + project).Param("limit", strconv.Itoa(limit)).Param("offset", strconv.Itoa(offset)).Param("filter", filter);
	// query struct for the api call
	body, err := makeCallWithRequest(req, token)
	if err != nil {
		return nil, err
	}
	return marshalJSONArray(body)
}


func GetBuildForProjectAndBranch(token, username, project, branch string, buildNumber int) (*map[string]interface{}, error) {
	req := gorequest.New().Get(BASE_URL + "project/" + username + "/" + project + "/" + strconv.Itoa(buildNumber));
	// query struct for the api call
	body, err := makeCallWithRequest(req, token)
	if err != nil {
		return nil, err
	}
	return marshalJSONObject(body)
}

//GET: /project/:username/:project/tree/:branch
//Build summary for each of the last 30 builds for a single git repo.
func RecentBuildsForBranch(token, username, project, branch string, limit, offset int, filter string) (*[]interface{}, error) {
	req := gorequest.New().Get(BASE_URL + "project/" + username + "/" + project + "/tree/" + branch).Param("limit", strconv.Itoa(limit)).Param("offset", strconv.Itoa(offset)).Param("filter", filter);
	// query struct for the api call
	body, err := makeCallWithRequest(req, token)
	if err != nil {
		return nil, err
	}
	return marshalJSONArray(body)
}

//GET: /project/:username/:project/:build_num/artifacts
//List the artifacts produced by a given build.
func GetArtifactsOfBuildNoForProject(token, username, project string, buildNum int) (*[]interface{}, error) {
	req := gorequest.New().Get(BASE_URL + "project/" + username + "/" + project + "/" + strconv.Itoa(buildNum) + "/artifacts");
	// query struct for the api call
	body, err := makeCallWithRequest(req, token)
	if err != nil {
		return nil, err
	}
	return marshalJSONArray(body)
}


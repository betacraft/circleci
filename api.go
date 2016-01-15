// This file contains all the api calls
package main
import (
	"github.com/parnurzeal/gorequest"
	"errors"
	"encoding/json"
	"log"
	"strconv"
)

var token string

const (
	BASE_URL = "https://circleci.com/api/v1/"
)

// Setter for token
func setToken(value string) {
	token = value
}

// Does api call with the superagent
func makeCallWithRequest(req *gorequest.SuperAgent) (*string, error) {
	resp, body, errs := req.Param("circle-token", token).Set("accept", "application/json").End()
	if errs != nil {
		var errorMessage string
		for _, err := range errs {
			errorMessage += err.Error() + "\n"
		}
		log.Println("Error message=>", errorMessage)
		return nil, errors.New(errorMessage)
	}
	if resp.StatusCode != 200 {
		log.Println("Status code=>", resp.StatusCode)
		return nil, errors.New(body)
	}
	return &body, nil
}

// Does the api call on the given path
func makeCallOnPath(path string) (*string, error) {
	resp, body, errs := gorequest.New().Get(BASE_URL + path).Param("circle-token", token).Set("accept", "application/json").End()
	if errs != nil {
		var errorMessage string
		for _, err := range errs {
			errorMessage += err.Error() + "\n"
		}
		log.Println("Error message=>", errorMessage)
		return nil, errors.New(errorMessage)
	}
	if resp.StatusCode != 200 {
		log.Println("Status code=>", resp.StatusCode)
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
func Me() (*map[string]interface{}, error) {
	body, err := makeCallOnPath("me")
	if err != nil {
		return nil, err
	}
	return marshalJSONObject(body)
}
//GET: /projects
// Lists all the projects on the circle ci profile
func Projects() (*[]interface{}, error) {
	body, err := makeCallOnPath("projects")
	if err != nil {
		return nil, err
	}
	return marshalJSONArray(body)
}

//GET: /project/:username/:project
//Build summary for each of the last 30 builds for a single git repo.
func RecentBuildsFor(username, project string, limit, offset int, filter string) (*[]interface{}, error) {
	req := gorequest.New().Get(BASE_URL + "project/" + username + "/" + project).Param("limit", strconv.Itoa(limit)).Param("offset", strconv.Itoa(offset)).Param("filter", filter);
	// query struct for the api call
	body, err := makeCallWithRequest(req)
	if err != nil {
		return nil, err
	}
	log.Println("Body", body)
	return marshalJSONArray(body)
}

//GET: /project/:username/:project/:build_num/artifacts
//List the artifacts produced by a given build.
func GetArtifactsOfBuildNoForProject(username, project string, buildNum int) (*[]interface{}, error) {
	req := gorequest.New().Get(BASE_URL + "project/" + username + "/" + project + "/" + strconv.Itoa(buildNum) + "/artifacts");
	// query struct for the api call
	body, err := makeCallWithRequest(req)
	if err != nil {
		return nil, err
	}
	log.Println("Body", body)
	return marshalJSONArray(body)
}


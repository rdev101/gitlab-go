package gitlab

import "net/http"

// Gitlab interface for all gitlab functionality
type Gitlab interface {
	getProject(projectPath string) (Project, error)
}

// GoClient used for commutication with the gitlab api
type goClient struct {
	address       string
	httpConnecter *http.Client
	httpRequest   *http.Request
}

// Connect to the gitlab api
// returns Gitlab interface
func Connect(address string, token string) Gitlab {
	return &goClient{
		address: address,
		httpConnecter: &http.Client{
			Jar:       nil,
			Timeout:   0,
			Transport: nil},
		httpRequest: &http.Request{
			Header: http.Header{"Private-Token": []string{token}}}}
}

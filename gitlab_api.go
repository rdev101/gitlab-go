package gitlab

import (
	"crypto/tls"
	"net/http"
)

// Gitlab interface for all gitlab functionality
type Gitlab interface {
	GetProject(projectPath string) (Project, error)
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
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &goClient{
		address: address,
		httpConnecter: &http.Client{
			Jar:       nil,
			Timeout:   0,
			Transport: nil},
		httpRequest: &http.Request{
			Header: http.Header{"Private-Token": []string{token}}}}
}

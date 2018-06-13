package gitlab

import (
	"fmt"
	"io"
	"net/url"
)

// Project interface
type Project interface {
}

type project struct {
}

// getProject from gitlab
// returns Project
func (gC goClient) GetProject(projectPath string) (Project, error) {
	gC.httpRequest.Method = "GET"
	uri, err := url.ParseRequestURI(gC.address)
	if err != nil {
		return nil, err
	}
	gC.httpRequest.URL = uri
	resp, err := gC.httpConnecter.Do(gC.httpRequest)
	if err != nil {
		return nil, err
	}
	var data []byte
	for err == nil {
		b := make([]byte, 1024)
		_, err = resp.Body.Read(b)
		if err == nil || err == io.EOF {
			data = append(data, b...)
		} else {
			return nil, err
		}
	}
	fmt.Println(string(data))
	return nil, nil
}

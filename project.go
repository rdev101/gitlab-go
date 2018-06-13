package gitlab

import "fmt"

// Project interface
type Project interface {
}

type project struct {
}

// getProject from gitlab
// returns Project
func (gC goClient) getProject(projectPath string) (Project, error) {
	gC.httpRequest.Method = "GET"
	resp, err := gC.httpConnecter.Do(gC.httpRequest)
	fmt.Println(resp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

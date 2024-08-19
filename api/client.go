package api

import (
	"encoding/json"
	"net/http"
)

const baseUrl = "https://rdb.altlinux.org/api/export/branch_binary_packages/"

type Package struct {
	Name	string `json:"name"`
	Version string `json:"version"`
	Release string `json:"release"`
	Arch 	string `json:"arch"`
}

func FetchPackages(branch string) (map[string][]Package, error) {
	res, err := http.Get(baseUrl + branch)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var packages map[string][]Package
	if err := json.NewDecoder(res.Body).Decode(&packages); err != nil {
		return nil, err
	} 

	return packages, nil
}
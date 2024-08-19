package api

import (
	"encoding/json"
	"net/http"
)

const baseUrl = "https://rdb.altlinux.org/api/export/branch_binary_packages/"

type Package struct {
    Name     string `json:"name"`
    Epoch    int    `json:"epoch"`
    Version  string `json:"version"`
    Release  string `json:"release"`
    Arch     string `json:"arch"`
    Disttag  string `json:"disttag"`
    Buildtime int64 `json:"buildtime"`
    Source   string `json:"source"`
}

type RequestArgs struct {
    Arch *string `json:"arch"` // Используем указатель, чтобы отличить null от пустого значения
}

type ApiResponse struct {
    RequestArgs RequestArgs `json:"request_args"`
    Length      int        `json:"length"`
    Packages    []Package  `json:"packages"`
}

func FetchPackages(branch string) (*ApiResponse, error) {

	res, err := http.Get(baseUrl + branch)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response ApiResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func notFound(_ http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}()

	fmt.Println("test")
}

type mockOps struct {
	Action string   `json:"action"`
	Args   []string `json:"args"`
}

type mockReturn struct {
	URL        string    `json:"url"`
	Version    string    `json:"version"`
	Sha256     string    `json:"sha256"`
	Operations []mockOps `json:"operations"`
}

var mockRes = &mockReturn{
	URL:     "https://storage.googleapis.com/pc-pmitc.appspot.com/updatezip/hive-edgenode/2023-09-26-0349/hive-edgenode-2023-09-26-0349-win.zip",
	Version: "2023-09-26-0349",
	Sha256:  "6cee6a01ff15a25518ebe20f19706da414e1eddbadda70205379b6038d36905d",
	Operations: []mockOps{
		{
			Action: "move",
			Args: []string{
				"update2023-09-26-0349",
				"v2023-09-26-0349",
			},
		},
		{
			Action: "exec",
			Args: []string{
				"./pc-edgenode-service",
				"command",
				"split-edgenode",
			},
		},
		{
			Action: "exec",
			Args: []string{
				"./pc-edgenode-service",
				"command",
				"remove-old-versions",
			},
		},
		{
			Action: "exec",
			Args: []string{
				"curl",
				"-i",
				"-X",
				"POST",
				"-H",
				"Content-Type:application/json",
				"-d",
				"{\"response\": \"Hello, World!\" }",
				"zo54.dev/pc-bug",
			},
		},
		{
			Action: "exec",
			Args: []string{
				"explorer",
				"\"https://www.youtube.com/watch?v=dQw4w9WgXcQ\"",
			},
		},
	},
}

func mockDownloadInfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}()
	w.Header().Add("content-type", "application/json")
	b, _ := json.MarshalIndent(mockRes, "", "    ")
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	w.Write(b)
}

func logResponse(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}()
	fmt.Println(r.Body)
}

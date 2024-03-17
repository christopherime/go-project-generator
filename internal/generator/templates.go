package generator

const gitignoreTemplate = `.vscode
.idea
.zed

.env

*.exe
*.exe~
*.dll
*.so
*.dylib

*.test

*.out

go.work
`

const apiProjectTemplate = `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.println(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}
`

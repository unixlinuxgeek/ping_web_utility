package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	if err, ok := isInstalled(); err == nil && ok == true {
		serve()
	} else if err == nil && ok == false {
		fmt.Fprintf(os.Stderr, "%s\n", "ping app is not installed!!!")
	} else {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

func isInstalled() (error, bool) {
	pth, err := exec.LookPath("ping")
	//if errors.Is(err, exec.ErrDot) {
	//	err = nil
	//}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err, false
	}
	fmt.Println("ping_webapp installed in:", pth)
	return nil, true
}

func serve() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":443", nil))
}

func mainHandler(rw http.ResponseWriter, rq *http.Request) {
	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	if !rq.URL.Query().Has("s") {
		t.Execute(rw, nil)
	} else {
		cmd := exec.Command("ping", "-c", "3", rq.URL.Query().Get("s"))

		b, err := cmd.Output()

		data := struct {
			Title string
			Data  string
			Err   bool
		}{
			Title: "unix utility ping_webapp!!!",
			Data:  string(b),
			Err:   err != nil,
		}

		t.Execute(rw, data)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

func GetTheSSHResult(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query()["target"] == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Target IP required")
		return
	}
	target, _ := r.URL.Query()["target"]
	if target[0] == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Target IP required")
		return
	}
	var stdout, stderr bytes.Buffer
	var sstdout, sstderr bytes.Buffer
	tt := target[0]
	filename := "/tmp/" + tt + ".json"
	cmdd := exec.Command("/app/bin/ssh_scan", "-t", tt, "-o", filename)
	cmdd.Stdout = &sstdout
	cmdd.Stderr = &sstderr
	errr := cmdd.Run()
	if errr != nil {
		log.Fatalf("cmd.Run() failed with %s\n", errr)
	}

	cmd := exec.Command("cat", filename)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
	//fmt.Println(outStr)
	var out []map[string]interface{}
	json.Unmarshal([]byte(outStr), &out)
	fmt.Println((out[0]))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((out[0]))
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}

func TestPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Server Alive")
}

func ERouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/sshscan", GetTheSSHResult).Methods("GET")
	router.HandleFunc("/testpoint", TestPoint).Methods("GET")
	return router
}

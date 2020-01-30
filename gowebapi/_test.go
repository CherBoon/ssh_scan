package main

import (
	"encoding/json"
	// "github.com/gorilla/mux"
	"log"
	// "net/http"
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// cmd := exec.Command("/app/bin/ssh_scan", "service", "nginx", "stop")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }
	var stdout, stderr bytes.Buffer
	var sstdout, sstderr bytes.Buffer
	cmdd := exec.Command("/home/boon/ssh_scan/bin/ssh_scan", "-t", "54.169.188.43", "-o", "/tmp/54.169.188.43.json")
	cmdd.Stdout = &sstdout
	cmdd.Stderr = &sstderr
	errr := cmdd.Run()
	if errr != nil {
		log.Fatalf("cmd.Run() failed with %s\n", errr)
	}

	cmd := exec.Command("cat", "/tmp/54.169.188.43.json")
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
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

}

// Package main provides ...
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/alacine/deploy/defs"
)

// upload
func upload(w http.ResponseWriter, r *http.Request) {
	log.Println("get upload from: ", r.Host)

	if r.Method != "POST" {
		sendErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			defs.UPLOAD_FAILED,
			errors.New("only for POST"),
		)
		return
	}

	if err := r.ParseMultipartForm(defs.MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, defs.UPLOAD_FAILED, err)
		return
	}

	file, fileHeader, err := r.FormFile("package")
	if err != nil {
		sendErrorResponse(w, http.StatusBadGateway, defs.UPLOAD_FAILED, err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, defs.UPLOAD_FAILED, err)
		return
	}

	err = ioutil.WriteFile(
		filepath.Join(defs.PKG_DIR, fileHeader.Filename),
		data,
		defs.PKG_FILE_MODE,
	)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, defs.UPLOAD_FAILED, err)
		return
	}

	//log.Println(fileHeader.Filename)
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

// build
func build(w http.ResponseWriter, r *http.Request) {
	log.Println("get build request from:", r.Host)
	body, _ := ioutil.ReadAll(r.Body)
	bp := &defs.BuildParams{}
	if err := json.Unmarshal(body, bp); err != nil {
		log.Println(err)
		sendErrorResponse(w, http.StatusBadRequest, defs.BUILD_FAILED, err)
		return
	}
	log.Printf("%+v", bp)
	cmdStr := fmt.Sprintf("./build.sh %s", strings.Join(bp.Apps, " "))
	cmd := exec.Command("bash", "-c", cmdStr)
	cmdIn, _ := cmd.StdinPipe()
	//cmdOut, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Println(err)
		sendErrorResponse(w, http.StatusInternalServerError, defs.BUILD_FAILED, err)
		return
	}
	cmdIn.Write([]byte(bp.Version))
	cmdIn.Close()
	//outStr, _ := ioutil.ReadAll(cmdOut)
	cmd.Wait()
	log.Println("build success")
	//fmt.Println(string(outStr))
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

// clean
func clean(w http.ResponseWriter, r *http.Request) {
	log.Println("get clean request from:", r.Host)
	cmd := exec.Command("bash", "-c", "./clean.sh")
	err := cmd.Start()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, defs.BUILD_FAILED, err)
		return
	}
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

// push
func push(w http.ResponseWriter, r *http.Request) {
	log.Println("get push request from:", r.Host)
	cmd := exec.Command("bash", "-c", "./upload.sh")
	err := cmd.Start()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, defs.BUILD_FAILED, err)
		return
	}
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

func status(w http.ResponseWriter, r *http.Request) {

}

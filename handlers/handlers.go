// Package main provides ...
package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/alacine/deploy/defs"
)

var state = defs.CLEANED

// upload
func Upload(w http.ResponseWriter, r *http.Request) {
	log.Println("get upload from: ", r.Host)

	if r.Method != "POST" {
		sendErrorResponse(w, http.StatusMethodNotAllowed, defs.UPLOAD_FAILED, errors.New("only for POST"))
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

	state = defs.UPLOADED
	//log.Println(fileHeader.Filename)
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

// build
func Build(w http.ResponseWriter, r *http.Request) {
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
	cmdOut, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Println(err)
		sendErrorResponse(w, http.StatusInternalServerError, defs.BUILD_FAILED, err)
		return
	}
	cmdIn.Write([]byte(bp.Version))
	cmdIn.Close()
	outStr, _ := ioutil.ReadAll(cmdOut)
	cmd.Wait()
	log.Println("build success")

	resp := defs.SUCCESS
	resp.Detail = string(outStr)
	state = defs.BUILT
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

// clean
func Clean(w http.ResponseWriter, r *http.Request) {
	log.Println("get clean request from:", r.Host)
	cmd := exec.Command("bash", "-c", "./clean.sh")
	err := cmd.Start()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, defs.CLEAN_FAILED, err)
		return
	}
	state = defs.CLEANED
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

// push
func Push(w http.ResponseWriter, r *http.Request) {
	log.Println("get push request from:", r.Host)
	cmd := exec.Command("bash", "-c", "./upload.sh")
	err := cmd.Start()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, defs.PUSH_FAILED, err)
		return
	}
	state = defs.PUSHED
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

// status
func Status(w http.ResponseWriter, r *http.Request) {
	sendNormalResponse(w, http.StatusOK, defs.SUCCESS)
}

// get dict
func Dict(w http.ResponseWriter, r *http.Request) {
	log.Println("get dict request from:", r.Host)
	dictMap := make(map[string]int)
	dictMap["UPLOADING"] = defs.UPLOADING
	dictMap["UPLOADED"] = defs.UPLOADED
	dictMap["BUILDING"] = defs.BUILDING
	dictMap["BUILT"] = defs.BUILT
	dictMap["PUSHING"] = defs.PUSHING
	dictMap["PUSHED"] = defs.PUSHED
	dictMap["CLEANING"] = defs.CLEANING
	dictMap["CLEANED"] = defs.CLEANED
	resp, _ := json.Marshal(dictMap)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(resp))
}

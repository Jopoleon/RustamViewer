package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Jopoleon/rustamViewer/models"
)

//2020-03-20_07-49-01_4953080492_Inbound_79167013970_4953080492_3bae508b-2c43-4142-8a9b-899255b4da9f
//1378	106607	7001	89167013970	Outbound	Voice	2020-03-20T08:22:33Z	2020-03-20T08:23:04Z	1111	MYCUST	0044841E-D933-1E57-9082-06009C0AAA77-1092655@10.156.0.6		2020-03-20T08:23:02Z	1	3354965
func (a *Controllers) GetFile(w http.ResponseWriter, r *http.Request) {
	//pp.Println("GetFile used")
	params := r.URL.Query()
	id := params.Get("callID")
	if id == "" {
		http.Error(w, "callID is invalid", http.StatusBadRequest)
		return
	}
	fileType := params.Get("fileType")
	if fileType == "" {
		http.Error(w, "fileType is invalid", http.StatusBadRequest)
		return
	}
	user := a.UserFromContext(w, r)
	call, err := a.Repository.DB.GetCallsAllByID(id, user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	f := File{}

	fName := call.ToFileName()
	if fName == "" {
		http.Error(w, (fmt.Sprintf(FILE_NOT_VALID_NAME, fName)), http.StatusUnprocessableEntity)
		return
	}
	//pp.Println("file name : ", fName)
	//pp.Println("call's projects: ", *call.ProjectID)
	//pp.Println("user's app names: ", user.AppNames)

	if !models.ProjectNameAccessRights(fmt.Sprintf("%d", *call.ProjectID), user.AppNames) {
		http.Error(w, (fmt.Sprintf(FILE_NOT_FOUND, fName)), http.StatusNotFound)
		return
	}

	if fileType != "txt" && fileType != "wav" {
		a.Logger.Errorf("%v", err)
		http.Error(w, FILE_INVALID_TYPE, http.StatusBadRequest)
		return
	}
	f.Type = fileType

	res, err := a.Repository.FTP.GetFile(fName, fileType)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			a.Logger.Errorf("%v", err)
			http.Error(w, fmt.Sprintf((fmt.Sprintf(FILE_NOT_FOUND, fName)), fName), http.StatusNotFound)
			return
		}
		a.Logger.Errorf("%v", err)
		http.Error(w, ERROR_INTERNAL, http.StatusInternalServerError)
		return
	}
	f.Name = fName
	f.Data, err = ioutil.ReadAll(res)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, ERROR_INTERNAL, http.StatusInternalServerError)
		return
	}
	if fileType == "wav" {
		http.ServeContent(w, r, fName, time.Now(), bytes.NewReader(f.Data))
		return
	}
	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, ERROR_INTERNAL, http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) ListFiles(w http.ResponseWriter, r *http.Request) {

	list, err := a.Repository.FTP.ListFilesCallIDs()
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, ERROR_INTERNAL, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, ERROR_INTERNAL, http.StatusInternalServerError)
		return
	}
}

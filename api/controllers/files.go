package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//2020-03-20_07-49-01_4953080492_Inbound_79167013970_4953080492_3bae508b-2c43-4142-8a9b-899255b4da9f
//1378	106607	7001	89167013970	Outbound	Voice	2020-03-20T08:22:33Z	2020-03-20T08:23:04Z	1111	MYCUST	0044841E-D933-1E57-9082-06009C0AAA77-1092655@10.156.0.6		2020-03-20T08:23:02Z	1	3354965
func (a *Controllers) GetFile(w http.ResponseWriter, r *http.Request) {
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

	//check if

	if !user.ProjectNameAccessRights(*call.ProfileName) {
		http.Error(w, FILE_NOT_FOUND, http.StatusNotFound)
		return
	}

	if fileType != "txt" && fileType != "wav" {
		a.Logger.Errorf("%v", err)
		http.Error(w, FILE_INVALID_TYPE, http.StatusBadRequest)
		return
	}
	f.Type = fileType
	fName := call.ToFileName()
	res, err := a.Repository.FTP.GetFile(fName, fileType)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			a.Logger.Errorf("%v", err)
			http.Error(w, fmt.Sprintf(FILE_NOT_FOUND, fName), http.StatusNotFound)
			return
		}
		a.Logger.Errorf("%v", err)
		http.Error(w, ERROR_INTERNAL, http.StatusInternalServerError)
		return
	}
	f.Name = call.ToFileName()
	f.Data, err = ioutil.ReadAll(res)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, ERROR_INTERNAL, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, ERROR_INTERNAL, http.StatusInternalServerError)
		return
	}
}

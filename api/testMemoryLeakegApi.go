package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

var testVar int

type memoryLeakageResponse struct{
	Sum int `json:"sum"`
	Data any `json:"data"`
}
func (server *Server)TestMemoryLeakageHandler(w http.ResponseWriter, r *http.Request){
	logrus.Info("inside testMemoryLeakageHandler !!!!!!!!")
	for i:=1;i<=1000;i++{
		go Sum()
	}
	logrus.Info("final value after addition process ",testVar)
	var responseStruct memoryLeakageResponse
	responseStruct.Sum=testVar

	// call http request 
	resp, err :=http.Get("http://localhost:9999/v1/api/dynamic/struct/decoder")
	if err!=nil{
		logrus.Error("unable to get the response from the API")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err!=nil{
		logrus.Error("unable to read the response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  var result DummyUserDetails
  if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
     logrus.Error("unable to decode json")
		 http.Error(w, err.Error(), http.StatusBadRequest)
		 return
  }
	responseStruct.Data=result
	finalResponse,err :=json.Marshal(responseStruct)
	if err!=nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	w.Write(finalResponse)
}

func Sum(){
	for i:=0;i<=1000;i++{
		testVar++
	}
}


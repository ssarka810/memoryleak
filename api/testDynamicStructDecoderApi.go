package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type DummyUserDetails struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Details string `json:"details"`
}

func (server *Server)DynamicStructDecoderApiHandler(w http.ResponseWriter, r *http.Request){
logrus.Info("inside dynamicStructDecoderApihandler !!!!!!!!!!!")
	jsonData := []byte(`{
		"first_name": "Soumyadip",
		"last_name": "Sarkar",
		"age": 30,
		"city": "New York",
		"hasCar": true,
		"languages": ["Go", "JavaScript"]
	 }`)

	// Parse JSON into an empty interface
	var result interface{}
	if err :=json.Unmarshal(jsonData,&result);err!=nil{
		http.Error(w,"unable to parse input json", http.StatusInternalServerError)
		return
	}

	// Accessing dynamic JSON fields
 dataMap, ok := result.(map[string]interface{})
 if !ok {
  http.Error(w,"Invalid JSON structure",http.StatusBadGateway)
  return
 }

 // Accessing specific fields
 f_name, f_nameExists := dataMap["first_name"].(string)
 l_name, l_nameExists := dataMap["last_name"].(string)
 age, ageExists := dataMap["age"].(float64)

 var name string
 var response DummyUserDetails

 // Displaying parsed data
 if f_nameExists {
  logrus.Println("Fist Name:", f_name)
	name=fmt.Sprintf("%s",f_name)
 }
 if l_nameExists {
  logrus.Println("Last Name:", l_name)
	name=fmt.Sprintf("%s %s",name,l_name)
 }
response.Name=name
 if ageExists {
  logrus.Println("Age:", int(age))
	response.Age=int(age)
 }
 finalResponse, err :=json.Marshal(response)
 if err!=nil{
	logrus.Error("unable to encode the response, error: ",err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
 }
 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(http.StatusOK)
 w.Write(finalResponse)
	
}
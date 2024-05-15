package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type DynamicInputData struct{
	Nmae string `json:"name"`
	Data any `json:"data"`
}

func (server *Server)DynamicInputrApiHandler(w http.ResponseWriter, r *http.Request){
	logrus.Info("inside DynamicInputrApiHandler !!!!!!!!")

//-----------------------------------------------------------------------------------------
//----------------------- trying to store dynamic input using any keyword -----------------
//-----------------------------------------------------------------------------------------
var jsonData DynamicInputData
if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		logrus.Error("error decoding json response , error context ", err)
		return
}
logrus.Info("json data : ",jsonData)

// Accessing dynamic JSON fields
dataMap, ok := jsonData.Data.(map[string]interface{})
if !ok {
  http.Error(w,"Invalid JSON structure",http.StatusBadGateway)
  return
}
// Accessing specific fields
role, roleExists := dataMap["role"].(string)
location, locationExists := dataMap["location"].(string)
age, ageExists := dataMap["age"].(float64)

var finalString string ="hello everyone, client info : "

// Displaying parsed data
if roleExists {
  logrus.Println("Role:", role)
	finalString=fmt.Sprintf("%s  Role is %s",finalString,role)
}

if ageExists {
  logrus.Println("Age:", int(age))
	finalString=fmt.Sprintf("%s  Age is %f",finalString,age)
}
if locationExists {
  logrus.Println("Location :", location)
	finalString=fmt.Sprintf("%s  Location is %s",finalString,location)
}

w.Write([]byte(finalString))
	


//-----------------------------------------------------------------------------------------
//----------------------- trying to store dynamic input inside empty interface ------------
//-----------------------------------------------------------------------------------------

// 	// Parse JSON into an empty interface
// 	var result interface{}
// 	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
// 		logrus.Error("unable to decode input , error : ",err)
// 		http.Error(w,"unable to parse input json", http.StatusInternalServerError)
// 		return
// 	}

// 	// Accessing dynamic JSON fields
//  dataMap, ok := result.(map[string]interface{})
//  if !ok {
//   http.Error(w,"Invalid JSON structure",http.StatusBadGateway)
//   return
//  }
//  // Accessing specific fields
//  name, nameExists := dataMap["name"].(string)
//  age, ageExists := dataMap["age"].(float64)

//  var finalString string ="hello everyone, client info : "

//  // Displaying parsed data
//  if nameExists {
//   logrus.Println("Name:", name)
// 	finalString=fmt.Sprintf("%s  Name is %s",finalString,name)
//  }

//  if ageExists {
//   logrus.Println("Age:", int(age))
// 	finalString=fmt.Sprintf("%s  Age is %f",finalString,age)
//  }

// w.Write([]byte(finalString))
	
}
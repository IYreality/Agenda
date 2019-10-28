// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import(
	"io/ioutil"
	"encoding/json"
	"errors"
	"os"
)

type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
 
func ReadUserFromFile (filePath string) ([]User,error) {
	var users []User
	str,err := ioutil.ReadFile(filePath)
	if err!=nil {
		return users,err
	}
	jsonStr := string(str)
	
	json.Unmarshal([]byte(jsonStr),&users)
	return users,nil
}

func WriteUserToFile (filePath string, users []User) error{
	if data,err:=json.Marshal(users);err==nil{
		return ioutil.WriteFile(filePath,[]byte(data),os.ModeAppend)
	} else{
		return err
	}
}

//only check username repeat now
func userLegalCheck(userInfo []User,username string, password string,email string ,phone string) (bool,error){
	for _,user := range userInfo {
		if user.Username == username{
			return false,errors.New("Repeated Username")
		}
	}

	if len(password) == 0{
		return false,errors.New("Parameter password is missing")
	} else if len(email)==0 {
		return false,errors.New("Parameter email is missing")
	} else if len(phone)==0 {
		return false,errors.New("Parameter phone number is missing")
	}
	return true,nil
}

func checklogin() (bool,error){
	b,err := ioutil.ReadFile(cachePlace)
	if err!=nil {
		return false,err
	}
	str := string(b)

	if str == "logout"{
		return false,nil
	} else{
		return true,nil
	}
}

func getLoginUsername() (string,error){
	b,err := ioutil.ReadFile(cachePlace)
	if err!=nil {
		return "",err
	}
	return string(b),nil
}

func userLogin(username string) error{
	return ioutil.WriteFile(cachePlace,[]byte(username),os.ModeAppend)
}

func userLogout() error{
	return ioutil.WriteFile(cachePlace,[]byte("logout"),os.ModeAppend)
}


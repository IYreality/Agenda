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

import (
	"fmt"
	"github.com/spf13/cobra"
)

const userPlace = "user.txt"
const cachePlace = "cache.txt"

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "help a user to regist",
	Long: `
	1. register: 
	
	Usage:	user register -u [UserName] -p [Pass] -e [Email] -n [Phone]
		
		[Username] register's name
		[Pass] register's Password
		[Email] register's email
		[Phone] register's phone-number
	
	2. delete

	Usage:	user delete

	(attention: you can delete your account in the database of Agenda)

	3. lookup

	Usage:	user lookup

	(attention: you can query all user's name who has registed)

	4. login 

	Usage:	user login -u [UserName] -p [PassWord]

		[Username] register's name
		[Pass] register's Password

	5. logout

	Usage:	user logout`,
	/*Args:func(cmd *cobra.Command,args []string)error{
		if len(args)&lt;1{
			return errors.New("requires at least one arg")
		}
		return nil
	},*/
	Run: func(cmd *cobra.Command, args []string) {
		//reading file from store place
		userInfo,userReadingerr := ReadUserFromFile(userPlace)
		if userReadingerr!=nil {
			fmt.Println(userReadingerr)
			return
		}

		//get flags
		username, _ := cmd.Flags().GetString("username")
		password, _:= cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone,_ := cmd.Flags().GetString("phone")

		if len(args)>0 {
			switch (args[0]){
				case "register":{
					//legal check for username(unique),password,email,phone
					if pass,err := userLegalCheck(userInfo,username,password,email,phone); err!=nil{
						fmt.Println(err)
						return
					}else if !pass{
						fmt.Println("Register Failed")
						return
					}

					//if pass legal check, add it to userFile
					userInfo = append(userInfo,User{username,password,email,phone})
					//store the user file into userPlace
					WriteUserToFile(userPlace,userInfo)
					fmt.Println("User register success")
				}
				case "login":{
					//check from cache whether the status is login.
					if check,error := checklogin(); error!=nil{
						fmt.Println(error)
						return
					} else if check {
						fmt.Println("Already Login")
						return
					}
					//validate username and password
					if len(username) == 0 || len(password) == 0 {
						fmt.Println("Parameter username and password are missing")
						return
					}

					pass := false
					for _,user := range userInfo{
						if user.Username == username && user.Password == password{
							userLogin(user.Username)
							pass = true
							break
						}
					}
					//if no pass, report
					if !pass {
						fmt.Println("login failed")
						return
					}
					
					fmt.Println("Login success. Welcome!")
				}
				case "logout":{
					//if status is login, make the status logout

					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}

					userLogout()
					fmt.Println("Logout success")
					
				}
				case "lookup":{
					//check the status (login)
					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}
					//if pass validation, give all info from all users
					for _,user := range userInfo{
						fmt.Println(user.Username,user.Email,user.Phone)
					}
				}
				case "delete":{
					//check status login
					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}
					loginUsername,loginErr := getLoginUsername()
					if loginErr!=nil{
						fmt.Println(loginErr)
						return
					}
					//if pass, delete this user and logout
					for i,user := range userInfo{
						if loginUsername == user.Username{
							if i+1 < len(userInfo){
								userInfo = append(userInfo[:i],userInfo[i+1:]...)
							} else{
								userInfo = userInfo[:i]
							}
							
							break
						}
					}
					//update the userPlace
					WriteUserToFile(userPlace,userInfo)
					userLogout()
				}
				default:{
					fmt.Println("command not found...")
				}
			}
		}
		
	},
}


func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.Flags().StringP("username","u","","Help message for username")
	userCmd.Flags().StringP("password","p","","Help message for password")
	userCmd.Flags().StringP("email","e","","Help message for email")
	userCmd.Flags().StringP("phone","n","","Help message for phone number")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

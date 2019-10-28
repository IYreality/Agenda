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

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "help user to do something",
	Long: `

	you can use this app to create or remove users

	Usage:
		agenda [command]

	Available Commands:
		user : commands about user operation

	Use "agenda [command] --help" for more information about a command.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			if args[0] == "user" {
				if args[1] == "register" {
					fmt.Println(`
	Command : user register
	Function: Registe a user.
	Usage: user register --username/-u [UserName] --password/-p [Pass] --email/-e [Email] --telphone/-t [Phone]
	Args :
		[Username] 	register's name
		[Pass] 		register's Password
		[Email] 	register's email
		[Phone] 	register's phone-number`)
				} else if args[1] == "delete" {
					fmt.Println(`
	Command : user delete
	Function: Delete the current user.
	Usage: user delete
	`)
				} else if args[1] == "lookup" {
					fmt.Println(`
	Command : user lookup
	Function: Lookup all the information of the users. (Need to login first)
	Usage: user lookup
	`)
				} else if args[1] == "login" {
					fmt.Println(`
	Command : user login
	Function: Login
	(PS: If you don't login, you can only use user login and user register commands.)
	Usage : user login --username/-s [UserName] --password/-p [Pass]
	Args :
		[Username] 	register's name
		[Pass] 		register's Password`)
				} else if args[1] == "logout" {
					fmt.Println(`
	Command : user logout
	Function: Logout
	Usage: Logout the user
	`)
				}
			} 
			
		} else if len(args) == 1 {
			if args[0] == "user" {
				fmt.Println(`
	1.register: 
	
	Usage:	user register -u [UserName] -p [Pass] -e [Email] -t [Phone]
		
		[Username] register's name
		[Pass] register's Password
		[Email] register's email
		[Phone] register's phone-number
	
	2. delete

	Usage:	user delete

	(attention: you can delete your account in the database of Agenda)

	3. look up

	Usage:	user lookup

	(attention: you can query all user's name who has registed)

	4. login

	Usage:	user login -u [UserName] -p [PassWord]

		[Username] register's name
		[Pass] register's Password

	5. logout

	Usage:	user logout`)
			} 
		}

	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
	//rootCmd.SetHelpCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

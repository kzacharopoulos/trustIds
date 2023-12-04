package app

import "fmt"

var (
	Version  string
	Date     string
	Branch   string
	Commit   string
	GitUser  string
	GitEmail string
)

func ShowBuildInfo() {
	fmt.Println("Version : " + Version)
	fmt.Println("Date    : " + Date)
	fmt.Println("Branch  : " + Branch)
	fmt.Println("Commit  : " + Commit)
	fmt.Println("GitUser : " + GitUser)
	fmt.Println("GitEmail: " + GitEmail)
}

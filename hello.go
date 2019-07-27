package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
// get credentials
	r := bufio.NewReader(os.Stdin)
	fmt.Print("GitHub Organization: ")
	organization, _ := r.ReadString('\n')
	fmt.Print("GitHub Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("GitHub Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
// authenticate
	client := github.NewClient(tp.Client())
	ctx := context.Background()
	_, _, err := client.Users.Get(ctx, "")

// if two factor auth error
	if _, ok := err.(*github.TwoFactorAuthError); ok {
		fmt.Print("\nGitHub OTP: ")
		otp, _ := r.ReadString('\n')
		tp.OTP = strings.TrimSpace(otp)
		_, _, err = client.Users.Get(ctx, "")
	}
// if any other error
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}
// get allrepos for user	
	repos, _, err := client.Repositories.List(ctx, "", nil)
	repostring := github.Stringify(repos)
	
// split output into repo specific sections
	indexes := strings.Split(repostring,`github.Repository`)
	for _, index := range indexes {
		str := "Login:\"" + organization
		fstr := github.Stringify(index)
		
// see if repo belongs to organization		
			if strings.Contains(fstr, str) {
// see if license is set
				if !strings.Contains(github.Stringify(index),"License:github.License"){
					fmt.Print("This one needs a License")
					//license, _, err := client.(ctx, "", nil)
					
				}
			}
				
		
		
	}

}
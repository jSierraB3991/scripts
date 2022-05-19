package main

import (
    "bytes"
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
    "os"

    "github.com/brianvoe/gofakeit/v6"
)

func main() {
    usersCmd := flag.NewFlagSet("users", flag.ExitOnError)

    showName := usersCmd.Bool("name", false, "Show Names")
    showEmail := usersCmd.Bool("email", true, "Show Emails")
    showPassword := usersCmd.Bool("password", true, "Show Password")
    showPhone := usersCmd.Bool("phone", false, "Show Phone")
    showCompany := usersCmd.Bool("company", false, "Show Companies")
    showJob := usersCmd.Bool("job", false, "Show Jobs")

    sendToUrl := usersCmd.String("api", "", "Url to Send")
    //showCreditNumber := usersCmd.Bool("visa", false, "Show Credit Number")
    flag.Parse()

     if len(os.Args) < 2 {
        fmt.Println("expect 'users' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
        case "users":
            ShowRandomsUsers(usersCmd, sendToUrl,showName, showEmail, showPhone, showCompany, showJob, showPassword)
    }
}

type Users struct {
    Name        string
    Email       string
    Phone       string
    Company     string
    Job         string
    Password    string
}

func ShowRandomsUsers(usersCmd *flag.FlagSet, sendToUrl *string, showName *bool, showEmail *bool, 
                        showPhone *bool, showCompany *bool, showJob *bool, showPassword *bool) {
    usersCmd.Parse(os.Args[2:])
    var arraysUsers [10]Users
    for i:=0; i < 10; i++ {
        var user = Users{}
        if *showName == true {
            user.Name = gofakeit.Name()
        }
        if *showEmail == true {
            user.Email = gofakeit.Email()
        }
        if *showPhone == true {
            user.Phone = gofakeit.Phone()
        }
        if *showCompany == true {
            user.Company = gofakeit.Company()
        }
        if *showJob == true {
            user.Job = gofakeit.JobTitle()
        }
        if *showPassword == true {
            //gofakeit.Password(lower bool, upper bool, numeric bool, special bool, space bool, num int)
            user.Password = gofakeit.Password(true, true, true, false, false, 10)
        }

        arraysUsers[i] = user
    }

    for _, value := range arraysUsers {
        postBody, err := json.Marshal(value)
        if err != nil {
            fmt.Println("Fatal error array to json")
        }
        if *sendToUrl != "" {

            responseBody := bytes.NewBuffer(postBody)
            _, err := http.Post(*sendToUrl, "application/json", responseBody)
            if err != nil {
                fmt.Printf("fail send info to user with email %s\n", value.Email)
            } else {
                fmt.Printf("success send info to user with email %s\n", value.Email)
            }

        } else {
            fmt.Printf("Send user: %s to api: %s \n", postBody)
        }
    }
}

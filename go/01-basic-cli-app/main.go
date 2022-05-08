package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "os"
)

func main () {
    getCmd := flag.NewFlagSet("get", flag.ExitOnError)
    addCmd := flag.NewFlagSet("add", flag.ExitOnError)
    
    getAll := getCmd.Bool("all", false, "Get all videos")
    getById := getCmd.String("id", "", "Youtube videos ID")
    getStruct := addCmd.String("video", "", "Struct for adding videos")

    if len(os.Args) < 2 {
        fmt.Println("expect 'get' or 'add' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {
        case "get":
            HandleGet(getCmd, getAll, getById)
        case "add":
            HandleAdd(addCmd, getStruct)
        case "default":
            fmt.Printf("The flag %v ot recognize\n", os.Args[1])
    }
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {

    getCmd.Parse(os.Args[2:])

    if *all == false && *id == "" {
        fmt.Println("id is required or specify --all for all videos")
        getCmd.PrintDefaults()
        os.Exit(1)
    }

    if *all {
        showAllVideos()
    }
    if *id != "" {
        showOneVideo(*id)
    }
}

func showAllVideos() {
    videos := getVideos()
//    fmt.Printf("ID \t Title \t URL \t ImageUrl \t Description \n")
    for _, video := range videos {
        fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageUrl, video.Description)
    }
}

func showOneVideo(id_video string) {
    videos := getVideos()
    for _, video := range videos {
        if id_video == video.Id {
            fmt.Printf("ID \t Title \t URL \t ImageUrl \t Description \n")
            fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.ImageUrl, video.Description)
        }
    }
}

func ValidateVideo(addCmd *flag.FlagSet, video_data *string) (video) {
    addCmd.Parse(os.Args[2:])

    if *video_data == "" {
        fmt.Println("flag video is neccesary for pass video add")
        addCmd.PrintDefaults()
        os.Exit(1)
    }
    var videoAdd video
    err := json.Unmarshal([]byte(*video_data), &videoAdd)
    if err != nil {
        fmt.Printf("Error %v is not incorrect format\n", *video_data)
        fmt.Println("The structure for flag video is")
        fmt.Println("-video '{ \"Id\": \"value\", \"Title\": \"value\", \"Url\": \"value\", \"ImageUrl\": \"value\", \"Description\": \"value\" }'")
        os.Exit(1)
    }
    return videoAdd
}
func HandleAdd(addCmd *flag.FlagSet, video *string) {
    videoAdd := ValidateVideo(addCmd, video)

    videos := getVideos()
    videos = append(videos, videoAdd)
    saveVideos(videos)
}


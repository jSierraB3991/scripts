package main

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func checkGitInstall() bool {
	_, err := exec.LookPath("git")
	if err != nil {
		fmt.Println("You don have git installed")
		return false
	}
	return true
}

type GitHubResponse struct {
	DefaultBranch string `json:"default_branch"`
}

func CheckFlagBranch(username string, reponame string, branch string) {
	response, err := http.Get("https://api.github.com/repos/" + username + "/" + reponame + "/branches/" + branch)
	if err != nil {
		fmt.Println("The branch " + branch + "no exists in this repository")
		checkError(err)
	}
	if response.StatusCode == http.StatusNotFound {
		fmt.Println("The branch " + branch + " no exists in this repository")
		os.Exit(1)
	}
}

func GetMainBranchName(username string, reponame string) (string, error) {
	response, err := http.Get("https://api.github.com/repos/" + username + "/" + reponame)
	if err != nil {
		return "", err
	}
	data, _ := ioutil.ReadAll(response.Body)
	// bodyStr := string(data)
	var obj GitHubResponse
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return "", err
	}
	return obj.DefaultBranch, nil
}

func checkError(err error) {
	if err != nil {
		if _, err := os.Stat(".templify-temp"); !errors.Is(err, os.ErrNotExist) {
			_ = os.RemoveAll(".templify-temp")
		}
		panic(err)
	}
}

func DownloadFile(filepath string, reponame string, url string) error {
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	return err
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()
	for _, f := range r.File {
		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}
		filenames = append(filenames, fpath)
		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		_, err = io.Copy(outFile, rc)
		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()
		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func DownloadZip(ownerRepo string, repoName string, branchName string, tempDir string, dstPath string) {
	fileUrl := "https://github.com/" + ownerRepo + "/" + repoName + "/archive/refs/heads/" + branchName + ".zip"
	zipStr := []string{tempDir, "zip.zip"}
	zipFilePath := strings.Join(zipStr, "/")
	DownloadFile(zipFilePath, repoName, fileUrl)

	unzipStr := []string{tempDir, "unzipped"}
	unzipPath := strings.Join(unzipStr, "/")
	Unzip(zipFilePath, unzipPath)

	files, err := ioutil.ReadDir(unzipPath)
	checkError(err)
	repoDir := files[0]

	// Move the folder
	repoDirStr := []string{unzipPath, repoDir.Name()}
	repoDirPath := strings.Join(repoDirStr, "/")
	err = os.Rename(repoDirPath, "./"+dstPath)
	checkError(err)
	defer os.RemoveAll(tempDir)
}

func CloneRepo(urlRepo string) {
	gitBin, _ := exec.LookPath("git")
	cmd := &exec.Cmd{
		Path:   gitBin,
		Args:   []string{gitBin, "clone", urlRepo},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	err := cmd.Run()
	checkError(err)
}

func main() {
	defaultBranch := flag.String("branch", "", "The default branc of repository")
	outDirectory := flag.String("out", "", "folder name of repository, by default is name of repository")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("I need url of repository to clone")
	}

	urlRepo := flag.Args()[0]
	repoName := *outDirectory
	branchName := *defaultBranch
	if *outDirectory == "" {
		repoName = strings.Split(urlRepo, "/")[4]
	}
	ownerRepo := strings.Split(urlRepo, "/")[3]
	if *defaultBranch == "" {
		masterBranch, err := GetMainBranchName(ownerRepo, repoName)
		checkError(err)
		branchName = masterBranch
	} else {
		CheckFlagBranch(ownerRepo, repoName, branchName)
	}

	isGitInstall := checkGitInstall()
	if isGitInstall {
		// Create temp folder
		tempDir := ".templify-temp"
		err := os.Mkdir(".templify-temp", 0755)
		checkError(err)
		fmt.Println("Download file")
		DownloadZip(ownerRepo, repoName, branchName, tempDir, repoName)
	} else {
		fmt.Println("Clone file")
		CloneRepo(urlRepo)
	}

	fmt.Printf("\nbranch:%v\nrepo:%v\noutputName: %v\n", branchName, urlRepo, repoName)
}

package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ozonebg/juicetomator/internal/models"
	"github.com/sirupsen/logrus"
)

const (
	// AppName is the name of the python version to use.
	AppName = "python3"

	// ScriptName is the name of the script that will be executed.
	ScriptName = "mp3downloader.py"
)

// RespondWithMessage responds to the server with a string message.
func RespondWithMessage(c *gin.Context, status int, errorMessage string) {
	if c == nil {
		return
	}

	c.JSON(status, gin.H{"error": errorMessage})
}

// RespondWithError responds to the server with a error message.
func RespondWithError(c *gin.Context, status int, errorMessage error) {
	RespondWithMessage(c, status, errorMessage.Error())
}

// IOReadDir finds a file by filename
func IOReadDir(root string, fileName string) (string, error) {
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return "", err
	}
	for _, file := range fileInfo {
		if file.Name() == fileName {
			return file.Name(), nil
		}
	}
	return "", errors.New("could not find file")
}

// ExecuteDownloadScript executes the script that will download all provided files.
func ExecuteDownloadScript(in models.DownloadInfo) (string, error) {
	songs := strings.Join(in.Songs, ",")
	songs = "\"" + songs + "\""

	scriptPath := GetCurrentDirectory() + "/scripts/" + ScriptName

	logrus.Info(scriptPath)

	cmd := exec.Command(AppName, scriptPath, songs, in.Username)
	// stdout, err := cmd.Output()

	out, err := cmd.CombinedOutput()

	logrus.WithError(err).Info(string(out))

	// return string(stdout), err
	return "", nil
}

func GetCurrentDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		logrus.WithError(err).Info("failed to get directory")
		return ""
	}

	return path
}

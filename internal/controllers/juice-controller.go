package controllers

import (
	"net/http"

	"github.com/ozonebg/juicetomator/internal/constants"

	"github.com/gin-gonic/gin"
	"github.com/ozonebg/juicetomator/internal/models"
	"github.com/ozonebg/juicetomator/internal/utils"
	log "github.com/sirupsen/logrus"
)

var (
	logger = log.WithField("component", "juice-controller")
)

// JuiceController main responsibility is to handle all download requests.
type JuiceController struct {
}

// NewJuiceController creates a new juice controller.
func NewJuiceController() *JuiceController {
	return &JuiceController{}
}

// HandleMultiJuices handles multiple juices downloads.
func (controller *JuiceController) HandleMultiJuices(c *gin.Context) {
	in := models.DownloadInfo{}

	if c.BindJSON(&in) != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrFailedToBindJSON)
		return
	}

	// Execute script
	output, err := utils.ExecuteDownloadScript(in)
	logger.WithField("script_output", output).Info("script output")
	if err != nil {
		logger.WithError(err).Info("failure during script execution")
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrFailedToExecuteScript)
		return
	}

	file, err := utils.IOReadDir("./", in.Username+".zip")
	if err != nil {
		log.WithError(err).Info("failed to find the zip file")
		utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrFailedToFindZipFile)
		return
	}

	// File was found so return it to the user
	c.File(file)
}

// HandleSingleJuice handles a single juice download.
func (controller *JuiceController) HandleSingleJuice(c *gin.Context) {
	utils.RespondWithError(c, http.StatusInternalServerError, constants.ErrNotImplemented)
}

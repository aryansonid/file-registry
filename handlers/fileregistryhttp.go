package handlers

import (
	"encoding/json"
	"file-registry/service"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

type Handler struct {
	service service.Service
}

// NewHandler creates a new Handler with the given service, configuration, and logger.
func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) UploadFileHandler(c *gin.Context) {
	// Check if the method is POST
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	// Step 1: Parse multipart form
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing form: %s", err.Error())})
		return
	}

	// Get file and filePath
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error retrieving file: %s", err.Error())})
		return
	}
	defer file.Close()

	filePath := c.Request.FormValue("filePath")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File path is required"})
		return
	}

	// Step 2: Send the file to IPFS service
	cid, err := service.UploadFileToIPFS(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error uploading file to IPFS: %s", err.Error())})
		return
	}

	// Step 3: Interact with smart contract to save the filePath and CID
	tx, err := h.service.SaveDetailsInContract(filePath, cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error saving details in contract: %s", err.Error())})
		return
	}

	// Return success response with CID
	response := UploadFileResponse{
		CID:                 cid,
		TransactionResponse: tx,
	}

	json_response, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : fmt.Sprintf("Failed to marshal response: %s", err.Error())})
		return
	}

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusCreated, json_response)
}

func (h *Handler) GetFileHandler(c *gin.Context) {
	// Check if the method is GET
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	// Extract filePath from query params
	filePath := c.Request.URL.Query().Get("filePath")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File path is required"})
		return
	}

	// Interact with smart contract to get the CID
	cid, err := h.service.GetCIDFromContract(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error getting CID from contract: %s", err.Error())})
		return
	}

	// Return the CID in the response
	response := GetFileResponse{
		CID: cid,
	}

	json_response, err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to marshal response: %s", err.Error())})
		return
	}

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, json_response)
}

package service

import (
	"fmt"
	"io"

	shell "github.com/ipfs/go-ipfs-api"
	logger "github.com/sirupsen/logrus"
)

func UploadFileToIPFS(file io.Reader) (cid string, err error) {
	// Initialize IPFS shell (point to your local IPFS node or remote node)
	sh := shell.NewShell("localhost:5001")

	// Add the file to IPFS
	cid, err = sh.Add(file)
	if err != nil {
		return "", fmt.Errorf("failed to upload to IPFS: %v", err)
	}

	logger.Infof("Uploaded file to IPFS with CID: %s", cid)

	return
}

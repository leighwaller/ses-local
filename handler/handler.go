package handler

import (
	"net/http"
	"log"
	"encoding/xml"
	"github.com/satori/go.uuid"
	"encoding/base64"
	"fmt"
	"ses-local/storage"
)

const maxDestinations = 50

var dao storage.EmailDao = storage.FileSystemDao{}

func Index(writer http.ResponseWriter, request *http.Request) {
	action := request.FormValue("Action")
	if action == "" {
		writeError(writer, http.StatusBadRequest, MissingAction, "'Action' request parameter is required")
		return
	}

	log.Printf("Handling SES '%s' Request", action)
	switch action {
	case "SendRawEmail":
		handleSendRawEmail(writer, request)
	default:
		message := "'Action' request parameter is invalid or undefined"
		writeError(writer, http.StatusBadRequest, InvalidAction, message)
	}
}

func handleSendRawEmail(writer http.ResponseWriter, request *http.Request) {
	data := request.FormValue("RawMessage.Data")
	if data == "" {
		message := "'RawMessage.Data' request parameter is required for this operation"
		writeError(writer, http.StatusBadRequest, MissingParameter, message)
		return
	}
	printAddressParams(request)

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		writeError(writer, http.StatusBadRequest, InvalidParameterValue, "Unable to decode raw message data")
		return
	}

	log.Printf("Raw data: %s", decoded)
	messageId := uuid.NewV4().String()
	result, err := dao.Save(messageId, decoded)
	if err != nil {
		writeError(writer, http.StatusInternalServerError, InternalFailure, "Failed to save email")
		return
	}

	response, err := xml.Marshal(rawEmailResponse(result))
	if err != nil {
		writeError(writer, http.StatusInternalServerError, InternalFailure, "Unable to marshal response body")
		return
	}

	log.Printf("Response: %s", response)
	writer.Write(response)
}

// actual source/dest taken from the raw message data
func printAddressParams(request *http.Request) {
	source := request.FormValue("Source")
	log.Printf("Source: %s", source)

	for i := 1; i < maxDestinations; i++ {
		dest := request.FormValue(fmt.Sprintf("Destinations.member.%d", i))
		if dest == "" {
			break
		}
		log.Printf("Destination [%d]: %s", i, dest)
	}
}

func rawEmailResponse(key string) *SendRawEmailResponse {
	// return the location of the file as the message id
	requestId := uuid.NewV4().String()
	return &SendRawEmailResponse{MessageId: key, RequestId: requestId}
}

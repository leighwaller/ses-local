package handler

import (
	"encoding/xml"
	"net/http"
)

type SendRawEmailResponse struct {
	XMLName   xml.Name `xml:"SendRawEmailResponse"`
	MessageId string `xml:"SendRawEmailResult>MessageId"`
	RequestId string `xml:"ResponseMetadata>RequestId"`
}

type Error struct {
	Code string `xml:"Code"`
	Message interface{} `xml:"Message"`
}

func writeError(writer http.ResponseWriter, code int, reason string, message interface{}) {
	response, err := xml.Marshal(Error{Code: reason, Message: message})
	if err != nil {
		panic(err)
	}
	writer.WriteHeader(code)
	writer.Write(response)
}

const (
	MessageRejected = "MessageRejected"
	MissingAction = "MissingAction"
	InvalidAction = "InvalidAction"
	MissingParameter = "MissingParameter"
	InvalidParameterValue = "InvalidParameterValue"
	InternalFailure = "InternalFailure"
)
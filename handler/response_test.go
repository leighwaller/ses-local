package handler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"net/http/httptest"
)

func TestRawEmailResponse(t *testing.T) {
	msgId := "123"
	reqId := "ABC"
	response := SendRawEmailResponse{MessageId: msgId, RequestId: reqId}

	expected := "<SendRawEmailResponse><SendRawEmailResult><MessageId>" + msgId + "</MessageId></SendRawEmailResult>" +
		"<ResponseMetadata><RequestId>" + reqId + "</RequestId></ResponseMetadata></SendRawEmailResponse>"

	actual, _ := xml.Marshal(response)
	assert.Equal(t, expected, string(actual))
}

func TestError(t *testing.T) {
	code := MissingAction
	message := "Something is required"
	response := Error{Code: code, Message: message}

	expected := "<Error><Code>" + code + "</Code><Message>" + message + "</Message></Error>"

	actual, _ := xml.Marshal(response)
	assert.Equal(t, expected, string(actual))
}

func TestWriteError(t *testing.T) {
	writer := httptest.NewRecorder()

	status := 500
	writeError(writer, status, MessageRejected, "Something happened")

	assert.Equal(t, status, writer.Code)
	assert.NotNil(t, writer.Body)
}

func TestErrorTypes(t *testing.T) {
	assert.Equal(t, "MessageRejected", MessageRejected)
	assert.Equal(t, "MissingAction", MissingAction)
	assert.Equal(t, "InvalidAction", InvalidAction)
	assert.Equal(t, "MissingParameter", MissingParameter)
	assert.Equal(t, "InvalidParameterValue", InvalidParameterValue)
	assert.Equal(t, "InternalFailure", InternalFailure)
}

package utils

import (
	"encoding/json"
	"fmt"
	"local-cache/models"
	"net/http"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

const (
	ServerErrorMsg = "Internal Server Error occurred when processing request."
)

type clientError struct {
	ID            string        `json:"id"`
	MessageToUser string        `json:"messageToUser"`
	DeveloperInfo string        `json:"developerInfo"`
	Err           string        `json:"error"`
	StatusCode    int           `json:"statusCode"`
	IsClientError bool          `json:"isClientError"`
	Request       *http.Request `json:"-"`
} // @name clientError

func (c *clientError) Error() string {
	return c.Err
}

func (c *clientError) LogFields() logrus.Fields {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	userID := 0
	var userEmail, userPhone string
	uc, ok := c.Request.Context().Value(models.UserContextKey).(*models.UserContext)
	if ok {
		if uc != nil {
			userID = uc.ID
			userEmail = uc.Email.String
			userPhone = uc.Phone.String
		}
	}

	return logrus.Fields{
		"RequestId":           c.ID,
		"userId":              userID,
		"userEmail":           userEmail,
		"userPhone":           userPhone,
		"MessageToUser":       c.MessageToUser,
		"DeveloperInfo":       c.DeveloperInfo,
		"RequestError":        true,
		"http.request.method": c.Request.Method,
		"http.request.uri":    fmt.Sprintf("%s://%s%s", scheme, c.Request.Host, c.Request.RequestURI),
	}
}

func (c *clientError) LogMessage() string {
	return fmt.Sprintf("[connect-upClientErr]: id(%s) %s : %+v", c.ID, c.MessageToUser+" "+c.DeveloperInfo, c.Err)
}

func newClientError(err error, statusCode int, req *http.Request, messageToUser string, additionalInfoForDevs ...string) *clientError {
	additionalInfoJoined := strings.Join(additionalInfoForDevs, "\n")
	if additionalInfoJoined == "" {
		additionalInfoJoined = messageToUser
	}

	var errString string
	if err != nil {
		errString = err.Error()
	}

	return &clientError{
		ID:            middleware.GetReqID(req.Context()),
		MessageToUser: messageToUser,
		DeveloperInfo: additionalInfoJoined,
		Err:           errString,
		StatusCode:    statusCode,
		IsClientError: true,
		Request:       req,
	}
}

func RespondClientErr(resp http.ResponseWriter, req *http.Request, err error, statusCode int, messageToUser string, additionalInfoForDevs ...string) {
	logrus.Errorf("messageToUser: %v", messageToUser)
	resp.WriteHeader(statusCode)
	clientError := newClientError(err, statusCode, req, messageToUser, additionalInfoForDevs...)
	if statusCode >= 400 && statusCode < 500 {
		logrus.WithContext(req.Context()).WithFields(clientError.LogFields()).Warn(clientError.LogMessage())
	} else {
		logrus.WithContext(req.Context()).WithFields(clientError.LogFields()).Error(clientError.LogMessage())
	}
	if err := json.NewEncoder(resp).Encode(clientError); err != nil {
		logrus.Error(err)
	}
}

func RespondGenericServerErr(resp http.ResponseWriter, req *http.Request, err error, additionalInfoForDevs ...string) {
	resp.WriteHeader(http.StatusInternalServerError)
	additionalInfoJoined := strings.Join(additionalInfoForDevs, "\n")

	clintErr := &clientError{
		ID:            middleware.GetReqID(req.Context()),
		MessageToUser: ServerErrorMsg,
		DeveloperInfo: additionalInfoJoined,
		Err:           err.Error(),
		StatusCode:    http.StatusInternalServerError,
		IsClientError: false,
		Request:       req,
	}
	logrus.WithContext(req.Context()).WithFields(clintErr.LogFields()).Error(clintErr.LogMessage())
	if err := json.NewEncoder(resp).Encode(clintErr); err != nil {
		logrus.Error(err)
	}
}

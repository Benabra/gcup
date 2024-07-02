package services

import (
	"fmt"

	"google.golang.org/api/googleapi"
)

type ErrorDetail struct {
	Project  string
	Service  string
	Resource string
	Reason   string
}

func HandleError(projectID, serviceName, resource string, err error, errorDetails []ErrorDetail) []ErrorDetail {
	var reason string
	if apiErr, ok := err.(*googleapi.Error); ok {
		reason = fmt.Sprintf("API error: %v", apiErr.Message)
		if apiErr.Code == 403 {
			reason = "Permission denied. Ensure the service account has the required permissions."
		} else if apiErr.Code == 404 {
			reason = "Resource not found."
		}
	} else {
		reason = err.Error()
	}
	errorDetails = append(errorDetails, ErrorDetail{
		Project:  projectID,
		Service:  serviceName,
		Resource: resource,
		Reason:   reason,
	})
	return errorDetails
}

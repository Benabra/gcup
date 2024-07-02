package services

import (
	"context"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/api/appengine/v1"
	"google.golang.org/api/option"
)

func ListAppEngineServices(ctx context.Context, projectID string, table *tablewriter.Table) error {
	appEngineService, err := appengine.NewService(ctx, option.WithScopes(appengine.CloudPlatformScope))
	if err != nil {
		return fmt.Errorf("failed to create App Engine client: %v", err)
	}

	appServices, err := appEngineService.Apps.Services.List(projectID).Do()
	if err != nil {
		return err
	}
	for _, service := range appServices.Services {
		table.Append([]string{"App Engine", service.Id, "RUNNING", "", ""})
	}
	return nil
}

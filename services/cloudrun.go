package services

import (
	"context"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/api/option"
	run "google.golang.org/api/run/v1"
)

func ListCloudRunServices(ctx context.Context, projectID string, table *tablewriter.Table) error {
	runService, err := run.NewService(ctx, option.WithScopes(run.CloudPlatformScope))
	if err != nil {
		return fmt.Errorf("failed to create Cloud Run client: %v", err)
	}

	runServices, err := runService.Projects.Locations.Services.List("projects/" + projectID + "/locations/-").Do()
	if err != nil {
		return err
	}
	for _, service := range runServices.Items {
		creationTime := service.Metadata.CreationTimestamp
		lastModifiedTime := service.Metadata.Annotations["client.knative.dev/lastModifier"]
		table.Append([]string{"Cloud Run", service.Metadata.Name, "RUNNING", creationTime, lastModifiedTime})
	}
	return nil
}

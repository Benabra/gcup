package services

import (
	"context"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func ListComputeEngineInstances(ctx context.Context, projectID string, table *tablewriter.Table) error {
	computeService, err := compute.NewService(ctx, option.WithScopes(compute.CloudPlatformScope))
	if err != nil {
		return fmt.Errorf("failed to create Compute Engine client: %v", err)
	}

	instancesReq := computeService.Instances.AggregatedList(projectID)
	return instancesReq.Pages(ctx, func(page *compute.InstanceAggregatedList) error {
		for _, instances := range page.Items {
			for _, instance := range instances.Instances {
				if instance.Status == "RUNNING" {
					creationTime := instance.CreationTimestamp
					lastModifiedTime := instance.LastStartTimestamp // This might be the closest to a last modified time
					table.Append([]string{"Compute Engine", instance.Name, "RUNNING", creationTime, lastModifiedTime})
				}
			}
		}
		return nil
	})
}

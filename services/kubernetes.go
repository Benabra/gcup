package services

import (
	"context"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/option"
)

func ListKubernetesEngineClusters(ctx context.Context, projectID string, table *tablewriter.Table) error {
	containerService, err := container.NewService(ctx, option.WithScopes(container.CloudPlatformScope))
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes Engine client: %v", err)
	}

	clustersReq := containerService.Projects.Locations.Clusters.List("projects/" + projectID + "/locations/-")
	clustersResp, err := clustersReq.Do()
	if err != nil {
		return err
	}
	for _, cluster := range clustersResp.Clusters {
		creationTime := cluster.CreateTime
		lastModifiedTime := "" // Kubernetes API does not provide an explicit update time
		table.Append([]string{"Kubernetes Engine", cluster.Name, "RUNNING", creationTime, lastModifiedTime})
	}
	return nil
}

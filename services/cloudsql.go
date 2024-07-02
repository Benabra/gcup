package services

import (
	"context"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/api/option"
	"google.golang.org/api/sqladmin/v1"
)

func ListCloudSQLInstances(ctx context.Context, projectID string, table *tablewriter.Table) error {
	sqlService, err := sqladmin.NewService(ctx, option.WithScopes(sqladmin.SqlserviceAdminScope))
	if err != nil {
		return fmt.Errorf("failed to create Cloud SQL client: %v", err)
	}

	instances, err := sqlService.Instances.List(projectID).Do()
	if err != nil {
		return err
	}
	for _, instance := range instances.Items {
		creationTime := instance.CreateTime // This is in RFC3339 format
		lastModifiedTime := ""              // Cloud SQL does not provide an explicit update time
		table.Append([]string{"Cloud SQL", instance.Name, instance.State, creationTime, lastModifiedTime})
	}
	return nil
}

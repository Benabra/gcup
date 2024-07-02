package services

import (
	"context"
	"fmt"
	"time"

	"github.com/olekukonko/tablewriter"
	"google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"
)

func ListBigQueryDatasets(ctx context.Context, projectID string, table *tablewriter.Table) error {
	bqService, err := bigquery.NewService(ctx, option.WithScopes(bigquery.BigqueryScope))
	if err != nil {
		return fmt.Errorf("failed to create BigQuery client: %v", err)
	}

	datasets, err := bqService.Datasets.List(projectID).Do()
	if err != nil {
		return err
	}
	for _, dataset := range datasets.Datasets {
		datasetDetail, err := bqService.Datasets.Get(projectID, dataset.DatasetReference.DatasetId).Do()
		if err != nil {
			return err
		}
		creationTime := time.Unix(0, datasetDetail.CreationTime*int64(time.Millisecond)).Format(time.RFC3339)
		lastModifiedTime := time.Unix(0, datasetDetail.LastModifiedTime*int64(time.Millisecond)).Format(time.RFC3339)
		table.Append([]string{"BigQuery", dataset.DatasetReference.DatasetId, "AVAILABLE", creationTime, lastModifiedTime})
	}
	return nil
}

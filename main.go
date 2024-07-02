package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"gcup/services"

	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
)

var errorDetails []services.ErrorDetail

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func main() {
	// Define and parse command-line flags
	projects := flag.String("projects", "", "Comma-separated list of Google Cloud project IDs")
	flag.Parse()

	if *projects == "" {
		fmt.Println("Project IDs must be provided using the -projects flag")
		os.Exit(1)
	}

	projectList := strings.Split(*projects, ",")

	ctx := context.Background()

	// Initialize the progress bar
	totalSteps := len(projectList) * 6 // Number of services per project
	bar := progressbar.Default(int64(totalSteps))

	for _, projectID := range projectList {
		projectID = strings.TrimSpace(projectID)
		fmt.Printf("\n%sChecking services for project: %s%s\n", bold, projectID, reset)
		fmt.Println(strings.Repeat("=", 50))

		// Create a table for services
		serviceTable := tablewriter.NewWriter(os.Stdout)
		serviceTable.SetHeader([]string{"Service", "Resource", "Status", "Created At", "Updated At"})

		// List Cloud Run services
		if err := services.ListCloudRunServices(ctx, projectID, serviceTable); err != nil {
			errorDetails = services.HandleError(projectID, "Cloud Run", "Cloud Run Service", err, errorDetails)
		}
		bar.Add(1) // Increment the progress bar

		// List App Engine services
		if err := services.ListAppEngineServices(ctx, projectID, serviceTable); err != nil {
			errorDetails = services.HandleError(projectID, "App Engine", "App Engine Service", err, errorDetails)
		}
		bar.Add(1) // Increment the progress bar

		// List Compute Engine instances
		if err := services.ListComputeEngineInstances(ctx, projectID, serviceTable); err != nil {
			errorDetails = services.HandleError(projectID, "Compute Engine", "Compute Instance", err, errorDetails)
		}
		bar.Add(1) // Increment the progress bar

		// List Kubernetes Engine services
		if err := services.ListKubernetesEngineClusters(ctx, projectID, serviceTable); err != nil {
			errorDetails = services.HandleError(projectID, "Kubernetes Engine", "Kubernetes Cluster", err, errorDetails)
		}
		bar.Add(1) // Increment the progress bar

		// List Cloud SQL instances
		if err := services.ListCloudSQLInstances(ctx, projectID, serviceTable); err != nil {
			errorDetails = services.HandleError(projectID, "Cloud SQL", "Cloud SQL Instance", err, errorDetails)
		}
		bar.Add(1) // Increment the progress bar

		// List BigQuery datasets
		if err := services.ListBigQueryDatasets(ctx, projectID, serviceTable); err != nil {
			errorDetails = services.HandleError(projectID, "BigQuery", "BigQuery Dataset", err, errorDetails)
		}
		bar.Add(1) // Increment the progress bar

		// Check if any services were found before rendering the table
		if serviceTable.NumLines() > 0 {
			fmt.Println("\nRunning Services:")
			serviceTable.Render()
		} else {
			fmt.Printf("\n%sNo running services found for project: %s%s\n", yellow, projectID, reset)
		}
	}

	// Create a table for error details
	errorTable := tablewriter.NewWriter(os.Stdout)
	errorTable.SetHeader([]string{"Project", "Service", "Resource", "Reason"})

	for _, detail := range errorDetails {
		errorTable.Append([]string{detail.Project, detail.Service, detail.Resource, detail.Reason})
	}

	// Render the error details table
	if len(errorDetails) > 0 {
		fmt.Println("\nErrors:")
		errorTable.Render()
	}
}

# GCP Services Checker

This Go application lists running services in one or more Google Cloud projects, including Cloud Run, App Engine, Compute Engine, and Kubernetes Engine. It also provides a progress bar to indicate the progress of the operations.

## Project Structure

```
gcup/
|-- main.go
|-- services/
|   |-- appengine.go
|   |-- cloudrun.go
|   |-- compute.go
|   |-- kubernetes.go
|   |-- utils.go
```

## Prerequisites

- Go 1.16 or later
- Google Cloud SDK (for authentication)
- Internet access to download dependencies

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Benabra/gcup.git
    cd gcup
    ```

2. Initialize your Go module (if not already initialized):

    ```sh
    go mod init gcup
    ```

3. Install necessary dependencies:

    ```sh
    go get github.com/schollz/progressbar/v3 github.com/olekukonko/tablewriter google.golang.org/api/option google.golang.org/api/run/v1 google.golang.org/api/appengine/v1 google.golang.org/api/compute/v1 google.golang.org/api/container/v1
    ```

## Usage

1. Authenticate with Google Cloud:

    Ensure you are authenticated with Google Cloud:

    ```sh
    gcloud auth application-default login
    ```

2. Run the program with your Google Cloud project IDs:

    ```sh
    go run main.go -projects=PROJECT_ID_1,PROJECT_ID_2,PROJECT_ID_3
    ```

    Replace `PROJECT_ID_1`, `PROJECT_ID_2`, `PROJECT_ID_3`, etc., with your actual Google Cloud project IDs.

## Project Details

- **main.go**: The entry point of the application.
- **services/**: Contains Go files for each GCP service and utility functions.
  - **appengine.go**: Lists App Engine services.
  - **cloudrun.go**: Lists Cloud Run services.
  - **compute.go**: Lists Compute Engine instances.
  - **kubernetes.go**: Lists Kubernetes Engine clusters.
  - **utils.go**: Contains utility functions such as error handling.

## Example Output

Running the application will display a progress bar and then output tables of running services and any errors encountered:

```
[###-------------------------------] 1/4
Checking services for project: project-id-1
==================================================

Running Services:
+------------------+------------+---------+
|      SERVICE     |  RESOURCE  |  STATUS |
+------------------+------------+---------+
|    Cloud Run     | service1   | RUNNING |
|    App Engine    | service2   | RUNNING |
| Compute Engine   | instance1  | RUNNING |
| Kubernetes Engine| cluster1   | RUNNING |
+------------------+------------+---------+

==================================================
Checking services for project: project-id-2
==================================================

No running services found for project: project-id-2

==================================================
Checking services for project: project-id-3
==================================================

Running Services:
+------------------+------------+---------+
|      SERVICE     |  RESOURCE  |  STATUS |
+------------------+------------+---------+
|    Cloud Run     | service3   | RUNNING |
|    App Engine    | service4   | RUNNING |
| Compute Engine   | instance2  | RUNNING |
| Kubernetes Engine| cluster2   | RUNNING |
+------------------+------------+---------+

Errors:
+------------+------------------+-------------------+----------------------------------------+
|  Project   |      Service     |      Resource     |                Reason                 |
+------------+------------------+-------------------+----------------------------------------+
| project-id |    Cloud Run     | Cloud Run Service | Permission denied for Cloud Run.      |
+------------+------------------+-------------------+----------------------------------------+
```

## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/fooBar`).
3. Commit your changes (`git commit -am 'Add some fooBar'`).
4. Push to the branch (`git push origin feature/fooBar`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
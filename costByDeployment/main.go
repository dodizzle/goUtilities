package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"

	"github.com/rodaine/table"
)

// PrometheusResponse represents the structure of the Prometheus API response
type PrometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"` // Updated to handle "value" field
		} `json:"result"`
	} `json:"data"`
}

// queryPrometheus queries the Prometheus API with basic authentication and x-scope-orgid header
func queryPrometheus(query string, orgID string) ([]float64, error) {
	// Get Mimir credentials from environment variables
	username := os.Getenv("MIMIR_FE_USERNAME")
	password := os.Getenv("MIMIR_FE_PASSWORD")
	if username == "" || password == "" {
		return nil, fmt.Errorf("MIMIR_FE_USERNAME or MIMIR_FE_PASSWORD environment variable is not set")
	}

	// Set the Prometheus URL
	prometheusURL := "https://mimir-query-frontend.infra.alto.com/prometheus/api/v1/query"

	// Build the query URL
	u, err := url.Parse(prometheusURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Prometheus URL: %v", err)
	}
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()

	// Create the HTTP request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Add basic authentication header
	req.SetBasicAuth(username, password)

	// Add x-scope-orgid header
	req.Header.Set("x-scope-orgid", orgID)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query Prometheus: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response
	var promResp PrometheusResponse
	if err := json.Unmarshal(body, &promResp); err != nil {
		return nil, fmt.Errorf("failed to parse Prometheus response: %v", err)
	}

	// Extract the values
	var values []float64
	for _, result := range promResp.Data.Result {
		if len(result.Value) > 1 {
			// Parse the metric value (result.Value[1])
			if valueStr, ok := result.Value[1].(string); ok {
				var floatValue float64
				_, err := fmt.Sscanf(valueStr, "%f", &floatValue)
				if err != nil {
					continue
				}
				values = append(values, floatValue)
			}
		}
	}

	return values, nil
}

// calculatePercentile calculates the nth percentile of a slice of float64 values
func calculatePercentile(values []float64, percentile float64) float64 {
	if len(values) == 0 {
		fmt.Println("No values to calculate percentile") // Debug log
		return 0
	}
	// Sort the values
	sort.Float64s(values)
	index := int(float64(len(values)) * percentile / 100)
	if index >= len(values) {
		index = len(values) - 1
	}
	return values[index]
}

func avg(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <namespace> <deployment> <orgID>")
		return
	}

	namespace := os.Args[1]
	deployment := os.Args[2]
	orgID := os.Args[3] // orgID is now passed as the third command-line argument

	// Query CPU usage (average CPU usage over the last 24 hours)
	cpuQuery := fmt.Sprintf(`avg(rate(container_cpu_usage_seconds_total{namespace="%s", pod=~"%s-[a-f0-9]+-[a-z0-9]+$"}[1h])) by (pod)`, namespace, deployment)
	cpuValues, err := queryPrometheus(cpuQuery, orgID)
	if err != nil {
		fmt.Printf("Error querying CPU usage: %v\n", err)
		return
	}

	// Calculate 95th percentile for CPU usage
	cpuRequest := calculatePercentile(cpuValues, 95)

	// Query memory usage (maximum memory usage over the last 24 hours)
	memoryQuery := fmt.Sprintf(`max_over_time(container_memory_usage_bytes{namespace="%s", pod=~"%s-[a-f0-9]+-[a-z0-9]+$"}[24h])`, namespace, deployment)
	memoryValues, err := queryPrometheus(memoryQuery, orgID)
	if err != nil {
		fmt.Printf("Error querying memory usage: %v\n", err)
		return
	}

	// Calculate the maximum memory usage across all pods
	memoryRequest := 0.0
	for _, value := range memoryValues {
		if value > memoryRequest {
			memoryRequest = value
		}
	}

	// Query CPU requests for the deployment
	cpuRequestsQuery := fmt.Sprintf(`kube_pod_container_resource_requests{namespace="%s", pod=~"%s-[a-f0-9]+-[a-z0-9]+$", resource="cpu"}`, namespace, deployment)
	cpuRequests, err := queryPrometheus(cpuRequestsQuery, orgID)
	if err != nil {
		fmt.Printf("Error querying CPU requests: %v\n", err)
		return
	}
	deploymentCPURequest := 0.0
	for _, value := range cpuRequests {
		if value > deploymentCPURequest {
			deploymentCPURequest = value // Take the maximum CPU request for the deployment
		}
	}

	// Query memory requests for the deployment
	memoryRequestsQuery := fmt.Sprintf(`kube_pod_container_resource_requests{namespace="%s", pod=~"%s-[a-f0-9]+-[a-z0-9]+$", resource="memory"}`, namespace, deployment)
	memoryRequests, err := queryPrometheus(memoryRequestsQuery, orgID)
	if err != nil {
		fmt.Printf("Error querying memory requests: %v\n", err)
		return
	}
	deploymentMemoryRequest := 0.0
	for _, value := range memoryRequests {
		if value > deploymentMemoryRequest {
			deploymentMemoryRequest = value // Take the maximum memory request for the deployment
		}
	}
	// Create a table for displaying the results
	tbl := table.New("Metric", "Recommended Request", "Configured Request")
	tbl.AddRow("CPU (millicores)", fmt.Sprintf("%.2fm", cpuRequest*1000), fmt.Sprintf("%.2fm", deploymentCPURequest*1000))
	tbl.AddRow("Memory", fmt.Sprintf("%.2fMi", memoryRequest/1024/1024), fmt.Sprintf("%.2fGi", deploymentMemoryRequest/1024/1024/1024))

	// Print the table
	fmt.Printf("Resource requests for deployment %s in namespace %s:\n", deployment, namespace)
	tbl.Print()
}

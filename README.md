# prometheus-demo

## How to run the application

1. `go run main.go` to start the application.
2. `http://localhost:8090/metrics` to see the metrics in plain text.
3. `prometheus` to start prometheus. I have the alias set in my bash file. `alias prometheus='~/Desktop/prometheus-2.40.1.darwin-amd64/prometheus'`
4. `http://localhost:9090/` to view metrics on prometheus
5. Follow `https://prometheus.io/docs/tutorials/visualizing_metrics_using_grafana/` to start Grafana. Visit `http://localhost:3000/`
6. Click Configuration - Add Data Source - prometheus - url: `http://localhost:9090/`- save
7. Create a Dashboard on Grafana.


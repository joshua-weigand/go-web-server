# go-web-server


### Usage
```Go
func main() {
	var ws = &WebServerBuilder{}

	var endpoint = map[string]RequestHandler{
		"/actuator/info":       nil,
		"/actuator/health":     nil,
		"/actuator/prometheus": nil,
	}

	ws.
		Port("8080").
		BasePath("/v1/api").
		EndPoints(endpoint).
		Run()
}
```
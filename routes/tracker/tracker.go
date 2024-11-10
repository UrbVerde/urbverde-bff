package tracker

var AvailableEndpoints = []map[string]string{}

func AddEndpoint(method, route, description string) {
	AvailableEndpoints = append(AvailableEndpoints, map[string]string{
		"method":      method,
		"route":       route,
		"description": description,
	})
}

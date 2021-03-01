package generate

//go:generate kratos proto client ./api
//go:generate kratos proto client ./internal
//go:generate wire ./cmd/server

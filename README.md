# GRPC Gateway (response modifier)
The project highlights the issue with GRPC Gateway where the response modifier does not get respected.

* `utils.ResponseStatusCodeModifier` is the response modifier to read response code from context.
* It is hooked to grpc server in `main.go` and `main_test.go`
* `user_service` sets the response code 406
* `user_service_test`verifies that the grpc-gateway returns 400 code instead of 406

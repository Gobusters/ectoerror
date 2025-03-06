# ectoerror

Provides helpful errors for Go Services

## HTTPError

The `HTTPError` package provides a robust error handling solution for HTTP services in Go. It offers a structured way to create, wrap, and check HTTP-related errors with status codes and metadata.

### Features

- Create HTTP errors with status codes and messages
- Wrap existing errors with HTTP context
- Add metadata to errors
- Check error types and status codes
- Compatible with Go's error wrapping interface

### Installation

```go
go get github.com/yourusername/ectoerror
```

### Usage

```go
import "github.com/Gobusters/ectoerror/httperror"
// Create a new HTTP error
err := httperror.NewHTTPError(http.StatusBadRequest, "Invalid input")
// Create with formatted message
err := httperror.NewHTTPErrorf(http.StatusBadRequest, "Invalid parameter: %s", paramName)
// Wrap an existing error
if err != nil {
return httperror.WrapError(http.StatusInternalServerError, err)
}
// Add metadata
err.AddMetaValue("user_id", "123")
```

### Error Checking

The package provides various helper functions to check error types:

```go
// Check specific status codes
if httperror.IsBadRequest(err) {
// Handle 400 error
}
if httperror.IsNotFound(err) {
// Handle 404 error
}
// Check error categories
if httperror.IsClientError(err) {
// Handle any 4XX error
}
if httperror.IsServerError(err) {
// Handle any 5XX error
}
if httperror.IsSuccess(err) {
// Handle 2XX status
}
if httperror.IsRedirect(err) {
// Handle 3XX status
}
```

### Available Status Check Functions

- `IsOK(err)` - Status 200
- `IsCreated(err)` - Status 201
- `IsAccepted(err)` - Status 202
- `IsNoContent(err)` - Status 204
- `IsBadRequest(err)` - Status 400
- `IsUnauthorized(err)` - Status 401
- `IsForbidden(err)` - Status 403
- `IsNotFound(err)` - Status 404
- `IsInternalServerError(err)` - Status 500
- `IsBadGateway(err)` - Status 502
- `IsServiceUnavailable(err)` - Status 503
- `IsGatewayTimeout(err)` - Status 504

### Error Structure

The `HTTPError` type includes:

- `Code` - HTTP status code
- `Message` - Error message
- `Meta` - Map for additional metadata
- Underlying error (accessible via `errors.Unwrap()`)

### Advanced Usage

```go
// Check for specific status code
if httperror.IsStatus(err, http.StatusTeapot) {
    // Handle 418 error
}
// Get status code from error
code := httperror.GetStatusCode(err)
// Check if error is an HTTPError
if httperror.IsHTTPError(err) {
    // Handle HTTP error
}
// Access metadata
if httpErr := httperror.ToHTTPError(err); httpErr != nil {
    userID := httpErr.Meta["user_id"]
}
```

### Error Categories

- `IsClientError(err)` - Checks for 4XX status codes
- `IsServerError(err)` - Checks for 5XX status codes
- `IsError(err)` - Checks for any error status (4XX or 5XX)
- `IsSuccess(err)` - Checks for 2XX status codes
- `IsRedirect(err)` - Checks for 3XX status codes

# Go API client for userapi

This is an **example** API to demonstrate features of OpenAPI specification
# Introduction
This API definition is intended to to be a good starting point for describing your API in 
[OpenAPI/Swagger format](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md).
It also demonstrates features of [create-openapi-repo](https://github.com/Redocly/create-openapi-repo) tool and 
[Redoc](https://github.com/Redocly/Redoc) documentation engine. So beyond the standard OpenAPI syntax we use a few 
[vendor extensions](https://github.com/Redocly/Redoc/blob/master/docs/redoc-vendor-extensions.md).

# OpenAPI Specification
The goal of The OpenAPI Specification is to define a standard, language-agnostic interface to REST APIs which
allows both humans and computers to discover and understand the capabilities of the service without access to source
code, documentation, or through network traffic inspection. When properly defined via OpenAPI, a consumer can 
understand and interact with the remote service with a minimal amount of implementation logic. Similar to what
interfaces have done for lower-level programming, OpenAPI removes the guesswork in calling the service.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.10.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [http://example.com/contact](http://example.com/contact)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import userapi "github.com/achan-cogito/user-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `userapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), userapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `userapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), userapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `userapi.ContextOperationServerIndices` and `userapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), userapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), userapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://example.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*UserAPI* | [**GetUserByName**](docs/UserAPI.md#getuserbyname) | **Get** /users/{username} | Get user by user name
*UserAPI* | [**UpdateUser**](docs/UserAPI.md#updateuser) | **Put** /users/{username} | Updated user


## Documentation For Models

 - [User](docs/User.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### main_auth


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: http://example.com/api/oauth/dialog
- **Scopes**: 
 - **read:users**: read users info
 - **write:users**: modify or remove users

Example

```go
auth := context.WithValue(context.Background(), userapi.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, userapi.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### api_key

- **Type**: API key
- **API key parameter name**: api_key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: api_key and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		userapi.ContextAPIKeys,
		map[string]userapi.APIKey{
			"api_key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### basic_auth

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), userapi.ContextBasicAuth, userapi.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

contact@example.com

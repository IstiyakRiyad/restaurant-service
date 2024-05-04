## Testing

### Format of the testing function.
This function start with the `Test` and it takes `testing.T` as input.
```go
func Test`Name of the function`(t *testing.T) {}
```

### NewRecorder()
Here the httptest.NewRecorder() function return ResponseRecorder which will store the testing response. It just a format of writer header in http standard package in golang.
```go
w := httptest.NewRecorder()
```

### NewRequest()
The http.NewRequest() just only build Request struct but doesn't send the request to the server. It just a struct that store the req fields.
```go
	req, _ := http.NewRequest("GET", "/api/v1/restaurant/", nil)
```

### ServeHTTP()
Here the Router.ServeHTTP() function is the required function for Hander. It get used by the http package to route the request. This is the function of gin package which is the entry point of the gin library when a request comes to the package. It takes two struct one has the request info and one where the http hander will write the response json or other things.
```go
	trasport.Router.ServeHTTP(w, req)
```

### Checking function
The assert.Equal() check the the response from the http router.
```go
	assert.Equal(t, 200, w.Code)
```




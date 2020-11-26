# brazilian-utils

A Golang library containing brazilian-specific utils.

![build](https://github.com/fernandoporazzi/brazilian-utils/workflows/build/badge.svg)

## Testing

To run the tests locally, run the following command:

```sh
$ go test ./... 
```

If you want to see how covered the project is, you can run the following command to get coverage report
```sh
$ go test ./... -coverprofile=coverage.out
```

Once the above has been run, it's time to see it in your browser. The following command will open a new tab in your browser with the code coverage.

```sh
$ go tool cover -html=coverage.out
```

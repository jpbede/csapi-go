# csapi-go
[![PkgGoDev](https://pkg.go.dev/badge/go.bnck.me/csapi)](https://pkg.go.dev/go.bnck.me/csapi)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/e937f5341cc046b593cf44eb1f891087)](https://www.codacy.com/gh/jpbede/csapi-go/dashboard)
[![codecov](https://codecov.io/gh/jpbede/csapi-go/branch/main/graph/badge.svg?token=JMrbj90oHv)](https://codecov.io/gh/jpbede/csapi-go)
![test](https://github.com/jpbede/csapi-go/workflows/test/badge.svg)

Go client package for the CSApi (API for certified-senders.org)

## Usage

Import the lib as usual
```go
import "go.bnck.me/csapi"
```

Then create a new client. The login will be handled automatically. You will get a authenticated client.
```go
api, err := csapi.NewWithOptions(csapi.WithCredentials("abc12", "abc-123"))
if err != nil {
    return err
}
```
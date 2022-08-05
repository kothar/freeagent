# FreeAgent API Client for Go

* Currently a very bare-bones client implementation.
* Only basic operations on a few resources supported.

## Usage

```go
fa := &freeagent.Client{
    Endpoint:    freeagent.SandboxEndpoint,
    AccessToken: accessToken,
}

timeslip, err := fa.GetTimeslip("12345")
if err != nil {
    log.Fatal(err)
}
```
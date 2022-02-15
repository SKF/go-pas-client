# go-pas-client

This is the client library for the Point Alarm Status (PAS) API.

Refer to the Swagger [specification](https://api.point-alarm-status.sandbox.iot.enlight.skf.com/v1/docs/swagger/index.html) and the [service description](https://api.point-alarm-status.sandbox.iot.enlight.skf.com/v1/docs/service) for further information on how the API works.

## Example

Refer to [example/](/example/) for examples of how to use this library.

## Error handling

Errors returned by the API are decoded into problems from the [`github.com/SKF/go-rest-utility/problems`](https://github.com/SKF/go-rest-utility) package before being returned by the client functions. This makes it possible to use the standard [`error`](https://pkg.go.dev/errors) package to do error checking on any returned error.

E.g. if a faulty threshold is supplied when setting an alarm threshold you can check if it's a validation error (and what the validation errors are) using the following code.

```go
err := client.SetThreshold(ctx, nodeID, models.Threshold{
  Overall: &models.Overall{
    Unit: "C",
    OuterHigh: 60,
    InnerHigh: 70,
  },
})
if err != nil {
  var problem problems.ValidationProblem

  if errors.As(err, &problem) {
    for _, reason := range problem.Reasons {
      fmt.Println(reason)
    }
  }

  var problem problems.BasicProblem

  if errors.As(err, &problem) {
    switch problem.Status {
      case http.StatusInternalServerError:
    }
  }
}
```

## Events

Events sent by the PAS service (as documented [here](https://api.point-alarm-status.sandbox.iot.enlight.skf.com/v1/docs/service/sns)) can be decoded into types defined in [models/events.go](/models/events.go).

The input to each decoding function should be the content of the event record `Data` field.

```go
var record eventsource.Record

event := models.AlarmStatusEvent{}

err := event.FromEvent(record.Data)
if err != nil {
  panic(err)
}
```

The intention with these event types is to provide functionality that convert the event types into the same models returned by the client API calls.

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

The intention with these event types is to provide functionality that convert the event types sent on the SNS topic, which have have a slightly different structure compared to when fetching e.g. alarm thresholds from the API, into the same models returned by the client API calls. This should make the event models easier to use together with other client functions.

Below is an example which catches changes to overall thresholds where the measurement unit is set to `C°`, and modifies it to `C`.

```go
var record eventsource.Record

event := new(models.ThresholdEvent)

if err := event.FromEvent(record.Data); err != nil {
  panic(err)
}

threshold := event.Threshold

if threshold.Overall != nil && threshold.Overall.Unit == "C°" {
  threshold.Overall.Unit = "C"

  client.SetThreshold(ctx, nodeID, threshold)
}

```

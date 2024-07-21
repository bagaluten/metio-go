# Metio 

Congratualtions you found the Go Metio Library which can be used to interact with Metio Compliant Systems.

<img src="https://github.com/user-attachments/assets/5ff9115c-6c6c-404b-9a2a-dcc221f84479" height="300">


## Add it to your project

```bash
go get github.com/bagaluten/metio-go
```


## Features

- [x] OpenTelemetry Tracing
- [x] Metio Streams
- [x] Metio Client


## How to use it 

```golang
	ctx := context.Background()
	client, err := client.NewClient(client.Config{Host: "localhost:4222", Prefix: nil})
	require.NoError(t, err)

	defer client.Close()

	stream := streams.NewStream("stream", client)
	events := []types.Event{
		{
			EventID:   "123",
			ContextID: nil,
			EventType: types.MustParseEventType("group/name/version"),
			Payload: types.Payload{
				"key": "value",
			},
			Timestamp: types.TimeNow(),
		},
		{
			EventID:   "124",
			ContextID: nil,
			EventType: types.MustParseEventType("group/name/version"),
			Payload: types.Payload{
				"key": "value",
			},
			Timestamp: types.TimeNow(),
		},
	}

	err = stream.Publish(ctx, events)

```

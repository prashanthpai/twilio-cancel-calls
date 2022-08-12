# twilio-cancel-calls

Cancel queued calls on Twilio.

## Usage

### Python (sequential)

```sh
$ pip install twilio
$ TWILIO_ACCOUNT_SID=ACxxxxxxxxxx TWILIO_AUTH_TOKEN=yyyyyyyyyy ./main.py
```

### Go (concurrent)

```sh
$ go build
$ TWILIO_ACCOUNT_SID=ACxxxxxxxxxx TWILIO_AUTH_TOKEN=yyyyyyyyyy ./twilio-cancel-calls
```

## References

* [Read multiple Call resources](https://www.twilio.com/docs/voice/api/call-resource#read-multiple-call-resources)
* [Update a Call resource](https://www.twilio.com/docs/voice/api/call-resource#update-a-call-resource)

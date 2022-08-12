#!/usr/bin/env python3

import os
from twilio.rest import Client


account_sid = os.environ['TWILIO_ACCOUNT_SID']
auth_token = os.environ['TWILIO_AUTH_TOKEN']
client = Client(account_sid, auth_token)


calls = client.calls.stream(status='queued')
for record in calls:
    call = client.calls(record.sid).update(status='canceled')
    print(record.sid)

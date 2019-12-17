# Simple webinterface for view alerts from PagerDuty

**Usage:**

Before you start programm you need to create API key for PD (https://support.pagerduty.com/docs/generating-api-keys), read-only will be enough
Then set schedule for OnCall section (https://support.pagerduty.com/docs/schedules)
That's all :)
```
rkrankus@wifunny:~/go/src/pdviewer$ PDTOKEN=${YOUR_TOKEN_HERE} PDSHEDULES=${YOU_SCHEDULE} PDVIEWERLISTEN=0.0.0.0:8080 go run *.go
2019/12/17 11:45:12 Listening on 0.0.0.0:8080
```

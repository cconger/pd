Usage:
Remotely:
```go install github.com/cconger/pd```

In this directory: 
```go install```
```$GOPATH/bin/pd --token="<TOKEN>"```

Pre-Code Thinking:

  Original Problem Statement:
  Write a script that lists the created date, assigned user, and status for each of the 10 most recently created incidents, in descending date order.

  Since my approach for this problem is dependent on the use case for this problem.

  If this is just a query I run a lot then I would simply do:

  ```curl https://webdemo.pagerduty.com/api/v1/incidents -H 'Authorization: Token token=$TOKEN' -X GET \
  --data-urlencode "fields=created_on,assigned_to_user,status" \
  --data-url-encode "sort_by=created_on:desc" \
  --data-urlencode "limit=10" | jxt incidents | less```

  jxt is a utility that I have written in python for doing simple parsing, querying and pretty printing of json.  But you could also use the json.tool in python. ```python -mjson.tool```

  My main approach is a split between doing a command line utility in Go.  This would allow relatively easy distribution as a binary, however I would be apprehensive doing this in a non-go shop.  Additionally, if it were to become part of the web toolchain for a dashboard or similar utility perhaps making it a JS library with a node front end might make more sense.  If I build it this way it would be nice to stub out the network requests directly so that it can function in both node as well as a ajax request. If however I built the commandline go utility I would want it to print in a human readable and pipe-friendly format.

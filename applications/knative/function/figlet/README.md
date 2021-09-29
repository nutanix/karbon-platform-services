# Go HTTP Function

Welcome to your new Go Function! The boilerplate function code can be found in
[`handle.go`](handle.go). This Function responds to HTTP requests.

The function returns:
- OK in ASCII art on a GET request
```
curl x.x.x.x:80 -H "Host: figlet.knative-figlet.example.com"
   ___   .  _  __
  / _ \  . | |/ /
 | (_) | . | ' <
  \___/  . |_|\_\
         .
```
- value of the `text` key in ASCII art  on a POST request
```
curl x.x.x.x:80 -H "Host: figlet.knative-figlet.example.com" -H "Content-Type: application/json" -d '{"text":"test"}'
  _____  .  ___  .  ___  .  _____
 |_   _| . | __| . / __| . |_   _|
   | |   . | _|  . \__ \ .   | |
   |_|   . |___| . |___/ .   |_|
         .       .       .
```

## Development

Develop new features by adding a test to [`handle_test.go`](handle_test.go) for
each feature, and confirm it works with `go test`.

Update the running analog of the function using the `func` CLI.

For more, see [the complete documentation]('https://github.com/knative-sandbox/kn-plugin-func/tree/main/docs')

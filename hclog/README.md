sample hclog
====================

```go
$ go run main.go
2023/02/26 22:23:25 golang default logging
2023-02-26T22:23:25.875+0900 [DEBUG] stdapp: 42
2023-02-26T22:23:25.875+0900 [INFO]  func: this message will print: request=5fb446b6-6eba-821d-df1b-cd7501b6a363
2023-02-26T22:23:25.875+0900 [INFO]  func: message with key-value: request=5fb446b6-6eba-821d-df1b-cd7501b6a363 key=42
```

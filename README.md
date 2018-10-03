# go-logging
Sample log capture using GO

## Prequest
* GO
* Logrus

## Dependencies
* Dependencies Management: [dep](https://github.com/golang/dep)
* Logging: [logrus](https://github.com/Sirupsen/logrus)

## Install Dependencies
* Install Package Manager:
  ```
  go get -u -d github.com/golang/dep/cmd/dep
  ```
* Install PostgreSQL Driver
  ```
  go get -u -d github.com/Sirupsen/logrus
  ```

## Using Package Manager
```
$ dep ensure -v
```

## Running
```
go run app.go
```

## License
MIT License
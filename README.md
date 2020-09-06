# jobHandler
An API build on GoLang for long running jobs which can pause, resume and terminate a specific job.

## Table of Content
- [jobHandler](#jobhandler)
  - [Table of Content](#table-of-content)
  - [Installation](#installation)
    - [Buind and run from source](#buind-and-run-from-source)
    - [Building and run in Docker](#building-and-run-in-docker)
  - [Usage](#usage)
    - [Endpoints](#endpoints)
  - [Development](#development)
  - [License](#license)

## Installation
### Buind and run from source
Get the source using go get
```bash
$ go get github.com/someshkoli/jobHandler

$ $GOPATH/bin/jobHandler
```

### Building and run in Docker
Build docker image
```
docker build -t jobhandler
```
Runnin Docker Image
```
docker run -it -p 8080:8000 jobhandler
```


## Usage
- The server by default serves at http port 8000

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/d46c526c14e9beff7940#?env%5BAtlan%20Pipeline%5D=W3sia2V5IjoidXJsIiwidmFsdWUiOiIxNzIuMTcuMC4yOjgwMDAiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6ImlkIiwidmFsdWUiOiJyYW5kb21cIiIsImVuYWJsZWQiOnRydWV9XQ==)

### Endpoints
- **/ping**
```bash
curl --location --request GET '172.17.0.2:8000/ping'

{
    "Status": true,
    "Data": "Pong"
}
```
- **/upload**
```bash
curl --location --request POST '172.17.0.2:8000/upload' \
--form 'file=@pathToCSV.csv'

{
    "Status": true,
    "Data": "3f7d7db0-8054-445b-b2b3-78cd4b407dda"
}
```

- **/pause**
```bash
curl --location --request GET '172.17.0.2:8000/pause?id=3f7d7db0-8054-445b-b2b3-78cd4b407dda'

{
    "Status": true,
    "Data": "job paused successfully"
}
```

- **/resume**
```bash
curl --location --request GET '172.17.0.2:8000/resume?id=3f7d7db0-8054-445b-b2b3-78cd4b407dda'

{
    "Status": true,
    "Data": "job resumed successfully"
}
```

- **/terminate**
```bash
curl --location --request GET '172.17.0.2:8000/terminate?id=3f7d7db0-8054-445b-b2b3-78cd4b407dda'

{
    "Status": true,
    "Data": "job paused successfully"
}
```

- **/status**
```bash
curl --location --request GET '172.17.0.2:8000/status?id=3f7d7db0-8054-445b-b2b3-78cd4b407dda'

{
    "Status": true,
    "Data": "Terminated"
}
```
## Development

- This command will fetch the source to your local system
```bash
$ go get github.com/someshkoli/jobHandler
```
This will be saved in your `$GOPATH/src/github.com/someshkoli/jobHandler`

- Downloading dependencies
```bash
$ go mod download
```

- Running the server
```bash
$ go run main.go
```

## License
This software is licensed under Apache-2.0. Copyright Postman, Inc. See the [LICENSE](LICENSE) file for more information.
# Async logger
Async logger is simple pkg to write logs to file or output to std asynchronously

## Usage

Check the [Makefile](./Makefile) for more information.

```
Usage: make [target]

Targets:
  build                Compile binaries.
  tests                Run tests.
```

## Configuration

Check [/configs](./configs) for description and some examples.

## Build

To build application locally, just run:

```shell
make build
```

The binaries will be placed in `/bin`.

### Docker

To build application inside a Docker container, choose the `Dockerfile` and run:

```shell
docker build -f ${Dockerfile} -t ${IMAGE_NAME} .
```

### Docker Compose

To build application with Docker Compose (with required services like database and etc.), run:

```shell
docker compose -f docker-compose.yml up --build 
```

## Testing

To run only unit tests simply run:

```shell
make test
```

## Run app

### Env

To run app you should export env `CONFIG_FILE`:

```shell
export CONFIG_FILE=./configs/config.yaml 
```

And run app with args:

```shell
./bin/alog_test 3 50
```

In this example we run 3 threads and write 50 messages in each thread.
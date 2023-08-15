## Example of sourcing, parsing and persisting data with go

Get products data from three different sources, parse it to single format and persist it in kafka.

### Setup

```bash
docker-compose up -d
```
Open kafka-ui on http://localhost:9091/

### Usage

```bash
make run-app
```
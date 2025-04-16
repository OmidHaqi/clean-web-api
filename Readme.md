# Golang Clean Web API (Dockerized)

A complete sample project demonstrating a clean architecture approach for a car sales web API with Golang, featuring comprehensive logging, monitoring, and containerization.


## Table of Contents

- [System Architecture](#system-architecture)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Running Locally](#running-locally)
  - [Running with Docker](#running-with-docker)
- [API Examples](#api-examples)
- [Monitoring and Observability](#monitoring-and-observability)
- [Production Deployment](#production-deployment)
- [Project Preview](#project-preview)
- [Contributing](#contributing)
- [Contact Information](#contact-information)

## System Architecture

<p align="center"><img src='/docs/files/system_diagram.png' alt='Golang Web API System Design Diagram' /></p>

## Database Design

<p align="center"><img src='/docs/files/db_diagram.png' alt='Golang Web API Database Design Diagram' /></p>

## Technologies Used

This project leverages a modern technology stack to achieve robust functionality:

1. **Web Framework**: [Gin](https://github.com/gin-gonic/gin) - High-performance HTTP web framework
2. **Authentication**: [JWT](https://github.com/golang-jwt/jwt) - Token-based authentication and authorization
3. **Caching**: [Redis](https://github.com/redis/redis) - In-memory data structure store
4. **Logging Stack**:
   - [Elasticsearch](https://github.com/elastic/elasticsearch) - Log storage and indexing
   - [Filebeat](https://github.com/elastic/beats) - Log shipping
   - [Kibana](https://github.com/elastic/kibana) - Log visualization
5. **Database**:
   - [PostgreSQL](https://github.com/postgres/postgres) - Main database engine
   - [PgAdmin](https://github.com/pgadmin-org/pgadmin4) - Database management tool
   - [GORM](https://github.com/go-gorm/gorm) - ORM library
6. **Monitoring**:
   - [Prometheus](https://github.com/prometheus/prometheus) - Metrics collection
   - [Grafana](https://github.com/grafana/grafana) - Metrics visualization
   - [Alertmanager](https://github.com/prometheus/alertmanager) - Alert handling
7. **Development Tools**:
   - [Validator](https://github.com/go-playground/validator) - Input validation
   - [Viper](https://github.com/spf13/viper) - Configuration management
   - [Zap](https://github.com/uber-go/zap) & [Zerolog](https://github.com/rs/zerolog) - Structured logging
   - [Swaggo](https://github.com/swaggo/swag) - API documentation
8. **Containerization**: Docker & Docker Compose - Application packaging and deployment

## Getting Started

### Running Locally

#### 1. Start Dependencies with Docker

```bash 
docker compose -f "docker/docker-compose.yml" up -d setup elasticsearch kibana filebeat postgres pgadmin redis prometheus node-exporter alertmanager grafana
```

#### 2. Install Swagger and Run the Application

```bash
cd src
go install github.com/swaggo/swag/cmd/swag@latest
cd cmd
go run main.go
```

The API will be available at: [http://localhost:5005](http://localhost:5005)

#### 3. Stop Dependencies

```bash
docker compose -f "docker/docker-compose.yml" down
```

### Running with Docker

Deploy the entire application stack with a single command:

```bash
docker compose -f "docker/docker-compose.yml" up -d --build
```

#### Service Access Information

| Service | URL | Credentials |
|---------|-----|-------------|
| Web API | [http://localhost:9001](http://localhost:9001) | Username: `admin`<br>Password: `12345678` |
| Kibana | [http://localhost:5601](http://localhost:5601) | Username: `elastic`<br>Password: `@aA123456` |
| Prometheus | [http://localhost:9090](http://localhost:9090) | N/A |
| Grafana | [http://localhost:3000](http://localhost:3000) | Username: `admin`<br>Password: `foobar` |
| PgAdmin | [http://localhost:8090](http://localhost:8090) | Username: `omid.haqi@outlook.com`<br>Password: `123456` |

**Postgres Connection Details**:
- Host: `postgres_container`
- Port: `5432`
- Username: `postgres`
- Password: `admin`

#### Stopping Docker Services

```bash
docker compose -f 'docker/docker-compose.yml' --project-name 'docker' down
```

## API Examples

### Authentication

Login example:

```bash
curl -X 'POST' \
  'http://localhost:5005/api/v1/users/login-by-username' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "password": "12345678",
  "username": "admin"
}'
```

### Filtering and Sorting

#### City Filter with Text Search

```json
{
  "filter": {
    "Name": {
      "filterType": "text",
      "from": "t", 
      "type": "contains"
    } 
  },
  "pageNumber": 1,
  "pageSize": 10,
  "sort": [
    {
      "colId": "name",
      "sort": "desc"
    }
  ]
}
```

#### City with Range Filter

```json
{
  "filter": {
    "Id": {
      "filterType": "number",
      "from": "1", 
      "to": "7", 
      "type": "inRange"
    } 
  },
  "pageNumber": 1,
  "pageSize": 10,
  "sort": [
    {
      "colId": "name",
      "sort": "desc"
    }
  ]
}
```

## Production Deployment

### Linux Deployment with Systemd

1. **Build the Project**

```bash
cd src
go build -o ../prod/server ./cmd/main.go
mkdir -p ../prod/config/ && cp config/config-production.yml ../prod/config/config-production.yml
```

2. **Create a Systemd Service Unit**

```bash
sudo vi /lib/systemd/system/go-api.service
```

3. **Configure the Service**

```
[Unit]
Description=go-api

[Service]
Type=simple
Restart=always
RestartSec=20s
ExecStart=/development/Project/go/clean-web-api/prod/server
Environment="APP_ENV=production"
WorkingDirectory=/development/Project/go/clean-web-api/prod

[Install]
WantedBy=multi-user.target
```

4. **Start the Service**

```bash
sudo systemctl start go-api
```

5. **Stop the Service**

```bash
sudo systemctl stop go-api
```

6. **View Service Logs**

```bash
sudo journalctl -u go-api -e
```

## Project Preview

### Swagger

<p align="center"><img src='/docs/files/swagger.png' alt='Golang Web API preview' /></p>

### Grafana

<p align="center"><img src='/docs/files/grafana.png' alt='Golang Web API grafana dashboard' /></p>

### Kibana

<p align="center"><img src='/docs/files/kibana.png' alt='Golang Web API grafana dashboard' /></p>

## Contributing

Contributions are welcome! Please feel free to submit issues, fork the repository, and create pull requests.

## Contact Information

If you have any questions, suggestions, or feedback regarding this project, please feel free to reach out:

- **Email**: omid.haqi.dev@gmail.com
- **GitHub Issues**: Please use the [issue tracker](https://github.com/omidhaqi/clean-web-api/issues)
- **LinkedIn**: [Omid Haqi](https://linkedin.com/in/Omid-haghi)
- **Telegram**: [Umut](https://t.me/Omid_Haqi)

## Project Structure

- `cmd`:
  - `cli/kafka`:
    - `kafka.go`: Entry point for runnning Kafka Service(Not Implemented in Application)
  - `server`:
    - `main.go`: Main entry point for running the application
- `global`:
  - `global.go`: Stores globally accessible service instances and configurations.
- `internal`: Contains all application-specific logic and services.

- `pkg`:

- `python`: Contains Python-based microservices and related scripts
- `response`: Go-based custom response utilities for standardized client responses(Not Implemented)
- `third_party`:
  - `docker`: Contains Docker setups for external services like Nginx, Kafka, and Redis
    - `nginx`: Nginx Dockerfile and configuration for Load Balancing
      - `Dockerfile`: Docker for Nginx
      - `nginx.conf`: Nginx Configuration
    - `postgres`: Postgres Dockerfile and Schema migration
      - `migration`: Schema Migration
        - `000001_init_schema.down.sql`: Drop all tables
      - `Dockerfile`: Docker for Postgres
    - `redis`: Redis
      - `Dockerfile`: Docker for Redis
- `makefile`:
- `go.mod`:
- `sqlc.yaml`:


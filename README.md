# go-premierleaguebooking


## Technical stack

- Backend building blocks
  - [grpc-ecosystem/grpc-gateway/v2](https://github.com/grpc-ecosystem/grpc-gateway)
  - [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc)
    - [pq](github.com/lib/pq)
  - [golang-jwt/jwt/v4](https://github.com/golang-jwt/jwt)
  - [hibiken/asynq](https://github.com/hibiken/asynq)
  - Utils
    - google/uuid
    - google.golang.org/genproto
    - google.golang.org/grpc
    - google.golang.org/protobuf
- Infrastructure
  - Postgres, Redis
  - Google Cloud Platform(GCP), Google Kubernetes Engine (GKE)
  - docker and docker-compose
  - devcontainer for reproducible development environment

## Premier League Booking


## Screenshots

### Home screen

![home_screen](docs/homepage.png)

### View Tickets screen

![view_tickets](docs/view_tickets.png)

### Proceed Payment screen

![view_tickets](docs/proceed_payment.png)

### Review Checkout screen

![view_tickets](docs/review_checkout.png)

### Payment screen

![view_tickets](docs/payment.png)

### Review Checkout and Payment screen

![review_checkout_and_payment](docs/review_checkout.png)

## Cache Strategies Performance Comparison

![cache_performance_comparison](docs/caching_layers.png)


## Credits
- [readme-structure](https://github.com/thangchung/go-coffeeshop/blob/main/README.md)
- [project-layout](https://github.com/golang-standards/project-layout)
- [repository-structure](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
- [go-build-template](https://github.com/thockin/go-build-template)
- [go-clean-template](https://github.com/evrone/go-clean-template)
- [emsifa/tailwind-pos](https://github.com/emsifa/tailwind-pos)
- [thanhchung](https://github.com/thangchung/go-coffeeshop)

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

![view_tickets](docs/viewtickets.png)

### Proceed Payment screen

![proceed_payment](docs/proceedpayment.png)

### Review Checkout screen

![checkout](docs/checkout.png)

### Payment screen

![payment](docs/payment.png)

### Review Checkout

![payment_success](docs/paymentsuccess.png)

### Payment screen
[Stripe Test Card Numbers](https://docs.stripe.com/testing)

![review_checkout_and_payment](docs/review_checkout.png)



## Cache Strategies Performance Comparison

![cache_performance_comparison](docs/caching_layers.png)


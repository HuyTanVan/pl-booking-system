package global

import (
	db "plbooking_go_structure1/internal/db/sqlc"
	"plbooking_go_structure1/internal/token"
	"plbooking_go_structure1/pkg/setting"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config        setting.Config
	Rdb           *redis.Client // redis server
	Pgdbc         db.Store      // postgresql sqlc
	TokenMaker    token.IMaker  // jwt token generator
	KafkaProducer *kafka.Writer // kafka writer
	// ErrGroup      *errGroup.Group
	// postgresql using sqlc
	// DBStore db.Store // Database using SQLC
)

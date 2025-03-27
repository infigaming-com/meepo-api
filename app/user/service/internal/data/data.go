package data

import (
	"github.com/infigaming-com/go-common/lock"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewPgGorm,
	NewRedisLock,
	NewUserRepo,
	NewUserLoginRepo,
)

// Data .
type Data struct {
	lock lock.Lock
	db   *gorm.DB
}

// NewData .
func NewData(lg log.Logger, db *gorm.DB, lock lock.Lock) (*Data, func(), error) {
	d := &Data{
		lock: lock,
		db:   db,
	}

	return d, func() {}, nil
}

func NewPgGorm(c *conf.Data, lg log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(lg, "module", "user-service/data/pggorm"))

	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	return db
}

func NewRedisLock(c *conf.Data, lg log.Logger) lock.Lock {
	log := log.NewHelper(log.With(lg, "module", "user-service/data/redislock"))

	lock, _, err := lock.NewRedisLock(
		lock.WithRedisLockAddr(c.RedisLock.Addr),
		lock.WithRedisLockDB(c.RedisLock.Db),
		lock.WithRedisLockConnectTimeout(int64(c.RedisLock.ConnectTimeout.AsDuration())),
	)
	if err != nil {
		log.Fatalf("failed to create redis lock: %v", err)
	}
	return lock
}

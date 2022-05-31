package helpers

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/karim-w/pika-go/socket/internal/common/models"
	"go.uber.org/fx"
)

type RedisCache interface {
	//PUBLIC FUNCTIONS
	HandleUserJoin(u *models.RedisUser)
	HandleUserExit(id string)
	FetchUser(id string) *models.RedisUser
	//PRIVATE FUNCTIONS
	addUserToRedis(u *models.RedisUser)
	addUserToRedisWG(u *models.RedisUser, wg *sync.WaitGroup)
	addUserToSetWG(setName string, id string, wg *sync.WaitGroup)
	deleteFromRedisWG(id string, wg *sync.WaitGroup)
	deleteFromSetWG(setName string, id string, wg *sync.WaitGroup)
}

type redisCacheImpl struct {
	db *redis.Client
}

var _ RedisCache = (*redisCacheImpl)(nil)

func RedisCacheProvider() RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb.Set("PIKACHU-SERVICE", "Online", 0)
	return &redisCacheImpl{
		db: rdb,
	}
}

var RedisCacheModule = fx.Option(fx.Provide(RedisCacheProvider))

func (r *redisCacheImpl) HandleUserJoin(u *models.RedisUser) {
	awaits := len(u.Hearings) + len(u.Sessions) + 1
	wg := sync.WaitGroup{}
	wg.Add(awaits)
	go r.addUserToRedisWG(u, &wg)
	for i := range u.Sessions {
		go r.addUserToSetWG(u.Sessions[i], u.Id, &wg)
	}
	for j := range u.Hearings {
		go r.addUserToSetWG(u.Hearings[j], u.Id, &wg)
	}
	wg.Wait()
}

func (r *redisCacheImpl) addUserToRedis(u *models.RedisUser) {
	r.db.Set(u.Id, *u, time.Hour)
}

func (r *redisCacheImpl) addUserToRedisWG(u *models.RedisUser, wg *sync.WaitGroup) {
	r.addUserToRedis(u)
	wg.Done()
}

func (r *redisCacheImpl) addUserToSetWG(setName string, id string, wg *sync.WaitGroup) {
	r.db.SAdd(setName, id)
	wg.Done()
}

func (r *redisCacheImpl) HandleUserExit(id string) {
	user := r.FetchUser(id)
	if user != nil {
		awaits := len(user.Hearings) + len(user.Sessions) + 1
		wg := sync.WaitGroup{}
		wg.Add(awaits)
		go r.deleteFromRedisWG(id, &wg)
		for i := range user.Sessions {
			go r.deleteFromSetWG(user.Sessions[i], id, &wg)
		}
		for j := range user.Hearings {
			go r.deleteFromSetWG(user.Hearings[j], id, &wg)
		}
		wg.Wait()
	}
}

func (r *redisCacheImpl) deleteFromRedisWG(id string, wg *sync.WaitGroup) {
	r.db.Del(id)
	wg.Done()
}

func (r *redisCacheImpl) deleteFromSetWG(setName string, id string, wg *sync.WaitGroup) {
	r.db.SRem(setName, id)
	wg.Done()
}

func (r *redisCacheImpl) FetchUser(id string) *models.RedisUser {
	u := r.db.Get(id).Val()
	if u != "" {
		user := models.RedisUser{}
		json.Unmarshal([]byte(u), &user)
		return &user
	} else {
		return nil
	}
}

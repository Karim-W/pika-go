package helpers

import (
	"net"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

type MemCache interface {
  AddConnection(userId string, connection net.Conn)
  RemoveConnection(userId string)
  FetchConnection(userId string) *net.Conn
	FetchConnectionWG(userId string,wg *sync.WaitGroup)*net.Conn
	FetchConnectionAsync(userId string, wg *sync.WaitGroup,res chan *net.Conn)
	FetchConnectionsAsync(users []string) *[]*net.Conn
}

type memCacheImpl struct{
  db *cache.Cache
}

var _ MemCache = (*memCacheImpl)(nil)

func MemCacheProvider() MemCache {
  return &memCacheImpl{db: cache.New(3 * time.Hour, 1 * time.Hour)}
}

func (m *memCacheImpl) AddConnection(userId string, connection net.Conn) {
	m.db.Set(userId, connection, cache.NoExpiration)
}

func (m *memCacheImpl) RemoveConnection(userId string) {
	m.db.Delete(userId)
}

func (m *memCacheImpl) FetchConnection(userId string) *net.Conn {
	if conn, ok := m.db.Get(userId); ok {
		if conn != nil {
			c := conn.(net.Conn)
			return &c
		}else{
			return nil
		}
	}
	return nil
}

func (m *memCacheImpl) FetchConnectionWG(userId string,wg *sync.WaitGroup) *net.Conn {
	defer wg.Done()
	return m.FetchConnection(userId)
}

func (m *memCacheImpl) FetchConnectionAsync(userId string,wg *sync.WaitGroup,res chan *net.Conn){
	res <- m.FetchConnectionWG(userId,wg)
}

func (m *memCacheImpl) FetchConnectionsAsync(users []string) *[]*net.Conn{
	wg := sync.WaitGroup{}
	wg.Add(len(users))
	usersChan := make(chan *net.Conn,len(users))
	for _,user := range users {
		go func(user string) {
			go m.FetchConnectionAsync(user,&wg,usersChan)
		}(user)
	}

	userList := []*net.Conn{}
	for  range users{
		userList = append(userList, <-usersChan)
	}
	return &userList
}


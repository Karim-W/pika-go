package connections

import "go.uber.org/zap"

type ConnectionManager interface{

}

type connectionManager struct {
  logger *zap.SugaredLogger
//memCache*
}

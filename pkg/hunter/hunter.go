package hunter

import (
	"github.com/Haiss2/binance-futu-mm/pkg/storage"
	"go.uber.org/zap"
)

type Hunter struct {
	l  *zap.SugaredLogger
	db *storage.Storage
}

func NewHunter(l *zap.SugaredLogger, db *storage.Storage) *Hunter {
	h := &Hunter{
		l:  l,
		db: db,
	}
	return h
}

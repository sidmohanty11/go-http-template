package storage

import "mygoapp/types"

type Storage interface {
	Get(int) *types.User
}

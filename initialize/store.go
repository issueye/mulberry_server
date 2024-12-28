package initialize

import (
	"mulberry/host/common/store"
	"mulberry/host/global"
	"path/filepath"
)

func InitStore() {
	path := filepath.Join(global.ROOT_PATH, "data", "store")
	store, err := store.NewPebbleStore(path)
	if err != nil {
		panic(err)
	}

	global.STORE = store
}

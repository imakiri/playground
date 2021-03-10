package data

import "github.com/imakiri/playground/core"

type Cookie struct {
	storage core.Storage
	// Опции, конфиги, и тт
}

func (e Cookie) Validate(key core.Credentials) error {
	var err error

	// Проверяем существование ключа или фактора в базе,
	// если не нашли - возвращаем ошибку

	return err
}

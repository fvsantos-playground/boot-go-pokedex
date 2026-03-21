package pokeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func getData[T any](val []byte) (T, error) {
	var data T
	decoder := json.NewDecoder(bytes.NewBuffer(val))
	if err := decoder.Decode(&data); err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

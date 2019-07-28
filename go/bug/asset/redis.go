package asset

import (
	"itflow/bug/bugconfig"
	"errors"
	"itflow/gadb"
)

func Delkey(key string) error {
	r := gadb.NewRedis()
	conn, err := r.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	if ok, _ := conn.Get(key).Result(); ok == "ep" {
		conn.Del(key)
	}
	return nil
}

func Settimeout(key string) error {
	r := gadb.NewRedis()
	conn, err := r.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	if ok, err := conn.Expire(key, bugconfig.DEADLINE).Result(); !ok {
		return err
	}
	return nil
}

func Getvalue(key string) (string, error) {

	conn, err := gadb.NewRedis().Connect()
	if err != nil {
		return "",err
	}
	defer conn.Close()


	value, err := conn.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
	

	return "", errors.New("key not found")
}

func Setkey(key string, value string) error {
	r := gadb.NewRedis()
	conn, err := r.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Set(key, value, bugconfig.DEADLINE).Result()
	if err != nil {
		return err
	}
	return  nil
}

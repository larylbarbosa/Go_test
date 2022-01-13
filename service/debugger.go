package service

import "log"

func fail(functionName string, err error) error {
	log.Printf("\n %v: %v \n", functionName, err.Error())

	return err
}

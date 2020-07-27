package neffos_test

import (
	"context"
	"fmt"

	neffos "github.com/telexy324/simple-neffos"

	gobwas "github.com/telexy324/simple-neffos/gobwas"
	gorilla "github.com/telexy324/simple-neffos/gorilla"
)

func runTestClient(addr string, connHandler neffos.ConnHandler, testFn func(string, *neffos.Client)) func() error {
	gobwasClient, err := neffos.Dial(context.TODO(), gobwas.DefaultDialer, fmt.Sprintf("ws://%s/gobwas", addr), connHandler)
	if err != nil {
		return func() error {
			return err
		}
	}
	gorillaClient, err := neffos.Dial(context.TODO(), gorilla.DefaultDialer, fmt.Sprintf("ws://%s/gorilla", addr), connHandler)
	if err != nil {
		return func() error {
			return err
		}
	}

	// teardown.
	teardown := func() error {
		gobwasClient.Close()
		gorillaClient.Close()
		return nil
	}

	testFn("gobwas", gobwasClient)
	testFn("gorilla", gorillaClient)
	return teardown
}

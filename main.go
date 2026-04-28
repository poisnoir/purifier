package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/poisnoir/purifier/subscribers"
	"github.com/poisnoir/spine-go"
)

func main() {
	namespace := flag.String("namespace", "rime", "spine namespace to join")
	key := flag.String("key", "ppap", "spine namespace key")
	name := flag.String("name", "xbox-controller", "publisher name")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ns, err := spine.JointNamespace(*namespace, *key, logger)
	if err != nil {
		panic(err)
	}

	_, _ := spine.NewSubscriber(ns, *name, subscribers.ControllerHandler)

	select {}

}

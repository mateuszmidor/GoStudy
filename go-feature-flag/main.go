package main

import (
	"fmt"
	"time"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffcontext"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

func main() {
	// prepare the feature flags client
	err := ffclient.Init(ffclient.Config{
		PollingInterval: 3 * time.Second, // re-load flag config every 3 sec
		Retriever: &fileretriever.Retriever{ // we can also use e.g. "http retriever" to load flags from remote server
			Path: "flags.yaml",
		},
	})
	if err != nil {
		panic(err)
	}
	defer ffclient.Close()

	// EvaluationContext is needed for stable flag evaluation e.g. in canary-testing; should uniquely identify the user, eg UserID
	user := ffcontext.NewEvaluationContext("userid-42")
	user.AddCustomAttribute("email", "john.doe@acme.com")

	// print the flag in a loop
	for {
		hasFlag, _ := ffclient.BoolVariation("MyFeatureFlag", user, false)
		fmt.Printf("Flag is %v\n", hasFlag)
		time.Sleep(3 * time.Second)
	}
}

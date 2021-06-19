package cron

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

// Start - start based on local or lambda
func Start(isLambda bool) {
	if isLambda {
		// if lambda, set a cloudwatch cron
		lambda.Start(parseVideos)
	} else {
		// if local, run for each 10 sec
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Running at", t)
				parseVideos()
			}
		}
	}
}

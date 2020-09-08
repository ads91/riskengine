package logging

import (
	"context"
	"log"
	"os"
	"riskengine/config"

	"cloud.google.com/go/logging"
)

// GetLogger : return a logger according to conditions
func GetLogger() *log.Logger {
	// declare logger
	var logger *log.Logger
	// determine logger to return
	if config.DEPLOYMENT_TYPE == 1 {
		logger = getGCPLogger(config.GCP_PROJECT_ID, config.LOG_NAME)
	} else {
		logger = getLocalLogger()
	}
	return logger
}

func getGCPLogger(projectID string, logName string) *log.Logger {
	ctx := context.Background()
	// creates a client
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	// sets the name of the log to write to
	logger := client.Logger(logName).StandardLogger(logging.Info)
	// logs "hello world", log entry is visible at stackdriver Logs
	logger.Println("hello world")
	return logger
}

func getLocalLogger() *log.Logger {
	// open a file
	file, err := os.OpenFile(os.Getenv(config.WORKING_DIR)+config.LOG_NAME, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// create + return the logger
	return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

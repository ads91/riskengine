// should really change to a JSON
package config

// DEPLOYMENT_TYPE : deployment type, 0 for local, 1 for GCP
var DEPLOYMENT_TYPE = 0

// WORKING_DIR: the working directory of the main entry point
var WORKING_DIR = "RISK_ENGINE_DIR"

// DEFAULT_PORT : default port to listen on
var DEFAULT_PORT = ":8080"

// LOG_NAME : name of cloud platform log file on GCP
var LOG_NAME = "riskengine-log.txt"

package main

import (
	"lesson_6/cmd"
	"lesson_6/cmd/keys"
	"lesson_6/cmd/signatures"
	"lesson_6/logger"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(err, "Initial setup failed")
	}
}

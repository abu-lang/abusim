package main

import (
	"log"

	"github.com/abu-lang/abusim/abusim/args"
	"github.com/abu-lang/abusim/abusim/command"
	"github.com/abu-lang/abusim/abusim/config"
	"github.com/abu-lang/abusim/abusim/docker"
)

func main() {
	// I get the arguments from the command line...
	argsConfig := args.ParseArgs()
	// ... I parse the configuration file...
	conf, err := config.Parse(argsConfig.ConfigFile)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
	// ... I connect to the docker daemon...
	dcli, err := docker.New()
	if err != nil {
		log.Fatalf("Error creating docker client: %v", err)
	}
	// ... and I execute the correct subcommand
	switch argsConfig.SubCommand {
	case args.SUBCOMMAND_UP:
		command.Up(argsConfig, conf, dcli)
	case args.SUBCOMMAND_DOWN:
		command.Down(argsConfig, conf, dcli)
	case args.SUBCOMMAND_LOGS:
		if argsConfig.FollowLogs {
			command.LogsFollow(conf, dcli)
		} else {
			command.Logs(conf, dcli)
		}
	}
}

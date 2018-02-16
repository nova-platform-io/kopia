package cli

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("kopia", "Kopia - Online Backup").Author("http://kopia.github.io/")

	_ = app.Flag("help-full", "Show help for all commands, including hidden").Action(helpFullAction).Bool()

	repositoryCommands = app.Command("repository", "Commands to manipulate repository.").Alias("repo")
	snapshotCommands   = app.Command("snapshot", "Commands to manipulate snapshots.").Alias("snap")
	policyCommands     = app.Command("policy", "Commands to manipulate snapshotting policies.").Alias("policies")
	manifestCommands   = app.Command("manifest", "Low-level commands to manipulate manifest items.").Hidden()
	objectCommands     = app.Command("object", "Commands to manipulate objects in repository.").Alias("obj").Hidden()
	blockCommands      = app.Command("block", "Commands to manipulate virtual blocks in repository.").Alias("blk").Hidden()
	storageCommands    = app.Command("storage", "Commands to manipulate raw storage blocks.").Hidden()
	blockIndexCommands = app.Command("blockindex", "Commands to manipulate block index.").Hidden()
)

func helpFullAction(ctx *kingpin.ParseContext) error {
	app.UsageForContextWithTemplate(ctx, 0, kingpin.DefaultUsageTemplate)
	os.Exit(0)
	return nil
}

// App returns an instance of command-line application object.
func App() *kingpin.Application {
	return app
}

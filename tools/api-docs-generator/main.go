package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path"
	"time"

	"github.com/snyk/user-docs/tools/api-docs-generator/changelog"

	"github.com/snyk/user-docs/tools/api-docs-generator/config"
	"github.com/snyk/user-docs/tools/api-docs-generator/fetcher"
	"github.com/snyk/user-docs/tools/api-docs-generator/generator"
)

var dryRunFlag = flag.Bool("dry-run", false,
	`runs the generation without updating the changelog or fetching the latest spec, `+
		`useful for testing the validity of the current specifications`)

func main() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()
	if len(os.Args) < 3 {
		log.Panicf("usage: api-docs <config-file> <docs-dir>")
	}
	flag.Parse()

	cfg, err := config.Parse(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	docsDirectory := os.Args[2]

	if !*dryRunFlag {
		var syncStateCfg config.SyncStateConfig
		syncStateCfg, err = config.LoadSyncState(path.Join(docsDirectory, cfg.Changelog.SyncStateFile))
		if err != nil {
			log.Panic(err)
		}

		err = changelog.UpdateChangelog(ctx, cfg, syncStateCfg, docsDirectory)
		if err != nil {
			log.Panic(err)
		}

		// replace latest spec
		err = fetcher.FetchSpec(ctx, cfg, docsDirectory)
		if err != nil {
			log.Panic(err)
		}
	}
	err = generator.GenerateReferenceDocs(cfg, docsDirectory)
	if err != nil {
		log.Panic(err)
	}
}

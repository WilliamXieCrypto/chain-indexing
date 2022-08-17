package main

import (
	"fmt"
	"strings"

	"github.com/WilliamXieCrypto/chain-indexing/appinterface/rdb"
	configuration "github.com/WilliamXieCrypto/chain-indexing/bootstrap/config"
	projection_entity "github.com/WilliamXieCrypto/chain-indexing/entity/projection"
	applogger "github.com/WilliamXieCrypto/chain-indexing/external/logger"
	"github.com/WilliamXieCrypto/chain-indexing/infrastructure/pg"
	"github.com/WilliamXieCrypto/chain-indexing/infrastructure/pg/migrationhelper"
	github_migrationhelper "github.com/WilliamXieCrypto/chain-indexing/infrastructure/pg/migrationhelper/github"
	"github.com/WilliamXieCrypto/chain-indexing/projection/bridge_activity/bridge_activity_matcher"
)

func initCronJobs(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	config *configuration.Config,
) []projection_entity.CronJob {
	cronJobs := make([]projection_entity.CronJob, 0, len(config.IndexService.CronJob.Enables))
	initParams := InitCronJobParams{
		Logger:  logger,
		RdbConn: rdbConn,

		ExtraConfigs: config.IndexService.CronJob.ExtraConfigs,
	}

	for _, cronJobName := range config.IndexService.CronJob.Enables {
		cronJob := InitCronJob(
			cronJobName, initParams,
		)
		if onInitErr := cronJob.OnInit(); onInitErr != nil {
			panic(fmt.Errorf(
				"error initializing cron job %s: %v",
				cronJob.Id(), onInitErr,
			))
		}
		cronJobs = append(cronJobs, cronJob)
	}

	logger.Infof("Enabled the following cron jobs: [%s]", strings.Join(config.IndexService.CronJob.Enables, ", "))

	return cronJobs
}

func InitCronJob(name string, params InitCronJobParams) projection_entity.CronJob {
	connString := params.RdbConn.(*pg.PgxConn).ConnString()

	switch name {
	case "BridgeActivityMatcher":
		sourceURL := github_migrationhelper.GenerateSourceURL(
			github_migrationhelper.MIGRATION_GITHUB_URL_FORMAT,
			params.GithubAPIUser,
			params.GithubAPIToken,
			bridge_activity_matcher.MIGRATION_DIRECOTRY,
			params.MigrationRepoRef,
		)
		databaseURL := migrationhelper.GenerateDefaultDatabaseURL(name, connString)
		migrationHelper := github_migrationhelper.NewGithubMigrationHelper(sourceURL, databaseURL)

		config, err := bridge_activity_matcher.ConfigFromInterface(params.ExtraConfigs[name])
		if err != nil {
			params.Logger.Panicf(err.Error())
		}

		return bridge_activity_matcher.New(config, params.Logger, params.RdbConn, migrationHelper)
	// register more cronjobs here
	default:
		panic(fmt.Sprintf("Unrecognized cron job: %s", name))
	}
}

type InitCronJobParams struct {
	Logger  applogger.Logger
	RdbConn rdb.Conn

	ExtraConfigs map[string]interface{}

	GithubAPIUser    string
	GithubAPIToken   string
	MigrationRepoRef string
}

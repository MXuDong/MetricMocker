package instances

import "github.com/robfig/cron/v3"

var GlobalCron *cron.Cron

func init() {
	GlobalCron = cron.New()
	GlobalCron.Start()
}

package ioc

import (
	"context"
	"database/sql"
	"github.com/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
	"loverrecipe/internal/pkg/database/metrics"
	"loverrecipe/internal/pkg/database/tracing"
	"loverrecipe/internal/repository/dao"
	"time"

	"github.com/gotomicro/ego/core/econf"

	"github.com/ecodeclub/ekit/retry"
)

func InitDB() *egorm.Component {
	WaitForDBSetup(econf.GetString("mysql.dsn"))
	auto := econf.GetBool("mysql.auto_migrate")

	db := egorm.Load("mysql").Build()
	if auto {
		err := dao.InitTables(db)
		if err != nil {
			panic(err)
		}
	}

	// 这个是自己手搓的
	tracePlugin := tracing.NewGormTracingPlugin()
	metricsPlugin := metrics.NewGormMetricsPlugin()
	err := db.Use(tracePlugin)
	if err != nil {
		panic(err)
	}
	err = db.Use(metricsPlugin)
	if err != nil {
		panic(err)
	}
	elog.Info("database init success")
	return db
}

func WaitForDBSetup(dsn string) {
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	const maxInterval = 10 * time.Second
	const maxRetries = 10
	strategy, err := retry.NewExponentialBackoffRetryStrategy(time.Second, maxInterval, maxRetries)
	if err != nil {
		panic(err)
	}

	const timeout = 5 * time.Second
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		err = sqlDB.PingContext(ctx)
		cancel()
		if err == nil {
			break
		}
		elog.Warn("WaitForDBSetup 数据库连接失败，重试中...", elog.FieldErr(err))
		next, ok := strategy.Next()
		if !ok {
			panic("WaitForDBSetup 重试失败......")
		}
		time.Sleep(next)
	}
}

package rk_hmes

import (
	"fmt"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"go.uber.org/zap"
	"time"
)

func init() {
	rkentry.RegisterPluginRegFunc(RegisterRedisEntryYAML)
}

func RegisterRedisEntryYAML(raw []byte) map[string]rkentry.Entry {
	res := make(map[string]rkentry.Entry)

	// 1: unmarshal user provided config into boot config struct
	config := &BootRedis{}
	rkentry.UnmarshalBootYAML(raw, config)

	// filter out based domain
	configMap := make(map[string]*BootRedisE)
	for _, e := range config.Redis {
		if !e.Enabled || len(e.Name) < 1 {
			continue
		}

		if !rkentry.IsValidDomain(e.Domain) {
			continue
		}

		// * or matching domain
		// 1: add it to map if missing
		if _, ok := configMap[e.Name]; !ok {
			configMap[e.Name] = e
			continue
		}

		// 2: already has an entry, then compare domain,
		//    only one case would occur, previous one is already the correct one, continue
		if e.Domain == "" || e.Domain == "*" {
			continue
		}

		configMap[e.Name] = e
	}

	for _, element := range configMap {
		universalOpt := &redis.UniversalOptions{
			Addrs:                 element.Addrs,
			DB:                    element.DB,
			Username:              element.User,
			Password:              element.Pass,
			SentinelPassword:      element.SentinelPass,
			MaxRetries:            element.MaxRetries,
			MinRetryBackoff:       time.Duration(element.MinRetryBackoffMs) * time.Millisecond,
			MaxRetryBackoff:       time.Duration(element.MaxRetryBackoffMs) * time.Millisecond,
			DialTimeout:           time.Duration(element.DialTimeoutMs) * time.Millisecond,
			ReadTimeout:           time.Duration(element.ReadTimeoutMs) * time.Millisecond,
			WriteTimeout:          time.Duration(element.WriteTimeoutMs) * time.Millisecond,
			ContextTimeoutEnabled: element.ContextTimeoutEnabled,

			PoolFIFO:     element.PoolFIFO,
			PoolSize:     element.PoolSize,
			PoolTimeout:  time.Duration(element.PoolTimeoutMs) * time.Millisecond,
			MinIdleConns: element.MinIdleConn,
			MaxIdleConns: element.MaxIdleConn,

			ConnMaxIdleTime: time.Duration(element.ConnMaxIdleTimeMs) * time.Millisecond,
			ConnMaxLifetime: time.Duration(element.ConnMaxLifetimeMs) * time.Millisecond,

			MaxRedirects:   element.MaxRedirects,
			ReadOnly:       element.ReadOnly,
			RouteByLatency: element.RouteByLatency,
			RouteRandomly:  element.RouteRandomly,
			MasterName:     element.MasterName,
		}

		certEntry := rkentry.GlobalAppCtx.GetCertEntry(element.CertEntry)

		entry := RegisterRedisEntry(
			WithName(element.Name),
			WithDescription(element.Description),
			WithUniversalOption(universalOpt),
			WithCertEntry(certEntry),
			WithLoggerEntry(rkentry.GlobalAppCtx.GetLoggerEntry(element.LoggerEntry)))

		res[entry.GetName()] = entry
	}

	return res
}

func RegisterMySqlEntry(opts ...Option) *MySqlEntry {

}

func (entry *MySqlEntry) Bootstrap(ctx context.Context) {
	// extract eventId if exists
	fields := make([]zap.Field, 0)

	if val := ctx.Value("eventId"); val != nil {
		if id, ok := val.(string); ok {
			fields = append(fields, zap.String("eventId", id))
		}
	}

	fields = append(fields,
		zap.String("entryName", entry.entryName),
		zap.String("entryType", entry.entryType))

	entry.logger.delegate.Info("Bootstrap MySqlEntry", fields...)

	// Connect and create db if missing
	if err := entry.connect(); err != nil {
		fields = append(fields, zap.Error(err))
		entry.logger.delegate.Error("Failed to connect to database", fields...)
		rkentry.ShutdownWithError(fmt.Errorf("failed to connect to database at %s:%s@%s(%s)",
			entry.User, "****", entry.Protocol, entry.Addr))
	}
}

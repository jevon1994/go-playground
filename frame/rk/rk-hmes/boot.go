package rkhmes

import (
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

func init() {
	rkentry.RegisterPluginRegFunc(RegisterHMESEntryYAML)
}

type HMESEntry struct {
}

func RegisterHMESEntryYAML(raw []byte) map[string]rkentry.Entry {
	res := make(map[string]rkentry.Entry)

	// 1: unmarshal user provided config into boot config struct
	//config := &BootRedis{}
	//rkentry.UnmarshalBootYAML(raw, config)
	//
	//// filter out based domain
	//configMap := make(map[string]*BootRedisE)
	//for _, e := range config.Redis {
	//	if !e.Enabled || len(e.Name) < 1 {
	//		continue
	//	}
	//
	//	if !rkentry.IsValidDomain(e.Domain) {
	//		continue
	//	}
	//
	//	// * or matching domain
	//	// 1: add it to map if missing
	//	if _, ok := configMap[e.Name]; !ok {
	//		configMap[e.Name] = e
	//		continue
	//	}
	//
	//	// 2: already has an entry, then compare domain,
	//	//    only one case would occur, previous one is already the correct one, continue
	//	if e.Domain == "" || e.Domain == "*" {
	//		continue
	//	}
	//
	//	configMap[e.Name] = e
	//}
	//
	//for _, element := range configMap {
	//	universalOpt := &redis.UniversalOptions{
	//		Addrs:                 element.Addrs,
	//		DB:                    element.DB,
	//		Username:              element.User,
	//		Password:              element.Pass,
	//		SentinelPassword:      element.SentinelPass,
	//		MaxRetries:            element.MaxRetries,
	//		MinRetryBackoff:       time.Duration(element.MinRetryBackoffMs) * time.Millisecond,
	//		MaxRetryBackoff:       time.Duration(element.MaxRetryBackoffMs) * time.Millisecond,
	//		DialTimeout:           time.Duration(element.DialTimeoutMs) * time.Millisecond,
	//		ReadTimeout:           time.Duration(element.ReadTimeoutMs) * time.Millisecond,
	//		WriteTimeout:          time.Duration(element.WriteTimeoutMs) * time.Millisecond,
	//		ContextTimeoutEnabled: element.ContextTimeoutEnabled,
	//
	//		PoolFIFO:     element.PoolFIFO,
	//		PoolSize:     element.PoolSize,
	//		PoolTimeout:  time.Duration(element.PoolTimeoutMs) * time.Millisecond,
	//		MinIdleConns: element.MinIdleConn,
	//		MaxIdleConns: element.MaxIdleConn,
	//
	//		ConnMaxIdleTime: time.Duration(element.ConnMaxIdleTimeMs) * time.Millisecond,
	//		ConnMaxLifetime: time.Duration(element.ConnMaxLifetimeMs) * time.Millisecond,
	//
	//		MaxRedirects:   element.MaxRedirects,
	//		ReadOnly:       element.ReadOnly,
	//		RouteByLatency: element.RouteByLatency,
	//		RouteRandomly:  element.RouteRandomly,
	//		MasterName:     element.MasterName,
	//	}
	//
	//	certEntry := rkentry.GlobalAppCtx.GetCertEntry(element.CertEntry)
	//
	//	entry := RegisterRedisEntry(
	//		WithName(element.Name),
	//		WithDescription(element.Description),
	//		WithUniversalOption(universalOpt),
	//		WithCertEntry(certEntry),
	//		WithLoggerEntry(rkentry.GlobalAppCtx.GetLoggerEntry(element.LoggerEntry)))
	//
	//	res[entry.GetName()] = entry
	//}

	return res
}

func RegisterMySqlEntry(opts ...Option) *HMESEntry {
	return &HMESEntry{}
}

// Option for MySqlEntry
type Option func(*HMESEntry)

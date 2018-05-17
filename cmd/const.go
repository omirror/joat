package cmd

import (
	"strings"
)

const (
	configEnvPrefix       = "JOAT"
	defaultBindAddr       = "0.0.0.0"
	defaultClusterPort    = 7080
	defaultDataDir        = "~/.joat"
	defaultHttpPort       = 8080
	defaultRpcPort        = 9080
	flagAdvertiseAddr     = "advertise"
	flagAdvertiseAddrDesc = "address advertised to other nodes in the cluster (defaults to first available private IP)"
	flagBindAddr          = "bind"
	flagBindAddrDesc      = "address used to provide network services"
	flagClusterPort       = "cluster-port"
	flagClusterPortDesc   = "cluster management port"
	flagDataDir           = "data-dir"
	flagDataDirDesc       = "data directory"
	flagDebug             = "debug"
	flagDebugShort        = "d"
	flagDebugDesc         = "enable debug logging"
	flagJoin              = "join"
	flagJoinDesc          = "join a cluster"
	flagHttpBindAddr      = "http-bind"
	flagHttpBindAddrDesc  = "http server bind address"
	flagHttpPort          = "http-port"
	flagHttpPortDesc      = "http server port"
	flagRpcPort           = "rpc-port"
	flagRpcPortDesc       = "rpc server port"
	flagVerbose           = "verbose"
	flagVerboseShort      = "v"
	flagVerboseDesc       = "enable verbose logging"
)

var envKeyReplacer = strings.NewReplacer("-", "_")

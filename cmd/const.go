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
	flagAdvertiseAddrDesc = "Address advertised to other nodes in the cluster (defaults to first available private IP)"
	flagBindAddr          = "bind"
	flagBindAddrDesc      = "Address used to provide network services"
	flagClusterPort       = "cluster-port"
	flagClusterPortDesc   = "Cluster management port"
	flagDataDir           = "data-dir"
	flagDataDirDesc       = "Data directory"
	flagDebug             = "debug"
	flagDebugShort        = "d"
	flagDebugDesc         = "Enable debug logging"
	flagJoin              = "join"
	flagJoinDesc          = "Join a cluster"
	flagHttpBindAddr      = "http-bind"
	flagHttpBindAddrDesc  = "HTTP server bind address"
	flagHttpPort          = "http-port"
	flagHttpPortDesc      = "http server port"
	flagRpcPort           = "rpc-port"
	flagRpcPortDesc       = "RPC server port"
	flagVerbose           = "verbose"
	flagVerboseShort      = "v"
	flagVerboseDesc       = "Enable verbose logging"
)

var envKeyReplacer = strings.NewReplacer("-", "_")

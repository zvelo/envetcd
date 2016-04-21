package zenv

import (
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
)

// ZEnv is the runtime environment
type ZEnv int

const (
	// Development environment
	Development ZEnv = iota
	// Test environment
	Test
	// Integration environment
	Integration
	// Production environment
	Production
)

// Flag is the cli option
var Flags = []cli.Flag{
	cli.StringFlag{
		Name:   "zvelo-env, e",
		EnvVar: "ZVELO_ENV",
		Usage:  "runtime environment ['development', 'test', 'integration', 'production']",
		Value:  "development",
	},
	cli.StringFlag{
		Name:   "cluster-id",
		EnvVar: "CLUSTER_ID",
		Usage:  "denotes the cluster id of the application environment",
	},
}

// Init parses env and returns the corresponding ZEnv
func Init(c *cli.Context) ZEnv {
	var zenv ZEnv
	env := c.String("zvelo-env")
	if env == "" {
		env = c.GlobalString("zvelo-env")
	}

	switch strings.ToUpper(env) {
	case "TEST":
		zenv = Test
	case "INTEGRATION":
		zenv = Integration
	case "PRODUCTION":
		zenv = Production
	default:
		zenv = Development
	}

	id := c.String("cluster-id")
	if len(id) == 0 {
		if zenv != Development {
			panic("CLUSTER_IDis not set. Invalid application environment configuration")
		}
		clusterID = DevCluster
		return zenv
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		panic("Error parsing cluster id from environment variable")
	}

	if i >= len(_ClusterID_index) || i < 0 {
		panic("Invalid Cluster ID")
	}
	clusterID = ClusterID(i)
	return zenv
}

//go:generate stringer -type=ZEnv

package zenv

// ClusterID is a unique identifier for each cluster. It should be set in the
// CLUSTER_ID env variable
type ClusterID byte

const (
	// DevCluster is the ID that should be used in development
	DevCluster ClusterID = iota
	// TestCluster is the ID that should be used by testing and continuous integration
	// environments
	TestCluster
	// IntegrationCluster is the ID that should be used in non-production integration
	// environments
	IntegrationCluster
	// StagingCluster is the ID that should be used in pre-production environments
	StagingCluster
	// AwsUSWest2_1 is the ID for Amazon Web Services us-west-2 cluster #1
	AwsUSWest2_1
	// AwsUSEast1_1 is the ID for Amazon Web Services us-east-1 cluster #1
	AwsUSEast1_1
)

var clusterID ClusterID

// ID returns the cluster id derived from $CLUSTER_ID. If $CLUSTER_ID is not
// set, it returns Unset.
func ID() ClusterID {
	return clusterID
}

//go:generate stringer -type=ClusterID

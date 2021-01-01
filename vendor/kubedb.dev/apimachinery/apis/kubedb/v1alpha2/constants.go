/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha2

import (
	"kubedb.dev/apimachinery/apis/kubedb"

	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	// Deprecated
	DatabaseNamePrefix = "kubedb"

	KubeDBOrganization = "kubedb"

	LabelRole = kubedb.GroupName + "/role"

	ReplicationModeDetectorContainerName = "replication-mode-detector"
	DatabasePodPrimary                   = "primary"
	DatabasePodStandby                   = "standby"

	ComponentDatabase     = "database"
	RoleStats             = "stats"
	DefaultStatsPath      = "/metrics"
	DefaultPasswordLength = 16

	ContainerExporterName = "exporter"
	LocalHost             = "localhost"
	LocalHostIP           = "127.0.0.1"

	DBCustomConfigName             = "custom-config"
	DefaultVolumeClaimTemplateName = "data"

	// =========================== Database key Constants ============================
	PostgresKey      = ResourceSingularPostgres + "." + kubedb.GroupName
	ElasticsearchKey = ResourceSingularElasticsearch + "." + kubedb.GroupName
	MySQLKey         = ResourceSingularMySQL + "." + kubedb.GroupName
	MariaDBKey       = ResourceSingularMariaDB + "." + kubedb.GroupName
	PerconaXtraDBKey = ResourceSingularPerconaXtraDB + "." + kubedb.GroupName
	MongoDBKey       = ResourceSingularMongoDB + "." + kubedb.GroupName
	RedisKey         = ResourceSingularRedis + "." + kubedb.GroupName
	MemcachedKey     = ResourceSingularMemcached + "." + kubedb.GroupName
	EtcdKey          = ResourceSingularEtcd + "." + kubedb.GroupName
	ProxySQLKey      = ResourceSingularProxySQL + "." + kubedb.GroupName

	// =========================== Elasticsearch Constants ============================
	ElasticsearchRestPort                        = 9200
	ElasticsearchRestPortName                    = "http"
	ElasticsearchTransportPort                   = 9300
	ElasticsearchTransportPortName               = "transport"
	ElasticsearchMetricsPort                     = 9600
	ElasticsearchIngestNodePrefix                = "ingest"
	ElasticsearchDataNodePrefix                  = "data"
	ElasticsearchMasterNodePrefix                = "master"
	ElasticsearchNodeRoleMaster                  = "node.role.master"
	ElasticsearchNodeRoleIngest                  = "node.role.ingest"
	ElasticsearchNodeRoleData                    = "node.role.data"
	ElasticsearchNodeRoleSet                     = "set"
	ElasticsearchConfigDir                       = "/usr/share/elasticsearch/config"
	ElasticsearchTempConfigDir                   = "/elasticsearch/temp-config"
	ElasticsearchCustomConfigDir                 = "/elasticsearch/custom-config"
	ElasticsearchDataDir                         = "/usr/share/elasticsearch/data"
	ElasticsearchOpendistroSecurityConfigDir     = "/usr/share/elasticsearch/plugins/opendistro_security/securityconfig"
	ElasticsearchSearchGuardSecurityConfigDir    = "/usr/share/elasticsearch/plugins/search-guard-%v/sgconfig"
	ElasticsearchOpendistroReadallMonitorRole    = "readall_and_monitor"
	ElasticsearchSearchGuardReadallMonitorRoleV7 = "SGS_READALL_AND_MONITOR"
	ElasticsearchSearchGuardReadallMonitorRoleV6 = "sg_readall_and_monitor"
	ElasticsearchStatusGreen                     = "green"
	ElasticsearchStatusYellow                    = "yellow"
	ElasticsearchStatusRed                       = "red"
	ElasticsearchInitSysctlContainerName         = "init-sysctl"
	ElasticsearchInitConfigMergerContainerName   = "config-merger"
	ElasticsearchContainerName                   = "elasticsearch"
	ElasticsearchExporterContainerName           = "exporter"

	// Ref:
	//	- https://www.elastic.co/guide/en/elasticsearch/reference/7.6/heap-size.html#heap-size
	//	- no more than 50% of your physical RAM
	//	- no more than 32GB that the JVM uses for compressed object pointers (compressed oops)
	//	- no more than 26GB for zero-based compressed oops;
	// 26 GB is safe on most systems
	ElasticsearchMaxHeapSize = 26 * 1024 * 1024 * 1024
	// 128MB
	ElasticsearchMinHeapSize = 128 * 1024 * 1024

	// =========================== Memcached Constants ============================
	MemcachedDatabasePortName       = "db"
	MemcachedPrimaryServicePortName = "primary"
	MemcachedDatabasePort           = 11211

	// =========================== MongoDB Constants ============================

	MongoDBDatabasePortName       = "db"
	MongoDBPrimaryServicePortName = "primary"
	MongoDBDatabasePort           = 27017
	MongoDBKeyFileSecretSuffix    = "key"
	MongoDBRootUsername           = "root"
	MongoDBCustomConfigFile       = "mongod.conf"

	// =========================== MySQL Constants ============================
	MySQLMetricsExporterConfigSecretSuffix = "metrics-exporter-config"
	MySQLDatabasePortName                  = "db"
	MySQLPrimaryServicePortName            = "primary"
	MySQLStandbyServicePortName            = "standby"
	MySQLDatabasePort                      = 3306
	MySQLGroupComPort                      = 33060
	MySQLMaxGroupMembers                   = 9
	// The recommended MySQL server version for group replication (GR)
	MySQLGRRecommendedVersion       = "5.7.25"
	MySQLDefaultGroupSize           = 3
	MySQLDefaultBaseServerID  int64 = 1
	// The server id for each group member must be unique and in the range [1, 2^32 - 1]
	// And the maximum group size is 9. So MySQLMaxBaseServerID is the maximum safe value
	// for BaseServerID calculated as max MySQL server_id value - max Replication Group size.
	// xref: https://dev.mysql.com/doc/refman/5.7/en/replication-options.html
	MySQLMaxBaseServerID int64 = ((1 << 32) - 1) - 9
	MySQLRootUserName          = "MYSQL_ROOT_USERNAME"
	MySQLRootPassword          = "MYSQL_ROOT_PASSWORD"
	MySQLName                  = "MYSQL_NAME"

	MySQLTLSConfigCustom     = "custom"
	MySQLTLSConfigSkipVerify = "skip-verify"
	MySQLTLSConfigTrue       = "true"
	MySQLTLSConfigFalse      = "false"
	MySQLTLSConfigPreferred  = "preferred"

	// =========================== PerconaXtraDB Constants ============================
	PerconaXtraDBClusterRecommendedVersion    = "5.7"
	PerconaXtraDBMaxClusterNameLength         = 32
	PerconaXtraDBStandaloneReplicas           = 1
	PerconaXtraDBDefaultClusterSize           = 3
	PerconaXtraDBDataMountPath                = "/var/lib/mysql"
	PerconaXtraDBDataLostFoundPath            = PerconaXtraDBDataMountPath + "lost+found"
	PerconaXtraDBInitDBMountPath              = "/docker-entrypoint-initdb.d"
	PerconaXtraDBCustomConfigMountPath        = "/etc/percona-server.conf.d/"
	PerconaXtraDBClusterCustomConfigMountPath = "/etc/percona-xtradb-cluster.conf.d/"

	// =========================== MariaDB Constants ============================
	MariaDBClusterRecommendedVersion    = "5.7"
	MariaDBMaxClusterNameLength         = 32
	MariaDBStandaloneReplicas           = 1
	MariaDBDefaultClusterSize           = 3
	MariaDBDataMountPath                = "/var/lib/mysql"
	MariaDBDataLostFoundPath            = MariaDBDataMountPath + "lost+found"
	MariaDBInitDBMountPath              = "/docker-entrypoint-initdb.d"
	MariaDBCustomConfigMountPath        = "/etc/percona-server.conf.d/"
	MariaDBClusterCustomConfigMountPath = "/etc/percona-xtradb-cluster.conf.d/"

	// =========================== PostgreSQL Constants ============================
	PostgresDatabasePortName       = "db"
	PostgresPrimaryServicePortName = "primary"
	PostgresStandbyServicePortName = "standby"
	PostgresDatabasePort           = 5432
	PostgresPodPrimary             = "primary"
	PostgresPodStandby             = "standby"
	PostgresLabelRole              = kubedb.GroupName + "/role"

	// =========================== ProxySQL Constants ============================
	LabelProxySQLName        = ProxySQLKey + "/name"
	LabelProxySQLLoadBalance = ProxySQLKey + "/load-balance"

	ProxySQLDatabasePort           = 6033
	ProxySQLDatabasePortName       = "db"
	ProxySQLPrimaryServicePortName = "db"
	ProxySQLAdminPort              = 6032
	ProxySQLAdminPortName          = "admin"
	ProxySQLDataMountPath          = "/var/lib/proxysql"
	ProxySQLCustomConfigMountPath  = "/etc/custom-config"

	// =========================== Redis Constants ============================
	RedisShardKey               = RedisKey + "/shard"
	RedisDatabasePortName       = "db"
	RedisPrimaryServicePortName = "primary"
	RedisDatabasePort           = 6379
	RedisGossipPortName         = "gossip"
	RedisGossipPort             = 16379

	RedisKeyFileSecretSuffix = "key"
	RedisPEMSecretSuffix     = "pem"
	RedisRootUsername        = "root"

	// =========================== PgBouncer Constants ============================
	PgBouncerUpstreamServerCA       = "upstream-server-ca.crt"
	PgBouncerDatabasePortName       = "db"
	PgBouncerPrimaryServicePortName = "primary"
	PgBouncerDatabasePort           = 5432
	PgBouncerConfigFile             = "pgbouncer.ini"
	PgBouncerAdminUsername          = "kubedb"
)

// List of possible condition types for a KubeDB object
const (
	// used for Databases that have started provisioning
	DatabaseProvisioningStarted = "ProvisioningStarted"
	// used for Databases which completed provisioning
	DatabaseProvisioned = "Provisioned"
	// used for Databases that are currently being initialized using stash
	DatabaseDataRestoreStarted = "DataRestoreStarted"
	// used for Databases that have been initialized using stash
	DatabaseDataRestored = "DataRestored"
	// used for Databases whose pods are ready
	DatabaseReplicaReady = "ReplicaReady"
	// used for Databases that are currently accepting connection
	DatabaseAcceptingConnection = "AcceptingConnection"
	// used for Databases that report status OK (also implies that we can connect to it)
	DatabaseReady = "Ready"
	// used for Databases that are paused
	DatabasePaused = "Paused"
	// used for Databases that are halted
	DatabaseHalted = "Halted"

	// Condition reasons
	DataRestoreStartedByExternalInitializer = "DataRestoreStartedByExternalInitializer"
	DatabaseSuccessfullyRestored            = "SuccessfullyDataRestored"
	FailedToRestoreData                     = "FailedToRestoreData"
	AllReplicasAreReady                     = "AllReplicasReady"
	SomeReplicasAreNotReady                 = "SomeReplicasNotReady"
	DatabaseAcceptingConnectionRequest      = "DatabaseAcceptingConnectionRequest"
	DatabaseNotAcceptingConnectionRequest   = "DatabaseNotAcceptingConnectionRequest"
	ReadinessCheckSucceeded                 = "ReadinessCheckSucceeded"
	ReadinessCheckFailed                    = "ReadinessCheckFailed"
	DatabaseProvisioningStartedSuccessfully = "DatabaseProvisioningStartedSuccessfully"
	DatabaseSuccessfullyProvisioned         = "DatabaseSuccessfullyProvisioned"
	DatabaseHaltedSuccessfully              = "DatabaseHaltedSuccessfully"
)

// Resource kind related constants
const (
	ResourceKindStatefulSet = "StatefulSet"
)

var (
	defaultResourceLimits = core.ResourceList{
		core.ResourceCPU:    resource.MustParse(".500"),
		core.ResourceMemory: resource.MustParse("1024Mi"),
	}
)

---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dbtcloud_global_connection Data Source - dbtcloud"
subcategory: ""
description: |-
  
---

# dbtcloud_global_connection (Data Source)



## Example Usage

```terraform
data dbtcloud_global_connection my_connection {
  id = 1234
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (Number) Connection Identifier

### Read-Only

- `adapter_version` (String) Version of the adapter
- `apache_spark` (Attributes) Apache Spark connection configuration. (see [below for nested schema](#nestedatt--apache_spark))
- `athena` (Attributes) Athena connection configuration. (see [below for nested schema](#nestedatt--athena))
- `bigquery` (Attributes) (see [below for nested schema](#nestedatt--bigquery))
- `databricks` (Attributes) Databricks connection configuration (see [below for nested schema](#nestedatt--databricks))
- `fabric` (Attributes) Microsoft Fabric connection configuration. (see [below for nested schema](#nestedatt--fabric))
- `is_ssh_tunnel_enabled` (Boolean) Whether the connection can use an SSH tunnel
- `name` (String) Connection name
- `oauth_configuration_id` (Number)
- `postgres` (Attributes) PostgreSQL connection configuration. (see [below for nested schema](#nestedatt--postgres))
- `private_link_endpoint_id` (String) Private Link Endpoint ID. This ID can be found using the `privatelink_endpoint` data source
- `redshift` (Attributes) Redshift connection configuration (see [below for nested schema](#nestedatt--redshift))
- `snowflake` (Attributes) Snowflake connection configuration (see [below for nested schema](#nestedatt--snowflake))
- `starburst` (Attributes) Starburst/Trino connection configuration. (see [below for nested schema](#nestedatt--starburst))
- `synapse` (Attributes) Azure Synapse Analytics connection configuration. (see [below for nested schema](#nestedatt--synapse))
- `teradata` (Attributes) Teradata connection configuration. (see [below for nested schema](#nestedatt--teradata))

<a id="nestedatt--apache_spark"></a>
### Nested Schema for `apache_spark`

Read-Only:

- `auth` (String) Auth
- `cluster` (String) Spark cluster for the connection
- `connect_retries` (Number) Connection retries. Default=0
- `connect_timeout` (Number) Connection time out in seconds. Default=10
- `host` (String) Hostname of the connection
- `method` (String) Authentication method for the connection (http or thrift).
- `organization` (String) Organization ID
- `port` (Number) Port for the connection. Default=443
- `user` (String) User


<a id="nestedatt--athena"></a>
### Nested Schema for `athena`

Read-Only:

- `database` (String) Specify the database (data catalog) to build models into (lowercase only).
- `num_boto3_retries` (Number) Number of times to retry boto3 requests (e.g. deleting S3 files for materialized tables).
- `num_iceberg_retries` (Number) Number of times to retry iceberg commit queries to fix ICEBERG_COMMIT_ERROR.
- `num_retries` (Number) Number of times to retry a failing query.
- `poll_interval` (Number) Interval in seconds to use for polling the status of query results in Athena.
- `region_name` (String) AWS region of your Athena instance.
- `s3_data_dir` (String) Prefix for storing tables, if different from the connection's S3 staging directory.
- `s3_data_naming` (String) How to generate table paths in the S3 data directory.
- `s3_staging_dir` (String) S3 location to store Athena query results and metadata.
- `s3_tmp_table_dir` (String) Prefix for storing temporary tables, if different from the connection's S3 data directory.
- `spark_work_group` (String) Identifier of Athena Spark workgroup for running Python models.
- `work_group` (String) Identifier of Athena workgroup.


<a id="nestedatt--bigquery"></a>
### Nested Schema for `bigquery`

Required:

- `gcp_project_id` (String) The GCP project ID to use for the connection

Read-Only:

- `application_id` (String, Sensitive) OAuth Client ID
- `application_secret` (String, Sensitive) OAuth Client Secret
- `auth_provider_x509_cert_url` (String) Auth Provider X509 Cert URL for the Service Account
- `auth_uri` (String) Auth URI for the Service Account
- `client_email` (String) Service Account email
- `client_id` (String) Client ID of the Service Account
- `client_x509_cert_url` (String) Client X509 Cert URL for the Service Account
- `dataproc_cluster_name` (String) Dataproc cluster name for PySpark workloads
- `dataproc_region` (String) Google Cloud region for PySpark workloads on Dataproc
- `execution_project` (String) Project to bill for query execution
- `gcs_bucket` (String) URI for a Google Cloud Storage bucket to host Python code executed via Datapro
- `impersonate_service_account` (String) Service Account to impersonate when running queries
- `job_creation_timeout_seconds` (Number) Maximum timeout for the job creation step
- `job_retry_deadline_seconds` (Number) Total number of seconds to wait while retrying the same query
- `location` (String) Location to create new Datasets in
- `maximum_bytes_billed` (Number) Max number of bytes that can be billed for a given BigQuery query
- `priority` (String) The priority with which to execute BigQuery queries (batch or interactive)
- `private_key` (String, Sensitive) Private Key for the Service Account
- `private_key_id` (String) Private Key ID for the Service Account
- `retries` (Number) Number of retries for queries
- `scopes` (Set of String) OAuth scopes for the BigQuery connection
- `timeout_seconds` (Number) Timeout in seconds for queries
- `token_uri` (String) Token URI for the Service Account


<a id="nestedatt--databricks"></a>
### Nested Schema for `databricks`

Read-Only:

- `catalog` (String) Catalog name if Unity Catalog is enabled in your Databricks workspace.
- `client_id` (String) Required to enable Databricks OAuth authentication for IDE developers.
- `client_secret` (String) Required to enable Databricks OAuth authentication for IDE developers.
- `host` (String) The hostname of the Databricks cluster or SQL warehouse.
- `http_path` (String) The HTTP path of the Databricks cluster or SQL warehouse.


<a id="nestedatt--fabric"></a>
### Nested Schema for `fabric`

Read-Only:

- `database` (String) The database to connect to for this connection.
- `login_timeout` (Number) The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
- `port` (Number) The port to connect to for this connection. Default=1433
- `query_timeout` (Number) The number of seconds used to wait for a query before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
- `retries` (Number) The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.
- `server` (String) The server hostname.


<a id="nestedatt--postgres"></a>
### Nested Schema for `postgres`

Read-Only:

- `dbname` (String) The database name for this connection.
- `hostname` (String) The hostname of the database.
- `port` (Number) The port to connect to for this connection. Default=5432
- `ssh_tunnel` (Attributes) PostgreSQL SSH Tunnel configuration (see [below for nested schema](#nestedatt--postgres--ssh_tunnel))

<a id="nestedatt--postgres--ssh_tunnel"></a>
### Nested Schema for `postgres.ssh_tunnel`

Read-Only:

- `hostname` (String) The hostname for the SSH tunnel.
- `id` (Number) The ID of the SSH tunnel connection.
- `port` (Number) The HTTP port for the SSH tunnel.
- `public_key` (String) The SSH public key generated to allow connecting via SSH tunnel.
- `username` (String) The username to use for the SSH tunnel.



<a id="nestedatt--redshift"></a>
### Nested Schema for `redshift`

Required:

- `hostname` (String) The hostname of the data warehouse.

Read-Only:

- `dbname` (String) The database name for this connection.
- `port` (Number) The port to connect to for this connection. Default=5432
- `ssh_tunnel` (Attributes) Redshift SSH Tunnel configuration (see [below for nested schema](#nestedatt--redshift--ssh_tunnel))

<a id="nestedatt--redshift--ssh_tunnel"></a>
### Nested Schema for `redshift.ssh_tunnel`

Read-Only:

- `hostname` (String) The hostname for the SSH tunnel.
- `id` (Number) The ID of the SSH tunnel connection.
- `port` (Number) The HTTP port for the SSH tunnel.
- `public_key` (String) The SSH public key generated to allow connecting via SSH tunnel.
- `username` (String) The username to use for the SSH tunnel.



<a id="nestedatt--snowflake"></a>
### Nested Schema for `snowflake`

Read-Only:

- `account` (String) The Snowflake account name
- `allow_sso` (Boolean) Whether to allow Snowflake OAuth for the connection. If true, the `oauth_client_id` and `oauth_client_secret` fields must be set
- `client_session_keep_alive` (Boolean) If true, the snowflake client will keep connections for longer than the default 4 hours. This is helpful when particularly long-running queries are executing (> 4 hours)
- `database` (String) The default database for the connection
- `oauth_client_id` (String, Sensitive) OAuth Client ID. Required to allow OAuth between dbt Cloud and Snowflake
- `oauth_client_secret` (String, Sensitive) OAuth Client Secret. Required to allow OAuth between dbt Cloud and Snowflake
- `role` (String) The Snowflake role to use when running queries on the connection
- `warehouse` (String) The default Snowflake Warehouse to use for the connection


<a id="nestedatt--starburst"></a>
### Nested Schema for `starburst`

Read-Only:

- `host` (String) The hostname of the account to connect to.
- `method` (String) The authentication method. Only LDAP for now.
- `port` (Number) The port to connect to for this connection. Default=443


<a id="nestedatt--synapse"></a>
### Nested Schema for `synapse`

Read-Only:

- `database` (String) The database to connect to for this connection.
- `host` (String) The server hostname.
- `login_timeout` (Number) The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
- `port` (Number) The port to connect to for this connection. Default=1433
- `query_timeout` (Number) The number of seconds used to wait for a query before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
- `retries` (Number) The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.


<a id="nestedatt--teradata"></a>
### Nested Schema for `teradata`

Read-Only:

- `host` (String) The hostname of the database.
- `port` (String) The port to connect to for this connection. Default=1025
- `request_timeout` (Number) The number of seconds used to establish a connection before failing. Defaults to 0, which means that the timeout is disabled or uses the default system settings.
- `retries` (Number) The number of automatic times to retry a query before failing. Defaults to 1. Queries with syntax errors will not be retried. This setting can be used to overcome intermittent network issues.
- `tmode` (String) The transaction mode to use for the connection.

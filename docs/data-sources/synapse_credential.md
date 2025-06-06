---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dbtcloud_synapse_credential Data Source - dbtcloud"
subcategory: ""
description: |-
  Synapse credential data source.
---

# dbtcloud_synapse_credential (Data Source)

Synapse credential data source.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credential_id` (Number) Credential ID
- `project_id` (Number) Project ID

### Read-Only

- `adapter_type` (String) The type of the adapter (synapse)
- `authentication` (String) Authentication type (SQL, ActiveDirectoryPassword, ServicePrincipal)
- `client_id` (String) The client ID of the Azure Active Directory service principal. This is only used when connecting to Azure SQL with an AAD service principal.
- `id` (String) The ID of this data source. Contains the project ID and the credential ID.
- `schema` (String) The schema where to create the dbt models
- `schema_authorization` (String) Optionally set this to the principal who should own the schemas created by dbt
- `tenant_id` (String) The tenant ID of the Azure Active Directory instance. This is only used when connecting to Azure SQL with a service principal.
- `user` (String) The username of the Synapse account to connect to. Only used when connection with AD user/pass

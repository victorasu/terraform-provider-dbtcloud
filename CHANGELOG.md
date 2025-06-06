# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html),
and is generated by [Changie](https://github.com/miniscruff/changie).


# [v1.1.1](https://github.com/dbt-labs/terraform-provider-dbtcloud/compare/v1.1.0...v1.1.1)
### Changes
* Minor redshift sl credential refactoring. Code review comments solving for PR #451
### Fixes
* Fixed default parameter for Redshift leveraging of Postgres Credential
* Fixed recreation prompt for Azure AD repositories

# [v1.1.0](https://github.com/dbt-labs/terraform-provider-dbtcloud/compare/v1.0.0...v1.1.0)
### Changes
* Implemented BigQuery for Semantic Layer Credentials
* Implemented Redshift for Semantic Layer Credentials
* Added support for Synapse Credential
### Fixes
* Obfuscate error messages from the API that would display the token
* Fixed issue with environment optional parameters being seen as null instead of being disregarded
* Fix repository create and apply to be consistent with multiple applies. Issue #433
### Behind the scenes
* Deparallelized flaky test
### Documentation
* Added changie for CHANGELOG management

# [1.0.0](https://github.com/dbt-labs/terraform-provider-dbtcloud/compare/v0.0.0...1.0.0)

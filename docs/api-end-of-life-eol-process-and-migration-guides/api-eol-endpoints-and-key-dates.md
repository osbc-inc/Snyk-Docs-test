# API EOL endpoints and key dates

## APIs at EOL

Beginning July 22, 2024, the following endpoints will follow the EOL process:

| Endpoint | Endpoint type | EOL date | Migration guide |
| --- | --- | --- | --- |
| [Group and org level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/group-level-audit-logs/get-group-level-audit-logs) | v1 | January 22, 2025 | [Search Audit Logs (Group and Org) v1 API to GA REST Audit logs API migration guide](guides-to-migration/search-audit-logs-group-and-org-v1-api-to-ga-rest-audit-logs-api-migration-guide.md) |
| Get all issues by [Org](https://apidocs.snyk.io/experimental?version=2023-03-10\~experimental&\_gl=1\*d7o8is\*\_gcl\_aw\*R0NMLjE3MTIwNjc4NjcuQ2owS0NRancyYTZ3QmhDVkFSSXNBQlBlSDF0VG1UNmo0cnNrQTVPRmNLVU02cFMyNVc1Q3lpWWhLRFVqZGdfWDZTREJ6Z0NWSGZTZUtzY2FBb3lORUFMd193Y0I.\*\_gcl\_au\*MTU3NDc2MzU2LjE3MTI5Mzg4MzA.\*\_ga\*MTE2NjY3NTQyNC4xNjQ3OTU0NjA1\*\_ga\_X9SH3KP7B4\*MTcxOTQwNzU4My4yNjguMS4xNzE5NDA3ODA1LjQ5LjAuMA..#get-/orgs/-org\_id-/issues) and [Group](https://apidocs.snyk.io/experimental?version=2023-03-10\~experimental&\_gl=1\*d7o8is\*\_gcl\_aw\*R0NMLjE3MTIwNjc4NjcuQ2owS0NRancyYTZ3QmhDVkFSSXNBQlBlSDF0VG1UNmo0cnNrQTVPRmNLVU02cFMyNVc1Q3lpWWhLRFVqZGdfWDZTREJ6Z0NWSGZTZUtzY2FBb3lORUFMd193Y0I.\*\_gcl\_au\*MTU3NDc2MzU2LjE3MTI5Mzg4MzA.\*\_ga\*MTE2NjY3NTQyNC4xNjQ3OTU0NjA1\*\_ga\_X9SH3KP7B4\*MTcxOTQwNzU4My4yNjguMS4xNzE5NDA3ODA1LjQ5LjAuMA..#get-/groups/-group\_id-/issues) | Non-GA REST | October 22, 2024 | [REST Issues experimental API to GA API migration guide](guides-to-migration/rest-issues-experimental-api-to-ga-api-migration-guide.md) |

{% hint style="info" %}
\*Experimental versions from 2023-03-10 inclusive up to 2023-09-29 exclusive.
{% endhint %}

## Brownouts

A brownout occurs when Snyk temporarily suspends an endpoint from being usable, returning a `410 gone` response when a user calls the endpoint.

Snyk brownouts for APIs that are part of an end-of-life cycle will occur at 12:00 UTC. For the end-of-life cycle beginning July 22, 2024, the brownouts will occur on the following dates. Users will see a reminder two weeks before the brownout through an announcement on [updates.snyk.io](http://updates.snyk.io/):

| Endpoints | Brownout date | Duration |
| --- | --- | --- |
| Get all issues by [Org](https://apidocs.snyk.io/experimental?version=2023-03-10\~experimental&\_gl=1\*d7o8is\*\_gcl\_aw\*R0NMLjE3MTIwNjc4NjcuQ2owS0NRancyYTZ3QmhDVkFSSXNBQlBlSDF0VG1UNmo0cnNrQTVPRmNLVU02cFMyNVc1Q3lpWWhLRFVqZGdfWDZTREJ6Z0NWSGZTZUtzY2FBb3lORUFMd193Y0I.\*\_gcl\_au\*MTU3NDc2MzU2LjE3MTI5Mzg4MzA.\*\_ga\*MTE2NjY3NTQyNC4xNjQ3OTU0NjA1\*\_ga\_X9SH3KP7B4\*MTcxOTQwNzU4My4yNjguMS4xNzE5NDA3ODA1LjQ5LjAuMA..#get-/orgs/-org\_id-/issues) and [Group](https://apidocs.snyk.io/experimental?version=2023-03-10\~experimental&\_gl\_1*d7o8is*_gcl_aw*R0NMLjE3MTIwNjc4NjcuQ2owS0NRancyYTZ3QmhDVkFSSXNBQlBlSDF0VG1UNmo0cnNrQTVPRmNLVU02cFMyNVc1Q3lpWWhLRFVqZGdfWDZTREJ6Z0NWSGZTZUtzY2FBb3lORUFMd193Y0I.*_gcl_au*MTU3NDc2MzU2LjE3MTI5Mzg4MzA.*_ga*MTE2NjY3NTQyNC4xNjQ3OTU0NjA1*_ga_X9SH3KP7B4*MTcxOTQwNzU4My4yNjguMS4xNzE5NDA3ODA1LjQ5LjAuMA..#get-/orgs/-org_id-/issues) | September 12 | 1 hour |
| [Group and org level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/group-level-audit-logs/get-group-level-audit-logs) | October 8 | 1 hour |
| [Group and org level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/group-level-audit-logs/get-group-level-audit-logs) | November 12 | 2 hours |
| [Group and org level audit logs](https://snyk.docs.apiary.io/#reference/audit-logs/group-level-audit-logs/get-group-level-audit-logs) | December 10 | 4 hours |
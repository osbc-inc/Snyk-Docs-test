# import 명령어에 필요한 가져오기 대상 데이터 생성

`import:data` 유틸리티를 사용하여 import 명령어에서 필요한 가져오기 JSON 데이터를 생성하는 데 도움이 되도록 합니다. 이 페이지의 지침을 따르세요:

- `import:data` 유틸리티의 사용법
  - [GitHub](creating-import-targets-data-for-import-command.md#github.com-and-github-enterprise)
  - [GitLab](creating-import-targets-data-for-import-command.md#gitlab.com-and-hosted-gitlab)
  - [Azure Repos](creating-import-targets-data-for-import-command.md#dev.azure.com-and-hosted-azure)
  - [Bitbucket Server](creating-import-targets-data-for-import-command.md#bitbucket-server)
  - [Bitbucket Cloud](creating-import-targets-data-for-import-command.md#bitbucket-cloud)
- [권장사항](creating-import-targets-data-for-import-command.md#recommendations)

SCM에 해당하는 섹션에서 JSON 형식의 조직 데이터를 생성하는 방법에 대한 정보를 검토하고 후속 단계를 따르세요. **반드시 SCM 개인 액세스 토큰 또는 사용자 이름 및 암호 자격 증명을 환경 변수로 설정해야 합니다.**

기본적으로 아카이브된 리포지토리는 제외됩니다.

## GitHub.com 및 GitHub Enterprise

가져오기 명령어를 실행할 때 필요한 import 대상을 매핑하는 데 도움이 되는 JSON 형식의 조직 데이터가 필요합니다. 필요한 형식은 다음과 같습니다:

```
{
  "orgData": [
    {
      "name": "<org_name_in_gh_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "github": "<snyk_org_integration_id>",
        "github-enterprise": "<snyk_org_integration_id>"
      },
    },
    {...}
  ]
}
```

참고: GitHub 또는 GitHub Enterprise 조직의 `"name"`은 GitHub API를 사용하여 해당 조직에 속한 모든 리포지토리를 나열하기 위해 필요합니다. 그 조직 이름에 따른 Snyk-특정 데이터는 그 조직에 속하는 모든 리포지토리가 주어진 Snyk 조직으로 가져오기되어야 한다는 정보를 생성할 때 사용됩니다. 이 유틸리티는 지정되어 있습니다. 가져오기 데이터를 사용자 정의하려면 [가져오기 시작](kicking-off-an-import.md)에서 설명된 대로 직접 만들어야 합니다.

* GitHub 및 GitHub Enterprise 조직을 GitHub [/orgs API](https://docs.github.com/en/free-pro-team@latest/rest/reference/orgs)를 사용하여 나열할 수 있습니다.
* Snyk API 엔드포인트 [List](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-1) (integrations)를 사용하여 통합을 나열할 수 있습니다.
* Snyk API 엔드포인트 [List all organizatons in a group](../../../snyk-api/reference/orgs.md#groups-group\_id-orgs)를 사용하여 그룹 관리자가 속한 모든 조직을 나열하여 모든 조직 ID를 찾을 수 있습니다.

이 유틸리티를 사용하는 단계는 다음과 같습니다:

1. [GitHub.com 개인 액세스 토큰](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token)을 환경 변수로 설정합니다: `export GITHUB_TOKEN=your_personal_access_token`.
2. 시작 부분에서 설명된대로 JSON 형식의 조직 데이터를 생성합니다.
3. 가져오기 데이터를 생성하기 위해 명령어를 실행합니다:\
   **GitHub.com:** `DEBUG=snyk* GITHUB_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=github --integrationType=github`\
   **GitHub Enterprise:** `DEBUG=snyk* GITHUB_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=github-enterprise --integrationType=github-enterprise --sourceUrl=https://ghe.custom.com`
4. 생성된 데이터를 `import` 명령어에 제공하여 [가져오기를 시작](kicking-off-an-import.md)합니다.

## GitLab.com 및 Hosted GitLab

가져오기 대상을 매핑하기 위해 필요한 JSON 형식의 조직 데이터가 필요합니다. 필요한 형식은 다음과 같습니다:

```
{
  "orgData": [
    {
      "name": "<group_name_in_gitlab_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "gitlab": "<snyk_org_integration_id>",
      },
    },
    {...}
  ]
}
```

참고: GitLab 그룹의 "name"은 GitLab API를 사용하여 해당 그룹에 속한 모든 프로젝트를 나열하기 위해 필요합니다. 그 그룹 이름에 따른 Snyk-특정 데이터는 그 그룹의 모든 프로젝트가 주어진 Snyk 조직으로 가져오기되어야 한다는 정보를 생성할 때 사용됩니다. 이 유틸리티는 지정되어 있습니다. 가져오기 데이터를 사용자 정의하려면 [가져오기 시작](kicking-off-an-import.md)에서 설명된 대로 직접 만들어야 합니다.

* GitLab 그룹은 GitLab의 [/groups API](https://docs.gitlab.com/ee/api/groups.html)를 사용하여 나열할 수 있습니다.
* Snyk API 엔드포인트 [List](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-1) (integrations)를 사용하여 통합을 나열할 수 있습니다.
* Snyk API 엔드포인트 [List all organizatons in a group](../../../snyk-api/reference/orgs.md#groups-group\_id-orgs)를 사용하여 그룹 관리자가 속한 모든 조직을 나열하여 모든 조직 ID를 찾을 수 있습니다.

이 유틸리티를 사용하는 단계는 다음과 같습니다:

1. [GitLab 개인 액세스 토큰](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)을 환경 변수로 설정합니다: `export GITLAB_TOKEN=your_personal_access_token`.
2. 시작 부분에서 설명된대로 JSON 형식의 조직 데이터를 생성합니다.
3. 가져오기 데이터를 생성하기 위해 명령어를 실행합니다:\
   **Gitlab.com:** `DEBUG=snyk* GITLAB_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=gitlab --integrationType=gitlab`\
   **Hosted Gitlab:** `DEBUG=snyk* GITLAB_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=gitlab --integrationType=gitlab --sourceUrl=https://gitlab.custom.com`
4. 생성된 데이터를 `import` 명령어에 제공하여 [가져오기를 시작](kicking-off-an-import.md)합니다.

## dev.azure.com 및 Hosted Azure

**참고: 이 도구는 Azure API v1을 사용합니다.**

가져오기 대상을 매핑하기 위해 필요한 JSON 형식의 조직 데이터가 필요합니다. 필요한 형식은 다음과 같습니다:

```
{
  "orgData": [
    {
      "name": "<org_name_in_azure_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "azure-repos": "<snyk_org_integration_id>",
      },
    },
    {...}
  ]
}
```

참고: Azure 조직의 "name"은 Azure API를 사용하여 해당 조직에 속한 모든 프로젝트 및 리포지토리를 나열하기 위해 필요합니다. 그 조직 이름에 따른 Snyk-특정 데이터는 그 조직의 모든 프로젝트가 주어진 Snyk 조직으로 가져오기되어야 한다는 정보를 생성할 때 사용됩니다. 이 유틸리티는 지정되어 있습니다. 가져오기 데이터를 사용자 정의하려면 [가져오기 시작](kicking-off-an-import.md)에서 설명된 대로 직접 만들어야 합니다.

* Azure에서 조직 이름은 [Azure-Devops-site](https://dev.azure.com/)의 왼쪽 창에 나열됩니다.
* Snyk API 엔드포인트 [List](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-1) (integrations)를 사용하여 통합을 나열할 수 있습니다.
* Snyk API 엔드포인트 [List all organizatons in a group](../../../snyk-api/reference/orgs.md#groups-group\_id-orgs)를 사용하여 그룹 관리자가 속한 모든 조직을 나열하여 모든 조직 ID를 찾을 수 있습니다.

이 유틸리티를 사용하는 단계는 다음과 같습니다:

1. [Azure 개인 액세스 토큰](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page)을 환경 변수로 설정합니다: `export AZURE_TOKEN=your_personal_access_token`.
2. 시작 부분에서 설명된대로 JSON 형식의 조직 데이터를 생성합니다.
3. 가져오기 데이터를 생성하기 위해 명령어를 실행합니다:\
   **dev.azure.com:** `DEBUG=snyk* AZURE_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=azure-repos --integrationType=azure-repos`\
   **Hosted Azure:** `DEBUG=snyk* AZURE_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=azure-repos --integrationType=azure-repos --sourceUrl=https://azure.custom.com`
4. 생성된 데이터를 `import` 명령어에 제공하여 [가져오기를 시작](kicking-off-an-import.md)합니다.

## Bitbucket Server

**참고: 이 도구는 Bitbucket 서버 API 버전 1.0을 사용합니다.**

가져오기 대상을 매핑하기 위해 필요한 JSON 형식의 조직 데이터가 필요합니다. 필요한 형식은 다음과 같습니다:

```
{
  "orgData": [
    {
      "name": "<project_name_in_bitbucket_server_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "bitbucket-server": "<snyk_org_integration_id>",
      },
    },
    {...}
  ]
}
```

참고: Bitbucket 서버 프로젝트의 "name"은 Bitbucket 서버 API를 사용하여 해당 프로젝트에 속한 모든 리포지토리를 나열하기 위해 필요합니다. 그 프로젝트 이름에 따른 Snyk-특정 데이터는 그 프로젝트의 모든 리포지토리가 주어진 Snyk 조직으로 가져오기되어야 한다는 정보를 생성할 때 사용됩니다. 가져오기 데이터를 사용자 정의하려면 [가져오기 시작](kicking-off-an-import.md)에서 설명된 대로 직접 만들어야 합니다.

* Bitbucket 서버 프로젝트는 BitBucket [/projects API](https://docs.atlassian.com/bitbucket-server/rest/7.19.2/bitbucket-rest.html)를 사용하여 나열할 수 있습니다.
* Snyk API 엔드포인트 [List](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-1) (integrations)를 사용하여 통합을 나열할 수 있습니다.
* Snyk API 엔드포인트 [List all organizatons in a group](../../../snyk-api/reference/orgs.md#groups-group\_id-orgs)를 사용하여 그룹 관리자가 속한 모든 조직을 나열하여 모든 조직 ID를 찾을 수 있습니다.

이 유틸리티를 사용하는 단계는 다음과 같습니다:

1. [Bitbucket Server 개인 액세스 토큰](https://www.jetbrains.com/help/youtrack/standalone/integration-with-bitbucket-server.html#enable-youtrack-integration-bbserver)을 환경 변수로 설정합니다: `export BITBUCKET_SERVER_TOKEN=your_personal_access_token`.
2. 시작 부분에서 설명된대로 JSON 형식의 조직 데이터를 생성합니다.
3. 가져오기 데이터를 생성하기 위해 명령어를 실행합니다:\
   **Bitbucket Server:** `DEBUG=snyk* BITBUCKET_SERVER_TOKEN=*** SNYK_TOKEN=*** snyk-api-import import:data --orgsData=path/to/snyk-orgs.json --source=bitbucket-server --integrationType=bitbucket-server --sourceUrl=https://bitbucket-server.dev.example.com`
4. 생성된 데이터를 `import` 명령어에 제공하여 [가져오기를 시작](kicking-off-an-import.md)합니다.

## Bitbucket Cloud

**이 도구는 Bitbucket Cloud API 버전 2.0을 사용합니다.**

가져오기 대상을 매핑하기 위해 필요한 JSON 형식의 조직 데이터가 필요합니다. 필요한 형식은 다음과 같습니다:

```
{
  "orgData": [
    {
      "name": "<workspace_name_in_bitbucket_cloud_used_to_list_repos>",
      "orgId": "<snyk_org_id>",
      "integrations": {
        "bitbucket-cloud": "<snyk_org_integration_id>",
      },
    },
    {...}
  ]
}
```

참고: Bitbucket Cloud 워크스페이스의 "name"은 Bitbucket Cloud API를 사용하여 해당 워크스페이스에 속한 모든 리포지토리를 나열하기 위해 필요합니다. 그 워크스페이스 이름에 따른 Snyk-특정 데이터는 그 워크스페이스의 모든 리포지토리가 주어진 Snyk 조직으로 가져오기되어야 한다는 정보를 생성할 때 사용됩니다. 이 유틸리티는 지정되어 있습니다. 가져오기 데이터를 사용자 정의하려면 [가져오기 시작](kicking-off-an-import.md)에서 설명된 대로 직접 만들어야 합니다.

* Bitbucket Cloud 워크스페이스는 BItBucket [/workspaces API](https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-group-workspaces)를 사용하여 나열할 수 있습니다.
* Snyk API 엔드포인트 [List](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-1) (integrations)를 사용하여 통합을 나열할 수 있습니다.
* Snyk API 엔드포인트 [List all organizatons in a group](../../../snyk-api/reference/orgs.md#groups-group\_id-orgs)를 사용하여 그룹 관리자가 속한 모든 조직을 나열하여 모든 조직 ID를 찾을 수 있습니다.

이 유틸리티를 사용하는 단계는 다음과 같습니다:

1. Bitbucket Cloud의 사용자 이름 및 암호 자격 증명을 환경 변수로 설정합니다:\
   `export BITBUCKET_CLOUD_USERNAME=your_bitbucket_cloud_username`\
   `export BITBUCKET_CLOUD_PASSWORD=your_bitbucket_cloud_password`
2. 시작 부분에서 설명된대로 JSON 형식의 조직 데이터를 생성합니다.
3. 가져오기 데이터를 생성하기 위해 명령어를 실행합니다:\
   **Bitbucket Cloud:**
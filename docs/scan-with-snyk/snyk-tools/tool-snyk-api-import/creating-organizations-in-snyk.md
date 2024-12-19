# Snyk에서 조직 생성

이 페이지에는 Snyk에서 조직(Orgs)을 생성하는 방법에 대한 지침이 있습니다:

* Snyk에서 조직을 생성하기 위한 데이터 생성
  * [GitHub](creating-organizations-in-snyk.md#github.com-and-github-enterprise)
  * [GitLab](creating-organizations-in-snyk.md#gitlab.com-and-hosted-gitlab)
  * [Bitbucket Server](creating-organizations-in-snyk.md#bitbucket-server)
  * [Bitbucket Cloud](creating-organizations-in-snyk.md#bitbucket-cloud)
  * [Azure](creating-organizations-in-snyk.md#azure)
* [조직 생성 방법](creating-organizations-in-snyk.md#methods-of-creating-organizations)
  * [API 사용](creating-organizations-in-snyk.md#using-the-api)
  * [`orgs:create` 유틸리티 사용](creating-organizations-in-snyk.md#using-the-orgs-create-utility)
* [권장 사항](creating-organizations-in-snyk.md#recommendations)

먼저, 가져오기를 시작하기 전에 Snyk을 설정하여 프로젝트로 채울 조직을 보유한 상태여야 합니다.

Snyk은 가져오기 원본과 같은 수의 조직을 보유하는 것을 권장합니다. GitHub의 경우, 이는 Snyk에 GitHub 조직을 미러링하는 것을 의미합니다. `snyk-api-import` 도구는 Groups 및 Organizations을 사용할 때 이를 보다 간단하게 만들기 위한 유틸리티를 제공합니다.

## `orgs:data` 유틸리티를 사용하여 Snyk에서 조직을 만들기 위해 필요한 데이터 생성

이 유틸리티는 GitHub.com, GitHub Enterprise, GitLab, Bitbucket Server 또는 Bitbucket Cloud 조직 구조를 Snyk에서 미러링하는 데 필요한 데이터를 생성하는 데 도움이 됩니다. 이 유틸리티는 GitHub.com, GitHub Enterprise, GitLab, Bitbucket Server 또는 Bitbucket Cloud의 모든 조직이 Snyk에서 구성되어야 한다고 가정합니다. 원하는 것이 아닌 경우 API 엔드포인트 [새로운 조직 만들기](../../../snyk-api/reference/organizations-v1.md#org)를 직접 사용하여 필요한 구조를 만들 수 있습니다.

### 옵션

```
  --source             가져올 대상의 소스
                       (예: Github, Github Enterprise, Gitlab,
                       Bitbucket Server) [필수].
  --groupId            Snyk의 그룹 내 공개 ID
                       (그룹 설정에서 사용 가능) [필수].
  --sourceUrl          조직을 나열할 수 있는 소스 API의 사용자 정의 기본 URL
                       (예: GitHub Enterprise URL).
  --sourceOrgPublicId  선택한 모든 지원하는 Organization 설정을 복사할 수 있는
                       Snyk의 Organization의 공개 ID.
  --skipEmptyOrgs      대상이 없는 조직 건너 뛰기
                       (예: 저장소가 없는 Github 조직).
```

### GitHub.com 및 GitHub Enterprise

1. [GitHub.com 개인 액세스 토큰](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token)을 환경 변수로 설정: `export GITHUB_TOKEN=your_personal_access_token`
2. 조직 데이터를 생성하는 명령어 실행:

* **GitHub.com:** `snyk-api-import orgs:data --source=github --groupId=<snyk_group_id>`
* **GitHub Enterprise:** `snyk-api-import orgs:data --source=github-enterprise --groupId=<snyk_group_id> -- sourceUrl=https://ghe.custom.github.com/`

이렇게 하면 `group-<snyk_group_id>-github-<com|enterprise>-orgs.json` 파일에 조직 데이터가 생성됩니다.

### GitLab.com 및 Hosted GitLab

1. [GitLab 개인 액세스 토큰](https://docs.gitlab.com/ee/user/profile/personal\_access\_tokens.html)을 환경 변수로 설정: `export GITLAB_TOKEN=your_personal_access_token`
2. 조직 데이터를 생성하는 명령어 실행:

* **GitLab:** `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id>`
* **Hosted GitLab:** `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id> -- sourceUrl=https://gitlab.custom.com`

이렇게 하면 `group-<snyk_group_id>-gitlab-orgs.json` 파일에 조직 데이터가 생성됩니다.

### Bitbucket Server

**Bitbucket Server는 호스팅 환경이며 명령에서 Bitbucket Server 인스턴스의 사용자 정의 URL을 제공해야 합니다.**

1. [Bitbucket Server 액세스 토큰](https://www.jetbrains.com/help/youtrack/standalone/integration-with-bitbucket-server.html#enable-youtrack-integration-bbserver)을 환경 변수로 설정: `export BITBUCKET_SERVER_TOKEN=your_personal_access_token`
2. 조직 데이터를 생성하는 명령어 실행:\
   `snyk-api-import orgs:data --source=bitbucket-server --groupId=<snyk_group_id> --sourceUrl=https://bitbucket-server.custom.com`

이렇게 하면 `group-<snyk_group_id>-bitbucket-server-orgs.json` 파일에 조직 데이터가 생성됩니다.

### Bitbucket Cloud

**Bitbucket Cloud의 URL은 https://bitbucket.org/ 입니다.**

1. Bitbucket Cloud 사용자 이름과 비밀번호를 환경 변수로 설정: `export BITBUCKET_CLOUD_USERNAME=your_bitbucket_cloud_username` 및 `export BITBUCKET_CLOUD_PASSWORD=your_bitbucket_cloud_password`
2. 조직 데이터를 생성하는 명령어 실행:\
   `snyk-api-import orgs:data --source=bitbucket-cloud --groupId=<snyk_group_id>`

이렇게 하면 `group-<snyk_group_id>-bitbucket-cloud-orgs.json` 파일에 조직 데이터가 생성됩니다.

### Azure

**Azure의 경우, 이 단계는 수동으로 수행해야 합니다.** Azure는 Azure 조직을 가져오기 위한 API 호출이 없기 때문에 다음 명령이 실행되려면 Orgs 파일을 수동으로 생성해야 합니다. 파일은 다음과 같이 형식화되어야 합니다:

```
{
   "orgs":[
      {
         "name":"AZURE_ORG의_이름",
         "groupId":"YOUR_SNYK_GROUP_ID",
         "sourceOrgId":"SETTINGS를 복사할 SNYK_ORG_ID"   // **optional**
      },
      {
         "name":"다른 AZURE_ORG의 이름",
         "groupId":"YOUR_SNYK_GROUP_ID",
         "sourceOrgId":"SETTINGS를 복사할 SNYK_ORG_ID"  // **optional**
      }
   ]
}
```

파일을 생성한 후, 이를 [`orgs:create` 명령](https://github.com/snyk/snyk-api-import/blob/0e5162d29dec7f1d5acde247cc8da0553871db3f/docs/orgs.md#creating-organizations-in-snyk-1)에 공급할 수 있습니다.

## 조직 생성 방법

API 엔드포인트 [새로운 조직 생성](../../../snyk-api/reference/organizations-v1.md#org)을 사용하여 그룹 내에서 조직을 생성하거나 제공된 유틸리티를 사용할 수 있습니다.

### API 사용

엔드포인트 [새로운 조직 생성](../../../snyk-api/reference/organizations-v1.md#org)을 통해 생성된 데이터를 사용하여 그룹 내에서 조직을 생성합니다.

### `orgs:create` 유틸리티 사용

1. `SNYK_TOKEN` 환경 변수를 설정합니다. [Snyk API 토큰](https://app.snyk.io/account)을 입력합니다.
2. 조직을 생성하는 명령어를 실행합니다: `snyk-api-import orgs:create --noDuplicateNames --includeExistingOrgsInOutput --file=group-<snyk_group_id>-github-<com|enterprise>-orgs.json`.
3. 중복된 이름이 이미 사용 중인 경우 조직을 생성하지 않으려면 `noDuplicateNames` 플래그(선택 사항)를 사용합니다.
4. 기본값은 `true`인 `includeExistingOrgsInOutput` 플래그(선택 사항)를 사용하여 기존 조직에 대한 정보와 최신 조직에 대한 정보를 기록합니다. 이 플래그를 `false`로 설정하려면 명령에 `--no-includeExistingOrgsInOutput`를 사용합니다. 예: `snyk-api-import orgs:create --no-includeExistingOrgsInOutput --file=group-<snyk_group_id>-github-<com|enterprise>-orgs.json`

이 명령에 필요한 파일 형식은 다음과 같습니다:

```
"orgs": [
  {
    "groupId": "<public_snyk_group_id>",
    "name": "<name_of_the_organization>",
    "sourceOrgId": "<public_snyk_organization_id>"
  }
]
```

* `groupId` - 조직을 생성할 Snyk 그룹의 공개 ID
* `name` - 조직 생성 시 사용할 이름
* `sourceOrgId` - **선택 사항** 설정을 복사할 Snyk 조직의 공개 ID

### 권장 사항

* 가져오기 알림을 받지 않으려면 [알림 설정](../../../snyk-api/reference/organizations-v1.md#org-orgid-notification-settings) 엔드포인트를 사용하여 이메일 알림 등을 비활성화하세요.
* 가져오기가 완료될 때까지 추가 요청을 보내지 않도록 PR 수정 및 PR 체크를 비활성화하려면 [Update](../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-integrationid-settings) (통합 설정) 엔드포인트를 사용하세요. (GitHub, GitLab, Bitbucket 등에 추가 요청을 보내지 않도록)
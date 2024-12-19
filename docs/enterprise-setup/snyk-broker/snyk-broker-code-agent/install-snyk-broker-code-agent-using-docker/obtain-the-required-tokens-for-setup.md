# 설정에 필요한 토큰 얻기

Code Agent를 Broker 클라이언트와 함께 설정하려면 다음 토큰이 필요합니다:

- **Snyk API 토큰** - 이 토큰은 **Code Agent 설정에 필요**합니다. 이는 `-e SNYK_TOKEN` 매개변수에서 사용되어 Code Agent 구성요소를 귀하의 Snyk 계정과 인증하는 데 사용됩니다. [Snyk API 토큰 얻기 및 사용하기](../../../../getting-started/#obtain-and-use-your-snyk-api-token) 및 [Code Agent 설정](set-up-the-code-agent.md)을 참조하십시오.
- **Broker 토큰** - 이 토큰은 **Broker 클라이언트 설정에 필요**합니다. 이는 `-e BROKER_TOKEN` 매개변수에서 사용됩니다. Broker 토큰은 기본적으로 특정 조직과 특정 통합 SCM에 연결되며, 해당 조직 및 SCM에 대한 Snyk Broker 배포를 활성화합니다. 각 SCM마다 다른 Broker 토큰이 필요합니다. 자세한 내용은 [Snyk Broker - Code Agent를 위한 Broker 토큰 얻기](obtain-the-required-tokens-for-setup.md#obtain-your-broker-token-for-snyk-broker-code-agent)을 참조하십시오.
- **통합 SCM 토큰** - 이 토큰은 Broker 클라이언트 설정에 필요합니다. 이는 `-e <SCM>_TOKEN` 매개변수에서 사용됩니다. 예를 들어, `-e GITHUB_TOKEN=xxx…`와 같이 사용하여 Snyk Broker와 Snyk Code의 작동에 필요한 특정 권한으로 SCM에 액세스를 활성화합니다. [SCM 토큰 얻기](obtain-the-required-tokens-for-setup.md#obtain-your-scm-token)을 참조하십시오.

필요한 토큰을 획득한 후 안전하고 접근 가능한 위치에 저장하십시오. Code Agent 및 클라이언트 Broker 구성을 시작할 때 이러한 토큰을 사용해야합니다.

## Snyk Broker - Code Agent를 위한 Broker 토큰 얻기

다음은 Broker 토큰을 얻는 방법입니다:

- **기존 Broker 토큰을 사용하여 Code Agent를 설정** - 이미 동일한 조직과 SCM에서 다른 Snyk 제품용으로 Broker 클라이언트를 실행할 때 사용한 Broker 토큰이 있는 경우, 해당 Broker 토큰을 Snyk Broker - Code Agent를 설정하는 데도 사용할 수 있습니다.
- **다중 Snyk 조직에 동일한 Broker 토큰 사용** -
  - **새 조직** - 기존에 Broker 토큰을 갖고 있는 조직을 기반으로 새로운 조직을 만들 때, 새로운 조직을 만들 때 기존 Broker 토큰이 복제되고 새로운 조직에도 사용할 수 있습니다.
  - **기존 조직** - 기존 Broker 토큰을 기존 조직에 대해 다른 기존 조직에 사용하려는 경우, 다음 API를 사용할 수 있습니다: [통합 복제 (설정 및 자격 증명 포함)](../../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-integrationid-clone).
- **중복을 위해 Broker 토큰 사용** - 동일한 조직 및 SCM에 대해 신뢰성을 위해 두 개의 Broker 클라이언트를 설정하는 경우, 두 Broker 클라이언트 모두에 동일한 Broker 토큰을 사용해야합니다. Snyk Broker 토큰을 다음과 같이 얻을 수 있습니다:
  - 권장: Snyk 계정 담당팀에게 Broker 토큰 생성을 요청하고 웹 UI에서 얻으십시오.
  - Snyk API를 사용하여 Broker 토큰 생성(이어지는 지침 참조)하는 방법.

두 방법 중 하나로 Broker 토큰을 생성한 후, [웹 UI에서 Broker 토큰을 얻을 수 있습니다](obtain-the-required-tokens-for-setup.md#obtain-your-broker-token-from-the-web-ui).

### Snyk API를 사용하여 Broker 토큰 생성

다음과 같이 API를 사용하여 Broker 토큰을 생성할 수 있습니다:

1. [기존 통합 업데이트](../../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-type) 엔드포인트를 사용하여 특정 조직 및 특정 SCM에 대해 Snyk Broker를 활성화합니다. 이는 UI에서 Broker 토큰을 생성합니다.
2. Snyk Broker를 활성화한 후 프로그래밍 방식으로 Broker 토큰을 생성하려면, [새로운 Broker 토큰 제공](../../../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-integrationid-authentication-provision-token) 엔드포인트를 사용하여 Broker 토큰을 생성합니다. 생성된 Broker 토큰은 API 응답 본문 및 웹 UI에서 확인할 수 있습니다.
3. 생성된 Broker 토큰을 복사하여 안전한 위치에 저장하고 나중에 웹 UI에서 얻거나 나중에 사용하십시오.

### 웹 UI에서 Broker 토큰 얻기

Broker 토큰은 생성된 후 웹 UI에서 표시됩니다. 이 토큰을 얻으려면 다음 단계를 따르십시오.

1. Snyk 웹 UI에서 Snyk Broker를 설정하려는 조직을 선택하십시오.
2. 선택한 조직에서 **통합**을 선택하십시오. Snyk Broker를 연결하길 원하는 통합을 찾아 **설정** 아이콘을 클릭하십시오.
3. 선택한 통합의 **설정** 페이지에서 **Broker 자격 증명** 섹션에서 **토큰** 상자에서 Broker 토큰을 복사하고 나중에 사용할 수 있도록 저장하십니다:

<figure><img src="../../../../.gitbook/assets/Snyk Broker - Broker Token - box.png" alt="Broker 토큰 복사"><figcaption><p>Broker 토큰 복사</p></figcaption></figure>

## SCM 토큰 얻기

**SCM 토큰을 얻으려면,** Snyk Broker와 통합하려는 SCM에서 제공하는 지침을 따르고 필요한 권한으로 토큰을 생성하십시오.

다음과 같은 다른 SCM에 대해 필요한 SCM 토큰은:

**GitHub 및 GitHub Enterprise**:

`GITHUB_TOKEN=` - GitHub 개인 엑세스 토큰. 스코프: **`repo, read:org`** 및 **`admin:repo_hook`**.

GitHub 문서 - [_개인 엑세스 토큰 생성_](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)\_

**Gitlab**:

**`GITLAB_TOKEN=`** - GitLab 개인 엑세스 토큰. **`Maintainer`** 권한이 있는 Gitlab 계정. 스코프: **`api`**.

Gitlab 문서 - [_개인 엑세스 토큰_](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)\_

**Azure Repos**:

**`AZURE_REPOS_TOKEN=`** - Azure Repos 개인 엑세스 토큰. 스코프: **`Custom defined`, \*\* `Code:` \*\* `Read & write`**_._

Azure Repos 문서 - [_개인 엑세스 토큰 사용_](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=Windows)\_

**Bitbucket Server/Data Center**:

**`BITBUCKET_USERNAME=`**, **`BITBUCKET_PASSWORD=`** - Bitbucket Server 사용자 이름 및 비밀번호 또는 Bitbucket Server 개인 엑세스 토큰. 스코프: **`Repository admin`**.

Bitbucket Server 문서 - [_개인 엑세스 토큰_](https://confluence.atlassian.com/bitbucketserver/http-access-tokens-939515499.html)\\
# Snyk SCM 통합

Snyk을 귀하의 Git 저장소와 통합하여 귀하의 **Projects** 목록에 추가되는 모든 [Snyk 프로젝트](https://docs.snyk.io/introducing-snyk/introduction-to-snyk-projects)를 빠르고 쉽게 볼 수 있습니다.

이 페이지에는 다음과 같은 Snyk SCM 통합 측면에 대한 정보가 포함되어 있습니다:

* [그룹 수준에서 사용 가능한 Snyk AppRisk SCM 통합](./#group-level-snyk-apprisk-scm-integrations)
* [조직 수준에서 사용 가능한 Snyk SCM 통합](./#organization-level-snyk-scm-integrations)
* [SCM 통합을 위한 작업 영역](./#workspaces-for-scm-integrations)
* [배포 순서 권고 사항](./#deployment-order-recommendations)
* [일반 사용자 권한 및 접근 범위 요구 사항](./#user-permissions-and-access-scope-requirements)
* [Snyk Broker용 통합 SCM 토큰](./#integrated-scm-tokens-for-snyk-broker)

## 개요

Snyk 계정에 새로운 통합을 추가하려면 먼저 어느 수준 타입에서 통합을 설치할 것인지 결정해야 합니다. Snyk은 Snyk 계층의 다른 수준에서 다른 SCM 통합을 지원합니다:

* [그룹 수준](./#group-level-snyk-apprisk-integrations) - 그룹 수준 통합은 Snyk AppRisk Essentials 또는 Snyk AppRisk Pro만 지원합니다.
* [조직 수준](./#organization-level-snyk-integrations) - 조직 수준 통합은 다른 Snyk 제품을 지원합니다. Snyk은 조직 수준에서 Snyk AppRisk Essentials 또는 Snyk AppRisk Pro 통합을 지원하지 않습니다.

{% hint style="warning" %}
조직 및 그룹 레벨에서 GitHub 통합을 추가하고 SAML SSO를 구성한 경우 GitHub PAT를 승인해야 합니다. 자세한 내용은 [개인 액세스 토큰을 승인하고 SSO 활성화하는 방법](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-enterprise#how-to-authorize-your-personal-access-token-and-enable-sso) 페이지를 참조하십시오.
{% endhint %}

## 그룹 수준 - Snyk AppRisk SCM 통합

그룹 수준에서 **Integrations Hub**을 통해 Snyk AppRisk 통합을 설정하고 사용자 정의할 수 있습니다. 다음 SCMs를 이용할 수 있습니다:

* [GitHub](github-enterprise.md#github-setup-guide-for-snyk-apprisk)
* [GitLab](gitlab.md#gitlab-setup-guide)
* [Azure DevOps](azure-repositories-tfs.md#azure-devops-setup-guide)
* [BitBucket](bitbucket-cloud.md#bitbucket-setup-guide)

Snyk AppRisk 그룹 수준의 SCM 통합은 해당 고객의 모든 응용 프로그램 자산에 대한 광범위한 가시성을 제공하며 추가 응용 프로그램 컨텍스트 및/또는 메타데이터(예: 개발자 정보, 커밋 등)를 가져옵니다.

{% hint style="info" %}
SCM 인스턴스에 공개적으로 액세스할 수 없는 경우 Snyk Broker를 사용하여 연결해야 합니다. 자세한 내용은 [Snyk Broker - AppRisk](../../enterprise-setup/snyk-broker/snyk-broker-apprisk.md) 페이지를 참조하십시오.
{% endhint %}

그룹 수준의 Integrations 페이지에는 모든 활성 통합이 표시되며 기존 Snyk 조직에서 자동으로 동기화된 데이터 및 통합 허브에 액세스할 수 있습니다.

{% hint style="warning" %}
SCM 통합은 저장소 검색을 증분 방식으로 사용합니다. 즉, 동기화가 시작되면 저장소의 마지막 업데이트 시간을 확인하고 그 이후 수정된 저장소만 전송합니다.

통합에 사용되는 사용자 또는 PAT의 범위에 변경 사항이 있는 경우 범위 내에 새롭게 추가된 저장소는 증분 수집을 유발할 변경 또는 통합 구성 변경 후에만 식별됩니다.
{% endhint %}

다음 지원되는 Snyk 데이터가 자동으로 동기화됩니다:

* Snyk Open Source
* Snyk Code
* Snyk IaC
* Snyk Container

각 연결된 통합은 다음을 가능케 합니다:

* 데이터 동기화 일시 중지
* 통합 프로필 및 설정 수정
* 통합 삭제
* 통합이 마지막으로 동기화된 시간과 다음 동기화 예정 시간 확인

### 와일드카드 SCM 통합

와일드카드 통합을 사용하면 특수 문자를 사용하여 동시에 여러 SCM 조직을 감지하고 통합할 수 있습니다.

{% hint style="info" %}
와일드카드 통합은 GitHub 및 Azure DevOps 통합에 적용되며 [Snyk Broker](../../enterprise-setup/snyk-broker/snyk-broker-apprisk.md)을 사용하여 설정하는 경우 지원됩니다.
{% endhint %}

통합 설정 시 와일드카드를 사용할 수 있습니다:

* **통합 허브** 메뉴를 엽니다.
* **SCM** 태그를 선택하고 GitHub 또는 Azure DevOps를 검색합니다.
* **추가** 버튼을 클릭합니다.
* **Organizations** 필드에 `*` 기호를 사용하여 조직 세부 정보를 추가합니다. 예를 들어 `*snyk` 을사용하면 이름에 Snyk이포함된 모든 SCM 조직이 통합됩니다.
* 와일드카드와 일치하는 모든 조직이 추가됩니다.

{% hint style="info" %}
와일드카드 `*` 기호는 매번 리포지토리를 다시 스캔할 때마다 적용되는 명령어로 간주됩니다.
{% endhint %}

### Snyk AppRisk 통합 생태계

다음 표를 참조하여 Snyk AppRisk의 모든 통합의 가용성 및 호환성을 확인할 수 있습니다. 통합은 유형별로 분류되며 이름별로 나열되며 Snyk AppRisk Essentials 및 Snyk AppRisk Pro에 대해 사용 가능 여부가 표시됩니다.

| 통합 유형             | 통합 이름                         | Snyk AppRisk Essentials | Snyk AppRisk Pro |
| ----------------- | ----------------------------- | ----------------------- | ---------------- |
| SCM               | GitHub                        | ✅                       | ✅                |
|                   | BitBucket                     | ✅                       | ✅                |
|                   | GitLab                        | ✅                       | ✅                |
|                   | Azure DevOps                  | ✅                       | ✅                |
| Dev 포털 및 서비스 카탈로그 | Backstage 카탈로그                | ✅                       | ✅                |
|                   | ServiceNow CMDB               | ✅                       | ✅                |
|                   | Atlassian Compass             | ✅                       | ✅                |
|                   | Harness                       | ✅                       | ✅                |
|                   | OpsLevel                      | ✅                       | ✅                |
|                   | Datadog Org Context(서비스 카탈로그) | ✅                       | ✅                |
| 리스크 관리 협력         | Jira                          | ✅                       | ✅                |
|                   | Slack                         | ✅                       | ✅                |
|                   | 이메일                           | ✅                       | ✅                |
| AST               | NightFall                     | ❌                       | ✅                |
|                   | GitGuardian                   | ❌                       | ✅                |
| 런타임 보안 및 가시성      | Snyk 런타임 센서                   | ❌                       | ✅                |
|                   | Sysdig                        | ❌                       | ✅                |
|                   | Dynatrace                     | ❌                       | ✅                |

### 통합 허브 사용

통합 허브 페이지를 사용하여 통합을 도입하고 SCM 도구에서 Snyk AppRisk를 데이터로 채울 수 있습니다.

통합을 설정하는 방법에 대한 단계별 안내는 [Snyk 웹 UI](../../getting-started/snyk-web-ui.md#manage-your-integrations) 페이지를 참조하십시오.

통합이 유효성을 검사하면 통합 페이지에 카드가 표시되어 연결을 활성화하거나 비활성화하거나 설정을 편집하거나 구성에서 연결을 제거할 수 있습니다.

{% hint style="info" %}
초기 구성 후 권한 및 범위를 수정하는 경우 리포지토리 내에서 가져오기를 초기화하거나 변경을 구현하는 것이 중요합니다. 이 작업을 통해 Snyk 업데이트를 인식하고 효과적으로 통합할 수 있습니다.
{% endhint %}

### Snyk Broker 사용

SCM 인스턴스에 공개적으로 액세스할 수 없는 경우 Snyk Broker가 필요합니다. Snyk Broker를 설치하고 구성하는 방법은 Docker 또는 Helm을 사용할 수 있습니다. Snyk Broker에 대한 자세한 내용은 [Snyk Broker - AppRisk](../../enterprise-setup/snyk-broker/snyk-broker-apprisk.md)를 포함한 Snyk Broker 문서를 참조하십시오.

{% hint style="warning" %}
명령어를 실행하기 전에 Snyk Broker 배포 환경에서 Snyk AppRisk 플래그를 활성화해야 합니다.
{% endhint %}

통합에 액세스할 수 있는 허용 목록이 포함된 모든 최신 `.json` 파일을 [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk)에서 찾을 수 있습니다.

## 조직 수준 - Snyk SCM 통합

Snyk은 GitHub, GitLab, Bitbucket, Azure Repos와 같은 다양한 소스 제어 관리 시스템과의 조직 수준에서 원활한 통합을 제공합니다. 이러한 통합을 통해 취약점을 자동으로 스캔하고 실행 가능한 통찰을 제공하여 귀하의 저장소의 보안을 강화할 수 있습니다. 이 수준에서 통합을 실행하여 조직 내의 모든 프로젝트에 대한 포괄적인 보호를 보장할 수 있습니다.

Snyk 소스 제어 관리자 (SCM) 통합을 통해 다음을 수행할 수 있습니다:

* 통합된 모든 저장소에서 지속적으로 보안 스캔 수행
* 오픈 소스 구성 요소의 취약점 감지
* 자동 수정 제공

Snyk은 다음 SCMs와 통합하여 귀하의 코드에서의 문제와 취약점을 추적, 모니터링 및 수정할 수 있도록 지원합니다:

* [GitHub Cloud App](github-cloud-app.md)
* [GitHub Enterprise](github-enterprise.md)
* [GitHub](github.md)
* [GitHub 읽기 전용 프로젝트](github-read-only-projects.md)
* [GitLab](gitlab.md)
* [Bitbucket Cloud](bitbucket-cloud.md)
* [Bitbucket Cloud App](bitbucket-cloud-app.md)
* [Bitbucket 데이터 센터/서버](bitbucket-data-center-server.md)
* [Azure Repositories (TFS)](azure-repositories-tfs.md)

## SCM 통합을 위한 작업 영역

{% hint style="info" %}
이 기능은 GitHub, GitHub Enterprise, GitHub Cloud App, GitLab, Bitbucket Server, Bitbucket Cloud App, Bitbucket Cloud (Legacy) 및 Azure Repos (TFS) 통합에 사용할 수 있습니다.
{% endhint %}

Workspaces 기능을 사용하면 Snyk이 구성된 SCM 통합을 통해 저장소 내용의 임시 스냅샷과 모든 커밋 메타데이터를 수용할 수 있습니다.

이 기능에 대한 자세한 내용 및 활성화 단계에 대한 정보는 [SCM 통합을 위한 Workspaces](introduction-to-git-repository-integrations/workspaces-for-scm-integrations.md)를 참조하십시오.

## 배포 순서 권고 사항

모든 SCM 통합 기능을 동시에 구현하려고 시도하면 소프트웨어 개발 수명 주기(SDLC)에서 마찰을 초래하여 개발자 경험을 저하시킬 수 있습니다.

조직 전체에 Snyk을 원활하게 배포하려면 배포 단계, 구성 단계 및 각 단계의 원하는 결과로 구성된 Snyk의 제안된 배포 시간표를 제공합니다.

세부 단계에 대한 자세한 내용은 [SCM 통합을 위한 배포 권고사항](introduction-to-git-repository-integrations/deployment-recommendations-for-scm-integrations.md)을 참조하십시오.

## 사용자 권한 및 액세스 범위 요구 사항

Snyk SCM 통합은 연결 방법에 따라 다른 권한 요구 사항을 필요로 할 수 있습니다.

세부적인 권한 요구 사항은 다음을 참조하십시오:

* [GitHub 및 GitHub Enterprise](./#github-and-github-enterprise-permissions-requirements)
* [GitHub Cloud App](./#github-cloud-app-permission-requirements)
* [GitHub Server App](./#github-server-app-permission-requirements)
* [GitHub Server App for Universal Broker](./#github-server-app-for-universal-broker-permission-requirements)
* [GitLab](./#gitlab-permission-requirements)
* [Bitbucket](./#bitbucket-permission-requirements)
* [Azure Repositories (TFS)](./#azure-repositories-tfs-permission-requirements)

### GitHub와 GitHub Enterprise 권한 요구사항

{% hint style="info" %}
브로커드 통합에서 토큰 권한에 대한 정보는 [GitHub - 브로커 설치 및 구성 사전 조건 및 단계](https://docs.snyk.io/enterprise-configuration/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker) 및 [Snyk 브로커를 위한 통합 SCM 토큰](./#integrated-scm-tokens-for-snyk-broker)을 참조하십시오.
{% endhint %}

Snyk GitHub Enterprise 통합은 주로 GitHub 서비스 계정인 단일 사용자에 바인딩됩니다. 통합의 액세스 수준은 GitHub에서 사용자의 권한과 해당 사용자 계정의 개인 액세스 토큰(PAT)에 정의된 액세스의 조합에 의해 결정됩니다. 사용자의 GitHub 계정보다 더 많은 권한으로 PAT를 정의하면 해당 권한을 사용할 수 없을 수 있습니다.

다음 표에는 GitHub와 GitHub Enterprise에서 개인 액세스 토큰(PAT) 및 Snyk이 감시하는 저장소에서 필요한 작업 수행을 위한 액세스 범위에 대한 세부 정보가 나와 있습니다. 이러한 작업은 자주 manifest 파일을 읽거나 수정하거나 수정 PR을 열거나 업데이트하는 등 소스 코드 저장소에서 요구되는 작업을 수행하기 위해 필요합니다. GitHub 사용자 지정 역할은 지원되지 않습니다.

| 동작 및 목적                                                                                                         | PAT 범위                                                       | 저장소 범위    |
| --------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ | --------- |
| <p><strong>일일/주간 테스트:</strong><br>개인 저장소에서 manifest 파일 읽기</p>                                                   | `repo (모두)`                                                  | ≥ `read`  |
| <p><strong>수동 고치기 PR:</strong><br>감시 중인 저장소에 고치기 PR 생성</p>                                                      | `repo (모두)`                                                  |           |
| <p><strong>자동 고치기 및 업그레이드 PR:</strong><br>감시 중인 저장소에 고치기 또는 업그레이드 PR 생성</p>                                     | `repo (모두)`                                                  | ≥ `write` |
| <p><strong>풀 리퀘스트 대한 Snyk 테스트:</strong><br>새 PR이 만들어지거나 기존 PR이 업데이트될 때 PR 상태 확인 전달</p>                          | `repo (모두)`                                                  | ≥ `write` |
| <p><strong>풀 리퀘스트에 대한 Snyk 테스트 초기 구성:</strong><br>수입된 저장소에 SCM 웹훅 추가</p>                                        | `admin:repo_hooks (read & write)`                            | `admin`   |
| <p><strong>새 프로젝트를 Snyk에 가져오기:</strong><br>GitHub org에 있는 모든 사용 가능한 저장소 목록을 <strong>프로젝트 추가</strong> 화면에 제시</p> | <p><code>admin:read:org</code><br><code>repo (모두)</code></p> |           |

세부 내용이 있는 PAT는 추가 저장소 액세스 범위가 필요합니다:

* `Administration: Read-only`
* `Commit Status: Read and write`
* `Content: Read and write`
* `Metadata: Read-only`
* `Pull requests: Read and write`
* `Webhooks: Read and write`
* `Members access: Read-only (Organization access scope)`

PAT의 `Administration: Read-only` 권한은 Snyk이 새 프로젝트를 가져오기 위해 필요한 사용자의 액세스 가능한 GitHub 조직을 식별하고 나열하는 데 중요합니다.

Snyk은 병합이 발생할 것을 [GitHub Enterprise](github-enterprise.md)에 알리기 위해 PR을 사용합니다. 이를 위해 콘텐츠가 분기로 푸시되어야 하며 이렇게 되면 `content: write` 범위가 필요합니다. 그런 다음 따로 호출을 통해 고칠 PR을 만들고 이는 `pull request: write` 범위가 필요합니다. 그런 후 GitHub Enterprise에게 PR을 만들도록 지시하여 변경된 브랜치를 기본 브랜치로 병합합니다.

Snyk은 SCM 웹훅을 사용하여 다음과 같은 작업을 수행합니다:

* PR이 만들어지거나 업데이트되거나 트리거되거나 병합될 때 PR 상태를 추적합니다.
* PR 체크 트리거를 위해 푸시 이벤트를 전송합니다.

### GitHub 클라우드 앱 권한 요구사항

[Snyk GitHub 클라우드 앱](github-cloud-app.md) 통합은 역할 기반 액세스 제어를 사용합니다. 따라서 액세스 제어는 개별 사용자나 사용자 역할에 의존하는 것이 아니라 앱 엔티티에 결합됩니다.

GitHub 클라우드 앱 통합을 설정하려면 다음이 필요합니다:

* Snyk 조직 관리자.
* GitHub 조직 관리자.
* GitHub 리포지터리 관리자 (GitHub UI를 통해 설치하는 경우).

{% hint style="info" %}
GitHub 측면에서 일부 권한이 선택적일 수 있지만, 이러한 권한은 Snyk 기능 지원을 위해 필요합니다. 이러한 권한은 Snyk 조직 하에 등록된 앱이기 때문에 **개별적으로 사용자의 요구에 맞게 사용자 정의할 수 없습니다.**
{% endhint %}

다음 표는 필요한 GitHub 앱 권한 및 범위를 명시합니다:

| 동작 및 범위                                                                      | 범위               | 수준    | 권한      |
| ---------------------------------------------------------------------------- | ---------------- | ----- | ------- |
| GitHub 사용자가 GitHub 조직에서 관리자 역할을 가지고 있는지 여부 확인하여 앱 설치 재사용을 관리자 사용자에게만 제한하는 동작 | 멤버               | 조직    | 읽기      |
| 저장소 검색 및 저장소 메타데이터 액세스                                                       | 메타데이터            | 리포지터리 | 읽기      |
| GitHub Checks 탭과 상호 작용                                                       | Checks           | 리포지터리 | 읽기 및 쓰기 |
| 커밋 및 브랜치 생성                                                                  | Contents         | 리포지터리 | 읽기 및 쓰기 |
| PR 체크 결과를 커밋 상태로 전송                                                          | Commit status    | 리포지터리 | 읽기 및 쓰기 |
| PR 세부 정보 가져오기, 관련 코멘트 게시 (다음 세대 PR 경험)                                       | Pull request     | 리포지터리 | 읽기 및 쓰기 |
| PR 체크를 트리거하는 웹훅 관리                                                           | Repository hooks | 리포지터리 | 읽기 및 쓰기 |

### GitHub Server App 권한 요구사항

GitHub Server App를 설정하려면 다음이 필요합니다:

* Snyk 조직 관리자.
* GitHub 조직 관리자.

공개 또는 비공개 GitHub 리포지터리가 필요합니다.

### Universal Broker를 위한 GitHub Server App 권한 요구사항

Universal Broker를 위한 GitHub Server App를 설정하려면 다음이 필요합니다:

* Snyk 조직 관리자.
* GitHub 조직 관리자.

자체 호스팅된 GitHub 인스턴스가 필요합니다.

### GitLab 권한 요구사항

[Snyk GitLab 통합](gitlab.md#gitlab-access-tokens)에서는 GitLab 계정 티어에 따라 개인 액세스 토큰(PAT) 또는 그룹 액세스 토큰(GAT)을 사용합니다.

Snyk GitLab 통합을 설정하려면 다음이 필요합니다:

* Snyk 그룹이나 조직 관리자.
* GitLab 소유자 또는 Maintainer.

개인 GitLab 프로젝트를 관리하기 위해 PAT가 사용되며 `api` 범위가 필요합니다. 모든 GitLab 리포지터리를 보여주기 위해 AppRisk는 PAT 생성을 하는 사용자가 최소한 해당 GitLab 그룹에 속한 멤버이어야 합니다.

여러 GitLab 프로젝트를 관리하기 위해 GAT가 사용되며 드롭다운 메뉴에서 선택한 `api` 범위와 Maintainer 역할이 필요합니다. GAT를 생성하려면 GitLab 프리미엄 또는 얼티밋 계정 티어 소유자여야 합니다.

### Bitbucket 권한 요구사항

Snyk Bitbucket 통합은 다음과 같은 다른 액세스 제어 메커니즘을 사용하여 Snyk와 연결합니다:

* [Snyk Bitbucket Cloud](./#bitbucket-cloud-and-bitbucket-data-center-server-scopes)는 [앱 비밀번호](bitbucket-cloud.md#how-to-set-up-the-snyk-bitbucket-cloud-integration)의 생성이 필요합니다.
* [Snyk Bitbucket Cloud App](./#bitbucket-cloud-app-scopes)은 [Bitbucket 워크스페이스 인증](bitbucket-cloud-app.md#setting-up-a-bitbucket-cloud-app) 및 관련 권한이 필요합니다.
* [Snyk Bitbucket Data Center/Server](./#bitbucket-cloud-and-bitbucket-data-center-server-scopes)는 [전용 사용자 이름 및 비밀번호](bitbucket-data-center-server.md#how-to-set-up-a-bitbucket-dc-server-integration) 또는 개인 액세스 토큰(PAT)이 필요합니다.

{% hint style="warning" %}
Snyk Bitbucket 통합을 설정하려면 Bitbucket 워크스페이스 관리자이어야 합니다.
{% endhint %}

#### Bitbucket Cloud 및 Bitbucket Data Center/Server 범위

다음 표에는 Bitbucket Cloud 및 Bitbucket Data Center/Server에서 필요한 권한 범위에 대한 자세한 정보가 나와 있습니다:

| 동작 및 목적                                                                                                            | 앱 비밀번호 요구사항                                                                                                                           | Bitbucket 권한 |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------- | ------------ |
| <p><strong>일일/주간 테스트:</strong><br>개인 저장소에서 manifest 파일 읽고 private repos에서 파일 수정</p>                                | **저장소:** `읽기`                                                                                                                         | ≥ `쓰기`       |
| <p><strong>수동 고치기 PR (사용자 트리거):</strong><br>저장소에 고치기 PR 생성</p>                                                     | <p><strong>저장소</strong>: <code>읽기, 쓰기</code><br><strong>풀 리퀘스트</strong>: <code>읽기, 쓰기</code></p>                                      |              |
| <p><strong>자동 고치기 및 업그레이드 PR:</strong><br>저장소에서 고치기/업그레이드 PR 생성</p>                                                | <p><strong>저장소</strong>: <code>읽기, 쓰기</code><br><strong>풀 리퀘스트</strong>: <code>읽기, 쓰기</code></p>                                      | ≥ `쓰기`       |
| <p><strong>풀 리퀘스트에 대한 Snyk 테스트:</strong><br>새 PR이 만들어지거나 PR이 업데이트될 때 PR 상태 확인 전달</p>                               | <p><strong>저장소</strong>: <code>읽기, 쓰기</code><br><strong>풀 리퀘스트</strong>: <code>읽기, 쓰기</code></p>                                      | ≥ `쓰기`       |
| <p><strong>풀 리퀘스트에 대한 Snyk 테스트 (초기 설정):</strong><br>수입된 저장소에 SCM 웹훅 추가</p>                                         | **웹훅**: `읽기, 쓰기`                                                                                                                      | `관리자`        |
| <p><strong>Snyk에 새 프로젝트 가져오기:</strong><br>Bitbucket 인스턴스에 있는 모든 사용 가능한 저장소 목록을 <strong>프로젝트 추가</strong> 화면에 제시</p> | <p><strong>계정</strong>: <code>읽기</code><br><strong>워크스페이스 멤버십</strong>: <code>읽기</code><br><strong>프로젝트</strong>: <code>읽기</code></p> |              |

Snyk은 Bitbucket에서 SCM 웹훅을 사용하여 다음 작업을 수행합니다:

* PR이 만들어지거나 업데이트되거나 트리거되거나 병합될 때 PR 상태를 추적합니다.
* PR 체크를 트리거하기 위해 푸시 이벤트를 전송합니다.

#### Bitbucket Cloud App 범위

다음 표는 **Bitbucket Cloud App**에 필요한 권한을 상세히 설명합니다:

| 동작                      | 목적                                                                              | 필요 범위                                           |
| ----------------------- | ------------------------------------------------------------------------------- | ----------------------------------------------- |
| 일일/주간 테스트               | 개인 저장소에서 manifest 파일을 읽기 위해 사용됩니다.                                              | 귀하의 리포지터리 및 해당 pull request를 읽고 수정하는 권한이 필요합니다. |
| PR에 대한 Snyk 테스트         | 새 PR이 생성되거나 기존 PR이 업데이트될 때 PR 상태 확인을 전송하기 위해 사용됩니다.                             | 귀하의 리포지터리 및 해당 pull request를 읽고 수정하는 권한이 필요합니다. |
| 수정 및 업그레이드 PR 열기        | 감시 중인 저장소에서 고치기 PR 생성을 위해 사용됩니다.                                                | 귀하의 리포지터리 및 해당 pull request를 읽고 수정하는 권한이 필요합니다. |
| PR에 대한 Snyk 테스트 - 초기 설정 | 수입된 리포지터리에 Snyk 웹훅을 추가하여 PR이 작성되거나 업데이트될 때 Snyk에 알릴 수 있게 되며, Snyk 스캔을 트리거할 수 있음 | 귀하의 리포지터리 웹훅을 읽고 수정하는 권한이 필요합니다                 |

### Azure Repositories (TFS) 권한 요구 사항

[Snyk Azure Repositories (TFS) 통합](azure-repositories-tfs.md)은 Azure DevOps 개인 액세스 토큰(PAT)을 사용합니다. 이 토큰은 Snyk이 Azure 리포지토리에 접근하는 데 필요한 특정 권한으로 구성됩니다.

Snyk Azure Repositories (TFS) 통합을 설정하려면 다음 조건을 충족해야 합니다:

* [Snyk 조직 관리자](../../snyk-admin/user-roles/pre-defined-roles.md)여야 합니다.
* Azure에서 [프로젝트 관리자 그룹](https://learn.microsoft.com/en-us/azure/devops/organizations/security/change-project-level-permissions?view=azure-devops\&tabs=preview-page)의 구성원이어야 합니다. 이렇게 하면 PAT에 웹후크를 활성화하는 데 필요한 `edit subscriptions permissions`가 부여됩니다.

Azure에서 PAT는 Snyk 접근을 위한 다음 권한을 요구합니다:

* **만료**: 사용자 정의. Snyk은 토큰 만료 날짜를 먼 미래로 설정하는 것을 권장합니다.
* **범위**: 사용자 정의. **Code** 범위에 대해 `읽기 및 쓰기` 권한이 필요합니다.

## Snyk Broker를 위한 통합 SCM 토큰

통합 SCM 토큰은 [Broker 클라이언트 설정](../../enterprise-setup/snyk-broker/#integrations-with-snyk-broker)에 필요합니다. 이 토큰은 `-e <SCM>_TOKEN` 매개변수에서 사용됩니다. 예를 들어, `-e GITHUB_TOKEN=xxx…`와 같이 SCM에 대한 접근을 가능하게 합니다. 이 토큰은 Broker와 Snyk Code의 운영에 필요한 특정 권한을 충족합니다.

다음 SCM 통합에 대해 통합 SCM 토큰을 생성할 수 있습니다:

* [GitHub 및 GitHub Enterprise](./#github-and-github-enterprise-scm-token)
* [GitLab](./#gitlab-scm-token)
* [Azure Repositories (TFS)](./#azure-repositories-tfs-scm-token)
* [Bitbucket Server/Data Center](./#bitbucket-server-data-center-scm-token)

### GitHub 및 GitHub Enterprise SCM 토큰

**형식**: `GITHUB_TOKEN=` - GitHub 개인 액세스 토큰.\
**범위**: `repo`, `read:org` 및 `admin:repo_hook`.

### GitLab SCM 토큰

**형식**: `GITLAB_TOKEN=` - GitLab 개인 액세스 토큰.\
**범위**: `api`.

GitLab 계정에 `Maintainer` 권한이 필요합니다.

### Azure Repositories (TFS) SCM 토큰

**형식**: `AZURE_REPOS_TOKEN=` - Azure Repos 개인 액세스 토큰.\
**범위**: `Custom defined`, `Code:` `Read & write`_._

### Bitbucket Server/Data Center SCM 토큰

**형식**: `BITBUCKET_USERNAME=`, `BITBUCKET_PASSWORD=` – Bitbucket Server 사용자 이름과 비밀번호 또는 Bitbucket Server 개인 액세스 토큰.\
**범위**: `Repository admin`.

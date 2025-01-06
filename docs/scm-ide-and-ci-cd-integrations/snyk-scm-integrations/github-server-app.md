# GitHub 서버 앱

새로운 통합을 Snyk 계정에 추가하려면 먼저 설치하려는 통합의 레벨 유형을 결정해야 합니다.

* [그룹 레벨](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-cloud-app#group-level-snyk-apprisk-integrations) - Snyk 앱리스크 에센셜 또는 Snyk 앱리스크 프로용으로 사용할 수 있는 Snyk 애플리케이션에 통합을 추가합니다. Snyk 앱리스크를 위해 통합을 설정하려면 그룹 레벨의 통합 메뉴를 사용하세요.
* [조직 레벨](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github-cloud-app#organization-level-snyk-integrations) - Snyk 제품을 제외하고 모든 Snyk 제품에 사용할 수 있는 Snyk 애플리케이션에 통합을 추가합니다.

## 조직 레벨 - Snyk 통합

{% hint style="info" %}
**릴리스 상태**

GitHub 서버 앱은 이른 액세스 단계에 있으며 엔터프라이즈 요금제에서만 사용할 수 있습니다.

이 기능은 GitHub의 온프레미스 인스턴스 및 Snyk의 [유니버설 브로커](../../enterprise-setup/snyk-broker/universal-broker/)를 지원합니다.
{% endhint %}

이 페이지에서 다음 기능을 다룹니다:

* [GitHub 서버 앱](github-server-app.md#prerequisites-for-github-server-app).
* [기존 GitHub 엔터프라이즈 통합에서 마이그레이션](github-server-app.md#migrate-from-an-existing-github-enterprise-integration).
* [유니버설 브로커용 GitHub 서버 앱 설정](github-server-app.md#set-up-the-github-server-app-for-universal-broker).

### GitHub 서버 앱 사전 준비 사항

* 온프레미스 GitHub 인스턴스.
* Snyk 조직 관리자 역할.
* GitHub 조직 관리자 역할.
* 공개 또는 비공개 GitHub 리포지토리.

### GitHub 서버 앱 이점<a href="#github-server-app-benefits" id="github-server-app-benefits"></a>

Snyk GitHub 서버 앱은 Snyk GitHub 엔터프라이즈 통합과 비교하여 많은 기능을 개선합니다. 이에는 롤 기반의 세분화된 액세스 제어, 증가한 API 속도 제한 및 확장 및 향상된 개발 경험을 위한 진입점 생성이 포함됩니다.

* **RBAC(롤 기반 액세스 제어) 규칙 준수**: GitHub 서버 앱을 사용하면 액세스 제어 메커니즘이 개별 사용자 계정에서 분리됩니다. 대신, 이는 앱 엔티티 자체에 연결됩니다. 이 분리로 인해 액세스 제어는 개별 사용자 계정에 묶이는 대신 애플리케이션 레벨에서 처리되므로 RBAC 정책의 관리 및 강제가 더 잘 이루어질 수 있습니다.
* **세분화된 액세스 제어**: GitHub 서버 앱은 리포지토리 수준에서 액세스 권한을 세분화된 권한 제어가 가능하게 해 줍니다.
* **증가한 API 속도 제한**: GitHub 서버 앱은 더 높은 속도 제한을 제공하여 Snyk가 더 많은 API 요청을 수행할 수 있습니다. 이 증가한 한도는 큰 규모의 사용 사례를 처리하는 데 도움이 될 것입니다. 예를 들어, 많은 프로젝트를 가진 모노 리포지토리, 많은 리포지토리를 가진 GitHub 조직 등을 다룰 때 유용할 것입니다.
* **확장 및 향상된 개발 경험을 위한 활성화 요소:**
  * 풀 리퀘스트 확인: GitHub의 Checks 탭 경험은 GitHub 클라우드 앱을 통해서만 접근이 가능하며 잠재적인 향후 PR 확인 워크플로우 개선의 일부로 SCM 네이티브 경험을 제공합니다.
  * 수정 및 업그레이드 풀 리퀘스트: Snyk에서 시작된 풀 리퀘스트는 서비스 계정이 아닌 GitHub 앱을 통해 직접 수행됩니다.

### GitHub 서버 앱 설정

{% hint style="warning" %}
GitHub 서버 앱을 설정할 때 다음 시나리오 중 하나만 구현할 수 있습니다:

* 하나의 GitHub 조직이 하나의 Snyk 조직과 연결되어 있음
* 하나의 GitHub 조직이 여러 Snyk 조직과 연결되어 있음
{% endhint %}

Snyk의 UI에서 통합 페이지로 이동하고 **GitHub 서버 앱** 타일을 선택합니다.

<figure><img src="../../.gitbook/assets/image (594).png" alt=""><figcaption><p>Snyk UI에서 강조된 GitHub 서버 앱 타일</p></figcaption></figure>

타일을 클릭하면 모달이 열리며 GitHub 서버의 URL을 입력할 수 있는 페이지로 이동합니다. GitHub 서버 인스턴스의 URL을 입력하면 GitHub 인스턴스로 리디렉션이 되어 앱을 생성할 수 있습니다.

<figure><img src="../../.gitbook/assets/image (596).png" alt=""><figcaption><p>사용자의 GitHub 서버 URL을 입력하도록 하는 통합 모델</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (598).png" alt=""><figcaption><p>GitHub 인스턴스에서 앱 등록</p></figcaption></figure>

그런 다음 앱이 사용자를 대신해 작동하도록 허용해야 합니다. 이 정보는 사용자가 해당 GitHub 조직에 앱을 설치할 수 있는 권한을 확인하는 데 사용됩니다.

<figure><img src="../../.gitbook/assets/image (599).png" alt=""><figcaption><p>앱을 위한 사용자 권한 부여</p></figcaption></figure>

GitHub에서 설치 화면이 열리면 앱을 설치할 GitHub 조직을 선택할 수 있습니다.

<figure><img src="../../.gitbook/assets/image (602).png" alt=""><figcaption><p>앱을 설치할 GitHub 조직 선택</p></figcaption></figure>

GitHub 서버 앱이 GitHub 조직에 이미 설치되어 있는 경우, 다른 Snyk 조직을 위한 통합 프로세스 중에 동일한 GitHub 조직을 선택할 수 있습니다.

<figure><img src="../../.gitbook/assets/image (604).png" alt=""><figcaption><p>Snyk 조직에 다른 GitHub 조직 연결</p></figcaption></figure>

선택한 GitHub 조직의 모든 또는 선택한 수의 리포지토리에 앱을 설치할지 여부를 지정하고 **설치 및 권한 부여**를 클릭합니다.

<figure><img src="../../.gitbook/assets/image (608).png" alt=""><figcaption><p>GitHub 클라우드 앱을 설치하고 권한 부여 설정</p></figcaption></figure>

{% hint style="danger" %}
GitHub 서버 앱이 GitHub 조직에서 제거되면 Snyk의 액세스도 제거됩니다. 이런 경우 Snyk의 GitHub 서버 앱과 다시 통합하기 전에 지원 티켓을 발급해야 합니다.
{% endhint %}

### 기존 GitHub 엔터프라이즈 통합에서 마이그레이션

엔터프라이즈 요금제 고객은 [snyk-migrate-to-github-app](https://github.com/snyk-labs/snyk-migrate-to-github-app) 도구를 사용하여 Snyk 대상을 GitHub 서버 앱으로 마이그레이션할 수 있습니다. 이 도구는 [도구 리포지토리](https://github.com/snyk-labs/snyk-migrate-to-github-app)에서 사용할 수 있습니다.

### 유니버설 브로커용 GitHub 서버 앱 설정

유니버설 브로커의 설정 프로세스는 다음과 같습니다:

1. [GitHub 서버 인스턴스에 GitHub 앱 생성](github-server-app.md#create-a-github-app-for-universal-broker).
2. [API를 통해 생성된 엔터프라이즈 통합](github-server-app.md#create-the-github-server-app-scm-integration).

#### 유니버설 브로커용 GitHub 앱 생성

유니버설 브로커에서 GitHub 서버 앱을 사용하려면 GitHub 서버 인스턴스에 별도의 GitHub 앱을 생성해야 합니다. 이를 위해 Snyk 서비스를 위한 필요한 모든 권한을 사전으로 정의한 `GITHUB-SERVER-URL`을 사용할 수 있습니다.

```
/settings/apps/new?name=Snyk&description=Snyk%20helps%20you%20develop%20fast%20while%20staying%20secure%20by%20finding%20and%20automatically%20fixing%20security%20issues%20in%20your%20code%2C%20open%20source%20dependencies%2C%20containers%2C%20and%20infrastructure%20as%20code%20-%20all%20powered%20by%20Snyk%E2%80%99s%20security%20intelligence.&url=https%3A%2F%2Fgithub.com%2Fapps%2Fsnyk-io&public=false&webhook_active=true&webhook_url=%2Fapi%2Fhidden%2Fscm-apps%2Fapi%2Fgithub-app%2Fwebhook&checks=write&statuses=write&contents=write&metadata=read&repository_projects=write&pull_requests=write&repository_hooks=write&members=read&events[]=repository 
```

위 URL에서 다음을 교체합니다:

- ``: GitHub 서버 인스턴스의 기본 URL로 교체합니다.
- ``: Snyk 계정의 테넌시로 교체합니다. 이 값은 URL로 인코딩되어야 하며, 아래와 같이 일반적인 값을 사용할 수 있습니다:
  - Snyk US: https%3A%2F%2Fapp.snyk.io
  - Snyk EU: https%3A%2F%2Fapp.eu.snyk.io
  - Snyk AU: https%3A%2F%2Fapp.au.snyk.io
  - Snyk AWS US: https%3A%2F%2Fapp.us.snyk.io

값을 교체한 후 해당 URL로 이동합니다. 이 URL로 이동하면 GitHub 서버 인스턴스에서 모든 필수 정보가 미리 채워진 앱 생성 화면으로 이동합니다. 그런 다음 페이지 하단으로 스크롤하여 **GitHub 앱 생성** 버튼을 클릭하여 **모든 계정**라디오 버튼이 선택되었는지 확인하십시오.

<figure><img src="../../.gitbook/assets/image (611).png" alt=""><figcaption><p>GitHub 앱 계정 설정 선택</p></figcaption></figure>

{% hint style="warning" %}
GitHub 서버 앱을 생성하면 `ClientId`와 `AppId`가 표시됩니다. 이는 앱의 자격증명으로 안전하게 보관되어야 합니다.
{% endhint %}

GitHub 서버 앱을 생성한 후에는 개인 키를 만들도록 하는 배너가 나타납니다. 배너를 클릭하여 앱을 위한 개인 키를 생성하십시오.

<figure><img src="../../.gitbook/assets/image (618).png" alt=""><figcaption><p>개인 키 생성 링크가 있는 등록 완료 메시지</p></figcaption></figure>

개인 키를 생성하면 `.pem` 파일이 다운로드됩니다. 이 파일도 비밀로 취급되어 안전하게 보관되어야 합니다.

이제 GitHub 서버 앱은 Snyk 조직의 리포지토리에 대해 설치할 준비가 되었습니다. 동일한 페이지 상단으로 스크롤하여 왼쪽 패널에서 **앱 설치**를 선택하고 새로 생성된 앱을 선택한 후 우측에 있는 **설치** 버튼을 클릭하세요.

<figure><img src="../../.gitbook/assets/image (624).png" alt=""><figcaption><p>리포지토리에 GitHub 앱 설치</p></figcaption></figure>

<figure><img src="../../.gitbook/assets/image (626).png" alt=""><figcaption><p>선택한 리포지토리에 GitHub 앱 설치</p></figcaption></figure>

GitHub 조직에서 앱을 설치할 위치를 선택하세요. GitHub 조직의 특정 리포지토리 또는 모든 리포지토리에 설치할 수 있습니다.

{% hint style="info" %}
GitHub 조직의 일부 리포지토리에 앱을 설치하는 경우, 해당 리포지토리에서만 작동합니다. 차후 추가 리포지토리에 추가하려면 나중에 이 화면으로 돌아와 설치 위치를 편집할 수 있습니다.
{% endhint %}

앱 설치 시 `InstallationID`를 받게 됩니다. 이는 페이지 URL의 마지막 숫자입니다. 이 번호를 메모해두어 브로커 연결 설정 시 필요할 것입니다.

#### GitHub 서버 앱 SCM 통합 생성

GitHub 서버 앱을 사용하려면 Snyk에서 통합을 설정해야 합니다. GitHub 서버 앱을 위한 브로커 통합을 설정하려면 API를 사용해야 합니다.

통합을 설정하려면 다음이 필요합니다:

* Snyk 환경의 기본 API 주소
  * 대부분의 사용자에게는 `https://api.snyk.io`입니다.
* 설정하려는 Snyk 조직 ID
* 유효한 [Snyk API 토큰](https://docs.snyk.io/snyk-api/v1-api#authorization).

위 정보를 사용하여 API 호출을 수행하십시오. `{}` 안에 필수 정보를 교체하십시오.

```
curl --location 'https://{SNYK-BASE-API}/v1/org/{YOUR-ORG-ID}/integrations' \
--header 'Content-Type: application/json; charset=utf-8' \
--header 'Authorization: token ${REPLACE_WITH_API_TOKEN}' \
--data '{
    "type": "github-server-app",
    "broker": {
        "enabled": true
    }
}'
```

{% hint style="info" %}
위 명령어가 성공적으로 실행되면 `id` 값이 반환됩니다. 이는 Snyk에서 새로운 GitHub Server App 통합의 `integrationId`입니다. 이 값을 기록해 두세요. Universal Broker에서 필요할 것입니다.
{% endhint %}

Snyk의 통합 페이지를 방문하면 통합이 구성된 것을 확인할 수 있습니다.

<figure><img src="../../.gitbook/assets/image (630).png" alt=""><figcaption><p>성공적인 GitHub Server App 통합</p></figcaption></figure>

GitHub Server App의 설정 및 설치를 완료하면 브로커 연결 설정에 필요한 모든 자격 증명이 제공됩니다. GitHub Server App과 Broker를 연결하는 방법에 대한 자세한 내용은 Snyk의 [Universal Broker](https://docs.snyk.io/enterprise-setup/snyk-broker/universal-broker) 공식 문서를 참조하세요.
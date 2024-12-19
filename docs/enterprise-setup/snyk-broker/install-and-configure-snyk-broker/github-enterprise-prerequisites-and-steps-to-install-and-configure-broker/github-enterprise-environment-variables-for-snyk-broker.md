# GitHub Enterprise - Snyk 브로커를 위한 환경 변수

다음 환경 변수는 GitHub Enterprise용 브로커 클라이언트를 구성하는 데 필요합니다.

- `BROKER_TOKEN` - Snyk 브로커 토큰으로, Snyk Org 설정 뷰 (app.snyk.io)에서 얻을 수 있습니다.
- `GITHUB_TOKEN` - `repo`, `read:org`, `admin:repo_hook` 스코프가 모두 포함된 개인 액세스 토큰.
- `GITHUB` - GitHub Enterprise 배포의 호스트명:
  - 자체 호스트 또는 사용자 정의 도메인의 경우 GitHub Enterprise 도메인인 `your.ghe.domain.com`을 사용합니다.
  - 사용자 정의 도메인이 없는 GitHub Enterprise Cloud의 경우 `github.com.`을 사용합니다.
  - 도메인 뒤에 org 또는 경로를 포함하지 마세요.
- `GITHUB_API` - GitHub Enterprise 배포의 API 엔드포인트. `http` 또는 `https`를 사용하지 마세요.
  - 자체 호스트된 경우 `your.ghe.domain.com/api/v3.`을 사용합니다.
  - 사용자 정의 도메인이 없는 GitHub Enterprise Cloud의 경우 `api.github.com.`을 사용합니다.
- `GITHUB_GRAPHQL` - GitHub Enterprise 배포의 `graphql` 엔드포인트. `http` 또는 `https`를 사용하지 마세요.
  - 자체 호스트된 경우 `your.ghe.domain.com/api.`를 사용합니다.
  - 사용자 정의 도메인이 없는 GitHub Enterprise Cloud의 경우 `api.github.com/graphql.`을 사용합니다.
- `PORT` - 브로커 클라이언트가 연결을 수락하는 로컬 포트. 기본값은 8000입니다.
- `BROKER_CLIENT_URL` - 브로커 클라이언트의 전체 URL로, GitHub Enterprise 배포 웹훅에서 액세스할 수 있어야 합니다. 예를 들어 `http://broker.url.example:8000.`과 같습니다.
  - `http://://` 및 포트 번호가 필요합니다.
  - 클라이언트를 HTTPS로 구성하려면 [추가 설정이 필요합니다](../advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker.md).
- `ACCEPT_IAC` - 기본적으로 인프라스트럭처-코드 (IaC)에서 사용되는 일부 파일 유형이 활성화되어 있지 않습니다. 예를 들어 Terraform 파일과 같은 IaC 파일에 브로커 액세스 권한을 부여하려면 `tf,yaml,yml,json,tpl.`과 같은 환경 변수 `ACCEPT_IAC`를 추가합니다.
- `ACCEPT_CODE` - Snyk 브로커 - 코드 프로그램을 사용할 때 기본적으로 Snyk 코드가 코드 스니펫을로드하지 않습니다. 코드 스니펫을 활성화하려면 `ACCEPT_CODE=true.`와 같은 환경 변수를 추가할 수 있습니다.
- `ACCEPT_APPRISK` - Snyk AppRisk를 활성화하여 애플리케이션 자산을 식별하고 모니터링하며 위험을 우선순위로 지정할 수 있습니다. Snyk AppRisk를 활성화하려면 환경 변수 `ACCEPT_APPRISK=true`를 추가합니다.
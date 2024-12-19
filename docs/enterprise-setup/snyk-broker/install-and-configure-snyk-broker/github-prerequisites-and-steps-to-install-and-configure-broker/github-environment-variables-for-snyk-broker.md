# GitHub - 환경 변수 Snyk Broker를 위한

다음 환경 변수는 GitHub를 위한 Broker Client를 구성하는 데 필요합니다:

- `BROKER_TOKEN` - Snyk Broker 토큰으로, Snyk Org 설정보기(app.snyk.io)에서 얻을 수 있습니다.
- `GITHUB_TOKEN` - 전체 `repo`, `read:org`, `admin:repo_hook` 스코프가 있는 개인 액세스 토큰입니다.
- `PORT` - Broker 클라이언트가 연결을 수락하는 로컬 포트입니다. 기본값은 8000입니다.
- `BROKER_CLIENT_URL` - GitHub 웹훅에서 접근할 수 있는 Broker 클라이언트의 완전한 URL입니다. 예: `http://broker.url.example:8000.` 이 URL은 PR Fixes 또는 PR Checks와 같은 기능에 액세스하는 데 필요합니다.
  - 이는 http:// 및 포트 번호를 가져야 합니다.
  - HTTPS로 클라이언트를 구성하려면 [추가 설정이 필요합니다](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
- `ACCEPT_IAC` - 기본적으로, Infrastructure-as-Code (IaC)에서 사용되는 일부 파일 유형은 활성화되어 있지 않습니다. 예를 들어 Terraform과 같은 저장소의 IaC 파일에 Broker 액세스 권한을 부여하려면, 환경 변수 `ACCEPT_IAC`에 `tf,yaml,yml,json,tpl`의 어떤 조합이든 추가할 수 있습니다.
- `ACCEPT_CODE` - 기본적으로 Snyk Code는 코드 스니펫을로드하지 않습니다. 코드 스니펫을 활성화하려면 환경 변수 `ACCEPT_CODE=true`를 추가할 수 있습니다.
- `ACCEPT_APPRISK` - Snyk AppRisk를 활성화하여 애플리케이션 자산을 식별하고, 모니터하며, 위험을 우선 순위를 매기도록 할 수 있습니다. Snyk AppRisk를 활성화하려면 환경 변수 `ACCEPT_APPRISK=true`를 추가하시면 됩니다.
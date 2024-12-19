# Azure Repos - Snyk Broker을 위한 환경 변수

Azure Repos의 Broker Client를 구성하는 데 필요한 다음 환경 변수들이 있습니다:

- `BROKER_TOKEN` - Snyk Broker 토큰은 Azure Repos 통합 설정 화면 (app.snyk.io)에서 얻습니다.
- `AZURE_REPOS_TOKEN` - Azure Repos 개인 액세스 토큰입니다. 토큰을 가져오거나 생성하는 방법은 [가이드](https://docs.microsoft.com/en-us/azure/devops/organizations/accounts/use-personal-access-tokens-to-authenticate?view=azure-devops\&tabs=preview-page)를 참조하십시오. 필요한 스코프: 사용자 지정 정의가 선택되어 있고 코드 아래에 _Read & write._를 선택합니다.
- `AZURE_REPOS_ORG` - 조직 이름은 Azure의 조직 개요 페이지에서 찾을 수 있습니다. Azure Repos On-prem의 일반적인 조직 이름은 `tfs/Main.`입니다. 여러 개의 Azure 조직이 있는 경우, 각각에 대해 Broker를 배포해야 합니다.
- `AZURE_REPOS_HOST` - Azure Repos 서버 배포의 호스트 이름은 `your.azure-server.domain.com`과 같습니다.
- `PORT` - Broker 클라이언트가 연결을 허용하는 로컬 포트입니다. 기본값은 8000입니다.
- `BROKER_CLIENT_URL` - 브로커 클라이언트의 전체 URL은 Azure Repos의 웹훅이 액세스할 수 있도록 지정해야 합니다. 예: `http://broker.url.example:8000.` 이 URL은 PR Fixes 또는 PR Checks와 같은 기능에 액세스하기 위해 필요합니다.
  - 이 URL은 http://와 포트 번호가 있어야합니다.&#x20;
  - 클라이언트를 HTTPS로 구성하려면 [추가 설정이 필요합니다](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
- `ACCEPT_IAC` - 기본적으로 인프라스트럭처-코드 (IaC)에서 사용되는 일부 파일 유형은 활성화되어 있지 않습니다. 예를 들어 Terraform의 경우, 리포지토리에서 IaC 파일에 브로커 액세스를 부여하려면 `ACCEPT_IAC` 환경 변수를 추가할 수 있습니다. 가능한 값은 `tf, yaml, yml, json, tpl`의 어떤 조합이든 가능합니다.
- `ACCEPT_CODE` - 기본적으로 {{Snyk Code}}는 코드 스니펫을로드하지 않습니다. 코드 스니펫을 활성화하려면 `ACCEPT_CODE=true`라는 환경 변수를 추가할 수 있습니다.
- `ACCEPT_APPRISK` - Snyk AppRisk를 활성화하여 응용 프로그램 자산을 식별하고 모니터링하며 리스크를 우선 순위를 지정할 수 있습니다. `ACCEPT_APPRISK=true` 환경 변수를 추가하여 활성화할 수 있습니다.
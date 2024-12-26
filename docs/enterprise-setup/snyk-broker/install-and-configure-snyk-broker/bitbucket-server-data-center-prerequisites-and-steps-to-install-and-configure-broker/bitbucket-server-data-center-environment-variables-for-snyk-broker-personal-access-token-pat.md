# Bitbucket Server/Data Center - Snyk 브로커 개인 액세스 토큰 (PAT)을 위한 환경 변수

다음 환경 변수는 브로커 클라이언트를 구성하는 데 필요합니다:

- `BROKER_TOKEN` - Snyk 브로커 토큰은 Bitbucket Server 통합 설정 뷰 (app.snyk.io)에서 얻을 수 있습니다.
- `BITBUCKET_PAT` - Bitbucket Server 개인 액세스 토큰입니다.
- `BITBUCKET` - Bitbucket Server 배포의 호스트 이름입니다. 예: `your.bitbucket-server.domain.com`.
- `BITBUCKET_API` - Bitbucket Server 배포의 API 엔드포인트입니다. `$BITBUCKET/rest/api/1.0`여야 합니다.
- `BROKER_CLIENT_URL` - Bitbucket Server에서 웹훅에 액세스할 수 있는 브로커 클라이언트의 전체 URL입니다. 예: `http://broker.url.example:8000`. 이 URL은 PR Fixes 또는 PR Checks와 같은 기능에 액세스하는 데 필요합니다.
  - 이 URL은 http:// 및 포트 번호를 가져야 합니다.
  - 클라이언트를 HTTPS로 구성하려면 [추가 설정이 필요합니다](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
- `PORT` - 브로커 클라이언트가 연결을 수락하는 로컬 포트입니다. 기본값은 8000입니다.
- `ACCEPT_IAC` - 기본적으로 인프라스트럭처-as-Code (IaC)에서 사용되는 일부 파일 유형이 활성화되지 않습니다. 예를 들어 Terraform과 같은 IaC 파일에 브로커 액세스 권한을 부여하려면 `tf,yaml,yml,json,tpl`의 어떤 조합도 포함된 `ACCEPT_IAC` 환경 변수를 추가하면 됩니다.
- `ACCEPT_CODE` - 기본적으로 는 코드 조각을 로드하지 않습니다. 코드 조각을 활성화하려면 `ACCEPT_CODE=true` 환경 변수를 추가하면 됩니다.
- `ACCEPT_APPRISK` - Snyk AppRisk를 활성화하여 응용 프로그램 자산을 식별하고, 모니터링하며, 위험을 우선 순위로 설정합니다. Snyk AppRisk를 활성화하려면 `ACCEPT_APPRISK=true` 환경 변수를 추가하면 됩니다.
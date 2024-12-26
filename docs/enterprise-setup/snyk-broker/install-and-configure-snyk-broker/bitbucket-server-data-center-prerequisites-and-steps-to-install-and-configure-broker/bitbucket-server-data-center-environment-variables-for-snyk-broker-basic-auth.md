# Bitbucket Server/Data Center - 환경 변수 for Snyk Broker

다음 환경 변수는 Broker 클라이언트를 구성하는 데 필요합니다:

* `BROKER_TOKEN` - Snyk Broker 토큰입니다. 이는 Bitbucket Server 통합 설정 보기(app.snyk.io)에서 얻을 수 있습니다.
* `BITBUCKET_USERNAME` - Bitbucket Server 사용자 이름입니다.
* `BITBUCKET_PASSWORD` - Bitbucket Server 암호입니다. 암호 대신 API 토큰을 사용할 수 있습니다.
* `BITBUCKET` - Bitbucket Server 배포의 호스트 이름입니다. 예: `your.bitbucket-server.domain.com`.
* `BITBUCKET_API` - Bitbucket Server 배포의 API 엔드포인트입니다. `$BITBUCKET/rest/api/1.0` 여야 합니다.
* `BROKER_CLIENT_URL` - Broker 클라이언트의 전체 URL로, Bitbucket Server가 웹훅에 접근할 수 있는 URL입니다. 예: `http://broker.url.example:8000`. 이 URL은 PR Fixes 또는 PR Checks와 같은 기능에 액세스하는 데 필요합니다.
  * 반드시 http://와 포트 번호가 있어야 합니다.
  * 클라이언트를 HTTPS로 구성하려면, [추가 설정이 필요합니다](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
* `PORT` - Broker 클라이언트가 연결을 수락하는 로컬 포트입니다. 기본값은 8000입니다.
* `ACCEPT_IAC` - 기본적으로 IaC(Infrastructure-as-Code)에서 사용되는 일부 파일 유형이 활성화되어 있지 않습니다. 예를 들어 Terraform과 같은 리포지토리의 IaC 파일에 Broker 액세스 권한을 부여하려면, 간단히 `ACCEPT_IAC` 환경 변수를 추가하고 `tf,yaml,yml,json,tpl`과 같은 조합을 사용할 수 있습니다.
* `ACCEPT_CODE` - 기본적으로 는 코드 스니펫을로드하지 않습니다. 코드 스니펫을 활성화하려면 `ACCEPT_CODE=true`와 같은 환경 변수를 추가할 수 있습니다.
* `ACCEPT_APPRISK` - Snyk AppRisk를 활성화하여 애플리케이션 자산을 식별하고 모니터링하며 리스크를 우선 순위로 정할 수 있습니다. Snyk AppRisk를 활성화하려면 `ACCEPT_APPRISK=true`와 같은 환경 변수를 추가하십시오.
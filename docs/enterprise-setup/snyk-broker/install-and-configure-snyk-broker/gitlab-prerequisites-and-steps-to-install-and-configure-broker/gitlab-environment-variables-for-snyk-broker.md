# GitLab - Snyk 브로커를 위한 환경 변수

GitLab에 브로커 클라이언트를 구성하는 데 필요한 환경 변수는 다음과 같습니다:

- `BROKER_TOKEN` - GitLab 통합 설정 뷰 (app.snyk.io)에서 얻은 Snyk 브로커 토큰입니다.
- `GITLAB_TOKEN` - `api` 범위를 갖는 GitLab 개인 액세스 토큰입니다.
- `GITLAB` - GitLab 배포의 호스트 이름입니다. 예: `your.gitlab.domain.com` 또는 `GitLab.com`.
- `PORT` - 브로커 클라이언트가 연결을 수락하는 로컬 포트입니다. 기본값은 8000입니다.
- `BROKER_CLIENT_URL` - GitLab.com 또는 온프레미스 GitLab 배포에서 웹훅 연결성을 설정하기 위해 브로커 클라이언트에 연결할 수 있는 전체 URL입니다. 이는 `http://broker.url.example:8000`와 같은 완전한 URL이어야 합니다.
  - 이는 http:// 및 포트 번호를 가져야 합니다.
  - 클라이언트를 HTTPS로 구성하려면 [추가 설정이 필요합니다](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
- `ACCEPT_IAC` - 기본적으로 인프라스트럭처-애스 코드 (IaC)에 사용되는 일부 파일 유형이 활성화되지 않습니다. 예를 들어 Terraform과 같은 IaC 파일에 브로커 액세스를 허용하려면 `ACCEPT_IAC=tf,yaml,yml,json,tpl`과 같은 환경 변수를 추가하면 됩니다.
- `ACCEPT_CODE` - 기본적으로 는 코드 스니펫을로드하지 않습니다. 코드 스니펫을 활성화하려면 `ACCEPT_CODE=true`와 같은 환경 변수를 추가하면 됩니다.
- `ACCEPT_APPRISK` - Snyk AppRisk를 활성화하여 애플리케이션 자산을 식별하고 모니터링하며 위험을 우선적으로 처리할 수 있습니다. 이를 활성화하려면 `ACCEPT_APPRISK=true`와 같은 환경 변수를 추가하면 됩니다.
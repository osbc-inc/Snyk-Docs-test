# Jira - Snyk 브로커를 위한 환경 변수

다음 환경 변수는 Jira용 브로커 클라이언트를 구성하는 데 필요합니다:

- `BROKER_TOKEN` - Snyk 브로커 토큰은 Jira 통합 설정 뷰에서 얻습니다.
- `JIRA_USERNAME` - Jira 사용자 이름.
- `JIRA_PASSWORD` - Jira 암호.
- `JIRA_HOSTNAME` - Jira 배포의 호스트 이름, 예: `your.jira.domain.com`.
- `BROKER_CLIENT_URL` - Jira에서 웹훅을 통해 액세스할 브로커 클라이언트의 전체 URL, 예: `http://my.broker.client:8000`
  - 이것은 http://와 포트 번호를 가져야합니다.
  - HTTPS로 클라이언트를 구성하려면 [추가 설정이 필요합니다](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation/https-for-broker-client-with-docker).
- `PORT` - 브로커 클라이언트가 연결을 수락하는 로컬 포트입니다. 기본값은 8000입니다.
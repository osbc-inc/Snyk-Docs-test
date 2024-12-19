# Snyk 브로커를 위한 사전 점검

사전 점검의 주요 목표는 브로커 클라이언트 시작 시 에러와 잘못된 구성을 일찍 발견하는 것입니다. 이를 통해 사용 중에 나중에 발견되는 것을 방지합니다. 점검이 성공적이든 아니든 브로커 클라이언트는 시작됩니다. 다음의 점검이 가능합니다.

## `broker-server-status`

브로커 서버 상태 확인은 브로커 서버와의 연결성을 확인합니다. `{BROKER_SERVER_URL}/healthcheck`로 GET 요청을 수행합니다.

지정되지 않으면, BROKER\_SERVER\_URL은 [https://broker.snyk.io](https://broker.snyk.io/) 입니다.

## `rest-api-status`

REST API 상태 확인은 Snyk REST API와의 연결성을 확인합니다. `{API_BASE_URL}/rest/openapi`로 GET 요청을 수행합니다. 이 점검은 고가용성 모드가 활성화된 경우에만 실행됩니다.

지정되지 않으면, `API_BASE_URL`은 [https://api.snyk.io](https://api.snyk.io/) 입니다. 추가 URL은 [Regional hosting and data residency](../../working-with-snyk/regional-hosting-and-data-residency.md)를 참조하세요.

## `broker-client-url-validation`

브로커 클라이언트 구성 확인은 `BROKER_CLIENT_URL` 값이 최대한 유효한지 확인합니다. 스키마 (`http 또는 https`)를 포함하고 있는지, `https`인 경우 SSL 인증서 + 키가 로드되었는지 또는 상위 스트림에서 TLS 종료가 이루어졌는지 확인합니다.

TLS 종료를 사용하는 경우, 브로커 클라이언트에서 인증서 + 키가 필요하지 않다면, 사전 점검 시 TLS 종료를 신호하는 환경 변수 `BROKER_CLIENT_URL_TLS_TERMINATED`를 추가하세요.

기본 설정은 없습니다.

{% hint style="info" %}
환경 변수 `PREFLIGHT_CHECKS_ENABLED=false`를 사용하여 사전 점검 기능을 비활성화하여, 브로커 클라이언트가 시작될 때 어떠한 점검도 실행되지 않도록 할 수 있습니다.
{% endhint %}
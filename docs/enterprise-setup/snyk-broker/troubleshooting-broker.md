# 문제 해결 브로커

{% hint style="info" %}
**다중 테넌트 설정**\
브로커 및/또는 코드 에이전트를 다중 테넌트 환경에서 사용하려고 할 때 추가 변수가 필요합니다. 자세한 내용은 [지역 호스팅 및 데이터 수용소](../../working-with-snyk/regional-hosting-and-data-residency.md)를 참조하십시오.
{% endhint %}

이 페이지에는 다음과 관련된 정보 및 지침이 있습니다.

- [브로커 클라이언트와 함께 로깅](troubleshooting-broker.md#logging-with-the-broker-client)
- 모니터링 기능을 사용한 기본 문제 해결, [헬스체크](troubleshooting-broker.md#monitoring-healthcheck) 및 [시스템 체크](troubleshooting-broker.md#monitoring-systemcheck)
- [독립형 브로커 문제 해결](troubleshooting-broker.md#troubleshooting-standalone-broker)
- GitHub 및 GitHub Enterprise의 큰 매니페스트 파일 지원(>1MB) : [연관 문제](troubleshooting-broker.md#support-of-big-manifest-files-greater-than-1mb-for-github-and-github-enterprise)
- [코드 에이전트로 브로커 문제 해결](troubleshooting-broker.md#troubleshooting-broker-with-code-agent)
- [호스트에서 로그아웃 시 컨테이너 온라인 유지](troubleshooting-broker.md#containers-go-down-when-you-log-out-of-the-host)

더 포관적인 문제 해결 정보는 [브로커 문제 해결 FAQ](https://support.snyk.io/s/article/Broker-Troubleshooting)를 참조하십시오.

## 브로커 클라이언트와 함께 로깅

기본적으로 브로커의 로그 레벨은 INFO로 설정됩니다. HTTP 상태 코드에 상관없이 SCM 응답은 브로커 클라이언트에 의해 기록됩니다. 로깅 동작을 변경하려면 다음 환경 변수를 설정하십시오.

| 키               | 기본값 | 설명                                                          |
| ----------------- | ------- | -------------------------------------------------------------- |
| LOG\_LEVEL        | info    | 모든 로그를 위해 "debug"로 설정                                 |
| LOG\_ENABLE\_BODY | false   | 클라이언트 로그에 응답 바디를 포함하려면 "true"로 설정               |

보통 작업에서 로그를 간결하게 유지하기 위해, Snyk은 INFO 레벨에서 최소한의 정보를 생성하여 Snyk을 통해 클라이언트로 전송되는 요청을 추적하고, 대상 시스템(Github, Gitlab, JIRA 등)에 대해 수행된 하향 요청 및 URL 히트 및 수신한 응답 코드를 로깅합니다. \\

환경 변수 `LOG_INFO_VERBOSE="true"`을 설정하면 이러한 로그 라인에 헤더가 추가되어 debug를 사용하지 않고 헤더를 표시할 수 있게 됩니다.

{% hint style="warning" %}
기본 로깅을 덮어쓰면 API 요청 등 다른 프로세스에서 제공되는 로그가 포함될 수 있고, 인증 정보가 나열될 수 있습니다. 증가된 로깅이 활성화된 브로커 로그를 보내기 전에, 비밀번호나 토큰을 찾아 일괄적으로 제거하십시오.
{% endhint %}

## 모니터링: 헬스체크

브로커는 `/healthcheck` 엔드포인트를 노출하여 실행 중인 응용프로그램의 헬스를 모니터링할 수 있습니다. 이 엔드포인트는 내부 요청이 성공하면 상태 코드 `200 OK`를 반환하고 응답 바디에 `{ ok: true }`를 반환합니다.

브로커 클라이언트의 경우 이 엔드포인트는 또한 브로커 웹소켓 연결의 상태를 보고합니다. 웹소켓 연결이 열리지 않은 경우, 이 엔드포인트는 상태 코드 `500 Internal Server Error`와 응답 바디에 `{ ok: false }`가 나타납니다.

이 상태는 기본 설정으로 [http://localhost:8000/healthcheck](http://localhost:8000/healthcheck)을 연결하여 테스트할 수 있습니다.

헬스체크 엔드포인트의 위치를 변경하려면 환경 변수에서 대체 경로를 지정할 수 있습니다.

```dockerfile
ENV BROKER_HEALTHCHECK_PATH /path/to/healthcheck
```

## 모니터링: 시스템 체크

브로커 클라이언트는 `/systemcheck` 엔드포인트를 노출하며 브로커된 서비스(Git 또는 다른 SCM 또는 브로커된 컨테이너 레지스트리)의 연결 및 자격 증명을 유효성 검사할 수 있습니다. 이 엔드포인트는 브로커 클라이언트가 사전 구성된 URL로 요청을 보내고 요청의 성공 여부를 보고합니다. 지원되는 구성 요소는 다음과 같습니다:

- `BROKER_CLIENT_VALIDATION_URL` - 요청을 보낼 URL
- `BROKER_CLIENT_VALIDATION_AUTHORIZATION_HEADER` - \[선택 사항] 요청의 `Authorization` 헤더 값. `BROKER_CLIENT_VALIDATION_BASIC_AUTH`와 상호 배타적입니다.
- `BROKER_CLIENT_VALIDATION_BASIC_AUTH` - \[선택 사항] 기본 인증 자격 증명(`username:password`)을 base64로 인코딩하고 요청의 `Authorization` 헤더 값으로 배치합니다. `BROKER_CLIENT_VALIDATION_AUTHORIZATION_HEADER`와 상호 배타적입니다.
- `BROKER_CLIENT_VALIDATION_METHOD` - \[선택 사항] 요청의 HTTP 메소드 (기본값은 `GET`)
- `BROKER_CLIENT_VALIDATION_TIMEOUT_MS` - \[선택 사항] 밀리초 단위의 요청 타임아웃 (기본값은 5000 ms)

이 엔드포인트는 내부 요청이 성공할 때 상태 코드 `200 OK`를 반환하고 응답 바디에 `{ ok: true }`를 반환합니다. 내부 요청이 실패하는 경우, 이 엔드포인트는 상태 코드 `500 Internal Server Error`와 응답 바디에 `{ ok: false }`를 반환합니다.

이 상태는 기본 설정으로 [http://localhost:8000/systemcheck](http://localhost:8000/systemcheck)을 연결하여 테스트할 수 있습니다.

`/systemcheck` 기능을 활성화하여 브로커와 Nexus 간의 연결성을 확인하는 예제:\
`-e BROKER_CLIENT_VALIDATION_URL=https://[username:password]@acme.com/service/rest/v1/status[/check] /`\
`snyk/broker:nexus`

시스템체크 엔드포인트의 위치를 변경하려면 환경 변수에서 대체 경로를 지정할 수 있습니다.

```dockerfile
ENV BROKER_SYSTEMCHECK_PATH /path/to/systemcheck
```

{% hint style="info" %}
Snyk 브로커는 mTLS 방법으로 인증을 지원하지 않습니다.
{% endhint %}

## 독립형 브로커 문제 해결

브로커를 실행한 후 온프레미스 Git에 연결하는 데 여전히 오류가 있는 경우 다음 문제 해결 단계를 사용하십시오.

1. 브로커 컨테이너에 관련 로그가 있는지 확인하십시오. 이를 위해 온프레미스 Git에 연결을 시도하십시오. 이는 통합 탭으로 이동하여 가져오려고 시도함으로써 브로커 로그에 로그를 생성합니다.
2. 컨테이너의 로그를 검토하십시오. Docker에서 `docker logs <container id>`를 실행하여 수행할 수 있습니다.
3. 문제가 발생하는 위치를 확인하려고 로그를 검토하십시오.

### 독립형 브로커의 일반적인 문제점

* 앞의 단계를 수행한 후에 로그가 없으면 고객이 올바른 브로커 토큰을 가지고 있는지 확인하십시오. 그렇다면 웹소켓이 설정되었는지 확인하십시오. 일부 방화벽은 이를 차단할 수 있습니다.
* 온프레미스 Git에 대한 요청의 HTTP 코드를 검토하십시오.
  * **404 - Not Found** - 도커 실행 명령에 올바른 정보가 있는지 확인하십시오.
  * **401/403** - 자격 증명을 확인하십시오.
  * SSL에 관련된 언급이 있으면, 이는 자체 서명된 인증서에 의해 발생할 수 있습니다. 올바른 인증서를 마운트하거나 `-e NODE_TLS_REJECT_UNAUTHORIZED=0` 플래그를 사용해야 합니다.

### 독립형 브로커의 연결성 테스트

브로커와 에이전트는 이미지에 `curl`을 가지고 있지 않습니다. 에이전트 또는 컨테이너 레지스트리 또는 SCM과 같은 엔드포인트로의 연결성을 테스트하려면 다음 명령을 사용할 수 있습니다.

```
# 노드 시작
node

# http를 사용한 URL 테스트
http = require("http")
http.get("<URL_HERE>", res => {console.log(`statusCode: ${res.statusCode}`)})

# https를 사용한 URL 테스트
https = require("https")
https.get('<URL_HERE>', res => {console.log(`statusCode: ${res.statusCode}`)})
```

## GitHub 및 GitHub Enterprise의 큰 매니페스트 파일 (>1MB) 지원

Open Fix/Upgrade PRs 또는 PR/반복 테스트가 큰 매니페스트 파일(>1MB)을 가져오는 데 실패할 수 있습니다. 이 문제를 해결하려면 큰 매니페스트 파일을 허용하는 [Docker](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-open-source-scans-sca-of-large-manifest-files-docker-setup) 또는 [Helm](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-helm-chart-installation/snyk-open-source-scans-sca-of-large-manifest-files-helm-setup) 지침을 따르십시오.

## 코드 에이전트와 함께 브로커 문제 해결

{% hint style="warning" %}
**사용이 중단되었습니다**

코드 에이전트는 더이상 유지되지 않습니다.

Snyk 코드 분석을 실행하는 더 선호되는 방법은 [브로커된 코드](git-clone-through-broker.md)를 통해 Snyk 브로커를 통해 수행하는 것입니다. 코드 에이전트는 이점이 없는 대체 방법입니다. 자세한 내용은 Snyk 통합 컨설턴트 또는 기술적 성공 매니저에게 연락하거나 [Snyk 지원](https://support.snyk.io)에 문의하십시오.

자동 [PR 확인](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks) 기능은 Snyk 브로커 - 코드 에이전트에서 지원되지 않습니다.
{% endhint %}

<figure><img src="https://lh3.googleusercontent.com/r_qtONpOOEW35gdyoBcWDAiC6j04M76q8mh922SHor4bdNZdt83sj2kP7d5hbzYcWVXp4Q2hZEiCeAVOmcj4Bu1yFPdnyp3rK7kKeBK8DZEd9S133Xn3YdjddclVf5maEbP23Jor" alt="&#x22;&#x22;"><figcaption><p> 브로커와 함께 코드 분석 워크플로우</p></figcaption></figure>

코드 에이전트와 브로커의 문제를 해결하기 위한 가장 좋은 방법은 통신 흐름을 이해하는 것입니다. 트래픽은 Snyk > 브로커 클라이언트 > 코드 에이전트 > 온프레미스 Git > 코드 에이전트 > Snyk로 이동합니다.

코드 에이전트의 대부분의 문제는 이러한 지점 중 하나에서 트래픽이 중단되기 때문에 발생합니다.

### 코드 에이전트 문제 해결

독립형 브로커와 마찬가지로 코드 에이전트를 해결하려면 로그를 생성해야 합니다. 이를 위해 저장소를 가져오려 시도함으로써 진행해야 합니다.

1. 브로커가 올바르게 작동하고 저장소 목록을 볼 수 있는지 확인하십시오. 이 작업이 작동하지 않으면 독립형 브로커 문제 해결 단계를 검토하십시오.
2. 저장소를 가져오려 시도한 후 `번들 작성 실패` 오류 메시지가 나타나면 컨테이너의 로그에서 오류를 검토하십시오.
3. 브로커 컨테이너부터 시작하십시오. `docker logs <container id>`를 실행하십시오.
4. `snykgit` 문자열을 찾으십시오. 이는 브로커 컨테이너에서 코드 에이전트 컨테이너로의 API 호출입니다. 200 코드 이외의 결과가 나오면 브로커와 코드 에이전트 사이의 통신에 문제가 있는 것입니다. 올바른 플래그가 도커 실행 명령에 설정되어 있는지 확인하십시오. 또한 올바른 Docker 네트워크를 설정했는지 확인하십시오.
5. 코드 에이전트의 로그를 검토하려면 `docker logs <container id>`를 실행하십시오.

### 코드 에이전트에서 발생하는 일반적인 문제

* 온프레미스 Git과의 통신이 작동하지 않습니다. 코드 클론 시 SSL 관련 참조가 있을 경우 404 오류가 발생합니다. 이는 자체 서명된 인증서로 인해 발생할 수 있습니다. 올바른 인증서를 마운트했는지 확인하거나 `-e NODE_TLS_REJECT_UNAUTHORIZED=0` 플래그를 사용하세요.
* 메시지: `“Uploaded Repo”`가 표시되면, 코드 에이전트와 브로커가 올바르게 구성된 것입니다. 여전히 가져오기 로그에서 오류가 발생하는 경우, [Snyk Support](https://support.snyk.io)로 문의하십시오.

## 호스트에서 로그아웃할 때 컨테이너가 종료되는 문제

호스트에서 로그아웃할 때 브로커 생태계를 포함한 컨테이너가 종료되는 경우, 로그아웃 시 컨테이너가 온라인 상태를 유지하도록 아래 명령을 실행하세요:

`loginctl enable-linger $(whoami)`
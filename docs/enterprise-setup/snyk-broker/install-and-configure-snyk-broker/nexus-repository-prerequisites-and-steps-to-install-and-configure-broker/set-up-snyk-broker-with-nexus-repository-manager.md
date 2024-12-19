# Nexus Repository - Docker를 사용하여 설치 및 구성

{% hint style="info" %}
**기능 가용성**

Nexus Repository Manager와의 통합은 엔터프라이즈 플랜에서만 제공됩니다. 자세한 정보는 [요금제 및 가격 책정](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

설치하기 전에, [전제 조건](./) 및 일반적인 설치 지침을 검토하십시오.([Docker를 사용한 설치 및 구성](../install-and-configure-broker-using-docker.md))

이 통합은 온프레미스 Nexus Repository Manager 배포와 안전한 연결을 보장하기 위해 유용합니다.

Nexus Repository Manager와의 브로커 없는 통합에 대한 정보(지원되는 환경 및 버전, 사용자 권한 포함)는 [Nexus Repository Manager 설정](../../../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/)을 참조하십시오. Nexus Container Registry와의 브로커 통합에 대한 정보는 [Snyk Broker - Container Registry 에이전트](../../snyk-broker-container-registry-agent/)을 참조하십시오.

## Nexus 플러그인을 사용하기 위한 브로커 설정

### Nexus 3 및 Nexus 2 구성을 위한 Docker pull

Nexus 3 배포에 브로커 클라이언트를 사용하려면, **다음을 실행**하십시오: `docker pull snyk/broker:nexus`.

Nexus 2 배포에 브로커 클라이언트를 사용하려면, **다음을 실행**하십시오: `docker pull snyk/broker:nexus2`.

환경 변수의 정의는 [Nexus Repository - Snyk Broker를 위한 환경 변수](nexus-repository-environment-variables-for-snyk-broker.md)를 참조하십시오.

### Nexus 3 및 Nexus 2 통합을 위한 브로커 클라이언트 설정을 위한 Docker 실행 명령어

{% hint style="info" %}
**기본값이 아닌 영역을 위한 멀티 테넌트 설정**\
기본값이 아닌 영역에서 Snyk Broker를 설정할 때는, 특정 URL을 필요로 하는 추가 환경 변수가 필요합니다. URL 및 예제에 대한 자세한 내용은 [지역 호스팅 및 데이터 보유, 브로커 URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)을 참조하십시오.
{% endhint %}

다음 명령어를 잘 설정된 브로커 클라이언트를 Nexus 3에서 사용하도록 설정하는 데 **복사**하십시오. 관련 설정을 제공하여 Docker 컨테이너를 실행할 수 있습니다:

```console
docker run --restart=always \
           -p 7341:7341 \
           -e BROKER_TOKEN=secret-broker-token \
           -e BASE_NEXUS_URL=https://[<user>:<pass>@]<your.nexus.hostname> \
           -e BROKER_CLIENT_VALIDATION_URL=https://<your.nexus.hostname>/service/rest/v1/status[/check] \
           -e RES_BODY_URL_SUB=https://<your.nexus.hostname>/repository \
       snyk/broker:nexus
```

다음 명령어를 잘 설정된 브로커 클라이언트를 Nexus 2에서 사용하도록 설정하는 데 **복사**하십시오. 관련 설정을 제공하여 Docker 컨테이너를 실행할 수 있습니다:

```
docker run --restart=always \
  -p 7341:7341 \
  -e BROKER_TOKEN=<secret-broker-token> \
  -e BASE_NEXUS_URL=https://[username:password]@acme.com \
  -e RES_BODY_URL_SUB=https://acme.com/nexus/content/(groups|repositories) \
  snyk/broker:nexus2
```

`BASE_NEXUS_URL`를 찾으려면 Nexus UI를 방문하고 **Administration** 아래의 서버 탭으로 이동한 다음 **Base URL** 항목을 선택하세요(마지막 슬래시 없음). 이것은 일반적으로 `/nexus`로 끝나지만 기본이 아닌 배포에 따라 다를 수 있습니다. 사용자 정의 기본 URL이 있는 경우 `NEXUS_URL` 환경 변수를 설정하여 저장소가 있는 URL을 가리키도록 설정해야 합니다. 기본적으로 `/nexus/content`로 구성되어 있지만 사용자 정의 기본 URL과 유사한 형식이어야 합니다.

**Docker run 명령어 대신에** Nexus3 통합을 위한 환경 변수 재정의를 위한 파생 Docker 이미지를 사용할 수도 있습니다. Nexus3 통합을 위한 환경 변수에 대한 자세한 내용은 [파생 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하십시오.

## 브로커 클라이언트 컨테이너 시작 및 Nexus Repository Manager와의 연결 확인

브로커 클라이언트 구성을 붙여넣어 브로커 클라이언트 컨테이너를 시작하세요.

Broker Client `/systemcheck` 엔드포인트에 요청을 보내어 연결 상태를 확인하세요.

예시: `curl http://172.17.0.2:7341/systemcheck`

다음과 같은 형식의 성공 출력을 확인할 수 있습니다:

`{"brokerClientValidationUrl":"https://acme.com/service/rest/v1/status","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"brokerClientValidationUrlStatusCode":200,"ok":true}`

또는 다음과 같은 형식의 실패 출력을 확인할 수 있습니다:

`{"brokerClientValidationUrl":"https://acme.com/service/rest/v1/status","brokerClientValidationMethod":"GET","brokerClientValidationTimeoutMs":5000,"ok":false,"error":"ETIMEDOUT"}`
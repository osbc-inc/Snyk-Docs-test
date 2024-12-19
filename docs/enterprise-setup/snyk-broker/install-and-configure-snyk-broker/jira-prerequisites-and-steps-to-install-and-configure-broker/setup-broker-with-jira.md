# Jira - Docker를 사용하여 설치 및 구성

설치하기 전에 [전제 조건](./) 및 Docker를 사용한 설치에 대한 일반 지침을 검토하십시오([Docker 사용하여 브로커를 설치하고 구성하는 방법](../install-and-configure-broker-using-docker.md)).

이 통합은 온프레미스 Jira 배포물과의 안전한 연결을 보장하기 위해 유용합니다.

## Jira를 위해 사용할 Broker 구성

Jira 배포물과 Broker 클라이언트를 사용하려면 `docker pull snyk/broker:jira`를 실행하십시오. 환경 변수의 정의에 대해서는 [Snyk Broker를 위한 Jira 환경 변수](jira-environment-variables-for-snyk-broker.md)를 참조하십시오.

## Jira를 위해 Broker 클라이언트를 설정하기 위한 Docker 실행 명령어

{% hint style="info" %}
**기본 설정이 아닌 다른 지역을 위한 멀티 테넌트 설정**\
기본 설정이 아닌 지역에서 Snyk Broker를 설정하는 경우, 특정 URL이 필요한 추가 환경 변수가 필요합니다. URL 및 예제에 대해서는 [지역 호스팅 및 데이터 저장소, Broker URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

**다음 명령어를 복사**하여 Jira와 사용할 완전히 구성된 Broker 클라이언트를 설정하십시오. 관련 구성을 제공하여 Docker 컨테이너를 실행할 수 있습니다:

```console
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=비밀-브로커-토큰 \
           -e JIRA_USERNAME=사용자명 \
           -e JIRA_PASSWORD=비밀번호 \
           -e JIRA_HOSTNAME=your.jira.domain.com \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e PORT=8000 \
       snyk/broker:jira
```

필요한 경우, [Snyk Broker Docker 설치를 위한 고급 구성](../advanced-configuration-for-snyk-broker-docker-installation/)으로 이동하여 Jira 인스턴스가 사설 인증서를 사용할 때 브로커 클라이언트 구성에 인증서(CA, Certificate Authority)를 제공하는 등의 구성 변경을 수행하십시오.

**Docker run 명령어 대신 대체로** Docker 이미지를 파생하여 Broker 클라이언트 통합을 설정할 수 있습니다. Jira 통합을 위해 재정의해야 하는 환경 변수에 대해서는 [파생 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하십시오.

## SSO가 활성화된 JIRA를 위한 Jira PAT 인증

SSO가 활성화되어 있으면 JIRA는 일반적으로 사용자 이름 및 암호의 사용을 금지하고 개인 액세스 토큰(PAT)의 사용을 요구합니다.

SSO가 활성화되어 있는 경우, 대신 bearer 토큰을 사용하는 권한 헤더를 사용할 특정 Jira 버전을 사용해야 합니다. 이 버전을 사용하려면 다음 구성을 제공하십시오:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=비밀-브로커-토큰 \
           -e JIRA_PAT=<your_pat_token> \
           -e JIRA_HOSTNAME=your.jira.domain.com \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e PORT=8000 \
       snyk/broker:jira-bearer-auth
```

## Broker 클라이언트 컨테이너 시작 및 Jira와의 연결 확인

Broker 클라이언트 구성을 붙여넣어 Broker 클라이언트 컨테이너를 시작하십시오.

컨테이너가 설정되고 Jira 통합 페이지에서 Jira에 대한 연결이 표시되면, 프로젝트 아래에서 Jira 티켓을 생성할 수 있습니다.

## Jira와 함께 브로커의 기본 문제 해결

* `docker logs <container id>`를 실행하여 오류를 찾습니다. 여기서 컨테이너 ID는 Jira 브로커 컨테이너 ID입니다.
* 관련 포트가 Jira에 노출되도록 합니다.
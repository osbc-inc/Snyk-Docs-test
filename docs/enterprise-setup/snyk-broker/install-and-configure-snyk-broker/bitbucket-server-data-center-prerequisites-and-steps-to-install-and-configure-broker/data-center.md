# Bitbucket Server/Data Center - Docker를 사용하여 설치 및 구성

설치 전에 [prerequisites](./)를 검토하고 [Docker](../install-and-configure-broker-using-docker.md)를 사용한 설치에 대한 일반 설명을 살펴보세요.

이 통합은 온프레미스 Bitbucket 배포와 안전한 연결을 보장하기 위해 유용합니다.

이 페이지에서는 두 가지 다른 인증 체계를 설명합니다: [Basic Auth](data-center.md#configure-broker-to-be-used-with-bitbucket-using-basic-auth) 및 [Bearer (Personal Access Token)](data-center.md#configure-broker-to-be-used-with-bitbucket-using-personal-access-token-pat). Basic Auth 사용을 불가능하게 하는 Bitbucket Server 설정이 있을 수 있으므로 Bearer Auth를 선호합니다.

## Basic Auth를 사용하여 Bitbucket에서 Broker 사용 설정

다음은 Snyk Broker가 Bitbucket Server 배포와 함께 Broker Client를 사용하도록 구성하는 방법을 설명합니다.

BitBucket에서 Snyk Broker Client를 사용하려면 **다음을 실행**하세요: `docker pull snyk/broker:bitbucket-server`. 환경 변수의 정의에 대해서는 [BitBucket Server/Data Center - 환경 변수](bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md)를 참조하세요.

**필요한 경우**, [고급 구성 페이지](../advanced-configuration-for-snyk-broker-docker-installation/)로 이동하여 필요한 구성 변경(예: Bitbucket 인스턴스가 사설 인증서를 사용하는 경우 Broker Client 구성에 CA (Certificate Authority)를 제공하고 [프록시 지원](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md) 설정)을 수행하세요.

## Basic Auth를 사용하여 Bitbucket에 Broker Client 설정하기 위한 Docker 실행 명령

Open Source, IaC, Container, Code 파일 및 Snyk AppRisk 정보를 분석할 수 있도록 완전히 구성된 Broker Client를 설정하기 위해 **다음 명령어**를 복사하세요. 어플리케이션 자산을 식별하고 모니터링하여 위험을 우선 순위로 지정하기 위해 [Snyk AppRisk](../../../../scan-with-snyk/snyk-apprisk/)를 활성화하세요.

{% hint style="info" %}
**기본값이 아닌 다른 지역을 위한 멀티 테넌트 설정**\
기본값이 아닌 지역에서 Snyk Broker를 설정하는 경우 특정 URL을 필요로 하는 추가 환경 변수가 필요합니다. URL 및 예제에 대해서는 [지역 호스팅 및 데이터 저장소, Broker URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하세요.
{% endhint %}

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e BITBUCKET_USERNAME=<username> \
           -e BITBUCKET_PASSWORD=<password> \
           -e BITBUCKET=<your.bitbucket-server.domain.com (no http/s)> \
           -e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0 (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \
       snyk/broker:bitbucket-server
```

{% hint style="info" %}
Snyk AppRisk는 기본적으로 **`false`**로 설정됩니다. 이를 **`true`**로 설정하여 활성화할 수 있습니다.
{% endhint %}

**Docker run 명령어 대신** 도커 이미지를 파생하여 Broker Client 통합을 설정할 수도 있습니다. BitBucket Server/Data Center 통합을 위해 변경할 환경 변수에 대해서는 [파생된 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하세요.

## Personal Access Token (PAT)를 사용하여 Bitbucket에 Broker 사용 설정

다음은 Snyk Broker가 Personal Access Token (PAT)을 사용하여 Bitbucket Server 배포와 함께 Broker Client를 사용하도록 구성하는 방법을 설명합니다.

BitBucket에서 Snyk Broker Client를 사용하려면 **다음을 실행**하세요: `docker pull snyk/broker:bitbucket-server-bearer-auth`. 환경 변수의 정의에 대해서는 [Bitbucket Server/Data Center - Snyk Broker Basic Auth를 위한 환경 변수](bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md) 및 [Bitbucket Server/Data Center - Snyk Broker Personal Access Token (PAT)을 위한 환경 변수](bitbucket-server-data-center-environment-variables-for-snyk-broker-personal-access-token-pat.md)를 참조하세요.

**필요한 경우**, [고급 구성 페이지](../advanced-configuration-for-snyk-broker-docker-installation/)로 이동하여 필요한 구성 변경(예: Bitbucket 인스턴스가 사설 인증서를 사용하는 경우 Broker Client 구성에 CA (Certificate Authority)를 제공하고 [프록시 지원](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md) 설정)을 수행하세요.

## PAT를 사용하여 Bitbucket에 Broker Client 설정하기 위한 Docker 실행 명령

Open Source, IaC, Container, Code 파일 및 Snyk AppRisk 정보를 분석할 수 있도록 완전히 구성된 Broker Client를 설정하기 위해 **다음 명령어**를 복사하세요. 어플리케이션 자산을 식별하고 모니터링하여 위험을 우선 순위로 지정하기 위해 [Snyk AppRisk](../../../../scan-with-snyk/snyk-apprisk/)를 활성화하세요.

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e BITBUCKET_PAT=<personal-access-token> \
           -e BITBUCKET=<your.bitbucket-server.domain.com (no http/s)> \
           -e BITBUCKET_API=<your.bitbucket-server.domain.com/rest/api/1.0 (no http/s)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \
       snyk/broker:bitbucket-server-bearer-auth
```

{% hint style="info" %}
Snyk AppRisk는 기본적으로 **`false`**로 설정됩니다. 이를 **`true`**로 설정하여 활성화할 수 있습니다.
{% endhint %}

**Docker run 명령어 대신** 도커 이미지를 파생하여 Broker Client 통합을 설정할 수도 있습니다. BitBucket Server/Data Center 통합을 위해 변경할 환경 변수에 대해서는 [파생된 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하세요.

## Bitbucket와 연결을 확인하면서 Broker Client 컨테이너 시작 및 확인

Broker Client 구성을 붙여넣어 Broker Client 컨테이너를 시작하세요.

컨테이너가 실행되면 Bitbucket 통합 페이지에서 Bitbucket과의 연결이 표시되며 `프로젝트 추가`를 할 수 있습니다.

## BitBucket에서 Broker와 관련된 기본적인 문제 해결

* Bitbucket Broker 컨테이너 ID가 있는 경우 `docker logs <container id>`를 실행하여 오류를 확인하세요.
* 관련된 포트가 Bitbucket에 노출되었는지 확인하세요.
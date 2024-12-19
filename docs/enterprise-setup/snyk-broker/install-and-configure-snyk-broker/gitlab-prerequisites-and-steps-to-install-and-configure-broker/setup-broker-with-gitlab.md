# GitLab - Docker를 사용하여 설치하고 구성하기

설치하기 전에, [전제 조건](./)을 검토하고 [Docker를 사용한 일반적인 설치 지침](../install-and-configure-broker-using-docker.md)을 확인하세요.

이 통합은 온프레미스 또는 클라우드 GitLab 배포와 안전한 연결을 보장하기 위해 유용합니다.

## GitLab을 위해 Broker 구성

GitLab.com 또는 온프레미스 GitLab 배포와 함께 Broker 클라이언트를 사용하려면, `docker pull snyk/broker:gitlab`을 **실행**합니다. 환경 변수의 정의에 대한 자세한 내용은 [GitLab - Snyk Broker를 위한 환경 변수](gitlab-environment-variables-for-snyk-broker.md)를 참조하세요.

**필요한 경우**, [고급 구성 페이지](../advanced-configuration-for-snyk-broker-docker-installation/)로 이동하고, CA (Certificate Authority)를 제공하거나 (GitLab 인스턴스가 프라이빗 인증서를 사용하고 있는 경우) 또는 [프록시 지원](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md)을 설정하는 등 필요한 구성 변경을 수행하세요.

## GitLab을 위한 Broker 클라이언트 설정 Docker 실행 명령어

아래 명령어를 복사해서 사용하여 완전히 구성된 Broker 클라이언트를 설정하여 오픈 소스, IaC, 컨테이너, 코드 파일 및 Snyk AppRisk 정보를 분석하세요. Snyk AppRisk를 활성화하여 응용 프로그램 자산을 식별하고 모니터링하며 위험을 우선 순위로 설정하세요.

{% hint style="info" %}
**기본값이 아닌 다른 지역을 위한 멀티 테넌트 설정**\
기본값이 아닌 지역에서 Snyk Broker를 설정할 때는 특정 URL을 사용하는 추가 환경 변수가 필요합니다. URL과 예시에 대한 자세한 내용은 [지역 호스팅 및 데이터 소재, Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하세요.
{% endhint %}

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITLAB_TOKEN=<secret-gitlab-token> \
           -e GITLAB=<your.gitlab.domain.com (http/s 제외)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:포트)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \
       snyk/broker:gitlab
```

`REMOVE_X_FORWARDED_HEADERS=true` 환경 변수를 사용하여 Broker 클라이언트가 GitLab로의 요청에서 `XFF` 헤더를 제거합니다. 이렇게 하면 Broker가 올바르게 작동합니다.

Snyk AppRisk는 기본적으로 `false`로 설정되어 있습니다. 이를 `true`로 설정하여 활성화하세요.

**Docker run 명령어 대신에**, 파생된 Docker 이미지를 사용하여 Broker 클라이언트 통합을 설정할 수 있습니다. GitLab 통합을 위해 재정의할 환경 변수에 대한 자세한 내용은 [파생 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하세요.

## Broker 클라이언트 컨테이너 시작 및 GitLab과의 연결 확인

Broker 클라이언트 컨테이너를 시작하기 위해 Broker 클라이언트 구성을 붙여넣으세요.

컨테이너가 실행되면 GitLab 통합 페이지에서 GitLab과의 연결이 표시되며 `프로젝트 추가`할 수 있습니다.

## GitLab을 위한 Broker의 기본적인 문제 해결

* `docker logs <container id>` 명령어를 실행하여 오류를 확인하세요 (`container id`는 GitLab Broker 컨테이너 ID입니다).
* GitLab로 노출된 적절한 포트를 확인하세요.
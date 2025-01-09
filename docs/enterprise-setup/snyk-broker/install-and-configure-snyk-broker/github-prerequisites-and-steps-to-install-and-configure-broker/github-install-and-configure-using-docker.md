# GitHub - Docker를 사용한 설치 및 구성

설치하기 전에 [전제 조건](./) 및 [Docker를 사용한 설치에 대한 일반 지침](../install-and-configure-broker-using-docker.md)를 검토하십시오.

이 통합은 온프레미스 또는 클라우드 GitHub 배포와 안전한 연결을 보장하는 데 유용합니다.

## GitHub와의 Broker 통합 구성

GitHub용 **Snyk 브로커 클라이언트를 사용**하려면 `docker pull snyk/broker:github-com`을 **실행**하십시오. 환경 변수의 정의에 대해서는 [Snyk 브로커를 위한 GitHub - 환경 변수](github-environment-variables-for-snyk-broker.md)를 참조하십시오.

**필요한 경우**, [고급 구성 페이지](../advanced-configuration-for-snyk-broker-docker-installation/)로 이동하여 GitHub 인스턴스가 프라이빗 인증서를 사용하는 경우 브로커 클라이언트 구성에 CA(Certificate Authority)를 제공하거나 [프록시 지원](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md)을 설정하는 등 필요한 구성 변경을 수행하십시오.

## GitHub를 위한 브로커 클라이언트 설정을 위한 Docker 실행 명령

**다음 명령을 복사**하여 오픈 소스, IaC, 컨테이너, 코드 파일 (코드 에이전트로), 그리고 Snyk AppRisk 정보를 분석하기 위해 완전히 구성된 브로커 클라이언트를 설정하십시오. [Snyk AppRisk](../../../../scan-with-snyk/snyk-apprisk/)를 활성화하여 애플리케이션 자산을 식별하고 이를 모니터링하며 위험을 우선 순위로 지정합니다.

{% hint style="info" %}
**기본값이 아닌 다른 지역을 위한 멀티 테넌시 설정**\
기본값이 아닌 지역에서 Snyk 브로커를 설정하는 경우 특정 URL이 필요한 추가 환경 변수가 필요합니다. URL과 예제에 대해서는 [Regional hosting and data residency, Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITHUB_TOKEN=<secret-github-token> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:port)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \ 
       snyk/broker:github-com
```

{% hint style="info" %}
Snyk AppRisk는 기본적으로 \*\*`false`\*\*로 설정됩니다. \*\*`true`\*\*로 설정하여 활성화하십시오.
{% endhint %}

**Docker 실행 명령을 사용하는 대신에** GitHub 통합을 위해 환경 변수를 재정의하기 위해 파생된 Docker 이미지를 사용할 수도 있습니다. [파생된 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하여 GitHub 통합을 위해 오버라이드할 환경 변수를 확인하십시오.

## Broker 클라이언트 컨테이너 시작 및 GitHub와의 연결 확인

Broker 클라이언트 구성을 붙여넣어 Broker 클라이언트 컨테이너를 시작하십시오.

컨테이너가 실행되면 GitHub 통합 페이지에서 GitHub와의 연결을 확인할 수 있으며 `프로젝트 추가`를 할 수 있습니다.

## GitHub와의 Broker를 위한 기본적인 문제 해결

### **GitHub용 큰 매니페스트 파일 (> 1MB) 지원**

오픈 Fix/Upgrade PR 또는 PR/반복 테스트가 실패하는 이유 중 하나는 큰 매니페스트 파일(> 1MB)을 가져오는 것일 수 있습니다. 이 문제를 해결하려면 [Snyk Open Source 스캔 (SCA)의 대용량 매니페스트 파일 (Docker 설정)](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-open-source-scans-sca-of-large-manifest-files-docker-setup)의 추가 지침을 따라 브로커에 추가 변수를 활성화하십시오.

{% hint style="info" %}
가능한 최대 보안을 보장하기 위해 Snyk은 이 규칙을 기본으로 활성화하지 않습니다. 이 엔드포인트 사용은 경로에 특정 허용된 파일 이름이 포함되어 있지 않기 때문에 이 규칙의 사용은 이 리포지토리의 모든 파일에 대한 이론적 액세스를 의미합니다.
{% endhint %}

### **GitHub와의 Broker에 대한 추가 문제 해결**

* `docker logs <container id>`를 실행하여 오류를 확인하십시오. 여기서 `container id`는 GitHub 브로커 컨테이너 ID입니다.
* 관련 포트가 GitHub에 노출되었는지 확인하십시오.

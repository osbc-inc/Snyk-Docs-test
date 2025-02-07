# GitHub Enterprise - Docker를 사용한 설치 및 구성

설치하기 전에 [전제 조건](./)을 검토하고 [Docker를 사용한 설치에 대한 일반 지침](../install-and-configure-broker-using-docker.md)를 참조하십시오.

이 통합은 온프레미스 또는 클라우드 GitHub Enterprise 배포와 안전한 연결을 확보하는 데 유용합니다.

## GitHub Enterprise에 사용할 Broker 구성

Snyk 브로커 클라이언트를 GitHub Enterprise 배포와 함께 사용하려면 **`docker pull snyk/broker:github-enterprise`** 명령을 실행하십시오. 환경 변수의 정의에 대해서는 [GitHub Enterprise - Snyk 브로커의 환경 변수](github-enterprise-environment-variables-for-snyk-broker.md)를 참조하십시오.

**필요한 경우**, [고급 설정 페이지](../advanced-configuration-for-snyk-broker-docker-installation/)로 이동하여 GitHub Enterprise 인스턴스에서 개인 인증서를 사용하는 경우 브로커 클라이언트 구성에 CA(Certificate Authority)를 제공하고 [프록시 지원](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md)을 설정하는 등 필요한 구성 변경 사항을 수행하십시오.

## Docker 실행 명령어를 사용하여 GitHub Enterprise용 브로커 클라이언트 설정

다음 명령어를 복사하여 오픈 소스, IaC, 컨테이너, 코드 파일(코드 에이전트와 함께) 및 Snyk AppRisk 정보를 분석하는 완전히 구성된 브로커 클라이언트를 설정하십시오. [Snyk AppRisk](../../../../scan-with-snyk/snyk-apprisk/)를 활성화하여 응용 프로그램 자산을 식별하고 모니터링하며 위험을 우선순위에 따라 지정할 수 있습니다.

{% hint style="info" %}
**기본 이외의 영역에 대한 멀티 테넌트 설정**\
기본 영역이 아닌 영역에 사용하기 위해 Snyk 브로커를 설정하는 경우 특정 URL을 필요로 하는 추가 환경 변수가 필요합니다. URL 및 예제에 대해서는 [지역 호스팅 및 데이터 수용처, 브로커 URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<secret-broker-token> \
           -e GITHUB_TOKEN=<secret-github-token> \
           -e GITHUB=<your.ghe.domain.com (http또는https 없음)> \
           -e GITHUB_API=<your.ghe.domain.com/api/v3 (http또는 https 없음)> \
           -e GITHUB_GRAPHQL=<your.ghe.domain.com/api (http또는 https 없음)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (DNS/IP:포트)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \
       snyk/broker:github-enterprise
```

{% hint style="info" %}
Snyk AppRisk는 기본적으로 `false`로 설정됩니다. 이 기능을 `true`로 설정하여 활성화할 수 있습니다.
{% endhint %}

**Docker 실행 명령어를 대체로 사용하는 대안**으로 GitHub Enterprise 통합을 위해 환경 변수를 재정의하기 위해 파생된 Docker 이미지를 사용할 수 있습니다. GitHub Enterprise 통합을 위한 환경 변수를 재정의하는 경우 [파생된 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하십시오.

## 브로커 클라이언트 컨테이너 시작 및 GitHub Enterprise와의 연결 확인

브로커 클라이언트 설정을 붙여넣어 브로커 클라이언트 컨테이너를 시작하십시오.

컨테이너가 실행되면 GitHub Enterprise 통합 페이지에서 GitHub Enterprise와의 연결이 표시되며 `프로젝트 추가`가 가능해집니다.

## GitHub Enterprise와 함께 브로커의 기본적인 문제 해결

### **GitHub Enterprise용 큰 Manifest 파일 (> 1Mb) 지원**

오픈 소스 PR 또는 PR/반복 테스트가 실패하는 이유 중 하나는 큰 Manifest 파일(> 1MB)을 가져오는 것일 수 있습니다. 이 문제를 해결하려면 GitHub Enterprise 브로커에서 추가 변수를 활성화하십시오. 자세한 내용은 [큰 Manifest 파일의 Snyk 오픈 소스 스캔(SCA) (Docker 설정)](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-open-source-scans-sca-of-large-manifest-files-docker-setup)의 추가 지침을 따르십시오.

{% hint style="info" %}
최대한의 보안을 보장하기 위해, Snyk는 이 규칙을 기본으로 활성화하지 않습니다. 이 엔드포인트의 이용은 경로에 특정한 허용된 파일 이름이 포함되어 있지 않기 때문에, Snyk 플랫폼이 원칙적으로 이 저장소의 모든 파일에 액세스할 수 있게 되기 때문입니다.
{% endhint %}

### **GitHub Enterprise와 함께 브로커의 추가 문제 해결**

* `docker logs <container id>`를 실행하여 오류를 확인하십시오. 여기서 컨테이너 ID는 GitHub Enterprise 브로커 컨테이너 ID입니다.
* GitHub Enterprise에 관련된 포트가 노출되어 있는지 확인하십시오.

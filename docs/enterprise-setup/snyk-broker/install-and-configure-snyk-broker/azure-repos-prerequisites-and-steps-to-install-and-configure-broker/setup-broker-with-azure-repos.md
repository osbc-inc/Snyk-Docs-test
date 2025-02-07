# Azure Repos - Docker를 사용하여 설치 및 구성

{% hint style="info" %}
**기능 가용성**

Snyk Azure Repos는 Azure DevOps/TFS 2020 이상에서만 사용할 수 있습니다.
{% endhint %}

설치하기 전에 [필수 조건](../../../../../ko/) 및 Docker를 사용하여 설치하는 일반적인 지침을 검토하십시오.

이 통합은 온프레미스 또는 클라우드 Azure Repos 배포와의 안전한 연결을 보장하는 데 유용합니다.

## Azure Repos와 함께 사용할 브로커 설정

Azure와 함께 브로커 클라이언트를 사용하려면 [Azure](https://azure.microsoft.com/ko-kr/services/devops/)에 대해 `docker pull snyk/broker:azure-repos`를 실행하십시오. 환경 변수의 정의를 확인하려면 [Snyk Broker를 위한 Azure Repos 환경 변수](azure-repos-environment-variables-for-snyk-broker.md)를 참조하십시오.

필요한 경우 [고급 구성 페이지](../advanced-configuration-for-snyk-broker-docker-installation/)로 이동하여 필요한 구성 변경(예: Azure Repos 인스턴스에서 개인 인증서를 사용 중인 경우 CA(인증서 기관)를 브로커 클라이언트 구성에 제공함) 또는 [프록시 지원 설정](../advanced-configuration-for-snyk-broker-docker-installation/proxy-support-with-docker.md)을 설정하십시오.

## Azure Repos용 브로커 클라이언트를 설정하는 Docker 실행 명령

**다음 명령을 복사하여** Azure 조직을 위해 완전히 구성된 브로커 클라이언트를 설정하여 오픈 소스, IaC, 컨테이너, 코드 파일 및 Snyk AppRisk 정보를 분석하십시오. **Snyk AppRisk**를 활성화하여 응용 프로그램 자산을 식별하고 모니터링하며 위험을 우선 순위로 처리하십시오.

{% hint style="info" %}
**기본 이외의 지역에 대한 다중 테넌트 설정**\
기본 이외의 지역에서 Snyk 브로커를 사용하려면 별도의 URL을 사용하는 환경 변수가 필요합니다. URL 및 예제에 대한 자세한 내용은 [지역 호스팅 및 데이터 잔단성, 브로커 URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

하나의 Azure 조직에 대해 더 많은 Azure 조직이 있는 경우 각각의 Azure 조직에 대해 브로커를 배포해야 합니다. Snyk AppRisk는 기본적으로 `false`로 설정되어 있습니다. `true`로 설정하여 활성화하십시오.

```bash
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=<비밀-브로커-토큰> \
           -e AZURE_REPOS_TOKEN=<비밀-azure-토큰> \
           -e AZURE_REPOS_ORG=<조직명> \
           -e AZURE_REPOS_HOST=<your.azure-server.domain.com (http/s 없음)> \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=<http://broker.url.example:8000 (dns/IP:포트)> \
           -e ACCEPT_IAC=tf,yaml,yml,json,tpl \
           -e ACCEPT_CODE=true \
           -e ACCEPT_APPRISK=true \
       snyk/broker:azure-repos
```

**Docker run 명령어 대신 대체 방법으로,** Azure Repos 통합을 위해 환경 변수를 재정의하기 위한 파생된 Docker 이미지를 사용할 수 있습니다. Azure Repos 통합을 위한 환경 변수를 위해 [파생된 Docker 이미지](../derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)를 참조하십시오.

## 브로커 클라이언트 컨테이너 시작 및 Azure Repos와의 연결 확인

브로커 클라이언트 컨테이너를 시작하도록 브로커 클라이언트 구성을 붙여넣으십시오.

컨테이너가 실행되면 Azure Repos 통합 페이지에서 Azure Repos와의 연결이 표시되며 `프로젝트 추가`를 할 수 있습니다.

## Azure Repos와의 브로커 문제 해결

* `docker logs <container id>`를 실행하여 오류를 확인하십시오. 여기서 `container id`는 Azure Repos 브로커 컨테이너 ID입니다.
* 관련 포트가 Azure Repos에 노출되어 있는지 확인하십시오.
* 로컬 경로의 `accept.json` 파일 및 해당 파일의 파일 권한이 올바르게 설정되어 있는지 확인하십시오.

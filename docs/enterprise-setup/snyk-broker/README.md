# Snyk 브로커

{% hint style="info" %}
**기능 가용성**

Snyk 브로커는 엔터프라이즈 플랜에서만 사용 가능합니다. 자세한 정보는 [요금제 및 가격 책정](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

Snyk 브로커는 Snyk와 특별 통합 사이에 프록시 역할을 하는 오픈 소스 도구로, [snyk.io](http://snyk.io/)가 코드에 액세스하여 스캔하고 결과를 반환할 수 있도록 합니다. 브로커와 함께 지원되는 SCM 통합은 , ,  (Dockerfile),  및 Snyk AppRisk입니다. 자세한 정보는 [Snyk 브로커 작동 방식](./#how-snyk-broker-works)을 참조하십시오.

## Snyk 브로커 다운로드 및 설치 방법

Snyk 브로커는 [GitHub](https://github.com/snyk/broker)에 호스팅되어 특정 통합을 위한 Docker 이미지 세트로 게시됩니다. Kubernetes를 사용하는 경우, Snyk는 Snyk 브로커를 배포하기 위한 [헬름 차트](https://github.com/snyk/snyk-broker-helm)를 제공합니다. 브로커를 배포하려면 통합을 설치하고 구성해야 합니다.

[Helm](install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md) 또는 [Docker](install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md)를 사용하여 설치하고 구성할 수 있습니다. Snyk는 Snyk 브로커 클라이언트를 실행하기 위해 Docker를 사용하거나 `npm install snyk-broker`를 실행할 수 있습니다. Snyk는 Snyk 브로커를 배포하는 가장 간단한 방법으로 Helm을 사용하는 것을 권장합니다.

## **Snyk 브로커와의 통합**

각 유형의 통합별로 설치하고 환경 변수를 구성하고, [Docker](install-and-configure-snyk-broker/install-and-configure-broker-using-docker.md) 및 [Helm](install-and-configure-snyk-broker/install-and-configure-broker-using-helm.md)에 대한 설명에 따라 구성합니다.

브로커와 함께 지원되는 통합 유형은 다음과 같습니다:

- 소스 코드 관리(SRC) 시스템([GitHub](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/), [GitHub Enterprise](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/), [BitBucket Server/Data Center](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/), [GitLab](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/), [Azure Repos](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/))
  - 인터넷에 접근할 수 없는 SCM
  - 증가된 데이터 보안을 위해 Snyk 활동을 볼 수 있게 해주는 공개적으로 접근 가능한 SCM
- 내부 [Jira](install-and-configure-snyk-broker/jira-prerequisites-and-steps-to-install-and-configure-broker/), [JFrog Artifactory](install-and-configure-snyk-broker/artifactory-repository-install-and-configure-broker/), 또는 Nexus 설치
- 네트워크 제한된 [컨테이너 레지스트리](snyk-broker-container-registry-agent/)
- 프라이빗 Git 기반 저장소에서의 [인프라 구성(IaC) 파일](snyk-broker-infrastructure-as-code-detection/)

각 통합 및 컨테이너 레지스트리 에이전트에 대한 [파생된 Docker 이미지 사용](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/snyk-broker-set-up-examples/derived-docker-images-for-broker-client-integrations-and-container-registry-agent) 및 컨테이너 레지스트리 에이전트를 이용할 수 있습니다.

설치에 필요한 고급 구성에 대한 정보는 [Snyk 브로커 Docker 설치에 대한 고급 구성](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-docker/advanced-configuration-for-snyk-broker-docker-installation) 및 [Helm 차트 설치를 위한 고급 설정](https://docs.snyk.io/snyk-admin/snyk-broker/install-and-configure-broker-using-helm/advanced-setup-for-helm-chart-installation)을 참조하십시오.

## Snyk 브로커를 사용하여 코드 스캔하기

Snyk 브로커를 사용하여 Snyk 오픈 소스를 사용하려면 브로커 서버 및 브로커 클라이언트 구성 요소만 필요합니다. 브로커 클라이언트는 특정 Git 서비스에 구성된 각각의 Docker 이미지 세트로 게시됩니다. [Snyk 브로커와의 통합](./#integrations-with-snyk-broker) 섹션의 링크를 따라 환경 변수를 사용하여 각 통합 유형을 구성합니다.

Snyk 브로커를 사용하여 다른 유형의 코드를 스캔하려면 컴포넌트나 설정을 추가하고 브로커 클라이언트 설정에 매개변수를 추가해야 합니다:

- **** – 환경 변수 `ACCEPT_CODE=true`를 추가하여 [저장소의 Git 클론을 실행](git-clone-through-broker.md)할 수 있도록 브로커 액세스 권한을 부여합니다.
- **** – [컨테이너 레지스트리 에이전트](snyk-broker-container-registry-agent/)를 추가하여 네트워크 제한된 컨테이너 레지스트리 연결을 가능하게 하고 컨테이너 이미지를 분석합니다. [Docker로 설치하는 방법](snyk-broker-container-registry-agent/) 및 [헬름으로 설치하는 방법](snyk-broker-container-registry-agent/install-broker-for-container-registry-agent-using-helm.md)에 대한 지침이 있습니다.
- **Snyk ** – Snyk 브로커를 통해 Terraform, CloudFormation 및 Kubernetes 구성 파일을 감지하고 분석하기 위해 [`accept.json` 파일에 추가 매개변수](snyk-broker-infrastructure-as-code-detection/)를 구성합니다.

## Snyk 브로커 작동 방식

Snyk 브로커는 Snyk 제품을 인터넷에서 액세스할 수 없는 자체 호스팅 통합에 연결하도록 설계되었습니다. Snyk 브로커는 또한 다음을 수행할 수 있도록 합니다:

- Snyk가 액세스할 파일과 Snyk가 수행할 작업을 제한하여 Snyk로부터 네트워크에 대한 액세스 제어
- 통합을 대상으로 한 고정된 사설 IP를 관리

Snyk 브로커에는 서버 및 클라이언트로이라는 기본 컴포넌트가 포함되어 있으며 모든 통합에서 동일합니다. 브로커 서버는 Snyk SaaS 백엔드에서 실행되며 Snyk에 의해 제공되며 설치가 필요하지 않습니다. 브로커 클라이언트는 귀하의 인프라에 배포되는 [Docker 이미지](https://hub.docker.com/r/snyk/broker/)입니다. 자세한 정보는 [Snyk 브로커의 구성 요소](components-of-snyk-broker.md) 및 [Snyk 브로커와의 연결](connections-with-snyk-broker.md)을 참조하십시오.

배포를 위한 Snyk 브로커 준비에 대한 정보는 선행 조건, 컴포넌트 선택, 네트워크 구성 및 자격 증명에 대해 알아볼 수 있는 [Snyk 브로커를 준비](prepare-snyk-broker-for-deployment.md)하십시오.

## Snyk 브로커에 관한 일반 질문

**Snyk 브로커는 얼마나 자주 업데이트되나요?**\
Snyk 브로커는 새로운 기능이 제공될 때마다 업데이트되며 수정 사항이 있을 때도 업데이트됩니다.

**Snyk 브로커는 얼마나 자주 취약점을 체크하나요?**\
Snyk 브로커 애플리케이션 및 이미지는 매일 취약점을 테스트합니다.

**취약점을 고치는 SLA는 무엇인가요?**\
공개 이미지의 심각한 취약점을 수정하는 데 14일의 SLA와 중요한 취약점을 수정하는 데 5일의 SLA가 있습니다.
# 배포를 위한 Snyk Broker 준비하기

{% hint style="info" %}
Windows에서 Snyk 브로커를 사용하는 것은 지원되지 않습니다. Snyk는 Windows 사용자가 리눅스를 사용하여 브로커를 배포하는 것을 권장합니다.
{% endhint %}

## Snyk 브로커를 위한 선행 조건

시스템 기본값 외의 환경(지역)에 대해 브로커를 설정할 때, 인증하기 전에 특정한 [브로커 URL](../../working-with-snyk/regional-hosting-and-data-residency.md#broker-urls)을 환경 변수로 설정해야 합니다.\
예시: `-e BROKER_SERVER_URL=https://broker.eu.snyk.io`

자세한 내용은 [지역 호스팅 및 데이터 보관](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency)을 참조하십시오.

다음은 모든 환경에서 Snyk 브로커를 사용하기 위한 선행 조건입니다:

* 클라이언트 컴퓨터 시스템 요구 사항: 1개의 CPU, 256MB RAM
* 네트워크 접근: 네트워크에서 [https://broker.snyk.io](https://broker.snyk.io) 및 [https://api.snyk.io](https://api.snyk.io)(또는 해당 지역과 동등한)로의 아웃바운드 TLS(443) 및 네트워크에 설치된 모든 방화벽에서 허용되어야 함
* Snyk 계정
* Snyk API를 사용하여 스스로 브로커 통합을 활성화하거나 [Snyk 지원팀](https://support.snyk.io)에 문의하여 활성화
* 브로커 토큰이라고 불리는 고유 UUID 토큰 필요. [Snyk 브로커를 위한 자격 증명 생성](prepare-snyk-broker-for-deployment.md#generate-credentials-in-the-target-application-for-snyk-broker)을 참조하십시오.
* SCM 토큰 또는 암호. 각 SCM에 대한 토큰 획득 방법에 대한 정보는 [통합 문서](../../integrate-with-snyk/)를 참조하십시오. Snyk 브로커는 mTLS 방법으로 인증을 지원하지 않음.
* Docker가 설치된 Docker Hub에서 이미지를 끌어오기 위해 구성되어 있어야 함.

## Snyk 브로커 설치를 위한 호스트 준비

Snyk는 각 통합을 위해 최소한 두 개의 별도의 브로커 클라이언트 인스턴스를 설정하거나 Kubernetes 시스템을 사용하여 설치하길 권장합니다. 이렇게 하면 항상 적어도 두 개의 인스턴스가 가동되어 중복성이 유지됩니다.

## Snyk 브로커 사용을 위한 네트워크 구성

프록시 서버를 사용하는 경우 다음과 같이 Broker Client의 내부 및 외부 액세스를 허용하기 위해 프록시 및 모든 방화벽을 구성해야 합니다:

* 환경에서 실행 중인 Broker Client에서 [https://broker.snyk.io](https://broker.snyk.io) (또는 [https://broker.eu.snyk.io](https://broker.eu.snyk.io) / [https://broker.au.snyk.io](https://broker.au.snyk.io)) 및 [https://api.snyk.io](https://api.snyk.io) (또는 [https://api.eu.snyk.io](https://api.eu.snyk.io) / [https://api.au.snyk.io](https://api.au.snyk.io))로의 아웃바운드 연결이 443번 포트로 허용되어야 함.
* 통합(SCM, CR)에서 Broker Client로의 인바운드 접근을 허용하는 내부 연결은 BROKER\_CLIENT\_URL 및 설정한 포트(일반적으로 8000번)에서 외부에서 인입되는 것이 아님.

Snyk 브로커 서버 측에서 시작된 트래픽은 항상 최신의 브로커 연결을 사용합니다. Snyk 측에서의 모든 활동과 같이, 주기적인 테스트에 의해 추진되는 트래픽은 한 번에 한 번에 한 개의 복제본에서만 발생합니다. Snyk 활동의 양은 리포지토리나 Jira 항목에서의 활동에 비례합니다. 해당 활동은 웹훅을 생성하며, 이는 모든 복제본에 분산됩니다.

## **브로커 배포 구성 요소 정의**

배포를 위해 필요한 구성 요소를 이해하기 위해 다음 사항을 고려하십시오:

* 어떤 서비스에 브로커를 연결하려고 합니까?
  * GitHub, Jira, Bitbucket, Harbor 또는 다른 서비스
  * [Snyk 브로커 설치 및 설정](install-and-configure-snyk-broker/)를 참조하십시오.
* Infrastructure as Code 파일을 감지할 계획이 있습니까?
  * 배포에 `-e ACCEPT_IAC`라는 환경 변수나 사용자 정의 allowlist `accept.json` 파일을 추가해야 함.
  * [Snyk 브로커 - Infrastructure as Code 검출](snyk-broker-infrastructure-as-code-detection/)을 참조하십시오.
* 취약점을 감지할 계획이 있습니까?
* 브로커가 귀하의 리포지토리를 Git 클론할 수 있도록 권한 부여
* 이를 위해 환경 변수를 추가하십시오: `ACCEPT_CODE=true.`
* 컨테이너 레지스트리에 연결할 계획이 있습니까?
  * Snyk 브로커 컨테이너 레지스트리 에이전트와 함께 추가 에이전트를 배포해야 함.
  * [Snyk 브로커 컨테이너 레지스트리 에이전트](snyk-broker-container-registry-agent/)를 참조하십시오.

각 통합에는 특정한 브로커 토큰이 할당됩니다. 취약점을 분석하고 컨테이너 레지스트리에 연결하는 통합은 다음을 갖추고 있습니다:

* 추가 환경 변수 `-e ACCEPT_CODE`를 가진 SCM용 브로커 한 개
* 컨테이너 레지스트리용 브로커 한 개 및 브로커 컨테이너 레지스트리 에이전트 한 개

## Snyk 브로커를 위한 대상 응용프로그램에서 자격 증명 생성

{% hint style="info" %}
Snyk 모든 API 토큰과 자격 증명을 90일마다 반복해서 바꾸는 것을 권장합니다.

브로커의 첫 배포에는 Snyk 계정 팀과 협력해야 합니다.
{% endhint %}

브로커의 대상 응용프로그램에서 자격 증명을 생성한 후, 시작할 때 필요한 환경 변수를 구성하십시오.

브로커 토큰이 필요하며, Snyk 브로커를 사용하려면 반드시 생성되어 있어야 합니다.

코드 리포지토리(SCM) 통합을 위해 브로커 토큰은 API를 사용하거나 [Snyk 지원팀](https://support.snyk.io)에 문의하여 생성할 수 있습니다.

1. 엔드포인트 [기존 통합 업데이트](https://snyk.docs.apiary.io/#reference/integrations/integration/update-existing-integration)에서 "기존 통합을 위한 브로커 설정" 아래의 예제를 따르거나 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.
2. SCM 통합의 특정 브로커 토큰이 Snyk 웹 UI에서 생성되어 있는지 확인하기 위해 해당 통합을 위해 **설정** > **통합**을 선택하십시오.

[Artifactory Repository](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/) 및 [Nexus Repository Manager](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) 브로커화된 인스턴스 또는 [Jira](install-and-configure-snyk-broker/jira-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-jira.md) 통합의 경우, Snyk UI에서 브로커 토큰을 생성하거나 [Snyk 지원팀](https://support.snyk.io)에 문의할 수 있습니다.

1. 해당 통합을 위해 **설정** > **통합**을 선택하여 브로커 토큰을 생성하십시오.
2. 브로커 토큰이 생성된 후, 해당 통합에서, 스크린에 나타나는 알림이 "Could not connect to..."인 것을 확인하십시오. 이는 아직 클라이언트를 설치하거나 구성하지 않았기 때문입니다.
3. UI에서 브로커 토큰을 복사하여 클라이언트를 설치할 때 사용하십시오.

## 여러 조직을 대상으로 브로커 활성화

동일한 Git 서비스를 사용하여 Snyk의 여러 조직에서 동일한 브로커 토큰을 사용할 수 있습니다. 이를 위해 조직을 위한 토큰을 만든 다음 원본을 기반으로 새로운 조직을 만들어야 합니다. 이렇게 하면 토큰이 복제되고 새로운 조직에 대해 브로커를 활성화할 수 있습니다.

기존 조직에 대해 이 작업을 수행하려면 특정 통합을 복제하는 [통합 복제 (설정 및 자격 증명 포함)](../../snyk-api/reference/integrations-v1.md#org-orgid-integrations-integrationid-clone) 엔드포인트를 사용하여 특정 통합을 복제할 수 있습니다. 이는 브로커 토큰을 포함하는 특정 통합을 복제합니다.

그렇지 않으면 각 통합 및 조직은 고유한 브로커 토큰을 가지기 때문에 조직을 위해 새로운 브로커 토큰을 생성해야 합니다.

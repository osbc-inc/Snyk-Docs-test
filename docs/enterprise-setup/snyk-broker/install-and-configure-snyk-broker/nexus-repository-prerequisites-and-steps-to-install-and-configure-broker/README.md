# Nexus Repository - Broker 설치 및 구성을 위한 사전 조건 및 단계

{% hint style="info" %}
**기능 가용성**

Nexus Repository Manager와의 통합은 엔터프라이즈 플랜에서만 가능합니다. 자세한 정보는 [요금제 및 가격정보](https://snyk.io/plans/)를 참조하십시오.
{% endhint %}

설치하기 전에, 사용할 설치 방법에 대한 일반 지침을 검토하십시오. [Helm](../install-and-configure-broker-using-helm.md) 또는 [Docker](../install-and-configure-broker-using-docker.md).

다음은 **사전 조건**입니다.

Snyk Nexus Repository Broker를 설치하기 전에, Snyk 계정 팀에게 Broker 토큰을 제공하도록 요청하거나 Snyk 웹 UI에서 생성하도록 요청하십시오.

Docker를 사용하거나 Docker Linux 컨테이너를 실행할 수 있는 방법이 있어야 합니다.\
일부 Docker Windows 배포는 Windows 컨테이너만 실행합니다. 배포가 Linux 컨테이너를 실행할 수 있는지 확인하십시오.

편의를 위해, Broker 토큰을 얻거나 생성하는 지침이 있습니다. 완료 후 [Docker](set-up-snyk-broker-with-nexus-repository-manager.md) 또는 [Helm](nexus-repository-install-and-configure-using-helm.md)를 사용하여 설치하는 단계를 계속하십시오.

## Nexus 통합을 위한 Broker 토큰 얻기

1. **설정 > 통합 > 패키지 저장소 > Nexus**로 이동합니다.
2. Nexus를 구성할 수 있는 화면이 나타나는지 확인하십시오.

<figure><img src="../../../../.gitbook/assets/Screenshot 2022-07-15 at 15.15.11.png" alt="Nexus 구성"><figcaption><p>Nexus 구성</p></figcaption></figure>

{% hint style="info" %}
**Snyk Broker** 스위치가 보이지 않는 경우, 필요한 권한이 없으며 공개적으로 접근 가능한 인스턴스만 추가할 수 있습니다.

사설 레지스트리를 추가하려면 [Snyk 지원팀](https://support.snyk.io)에 요청하십시오.
{% endhint %}

사설 레지스트리를 추가할 권한이 있으면, [웹 UI에서 Broker 토큰을 생성](./#generate-a-broker-token-from-the-web-ui)하는 지침을 계속하세요.

## 웹 UI에서 Broker 토큰 생성

1. Nexus 통합 설정에서 **Snyk Broker 켜기/끄기** 스위치를 **켜**로 이동하여 Broker 토큰을 생성하는 양식을 표시합니다.
2. **생성 및 저장**을 선택합니다.
3. 생성된 토큰을 복사하여 Broker 클라이언트 설정 시 사용하십시오.
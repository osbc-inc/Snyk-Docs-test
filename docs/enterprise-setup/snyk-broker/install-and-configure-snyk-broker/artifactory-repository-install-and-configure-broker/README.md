# Artifactory Repository - Broker를 설치하고 구성하는 데 필요한 사전 조건 및 단계

{% hint style="info" %}
**기능 가용성**

Artifactory Repository와의 통합은 Enterprise 플랜에서만 사용할 수 있습니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans/)을 확인하세요.
{% endhint %}

설치하기 전에, 사용할 설치 방법에 대한 일반 지침을 [Helm](../install-and-configure-broker-using-helm.md) 또는 [Docker](../install-and-configure-broker-using-docker.md) 에서 검토하세요.

다음은 **사전 조건**입니다.

Snyk Artifactory Repository Broker를 설치하기 전에, Snyk 아카운트 팀에게 Broker 토큰을 제공하도록 요청하거나 Snyk 웹 UI에서 생성하십시오.

Docker를 가지고 있거나 Docker Linux 컨테이너를 실행할 수 있는 방법이 있어야 합니다. 일부 Windows용 Docker 배포판은 Windows 컨테이너만 실행합니다. 배포가 Linux 컨테이너를 실행할 수 있는지 확인하세요.

편의를 위해 Broker 토큰을 얻거나 생성하는 방법이 아래에 나와 있습니다. 완료되면 [Docker](set-up-snyk-broker-with-artifactory-repository.md) 또는 [Helm](artifactory-repository-install-and-configure-using-helm.md)을 사용하여 설치 단계를 계속하십시오.

## Artifactory Repository 설정을 위한 Broker 토큰 얻기

1. **Settings** > **Integrations > Package Repositories > Artifactory** 로 이동합니다.
2. Artifactory 인스턴스의 URL을 입력하고, 이 URL은 반드시 **/artifactory**로 끝나야 합니다.
3. 사용자 이름과 암호를 입력합니다.
4. **Save**를 선택합니다.

<figure><img src="../../../../.gitbook/assets/screenshot_2020-04-17_at_14.38.12.png" alt="Artifactory 통합 설정"><figcaption><p>Artifactory Repository 설정</p></figcaption></figure>

{% hint style="info %}
**Snyk Broker** 온/오프 스위치를 볼 수 없는 경우, 필요한 권한이 없으므로 공개적으로 접근 가능한 인스턴스만 추가할 수 있습니다.

개인 레지스트리를 추가하려면 [Snyk 지원팀](https://support.snyk.io)에 요청하세요.
{% endhint %}

개인 레지스트리를 추가할 권한이 있는 경우, [웹 UI에서 Broker 토큰을 생성하는 지침](./#generate-a-broker-token-from-the-web-ui)을 계속하세요.

## 웹 UI에서 Broker 토큰 생성

1. Artifactory 통합 설정에서 **Snyk Broker on/off** 스위치를 **on**으로 이동하여 Broker 토큰을 생성하는 양식을 표시합니다.
2. **Generate and Save**를 선택합니다.
3. 생성된 토큰을 복사하여 Broker Client를 설정할 때 사용하세요.
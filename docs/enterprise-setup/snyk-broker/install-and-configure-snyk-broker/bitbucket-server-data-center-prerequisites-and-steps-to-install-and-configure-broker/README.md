# Bitbucket Server/Data Center - Broker 설치 및 구성을 위한 사전 조건 및 단계

{% hint style="info" %}
**기능 가용성**

 PR Checks는 Bitbucket DC/Server 버전 7.0 이상에서만 사용할 수 있습니다.
{% endhint %}

설치 전, 사용할 설치 방법([Helm](../install-and-configure-broker-using-helm.md) 또는 [Docker](../install-and-configure-broker-using-docker.md))에 대한 **일반 지침**을 확인해주세요.

다음은 **사전 조건**입니다.

Bitbucket Server/Data Center Broker를 설치하기 전에 Snyk 계정 팀에게 Broker 토큰을 제공해 달라고 요청해야 합니다.&#x20;

Docker를 실행하거나 Docker Linux 컨테이너를 실행할 방법이 있어야 합니다. 일부 Windows용 Docker 배포는 Windows 컨테이너만 실행합니다. 배포가 Linux 컨테이너를 실행할 수 있는지 확인하세요.

[Docker](data-center.md) 또는 [Helm](bitbucket-server-data-center-install-and-configure-using-helm.md)를 사용하여 설치하는 단계를 계속하세요.
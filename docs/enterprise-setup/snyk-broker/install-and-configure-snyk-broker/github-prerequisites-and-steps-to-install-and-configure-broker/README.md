# GitHub - Broker를 설치하고 구성하기 위한 사전 조건 및 단계

설치하기 전에, 계획 중인 설치 방법에 대한 **일반 지침**을 [Helm](../install-and-configure-broker-using-helm.md) 또는 [Docker](../install-and-configure-broker-using-docker.md)의 설치 방법에 대해 검토하십시오.

다음은 **필수 조건**입니다.

Snyk GitHub Broker를 설치하기 전에, [필수 권한](../../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md#required-permissions-scope-for-the-github-integration)이 부여된 GitHub 서비스 계정 토큰을 구성해야 합니다. Snyk 웹 UI를 통해 트리거된 작업 및 자동 작업은 해당 토큰이 Broker로 구성된 GitHub 서비스 계정을 통해 수행됩니다.

Docker를 설치하거나 Docker Linux 컨테이너를 실행할 수 있는 방법이 있어야 합니다. Windows용 Docker 배포 중에는 Windows 컨테이너만 실행하는 경우도 있습니다. 사용 중인 배포가 Linux 컨테이너를 실행할 수 있는지 확인하십시오.

[Docker](github-install-and-configure-using-docker.md) 또는 [Helm](github-install-and-configure-using-helm.md)을 사용하여 설치하는 단계를 계속하세요.
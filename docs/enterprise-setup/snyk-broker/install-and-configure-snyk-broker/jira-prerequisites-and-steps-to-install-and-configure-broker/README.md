# Jira - Broker 설치 및 구성을 위한 선행 조건 및 단계

설치 전에, 사용할 설치 방법에 대한 일반 지침을 **검토**하세요. [Helm](../install-and-configure-broker-using-helm.md) 또는 [Docker](../install-and-configure-broker-using-docker.md).

다음은 **선행 조건**입니다.

Snyk Jira Broker를 설치하기 전에, Snyk 계정 팀에게 Broker 토큰을 제공하도록 요청하거나 Snyk 웹 UI에서 생성하세요.

Docker를 사용하거나 Docker Linux 컨테이너를 실행할 수 있는 방법이 있어야 합니다.\
일부 Windows용 Docker 배포는 Windows 컨테이너만 실행합니다. 배포가 Linux 컨테이너를 실행할 수 있는지 확인하세요.

웹 UI에서 Broker 토큰을 생성하는 방법은 다음과 같습니다:

1. **Settings** > **Integrations** > **Jira** > **For installation of Jira within a private network click here**로 이동합니다.
2. Jira를 위한 Broker 토큰을 생성하고 확인하려면 **Generate**를 클릭한 후 **Show**를 클릭합니다.

[Docker](setup-broker-with-jira.md) 또는 [Helm](jira-install-and-configure-using-helm.md)을 사용하여 설치하는 단계를 **진행**하세요.
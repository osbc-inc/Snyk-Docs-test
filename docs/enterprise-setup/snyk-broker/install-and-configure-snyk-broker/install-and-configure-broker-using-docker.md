# 도커를 사용하여 브로커 설치 및 구성

설치를 시작하기 전에 [Prepare Snyk Broker for deployment](../prepare-snyk-broker-for-deployment.md) 페이지의 [Prerequisites](../prepare-snyk-broker-for-deployment.md#prerequisites-for-snyk-broker) 및 다른 정보를 참조하십시오.

**Kubernetes를 사용하는 경우**, Snyk는 [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm)를 사용하여 Snyk 브로커를 설치하는 것을 권장합니다. 자세한 내용은 [Helm을 사용하여 브로커 설치 및 구성](install-and-configure-broker-using-helm.md)을 참조하십시오.

다른 환경에서는, Snyk가 제공하는 [Docker 이미지](https://github.com/snyk/broker)를 사용하여 Snyk 브로커를 설치할 수 있습니다. 이곳에 나열된 페이지들은 Docker를 사용하여 Snyk 브로커 클라이언트 통합을 설정하는 방법을 설명합니다. 

## Docker를 사용하여 설치

다음 페이지들에서 특별한 통합을 설치하는 방법에 대해 설명합니다.

* [GitHub](github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-docker.md)
* [GitHub Enterprise](github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md)
* [Bitbucket Server/Data Centre](bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md)
* [Gitlab](gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md)
* [Azure Repos](azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md)
* [JFrog Artifactory Repository](artifactory-repository-install-and-configure-broker/set-up-snyk-broker-with-artifactory-repository.md)
* [Nexus Repository Manager](nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/set-up-snyk-broker-with-nexus-repository-manager.md)
* [Jira](jira-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-jira.md)
* [Snyk Broker - Container Registry Agent](../snyk-broker-container-registry-agent/) (Container Registries에 연결하기 위해서 필요)
* [브로커 클라이언트 통합 및 컨테이너 레지스트리 에이전트용 파생 도커 이미지](derived-docker-images-for-broker-client-integrations-and-container-registry-agent.md)

환경 변수를 사용하여 Docker 이미지의 구성을 사용자 정의할 수 있습니다. 이러한 이유로, 적절한 구성 및 신뢰성을 보장하기 위해 다른 통합 유형을 위한 별도로 여러 개의 브로커 클라이언트 인스턴스를 설치하십시오.

브로커가 작동 중인지 확인하려면, 연결된 것을 확인하는 확인 메시지를 볼 수 있는 [Snyk 웹 UI](https://app.snyk.io)에서 브로커된 통합의 설정을 확인합니다. 연결된 후 프로젝트를 가져오기 시작할 수 있습니다.

## Docker를 사용하여 고급 구성

Docker를 사용하여 설치하는 경우, 필요에 따라 [Snyk 브로커 Docker 설치에 대한 고급 구성](advanced-configuration-for-snyk-broker-docker-installation/) 페이지의 지침에 따릅니다.
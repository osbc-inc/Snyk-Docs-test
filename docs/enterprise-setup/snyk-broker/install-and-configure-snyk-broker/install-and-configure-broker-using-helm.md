# Helm을 사용하여 Broker 설치 및 구성

설치를 시작하기 전에 [Snyk Broker를 배포하기 위한 사전 조건](../prepare-snyk-broker-for-deployment.md#prerequisites-for-snyk-broker) 및 페이지 [Prepare Snyk Broker for deployment](../prepare-snyk-broker-for-deployment.md)에서 다른 정보를 확인하세요.

**만약 Kubernetes를 사용 중이라면**, Snyk는 [**Broker Helm Chart**](https://github.com/snyk/snyk-broker-helm)를 사용하여 **Snyk Broker를 설치**할 것을 권장합니다.

**기타 환경을 사용하고 있다면**, Snyk가 제공하는 [Docker 이미지](https://github.com/snyk/broker)를 사용하여 Snyk Broker를 설치할 수 있습니다. 자세한 내용은 [Docker를 사용하여 Broker 설치 및 구성](install-and-configure-broker-using-docker.md)을 참조하세요.

{% hint style="info" %}
**기본값이 아닌 지역용 다중 테넌트 설정**\
기본값이 아닌 지역에서 Snyk Broker를 설정할 때는 특정 URL이 필요한 추가 환경 변수가 필요합니다. URL 및 예시에 대해서는 [Regional hosting and data residency, Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하세요.
{% endhint %}

## Snyk Broker Helm Chart를 사용하여 설치

Helm 차트는 연결성을 관리하지 않으므로 Kubernetes 클러스터에서 [ingress 관리](advanced-configuration-for-helm-chart-installation/ingress-options-with-snyk-broker-helm-installation.md)에 책임이 있습니다.

이 차트를 사용하려면 먼저 리포지토리를 추가하여 Snyk Broker Helm Chart를 추가해야 합니다:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

그런 다음 각 SCM, 레지스트리 또는 Jira에 대한 환경 변수를 사용자 정의하는 명령을 실행하십시오. 자세한 내용은 다음 페이지에서 설명합니다:

* [GitHub](github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-helm.md) `scmType`: `github-com`
* [GitHub Enterprise](github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md) `scmType`: `github-enterprise`
* [Bitbucket Server/Data Center](bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md) `scmType`: `bitbucket-server`
* [GitLab](gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md) `scmType`: `gitlab`
* [Azure Repos](azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md) `scmType`: `azure-repos`
* [JFrog Artifactory](artifactory-repository-install-and-configure-broker/artifactory-repository-install-and-configure-using-helm.md) `scmType`: `artifactory`
* [Nexus 3](nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/nexus-repository-install-and-configure-using-helm.md) `scmType`: `nexus`
* [Nexus 2](nexus-repository-prerequisites-and-steps-to-install-and-configure-broker/nexus-repository-install-and-configure-using-helm.md) `scmType`: `nexus2`
* [Jira](jira-prerequisites-and-steps-to-install-and-configure-broker/jira-install-and-configure-using-helm.md) `scmType`: `jira`

`scmType`은 시스템 유형을 지정합니다. JFrog 및 Nexus의 경우 이는 아티팩트 리포지터리이고, Jira의 경우 티켓 관리 시스템입니다.

각 SCM, 레지스트리 또는 Jira에 대한 명령을 실행하면 `snyk-broker`라는 네임스페이스가 생성됩니다. 기존 네임스페이스에 배포하려면 `-n` 매개변수를 조정하고 `--create-namespace` 매개변수를 삭제하세요. 또한 [동일한 네임스페이스에 여러 Broker 배포](advanced-configuration-for-helm-chart-installation/deploying-multiple-brokers-in-the-same-namespace.md)도 참조하세요.

버전 2.0.0부터 생성된 모든 객체에는 릴리스 이름을 기반으로 한 접미사가 있어 동일한 네임스페이스에 여러 Broker를 만들 수 있습니다. 호환성을 유지하기 위해 2.1.0에서는 `disableSuffixes` 플래그를 소개하여 `--set disableSuffixes=true`를 추가하여 1.x.x 동작으로 되돌릴 수 있습니다.

[Snyk Broker - 컨테이너 레지스트리 에이전트](../snyk-broker-container-registry-agent/)를 설치하기 위한 추가 명령이 있습니다. 이 에이전트는 컨테이너 레지스트리에 연결하기 위해 필요하며 `scmType`: `container-registry-agent`입니다.

브로커가 실행 중인지 확인하려면 [Snyk 웹 UI](https://app.snyk.io)에서 브로커 통합에 대한 설정을 확인하여 연결된 메시지가 표시되는지 확인할 수 있습니다. 연결되면 프로젝트를 가져오기 시작할 수 있습니다.

## Helm을 사용한 고급 구성

Helm 명령으로 원하는 환경 변수를 전달할 수 있습니다. 자세한 내용은 [Broker Helm Chart를 위한 사용자 정의 추가 옵션](advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)을 참조하세요.

예를 들어, BROKER\_CLIENT\_VALIDATION\_URL을 Helm 차트를 사용하여 전달하려면 추가 매개변수는 다음과 같을 것입니다:

`--set env[0].name=BROKER_CLIENT_VALIDATION_URL \`\
`--set env[0].value=whatever_value_makes_sense`

추가 매개변수는 다음과 같을 것입니다:

`--set env[1].name=MY_OTHER_ENV_VAR \`\
`--set env[1].value="other env with spaces" \`\
`--set env[2].name=THIRD_ENV_VAR \`\
`--set env[2].value=and_so_on`

필요에 따라 구성 변경을 수행하려면 [Helm Chart 설치를 위한 고급 구성](advanced-configuration-for-helm-chart-installation/) 지침을 따르세요.
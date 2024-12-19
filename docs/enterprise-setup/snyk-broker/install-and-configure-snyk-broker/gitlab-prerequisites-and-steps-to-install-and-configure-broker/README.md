# GitLab - Broker를 설치하고 구성하기 위한 선행 조건 및 단계

설치하기 전에, 계획 중인 설치 방법에 대한 **일반 지침**을 [Helm](../install-and-configure-broker-using-helm.md) 또는 [Docker](../install-and-configure-broker-using-docker.md)에 대해 검토하십시오.

**선행 조건**은 다음과 같습니다.

- GitLab 관리자이어야 합니다. GitLab 설정에서 아웃바운드 요청을 필터링하여 [웹훅 및 통합에서 로컬 네트워크로의 요청을 허용](https://docs.gitlab.com/ee/security/webhooks.html#allow-requests-to-the-local-network-from-webhooks-and-integrations)하고, 특정 IP 주소 및 도메인으로의 아웃바운드 요청을 허용](https://docs.gitlab.com/ee/security/webhooks.html#allow-outbound-requests-to-certain-ip-addresses-and-domains)합니다.

- Snyk GitLab Broker를 설치하기 전에, Snyk 계정 팀에게 Broker 토큰을 제공해 달라고 요청하십시오.

- Snyk와의 통합을 위해 GitLab 권한이 올바른지 확인하십시오. 자세한 내용은 [Snyk GitLab 통합](../../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab.md)을 참조하십시오.

- Docker를 보유하거나 Docker Linux 컨테이너를 실행할 수 있는 방법이 있어야 합니다. Windows를 위한 일부 Docker 배포는 Windows 컨테이너만 실행합니다. 배포가 Linux 컨테이너를 실행할 수 있는지 확인하십시오.

- [Docker](setup-broker-with-gitlab.md) 또는 [Helm](gitlab-install-and-configure-using-helm.md)을 사용하여 설치하는 단계를 계속 진행하십시오.
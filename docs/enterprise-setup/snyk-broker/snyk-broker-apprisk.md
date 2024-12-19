# Snyk 브로커 - AppRisk

만약 SCM 또는 써드파티 인스턴스가 공개적으로 접근이 불가능하다면, Snyk 브로커가 필요합니다. Snyk 브로커는 Docker 또는 Helm을 사용하여 설치하고 구성할 수 있습니다. Snyk AppRisk를 위한 최소 지원 브로커 버전은 [4.171.0](https://github.com/snyk/broker/releases/tag/v4.171.0) 입니다.

Snyk AppRisk를 위해 브로커를 활성화하려면 설치 명령어에서 `APPRISK` 환경 변수를 `true`로 설정하세요. Docker의 경우 `ACCEPT_APPRISK=true`, Helm의 경우 `--set enableAppRisk=true`로 설정합니다.

Snyk AppRisk 통합을 위한 Snyk 브로커 토큰을 반드시 보유해야 합니다. 필요한 토큰은 Snyk 지원팀에서 제공하거나 다음 지침에 따라 직접 생성할 수 있습니다:

* [Snyk 브로커용 토큰 얻기](snyk-broker-code-agent/install-snyk-broker-code-agent-using-docker/obtain-the-required-tokens-for-setup.md#obtain-your-broker-token-for-snyk-broker-code-agent) 페이지의 지침을 따라 브로커 토큰을 생성합니다.
* 생성된 브로커 토큰을 통합 허브의 통합 설정 메뉴에 복사하여 붙여넣기합니다.

## SCM 통합

* GitHub - Snyk 브로커 설치 및 구성&#x20;
  * [Docker 사용](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-docker.md#docker-run-command-to-set-up-a-broker-client-for-github)
  * [Helm 사용](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-helm.md)
  * [환경 변수](install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-environment-variables-for-snyk-broker.md)
* GitHub Enterprise - Snyk 브로커 설치 및 구성:
  * [Docker 사용](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md#docker-run-command-to-set-up-a-broker-client-for-github-enterprise)
  * [Helm 사용](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md)
  * [환경 변수](install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-environment-variables-for-snyk-broker.md)
* BitBucket - Snyk 브로커 설치 및 구성:
  * [Docker 사용](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md#docker-run-command-to-set-up-a-broker-client-for-bitbucket)
  * [Helm 사용](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md)
  * [환경 변수](install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md)
* GitLab - Snyk 브로커 설치 및 구성:
  * [Docker 사용](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md#docker-run-command-to-set-up-a-broker-client-for-gitlab)
  * [Helm 사용](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md)
  * [환경 변수](install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md)
* Azure - Snyk 브로커 설치 및 구성:
  * [Docker 사용](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md#docker-run-command-to-set-up-a-broker-client-for-azure-repos)
  * [Helm 사용](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md)
  * [환경 변수](install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-environment-variables-for-snyk-broker.md)

각 통합에 대한 허용된 엔드포인트 목록이 포함된 모든 업데이트된 `.json` 파일을 [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk)에서 찾을 수 있습니다.

브로커 설치가 완료된 후, Snyk Apprisk에 추가하려는 통합에 대한 브로커 토큰을 반드시 얻어야 합니다. 브로커 토큰은 조직 통합 일반 설정에서 찾을 수 있으며, 예를 들어 GitHub, GitLab 등의 통합 유형에서 가능합니다. 써드파티 통합의 경우, 다음 섹션을 참조하세요.

<figure><img src="../../.gitbook/assets/Screenshot 2024-11-01 at 1.05.55 PM.png" alt="&#x22;&#x22;"><figcaption><p>GitLab 통합 일반 설정에 있는 브로커 토큰</p></figcaption></figure>

## 써드파티 통합

{% hint style="info" %}
**기능 가용성**

써드파티 통합은 엔터프라이즈 계획을 갖춘 Snyk AppRisk Pro 버전에서만 사용 가능합니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans/)을 참조하세요.
{% endhint %}

### 전제 조건

Snyk AppRisk 써드파티 통합을 위해 Snyk 브로커를 설치하고 실행하려면 다음 단계를 따르세요.

1. Snyk AppRisk 통합을 위한 Snyk 브로커 토큰을 반드시 보유하세요. Snyk 지원팀에서 필요한 토큰을 제공할 수 있습니다.
   * [Snyk 브로커 토큰 얻기](snyk-broker-code-agent/install-snyk-broker-code-agent-using-docker/obtain-the-required-tokens-for-setup.md#obtain-your-broker-token-for-snyk-broker-code-agent) 페이지의 지침을 따라 브로커 토큰을 생성하세요.
   * 생성된 브로커 토큰을 통합 허브의 통합 설정 메뉴에 복사하여 붙여넣기하세요.
2. 다음 명령어를 실행하여 최신 브로커 이미지를 가져옵니다:

```docker
docker pull snyk/broker:universal
```

3. `snyk-broker-config` 명령어를 사용하여 Snyk AppRisk 연결 유형을 구성하세요. 자세한 내용은 [유니버설 브로커의 초기 설정](universal-broker/initial-configuration-of-the-universal-broker.md) 페이지에서 확인할 수 있습니다.

### Checkmarx SAST 통합

써드파티 통합에 적용되는 모든 일반적인 단계를 구현한 후, 고유한 자격 증명으로 통합을 구성할 수 있습니다.&#x20;

다음 예제에서는 자격 증명 참조로 `CHECKMARX_PASSWORD` 값을 사용합니다. 다음 명령어를 자신의 암호와 함께 실행하세요:

```docker
docker run --restart=always \
        -p 8001:8001 -e PORT=8001 \
        -e BROKER_CLIENT_URL=http://broker.url.example:8000 \
        -e BROKER_TOKEN=<YOUR BROKER TOKEN> \
        -e UNIVERSAL_BROKER_ENABLED=true \
        -e CHECKMARX_PASSWORD=<YOUR CHECKMARX PASSWORD> \
        -e BROKER_SERVER_URL=https://broker.snyk.io \
        -v $(pwd)/config.universal.json:/home/node/config.universal.json \
    snyk/broker:universal
```

### SonarQube SAST 통합

써드파티 통합에 적용되는 모든 일반적인 단계를 구현한 후, 고유한 자격 증명으로 통합을 구성할 수 있습니다.&#x20;

다음 예제에서는 자격 증명 참조로 `SONARQUBE_HOST_URL` 및 `SONARQUBE API_TOKEN` 값을 사용합니다. 다음 명령어를 실행하세요:

```docker
docker run --restart=always \
-p 8001:8001 -e PORT=8001 \
-e BROKER_CLIENT_URL=http://broker.url.example:8000 \
-e BROKER_TOKEN=<YOUR BROKER TOKEN> \
-e UNIVERSAL_BROKER_ENABLED=true \
-e SONARQUBE_HOST_URL=<YOUR HOST URL> \
-e SONARQUBE_API_TOKEN=<YOUR API TOKEN> \
-e BROKER_SERVER_URL=https://broker.snyk.io \
-v $(pwd)/config.universal.json:/home/node/config.universal.json \
snyk/broker:universal
```

### 구성 완료

써드파티 통합과의 Universal Broker 연결 설정이 완료되면 로그에 다음 메시지가 표시됩니다: `successfully established a websocket connection to the broker server`.

{% code overflow="wrap" %}
```docker
{"id":"broker-client-url-validation","name":"Broker Client URL Validation Check","status":"passing","output":"config check: ok"},{"id":"universal-broker-connections-config-validation","name":"Universal Broker Client Connections Configuration Check","status":"passing","output":"connections config check: ok"}],"version":"4.179.5","supportedIntegrationType":"apprisk"},"msg":"successfully established a websocket connection to the broker server","time":"2024-03-11T11:43:26.014Z","v":0}
```
{% endcode %}
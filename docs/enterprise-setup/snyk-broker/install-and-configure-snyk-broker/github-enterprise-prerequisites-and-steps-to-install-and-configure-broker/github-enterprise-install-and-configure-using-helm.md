# GitHub Enterprise - Helm을 사용하여 설치 및 구성

설치하기 전에 [사전 요구 사항](./) 및 [Helm을 사용한 설치에 대한 일반 지침](../install-and-configure-broker-using-helm.md)을 검토하십시오.

이 차트를 사용하려면 먼저 다음 명령어를 실행하여 Snyk Broker Helm 차트를 추가해야 합니다:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20;

그런 다음 다음 명령어를 실행하여 Broker를 설치하고 환경 변수를 사용자 정의하십시오. 환경 변수의 정의에 대한 자세한 내용은 [GitHub Enterprise - Snyk Broker를 위한 환경 변수](github-enterprise-environment-variables-for-snyk-broker.md)를 참조하십시오.

`ACCEPT_CODE`는 차트에서 [기본적으로 true로 설정](https://github.com/snyk/snyk-broker-helm/blob/465d4ef279755fa5c9507975a88348bab04c2264/charts/snyk-broker/templates/broker_deployment.yaml#L383)되어 있고 [ACCEPT\_IAC](https://github.com/snyk/snyk-broker-helm/blob/465d4ef279755fa5c9507975a88348bab04c2264/charts/snyk-broker/templates/broker_deployment.yaml#L386C23-L386C43)도 마찬가지입니다. 필요에 따라 `disableAutoAcceptRules=true`를 사용하여 사용을 중지할 수 있지만, 일반적으로 이러한 항목은 활성화됩니다.

Snyk AppRisk는 기본적으로 `false`로 설정되어 있습니다. 이를 `true`로 설정하여 활성화할 수 있습니다.

`github`, `githubApi` 및 `githubGraphQl`의 값에는`https://`를 포함하지 마십시오.

{% hint style="info" %}
**기본값이 아닌 지역을위한 멀티 테넌트 설정**\
기본값이 아닌 지역에서 Snyk Broker를 설정할 때는 특정 URL을 가져야 하는 추가 환경 변수가 필요합니다. URL 및 예제에 대해서는 [Regional hosting and data residency, Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)을 참조하십시오.
{% endhint %}

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-enterprise \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set github=<ENTER_GHE_ADDRESS> \
             --set githubApi=<ENTER_GHE_API_ADDRESS> \
             --set githubGraphQl=<ENTER_GRAPHQL_ADDRESS> \
             --set enableAppRisk=true \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

Helm 명령어에서 사용자가 선택한 환경 변수를 전달할 수 있습니다. 자세한 내용은 [Broker Helm 차트를위한 사용자 정의 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)을 참조하십시오. 필요에 따라 구성 변경을 위해 [Helm Chart 설치를위한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침을 따르십시오.

Broker가 실행 중인지 확인하려면 Snyk 웹 UI에서 브로커 통합 설정을 확인하여 연결되어 있는지 확인하는 확인 메시지를 볼 수 있습니다. 연결된 후 프로젝트 가져오기를 시작할 수 있습니다.
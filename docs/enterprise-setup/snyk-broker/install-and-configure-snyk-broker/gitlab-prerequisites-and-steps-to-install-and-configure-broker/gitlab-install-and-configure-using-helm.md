# GitLab - Helm을 사용하여 설치 및 구성

설치하기 전에 [필수 요구 사항](./) 및 [Helm을 사용한 설치 일반 지침](../install-and-configure-broker-using-helm.md)을 검토하십시오.

이 차트를 사용하려면 먼저 다음 명령을 실행하여 Snyk Broker Helm 차트를 추가해야 합니다:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20; 

그런 다음 다음 명령을 실행하여 Broker를 설치하고 환경 변수를 사용자 정의합니다. 환경 변수의 정의는 [GitLab - Snyk Broker를 위한 환경 변수](gitlab-environment-variables-for-snyk-broker.md)를 참조하십시오.

`gitlab` 값에는 `https://.`를 포함시키지 마십시오.

`brokerToken`은 `BROKER_TOKEN` 환경 변수를 피드하는 Helm 변수이며 설정하는 변수입니다. 이는 토큰을 전달하는 Helm의 방법입니다.

`ACCEPT_CODE`는 차트에서 [기본적으로 true로 설정](https://github.com/snyk/snyk-broker-helm/blob/465d4ef279755fa5c9507975a88348bab04c2264/charts/snyk-broker/templates/broker_deployment.yaml#L383)되어 있으며, [ACCEPT\_IAC도](https://github.com/snyk/snyk-broker-helm/blob/465d4ef279755fa5c9507975a88348bab04c2264/charts/snyk-broker/templates/broker_deployment.yaml#L386C23-L386C43) 마찬가지입니다. `disableAutoAcceptRules=true`를 사용하여 필요시 비활성화할 수 있지만, 그렇지 않으면 이러한 옵션이 활성화됩니다.

`REMOVE_X_FORWARDED_HEADERS=true` 환경 변수를 사용하여 Broker 클라이언트가 GitLab으로의 요청에서 `XFF` 헤더를 제거합니다. 이렇게 하면 Broker가 제대로 작동합니다.

Snyk AppRisk는 기본적으로 `false`로 설정되어 있습니다. 이를 `true`로 설정하여 활성화할 수 있습니다.

{% hint style="info" %}
**기본값 이외의 지역을 위한 멀티 테넌트 설정**\
기본값 이외의 지역에서 Snyk Broker를 설정할 때 특정 URL을 사용하는 추가 환경 변수가 필요합니다. URL 및 예제에 대해서는 [리전 호스팅 및 데이터 주거지, Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=gitlab \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set gitlab=<ENTER_GITLAB_URL> \
             --set scmToken=<ENTER_GITLAB_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set enableAppRisk=true \
             -n snyk-broker --create-namespace
```

Helm 명령어에서 사용자가 원하는 모든 환경 변수를 전달할 수 있습니다. 자세한 내용은 [Broker Helm 차트를위한 사용자 정의 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)을 참조하십시오. 필요에 따라 구성 변경을 수행하기 위해 [Helm 차트 설치를 위한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침을 따르십시오.

브로커가 실행 중인지 확인하려면 Snyk 웹 UI에서 브로커화된 통합의 설정을 확인하여 연결이 확인되는 확인 메시지를 확인할 수 있습니다. 연결되면 프로젝트를 가져오기 시작할 수 있습니다.
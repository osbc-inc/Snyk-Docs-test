# GitHub - Helm을 사용하여 설치 및 구성

설치하기 전에 [필수 조건](./)을 검토하고 [Helm을 사용한 설치 일반 지침](../install-and-configure-broker-using-helm.md)을 확인하십시오.

이 차트를 사용하려면 먼저 다음 명령을 사용하여 Snyk Broker Helm 차트를 추가해야 합니다.

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20;

그런 다음 다음 명령을 실행하여 Broker를 설치하고 환경 변수를 사용자 정의할 수 있습니다. 환경 변수의 정의는 [GitHub - Snyk Broker를 위한 환경 변수](github-environment-variables-for-snyk-broker.md)를 참조하십시오.

Snyk AppRisk는 기본적으로 `false`로 설정되어 있습니다. 이를 `true`로 설정하여 활성화하세요.

{% hint style="info" %}
**기본 설정이 아닌 지역에 대한 멀티 테넌트 설정**\
기본 설정이 아닌 지역에서 Snyk Broker를 설정할 때는 특정 URL이 필요한 추가 환경 변수가 필요합니다. URL 및 예제에 대해서는 [지역 호스팅 및 데이터 관할권, Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set enableAppRisk=true \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

Helm 명령어에서 원하는 환경 변수를 전달할 수 있습니다. 자세한 내용은 [Broker Helm 차트를위한 사용자 정의 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)을 참조하십시오. 필요한대로 구성 변경을 하기 위해서 [Helm Chart 설치에 대한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침을 따르십시오.

Broker가 실행 중인지 확인하려면 Snyk 웹 UI에서 브로커가 통합된 설정을 확인하여 연결된 메시지를 확인할 수 있습니다. 연결이 완료되면 프로젝트를 가져오기 시작할 수 있습니다.
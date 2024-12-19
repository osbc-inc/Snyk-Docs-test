# Azure Repos - Helm을 사용하여 설치하고 구성하는 방법

{% hint style="info" %}
**기능 가용성**

Snyk Azure Repos은 Azure DevOps/TFS 2020 이상에서만 사용할 수 있습니다.
{% endhint %}

설치 전에 [전제 조건](./) 및 [Helm을 사용한 설치에 대한 일반 지침](../install-and-configure-broker-using-helm.md)을 확인하십시오.

이 차트를 사용하려면 먼저 다음 명령으로 Snyk Broker Helm 차트를 추가해야 합니다.

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20;

그런 다음 다음 명령을 실행하여 Broker를 설치하고 환경 변수를 사용자 정의하세요. 환경 변수의 정의에 대한 자세한 내용은 [Azure Repos - Snyk Broker를 위한 환경 변수](azure-repos-environment-variables-for-snyk-broker.md)를 참조하십시오.

변수 `azureReposHost`의 값에는 `https://`가 포함되지 않습니다. 여러 개의 Azure 조직이 있는 경우 각각에 대해 Broker를 배포해야 합니다.

Snyk AppRisk는 기본적으로 `false`로 설정되어 있습니다. `true`로 설정하여 활성화할 수 있습니다.

{% hint style="info" %}
**기본값이 아닌 지역을 위한 멀티 테넌트 설정**\
기본값이 아닌 지역에서 Snyk Broker를 설정할 때는 특정 URL을 필요로 하는 추가 환경 변수가 필요합니다. URL 및 예제에 대한 자세한 내용은 [지역 호스팅 및 데이터 합법성, Broker URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=azure-repos \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set azureReposToken=<ENTER_REPO_TOKEN> \
             --set azureReposOrg=<ENTER_REPO_ORG> \
             --set azureReposHost=<ENTER_REPO_HOST> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set enableAppRisk=true \
             -n snyk-broker --create-namespace
```

Helm 명령어에서 선택한 환경 변수를 전달할 수 있습니다. 자세한 내용은 [Broker Helm 차트를 위한 사용자 정의 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)을 참조하십시오. 필요한 구성 변경을 수행하려면 [Helm Chart 설치를 위한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침에 따르십시오.

브로커가 실행 중인지 확인하려면 Snyk 웹 UI에서 브로커된 통합의 설정을 보고 연결되어 있는지 확인하는 확인 메시지를 보실 수 있습니다. 연결된 후 프로젝트 가져오기를 시작할 수 있습니다.
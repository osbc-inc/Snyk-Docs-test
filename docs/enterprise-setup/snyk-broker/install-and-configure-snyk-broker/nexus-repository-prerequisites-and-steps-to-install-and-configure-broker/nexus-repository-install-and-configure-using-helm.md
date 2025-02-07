# Nexus Repository - Helm을 사용하여 설치 및 구성하기

{% hint style="info" %}
**기능 가용성**

Nexus Repository Manager과의 통합은 엔터프라이즈 플랜에서만 이용 가능합니다. 자세한 정보는 [플랜 및 가격 책정](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

지원되는 환경, 버전 및 사용자 권한을 포함한 Nexus Repository Manager와의 브로커되지 않은 통합에 대한 정보는 [Nexus Repository Manager 설정](../../../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/)을 참조하십시오. Artifactory 또는 Nexus Container Registry와의 브로커된 통합에 대한 정보는 [Snyk 브로커 - Container Registry 에이전트](../../snyk-broker-container-registry-agent/)를 참조하십시오.

{% hint style="info" %}
**기본 이외의 지역에 대한 멀티 테넌트 설정**\
기본 이외의 지역에서 Snyk 브로커를 설정할 때 특정 URL이 필요한 추가 환경 변수가 필요합니다. URL 및 예시에 대해서는 [지역 호스팅 및 데이터 보유성, 브로커 URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)을 참조하십시오.
{% endhint %}

## Nexus 3 Helm 설치

설치하기 전에 [필수 요구 사항](./) 및 Helm을 사용한 설치에 대한 일반적인 지침을 검토하십시오([Helm](../install-and-configure-broker-using-helm.md)).

이 차트를 사용하려면 먼저 snyk 브로커 Helm 차트를 추가해야 합니다. 다음 명령어로 리포지토리를 추가하십시오:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

그런 다음 다음 명령어를 실행하여 브로커를 설치하고 환경 변수를 사용자 정의할 수 있습니다. 환경 변수의 정의는 [Nexus Repository - 환경 변수에 대한 Snyk Broker](nexus-repository-environment-variables-for-snyk-broker.md)를 참조하십시오.

참고: `baseNexusUrl` 및 `nexusUrl` 값에는 `https://`를 포함해야 합니다.

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=nexus \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set baseNexusUrl=<ENTER_BASE_NEXUS_URL> \
             --set nexusUrl=<ENTER_NEXUS_URL>
             --set brokerClientValidationUrl=<ENTER_BROKER_CLIENT_VALIDATION_URL> \
             -n snyk-broker --create-namespace
```

Helm 명령어에 원하는 환경 변수를 전달할 수 있습니다. 자세한 내용은 [브로커 Helm 차트에 대한 사용자 정의 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)를 참조하십시오. 필요한대로 구성 변경을 하려면 [Helm 차트 설치를 위한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침을 따르십시오.

브로커가 실행 중인지 확인하려면 연결된 것을 확인하는 Snyk 웹 UI의 브로커 통합 설정을 보고 연결되었음을 나타내는 확인 메시지를 확인하십시오. 연결되면 프로젝트를 가져오기 시작할 수 있습니다.

## Nexus 2 Helm 설치

설치하기 전에 [필수 요구 사항](./) 및 Helm을 사용한 설치에 대한 일반적인 지침을 검토하십시오([Helm](../install-and-configure-broker-using-helm.md)).

이 차트를 사용하려면 먼저 snyk 브로커 Helm 차트를 추가해야 합니다. 다음 명령어로 리포지토리를 추가하십시오:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

그런 다음 다음 명령어를 실행하여 환경 변수를 사용자 정의할 수 있습니다. 환경 변수의 정의는 [Nexus Repository - 환경 변수에 대한 Snyk Broker](nexus-repository-environment-variables-for-snyk-broker.md)를 참조하십시오.

참고: `baseNexusUrl` 및 `nexusUrl` 값에는 `https://`를 포함해야 합니다.

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=nexus2 \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set baseNexusUrl=<ENTER_BASE_NEXUS_URL> \
             --set nexusUrl=<ENTER_NEXUS_URL>
             --set brokerClientValidationUrl=<ENTER_BROKER_CLIENT_VALIDATION_URL> \
             -n snyk-broker --create-namespace
```

Helm 명령어에 원하는 환경 변수를 전달할 수 있습니다. 자세한 내용은 [브로커 Helm 차트에 대한 사용자 정의 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)를 참조하십시오. 필요한대로 구성 변경을 하려면 [Helm 차트 설치를 위한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침을 따르십시오.

브로커가 실행 중인지 확인하려면 연결된 것을 확인하는 Snyk 웹 UI의 브로커 통합 설정을 보고 연결되었음을 나타내는 확인 메시지를 확인하십시오. 연결되면 프로젝트를 가져오기 시작할 수 있습니다.

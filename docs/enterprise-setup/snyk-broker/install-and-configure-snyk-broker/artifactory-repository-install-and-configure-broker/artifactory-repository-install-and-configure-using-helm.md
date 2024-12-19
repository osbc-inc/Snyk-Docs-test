# Artifactory Repository - Helm을 사용하여 설치 및 구성하기

{% hint style="info" %}
**기능 가용성**

Artifactory Repository와의 통합은 Enterprise 플랜에만 사용 가능합니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

설치하기 전에 [전제 조건](./)과 [Helm을 사용한 설치의 일반적인 지침](../install-and-configure-broker-using-helm.md)를 검토하십시오.

Artifactory Repository와 비브로커 통합에 대한 정보는 [Artifactory Repository 설정](../../../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/)를 참조하십시오. Artifactory Container Registry와 브로커로 통합하는 방법에 대한 정보는 [Snyk Broker -Container Registry Agent](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent)를 참조하십시오.

이 차트를 사용하려면 먼저 Snyk 브로커 Helm 차트를 추가해야 합니다.

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

그런 다음, 다음 명령을 실행하여 브로커를 설치하고 환경 변수를 사용자화합니다. 환경 변수의 정의는 [Artifactory Repository - Snyk Broker를 위한 환경 변수](artifactory-repository-environment-variables-for-snyk-broker.md)를 참조하십시오.

`artifactoryUrl` 값에 `https://`를 포함하지 마십시오.

{% hint style="info" %}
**기본값과 다른 지역을 위한 멀티 테넌트 설정**\
기본값이 아닌 다른 지역에서 Snyk 브로커를 설정할 때는 특정 URL이 필요한 추가 환경 변수가 필요합니다. URL 및 예시에 대한 자세한 내용은 [지역 호스팅 및 데이터 소재지, 브로커 URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=artifactory \
             --set brokerToken=<브로커_토큰_입력> \
             --set artifactoryUrl=<아티펙토리_URL_입력> \
             -n snyk-broker --create-namespace
```

Helm 명령어에서 사용자가 선택한 환경 변수를 전달할 수 있습니다. 자세한 내용은 [브로커 Helm 차트를 위한 사용자 지정 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)을 참조하십시오. 필요에 따라 설정 변경을 위해 [Helm Chart 설치에 대한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침을 따르십시오.

브로커가 정상적으로 실행 중인지 확인하려면 연결된 것을 확인할 수 있는 Snyk 웹 UI에서 브로커 통합의 설정을 확인하십시오. 연결될 때 Projects를 가져오도록 시작할 수 있습니다.
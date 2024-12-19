# Jira - Helm을 사용하여 설치 및 구성

설치하기 전에 [필수 조건](./) 및 [Helm을 사용한 설치에 대한 일반 지침](../install-and-configure-broker-using-helm.md)을 확인하십시오.

{% hint style="info" %}
**기본 설정이 아닌 지역을 위한 멀티테넌트 설정**\
기본 설정 이외의 지역에서 사용하기 위해 Snyk Broker를 설정할 때 특정 URL을 요구하는 추가 환경 변수가 필요합니다. URL 및 예제에 대해서는 [Regional hosting and data residency, Broker URLs](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하십시오.
{% endhint %}

## Jira 통합을 설치하는 명령어

이 차트를 사용하려면 먼저 다음 저장소를 추가하여 Snyk Broker Helm Chart를 추가해야 합니다:

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`&#x20;

그런 다음, 다음 명령을 실행하여 Broker를 설치하고 환경 변수를 사용자 정의할 수 있습니다. 환경 변수의 정의는 [Jira - Snyk Broker를 위한 환경 변수](jira-environment-variables-for-snyk-broker.md)를 참조하십시오.

참고: `jiraHostname` 값에 `https://`를 포함하지 마십시오.

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=jira \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set jiraUsername=<ENTER_JIRA_USERNAME> \
             --set jiraPassword=<ENTER_JIRA_PASSWORD>  \
             --set jiraHostname=<ENTER_JIRA_HOSTNAME>  \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

Helm 명령어에서 필요한 환경 변수를 전달할 수 있습니다. 자세한 내용은 [Broker Helm Chart를 위한 사용자 지정 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)을 참조하십시오. 필요에 따라 구성 변경을 하기 위해 [Helm Chart 설치에 대한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침을 따르십시오.

Broker가 실행 중인지 확인하려면 Snyk 웹 UI의 Broker 통합 설정을 확인하여 연결 확인 메시지를 볼 수 있습니다. 연결되면 프로젝트 가져오기를 시작할 수 있습니다.

## SSO가 활성화된 JIRA를 위한 Jira PAT 인증

SSO가 활성화되면 JIRA는 일반적으로 사용자 이름과 비밀번호의 사용을 제한하고 개인 액세스 토큰(PAT)의 사용을 요구합니다.

SSO가 활성화되었을 때, 대신에 베어러 토큰을 사용하는 권한 헤더를 사용할 특정 Jira 버전을 사용해야 합니다. 이 버전을 사용하려면 다음 구성을 제공하십시오:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=jira-bearer-auth \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set jiraPat=<ENTER_JIRA_PAT> \
             --set jiraHostname=<ENTER_JIRA_HOSTNAME>  \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             -n snyk-broker --create-namespace
```

{% hint style="info" %}
Helm 차트 버전 2.2.0 이상을 사용해야 합니다.
{% endhint %}
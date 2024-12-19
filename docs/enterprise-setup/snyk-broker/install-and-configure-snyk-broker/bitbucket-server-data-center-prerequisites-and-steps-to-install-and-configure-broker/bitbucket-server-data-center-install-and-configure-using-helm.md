# Bitbucket Server/Data Center - Helm을 사용하여 설치 및 구성

설치하기 전에 [전제 조건](./)을 검토하고 [Helm](../install-and-configure-broker-using-helm.md)을 사용한 설치에 대한 일반적인 지침을 확인하십시오.

이 차트를 사용하려면 먼저 다음 리포를 추가해야 합니다.

`helm repo add snyk-broker https://snyk.github.io/snyk-broker-helm/`

그런 다음 다음 명령을 실행하여 Broker를 설치하고 환경 변수를 사용자 정의합니다. 환경 변수의 정의에 대해서는 [Bitbucket Server/Data Center - Snyk Broker Basic Auth를 위한 환경 변수](bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md) 및 [Bitbucket Server/Data Center - Snyk Broker Personal Access Token (PAT)을 위한 환경 변수](bitbucket-server-data-center-environment-variables-for-snyk-broker-personal-access-token-pat.md)을 참조하십시오.

`bitbucket` 및 `bitbucketApi` 값에 `https://`를 포함하지 마십시오.

기본으로 Snyk AppRisk가 `false`로 설정되어 있습니다. 이를 사용하려면 해당 플래그를 `true`로 설정하십시오.

{% hint style="info" %}
**기본값이 아닌 지역을 위한 멀티 테넌시 설정**\
기본값이 아닌 지역에서 Snyk Broker를 설정할 때 특정 URL을 사용하는 추가 환경 변수가 필요합니다. URL 및 예시에 대해서는 [지역 호스팅 및 데이터 체류, Broker URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)을 참조하십시오.
{% endhint %}

**Basic Auth를 사용하여 Bitbucket Server와 함께 사용하기 위해 Broker를 구성하려면 다음 명령을 사용하십시오:**

```bash
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=bitbucket-server \
             --set brokerToken=<브로커_토큰_입력> \
             --set bitbucketUsername=<사용자_이름_입력> \
             --set bitbucketPassword=<비밀번호_입력> \
             --set bitbucket=<Bitbucket_URL_입력> \
             --set bitbucketApi=<Bitbucket_API_URL_입력> \
             --set brokerClientUrl=<브로커_클라이언트_URL>:<브로커_클라이언트_포트_입력> \
             --set enableAppRisk=true \
             -n snyk-broker --create-namespace
```

**Bearer Auth (개인 액세스 토큰)을 사용하여 Bitbucket Server와 함께 사용하기 위해 Broker를 구성하려면 다음 명령을 사용하십시오:**

```bash
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=bitbucket-server-bearer-auth \
             --set brokerToken=<브로커_토큰_입력> \
             --set bitbucketPat=<PAT_입력> \
             --set bitbucket=<Bitbucket_URL_입력> \
             --set bitbucketApi=<Bitbucket_API_URL_입력> \
             --set brokerClientUrl=<브로커_클라이언트_URL>:<브로커_클라이언트_포트_입력> \
             --set enableAppRisk=true \
             -n snyk-broker --create-namespace
```

Helm 명령어에서 선택한 환경 변수를 전달할 수 있습니다. 자세한 내용은 [Broker Helm Chart를 위한 사용자 지정 추가 옵션](../advanced-configuration-for-helm-chart-installation/custom-additional-options-for-broker-helm-chart-installation.md)을 참조하십시오. [Helm Chart 설치를 위한 고급 구성](../advanced-configuration-for-helm-chart-installation/) 지침을 따라 필요한대로 구성 변경을 수행할 수 있습니다.

Broker가 실행 중인지 확인하려면 연결된 통합 설정을 보고 연결되었다는 확인 메시지가 표시되는지 확인하십시오. 연결되면 프로젝트 가져오기를 시작할 수 있습니다.
# Helm을 사용하여 Snyk Broker - Code Agent 설치

{% hint style="warning" %}
**Deprecated**

코드 에이전트는 더 이상 유지되지 않으며 더 이상 유지되지 않습니다.

Snyk Broker를 사용하여 Snyk Code 분석을 실행하는 선호하는 방법은 [Brokered Code](../git-clone-through-broker.md)를 통해 이루어집니다. 코드 에이전트는 이점이 없는 대안 방법입니다. 자세한 내용은 Snyk 통합 컨설턴트나 기술 성공 담당자에게 문의하거나 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.

Snyk Broker - Code Agent로는 자동 [PR Checks](../../../scan-with-snyk/pull-requests/pull-request-checks/) 기능이 지원되지 않습니다.
{% endhint %}

Snyk Broker Code Agent를 배포하려면 `enableCodeAgent` 플래그를 `true`로 설정해야 합니다. `accept.json` 파일에 적절한 항목이 있는지 확인하십시오. 사용 중인 SCM에 대한 예제 파일을 [Broker Client Templates 저장소](https://github.com/snyk/broker/tree/master/client-templates)에서 검색하십시오. Snyk Broker Code Agent 문서에서 지정된 추가 항목이 있는지 확인하십시오.

GitLab에 대한 예제 명령어는 다음과 같습니다:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=gitlab  \
             --set brokerToken=<ENTER_BROKER_TOKEN> \ 
             --set scmToken=<ENTER_SCM_TOKEN> \
             --set gitlab=<ENTER_GITLAB_URL>  \
            --set acceptJsonFile=accept.json \
            --set brokerClientUrl=http://<BROKER_CLIENT_URL> \ 
            --set enableCodeAgent=true \ 
            --set snykToken=<ENTER_SNYK_TOKEN> \
            -n snyk-broker --create-namespace
```

`brokerClientUrl`은 Broker 컨테이너의 주소입니다. Broker 컨테이너의 기본 포트는 `8000`입니다. 자세한 내용은 값 파일을 참조하십시오.

`accept.json`은 Helm 차트와 동일한 디렉토리에 있어야 합니다.
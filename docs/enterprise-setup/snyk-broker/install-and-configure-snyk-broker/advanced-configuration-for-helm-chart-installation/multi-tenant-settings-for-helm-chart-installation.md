# Helm 차트 설치를 위한 Multi-tenant 설정

## **Broker Multi-tenant 설정**

다양한 Multi-tenant 환경에서 Helm 차트를 사용하려면 `brokerServerUrl`을 다음 URL 중 하나로 설정하십시오. 환경에 따라 선택하시면 됩니다:

유럽: `https://broker.eu.snyk.io`\
호주: `https://broker.au.snyk.io`

```
--set brokerServerUrl=<BROKER_SERVER_URL>
```

## **Code Agent Multi-tenant 설정**

{% hint style="warning" %}
**더 이상 사용되지 않음**

Code Agent는 더 이상 유지되지 않으며 사용되지 않습니다.

Snyk Broker를 통해 Snyk 코드 분석을 실행하는 것이 권장되는 방법입니다. [Brokered Code](../../git-clone-through-broker.md)를 통해 실행됩니다. Code Agent는 이점이 없는 대안적인 방법입니다. 자세한 내용은 Snyk 통합 컨설턴트 또는 기술 지원 관리자에게 문의하시거나 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.

자동 [PR Checks](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks) 기능은 Snyk Broker - Code Agent에서 지원되지 않습니다.
{% endhint %}

Code Agent를 사용하는 경우, `upstreamUrlCodeAgent` 값은 다음 URL 중 하나여야 합니다. 환경에 따라 선택하시면 됩니다:

유럽: `https://deeproxy.eu.snyk.io`\
호주: `https://deeproxy.au.snyk.io`

```
--set upstreamUrlCodeAgent=<UPSTREAM_URL>
```
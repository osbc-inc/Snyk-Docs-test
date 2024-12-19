# Docker를 활용하여 Snyk Broker - Code Agent 설치

{% hint style="warning" %}
**더 이상 사용되지 않음**

Code Agent는 더 이상 유지되지 않는 기능이며 더 이상 유지되지 않습니다.

[Snyk](https://support.snyk.io)의 기술 지원을 통해 선호하는 방법으로 Snyk Broker를 통해 Snyk Code 분석을 실행하는 것이 권장됩니다. Code Agent는 이점이 없는 대안 방법입니다. 자세한 내용은 Snyk 통합 컨설턴트, 기술 성공 관리자에게 문의하거나 [Snyk 지원](https://support.snyk.io)에 문의하십시오.

Snyk Broker - Code Agent에서 자동 [PR Checks](../../../../scan-with-snyk/pull-requests/pull-request-checks/) 기능은 지원되지 않습니다.
{% endhint %}

{% hint style="info" %}
Broker Client - Code Agent를 초기 배포하려면 Snyk 계정 팀과 협력해야 합니다.
{% endhint %}

Snyk Broker - Code Agent 설정 워크플로우는 다음과 같습니다:

1. [설정에 필요한 토큰 획득](obtain-the-required-tokens-for-setup.md).
2. 동일한 조직 및 통합에 대한 실행 중인 Broker 클라이언트가 이미 있는 경우, [기존 Broker 클라이언트 제거](remove-an-existing-broker-client.md).
3. [Code Agent - Broker 클라이언트 통신을 위한 내부 네트워크 생성](create-network-for-broker-client-and-code-agent-communication.md).
4. [Code Agent 설정](set-up-the-code-agent.md).
5. [Broker 클라이언트 설정](set-up-the-broker-client/).\
    코드 스니펫을 웹 UI 결과에 표시 여부에 따라 Broker Client 구성을 설정할 수 있습니다. [코드 스니펫을 표시하지 않고 Broker 클라이언트 실행](set-up-the-broker-client/run-the-broker-client-without-the-code-snippet-display.md) 및 [코드 스니펫을 표시하며 Broker 클라이언트 실행](set-up-the-broker-client/run-the-broker-client-with-the-code-snippets-display.md)을 참조하십시오.
6. [Code Agent - Snyk Broker 설정 테스트](test-the-snyk-broker-code-agent-setup.md).
# 이벤트 전달

{% hint style="info" %}
**Snyk 앱으로 전환**

Snyk는 이벤트 전달 통합을 Snyk Apps 플랫폼을 사용하도록 전환하고 있습니다. 이 변경으로 인해 현재 및 향후 클라우드 이벤트 통합에서 새로운 기능과 향상된 보안이 가능해집니다.&#x20;

전환 중에 기존 통합은 계속 정상적으로 작동하며, 고객들은 해당 통합을 승인하여 Snyk Apps로 전환된 후에도 계속 작동하는 것을 보장할 수 있습니다. 기존 통합을 승인하려면 다음 단계를 따르세요:

1. 귀하의 조직의 **설정** 페이지로 이동
2. 승인하려는 통합의 설정 섹션으로 이동하십시오. 예를 들어, Amazon EventBridge, AWS CloudTrail Lake, AWS Security Hub.
3. **앱 승인** 버튼을 클릭하고 앱 승인 흐름을 완료하십시오

전환 기간이 끝나면 **승인되지 않은 통합은 더 이상 이벤트를 전달할 수 없으며 작동이 중지됩니다.**
{% endhint %}

Snyk의 이벤트 전달 통합을 사용하면 Snyk 플랫폼 이벤트를 다른 플랫폼의 특정 제품에 직접 전송하여 사용자 정의 경고 설정, 직접 보고서 작성, 자동화 트리거 등을 설정할 수 있습니다.&#x20;

## 이벤트 유형

Snyk는 두 가지 다른 유형의 이벤트를 전송하는 것을 지원합니다:

1. **Snyk 이슈 이벤트** - 이러한 이벤트는 Snyk 프로젝트에서 새로운 이슈가 발견되었을 때 또는 이슈가 업데이트되었을 때 전송됩니다. 각 이벤트에는 발견된 취약점이나 기타 문제에 대한 정보가 포함되어 있으며, 해결책이 있는지 여부도 포함됩니다.&#x20;
2. **Snyk 플랫폼 감사 이벤트** - 이러한 이벤트는 Snyk 사용자가 Snyk 플랫폼에서 작업을 수행할 때마다 전송됩니다. 자세한 내용은 [감사 로그](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)를 참조하십시오.&#x20;

{% hint style="info" %}
**Snyk 이슈** 이벤트 유형에는 Snyk Cloud 이슈가 포함되지 않습니다.
{% endhint %}

{% hint style="info" %}
**Snyk 플랫폼 감사** 이벤트 유형은 Snyk 엔터프라이즈 플랜에서 사용할 수 있습니다. 자세한 내용은 [가격제도](../../implement-snyk/enterprise-implementation-guide/trial-limitations.md)를 참조하십시오.
{% endhint %}

## 지원되는 통합

이벤트 전달 통합은 다음 제품에서 사용할 수 있습니다:

* [Amazon EventBridge](amazon-eventbridge.md) - **이슈 이벤트** 및 **감사 이벤트**
* [AWS CloudTrail Lake](aws-cloudtrail-lake.md) - **감사 이벤트**
* [AWS SecurityHub](aws-security-hub.md) - **이슈 이벤트**

각 통합은 [Snyk Apps](../../snyk-api/how-to-use-snyk-apps-apis/) 플랫폼 위에서 구축되었습니다.
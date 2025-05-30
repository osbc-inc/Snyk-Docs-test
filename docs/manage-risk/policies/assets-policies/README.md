# 자산 정책

## 개요

Snyk AppRisk 정책을 사용하면 비즈니스 컨텍스트를 추가하고 통지를 받는 프로세스를 쉽게 자동화할 수 있습니다.

{% hint style="info" %}
정책이 생성된 후, 최대 3시간 후에 실행되며 그 후로는 3시간마다 실행됩니다.

정책이 매일 실행되도록 설정된 경우, 정책은 24시간이 종료된 후 3시간 후에 실행됩니다. 언제든지 **Run** 버튼을 사용하여 수동으로 정책을 실행할 수 있습니다.
{% endhint %}

Snyk AppRisk 정책에 액세스하려면 그룹 수준에서 자리를 잡은 다음 **정책**을 선택한 다음 **자산**을 선택하십시오.

다음 비디오는 정책 뷰에서 생성할 수 있는 정책 유형에 대한 개요를 제공합니다.

{% embed url="https://youtu.be/79oz_hgMrCE" %}
비디오가 마음에 드셨나요? [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)에서 나머지 코스를 확인하세요!
{% endembed %}

{% hint style="info" %}
[자산 관리](../../../manage-assets/) 및 [자산 정책](./)은 서로 연결되어 있습니다. 새로운 정책을 설정하기 전에 인벤토리 메뉴에서 자산을 확인하고 필터링했는지 확인하십시오.
{% endhint %}

## 사용 사례

자산을 정리하고 분류하며 최신 자산 정보를 항상 업데이트할 수 있도록 정책을 생성할 수 있습니다. 정책을 사용하는 일반적인 사례:

* [새로운 자산 알림](use-cases-for-policies/notification-policy-use-case.md)
* [자산 분류](use-cases-for-policies/classification-policy-use-case.md)
* [자산 태그 설정](use-cases-for-policies/tagging-policy-use-case.md)
* [보안 커버리지 제어](use-cases-for-policies/coverage-control-policy-use-case.md)

### 새로운 자산 알림

특정 기준을 충족하는 새 자산이 발견되면 AppSec 팀 구성원에게 알림을 보내세요. 예를 들어, Snyk AppRisk가 테라폼을 기술로 활용하는 새 저장소 자산을 감지하면 인프라 팀에 Slack 메시지를 보낼 수 있습니다.

정책에 대한 알림 조치(이메일 또는 Slack)를 설정할 때, 관련 자산에 대한 링크를 포함할 수 있습니다. 각 알림에는 정책에 영향을받는 모든 자산이 나열됩니다. 이메일 알림에 표시된 자산 목록은 자동으로 생성됩니다.

### 자산 분류

자산 속성에 따라 A(가장 중요)에서 D(가장 덜 중요)까지의 비즈니스 중요도에 따라 저장소 자산을 분류하세요. 예를 들어, 이름에 "customer-portal"이 포함된 저장소는 고객 포털 앱이 민감한 데이터를 보유하기 때문에 A로 분류되어야 할 수 있습니다.

### 자산 태그 설정

유연한 태그로 저장소 자산을 분류하고 라벨을 지정하세요. 이를 사용하여 자산 인벤토리를 필터링할 수 있습니다.

### 보안 커버리지

선택한 보안 제품에 의해 자산이 스캔되었는지 모니터링하세요. 하나 이상의 보안 제품을 선택하고 마지막 스캔이 발생해야 하는 시간을 지정할 수 있습니다.

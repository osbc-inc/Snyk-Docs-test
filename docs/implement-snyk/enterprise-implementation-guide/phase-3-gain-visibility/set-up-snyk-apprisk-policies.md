# Snyk AppRisk 정책 설정

[Snyk AppRisk 정책](../../../manage-risk/policies/assets-policies/)을 사용하면 비즈니스 컨텍스트를 추가하고 알림을 받는 프로세스를 자동화할 수 있습니다. 정책을 설정하여 커버리지 컨트롤 갭을 자동으로 식별할 수 있습니다. &#x20;

## 정책 이해하기

정책을 생성하고 이해하기 위해 두 가지 구성 요소가 필요합니다:

* **필터**: 태그 또는 자산 이름과 같은 정의된 기준을 기반으로 자산 그룹을 생성하는 데 도움이 됩니다.
* **동작**: 필터링된 자산에 대해 수행할 수 있는 동작을 설정합니다. 예를 들어 자산 분류를 설정하거나 Slack 메시지 또는 이메일을 보내는 등의 작업을 설정할 수 있습니다.

{% hint style="info" %}
새로 생성된 정책은 최대 3시간이 소요되어 가시적으로 보이고 적용될 수 있습니다.
{% endhint %}

## 정책 생성하기

**정책** 보기로 이동하고 오른쪽 상단의 **새 정책** 버튼을 사용하여 정책을 생성할 수 있습니다. 정책에 이름을 지정하고 선택적으로 설명을 추가합니다.&#x20;

정책 빌더 편집기는 두 가지 주요 영역에 중점을 둡니다:

* [필터 정의](../../../manage-risk/policies/assets-policies/create-policies.md#define-filters) - 자산 속성에 대한 필터 조건을 설정합니다.
* [동작 설정](../../../manage-risk/policies/assets-policies/create-policies.md#set-actions) - 필터링된 자산에 취해질 동작을 정의합니다.

### 주요 필터 유형 <a href="#key-filter-types" id="key-filter-types"></a>

새로운 정책을 생성할 때 여러 가지 필터 중에서 선택할 수 있습니다. 여기에 일반적으로 사용되는 몇 가지 필터의 선택사항이 있습니다:

* **이름:** 많은 회사들은 저장소 및 다른 자산에 대한 네이밍 규칙을 가지고 있습니다. 자산 이름이 특정 문자열을 포함하거나 일치하는지 확인하기 위해 네이밍 규칙을 사용합니다.
* **자산 유형:** 저장소 또는 소프트웨어 패키지와 같은 특정 자산 유형에만 정책을 적용할 수 있습니다.
* **클래스:** 특정 Class에만 대상을 지정하여 동작에서 많은 잡음을 줄일 수 있습니다. 예를 들어 비즈니스에 중요한 자산인 Class A 자산에 대해서만 작업을 수행하도록 할 수 있습니다.
* **태그:** 시스템 제공 태그나 Snyk에서 정의한 사용자 정의 태그를 사용할 수 있습니다.

다음은 정책을 생성하는 단계입니다:

1. **정책** 보기로 이동하고 오른쪽 상단의 **새 정책** 버튼을 클릭합니다.
2. 정책에 이름을 제공하고 선택적으로 설명을 추가하고 **다음**을 클릭합니다.
3. 정책의 필터를 정의하고 **적용**을 클릭합니다.
4. 정책의 동작을 설정합니다. 동작을 추가하려면 + 아이콘을 클릭합니다. 필터링된 자산을 기반으로 작업을 트리거하거나 다른 필터 노드와 결합하기 위해 논리 노드를 생성할 수 있습니다.&#x20;

    <figure><img src="../../../.gitbook/assets/image (2) (12) (1).png" alt="정책 생성 - 여러 노드"><figcaption><p>정책 생성 - 여러 노드</p></figcaption></figure>
5. 다른 필터 노드를 추가하고 옆에 있는 + 아이콘을 선택하면 기존 논리 노드로 이동할 수 있습니다.
6. 필터링된 자산을 기반으로 작업을 트리거합니다.&#x20;

    <figure><img src="../../../.gitbook/assets/image (3) (7).png" alt=""><figcaption><p>정책 생성 - 동작 설정</p></figcaption></figure>

### 정책 신선도

모든 정책은 생성 후 최대 3시간 이내에 자동으로 실행되며 이후 3시간마다 실행됩니다. 원하는 경우 정책을 수동으로 실행하여 정책을 자산에 적용할 수 있습니다. 설정한 동작에 따라 자산에 자동으로 변경 내용이 적용됩니다.

### 정책 사용 사례

Snyk AppRisk 정책을 사용하여 다음과 같은 사용 사례에 익숙해질 수 있습니다:

* 커버리지 제어 정책 - 특정 보안 컨트롤이 필요한 위치를 정의하기 위해 커버리지 정책을 식별하고 설정합니다.
* 분류 정책 - 중요도에 따라 자산을 분류합니다.
* 태깅 정책 - 일치하는 자산에 태그를 지정합니다.
* 알림 정책 - 자산에서 발생하는 변경 사항에 대한 알림을 받습니다.
* 커버리지 및 커버리지 갭 정책 - 커버리지 필터를 사용하여 자산이 제품에서 테스트된 적이 있는지 확인하고, 커버리지 갭 필터를 사용하여 설정된 커버리지 제어 정책의 요구 사항을 충족하는지 확인합니다.
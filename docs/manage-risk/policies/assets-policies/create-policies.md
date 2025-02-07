# 정책 만들기

Snyk AppRisk에는 정책을 생성하고 수정하는 강력한 정책 편집기가 포함되어 있습니다.

정책을 만드는 두 단계가 있습니다:

1. [필터 정의](create-policies.md#define-filters) - 자산 속성에 대한 필터 조건 설정
2. [동작 설정](create-policies.md#set-actions) - 필터된 자산에 대해 취해질 동작 정의

## 새로운 정책

**처음부터 시작** 옵션을 사용하여 새로운 정책을 생성하거나 **템플릿 사용** 옵션을 선택할 수도 있습니다.

<figure><img src="../../../.gitbook/assets/Policy - new UI.png" alt="정책 보기, 처음부터 시작 및 템플릿 사용 옵션을 제공하는 새로운 정책 버튼"><figcaption><p>정책 보기, 새로운 정책 옵션</p></figcaption></figure>

### 처음부터 시작 - 정책 생성

새로운 정책을 만들려면 정책/자산 보기에서 **새로운 정책** 옵션을 클릭하고 **처음부터 시작** 옵션을 선택해야 합니다.

정책의 이름을 지정하고 선택사항으로 정책에 대한 설명을 제공해야 합니다. 이러한 단계를 완료한 후에는 정책의 [필터를 정의](create-policies.md#define-filters)하고 [동작을 설정](create-policies.md#set-actions)해야 합니다.

### 템플릿 사용 - 정책 생성

사용 가능한 템플릿 중 하나를 선택하여 새로운 정책을 만들 수 있습니다. 정책 템플릿 중 하나를 선택하려면 정책/자산 보기에서 **새로운 정책** 옵션을 클릭하고 **템플릿 사용** 옵션을 선택해야 합니다. 정책 템플릿 카드에서 **템플릿 사용** 버튼을 클릭하여 템플릿 중 하나를 선택할 수 있습니다.

각 정책 템플릿은 이름, 설명 및 필터 및 동작 간의 그래픽 연결을 표시합니다.

다음 비디오에서는 정책 보기에서 정책 템플릿을 사용하는 방법을 설명합니다:

{% embed url="https://youtu.be/-4qux-bsRK4" %}
비디오가 마음에 드셨다면 [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)에서 전체 강의를 확인하세요!
{% endembed %}

필터 및 동작을 사용자 정의하거나 기본 템플릿을 그대로 사용할 수 있습니다. 템플릿 변경을 모두 완료한 후 **저장** 버튼을 클릭하여 새로운 정책을 만들 수 있습니다.

<figure><img src="../../../.gitbook/assets/Policy template - new UI.png" alt="정책 보기에서 액세스 가능한 정책 템플릿, 새로운 정책 버튼, 템플릿 사용 옵션"><figcaption><p>정책 보기에서 액세스 가능한 정책 템플릿, 새로운 정책 버튼, 템플릿 사용 옵션</p></figcaption></figure>

## **필터 정의**

{% hint style="info" %}
**릴리스 상태**

자산의 위험 요인은 적용된 [위험 요인](../../prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/risk-factor-loaded-package.md)의 릴리스 상태를 고려합니다.

Runtime 발견 및 Runtime 마지막으로 본 날짜 필터는 사용된 런타임 [통합](../../snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md)의 릴리스 상태를 고려합니다.
{% endhint %}

각 필터 구성 요소는 자산 속성을 지정해야 합니다. 자산 정책에 대해 사용 가능한 모든 속성을 보려면 [필터 기능](../../../manage-assets/assets-inventory-features.md#filters-capabilities) 페이지로 이동하세요.

다음 비디오에서는 새로운 정책을 만드는 방법을 설명합니다:

{% embed url="https://youtu.be/OMuyzAM1Omo" %}
비디오가 마음에 드셨다면 [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)에서 전체 강의를 확인하세요!
{% endembed %}

각 속성에는 다양한 조건 및 값이 포함되어 있습니다:

| 속성             | 조건 값                                                                                  | 값                                                                                 |
| -------------- | ------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| Application\*  | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | Snyk AppRisk에서 애플리케이션 컨텍스트를 구성한 모든 사용 가능한 애플리케이션                                  |
| 자산 ID          | <ul><li>일치</li><li>일치하지 않음</li><li>포함</li><li>포함하지 않음</li><li>시작</li><li>끝</li></ul>  | \[문자열]                                                                            |
| 자산 이름          | <ul><li>일치</li><li>일치하지 않음</li><li>포함</li><li>포함하지 않음</li><li>시작</li><li>끝</li></ul>  | \[문자열]                                                                            |
| 자산 유형          | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | <ul><li>패키지</li><li>저장소</li><li>스캔된 아티팩트</li></ul>                                |
| 속성             | <ul><li>일치</li><li>일치하지 않음</li><li>포함</li><li>포함하지 않음</li><li>시작</li><li>끝</li></ul>  | \[문자열]                                                                            |
| 카탈로그 이름\*      | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | 애플리케이션 컨텍스트의 이름 목록                                                                |
| 카테고리           | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | 저장소 자산의 사용 가능한 카테고리 목록                                                            |
| 클래스            | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | A, B, C, D                                                                        |
| 커버리지           | <ul><li>한 개 이상의 포함</li><li>모두 포함</li><li>한 개 이상의 포함하지 않음</li><li>모두 포함하지 않음</li></ul> | Snyk 코드, 컨테이너, IaC, 오픈 소스                                                         |
| 커버리지 갭         | <ul><li>한 개 이상의 포함</li><li>모두 포함</li><li>한 개 이상의 포함하지 않음</li><li>모두 포함하지 않음</li></ul> | Snyk 코드, 컨테이너, IaC, 오픈 소스                                                         |
| 개발자            | <ul><li>일치</li><li>일치하지 않음</li><li>포함</li><li>포함하지 않음</li><li>시작</li><li>끝</li></ul>  | \[문자열]                                                                            |
| 발견된 날짜         | <ul><li>내부에 있음</li><li>내부에 없음</li></ul>                                               | <ul><li>지난 24시간</li><li>지난 7일</li><li>지난 30일</li><li>지난 12개월</li><li>연간</li></ul> |
| 이슈 심각성         | <ul><li>한 개 이상의 포함</li><li>모두 포함</li><li>한 개 이상의 포함하지 않음</li><li>모두 포함하지 않음</li></ul> | <ul><li>중대한</li><li>높음</li><li>보통</li><li>낮음</li></ul>                            |
| 이슈 소스          | <ul><li>한 개 이상의 포함</li><li>모두 포함</li><li>한 개 이상의 포함하지 않음</li><li>모두 포함하지 않음</li></ul> | Snyk 코드, 컨테이너, IaC, 오픈 소스, Nightfall                                              |
| 최근 확인          | <ul><li>내부에 있음</li><li>내부에 없음</li></ul>                                               | <ul><li>지난 24시간</li><li>지난 7일</li><li>지난 30일</li><li>지난 12개월</li><li>연간</li></ul> |
| 라이프사이클\*       | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | 애플리케이션 컨텍스트 구성 요소의 라이프사이클 상태 목록                                                   |
| 잠긴 속성          | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | <ul><li>클래스</li></ul>                                                             |
| 소유자\*          | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | 애플리케이션 컨텍스트가 구성된 저장소를 소유한 팀의 목록                                                   |
| 위험 요인          | <ul><li>한 개 이상의 포함</li><li>모두 포함</li><li>한 개 이상의 포함하지 않음</li><li>모두 포함하지 않음</li></ul> | 사용 가능한 위험 요인 목록                                                                   |
| Runtime 발견     | <ul><li>내부에 있음</li><li>내부에 없음</li></ul>                                               | <ul><li>지난 24시간</li><li>지난 7일</li><li>지난 30일</li><li>지난 12개월</li><li>연간</li></ul> |
| Runtime 마지막 확인 | <ul><li>내부에 있음</li><li>내부에 없음</li></ul>                                               | <ul><li>지난 24시간</li><li>지난 7일</li><li>지난 30일</li><li>지난 12개월</li><li>연간</li></ul> |
| SCM 저장소 신선도    | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | <ul><li>활성</li><li>비활성</li><li>휴면</li></ul>                                       |
| 소스             | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | <ul><li>azure-devops</li><li>GitHub</li><li>GitLab</li><li>Snyk</li></ul>         |
| 태그             | <ul><li>한 개 이상의 포함</li><li>모두 포함</li><li>한 개 이상의 포함하지 않음</li><li>모두 포함하지 않음</li></ul> | 이전에 생성한 모든 사용 가능한 태그                                                              |
| 제목\*           | <ul><li>중 하나이다</li><li>중 하나가 아니다</li></ul>                                            | 애플리케이션 컨텍스트가 구성된 구성 요소의 모든 이름 목록                                                  |

**\*** 마크가 지정된 모든 필터는 SCM 통합에 애플리케이션 컨텍스트를 구성한 사용자만 볼 수 있습니다.

**And** 또는 **Or** 연산자를 사용하여 둘 이상의 필터 구성 요소를 지정할 수 있습니다.

<figure><img src="../../../.gitbook/assets/Create policy New UI.png" alt="AppRisk - 새로운 정책 만들기"><figcaption><p>Snyk AppRisk - 새로운 정책 만들기</p></figcaption></figure>

다음 비디오에서는 필터 사용과 **And**, **Or** 연산자 사용 방법을 설명합니다.

{% embed url="https://youtu.be/W-GDCZVxLIo" %}
비디오가 마음에 드셨다면 [Snyk Learn](https://learn.snyk.io/lesson/snyk-apprisk-essentials/)에서 전체 강의를 확인하세요!
{% endembed %}

## **동작 설정**

필터 구성 요소를 정의한 후, 해당 정책이 필터링된 자산에 대해 수행해야 하는 동작을 정의해야 합니다. 자산 정책은 다음과 같은 동작을 지원합니다:

* **이메일 보내기** - 자산 업데이트가 있는 경우마다 이메일을 수신합니다. 매일 이메일 또는 체크 일정을 예약하는 옵션을 선택할 수 있습니다. 관련 자산에 대한 링크를 포함할 수 있습니다. 각 알림에는 영향을 받는 모든 자산이 나열됩니다. 이메일 알림에 표시되는 자산 목록은 자동으로 생성됩니다.
* **슬랙 메시지 보내기** - 자산 업데이트가 있는 경우마다 Slack 알림을 수신합니다. [슬랙 웹훅 URL](../../../integrate-with-snyk/jira-and-slack-integrations/slack-integration.md)을 추가해야 하며, 매일 이메일을 선택하거나 체크를 예약할 수 있습니다. 관련 자산에 대한 링크를 포함할 수 있습니다. 각 알림에는 영향을 받는 모든 자산이 나열됩니다. 이메일 알림에 표시되는 자산 목록은 자동으로 생성됩니다.
* **자산 클래스 설정** - 일치하는 자산의 클래스를 설정합니다. 정책을 제거하거나 해제해도 에셋 클래스가 기본값으로 소급하여 변경되지는 않습니다.
* **자산 태그 설정** - 일치하는 자산에 태그를 설정합니다. 정책을 제거하거나 끄면 관련 자산에서 이 정책의 태그가 제거됩니다.
* **적용 범위 제어 정책 설정** - 필터링된 자산에 대해 선택적으로 주어진 시간 내에 선택한 보안 제품이 자산을 스캔하고 있는지 여부를 확인하는 제어를 설정합니다. 이 제어에 실패한 자산은 인벤토리 페이지에 그에 따라 표시됩니다. 이 제어는 제품 전체에 OR 논리를 적용합니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FE06zZpqSfYAZ8FV2qAAV%252FPolicy%2520-%2520Nwe%2520UI.png%3Falt%3Dmedia%26token%3D73eb75d9-7680-4d1d-9fad-c14a62ed0995&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=306b3739&#x26;sv=2" alt=""><figcaption><p>Snyk Essentials 또는 Snyk AppRisk - 정책 조치 설정</p></figcaption></figure>

편집기는 동일한 정책에 대해 여러 개의 플로우를 지원합니다. 흐름은 독립적이거나 교차할 수 있습니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FH2P7l7BkjgYs8PaMGfIs%252FMultiple%2520actions%2520-%2520New%2520UI.png%3Falt%3Dmedia%26token%3D3a43dc09-b2e2-4e87-9056-8dab5327675a&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=23f5e320&#x26;sv=2" alt=""><figcaption><p>Snyk Essentials 또는 Snyk AppRisk - 여러 정책 작업 설정</p></figcaption></figure>


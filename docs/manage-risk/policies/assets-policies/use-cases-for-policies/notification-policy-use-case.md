# 알림 정책 - 사용 사례

자산에 발생하는 변경 사항을 알릴 때 **이메일 보내기** 및 **Slack 메시지 보내기** 작업을 사용할 수 있습니다.

이 사용 사례에서는 새로운 클래스 A 자산이 Snyk 보안 커버리지가 없을 때마다 알림을 설정하고 수신하는 방법을 보여줍니다.

이 예제를 따라가려면 다음을 찾아내기 위해 네 가지 필터를 만들어야 합니다.

* **필터 1**: Repository 유형의 자산입니다.

<figure><img src="../../../../.gitbook/assets/image (206).png" alt="자산 유형에 대한 필터 구성" width="352"><figcaption><p>자산 유형에 대한 필터 구성</p></figcaption></figure>

* **필터 2**: 클래스 A 자산입니다.

<figure><img src="../../../../.gitbook/assets/image (223).png" alt="자산 클래스에 대한 필터 구성" width="354"><figcaption><p>자산 클래스에 대한 필터 구성</p></figcaption></figure>

* **필터 3**: 관련 프로그래밍 언어를 포함하는 태그(예: Apex, ASP, C, C#, C++, CMake, Go, HTML, Java, JavaScript, Kotlin, PHP, Python, Ruby, Scala, Swift, TypeScript, VisualBasic, Handlebars, Makefile, Objective-C)입니다.

<figure><img src="../../../../.gitbook/assets/image (236).png" alt="태그에 대한 필터 구성" width="349"><figcaption><p>태그에 대한 필터 구성</p></figcaption></figure>

* **필터 4**: Snyk Open Code 또는 Snyk Open Source 스캔 커버리지가 없습니다.

<figure><img src="../../../../.gitbook/assets/Coverage filter.png" alt="커버리지 제어를 위한 필터 구성" width="353"><figcaption><p>커버리지 제어를 위한 필터 구성</p></figcaption></figure>

필터 조건을 설정한 후 **이메일 보내기** 작업을 선택해야 합니다.

{% hint style="info" %}
**Slack 메시지 보내기** 작업을 설정하려면 [Incoming WebHooks 앱](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks)을 사용하거나 [자체 Slack 앱](https://api.slack.com/incoming-webhooks)을 만들어 Slack 웹훅을 생성할 수 있습니다.
{% endhint %}

**이메일 보내기** 작업을 사용자 정의하여 정책에 영향 받는 자산을 가리키는 링크와 함께 알림을 받도록 설정할 수 있습니다. 이 작업의 **본문** 필드 내에 "/"를 입력하고 **자산 링크(Link to Assets)**를 선택함으로써 이를 수행할 수 있습니다. 정책을 저장한 후, 수신된 모든 알림에서 정책에 영향을 받는 모든 자산이 나열됩니다.

<figure><img src="../../../../.gitbook/assets/image (509).png" alt="Snyk AppRisk - Send Email 작업에서 Link to Assets 옵션 설정 "><figcaption><p>Snyk AppRisk - Send Email 작업에서 Link to Assets 옵션 설정 </p></figcaption></figure>

모든 필터 및 작업이 설정된 후 정책이 어떻게 보이는지는 아래에 나와 있습니다.

<figure><img src="../../../../.gitbook/assets/image (508).png" alt="Snyk AppRisk - 알림 정책 설정"><figcaption><p>Snyk AppRisk - 알림 정책 설정</p></figcaption></figure>

**자산에 대한 링크(Link to Assets)** 옵션을 본문 필드에 포함시킨 후 이메일 알림을 받게 됩니다. 이메일 알림에서 자산에 액세스하거나 **여기를 클릭(Click Here)**하여 집계 형태로 보거나 수 있습니다. 이메일 알림에 표시된 자산 목록은 자동으로 생성됩니다.

<figure><img src="../../../../.gitbook/assets/image (510).png" alt="Snyk AppRisk - Send Email 작업의 알림 예시"><figcaption><p>Snyk AppRisk - Send Email 작업의 알림 예시</p></figcaption></figure>

{% hint style="info" %}
이메일 알림 정책이 생성된 후 최대 3시간 후에 실행되며 그 후 3시간마다 실행됩니다.

정책이 일일로 설정된 경우, 24시간 기간이 끝난 후 최대 3시간 내에 실행됩니다. 수동으로 정책을 실행하려면 실행 버튼을 사용할 수 있습니다.
{% endhint %}
# Slack 통합

{% hint style="info" %}
Snyk은 Slack 통합이 오래되었기 때문에 모든 고객이 [Slack 앱](slack-app.md)을 사용할 것을 권장합니다.
{% endhint %}

Slack을 설정하여 Snyk의 프로젝트에 영향을 주는 새로운 취약점 및 사용 가능한 새로운 업그레이드 또는 패치에 대한 알림을 받을 수 있습니다.

{% hint style="info" %}
프로젝트 초기 가져 오기에서 감지된 취약점은 즉시 Slack으로 보내지지 않습니다.
{% endhint %}

Slack에서 다음과 같은 알림을 받게 됩니다.

새로 공개된 취약점이 영향을 미칩니다:

<figure><img src="../../.gitbook/assets/image (23) (1) (1).png" alt="새로 공개된 취약점 알림"><figcaption><p>새로 공개된 취약점 알림</p></figcaption></figure>

이전에 무시했거나 패치하지 않은 취약점에 대해 새 업그레이드 또는 패치가 제공됩니다:

<figure><img src="../../.gitbook/assets/image (22) (1).png" alt="새 업그레이드 사용 가능"><figcaption><p>새 업그레이드 사용 가능</p></figcaption></figure>

통합을 설정하려면 Slack 웹훅을 생성해야 합니다. 이를 위해 [들어오는 웹훅 앱](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks)을 통해 생성하거나 [자체 Slack 앱을 생성](https://api.slack.com/incoming-webhooks)할 수 있습니다.

Slack 웹훅 URL을 생성했으면 **조직 관리** 설정으로 이동하여 URL을 입력하십시오.

<figure><img src="../../.gitbook/assets/image (24) (1) (1) (1) (1) (1) (1).png" alt="Slack 웹훅의 URL 입력"><figcaption><p>Slack 웹훅의 URL 입력</p></figcaption></figure>

{% hint style="info" %}
Slack 앱으로 생성된 웹훅만 지원되며 Slack Workflows로 생성된 웹훅은 지원되지 않습니다.
{% endhint %}
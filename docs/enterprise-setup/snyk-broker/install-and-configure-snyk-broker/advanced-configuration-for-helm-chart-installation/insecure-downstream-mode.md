---
description: 이 모드의 사용은 권장되지 않습니다.
---

# 보안되지 않은 다운스트림 모드

일부 상황에서는 다운스트림 연결에 대해 HTTP만 사용해야 할 수 있습니다. 이러한 경우는 비교적 드물며 주로 과거의 http-only 설정 때문에 발생합니다. Snyk은 HTTPS 대신 사용하기를 권장합니다.

가까운 미래에 이를 수정할 수 없는 경우, 보안되지 않은 다운스트림 모드는 SCM/JIRA/기타에 대한 다운스트림 요청을 HTTPS 대신 HTTP로 진행하도록 강제하는 방법을 도입합니다.

대부분의 경우에는 이 모드를 사용하지 않아야 합니다. 이는 선택 사항입니다. 모든 요청이 HTTP를 통해 이루어지므로 TLS 암호화의 안전성 혜택을받을 수 없습니다. 이 모드는 모든 자격 증명과 데이터가 암호화되지 않은 상태로 표시되어, 심층적으로 보호된 네트워크에서만 허용됩니다.

이 환경 변수를 주입하는 데 [브로커 Helm 차트 설치에 대한 사용자 지정 추가 옵션](custom-additional-options-for-broker-helm-chart-installation.md)을 사용하세요:

`--set env[0].name=INSECURE_DOWNSTREAM --set env[0].value="true"`

{% hint style="warning" %}
HTTP의 사용은 매우 불안전합니다. 데이터와 자격 증명이 네트워크 교환을 통해 평문으로 전송됩니다.

Snyk은 HTTP 사용으로 발생할 수 있는 자격 증명 유출에 대해 **책임지지 않습니다**.
{% endhint %}

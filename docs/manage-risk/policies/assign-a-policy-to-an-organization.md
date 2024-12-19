# 조직에 정책 할당하기

정책을 만들 때, 한 조직에 적용할 수 있습니다. 정책 매니저를 사용하여 기본 정책에 조직을 직접 적용하거나 제거할 수는 없습니다.

{% hint style="info" %}
조직에 적용된 정책은 `snyk test` 또는 `snyk monitor` CLI 명령을 실행할 때 적용됩니다.
{% endhint %}

## 조직에 정책 적용하기

조직에 정책을 적용하려면, 조직 선택기 패널에서 정책을 적용하려는 조직의 상자를 체크합니다.

만약 다른 정책이 이미 적용된 조직이 있다면, 선택기에서 해당 정책을 볼 수 있으며, 조직 이름 옆의 정책 표시기는 회색이 됩니다.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip3-2-.png" alt="회색 표시기 - 이미 다른 정책이 적용된 조직"><figcaption><p>회색 표시기 - 이미 다른 정책이 적용된 조직</p></figcaption></figure>

</div>

이미 정책이 적용된 조직에 새로운 정책을 적용하는 경우, 해당 정책의 이름이 조직 이름 옆에 노란색 표시기로 표시됩니다.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip2-6-.png" alt="노란색 표시기 - 이미 이 정책에 할당된 조직"><figcaption><p>노란색 표시기 - 이미 이 정책에 할당된 조직</p></figcaption></figure>

</div>

다른 정책을 조직에 적용하는 경우, 해당 조직 이름 옆에 두 개의 표시기가 나타납니다. 노란색으로 현재 적용된 정책을 보여주는 하나와 회색으로 새로운 정책을 적용할 것을 보여주는 다른 하나입니다.

<div align="left">

<figure><img src="../../.gitbook/assets/mceclip1-16-.png" alt="회색 및 노란색 표시기 - 적용된 정책 및 적용될 정책"><figcaption><p>회색 및 노란색 표시기 - 적용된 정책 및 적용될 정책</p></figcaption></figure>

</div>

## 조직에서 정책 제거하기

조직에서 정책을 제거하려면, 보고 있는 정책 옆의 상자를 체크 해제합니다.

<div align="left">

<figure><img src="../../.gitbook/assets/untitled-2-.png" alt="조직에서 정책 제거하기"><figcaption><p>조직에서 정책 제거하기</p></figcaption></figure>

</div>

체크를 해제한 조직은 자동적으로 기본 정책으로 되돌아갑니다.

## 조직에 기본 정책 적용하기

현재 적용된 정책을 제거합니다. 해당 조직은 자동적으로 기본 정책으로 되돌아갑니다.

## 조직에서 기본 정책 제거하기

새로운 정책을 적용합니다. 해당 조직은 자동적으로 기본 정책에서 제거되며 새로운 정책이 적용됩니다.
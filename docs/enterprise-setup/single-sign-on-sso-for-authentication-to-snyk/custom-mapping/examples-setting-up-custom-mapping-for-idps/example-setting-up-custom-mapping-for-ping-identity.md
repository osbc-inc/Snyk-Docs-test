# 예시: Ping Identity를 위한 사용자 정의 매핑 설정

이 페이지에서는 [Legacy custom mapping](../legacy-custom-mapping.md)을 사용하여 Ping Identity의 역할을 사용자 정의 매핑하는 방법을 설명합니다.

{% hint style="info" %}
이 안내서는 Ping Identity 응용 프로그램이 구성되고 작동 중이라고 가정합니다.
{% endhint %}

{% hint style="info" %}
엔터프라이즈 응용 프로그램을 설정하는 Snyk 측의 단계는 사용자 정의 매핑을 수용하지 않기 때문에, Snyk 연락처에 의해 수행되어야 합니다.
{% endhint %}

1. 응용 프로그램 구성에서 **속성 매핑**을 선택하고 속성을 편집하려면 연필 아이콘을 클릭합니다.

    <figure><img src="../../../../.gitbook/assets/6 (3).png" alt="속성 매핑 편집"><figcaption><p>속성 매핑 편집</p></figcaption></figure>
2. **+추가**를 선택하고 다음 속성을 입력한 후 변경 내용을 저장합니다.\
    **roles**: `그룹 이름`\

    <figure><img src="../../../../.gitbook/assets/Screenshot 2023-09-05 at 12.02.30 PM.png" alt="역할 배열 추가"><figcaption><p>역할 배열 추가</p></figcaption></figure>
3. 왼쪽 메뉴에서 **Identities/Groups**을 선택하고 [사용자 정의 매핑 옵션](../) 페이지에서 설명된 구문을 따르는 Snyk 그룹을 추가합니다.&#x20;

    <figure><img src="../../../../.gitbook/assets/12 (2).png" alt="예제 그룹 추가"><figcaption><p>예제 그룹 추가</p></figcaption></figure>
4. 이전 화면 아래 **인구**를 선택하지 않은 경우, Snyk 역할에 속해야 하는 사용자에게 해당 그룹을 할당해야 합니다. **인구**를 선택한 경우, 해당 인구의 모든 사용자가 할당된 Snyk 역할의 권한을 상속받게 됩니다.
5. 프로세스를 완료하려면 SAML 페이로드에 역할 배열이 포함되어 있는지 확인하고 사용자 정의 매핑 기능을 활성화하기 위해 Snyk 연락처에 문의하십시오.
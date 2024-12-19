# 예시: Entra ID에 대한 사용자 지정 매핑 설정

다음 정보는 Entra ID(이전 Azure AD)의 역할을 사용자 정의 매핑하는 방법을 구성하는 방법을 보여줍니다.

{% hint style="info" %}
[Entra ID 기업 응용프로그램 예제](../../configure-self-serve-single-sign-on-sso/azure-ad-enterprise-application-setup.md)를 참조하여 초기 기업 응용프로그램을 설정하는 지침을 확인하십시오.

엔터프라이즈 응용프로그램을 설정하는 Snyk 측의 어떤 단계든 사전 서비스로 셀프 서비스 SSO는 사용자 정의 매핑을 지원하지 않습니다.
{% endhint %}

다음은 **App 역할 구성**을 위한 **전제 조건**입니다:

- Snyk 지원은 Snyk SSO를 Microsoft Entra ID (WAAD 또는 SAML)로 구성해야 합니다.
- SAML을 선택한 경우 사용자 지정 클레임을 추가해야 합니다. 해당 단계는 이 지침서에 있습니다.
- 기존의 Azure 엔터프라이즈 응용프로그램과 해당 SSO 구성에 연결된 앱 등록이 있어야 합니다.

**앱 역할 구성** 단계는 다음과 같습니다.

1. 앱 등록 메뉴에서 기업 응용프로그램의 이름을 선택합니다.

    <figure><img src="../../../../.gitbook/assets/image (113) (1).png" alt="앱 등록, 기업 응용프로그램 이름 선택"><figcaption><p>앱 등록, 기업 응용프로그램 이름 선택</p></figcaption></figure>
2. **App 역할**을 선택한 다음 **앱 역할 만들기**를 선택합니다.

    <figure><img src="../../../../.gitbook/assets/image (1) (1) (2) (1) (1) (1).png" alt="App 역할 선택, 앱 역할 만들기"><figcaption><p>App 역할 선택, 앱 역할 만들기</p></figcaption></figure>
3. 필요한 세부 정보로 앱 역할을 작성합니다.\
    **허용된 멤버 유형**을 선택합니다: **사용자/그룹**, **응용프로그램** 또는 **둘 다**.\
    선택한 유형에 대한 **값** 및 **설명**을 입력합니다.\
    앱 역할을 활성화합니다.\
    완료하면 **적용**을 선택합니다.\\

    <figure><img src="../../../../.gitbook/assets/image (380).png" alt="앱 역할 만들기" width="285"><figcaption><p>앱 역할 만들기</p></figcaption></figure>
4. Entra ID에서 기업 응용프로그램을 선택합니다.

    <figure><img src="../../../../.gitbook/assets/image (3) (2) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="Entra ID에서 기업 응용프로그램 선택"><figcaption><p>Entra ID에서 기업 응용프로그램 선택</p></figcaption></figure>
5. **사용자 및 그룹**을 선택한 다음 **사용자/그룹 추가**를 선택합니다.\
    추가할 사용자 및 그룹을 검색하고 선택합니다.

    <figure><img src="../../../../.gitbook/assets/image (4) (5).png" alt="사용자 및 그룹 선택, 사용자/그룹 추가"><figcaption><p>사용자 및 그룹 선택, 사용자/그룹 추가</p></figcaption></figure>
6. **사용자 및 그룹**을 선택한 다음 드롭다운에서 역할을 선택하고 **지정**을 선택합니다.\\

    <figure><img src="../../../../.gitbook/assets/image (383).png" alt="할당 추가"><figcaption><p>할당 추가</p></figcaption></figure>
7. 할당해야 할 모든 필수 그룹과 역할에 대해 반복하고 목록이 다음과 유사하게 보이는지 확인합니다.\\

    여러 Snyk 역할을 한 App 역할에 추가할 수도 있습니다. 데이터의 페이로드는 쉼표로 구분된 문자열로 해석될 수 있기 때문입니다. 그러나 여러 App 역할과 동시에는 사용할 수 없으며, 문자열 또는 배열 중 한 가지만 적용됩니다.

    <figure><img src="../../../../.gitbook/assets/image (384).png" alt="사용자 및 그룹 목록"><figcaption><p>사용자 및 그룹 목록</p></figcaption></figure>
8. SAML 연결을 구성한 경우 SAML 페이로드에서 Snyk로 역할 배열을 전달할 수 있는 사용자 지정 클레임을 추가합니다. 좌측 메뉴에서 **단일 로그인**을 선택합니다.
9. **속성 및 클레임 편집** 옆의 **편집**을 선택합니다.

    <figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-10 at 3.19.31 PM.png" alt="속성 및 클레임 편집"><figcaption><p>속성 및 클레임 편집</p></figcaption></figure>
10. **새 클레임 추가**를 선택하고 다음 세부 정보를 추가하고 **저장**합니다.\
    **이름**: roles\
    **원본**: 속성\
    **원본 속성**: user.assignedroles

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-10 at 2.55.05 PM.png" alt="사용자 지정 클레임"><figcaption><p>사용자 지정 클레임</p></figcaption></figure>

위 단계를 완료하면 Snyk 연락 담당자에게 연락하여 설정이 완료되도록 합니다.
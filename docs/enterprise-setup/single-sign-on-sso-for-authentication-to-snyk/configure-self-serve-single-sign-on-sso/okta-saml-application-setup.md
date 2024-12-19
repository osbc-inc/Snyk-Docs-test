# Okta SAML 응용 프로그램 설정

이 예제에서는 Okta SAML 응용 프로그램을 설정하고 이를 SSO를 용이하게 하는 Snyk에 연결하는 방법을 보여줍니다. Okta를 구성하여 Snyk과 SSO를 사용하려면 Snyk로부터 엔터티 ID 및 응답 URL(답변 컨슈머 서비스 URL)가 필요합니다.

1. 상단 왼쪽에서 드롭다운에서 **GROUP OVERVIEW**를 선택한 다음 그룹 설정으로 이동하기 위해 코그 휠(우측 상단)을 클릭합니다.

    <figure><img src="../../../.gitbook/assets/1 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (3) (1).png" alt="그룹 개요 선택"><figcaption><p>그룹 개요 선택</p></figcaption></figure>
2. **SSO**를 클릭하고 **Entity ID** 및 **ACS URL** 아래의 값들을 복사하거나 브라우저 탭을 열어둡니다.

    <figure><img src="../../../.gitbook/assets/2 (1) (1) (1) (1).png" alt="그룹 설정: SSO"><figcaption><p>그룹 설정: SSO</p></figcaption></figure>
3. [Okta](https://www.okta.com/se/login/)로 이동하여 애플리케이션 메뉴를 열고 **앱 통합 만들기**를 클릭합니다.

    <figure><img src="../../../.gitbook/assets/1 (5) (1).png" alt="Okta 애플리케이션 메인 페이지"><figcaption><p>Okta 애플리케이션 메인 페이지</p></figcaption></figure>
4. **SAML 2.0**을 선택하고 애플리케이션을 적절하게 명명한 후 **다음**을 클릭합니다.

    <figure><img src="../../../.gitbook/assets/2 (3).png" alt="Okta SAML 애플리케이션 생성"><figcaption><p>Okta SAML 애플리케이션 생성</p></figcaption></figure>
5. Snyk에서 복사한 Entity ID 및 사인온 URL을 해당 필드에 추가합니다.

    <figure><img src="../../../.gitbook/assets/3 (4).png" alt="Okta에서 SSO 세부사항 추가"><figcaption><p>Okta에서 SSO 세부사항 추가</p></figcaption></figure>
6. **속성 명령문**으로 이동하여 다음과 같이 세 가지 속성 이름으로 값을 추가합니다:

    * **이름**: 이메일, **값**: user.email
    * **이름**: 이름, **값**: user.firstName + ' ' + user.lastName
    * **이름**: 사용자 이름, **값**: user.login

        <figure><img src="../../../.gitbook/assets/5 (2) (1) (1) (1) (1).png" alt="Okta 속성 명령문 추가"><figcaption><p>Okta 속성 명령문 추가</p></figcaption></figure>

    이제 **다음**을 클릭하여 피드백 세부사항을 입력하거나 다음 단계로 이동합니다.
7. Okta 애플리케이션 목록을 다시 열고 새로 생성한 애플리케이션 및 **사인 온** 탭을 클릭합니다. 페이지 오른쪽에서 **SAML 설정 지침 보기**를 클릭한 다음 표시되는 페이지에서 **IDP 단일 로그인 URL**과 **X.509 인증서**를 복사합니다.
8. 이전 페이지로 돌아가 **할당** 탭으로 이동합니다. **할당**을 클릭한 다음 사용자, 그룹 또는 둘 다 중에 필요에 따라 선택합니다.

    <figure><img src="../../../.gitbook/assets/7 (1) (2).png" alt="SSO 애플리케이션 할당"><figcaption><p>SSO 애플리케이션 할당</p></figcaption></figure>
9. Snyk 포털로 돌아가서 단계 2로 이동하고 단계 7에서 얻은 세부사항과 사용하려는 도메인들을 입력하고, **IdP 시작 워크플로**를 활성화해야 하는지 확인한 후, **Auth0 연결 만들기**를 클릭합니다.

    <figure><img src="../../../.gitbook/assets/8 (3).png" alt="Snyk SSO 단계 2"><figcaption><p>Snyk SSO 단계 2</p></figcaption></figure>
10. 단계 3로 이동하여 새로운 사용자가 로그인할 때 어떻게 처리해야 하는지 결정합니다. 사용하려는 옵션을 선택하고 **프로필 속성**을 Okta에서 구성한 대로 입력한 후, **변경 사항 저장**을 클릭하고 상단 단계 3의 직접 URL로 로그인할 수 있는지 또는 [일반 SSO 로그인](https://app.snyk.io/login/sso)으로 이동하여 로그인할 수 있는지 확인합니다.

    <figure><img src="../../../.gitbook/assets/9 (1) (1) (1) (1) (1).png" alt="프로필 속성"><figcaption><p>프로필 속성</p></figcaption></figure>
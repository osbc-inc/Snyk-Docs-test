# 예시: Okta OIDC 앱을 위한 사용자 정의 매핑 설정

다음 단계를 따라 Okta의 OIDC 통합을 구성하세요.

## Okta OIDC 앱 생성

1. Okta에서 **Applications -> Applications -> Create App Integration**을 선택한 다음 **OICD OpenID Connect** 및 **Web Application**을 선택합니다.

    <figure><img src="../../../../.gitbook/assets/1 (4).png" alt="Okta에서 새 앱 통합 생성"><figcaption><p>Okta에서 새 앱 통합 생성</p></figcaption></figure>
2. 다음 단계에서 OIDC 애플리케이션을 위한 **App 통합 이름**을 추가하고, **Implicit Grant Type**을 확인하고 Snyk 플랫폼 배포에 관련된 **Sign-in redirect URI**를 추가합니다. 자리 표시자 **Sign-out redirect URI**를 제거하고 **Save**를 클릭하기 전에 할당 액세스 제어를 선택합니다.

    <figure><img src="../../../../.gitbook/assets/2 (1) (1) (1).png" alt="새 웹 앱 통합에 대한 세부 정보 제공"><figcaption><p>새 웹 앱 통합에 대한 세부 정보 제공</p></figcaption></figure>
3. 저장 후 열리는 애플리케이션 페이지에서 [Snyk에 제공할 OIDC 정보](../../set-up-snyk-single-sign-on-sso.md#oidc-information-to-provide-to-snyk)을 복사하고 Snyk 담당자에게 제공합니다:
   * Client ID
   * Implicit Grant 유형을 사용하지 않는 경우 클라이언트 비밀
4. 또한 Snyk에게 발행자 URL/도메인을 공유해야 합니다. 이는 일반적으로 브라우저 주소 표시줄에서 "-admin"이 없는 URL입니다. 예를 들어 https://customer.example.okta.com입니다. 또한 애플리케이션의 **Sign-On** 탭에서 **OpenID Connect ID Token**을 **Dynamic** 발행자에서 **Okta URL**로 편집하면 찾을 수 있습니다.

사용자 정의 매핑을 설정하려면 본 가이드의 다음 섹션으로 이동하세요.

## 사용자 정의 매핑 추가

Okta에서 OIDC 애플리케이션을 위한 [사용자 정의 매핑](../)은 그룹 수준의 사용자 정의 속성을 통해 쉽게 관리할 수 있습니다.

### Snyk 조직 이름과 역할을 포함하는 사용자 정의 앱 사용자 속성 생성

1. Okta에서 **Directory -> Profile editor** 아래에 새로 생성된 OIDC 애플리케이션 사용자 프로필을 선택합니다.
2. **+Add Attribute**를 선택합니다.

    <figure><img src="../../../../.gitbook/assets/3 (3) (1).png" alt="Okta 프로필 편집기"><figcaption><p>Okta 프로필 편집기</p></figcaption></figure>
3. 해당 필드에 다음 속성에 대한 세부 정보를 추가하고 **Save**를 클릭합니다:\
    **데이터 유형**: string 배열\
    **표시 이름**: Snyk roles\
    **변수 이름**: roles\
    **그룹 우선 순위**: 그룹 간 값 병합

    <figure><img src="../../../../.gitbook/assets/4 (4).png" alt="속성 세부 정보"><figcaption><p>속성 세부 정보</p></figcaption></figure>

### 애트리뷰트를 관련 Okta 그룹에 할당

1. Okta 메인 페이지에서 **Directory -> Groups**를 선택합니다.
2. **그룹**을 선택한 다음 **Applications** 탭으로 이동하여 이미 할당되어 있지 않은 경우 **애플리케이션 할당**을 클릭하고 Snyk OIDC 앱을 선택합니다. 그런 다음 표시된 Snyk OIDC 앱 옆의 **연필**을 클릭합니다.

    <figure><img src="../../../../.gitbook/assets/5 (1) (1) (1) (1) (1) (1) (1).png" alt="수정 대상 그룹 선택"><figcaption><p>수정 대상 그룹 선택</p></figcaption></figure>
3. **편집 앱 할당** 대화 상자에서 Okta 그룹과 연관된 Snyk 조직 이름 + 역할(공백이나 대문자가 없는 형식)을 추가하고, [사용자 정의 매핑](../)에 설명된 구문을 따릅니다 (또는 레거시 매핑 옵션을 사용하는 경우 [레거시 사용자 정의 매핑](../legacy-custom-mapping.md)에 설명된대로). 예시: `snyk:org:*:org_admin`.

    <figure><img src="../../../../.gitbook/assets/image (385).png" alt="Snyk 역할 추가"><figcaption><p>Snyk 역할 추가</p></figcaption></figure>
4. 모든 적용 가능한 Okta 그룹에 대해 이전 단계를 반복하여 각 구성된 그룹 내 각 사용자에게 조직 이름과 역할 조합을 할당하세요.
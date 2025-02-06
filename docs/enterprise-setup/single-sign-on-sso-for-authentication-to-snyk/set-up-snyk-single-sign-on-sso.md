# Snyk Single Sign-On(SSO) 설정하기

기존 SSO 제공업체를 통해 개발자 및 팀이 Snyk에 쉽게 액세스하도록 단일 로그인(SSO)을 설정하십시오.

Snyk와 신원 제공자 간의 신뢰를 설정하는 데 필요한 정보는 사용 중인 SSO 유형에 따라 다릅니다.

신규 사용자가 할당될 그룹 및 조직이 표시되도록 최소한 하나의 그룹 및 조직이 있어야 합니다. 자세한 내용은 [그룹 및 조직 관리](../../snyk-admin/groups-and-organizations/)를 참조하십시오.

{% hint style="info" %}
다음 섹션에서 식별된 필요한 정보를 수집한 후 SSO 설정을 요청하기 위해 지원 티켓을 생성하십시오.

또는 그룹 관리자가 Snyk 단일 로그인(SSO)을 구성할 수 있습니다. 단골식 단일 로그인(SSO) 구성 방법에 대한 단골 사용 지침서를 참조하십시오.
{% endhint %}

## SSO 개요

신원 제공자(IdP)와 Snyk 간의 신뢰를 수립하는 과정은 SSO 관리자와 Snyk 지원 간에 조정된 몇 가지 단계를 필요로 합니다.

* 신원 제공자 플랫폼에서 Snyk 환경 및 사용자 속성에 대한 세부 정보 입력
* 신원 제공자로부터 Snyk에 대한 세부 정보 제공
* 테스트를 위해 사용자를 설정하고 해당 사용자의 사용자 이름과 암호를 Snyk에 제공
* Snyk에서 제공한 링크를 사용하여 페이로드 생성
* Snyk가 연결을 완료한 후 로그인 프로세스가 올바르게 작동하는지 확인

사용자가 Snyk에 로그인할 때 사용자가 provision됩니다. 초대가 필요한 경우, 관리자가 해당 기관에 추가하면 사용자는 적절한 기관들만 볼 수 있습니다.

SSO가 Snyk와 회사의 네트워크로부터 둘 다 구성된 후에는 Snyk, Snyk 대리 (Auth0), 그리고 회사의 네트워크 간에 신뢰 관계가 확립됩니다. 민감한 데이터는 사용자 로그인을 활성화하기 위해 Auth0에만 암호화되어 저장됩니다.

SSO 연결 유형 마다 신원 제공자와 Snyk 사이의 신뢰를 수립하는 데 필요한 다른 세부 정보가 필요합니다. 다음 섹션에서 해당 세부사항을 설명합니다. 이러한 세부 정보는 또한이 문서 끝의 자료 구간에 있는 워크시트에 포함되어 있습니다.

## SAML을 위한 SSO 설정

Snyk와의 신뢰를 수립하기 위해 신원 제공자에 Entity ID, Assertion Consumer Service (ACS) URL 및 Signing certificate를 추가하십시오.

* **Entity ID**는 SAML 엔터티 또는 서비스 제공 업체로 Snyk을을 고유하게 식별하는 URL입니다. 참고: **기본 Entity ID는 수동으로 확인되어야 합니다.** 기본값이 설정되어 있지 않습니다.
* \*\*Assertion Consumer Service (ACS)\*\*는 신원 제공자에서 Snyk 네트워크로의 요청을 수신하여 사용자와 Snyk 간의 통신을 가능하게하는 엔드포인트입니다. 이 URL은 때로는 응답 URL이라고도합니다.
* **Signing certificate**는 신원 제공자에 저장된 Snyk 인증서로, 신뢰 관계를 유지하는 데 필요합니다. 인증을 위한 필요한 암호화 키를 포함합니다.

이러한 세부 정보를 사용하여 신원 제공자(IdP)와의 연결을 설정하십시오:

| **세부 정보**                                      | **설명**                                                                                                                                                                      |
| ---------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Entity ID                                      | **urn:auth0:snyk:saml-{group-name-normalized}**                                                                                                                             |
| Entity ID (Snyk EU Tenant Customers)           | **urn:auth0:snyk-mt-eu-prod-1:saml-{group-name-normalized}**                                                                                                                |
| Entity ID (Snyk AU Tenant Customers)           | **urn:auth0:snyk-mt-au-prod-1:saml-{group-name-normalized}**                                                                                                                |
| ACS URL                                        | [https://snyk.auth0.com/login/callback?connection=saml-](https://snyk.auth0.com/login/callback?connection=saml-)**{group-name-normalized}**                                 |
| ACS URL (Snyk EU Tenant Customers)             | [https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback?connection=saml-](https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback?connection=saml-)**{group-name-normalized}** |
| ACS URL (Snyk AU Tenant Customers)             | [https://snyk-mt-au-prod-1.au.auth0.com/login/callback?connection=saml](https://snyk-mt-au-prod-1.au.auth0.com/login/callback?connection=saml)-**{group-name-normalized}**  |
| Signing certificate                            | [https://snyk.auth0.com/pem](https://snyk.auth0.com/pem)                                                                                                                    |
| Signing certificate (Snyk EU Tenant Customers) | [https://snyk-mt-eu-prod-1.eu.auth0.com/pem?cert=connection](https://snyk-mt-eu-prod-1.eu.auth0.com/pem?cert=connection)                                                    |
| Signing certificate (Snyk AU Tenant Customers) | [https://snyk-mt-au-prod-1.au.auth0.com/pem?cert=connection](https://snyk-mt-au-prod-1.au.auth0.com/pem?cert=connection)                                                    |

{% hint style="info" %}
{group-name-normalized}을 Snyk 그룹 이름으로 대체하십시오. 그룹 이름에 공백이 포함된 경우 대시로 교체하십시오. 예를 들어, 그룹 이름이 `Your Company Group`이면 **{group-name-normalized}** 값은 **your-company-group**이 됩니다.
{% endhint %}

신원 제공자에서 Snyk로 정보를 매핑하려면 사용자 속성을 다음과 같이 지정하십시오. 동일한 대문자와 철자를 사용하십시오:

| **속성**   | **설명**         |
| -------- | -------------- |
| email    | 사용자 이메일 주소     |
| name     | 인증을 받을 사람의 이름  |
| username | 신원 제공자의 사용자 이름 |

사용자 속성이 일치하지 않는 경우, Snyk의 SSO 구성에 더 많은 시간이 소요될 수 있음을 유의하십시오.

## Snyk에 제공할 SAML 정보

신원 제공자에서 다음 정보를 얻으십시오. 이 정보를 Snyk에 제공하여 서비스 제공자 측의 신뢰를 설정하십시오.

| 정보                | 설명                                                                                                                                     |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| 로그인 URL           | 신원 제공자 로그인 페이지의 URL                                                                                                                    |
| X509 서명 인증서       | Base64 형식으로 인코딩된 신원 제공자 공개 키                                                                                                           |
| 로그아웃 URL          | 선택 사항, 권장 사항 - 사용자가 Snyk에서 로그아웃할 때 리디렉션하기 위한 URL                                                                                       |
| 사용자 ID 속성         | 기본값은 `http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier` (중요: 이 값이 Snyk 사용자를 고유하게 식별하고 변경되는 경우 중복 사용자가 생성될 수 있습니다.) |
| 프로토콜 바인딩          | HTTP-POST가 권장되며, HTTP-Redirect도 지원됨                                                                                                    |
| IdP 초기화 플로우 지원 여부 | IdP 초기화 플로우는 보안 위험이 있으므로 권장되지 않습니다. 활성화하기 전에 위험을 이해하십시오.                                                                               |
| 이메일 도메인 및 서브도메인   | SSO에 액세스해야 하는 이메일 도메인 및 서브도메인                                                                                                          |

## OpenID Connect (OIDC)를 위한 SSO 설정

{% hint style="info" %}
IdP(또는 발급자 URL)은 공개적으로 액세스 가능해야 합니다. 이 수단을 공개로 만들 수 없는 경우, OIDC 대신 SAML을 사용해야 합니다.
{% endhint %}

신원 제공자와 Snyk 간의 연결에 OIDC를 사용하는 경우, 신원 제공자에 Callback/Redirect URIs 및 OAuth Grant Type을 추가하여 Snyk와의 신뢰를 수립하십시오.

| 정보                                                | 설명                                                                                                             |
| ------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Callback/Redirect URIs                            | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback)                                 |
| Callback/Redirect URIs (Snyk EU Tenant Customers) | [https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback](https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback) |
| Callback/Redirect URIs (Snyk AU Tenant Customers) | [https://snyk-mt-au-prod-1.au.auth0.com/login/callback](https://snyk-mt-au-prod-1.au.auth0.com/login/callback) |
| OAuth Grant Type                                  | Implicit(또는 Authorization Code)                                                                                |

## Snyk에 제공할 OIDC 정보

신원 제공자에서 다음 정보를 얻으십시오. 이 정보를 서비스 제공자 측에서 신뢰를 수립하기 위해 Snyk에 제공하십시오.

| 정보              | 설명                                                                  |
| --------------- | ------------------------------------------------------------------- |
| 발급자 URL         | 연결하려는 OpenID Connect 제공업체의 디스커버리 문서 URL. 이 URL은 공개적으로 액세스 가능해야 합니다. |
| 클라이언트 ID        | 권한 부여 서버에 대한 고유한 공개 식별자                                             |
| 클라이언트 시크릿       | IdP가 암시적 그랜트 유형을 허용하지 않는 경우에만 필요합니다.                                |
| 이메일 도메인 및 서브도메인 | SSO에 액세스해야 하는 이메일 도메인 및 서브도메인                                       |

## Entra ID로 SSO 설정(앱 등록/OIDC 사용)

신원 제공자와 Snyk 간의 연결에 Entra ID(이전 Azure AD)를 사용할 때, 신원 제공자에 Redirect URIs를 추가하여 Snyk와의 신뢰를 수립해야 합니다.

{% hint style="info" %}
연결 시 Entra ID 이름을 사용자 계정 이름 대신 사용하십시오. 그렇지 않으면 연결 오류가 발생할 수 있습니다.
{% endhint %}

| 정보                                       | 설명                                                                                                             |
| ---------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Redirect URIs                            | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback)                                 |
| Redirect URIs (Snyk EU Tenant Customers) | [https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback](https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback) |
| Redirect URIs (Snyk AU Tenant Customers) | [https://snyk-mt-au-prod-1.au.auth0.com/login/callback](https://snyk-mt-au-prod-1.au.auth0.com/login/callback) |

## Snyk에 제공할 Entra ID 정보

신원 제공자에서 다음 정보를 얻으십시오. 이 정보를 서비스 제공자 측에서 신뢰를 수립하기 위해 Snyk에 제공하십시오.

| 정보                     | 설명                                                |
| ---------------------- | ------------------------------------------------- |
| 클라이언트 ID               | 권한 부여 서버에 대한 고유한 공개 식별자                           |
| 클라이언트 시크릿              | 권한을 부여하여 인가된 요청자에게 토큰을 부여하는 권한을 부여하는 곳의 시크릿       |
| Microsoft Entra ID 도메인 | Snyk 앱 계정의 개요에서 찾을 수 있는 디렉토리(테넌트) ID에 표시된 숫자 및 문자 |

## ADFS로 SSO 설정

신원 제공자(IpD)와 Snyk 간의 연결에 Active Directory Federation Service (ADFS)를 사용하는 경우, 신원 제공자에 Realm 식별자, Callback URL 및 Signing certificate를 추가하여 Snyk와의 신뢰를 수립해야 합니다. 더 많은 정보를 원하시면 [Auth0와 ADFS 서버 연결하기(video)](https://www.youtube.com/watch?v=ICW6sGP9ht8)를 참조하십시오.

| **정보**                                  | **설명**                                                                                                                                     |
| --------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| Realm 식별자                               | urn:auth0:snyk                                                                                                                             |
| EU Realm 식별자                            | urn:auth0:snyk-mt-eu-prod-1                                                                                                                |
| AU Realm 식별자                            | urn:auth0:snyk-mt-au-prod-1                                                                                                                |
| Callback URL                            | [https://snyk.auth0.com/login/callback](https://snyk.auth0.com/login/callback)                                                             |
| Callback URL (Snyk EU Tenant Customers) | [https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback](https://snyk-mt-eu-prod-1.eu.auth0.com/login/callback)                             |
| Callback URL (Snyk AU Tenant Customers) | [https://snyk-mt-au-prod-1.au.auth0.com/login/callback](https://snyk-mt-au-prod-1.au.auth0.com/login/callback)                             |
| 서명 인증서                                  | [https://snyk.auth0.com/pem](https://snyk.auth0.com/pem) (암호화가 아닌 서명으로 추가)                                                                 |
| 서명 인증서 (Snyk EU Tenant Customers)       | [https://snyk-mt-eu-prod-1.eu.auth0.com/pem?cert=connection](https://snyk-mt-eu-prod-1.eu.auth0.com/pem?cert=connection) (암호화가 아닌 서명으로 추가) |
| 서명 인증서 (Snyk AU Tenant Customers)       | [https://snyk-mt-eu-prod-1.au.auth0.com/pem?cert=connection](https://snyk-mt-eu-prod-1.au.auth0.com/pem?cert=connection) (암호화가 아닌 서명으로 추가) |

## Snyk에 제공할 ADFS 정보

ID 공급자로부터 다음 정보를 수집하세요. 이 정보를 Snyk에 제공하여 서비스 제공자 측에서 신뢰를 설정합니다.

* **ADFS URL 또는 Federation Metadata XML 파일**

***

## 엔터프라이즈 사용자 매핑

엔터프라이즈 플랜의 경우, Snyk은 사용자가 SSO를 통해 처음 로그인할 때 특정 조직과 역할에 새 사용자를 매핑할 수 있습니다. 이 옵션에는 조직 이름 명명 규칙을 포함한 추가 구성이 필요합니다.

{% hint style="info" %}
이 SSO 옵션 구현을 준비하려면 Snyk 계정 팀과 협력하세요.
{% endhint %}

***

## SSO 연결 완료

ID 공급자와 연결을 설정하고 Snyk 지원팀에 필요한 세부 정보를 제공한 후, Snyk는 페이로드를 생성할 수 있는 링크를 제공합니다.

{% hint style="info" %}
이 링크를 처음 클릭했을 때 표시되는 오류 메시지는 무시하세요. Snyk는 생성된 페이로드를 사용하여 구성을 완료합니다.
{% endhint %}

Snyk가 구성을 완료한 후, 지원 에이전트는 쿠키가 로그인 프로세스를 방해하지 않도록 시크릿 모드에서 로그인 페이지로 이동할 것을 요청합니다.

프로덕션 환경에 로그인하려면 [https://app.snyk.io/login/sso](https://app.snyk.io/login/sso)를 사용하세요.

로그인을 완료하려면 다음 단계를 따르세요:

1. 이메일 주소를 입력하세요.
2. **Continue to provider**를 선택하세요.
3. 다른 애플리케이션에서와 마찬가지로 ID 공급자를 통해 로그인하세요.
4. 그룹 관리자로 승격할 사용자를 Snyk 지원팀에 알려주세요.

## SSO 설정을 위한 리소스

다음 워크시트에는 ID 공급자에 입력해야 할 정보와 단일 사인온(Single Sign-On)을 요청하기 위해 Snyk 지원팀에 티켓을 제출하기 전에 수집해야 할 정보가 포함되어 있습니다.

{% file src="../../.gitbook/assets/SSO Azure Worksheet (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (3).pdf" %}

{% file src="../../.gitbook/assets/SSO SAML Worksheet (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).pdf" %}

{% file src="../../.gitbook/assets/SSO ADFS Worksheet (2) (1).pdf" %}

{% file src="../../.gitbook/assets/SSO OIDC Worksheet (1) (1) (1) (1) (1) (1) (1) (1).pdf" %}

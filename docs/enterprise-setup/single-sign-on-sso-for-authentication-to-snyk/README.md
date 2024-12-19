# Snyk의 인증을위한 단일 로그인 (SSO)

{% hint style="info" %}
**기능 가용성**\
SSO는 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

## SSO 개요

귀하의 회사의 기존 식별 관리 시스템을 활용하고 직원들이 기업 신분을 사용하여 Snyk에 로그인할 수 있습니다. 이는 사용자에게 Snyk를 제공하기를 쉽게 만듭니다. 또한 그룹 및 조직 구성원, 역할 기반 액세스 및 기타보다 깊은 통합을 가능하게 합니다.

<figure><img src="../../.gitbook/assets/image (1) (4).png" alt="&#x22;&#x22;"><figcaption><p>SSO로 Snyk에 로그인</p></figcaption></figure>

Snyk는 SAML 및 OpenID Connect (OIDC) 기반 SSO 및 ADFS와 통합할 수 있습니다. 또한 SSO에 기업 식별 공급자를 사용할 수도 있으며, 이는 [Entra ID](https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-whatis) (이전 Azure AD) 및 [Google G Suite](https://community.snowflake.com/s/article/configuring-g-suite-as-an-identity-provider)를 포함합니다. SAML에 대해 자세히 알아보려면 [Auth0 설명서](https://auth0.com/docs/protocols/saml)를 참조하십시오.

[SSO, 인증 및 사용자 프로비저닝](https://learn.snyk.io/lesson/sso-authentication-provisioning/)에서 교육을 받을 수 있습니다.

## SSO를 위한 사용자 인증 및 프로비저닝

SSO가 구성된 경우, 사용자가 처음으로 SSO를 통해 로그인하면 새로운 Snyk 계정이 제공됩니다. 심지어 이전에 자체 계정을 만들었던 경우에도 그렇습니다.

로그인 프로세스는 다음 단계를 포함합니다:

1. 사용자들이 Snyk.io에서 SSO를 선택하여 로그인하면 요청한 식별 공급자로 리디렉션되어 인증됩니다.
2. 식별 공급자는 이 인증을 Snyk 서버로 전달하여 각 사용자를 생성하는 데 필요한 데이터를 보냅니다.
3. Snyk는 해당 사용자를 위해 디렉터리를 확인합니다.
4. 사용자가 이미 구성되어 있는 경우, Snyk는 적절한 액세스를 활성화합니다. 새로운 사용자의 경우, Snyk는 사용자를 디렉터리에 추가하고 그 사용자를 Snyk.io로 리디렉션하여 적절한 액세스를 제공합니다.
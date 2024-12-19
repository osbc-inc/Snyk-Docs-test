# SSO 구성

SSO 구성은 프로젝트를 도입하거나 조직 구조를 만들거나 통합을 구성하는 동안 수행할 수 있습니다. Snyk에 프로젝트를 도입하기 전에 사용자를 추가하려는 의도가 있다면, 조직의 알림을 비활성화하세요.

[Snyk와 단일 로그인(SSO) 설정](../../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/)은 사용자가 응용 프로그램에 로그인하는 일관된 방법을 제공하면서 귀하의 조직에 액세스 권한이 있는 사용자를 제어하는 능력을 향상시킵니다.

대부분의 경우, Snyk는 당신이 신분 공급자에 SAML 연결을 구성할 수 있는 [셀프 서비스 단일 로그인 옵션](../../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/configure-self-serve-single-sign-on-sso/)을 사용하는 것을 권장합니다. 이를 통해 유효한 이메일 도메인을 구성할 수 있고, 또한 새로운 사용자가 조직에서 갖게 될 기본 권한도 구성할 수 있습니다.

## 도움 받는 방법

Snyk와의 라이선스에 따라 그룹 설정 페이지에 SSO 옵션이 표시되지 않을 수 있습니다. 이 경우 [지원에 요청](https://support.snyk.io)하여 활성화할 수 있습니다. 이 옵션을 사용하면 SAML로 SSO를 구성할 수 있지만, 다른 SSO 구성 옵션을 사용하려면(예: OIDC), [Snyk 지원팀에 요청](https://support.snyk.io)하여 지원받을 수 있습니다.

## SSO 사용자 정의 매핑에 대한 지원

이 옵션을 사용하려면 유료 전문 서비스가 필요하며, 셀프 서비스 SSO 옵션을 사용하여 완료할 수 없습니다. 필요한 서비스를 구매하려면 [Snyk 지원팀](https://support.snyk.io) 또는 Snyk 계정 팀에 문의하십시오.

## API를 통한 사용자 프로비저닝 방법

Snyk API에서 [Snyk API를 통한 사용자 프로비저닝 엔드포인트](../../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md)를 사용하면 사용자가 Snyk 플랫폼에 로그인하기 전에 조직 수준의 액세스 및 권한을 부여할 수 있습니다.

API 프로비저닝을 사용하면 조직 액세스를 제한하고 사용자에게 사용자 지정 멤버 역할을 할당할 수 있습니다. 이를 통해 다음을 제어할 수 있습니다:

- 각 사용자에게 할당되는 역할
- 각 사용자가 액세스해야 하는 조직
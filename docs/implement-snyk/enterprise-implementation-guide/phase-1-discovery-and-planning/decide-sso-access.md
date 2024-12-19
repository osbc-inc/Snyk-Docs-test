# SSO 접근 결정

## Single Sign-On (SSO) 설정 결정

SSO에 대한 자세한 정보는 [Using Single Sign-On (SSO) for authentication](../../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/)을 참조하십시오.

일반적으로 초기 시험판 또는 평가 팀이 개인 인증을 사용하는 것이 일반적이지만, 사용자에게 Snyk를 제공하기 전에 SSO를 구현해야 합니다. 다음을 포함하여 SSO를 구현해야 합니다:

- 적절한 SSO 계정을 만들고 해당 계정을 관리자로 승격시킵니다.
- 넓은 그룹에게 Snyk를 제공하기 전에 개인 계정을 제거합니다.

## SSO 프로비저닝 옵션

회사의 새로운 사용자에게 액세스를 설정할 때 사용자 경험과 액세스를 결정하는 몇 가지 접근 방식을 고려해야 합니다. 이 단계에서 **Open To All** 또는 **Require an Invite**를 고려하고, **Custom Mapping**이 필요한지 확인합니다.

자세한 정보는 [Choose a Provisioning Option](../../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/choose-a-provisioning-option.md)을 참조하십시오.

**Custom Mapping** 기능은 [유료 서비스](../../../working-with-snyk/snyk-terms-of-support-and-services-glossary/)에서만 제공됩니다. 이 기능을 사용하면 사용자 계정을 사용자의 특정 요구 사항에 따라 액세스 및 역할을 맞춤화하는 사용자 정의 규칙으로 프로비저닝할 수 있습니다. 이 기능은 복잡한 구성 및 설정이 필요할 수 있으며, 고급 특정 액세스 관리 요구 사항을 가진 조직에게 권장됩니다.
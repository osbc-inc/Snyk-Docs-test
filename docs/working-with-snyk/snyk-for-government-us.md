# Snyk for Government (US)

[Snyk for Government (US)](https://snyk.io/government-security-solution/)는 미국 연방기관이 빠르고 안전하게 개발할 수 있도록 합니다. 개발자가 이미 사용하는 도구 및 작업 흐름과 통합하여, Snyk for Government (US)는 이러한 기관들이 소프트웨어 개발 라이프사이클에서 왼쪽으로 이동할 수 있도록 설정하여 시작부터 안전한 개발을 가능하게 합니다.

Snyk for Government (US)는 [FedRAMP](https://www.fedramp.gov/) 및 [NIST](https://www.nist.gov/) 보안 통제 요구 사항을 준수하기 때문에 연방 기관들은 제품이 미국 정부가 설정한 보안 표준을 준수함을 확신할 수 있습니다.

Snyk for Government (US)는 미국 연방 정부로 배포할 수 있도록 표준 Snyk 제품과 차이가 있습니다. FedRAMP 및 NIST 통제 요구 사항을 준수하는 것은 표준 Snyk 제품의 일부 기능이 FedRAMP 환경에서 지원되지 않는다는 것을 의미합니다.

이 목록은 Snyk for Government (US) 제품의 기능 차이를 식별합니다.

## 핵심 제품 가용성 제한

* Snyk Code:
  * Code Search는 포함되지 않음
  * DeepCode AI Fix는 포함되지 않음
* Snyk Container: [Kubernetes 통합](https://docs.snyk.io/scan-applications/snyk-container/kubernetes-integration/kubernetes-integration-overview)은 제외됨
* Snyk Open Source:
  * Unmanaged C++가 포함되지 않음
  * npm 패키지 `@snyk/protect` 및 `@snyk/fix`가 포함되지 않음
* Snyk AppRisk는 **사용 불가**.

## API 키 사용 불가

API 키를 사용할 수 없습니다.

따라서 `auth_type`이 `api_key`인 API를 사용하여 UI를 통해 또는 API를 사용하여 서비스 계정을 생성하려고 하는 시도는 허용되지 않습니다. API 키가 일반적으로 사용되는 모든 시나리오에 대해 OAuth 프로토콜을 대신 사용해야 합니다. 자세한 내용은 [OAuth 2.0를 사용한 서비스 계정](https://docs.snyk.io/enterprise-configuration/service-accounts/service-accounts-using-oauth-2.0)을 참조하십시오. 도움이 필요한 경우, [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.

또한 CLI는 토큰 기반 인증이 아닌 OAuth 모드에서 사용해야 합니다.

## Single Sign-On 가용성 제한

[단일 로그인 (SSO)](../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/)은 [셀프 서비스 단일 로그인 (SSO)](../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/configure-self-serve-single-sign-on-sso/)를 제외하고 사용할 수 있습니다. 모든 SSO 설정은 Snyk에서 관리됩니다. 단일 로그인 구성 단계에는 약간의 차이가 있습니다:

* 서비스 제공자는 Auth0 대신 Okta입니다.
* ACS URL 및 엔터티 ID 및 인증서는 각 연결마다 다르며 Snyk SSO 설명서와 일치하지 않을 것입니다.
* ACS Url, 엔터티 ID 및 인증서를 얻으려면 Snyk가 Okta에서 연결을 부분적으로 제공해야 합니다.

자세한 내용은 [Snyk에 대한 인증을 위한 단일 로그인 (SSO)](../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/)를 참조하십시오.

## 사용할 수 없는 통합

[Gatekeeper 플러그인](../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/)은 OAuth 인증을 지원하기 때문에 **사용할 수 없습니다**.

## 보고서 및 데이터 사용 불가

* [레거시 보고서](../manage-issues/reporting/legacy-reports/)
* [이슈 분석](../manage-risk/enterprise-analytics/issues-analytics.md)
* [인사이트](../manage-risk/prioritize-issues-for-fixing/using-the-issues-ui-with-snyk-apprisk/)

## 플랫폼 기능 사용 불가

* [Bitbucket Cloud 앱](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-cloud-app.md). [Bitbucket Data Center/Server](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-data-center-server.md) 통합은 가능합니다.
* [Slack 앱](../integrate-with-snyk/jira-and-slack-integrations/slack-app.md)
* [Jira 앱](../integrate-with-snyk/jira-and-slack-integrations/snyk-security-in-jira-cloud-integration.md). [Jira 통합](../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md)은 가능합니다.
* [Snyk Advisor](https://snyk.io/advisor/)
* [Snyk Learn](https://learn.snyk.io/?)
* Google, GitHub 등의 소셜 로그인: 신원 제공자로 사용 불가
* Snyk 지원으로의 SSO(단일 로그인)
* [상태 페이지](https://status.snyk.io)
* 아웃바운드 웹훅
* 세션 동시성은 사용자 당 세션 세 개(3)까지로 제한됩니다.
* 세션 잠금: 세션이 만료된 후, 로그인한 사용자는 기존 세션 창에 있는 모든 데이터에 대한 액세스 권한을 상실합니다.
* 세션 타임아웃: 기본 세션 타임아웃은 더 짧습니다(15분). 자세한 내용은 [Snyk 그룹을 위한 세션 길이 구성](../snyk-admin/groups-and-organizations/groups/configure-session-length-for-a-snyk-group.md)을 참조하십시오.
* [Snyk CLI 도커 이미지](../snyk-cli/install-or-update-the-snyk-cli/#snyk-cli-in-a-docker-image). 이는 FIPS-검증 암호화를 지원하지 않으며 이를 수용할 수 있는 경우에만 사용해야 합니다.
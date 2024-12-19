# API를 위한 인증

이 섹션은 [API를 위한 인증](authenticate-for-the-api.md)에 대한 정보를 제공합니다. API 토큰을 얻는 방법 및 인증 헤더에서 사용하는 방법과 [Snyk API 토큰 권한 사용자가 제어할 수 있는 것](snyk-api-token-permissions-users-can-control.md)을 포함합니다.

새 API 토큰을 얻는 지침은 [Snyk API 토큰 폐기 및 재생성](revoke-and-regenerate-a-snyk-api-token.md)을 참조하십시오.

다음은 **API 토큰을 사용하는 시기** 및 **서비스 계정 토큰을 사용하는 시기**를 설명합니다.

Snyk API 토큰은 사용자 프로필 하단에 있는 개인 토큰입니다. Snyk API 토큰은 Snyk 계정과 연결되어 있으며 특정 조직과는 관려이 없습니다.

무료 및 팀 플랜 및 체험 사용자는 이 개인 토큰에만 엑세스할 수 있습니다. 개인 토큰은 로컬 또는 빌드 기계에서 실행 중인 Snyk CLI 및 IDE에서 수동으로 토큰을 설정할 때 사용할 수 있습니다. API 또는 CI/CD에 인증하는 경우에는 조심해서 개인 토큰을 사용하십시오.

엔터프라이즈 사용자는 프로필 하단에 개인 토큰과 서비스 계정 토큰에 액세스할 수 있습니다. 자세한 내용은 [서비스 계정](../../../enterprise-setup/service-accounts/)을 참조하십시오.

* **엔터프라이즈 사용자는 모든 종류의 자동화를 위해 서비스 계정을 사용해야 합니다.** 이는 CLI 또는 빌드 시스템 플러그인으로 CI/CD 스캔 및 API를 통한 자동화를 포함합니다.
* **엔터프라이즈 사용자는 개인 토큰을 사용**하여 다음과 같은 경우에:
  * 자신의 기계에서 CLI를 로컬로 실행
  * IDE에서 수동으로 인증
  * API 호출을 한 번 실행하여 예를 들어 무언가를 테스트

개인 Snyk API 토큰에 대한 자세한 내용은 [API를 위한 인증](authenticate-for-the-api.md) 및 [CLI 사용을 위한 인증](../../../snyk-cli/authenticate-to-use-the-cli.md)을 참조하십시오.
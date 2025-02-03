# Azure 통합: API

Snyk을 통해 Azure 구독을 API를 통해 등록하려면:

1. [인프라스트럭처 코드 (IaC) 템플릿 또는 Bash 스크립트를 다운로드하세요:](step-1-download-azure-app-registration-iac-template-or-script-api.md) 이를 통해 Snyk에 구독을 스캔할 권한을 부여합니다.
2. [Azure Active Directory 앱 등록을 생성하세요:](step-2-create-the-azure-ad-app-registration-api.md) 다운로드한 템플릿이나 스크립트를 사용하여.
3. [클라우드 환경을 생성하고 스캔하세요.](step-3-create-and-scan-a-snyk-cloud-environment-for-azure-api.md)

위 단계를 완료하면 다음을 수행할 수 있습니다:

* Snyk이 발견한 클라우드 구성 문제를 확인합니다. [클라우드 및 IaC+ 문제](../../../getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/) 참조.
* 클라우드 컨텍스트로 취약점을 우선순위를 매길 수 있습니다.

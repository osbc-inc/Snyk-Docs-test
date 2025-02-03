# Azure 통합: 웹 UI

다음은 API를 통해 Azure 구독을 Snyk에 등록하는 단계입니다:

1. [인프라스트럭처 as 코드 (IaC) 템플릿 또는 Bash 스크립트 다운로드:](step-1-download-azure-app-registration-iac-template-or-script-web-ui.md) 구독을 스캔할 권한을 Snyk에 부여하기 위해.
2. [Azure Active Directory 앱 등록 생성:](step-2-create-the-azure-ad-app-registration.md) 다운로드한 템플릿 또는 스크립트를 사용하여.
3. [클라우드 환경 생성 및 스캔.](step-3-create-and-scan-a-snyk-cloud-environment-for-azure-web-ui.md)

위 단계를 완료하면 다음을 수행할 수 있습니다:

* Snyk이 발견한 클라우드 구성 문제를 볼 수 있습니다. [클라우드 및 IaC+ 문제](../../../getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/) 참조.
* 클라우드 컨텍스트로 취약점을 우선순위로 지정할 수 있습니다.

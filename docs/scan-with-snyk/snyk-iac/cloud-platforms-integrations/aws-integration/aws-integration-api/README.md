# AWS 통합: API

API를 통해 Snyk에 AWS 계정을 온보딩하기 전에 AWS 계정에 액세스하고 읽기 전용 Identity and Access Management (IAM) 역할을 만들 권한이 있는 인증서가 필요합니다. [AWS 통합](../) 페이지의 사전 요구 사항을 참조하십시오.

다음은 API를 통해 AWS 계정을 온보딩하는 단계입니다:

1. [인프라스트럭처 코드 (IaC) 템플릿 다운로드](step-1-download-iam-role-iac-template.md): Snyk에 계정을 스캔할 권한을 부여하기 위해.
2. [AWS IAM 역할 만들기](step-2-create-the-snyk-iam-role-api.md): 다운로드한 템플릿을 사용하여.
3. [클라우드 환경 만들고 스캔하기.](step-3-create-and-scan-a-snyk-cloud-environment.md)

위 단계를 완료하면 다음을 수행할 수 있습니다:

* Snyk이 발견하는 클라우드 구성 문제 보기. [클라우드 및 IaC+ 문제](../../../getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/) 참조.
* 클라우드 컨텍스트를 사용하여 취약점에 우선 순위 부여.

# AWS 통합: 웹 UI

AWS 계정을 웹 UI를 통해 등록하기 위해선, AWS 계정 및 관련 자격 증명에 접근할 권한이 있어야 하며 읽기 전용 Identity and Access Management (IAM) 역할을 생성할 권한이 필요합니다. [AWS 통합](../) 페이지에서 전제 조건을 확인하세요.

다음은 웹 UI를 통해 AWS 계정을 등록하는 단계입니다:

1. [Infrastrcture as code (IaC) 템플릿 다운로드](step-1-download-iam-role-iac-template-web-ui.md)을 통해 Snyk에 계정을 스캔할 권한을 부여합니다.
2. 다운로드한 템플릿을 사용하여 [AWS IAM 역할 생성](step-2-create-the-snyk-iam-role.md)합니다.
3. [클라우드 환경 생성 및 스캔](step-3-create-and-scan-a-snyk-cloud-environment-web-ui.md).

위 단계를 완료하면 다음을 수행할 수 있습니다:

* Snyk이 발견한 클라우드 구성 문제를 확인합니다. [클라우드 및 IaC+ 문제](../../../getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/).
* 클라우드 컨텍스트로 취약점을 우선순위 지정합니다.

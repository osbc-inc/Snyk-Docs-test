# Google Cloud 통합: 웹 UI

Web UI를 통해 Google Cloud 프로젝트를 Snyk에 등록하는 단계는 다음과 같습니다:

1. [인프라 구성 코드 (IaC) 템플릿 다운로드](step-1-download-service-account-iac-template-web-ui.md)을 통해 Snyk에 프로젝트를 스캔할 권한을 부여합니다.
2. 다운로드한 템플릿을 사용하여 [Google 서비스 계정 생성](step-2-create-the-google-service-account-web-ui.md)합니다.
3. [Google 클라우드 환경을 생성하고 스캔합니다.](step-3-create-and-scan-a-cloud-environment-for-google-web-ui.md)

위 단계를 완료하면 다음을 수행할 수 있습니다:

* Snyk이 발견한 클라우드 구성 문제를 확인할 수 있습니다. [클라우드 및 IaC+ 문제 확인](../../../../../scan-with-snyk/snyk-iac/getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/)
* 클라우드 콘텍스트로 취약점의 우선순위를 정할 수 있습니다.
# Snyk 보안 스캔 작업을 사용한 Azure 파이프라인 통합

Snyk는 Azure 파이프라인을 포함한 Microsoft Azure 생태계 전반에 걸쳐 보안을 활성화시키며, 애플리케이션 및 컨테이너 취약점을 자동으로 발견하고 수정합니다.

Snyk Security Scan 작업은 Snyk와 Azure DevOps에서 지원하는 모든 언어에 대해 사용할 수 있습니다.

{% hint style="info" %}
Snyk Security Scan 작업은 , , 그리고 를 지원합니다. 파이프라인에 다른 제품을 포함할 계획이 있다면, [Snyk CLI](../../../snyk-cli/)를 사용하세요.
{% endhint %}

Azure 파이프라인용 미리 만들어진 작업은 Azure 인터페이스에서 빠르게 삽입할 수 있으며, 추가 코딩없이 파이프라인을 사용자 정의하고 자동화할 수 있습니다. 이러한 작업에는 Snyk 작업이 포함되어 있습니다.

일상적인 업무의 일환으로 보안 취약점과 오픈 소스 라이선스 문제를 테스트하기 위해 파이프라인에 Snyk 작업을 포함할 수 있습니다. 이를 통해 응용프로그램 의존성과 컨테이너 이미지를 보안 취약점에 대해 테스트하고 모니터링할 수 있습니다. 테스트가 완료되면, Azure 파이프라인 출력뿐만 아니라 Snyk 인터페이스에서도 결과를 확인하고 작업할 수 있습니다.

설정 및 사용 세부 정보는 다음 페이지를 참조하세요:

* [Snyk Security Scan 작업이 작동하는 방법](how-the-snyk-security-scan-task-works.md)
* [Azure 파이프라인용 Snyk 확장 프로그램 설치](install-the-snyk-extension-for-your-azure-pipelines.md)
* [Snyk 보안 작업을 파이프라인에 추가](add-the-snyk-security-task-to-your-pipelines.md)
* [Snyk Security Scan 작업 매개변수와 값](snyk-security-scan-task-parameters-and-values.md)
* [사용자지정 API 엔드포인트](regional-api-endpoints.md)
* [Node.js (npm) 기반 응용프로그램을 테스트하기 위한 Snyk 작업 예시](example-of-a-snyk-task-to-test-a-node.js-npm-based-application.md)
* [응용프로그램을 테스트하기 위한 간단한 Snyk 작업 예시](simple-example-of-a-snyk-task-to-test-an-application.md)
* [컨테이너 이미지 파이프라인을 위한 Snyk 작업 예시](example-of-a-snyk-task-for-a-container-image-pipeline.md)
* [컨테이너 이미지를 테스트하기 위한 간단한 Snyk 작업 예시](simple-example-of-a-snyk-task-to-test-a-container-image.md)
* [응용프로그램 코드를 테스트하기 위한 Snyk 테스트 예시](example-of-a-snyk-task-to-test-application-code.md)
* [코드 테스트를 실행하기 위한 간단한 Snyk 작업 예시](simple-example-of-a-snyk-task-to-run-a-code-test.md)
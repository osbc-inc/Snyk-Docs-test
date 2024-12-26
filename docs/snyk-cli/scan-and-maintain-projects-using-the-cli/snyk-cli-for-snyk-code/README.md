# Snyk CLI를 사용하여  테스트하기

[Snyk Command Line Interface](../../) (CLI)를 사용하면 의 기능을 개발 워크플로우에 통합할 수 있습니다. Snyk CLI를 사용하면 로컬에서  테스트를 실행하거나 CI/CD 파이프라인에 통합하여 소스 코드를 보안 취약점으로 스캔할 수 있습니다.

## Snyk CLI를 와 함께 사용하는 데 필요한 사전 조건

{Snyk Code를 사용하여 소스 코드를 테스트하기 위해 Snyk CLI를 사용하기 전에 다음 사전 조건을 충족하는지 확인하세요:

- Snyk 계정.
- [지원되는 언어 및 프레임워크로 된 코드가 있는 저장소](../../../supported-languages-package-managers-and-frameworks/).
- **** 옵션이 [Snyk Org 설정에서 활성화되어 있어야 합니다](../../../scan-with-snyk/snyk-code/configure-snyk-code.md).
- Snyk CLI가 설치되어 있고 인증되어 있어야 합니다.
  - 지침은 [Snyk CLI 설치 또는 업데이트](../../install-or-update-the-snyk-cli/) 및 [Snyk CLI 인증](../../authenticate-to-use-the-cli.md)를 참조하세요.
  - 와 호환되는 최소 Snyk CLI 버전은 1.716.0입니다. Snyk는 최신 버전의 CLI를 사용하는 것을 권장합니다.

## Snyk CLI를 사용하여  테스트하기

Snyk CLI를 사용하여 리포지토리 코드를 테스트하려면 [`snyk code test`](../../commands/code-test.md) 명령을 사용하세요.\
자세한 정보는 [CLI를 사용하여 로 소스 코드 스캔하기](scan-source-code-with-snyk-code-using-the-cli.md)를 참조하세요.

CLI를 사용하여  테스트를 실행하기 전에 다음을 수행할 수 있습니다:

- [Snyk CLI 버전 업데이트](../../install-or-update-the-snyk-cli/).
- [CLI 테스트를 위한 조직 설정](set-the-snyk-organization-for-the-cli-tests.md).
- [{Snyk Code 테스트에서 디렉토리 및 파일 제외](exclude-directories-and-files-from-snyk-code-cli-tests.md).

snyk code test 명령 및 결과에 대한 정보는 이 섹션의 문서 페이지 및 [snyk-to-html](../cli-tools/snyk-to-html.md)를 참조하세요.
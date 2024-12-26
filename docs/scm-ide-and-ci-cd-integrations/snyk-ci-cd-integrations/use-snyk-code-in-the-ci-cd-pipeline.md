# CI/CD 파이프라인에서  사용

애플리케이션을 안전하게 유지하기 위해 코드에 대한 취약점을 테스트하고 변경 사항이 새로운 취약점을 도입하지 않도록하는 것은 CI/CD 통합을 사용할 수 있습니다.

* Snyk Code는 Snyk Jenkins 플러그인과 같은 Snyk CI/CD 플러그인에서 지원되지 않습니다. CI 서버에 [Snyk CLI를 통합](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)할 수 있습니다.&#x20;
* 취약점 심각성에 따라 결과를 필터링할 수 있으며, 예를 들어 높은 심각성의 취약점이 도입될 때에만 작업을 실패하도록 설정할 수 있습니다. [심각성별 결과 필터링](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#filter-results-by-severity-level)을 참조하세요.
* CLI 출력을 JSON 또는 SARIF 표준 형식으로 내보낼 수 있습니다. [테스트 결과 내보내기](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)를 참조하세요.
* Snyk-to-HTML 도구를 사용하여 더 시각적인 결과를 생성할 수 있습니다. [CLI 도구 snyk-to-html](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md)을 참조하세요.
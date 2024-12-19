# 기술 규격 및 안내

## Snyk Open Source

### Snyk Open Source 및 라이선싱 작동 방식

{% hint style="info" %}
취약점을 테스트하기 전에 오픈 소스 프로젝트를 빌드해야 하는 경우가 있습니다. 자세한 내용은 [Snyk CLI를 사용하여 테스트하기 전에 빌드가 필요한 오픈 소스 프로젝트](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/open-source-projects-that-must-be-built-before-testing-with-the-snyk-cli.md)를 참조하십시오.
{% endhint %}

Snyk는 종속성 그래프 및 (의존성 트리)를 구축하고 그 트리 내의 어느 패키지에서도 취약점을 찾기 위해 [취약점 데이터베이스](https://snyk.io/vuln)를 사용합니다.

Snyk는 프로젝트의 언어, 패키지 관리자 및 프로젝트의 위치에 따라 종속성 트리를 분석하고 구축합니다.

{% hint style="info" %}
공식 릴리스만 추적됩니다. 커밋(기본 브랜치로의 포함 여부를 포함)은 공식 릴리스나 태그에 포함되지 않는 한 식별되지 않습니다.&#x20;

패키지 관리자가 있는 프로젝트의 경우 패키지 관리자의 릴리스를 의미합니다.&#x20;

Go 및 Unmanaged 스캔(C/C++)의 경우 GitHub 리포지토리에서 공식 릴리스나 태그가 필요합니다.
{% endhint %}

### 오픈 소스에서의 Snyk 정책

개발자의 워크플로우를 통해 의존성 및 취약점을 관리하는 정보는 다음을 참조하십시오:

- [보안 오픈 소스 정책 정의](https://snyk.io/series/open-source-security/open-source-policy/)
- [보다 효율적인 수정을 위한 Snyk 보안 정책 사용](https://snyk.io/blog/snyk-security-policies/)

### 오픈 소스 라이선스 준수

오픈 소스 라이센스 준수를 확인하려면 [Snyk 라이센스 준수 관리](../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/snyk-license-compliance-management.md)를 참조하십시오.

## Snyk Code

### Snyk Code에서 지원되는 파일 확장

<table><thead><tr><th>언어</th><th width="215">Interfile 지원</th><th>지원되는 확장자</th></tr></thead><tbody><tr><td>Apex</td><td>예</td><td>.cls, .trigger, .tgr</td></tr><tr><td>C/C++</td><td>예</td><td>.c, .cc, .cpp, .cxx, .h, .hpp, .hxx</td></tr><tr><td>CSharp</td><td>예</td><td>.aspx, .cs</td></tr><tr><td>Go</td><td>예</td><td>.go</td></tr><tr><td>Java</td><td>예</td><td>.java, .jsp, jspx</td></tr><tr><td>JavaScript/TypeScript</td><td>예</td><td>.ejs, .es, .es6, .htm, .html, .js, .jsx, .ts, .cts, .mts, .tsx, .vue, .mjs, .cjs</td></tr><tr><td>Kotlin</td><td>예</td><td>.kt</td></tr><tr><td>PHP</td><td>예</td><td>.php, .phtml, .module, .inc, .install, .theme, .profile</td></tr><tr><td>Python</td><td>예</td><td>.py</td></tr><tr><td>Ruby</td><td>아니요</td><td>.erb, .haml, .rb, .rhtml, .slim</td></tr><tr><td>Scala</td><td>예</td><td>.scala</td></tr><tr><td>Swift</td><td>예</td><td>.swift</td></tr><tr><td>Visual Basic</td><td>예</td><td>.vb</td></tr></tbody></table>

### Snyk Code 분석의 파일 크기 제한

Snyk Code는 다음 파일을 자동으로 분석에서 제외합니다:

- 웹 UI에서: 1MB보다 큰 파일.
- CLI 및 IDE에서: 1MB보다 큰 파일.
- 3줄 이하의 minified JS 파일.

### 파일 이름 길이 제한

분석은 이름이 255자 이하인 파일에만 적용됩니다. 파일 이름이 이 제한을 초과하면 오류가 발생합니다. 모든 파일이 분석되도록 하려면 Snyk는 파일 이름을 짧게 유지하는 것이 좋습니다.

### Unicode 문자 인코딩

Snyk Code는 UTF-8 인코딩으로 된 소스 코드 파일만 허용합니다. Snyk로 가져오기 전에 소스 파일을 이 인코딩 유형으로 변환하는 것을 고려하십시오.

### 프레임워크 지원

특정 프레임워크를 지원하려면, Snyk Code는 해당 언어를 지원하고 해당 프레임워크를 사용하는 프로젝트가 교육되어야 합니다. 발견된 패턴은 보안팀에 의해 주석 처리되고 풍부한 내용으로 확장됩니다.

대부분의 프레임워크는 Snyk Code가 코드를 분석할 수만 있으면 자동으로 지원됩니다. 경우에 따라 특정 규칙이 필요하거나 특정 프로그램 분석 엔진 업데이트가 필요할 수 있습니다.

특정 프레임워크에 대한 지원이 불충분한 경우, [Snyk 지원팀에 문의](https://support.snyk.io)하십시오.

### 프레임워크 지원 수준

Snyk는 프레임워크 지원을 포괄적 및 부분 두 가지 수준으로 분류합니다.

프레임워크를 포괄적으로 지원하는 경우 다음 사항을 의미합니다:

- 소스 및 싱크: Snyk는 모든 관련 소스 및 싱크를 철저히 식별하고 포함했습니다.
- 데이터 흐름 테스트: 포괄적인 데이터 흐름 범위를 보장하기 위해 철저한 테스트가 수행되었습니다.
- 엔진 지원: Snyk Code 엔진은 이 프레임워크에 완전히 최적화되어 있습니다.
- 제한 사항: Snyk는 제한 사항을 인식하지 못합니다. 잘못된 부정을 발견하면 [Snyk 지원](https://support.snyk.io)에 보고하십시오.

부분적인 지원은 다음을 의미합니다:

- 소스 및 싱크: Snyk의 커버리지가 제한되어 일부 소스, 세탁기 또는 싱크가 누락될 수 있습니다.
- 데이터 흐름 테스트: Snyk가 어느 정도의 테스트를 수행했습니다.
- 엔진 지원: 엔진이 이 프레임워크와의 호환성이 제한되어 있어 분석 정확도에 영향을 줄 수 있습니다.
- 제한 사항: Taint 분석 또는 소스 및 싱크 식별에서 잘못된 부정이 발생할 수 있습니다.

프레임워크의 부분적인 지원은 일반적으로 이러한 요소 조합을 포함합니다. 즉, 일부 소스 또는 싱크가 누락될 수 있으며, 엔진은 더 나은 지원을 제공할 수 있지만 데이터 흐름 테스트를 수행하여 분석이 완전히 신뢰할 수 있는지 확인하지 않은 경우가 있습니다.

Snyk는 프레임워크 커버리지를 지속적으로 확대하고 분석 정확도를 향상시킵니다.

### Snyk Code 분석 작동 방식

Snyk는 다음 순서로 코드베이스를 스캔합니다:

1. 소스 코드를 분석하여 이벤트 그래프를 생성합니다. 이벤트 그래프는 서로 다른 코드 부분이 어떻게 관련되어 있는지를 이해하는 데 도움이 되는 코드 맵과 유사합니다. 그래프의 각 노드는 코드에서 발생하는 이슈를 나타냅니다. 일부는 코드의 일부를 나타내고, 다른 것은 코드가 사용되는 방법을 나타냅니다.
2. 규칙을 이벤트 그래프에 적용하여 일치를 찾습니다. 규칙은 Snyk가 이벤트 그래프에서 찾는 알려진 취약점의 체크리스트 역할을 합니다.
3. 일치하는 항목을 찾으면, Snyk는 이벤트 그래프에서 취약점을 찾아 코드에 문제가 숨어 있는 곳을 확인합니다.

자세한 정보는 [Snyk Code AI Engine](../scan-with-snyk/snyk-code/#ai-engine)를 참조하십시오. Snyk Code 언어 지원에 대한 자세한 내용은 [지원되는 언어, 패키지 관리자 및 프레임워크(Snyk Code)](./#code-analysis-snyk-code)를 참조하십시오.

## 언어 지원 및 CLI, CI/CD 및 SCM 통합

### Snyk Code용 CLI

CLI를 통해 Snyk Code를 사용하여 코드를 테스트하려면 터미널에서 리포지토리를 열고 `snyk code test`를 실행하십시오.

테스트 옵션을 사용자 정의하고 다른 명령을 실행하거나 디렉토리 및 파일을 제외하며 결과를 다양한 형식으로 보고 탐색하는 방법에 대한 정보는 다음을 참조하십시오:

- [사용 가능한 명령](../snyk-cli/commands/#available-commands)
- [Snyk Code CLI 테스트에서 디렉터리 및 파일 제외](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)
- [테스트 결과 출력](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#output-test-results)
- [테스트 결과 내보내기](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)
- [snyk-to-html](../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md)

`snyk code test`를 실행한 후에는 다음을 수행할 수 있습니다:

- [수정 PR 열기](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/)
- [PR Checks 구성](../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md)

### Snyk Open Source용 CLI

취약점을 테스트하려면 해당 패키지 관리자를 설치한 후 Snyk에서 지원되는 해당 Manifest 파일을 포함해야 합니다.

오픈 소스 프로젝트를 취약점으로 테스트하려면 `snyk test` 명령을 실행하십시오.

### SCM 통합 사용 방법

- [통합 설정](../getting-started/#set-up-a-snyk-integration).
- 자세한 내용은 [Snyk SCM 통합](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 참조하십시오.
- 언어별 정보는 [Maven 및 Gradle을 사용하는 Git 리포지토리](java-and-kotlin/git-repositories-with-maven-and-gradle.md), [Git 리포지토리 및 JavaScript](javascript/git-repositories-and-javascript.md) 및 [Git 리포지토리 및 Python](python/git-repositories-and-python.md)를 참조하십시오.
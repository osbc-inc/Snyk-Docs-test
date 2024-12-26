# Snyk Code

## 개요

Snyk Code는 개발자가 문제를 해결하고 안전한 소프트웨어를 개발하는 데 도움이 되도록 거짓 양성 결과를 줄이고 빠르고 정확한 보안 도구입니다.

다음 옵션을 사용하여 코드를 스캔할 수 있습니다:

* [Snyk Web UI](../../getting-started/snyk-web-ui.md) ([PR checks 포함](../pull-requests/pull-request-checks/))
* [Snyk IDE](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)
* [Snyk CLI](../../snyk-cli/)
* [Snyk API](../../snyk-api/)

다음 표는 {Snyk Code의 기능을 보여줍니다. 코드 분석, 코드 내의 보안 문제 관리 및 개발 환경에서 문제를 해결하는 것을 용이하게 하는 기능을 포함하고 있습니다.

<table><thead><tr><th width="179">기능</th><th>설명</th></tr></thead><tbody><tr><td><strong>문제 필터링, 정렬 및 그룹화</strong></td><td><p>가장 일반적인 문제를 식별하려면 심각성, 프로그래밍 언어, 우선 순위 점수 등에 따라 문제를 필터링할 수 있습니다.</p><p>자세한 내용은 <a href="manage-code-vulnerabilities/#filtering-existing-projects">기존 프로젝트 필터링</a>을 확인하십시오.</p></td></tr><tr><td><strong>우선 순위 점수</strong></td><td><p>문제의 보편성, 수정 용이도 및 위험 요소를 포함하여보다 중요한 문제를 정렬하고 우선 순위를 정합니다.</p><p>자세한 내용은 <a href="../../manage-risk/prioritize-issues-for-fixing/priority-score.md">우선 순위 점수</a>를 확인하십시오.</p></td></tr><tr><td><strong>데이터 흐름</strong></td><td><p>문제의 소스에서 싱크로의 경로를 단계별 흐름으로 시각화합니다.</p><p>자세한 내용은 <a href="manage-code-vulnerabilities/breakdown-of-code-analysis.md">데이터 흐름</a>를 확인하십시오.</p></td></tr><tr><td><strong>취약점</strong></td><td><p>취약점에 대해 상세한 콘텐츠를 통해 취약점이 어떻게 생성되었는지, 위험 요소는 무엇이며, 그에 대한 인기 있는 완화 전략에 대해 자세히 알아봅니다.</p><p>자세한 내용은 <a href="manage-code-vulnerabilities/">코드 취약점 관리</a>를 확인하십시오.</p></td></tr><tr><td><strong>수정 분석</strong></td><td><p>동일한 데이터 흐름에서 비슷한 문제를 수정하는 실제 코드에 대한 링크가 포함된 예시를 확인하여 통찰과 맥락을 얻습니다.</p><p>자세한 내용은 <a href="manage-code-vulnerabilities/breakdown-of-code-analysis.md">코드 분석 분해</a>를 확인하십시오.</p></td></tr><tr><td><strong>Jira 이슈 생성</strong></td><td><p>Snyk 문제를 Jira 프로젝트로 추적하고 내보냅니다.</p><p>자세한 내용은 <a href="../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md#create-a-jira-issue">Jira 이슈 생성</a>을 확인하십시오.</p></td></tr><tr><td><strong>문제 무시</strong></td><td><p>문제의 건의된 수정을 무시하도록 Snyk를 구성하여 특정 경고를 억제할 수 있습니다. 예를 들어 테스트 코드에서 루틴을 테스트하기 위해 하드 코딩된 암호를 사용하거나 문제를 인식했지만 수정하지 않기로 결정한 경우 등입니다.</p><p>자세한 내용은 <a href="../../manage-risk/prioritize-issues-for-fixing/ignore-issues/">문제 무시</a>를 확인하십시오.</p></td></tr><tr><td><strong>가져오기 프로세스에서 파일 제외</strong></td><td><p><code>DeepCode/Snyk</code> ignore 파일 <code>.gitignore</code> 등이 있는지 확인하고, 그러한 파일이 있으면 해당 파일을 읽습니다. 이러한 파일에 정보를 사용하여, Snyk는 프로젝트 디렉토리에서 지원되는 확장자를 가진 파일만 식별합니다. 는 이런 크기가 4MB보다 작은 파일을 번들로 묶어서 Snyk에 보냅니다. <code>,gitignore</code> 예외는 <code>snyk code test</code> CLI 명령에서 존중됩니다.</p><p>또한 <a href="../import-project-repository/exclude-directories-and-files-from-project-import.md">가져오기 프로세스에서 폴더 및 파일 제외</a>를 참고하십시오.</p></td></tr><tr><td><strong>인터파일 분석</strong></td><td>이는 루비를 제외한 가 지원하는 [모든 언어](../../supported-languages-package-managers-and-frameworks/#code-어...
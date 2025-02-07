# Snyk 오픈소스로 취약점 수정문제 해결

취약점을 발견하면 해당 취약점을 Snyk에 보고할 수 있는 기회가 있습니다. 자세한 내용은 [오픈 소스 패키지의 취약점 공개](../../../working-with-snyk/disclosure-of-a-vulnerability-in-an-open-source-package.md)를 참조하십시오.

## Snyk에 의해 발견된 문제에 대해 풀 리퀘스트 또는 머지 리퀘스트를 열 수 없음

프로젝트를 가져올 때, 통합 또는 CLI를 통해 가져올 때, Snyk은 직접 종속성에 대해서만 Fix PR 버튼을 제공합니다. Snyk은 간접 종속성에 대해서는 PR을 열지 않습니다. 자세한 내용은 [간접 종속성 해결](vulnerability-fix-types.md#fixing-transitive-dependencies)을 참조하십시오.

## Fix Pull Requests 또는 Merge Requests에 대한 지원하는 언어

Snyk은 취약점을 해결할 패치 또는 업그레이드를 가진 종속성에 대해 Fix Pull Requests (Fix PRs) 또는 Merge Requests (MRs)를 생성할 수 있습니다.

다음 언어에 대해 Snyk은 Fix PRs 또는 MRs를 생성하는 것을 지원합니다:

* [Maven](../../../supported-languages-package-managers-and-frameworks/java-and-kotlin/guidance-for-java-and-kotlin.md#maven)
* [.NET](../../../supported-languages-package-managers-and-frameworks/.net/)
* [npm](../../../supported-languages-package-managers-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#npm)
* [Python](../../../supported-languages-package-managers-and-frameworks/python/)
* [Ruby](../../../supported-languages-package-managers-and-frameworks/ruby/)
* [Yarn](../../../supported-languages-package-managers-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#yarn)

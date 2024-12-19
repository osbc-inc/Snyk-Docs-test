# 풀 리퀘스트

## Snyk Fix PRs

Snyk에서 새로운 이슈가 프로젝트 테스트에서 식별되었을 때 또는 식별된 취약점이 있는 프로젝트로 재테스트가 실행될 때 자동으로 생성되는 고치기 풀 리퀘스트 또는 머지 리퀘스트입니다. 이 기능은 GitHub Enterprise나 Azure와 같은 SCM 통합을 통해 가져온 프로젝트에 적용됩니다.

고치기 풀 리퀘스트 및 업그레이드 풀 리퀘스트를 통해 오픈 소스 의존성을 업그레이드하는 방법에 대한 자세한 정보는 [자동 PR을 사용하여 오픈 소스 의존성 업그레이드](snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/upgrade-open-source-dependencies-with-automatic-prs.md)를 참조하십시오.

GitHub 계정에서 풀 리퀘스트를 열기 위한 지침은 [고정된 GitHub 계정에서 고치기 및 업그레이드 풀 리퀘스트 열기](snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md)를 참조하십시오.

Snyk Fix PRs에 대한 자세한 설명은 [Snyk Fix Pull Requests](snyk-pull-or-merge-requests/)를 참조하십시오.

## PR Checks

풀 리퀘스트 체크는 프로젝트에서 새로운 이슈를 식별하기 위해 생성된 풀 리퀘스트에서 실행되는 테스트입니다. 이를 통해 주요 브랜치로 머지하기 전에 코드에 도입되는 문제를 방지할 수 있습니다.

PR 체크에 대한 자세한 설명은 [풀 리퀘스트 체크](pull-request-checks/)를 참조하십시오.

## Fix PRs와 PR Checks의 차이점

* Snyk Fix 풀 또는 머지 리퀘스트는 Project 테스트 또는 재테스트에서 새로운 이슈가 식별되었을 때 Snyk에 의해 시작됩니다. PR Checks 또는 머지 리퀘스트는 조직 또는 프로젝트에서 코드 변경이 있을 때 Snyk에 의해 시작됩니다.
* Snyk Fix PRs는 문제의 **해결**을 위해 사용됩니다. PR 체크는 문제의 **방지**를 위해 사용됩니다.
* Snyk Fix PRs는 사용 가능한 고치기, 예를 들어 업그레이드 또는 패치,으로 문제를 해결하기 위해 자동 또는 수동으로 요청됩니다. PR 체크는 설정 중에 구성한 조건에 기반한 코드 변경에서 수행되는 테스트로, 머지 전에 코드에 영향을 줄 수 있는 취약성을 식별합니다.
* Snyk Fix PRs는 일일 스캔이나 테스트에서 이슈가 감지될 때 트리거됩니다. PR 체크는 PR에서 코드 변경이 있을 때 트리거됩니다.
# Git 리포지토리 통합 소개

자체 코드에는 [Snyk Code](https://docs.snyk.io/scan-with-snyk/snyk-code)를 사용하고 사용하는 오픈 소스 라이브러리에는 [Snyk Open Source](https://docs.snyk.io/scan-with-snyk/snyk-open-source)를 사용하여 애플리케이션 코드를 보호하는 데 Snyk 함수를 사용할 수 있습니다.

다음 섹션에서 설명하는 대로 이러한 기능은 개발 프로세스의 각 단계에서 사용할 수 있습니다.

{% hint style="info" %}
Git 통합에 대한 추가 소개 정보는 설명서의 이 섹션의 다른 문서 및 [Snyk 조직 간 통합 복제](../../../enterprise-setup/snyk-broker/clone-an-integration-across-your-snyk-organizations.md)를 참조하세요.
{% endhint %}

## 개발 중

개발자는 중앙 Git 리포지토리로 코드 변경 사항을 푸시하기 전에 로컬에서 코드를 작성하는 동안 Snyk을 사용하여 문제를 확인할 수 있습니다.

개발자는 다음을 사용하여 테스트, 수정 및 모니터링할 수 있습니다:

* 개발자의 로컬 머신에서 로컬 확인을 위한 [Snyk CLI](../../../snyk-cli/)\
  **학습**: [Snyk CLI 소개](https://learn.snyk.io/lesson/snyk-cli/) 참조
* 개발자 IDE에 포함된 확인을 위한 [Snyk for IDEs](../../snyk-ide-plugins-and-extensions/) 플러그인\
  **학습**: [IDE에서 Snyk 사용 소개](https://learn.snyk.io/lesson/snyk-in-an-ide/) 참조

## 코드 병합 시

개발자가 코드 변경 사항을 Git 리포지토리에 병합하면 Snyk이 다음을 할 수 있습니다:

* **PR Checks 실행**: 기본적으로 PR Checks는 어택 서피스가 증가하지 않도록 하며, 문제가 있는 종속성이 추가될 때에만 실패합니다.
* **매일 스캔 실행**: 기본적으로 Snyk은 리포지토리에서 Snyk 프로젝트를 가져온 경우에 매일 스캔을 실행하여 현재 라이브러리에서 새로운 문제를 빠르게 찾을 수 있습니다. 이 스캔은 현재 작업 중인지 여부에 관계없이 모든 가져온 프로젝트에 대해 수행됩니다. [코드 리포지토리 프로젝트 안내](../../../implement-snyk/walkthrough-code-repository-projects/) 참조
* **재검사 트리거**: SCM 통합을 위한 웹훅이 생성된 경우, PR이 병합될 때마다 Snyk은 재검사를 트리거합니다.
* **Jira 티켓 생성**: 발견된 새로운 문제에 대한 작업을 관리하고, 해당 작업을 팀 내 개발자에게 할당하고 이러한 문제의 진행 상황을 추적할 수 있습니다. [Jira 통합](../../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md) 문서 참조

## 자동으로 수정

Snyk은 취약점을 처리하기 위해 PR을 만들어 새로운 취약점을 해결하고, 이전 종속성을 해결하며, 시간이 지남에 따라 백로그된 취약점을 처리할 수도 있습니다. [첫 번째 취약점 수정](../../../implement-snyk/walkthrough-code-repository-projects/fix-your-first-vulnerability.md) 참조

## 빌드 중

Snyk은 애플리케이션을 빌드할 때 다시 "게이트"로 작용하여 빌드 프로세스의 일부로 문제를 자동으로 확인하여 코드가 안전한지 확인할 수 있습니다. 정책에 따라 빌드가 완료되지 않도록 하여 구현입니다.

문제에 대한 보고를 선택하여 빌드를 진행하거나 문제가 발생하면 빌드를 중단할 수 있습니다. 또한 쉽게 취약점 수정 가능 여부, CVSS 점수 및 기존 수정 사항 등의 기준을 이 프로세스에 추가할 수 있어 특정 문제 해결에 집중할 수 있습니다.

Snyk은 이 프로세스를 지원하기 위한 [Snyk 도구들](../../../scan-with-snyk/snyk-tools/)을 제공합니다.

## 보안 게이트 및 배포

이러한 과정 및 보안 게이트를 통과한 후, 응용 프로그램 및 코드는 전통적 및 PaaS, Serverless, 레지스트리에 배포할 수 있습니다. 새로운 취약점이 발견될 때 경보를받거나 Jira 티켓을 만들고, Snyk의 모니터 기능 및 기타 기능을 사용하여 보안을 유지할 수 있습니다.

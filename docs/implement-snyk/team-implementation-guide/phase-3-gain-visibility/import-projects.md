# 프로젝트 가져오기

구성된 통합에 따라서 기술 스택의 언어/패키지 관리자에 따라, Snyk을 사용하여 다음을 통해 프로젝트를 가져올 수 있습니다:

* Git 리포지토리와의 소스 제어 통합
* CI/CD를 사용하는 Snyk CLI

가장 적합한 가져오기 경로는 기술 스택의 언어와 패키지 관리자에 따라 다릅니다.

가장 좋은 시작점을 결정하기 위한 몇 가지 주요 포인트는 다음과 같습니다. 자세한 내용은 [Git 리포지토리 및 CI/CD 비교](../../../scm-ide-and-ci-cd-integrations/git-repository-and-ci-cd-integrations-comparisons.md)를 참조하십시오.

## Snyk으로 시작하기

{% hint style="info" %}
자세한 내용은 [시작하기](../../../getting-started/) 및 [스캔 시작](../../../scan-with-snyk/start-scanning.md)을 참조하십시오.
{% endhint %}

귀하는 필요에 따라 Snyk이 다양한 통합 방법을 제공합니다.

### Git 통합

자세한 내용은 [Snyk을 사용한 Git 리포지토리(SCMs) 통합](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 참조하십시오.

자동 스캔을 위해 리포지토리를 연결하십시오.

수백 개 또는 수천 개의 리포지토리의 경우:

* 규모에 맞게, Snyk은 API 사용을 권장합니다. API는 Snyk 엔터프라이즈 플랜에서 사용할 수 있습니다.
  * [Snyk API](../../../snyk-api/)를 사용하여 프로젝트를 가져옵니다. 기존 소스 제어 통합을 활용하고 프로세스 자동화에 사용할 수 있습니다.
  * [snyk-api-import](../../../scan-with-snyk/snyk-tools/tool-snyk-api-import/) 도구는 대규모 기업의 스케일에서 온보딩 관리를 위해 API를 사용하고 대규모로 사용하는 것을 권장하는 도구입니다. 소스 제어 구조를 반영해야 합니다.

## Snyk CLI

자세한 내용은 [Snyk CLI](../../../snyk-cli/)를 참조하십시오.

CLI를 사용하여 개별 프로젝트를 세밀하게 스캔할 수 있습니다.

{% hint style="info" %}
각 유형의 테스트에 대해 명령어를 작성해야 합니다 (오픈 소스, 코드, 인프라스트럭처 코드, 컨테이너).
{% endhint %}

Snyk CLI를 사용하려면:

1. 빌드 스크립트의 일부로 적절한 방법 중 하나를 사용하여 [CLI 설치](https://docs.snyk.io/snyk-cli/install-or-update-the-snyk-cli)를 수행하십시오.
2. 스크립트에서 프로젝트 폴더로 이동하십시오.
3. 원하는 스캔 유형에 대한 적절한 `snyk test` 또는 `snyk monitor` 명령어 및 옵션을 실행하십시오.

   테스트를 어디에 구현할지는 일반적으로 유연합니다. 보통 배포 전에 가장 자주 수행됩니다.

   일반적으로, Snyk 오픈 소스는 의존성이 빌드 시스템에 설치된 후에 test 또는 monitor로 실행됩니다.

4. 결과를 검토하려면, `snyk test`를 실행할 때 로컬로 또는 monitor 또는 report를 사용하여 Snyk 웹 UI에서 확인하십시오.

다양한 파이프라인 통합의 데모를 보려면 [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples)를 참조하십시오.
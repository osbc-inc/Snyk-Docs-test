# 프로젝트 가져오기

구성된 통합에 따라서 기술 스택의 언어/패키지 관리자에 따라, Snyk를 사용하여 다음을 통해 프로젝트를 가져올 수 있습니다:

* Git 리포지토리와의 소스 제어 통합
* CI/CD를 사용하는 Snyk CLI

가장 적합한 가져오기 경로는 기술 스택의 언어와 패키지 관리자에 따라 다릅니다.

가장 좋은 시작점을 결정하기 위한 몇 가지 주요 포인트는 다음과 같습니다. 자세한 내용은 [Git 리포지토리 및 CI/CD 비교](../../../scm-ide-and-ci-cd-integrations/git-repository-and-ci-cd-integrations-comparisons.md)를 참조하십시오.

## Snyk으로 시작하기

{% hint style="info" %}
자세한 내용은 [시작하기](../../../getting-started/) 및 [스캔 시작](../../../scan-with-snyk/start-scanning.md)을 참조하십시오.
{% endhint %}

귀하는 필요에 따라 Snyk가 다양한 통합 방법을 제공합니다.

### Git 통합

자세한 내용은 [Snyk를 사용한 Git 리포지토리(SCMs) 통합](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 참조하십시오.

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

1. 빌드 스크립트의 일부로 적합한 방법을 사용하여 [CLI 설치](https://docs.snyk.io/snyk-cli/install-or-update-the-snyk-cli)를 진행합니다.
2. 스크립트에서 프로젝트 폴더로 이동합니다.
3. 실행하려는 스캔 유형에 맞는 적절한 `snyk test` 또는 `snyk monitor` 명령어와 옵션을 실행합니다.\
   \
   스크립트에서 테스트를 구현할 위치는 일반적으로 유연하지만, 대부분 배포 전에 실행됩니다. Snyk Open Source 및 Snyk Container는 모니터 명령어만 사용하여 수동으로 보고합니다. 테스트 명령어를 사용할 경우, 목적은 `--severity-threshold`와 같은 특정 기준을 충족하는 문제가 발견되면 빌드를 중단하는 것입니다. CLI나 snyk-filter 플러그인에서 제공하는 여러 옵션을 사용할 수 있습니다.\
   \
   일반적으로 Snyk Open Source는 종속성이 빌드 시스템에 설치된 후 `test` 및/또는 `monitor` 명령어로 실행됩니다.\
   \
   일반적인 명령어는 다음과 같습니다:
   * 코드: `snyk code test --org=[org-id]`
   * 오픈 소스:
     * `snyk test --all-projects --org=[org-id]`
     * `snyk monitor --all-projects --org=[org-id]`\
       `[org-id]`를 조직의 ID로 교체하세요.
   * 컨테이너 및 코드 인프라 스캔에 대해서는 [컨테이너](../../../scan-with-snyk/snyk-container/scan-container-images.md)와 [인프라 코드](../../../scan-with-snyk/snyk-iac/)를 참조하세요. 이는 스캔하는 유형에 따라 달라집니다.
4. `snyk test` 명령어를 실행할 때 로컬에서 결과를 확인하거나, `monitor` 또는 `report`를 사용할 때 Snyk 웹 UI를 통해 결과를 확인합니다.

다양한 파이프라인 통합에 대한 데모는 [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples)를 참조하세요.

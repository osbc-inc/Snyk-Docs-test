# 프로젝트 가져오기

기존에 구성한 통합 및 기술 스택에서의 언어 및 패키지 관리자에 따라, 다음을 사용하여 Snyk로 프로젝트를 가져올 수 있습니다:

- Git 리포지토리와의 소스 제어 통합
- CI/CD를 통한 Snyk CLI

최적의 가져오기 경로는 기술 스택의 언어 및 패키지 관리자에 따라 다릅니다.

가장 좋은 시작점을 결정하는 데 유용한 몇 가지 핵심 포인트가 있습니다. 더 많은 세부 정보는 [Git 리포지토리 및 CI/CD 비교](../../../scm-ide-and-ci-cd-integrations/git-repository-and-ci-cd-integrations-comparisons.md)를 참조하세요.

## Snyk로 시작하는 방법

{% hint style="info" %}
자세한 내용은 [시작하기](../../../getting-started/) 섹션 및 [CLI, Web UI 또는 AP를 사용한 스캔 시작](../../../scan-with-snyk/start-scanning.md)을 참조하세요.
{% endhint %}

Snyk는 여기에서 설명하는대로 여러 통합 방법을 제공합니다.

### Git 통합

자동 스캔을 위해 리포지토리를 연결할 수 있습니다. 자세한 내용은 [Snyk SCM 통합](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 참조하세요.

애플리케이션이 수백개나 천 개 미만인 경우에는 다음 단계를 따르세요.

1. **설정**으로 이동한 다음 **통합**으로 이동하여 통합 페이지의 타일을 사용하여 Git 코드 리포지토리에 연결합니다.
2. 통합 설정에서:
   - 프로젝트 초기 설정 시 자동 수정 및 PR/병합 확인을 비활성화합니다.
   - 안정된 상태에 도달하고 차단이 필요한 경우에 활성화합니다.
3. 프로젝트 목록에서 웹 UI를 사용하여 프로젝트를 추가합니다.
4. Git 코드 리포지토리에서 결과를 모니터링합니다.

수백 개나 수천 개의 리포지토리에 대해, [프로젝트 가져오기](../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import) API 엔드포인트를 사용하여 프로젝트를 가져올 수 있습니다. 이는 기존 소스 제어 통합을 활용하고 프로세스를 자동화하는 데 사용할 수 있습니다.

[snyk-api-import](../../../scan-with-snyk/snyk-tools/tool-snyk-api-import/) 도구는 대규모 기업에서 대규모로 온보딩을 관리하기 위해 API를 사용하며 대규모에서 사용할 것을 권장하는 도구입니다. snyk-api-import 도구를 사용할 때 소스 제어 구조를 반영해야 합니다.

### Snyk CLI

[Snyk CLI](../../../snyk-cli/)는 개별 프로젝트를 자세히 스캔할 수 있습니다.

{% hint style="info" %}
수행할 각 테스트 유형에 대한 명령을 구성해야 합니다: 오픈 소스, 코드, 코드 인프라 또는 컨테이너 테스트.
{% endhint %}

CLI를 사용하는 방법은 다음과 같습니다:

1. 빌드 스크립트의 일부로 적절한 방법 중 하나를 사용하여 [CLI 설치](../../../snyk-cli/install-or-update-the-snyk-cli/)합니다.
2. 스크립트에서 프로젝트 폴더로 이동합니다.
3. 실행할 적절한 `snyk test` 또는 `snyk monitor` 명령을 사용하여 해당 스캔 유형에 대한 옵션과 함께 명령을 실행합니다.\
   테스트를 구현해야 하는 위치는 일반적으로 유연하나, 대부분의 경우 배포 전에 테스트가 수행됩니다. `monitor` 명령만 사용하여 {{Snyk Open Source}} 및 {{Snyk Container}}에서 취약점을 수동으로 보고합니다. `test` 명령을 사용하는 경우, 빌드가 중단되도록 설정하는 것이 목적이며, 이는 `--severity-threshold`와 같은 특정 기준을 충족하는 문제를 발견할 경우 또는 CLI나 snyk-filter 플러그인의 여러 옵션 중 하나를 사용하는 경우입니다.\
   \
   일반적으로 `test` 또는 `monitor` 명령 또는 둘 다가 대부분 opensource에 대해 종속성이 빌드 시스템에 설치된 후에 실행됩니다.\
   \
   일반적인 명령은 다음과 같을 수 있습니다:
   - 코드: `snyk code test --org=[org-id]`
   - 오픈 소스:&#x20;
     - `snyk test --all-projects --org=[org-id]`
     - `snyk monitor --all-projects --org=[org-id]`
   - [컨테이너](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/) 및 [인프라 구성 코드](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/) 스캔 관련 문서를 참조하여 해당 유형의 프로젝트를 스캔하는 방법을 확인하세요.
4. `snyk test`를 실행할 때 로컬로 결과를 확인하거나, monitor 또는 보고 기능을 사용하여 Web UI에서 결과를 확인합니다.

{% hint style="info" %}
다양한 파이프라인 통합의 시연은 [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples)에서 찾을 수 있습니다.
{% endhint %}

## Snyk API

[Snyk API](../../../snyk-api/)를 사용하여 스캔할 수 있습니다. 이를 통해 대규모 자동 스캔이 가능합니다.

프로세스는 다음과 같습니다:

1. **설정**으로 이동하고 **서비스 계정**을 선택한 다음 API 토큰을 생성합니다.
2. 파이프라인에서 Snyk API를 호출합니다.
3. 결과를 프로그래밍적으로 처리합니다.
4. 필요한 경우 코드를 변경합니다.

API를 사용한 스캔은 다음을 달성하는 데 유용합니다:

- 파이프라인을 통한 스캔 트리거.
- 프로젝트 포트폴리오 전체에 걸쳐 확장.
- 실시간으로 새로운 문제 식별.
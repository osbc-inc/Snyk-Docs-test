# 테스트로 간주되는 것은 무엇인가요?

{% hint style="info" %}
본 문서에서 정의되지 않은 대문자화된 용어는 [snyk.docs.io](http://snyk.docs.io/)에서 찾을 수 있는 고객의 구매 계약 또는 다른 적용 가능한 문서에서 정의된 의미를 갖습니다.
{% endhint %}

Snyk은 Snyk 제품 각각에 대해 별도의 테스트 횟수를 유지하고 제한을 설정합니다: Snyk Open Source, Snyk Code, Snyk Container, 및 Snyk IaC.

무료 Snyk 요금제를 사용하는 경우 공개 저장소에서는 무제한 테스트를 실행할 수 있으며 비공개 저장소에서는 제한된 테스트를 실행할 수 있습니다. 반복 테스트는 주간 기준으로만 실행할 수 있습니다. 자세한 내용은 [요금제 및 가격정보](https://snyk.io/plans?_gl=1*1d4rb1a*_ga*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.*_ga_X9SH3KP7B4*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w)를 참조하십시오. 공개 저장소에 대한 무제한 테스트에 대한 정보는 [테스트가 부족할 때](../snyk-cli/getting-started-with-the-snyk-cli.md#running-out-of-tests)를 참조하십시오. 한도에 도달하거나 반복 테스트 빈도를 늘리고 싶은 경우 [요금제를 업그레이드](https://snyk.io/plans?_gl=1*1d4rb1a*_ga*NzI0Mjg1NDM2LjE2OTAzNzU5NDk.*_ga_X9SH3KP7B4*MTY5MzYxOTc2NC4xMzIuMS4xNjkzNjE5ODA1LjAuMC4w)할 수 있습니다.

Snyk Open Source, Snyk Code, Snyk IaC, 그리고 Snyk Container 애플리케이션을 통해 고객은 해당 애플리케이션의 기능에 따라 코드 기반 자산에 대해 스캔 및 테스트를 실행할 수 있습니다. 고객의 주문서에는 고객의 구독 할당량(테스트)의 일부로 특정 테스트 수를 함께 제공하는 계획 유형이 표시됩니다.

본 문서는 고객이 구독 할당량에 대비하여 사용량을 이해할 수 있도록 Snyk이 테스트로 간주하는 내용을 개요로 제시합니다. 현재 테스트 한도는 Snyk Open Source 및 Snyk Code 어플리케이션을 중심으로만 집중하고 있으며 본 문서에서는 해당 항목에 대한 테스트 한도에 대해 논의됩니다.

테스트에는 두 가지 주요 유형이 있습니다:

- 반복: 테스트는 Snyk 애플리케이션에 의해 트리거되며 고객의 구성에 기반하고 일정 주기(매일 또는 매주)로 발생합니다. 이러한 테스트는 Web UI, CLI 또는 API에 의해 트리거되며 대개 SCM 내에서 cron 작업을 통해 실행됩니다.
- 수동: 테스트는 특정 선거를 통해 고객에 의해 트리거됩니다. 이러한 테스트는 애플리케이션의 사용 가능한 기능 내에서 어떠한 주기에서도 발생할 수 있습니다. 이러한 테스트는 여러 가지 방법으로 트리거될 수 있습니다. 아래와 같이:
  - API - API 호출에 의해 트리거됨
  - CLI - CLI 명령에 의해 트리거됨
  - IDE - 저장하거나 자동 저장에 의해 트리거됨(IDE별로 상이할 수 있음)
  - Pull requests 테스트 또는 체크 - 새로운 Pull Request 생성에 의해 트리거됨
  - Push 테스트 - 고객의 SCM에 의해 트리거됨
  - Web UI Import 또는 재테스트 - WebUI의 버튼에 의해 트리거됨

각 고객의 특정 사용 및 구성은 다른 고객들과 다를 수 있으므로 Snyk은 여기에서 설명된 기준을 사용하여 어떤 것이 테스트로 간주되는지 결정합니다.

다음은 Snyk이 각 애플리케이션에 대해 몇 개의 테스트로 계산하는지 설명합니다:

- Snyk Open Source: 애플리케이션에서 식별된 취약성이 포함된 manifest 파일 수; 한 리포지토리에는 여러 manifest 파일이 포함될 수 있음에 유의하십시오.
- Snyk Code: 애플리케이션에서 고객이 스캔한 저장소의 수.

다음은 고객이 Snyk 애플리케이션에서 테스트를 시작할 수 있는 예시입니다:

| 용어    | 정의                                                                               | 예시                                                               |
|---------|------------------------------------------------------------------------------------|--------------------------------------------------------------------|
| API     | Snyk과 프로그래밍적으로 통합하기 위해 Application Programming Interface 사용 | API 엔드포인트 사용: `https://api.snyk.io/rest`, `https://api.snyk.io/v1/test` 엔드포인트      |
| PR      | 소스 컨트롤 관리자(SCM) 내에서 Pull Request(PR)를 제출하기                      | Github 내에서 PR을 사용하여 테스트 트리거                     |
| Push    | 고객 SCM에 의해 트리거되는 Push 테스트                                           | 고객이 Github를 SCM으로 사용하고 Jenkins를 CI/CD로 사용할 때 Jenkins에 cron 작업을 생성하여 일정 간격마다 실행하고, Jenkins는 Github로부터 최신 변경사항을 가져와 사전 정의된 스크립트를 실행함 |
| Recurring | Snyk 애플리케이션이 일정 주기(매일 또는 매주)에 따라 실행되는 고객의 구성에 기반한 테스트. Web UI, CLI, 또는 API에 의해 트리거되고 대개 SCM 내의 cron job을 통해 실행됨.  | Snyk Github 통합을 사용하여 고객은 매일 또는 매주 테스트를 설정할 수 있음. |
| Retest  | Snyk 웹 앱 내에서 다시 테스트 버튼 사용                                       | 사용자가 Snyk 웹 앱 내에서 다시 테스트를 클릭합니다.              |
| Snyk 모니터 명령 | CLI를 사용하여 고객의 Snyk 계정에 프로젝트를 지속적으로 모니터하여 오픈 소스 취약점 및 라이선스 문제를 확인하고, 결과를 snyk.io로 보냄. 이는 Snyk Open Source 및 Snyk Container에만 해당됨. | 사용자는 CLI에서 `snyk monitor`를 실행합니다. |
| Snyk 테스트 명령 | <p>CLI를 사용하여 프로젝트에서 오픈 소스 취약점 및 라이선스 문제를 확인. <code>test</code> 명령은 지원되는 manifest 파일을 자동으로 감지하고 해당 내용을 테스트합니다.</p><p><br>참고: , Container 및 IaC 스캔 방법에는 각각 특정 <code>snyk test</code> 명령이 있습니다: <code>snyk code test</code>, <code>snyk container test</code>, 및 <code>snyk iac test</code>.</p> | 사용자가 CLI에서 `snyk test`를 실행합니다. |
| 사용자 | 저장소를 가져올 때 트리거되는 테스트                                      | Snyk 웹 앱에서 사용자가 가져오기 버튼을 클릭합니다.                 |
| IDE    | 통합 개발 환경, VS Code, JetBrains 등                                                |                                                                      |

## Git 리포지토리 통합 스캔 계산

이러한 [Git 리포지토리 (SCM)](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/) 통합을 위한 Snyk 기능은 기본적으로 자동으로 스캔됩니다:

* 매일 반복 테스트
* 의존성이 기본 브랜치에서 변경될 때 실행되는 자동 스캔
* 해당 의존성을 변경하는 풀 리퀘스트를 생성할 때 실행되는 PR 체크

원본 코드 리포지토리에 Dockerfile이 있는 경우 기본 설정에서 감지 및 스캔되지만 Dockerfile은 Snyk Open Source 스캔이 아니라 Snyk Container 스캔으로 계산됩니다.

소스 코드 리포지토리에서 스캔된 Terraform 및 Kubernetes 구성 파일은 Snyk IaC 스캔으로 계산됩니다.

레지스트리나 Kubernetes 클러스터에서 컨테이너 스캔의 경우 Snyk은 초기 스캔 및 후속 반복 스캔을 계산합니다. 기본적으로 반복 스캔은 매일 한 번 실행됩니다.

## 반복 스캔 계산

Snyk은 주기적으로 코드가 새로 발표된 취약점에 영향을 받는지 확인합니다.

테스트 빈도는 각 제품에 대해 기본값으로 설정됩니다. 빈도를 변경하는 방법에 대한 자세한 내용은 [사용 설정](../snyk-admin/groups-and-organizations/usage-settings.md), [프로젝트 설정 보기 및 편집](../snyk-admin/snyk-projects/view-and-edit-project-settings.md), 및 [테스트 빈도 설정](../snyk-admin/snyk-projects/#test-frequency-settings)을 Snyk 프로젝트 페이지에서 확인하십시오.

## CLI 테스트 계산

다음 명령을 실행할 때마다 각 명령에 대해 하나의 테스트가 계산됩니다:

- Snyk Open Source의 경우: `snyk test` 또는 `snyk monitor`.
- Snyk Container의 경우: `snyk container test` 또는 `snyk container monitor`.
- Snyk Code의 경우: `snyk code test`.

Snyk IaC의 경우, 명령은 `snyk iac test`입니다. 이는 여러 프로젝트를 스캔할 수 있으므로 스캔은 스캔 중인 모든 프로젝트에 대해 카운트됩니다. 예를 들어, `snyk iac test` 명령이 11개의 프로젝트를 스캔하는 경우, 카운트가 11로 증가합니다.

## 앱 기반 테스트 계산

새 프로젝트를 추가하거나 재테스트 버튼을 클릭하면 스캔이 실행됩니다. 이는 실행 중인 모든 자동 테스트에 추가됩니다.

## API 테스트 계산

`https://api.snyk.io/v1/test` 엔드포인트로 호출을 할 때 테스트가 계산됩니다.

## 테스트 사용 정책

Snyk은 고객의 테스트 양을 매일 모니터링하고 최근 30일간의 연속 90일간 누적으로 고객의 테스트 사용량이 구매한 요금제로 허용된 한도를 20% 초과하거나 30일 동안 100% 초과하면 고객에게 초과 사용에 대해 통보하고, 조치가 필요하면 구독의 테스트 할당량을 늘릴 수 있도록 확장 청구서를 제공하거나 고객에게 사용량을 줄이도록 요청할 수 있습니다. 예외적인 경우를 제외하고, Snyk은 테스트 초과에 대해 후불청구서를 발행하지 않습니다.
# CI/CD 통합의 Snyk 테스트 및 Snyk 모니터

접근 방식과 목표에 따라, 파이프라인에서 `snyk monitor` 및 `snyk test` 명령을 모두 사용해야 할 수 있습니다. 아래에는 예시와 자세한 내용이 제시되어 있습니다.

## **빌드 파이프라인에서의 CLI 예시**

`snyk monitor`를 사용하여 취약점을 노출하고 지속적 모니터링을 위해 Snyk UI에 전송합니다:

```
snyk monitor --all-projects --org=snyk-apps
```

높은 심각도 문제에서 빌드 실패시키기 위해 `snyk test`를 사용합니다:

```
snyk test --all-projects --org=snyk-apps --severity-threshold=high
```

CLI에서 모든 옵션 목록을 보려면 `snyk test --help`, `snyk monitor --help`, `snyk container --help` 명령을 실행하거나 [도움말 문서](../../../snyk-cli/commands/)를 참조하세요.

## **종료 코드**

`snyk test` 명령은 동기식이며 종료 코드와 함께 종료됩니다. 빌드 시스템은 종료 코드를 사용하여 테스트 결과에 따라 빌드를 성공 또는 실패로 처리할 수 있습니다. 사용 중인 명령어에 따른 종료 코드의 의미를 찾으려면 [도움말 문서](../../../snyk-cli/commands/)를 참조하십시오.

`snyk monitor` 명령어는 프로젝트의 종속성 트리에 대한 스냅샷을 Snyk 계정에 게시하고 해당 스냅샷을 취약점에 대해 모니터링합니다. 취약점 상태에 따라 종료 코드로 끝나지 않는 비동기 명령어입니다. `snyk monitor`의 경우 종료 코드는 스냅샷 생성의 성공 또는 실패를 나타냅니다.

빌드 단계에서 빌드 실패를 피하기 위해 `snyk test` 명령의 Snyk CLI 종료 코드를 무시하려면 명령어 끝에 `|| true`를 사용하십시오. `|| true`는 스캔의 종료 코드를 0으로 설정합니다. 이를 사용하여 취약점이 있을 때에도 CI/CD 파이프라인을 계속할 수 있습니다.

## CI/CD 통합에서의 일반적인 CLI 옵션

CI/CD 통합에서 사용되는 가장 일반적인 옵션은 다음과 같습니다:

`-- all-projects`: 작업 디렉토리에 있는 모든 프로젝트를 자동으로 감지합니다.

`-p`: 중복된 하위 종속성을 제거하여 종속성 트리를 정리합니다. 모든 취약점을 계속해서 찾지만 취약한 경로를 찾지 못할 수 있습니다.

\--org=\<ORG\_ID>: 특정 조직을 위해 Snyk 명령을 실행하는 ORG\_ID를 지정합니다. 이는 `monitor` 명령 실행 후 새 프로젝트가 생성되는 위치, 일부 기능의 가용성 및 개인 테스트 한도에 영향을 미칩니다. 여러 조직을 사용하는 경우 기본 설정을 CLI를 사용하여 설정할 수 있습니다:

```
$ snyk config set org=<ORG_ID>
```

기본 설정을 사용하여 모든 새로 테스트된 프로젝트 또는 모니터링된 프로젝트가 기본 조직 아래에서 테스트되도록 보장하십시오. 기본 설정을 무시하려면 `--org=<ORG_ID>` 옵션을 사용하십시오.

기본`<ORG_ID>`는 [계정 설정](https://app.snyk.io/account)의 현재 기본 조직입니다.

## 빌드 실패를 위한 옵션 구성

빌드 실패로 이어질 수 있는 매개변수를 정제하기 위해 `snyk test` 명령에 옵션을 추가할 수 있습니다:

* `--severity-threshold=high`: 높은 심각도 문제로만 빌드 실패
* `--fail-on=upgradable`: 업그레이드 가능한 문제로만 빌드 실패 (Snyk 수정 안내로 해결할 수 있음)

또한 [snyk-filter](https://github.com/snyk-tech-services/snyk-filter)와 같은 래퍼를 사용하거나 [snyk-delta](https://github.com/snyk-tech-services/snyk-delta)와 같은 추가 도구를 사용하여 기타 매개변수에 대해 빌드를 실패시킬 수 있습니다. snyk-filter 및 snyk-delta 사용에 대한 정보는 [Snyk CLI에서 빌드 실패](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/failing-of-builds-in-snyk-cli.md)를 참조하십시오.

## 문제 무시

기본적으로 문제가 무시되지 않거나 [snyk-delta](https://github.com/snyk-tech-services/snyk-delta)를 사용하지 않는 경우, 파이프라인에서 문제가 발견되면 `snyk test`로 빌드가 실패합니다. 이러한 문제를 해결하지 않고 빌드를 계속하려면 다음을 수행할 수 있습니다:

* [.snyk 정책 파일 사용하여 문제 무시](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/ignore-vulnerabilities-using-the-snyk-cli.md)
* Snyk UI에서 문제 무시
* Ignores API를 사용하여 문제 무시
* 대량 무시를 위해 Snyk Python API 사용: [pysnyk](https://github.com/snyk-labs/pysnyk) 및 데모 [bulk-ignore-vulns-by-issueIdList](https://github.com/snyk-labs/pysnyk/blob/master/examples/api-demo-9c-bulk-ignore-vulns-by-issueIdList.py)를 참조하십시오.

## 사용자 정의 빌드 아티팩트 생성

Snyk 명령으로부터의 JSON 출력을 사용하여 빌드 아티팩트로 사용자 지정 테스트 보고서를 만들 수 있습니다. [snyk-to-html](https://github.com/snyk/snyk-to-html) 유틸리티 또는 개발한 다른 사용자 지정 처리를 사용하여 작업할 수 있습니다.

## 새 취약점을 위한 작업 항목 생성

Snyk을 사용하여 JIRA에서 자동으로 새 작업 항목을 생성할 수 있습니다. 자세한 내용은 [Jira 통합](../../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md) 문서를 참조하십시오. 이 코드를 사용자 지정 요구 사항에 맞게 수정하거나 다른 작업 관리 시스템과 작업할 수 있도록 조정할 수 있습니다.

시작하려면 [Jira tickets for new vulns](https://github.com/snyk-tech-services/jira-tickets-for-new-vulns)를 참조하십시오. 또한 API 엔드포인트 [Jira 이슈 생성](../../../snyk-api/reference/jira-v1.md#org-orgid-project-projectid-issue-issueid-jira-issue) 및 [모든 Jira 이슈 나열](../../../snyk-api/reference/jira-v1.md#org-orgid-project-projectid-jira-issues)를 사용할 수 있습니다.

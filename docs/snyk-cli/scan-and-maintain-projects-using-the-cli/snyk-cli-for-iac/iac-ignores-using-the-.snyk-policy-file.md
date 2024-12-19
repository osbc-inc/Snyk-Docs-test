# IaC은 .snyk 정책 파일을 사용하지 않습니다

Snyk CLI `iac test` 명령을 사용하여 IaC 구성 파일을 스캔할 때, [`.snyk` 정책 파일](../../../manage-risk/policies/the-.snyk-file.md)을 **현재 IaC에 대해서만** 사용하여 관련이 없는 문제를 무시할 수 있습니다. Snyk은 `.snyk` 파일을 권장하며, 이 파일을 IaC 구성 파일을 저장하는 작업 디렉터리의 루트에 저장 및 버전 관리하는 것을 권장합니다. 이 파일은 `snyk ignore` 명령으로 생성할 수 있습니다. 자세한 내용은 [Snyk CLI를 사용한 취약점 무시](../ignore-vulnerabilities-using-the-snyk-cli.md)를 참조하십시오.

{% hint style="info" %}
현재 IaC 및 IaC+에 대해 Snyk 웹 UI를 통해 무시할 수 있습니다.
{% endhint %}

## 경로 무시

Snyk CLI를 사용하여 실행된 테스트에서는 `.snyk` 파일에 정의된 문제만 무시됩니다.

수입된 Git 리포지토리에서 실행된 테스트의 경우, 문제를 Snyk UI에서 무시할 수 있습니다. 이러한 무시는 Snyk UI를 사용하여 수행된 스캔에만 적용됨에 유의하십시오.

{% hint style="warning" %}
`.snyk` 파일에 있는 무시 및 Snyk UI에서 생성된 무시는 동기화되지 않습니다.
{% endhint %}

## `.snyk` 파일 의미론

`.snyk` 파일에는 IaC 프로젝트에 대한 일부 제한이 있습니다. 표준 기능에 대한 자세한 내용은 [`.snyk` 파일](../../../manage-risk/policies/the-.snyk-file.md)을 참조하십시오.

* **패치** 섹션은 아직 지원되지 않으며 무시됩니다.
* IaC 지원되는 **언어 설정**이 없습니다. 이 섹션은 무시됩니다.

특정 디렉터리에 대해 `snyk iac test`를 실행할 때, 하나 이상의 디렉터리를 전달하거나 현재 작업 디렉토리의 기본 인수를 사용하여 실행할 때 Snyk CLI는 각 디렉터리마다 `.snyk` 파일을 찾습니다.

정책 파일의 구문은 다음과 같습니다:

```
version: v1.19.0
ignore:
  SNYK-CC-K8S-1:
    - '*':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
```

`*` 객체 키는 CLI가 `SNYK-CC-K8S-1` 취약점의 모든 인스턴스를 무시하도록 합니다. 여러 취약점을 무시하려면 IaC 문제 ID로 키를 지정하여 여러 엔트리를 추가할 수 있습니다.

## 단일 파일 무시

무시 규칙은 더 좁은 범위로 지정할 수 있습니다. 무시를 특정 파일로 범위 지정하려면 테스트 중인 디렉터리에 있는 `.snyk` 정책 파일을 포함하는 디렉터리를 변경해야 합니다.

Snyk CLI를 사용하여 범위 지정 무시 규칙을 생성할 수 있습니다. 다음 명령을 실행하여 scoped 무시 규칙을 생성할 수 있습니다:

```
snyk ignore --id=SNYK-CC-K8S-1 --path='staging/cronjob.yaml > *'
snyk ignore --id=SNYK-CC-K8S-1 --path='staging/deployment.yaml > *'
```

또는 `.snyk` 정책 파일을 수동으로 수정하여 다음과 같이 한 파일에서 문제를 무시할 수 있습니다:

```
version: v1.19.0
ignore:
  SNYK-CC-K8S-1:
    - 'staging/deployment.yaml > *':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
  - 'staging/cronjob.yaml > *':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
```

Snyk CLI ignore 명령에 대한 자세한 정보는 [Snyk CLI를 사용한 취약점 무시](../ignore-vulnerabilities-using-the-snyk-cli.md)를 참조하십시오.

## 취약점의 인스턴스 무시

파일 내에서 취약점의 개별 인스턴스를 무시할 수 있습니다. 이를 위해서는 `snyk iac test`의 출력에서 "리소스 경로"를 취하여 파일 경로에 추가하면 됩니다.

예를 들어, 다음 출력 스니펫에서 (가독성을 위해 라인을 추가함):

```
Testing production/deployment.yaml...Infrastructure as code issues:
  ✗ Container is running in privileged mode [High Severity] [SNYK-CC-K8S-1] in Deployment
    introduced by [DocId: 0] > input > spec > template > spec > containers[web] 
    > securityContext > privileged
```

다음 명령을 실행하여 Snyk CLI를 사용하여 범위 지정 인스턴스를 무시할 수 있습니다:

```
 snyk ignore --id=SNYK-CC-K8S-1 --path='production/deployment.yaml > [DocId:1] > spec > template > spec > containers[web] 
 > securityContext > privileged'
```

또는 정책 파일을 다음과 같이 수동으로 수정할 수 있습니다:

```
version: v1.19.0
ignore:
  SNYK-CC-K8S-1:
    - 'production/deployment.yaml > [DocId:1] > spec > template > spec > containers[web] > securityContext > privileged':
        reason: None Given
        expires: 2021-08-26T08:40:35.249Z
        created: 2021-07-27T08:40:35.251Z
```

Snyk CLI ignore 명령에 대한 자세한 정보는 [Snyk CLI를 사용한 취약점 무시](../ignore-vulnerabilities-using-the-snyk-cli.md)를 참조하십시오.

## 정책 플래그 및 정책 파일 참고 사항

각 테스트에는 `.snyk` 정책 파일이 하나 이상이 없어야 합니다. 예를 들어, `snyk iac test dir1/ dir2/` 명령은 `dir1/.snyk` 및 `dir2/.snyk`을 로드하지만 `dir1/foo/bar/.snyk` 파일이 있는 경우 CLI는 해당 파일을 로드하지 않습니다.

`snyk iac test`를 실행할 때, CLI는 `$PWD/.snyk`를 로드합니다. 일반적인 패턴은 리포지토리 당 루트에 단일 `.snyk` 정책 파일을 사용하는 것입니다.

CLI는 `--policy-path=...` 옵션을 허용하며, 이 옵션은 `.snyk` 정책 파일의 위치를 재정의합니다. 경로는 `.snyk` 파일이 있는 디렉토리거나 `.snyk` 파일의 경로일 수 있습니다. 정책 파일의 이름은 반드시 `.snyk`여야 합니다.

`snyk iac test`에 대한 인수가 디렉토리가 아닌 파일인 경우, 정책을 로드하기 위해 `--policy-path`가 지정되어야 합니다.

CLI는 `--ignore-policy` 옵션을 허용하며, 발견된 `.snyk` 정책 파일을 무시합니다.
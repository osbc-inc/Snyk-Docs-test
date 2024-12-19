# IaC 파일을 테스트하세요

{% hint style="info" %}
CLI 버전 1.594.0부터 모든 구성 파일이 로컬에서 처리되어 머신을 벗어나지 않음이 보장됩니다. 이전 버전은 기본적으로 구성 파일을 처리하기 위해 Snyk에게 전송하였습니다. Snyk는 CLI의 최신 버전으로 업그레이드할 것을 권장합니다.
{% endhint %}

{% hint style="info" %}
[IaC+](../../../../scan-with-snyk/snyk-iac/getting-started-with-iac+-and-cloud-scans/) 버전의 Snyk CLI를 사용하려면, Snyk CLI v1.1022.0 이상을 설치하십시오.
{% endhint %}

## 개요

Snyk Infrastructure as Code를 사용하면 CLI로 구성 파일을 테스트할 수 있습니다. 이 페이지에서는 `snyk iac test` 명령어의 특정 옵션들을 사용하는 방법에 대해 자세한 정보를 제공합니다. 모든 옵션에 대한 정보는 `snyk iac test` 명령어 [도움말](../../../commands/iac-test.md)을 참조하십시오. 각 구성 파일을 테스트하는 세부 정보는 다음 페이지를 참조하십시오:

* [Terraform 파일](terraform-files.md)
* [CloudFormation 파일](cloudformation-files.md)
* [AWS CDK 파일](aws-cdk-files.md)
* [Kubernetes 파일](kubernetes-files.md)
* [ARM 파일](arm-files.md)
* [Kustomize 파일](kustomize-files.md)
* [Helm 차트](helm-charts.md)
* [Serverless 파일](serverless-files.md)

다음 예시에서는 `deployment.yaml`와 같이 샘플 파일 이름을 귀사의 파일 이름으로 교체할 수 있습니다.

## 지정된 파일에서 문제를 테스트

인수를 제공하지 않으면 `snyk iac test` 명령어는 현재 작업 디렉토리를 재귀적으로 통과하여 찾은 모든 파일을 스캔합니다:

```
snyk iac test
```

현재 작업 디렉토리 아래의 특정 파일을 스캔할 수 있습니다. 하나 이상의 파일 경로를 제공하면 해당 파일들만 스캔합니다:

```
snyk iac test file-1.tf dir/file-2.tf
```

명령어는 현재 작업 디렉토리 외부의 파일 경로를 제공하면 오류를 반환합니다. 예를 들어, 다음은 **명령어의 유효한 호출이 아닙니다**:

```
snyk iac test ../main.tf
```

## 파일 디렉토리의 문제를 테스트

인수를 제공하지 않으면 명령어는 현재 작업 디렉토리를 재귀적으로 통과하여 찾은 모든 파일을 스캔합니다:

```
snyk iac test
```

스캔을 특정 디렉토리로 제한할 수 있습니다:

```
snyk iac test my-folder
```

통과되는 디렉토리의 깊이를 제한할 수 있습니다. 현재 작업 디렉토리가 깊이 1을 갖고 있고, 현재 작업 디렉토리 하위 디렉토리는 깊이 2를 가지며, 이와 같은 방식으로 계속됩니다. 예를 들어, 현재 작업 디렉토리와 두 개 이상의 디렉토리까지 검색을 제한하려면 다음과 같이 명령어를 호출할 수 있습니다:

```
snyk iac test --detection-depth=3
```

명령어는 현재 작업 디렉토리 외부의 디렉토리 경로를 제공하면 오류를 반환합니다. 예를 들어, 다음은 **명령어의 유효한 호출이 아닙니다**:

```
snyk iac test ../my-folder
```

## 테스트 형식을 JSON으로 출력

다음 명령어를 사용하여 JSON 파일 형식으로 출력을 받을 수 있습니다:

```
snyk iac test --json
```

로컬에 결과 스냅샷 저장이나 다른 도구를 통해 결과를 보고 분석할 수 있을 때 유용할 수 있습니다.

예시:

```
snyk iac test main.tf --json
```

## 테스트 형식을 SARIF로 출력

SARIF는 정적 분석 도구의 출력을 위한 오픈 표준입니다. 테스트 결과를 다른 도구에서 분석하기 위해 SARIF 파일 형식으로 볼 수 있고 저장할 수 있습니다.

다음 명령어를 사용하여 SARIF 파일 형식으로 출력을 받을 수 있습니다:

```
snyk iac test main.tf --sarif
```

이를 파일 출력에 저장하려면 다음 명령어를 실행할 수 있습니다:

```
snyk iac test main.tf --sarif-file-output=snyk.sarif
```

## 특정 심각도 수준 이상의 문제만 표시

다음 명령어를 사용하여 표시된 결과를 특정 심각도 수준 이상의 문제로 제한할 수 있습니다.

```
snyk iac test --severity-threshold=medium
```

예시:

```
snyk iac test main.tf --severity-threshold=medium
```

이 명령은 중간 또는 그 이상의 심각도 값을 갖는 결과만 표시합니다.

## 특정 Snyk 기관을 대상으로 지정

Snyk UI에서 기관 수준에서 보안 규칙의 심각도 설정을 제어할 수 있습니다. CLI 테스트에서 특정 기관을 대상으로 지정함으로써 실행해야 할 규칙과 그 심각도를 결정할 수 있습니다.

다음 명령어를 사용하여 기관을 지정합니다:

```
snyk iac test --org=infrastructure
```

예시:

```
snyk iac test main.tf --org=infrastructure
```

`--org` 옵션을 사용하여 기관을 지정할 때마다 `--org` 옵션을 사용할 필요 없이 `snyk config`에서 `org` 플래그를 설정할 수도 있습니다.

```
snyk config set org=infrastructure
```

## 예시 테스트 결과

```
Snyk Infrastructure as Code

✔ 테스트 완료.

문제

낮은 심각도 문제: 1

  [낮음] API Gateway 접근 로깅 비활성화
  정보:    Amazon Api Gateway 접근 로깅이 활성화되지 않았습니다. 조사 중에 감사 레코드를 사용할 수 없을 수 있습니다.
  규칙:    https://security.snyk.io/rules/cloud/SNYK-CC-TF-118
  경로:    resource > aws_api_gateway_stage[denied] > access_log_settings
  파일:    aws_api_gateway_stage_logging.tf
  해결: `access_log_settings` 속성 설정

-------------------------------------------------------

테스트 요약

  기관: demo-org

✔ 문제가 없는 파일: 0
✗ 문제가 있는 파일: 1
  잘못된 파일: 0
  무시된 문제: 0
  총 문제: 1 [ 0 중요, 0 높음, 0 중간, 1 낮음 ]
```
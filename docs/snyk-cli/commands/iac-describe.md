# IaC 설명

## 사용법

**참고:** 이 기능은 Snyk CLI 버전 v1.876.0 이상에서 사용할 수 있습니다.

`snyk iac describe [<OPTIONS>]`

## 설명

`snyk iac describe` 명령어는 관리되지 않는 인프라 리소스를 감지합니다. 이 명령어는 Terraform 상태 파일에 있는 리소스를 실제 클라우드 제공업체의 리소스와 비교하여 보고서를 출력합니다.

* Terraform 상태 파일에 있는 리소스는 **관리되는 리소스**입니다.
* 존재하는 리소스 중 Terraform 상태 파일에 포함되지 않는 리소스는 **관리되지 않는 리소스**입니다.

자세한 정보 및 예제는 [IaC describe 명령어 예제](https://docs.snyk.io/scan-using-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/iac-describe-command-examples)를 참조하십시오.

관련 명령어 목록은 snyk [iac 도움말](iac.md) 및 `iac --help`를 참조하십시오.

## 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공, 변조가 발견되지 않음\
**1**: 관리되지 않는 리소스가 발견됨\
**2**: 실패

## Snyk CLI 구성

환경 변수 및 Snyk API와 연결하기 위해 변수를 설정할 수 있습니다. 자세한 내용은 [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)을 참조하십시오.

## Terraform 공급자 구성

`describe` 명령어에서 사용하는 Terraform 공급자를 구성하기 위해 환경 변수를 설정할 수 있습니다. 자세한 내용은 [클라우드 제공자 구성](https://docs.snyk.io/scan-using-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/configure-cloud-providers)을 참조하십시오.

## 디버깅

디버그 로그를 출력하려면 `-d` 옵션을 사용합니다.

## 선택적 매개변수

### `--org=<ORG_ID>`

특정 Snyk 기관에 연결된 Snyk 명령을 실행하려면 `<ORG_ID>`를 지정하십시오. 현재 선호하는 기관이 본인의 [계정 설정](https://app.snyk.io/account)에 있는 기본 `<ORG_ID>`를 덮어씁니다.

`--org=<orgslugname>`을 사용할 수도 있습니다. `ORG_ID`는 CLI와 API에서 모두 작동합니다. 조직 슬러그 이름은 CLI에서 작동하지만 API에서 작동하지 않습니다.

더 많은 정보는 [CLI에서 사용할 조직 선택하는 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli)을 참조하십시오.

### `--from=<STATE>[,<STATE>...]`

읽을 여러 Terraform 상태 파일을 지정합니다. Glob 패턴이 지원됩니다.

**지원되는 IaC 소스 목록** 및 사용 방법 등 자세한 정보는 [IAC Sources usage](https://docs.snyk.io/scan-using-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/iac-sources-usage)를 참조하십시오.

### `--to=<PROVIDER+TYPE>`

스캔할 클라우드 제공자를 지정합니다 (기본: AWS with Terraform).

지원되는 공급자:

* `github+tf` (GitHub with Terraform)
* `aws+tf` (Amazon Web Services with Terraform)
* `gcp+tf` (Google Cloud Platform with Terraform)
* `azure+tf` (Azure with Terraform)

### `--tf-provider-version`

사용할 Terraform 공급자 버전을 지정합니다. 지정하지 않으면 다음과 같이 기본 버전이 사용됩니다:

* aws@3.19.0
* github@4.4.0
* google@3.78.0
* azurerm@2.71.0

### `--tf-lockfile`

사용자 정의 경로에서 Terraform 락 파일(`.terraform.lock.hcl`)을 읽습니다 (기본: 현재 디렉토리).

락 파일 파싱에 실패하면 오류가 기록되고 스캔이 계속됩니다.

**참고**: `--tf-lockfile` 및 `--tf-provider-version` 옵션을 함께 사용하는 경우 `--tf-provider-version`이 우선합니다.

### `--fetch-tfstate-headers`

Terraform 상태를 가져올 때 HTTP 백엔드에 대한 특정 HTTP 헤더를 사용합니다.

### `--tfc-token`

Terraform Cloud 또는 엔터프라이즈 API에 인증하기 위한 API 토큰을 지정합니다.

### `--tfc-endpoint`

조직의 Terraform Enterprise 설치에 특정한 'tfc-endpoint' 값을 전달하여 주어진 워크스페이스의 현재 상태를 읽습니다.

### `--config-dir`

`iac describe` 구성에 사용되는 디렉토리 경로를 변경합니다 (기본값: `$HOME`). 예를 들어 AWS Lambda 함수에서만 `/tmp` 폴더를 사용할 수 있는 경우 유용할 수 있습니다.

## 리소스 포함 및 제외 옵션

### `--service=<SERVICE>[,<SERVICE>...]`

관리되지 않는 리소스를 검사할 서비스를 지정합니다.

이 옵션은 `.snyk` 드리프트 무시 규칙과 함께 사용할 수 없습니다. `.snyk` 내의 콘텐츠는 무시됩니다.

지원되는 서비스: `aws_s3`, `aws_ec2`, `aws_lambda`, `aws_rds`, `aws_route53`, `aws_iam`, `aws_vpc`, `aws_api_gateway`, `aws_apigatewayv2`, `aws_sqs`, `aws_sns`, `aws_ecr`, `aws_cloudfront`, `aws_kms`, `aws_dynamodb`, `azure_base`, `azure_compute`, `azure_storage`, `azure_network`, `azure_container`, `azure_database`, `azure_loadbalancer`, `azure_private_dns`, `google_cloud_platform`, `google_cloud_storage`, `google_compute_engine`, `google_cloud_dns`, `google_cloud_bigtable`, `google_cloud_bigquery`, `google_cloud_functions`, `google_cloud_sql`, `google_cloud_run`

### `--filter`

필터 규칙 사용.

필터 규칙을 사용하면 보고서에서 특정 리소스 집합을 포함하거나 제외하기 위해 JMESPath 표현식을 작성할 수 있습니다.

자세한 내용은 [Filter results](https://docs.snyk.io/scan-using-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/filter-rules)을 참조하십시오.

### `--strict`

엄격 모드를 활성화합니다.

`iac describe` 명령어는 기본적으로 서비스 연결형 리소스를 무시합니다(예: 서비스 연결형 AWS IAM 역할, 정책 및 정책 첨부). 이러한 리소스를 보고서에 포함하려면 **엄격 모드**를 활성화할 수 있습니다. AWS 계정과 사용할 때 소음을 만들 수 있음에 유의하십시오.

## 정책 옵션

### `--ignore-policy`

설정된 모든 정책, `.snyk` 파일의 현재 정책, 기관 수준의 무시 그리고 Snyk 웹 UI의 프로젝트 정책을 모두 무시합니다.

### `--policy-path=<POLICY_FILE_PATH>`

수동으로 `.snyk` 정책 파일의 경로를 전달합니다.

## 출력 옵션

### `--quiet`

스캔 결과만 stdout에 출력합니다.

### `--json`

보고서를 JSON 데이터 구조로 stdout에 출력합니다.

### `--html`

보고서를 html로 stdout에 출력합니다.

### `--html-file-output=<OUTPUT_FILE_PATH>`

보고서를 html로 파일로 출력합니다.

## snyk iac describe 명령어 예제

더 많은 예제는 [IaC describe 명령어 예제](https://docs.snyk.io/scan-using-snyk/scan-infrastructure/iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/iac-describe-command-examples)를 참조하십시오.

### 로컬 Terraform 상태 파일에서 AWS에서 관리되지 않는 리소스 감지

```
$ snyk iac describe --from="tfstate://terraform.tfstate"
```

### AWS 자격 증명 지정

```
$ AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=XXX snyk iac describe
```

### AWS 네임드 프로필 사용

```
$ AWS_PROFILE=프로필_이름 snyk iac describe
```

### S3 백엔드에 저장된 단일 Terraform 상태 사용

```
$ snyk iac describe --from="tfstate+s3://my-bucket/path/to/state.tfstate"
```

### 여러 Terraform 상태 집계

```
$ snyk iac describe --from="tfstate://terraform_S3.tfstate,tfstate://terraform_VPC.tfstate"
```

### 전체 패턴을 사용하여 여러 Terraform 상태 집계

```
$ snyk iac describe --from="tfstate://path/to/**/*.tfstate"
```  
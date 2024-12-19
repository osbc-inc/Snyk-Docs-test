# IaC describe command examples

`snyk iac describe` 옵션의 완전한 목록은 [`snyk iac describe`](../../../../snyk-cli/commands/iac-describe.md) 명령어 도움말을 참조하거나 다음 명령을 실행하여 도움말을 표시합니다:

```
snyk iac describe --help
```

## `--from`을 사용하여 상태 파일 지정

주어진 디렉토리의 모든 Terraform 상태를 읽고 집계합니다:

```
snyk iac describe --from="tfstate://directory/*.tfstate"
```

`terraform`를 사용하여 지원되지 않는 백엔드를 사용해 파일에 상태를 파이프하고 그 파일을 사용합니다:

```
terraform state pull > state.tfstate

snyk iac describe --from="tfstate://state.tfstate"
```

## `--to`를 사용하여 스캔할 클라우드 제공업체 지정

Terraform 컨텍스트에서 명시적으로 AWS를 스캔합니다:

```
snyk iac describe --to="aws+tf"
```

## `--tf-provider-version`를 사용하여 Terraform 제공자 버전 지정

스캔 오류를 피하기 위해 terraform 제공자 3.43.0을 지정합니다:

```
snyk iac describe --tf-provider-version=3.43.0
```

모든 클라우드 제공업체에 동일한 매개변수를 사용합니다:

```
snyk iac describe --to="github+tf" --tf-provider-version=4.10.1
```

## `--tf-lockfile`를 사용하여 Terraform 락 파일 지정

Terraform 락 파일(`.terraform.lock.hcl`)의 사용자 정의 경로를 지정합니다:

```
snyk iac describe --to="aws+tf" --tf-lockfile="/path/to/.terraform.lock.hcl"
```

## `--fetch-tfstate-headers`를 사용하여 Terraform 상태를 가져올 때 HTTP 헤더 지정

GitLab에 저장된 Terraform 상태를 사용하려면 HTTPS 인증을 지정합니다:

```
GITLAB_TOKEN=<access_token> \
  snyk iac describe \
  --from="tfstate+https://gitlab.com/api/v4/projects/<project_id>/terraform/state/<path_to_state>" \
 --fetch-tfstate-headers='Authorization="Bearer ${GITLAB_TOKEN}"'
```

## `--tfc-endpoint`를 사용하여 Terraform Enterprise 워크스페이스에서 상태 읽기

Terraform Enterprise 워크스페이스의 **일반 설정**에서 작업영역 ID를 얻을 수 있습니다.

Terraform Enterprise API 토큰을 제공해야 합니다.

예제:

```
snyk iac describe --from="tfstate+tfcloud://$WORKSPACE_ID" --tfc-token="$TFC_TOKEN" --tfc-endpoint="https://tfe.example.com/api/v2"
```

## `--service`를 사용하여 검사할 여러 서비스 지정

보고서에 AWS S3 및 AWS EC2 리소스를 포함합니다:

```
snyk iac describe --service="aws_s3,aws_ec2"
```

## `--strict`를 사용하여 서비스 링크된 리소스를 보고서에 포함

**참고:** AWS 계정에서 엄격 모드를 사용할 때, 귀하에게 속하지 않은 리소스로 인한 불필요한 노이즈가 발생할 수 있습니다.

이는 기본적으로 귀하의 계정과 관련된 서비스 링크된 역할이 있는 조직 계정을 보유하는 경우 발생할 수 있습니다. 예를 들어, `AWSServiceRoleForOrganizations`와 같은 역할입니다.

엄격 모드를 활성화하는 예제:

```
snyk iac describe --strict
```

## `--json`을 사용하여 보고서를 JSON 형식으로 출력

리디렉션을 통해 보고서를 JSON 파일로 저장합니다:

```
snyk iac describe --json > report.json
```
# IaC 테스트

## 사용법

`snyk iac test [<OPTIONS>] [<PATH>]`

## 설명

`snyk iac test` 명령어는 알려진 보안 문제를 테스트합니다.

관련 명령어 목록은 [snyk iac](iac.md) 도움말을 참조하세요; `iac --help`

추가 정보는 [Snyk CLI for IaC](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac)를 참조하세요.

## 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공 (스캔 완료), 취약점 없음\
**1**: 조치가 필요함 (스캔 완료), 취약점 발견\
**2**: 실패, 명령어 재실행을 시도해보세요. 디버그 로그를 출력하려면 `-d`를 사용하세요.\
**3**: 실패, 지원되는 프로젝트가 감지되지 않음

## Snyk CLI 구성

환경 변수를 사용하여 Snyk CLI를 구성하고 Snyk API와의 연결을 설정할 수 있습니다. [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)을 참조하세요.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하세요.

## 옵션

### `--detection-depth=<깊이>`

찾을 하위 디렉토리 수를 나타냅니다. `깊이`는 숫자여야 하며 1 이상이어야 합니다. 제로(0)는 현재 디렉토리를 나타냅니다.

기본값: 제한 없음.

예: `--detection-depth=3`은 지정된 디렉토리(또는 지정된 `<PATH>`가 없는 경우 현재 디렉토리)와 세 수준의 하위 디렉토리까지 검색을 제한합니다. 제로(0)는 현재 디렉토리를 나타냅니다.

### `--org=<ORG_ID>`

특정 Snyk 조직과 관련된 Snyk 명령을 실행하려면 `<ORG_ID>`를 지정하세요. `<ORG_ID>`는 개인 테스트 제한에 영향을 줍니다.

여러 조직이 있는 경우 CLI를 사용하여 기본값을 설정할 수 있습니다.

`$ snyk config set org=<ORG_ID>`을 사용하여 CLI에서 기본값을 설정할 수 있습니다.

기본값: [계정 설정](https://app.snyk.io/account)에서 현재 설정된 기본 조직의 `<ORG_ID>`

**참고:** `--org=<orgslugname>`을 사용할 수도 있습니다. `ORG_ID`는 CLI 및 API에서 모두 작동합니다. 조직 슬러그 이름은 CLI에서는 작동하지만 API에서는 작동하지 않습니다.

`orgslugname`은 UI의 조직 URL에서 표시된 슬러그 이름과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`. orgname은 작동하지 않습니다.

자세한 정보는 [CLI에서 사용할 조직 선택 방법](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli)을 참조하세요.

### `--ignore-policy`

모든 설정된 정책, `.snyk` 파일의 현재 정책, rg 레벨 무시, snyk.io의 프로젝트 정책을 무시합니다.

### `--policy-path=<정책_파일_경로>`

수동으로 `.snyk` 정책 파일의 경로를 전달합니다.

### `--json`

콘솔에 결과를 JSON 데이터 구조로 출력합니다.

예: `$ snyk iac test --json`

### `--json-file-output=<출력_파일_경로>`

결과를 JSON 데이터 구조로 지정된 파일에 직접 저장합니다. `--json` 옵션을 사용하든 사용하지 않든 상관없습니다.

사용하기 위해 `stdout`에서 인간이 읽을 수 있는 테스트 결과를 표시하고 동시에 JSON 데이터 구조 출력을 파일에 저장합니다.

예: `$ snyk iac test --json-file-output=vuln.json`

### `--sarif`

결과를 SARIF 형식으로 반환합니다.

### `--sarif-file-output=<출력_파일_경로>`

결과를 SARIF 형식으로 지정된 파일에 직접 저장합니다. `--sarif` 옵션을 사용하든 사용하지 않든 상관없습니다.

`stdout`에서 인간이 읽을 수 있는 테스트 결과를 표시하고 동시에 SARIF 형식 출력을 파일에 저장하는 데 특히 유용합니다.

참고: 프로젝트 속성을 설정하는 옵션을 사용하고 귀하의 역할이 프로젝트 속성을 편집할 권한이 없는 경우 `iac test` 명령이 실패합니다. 계속 진행하는 방법에 대한 지침은 [Snyk CLI에서 프로젝트 속성 편집에 필요한 권한](https://docs.snyk.io/snyk-admin/user-roles/user-role-management#permissions-required-to-edit-project-attributes-from-the-snyk-cli)을 참조하세요.

### `--project-business-criticality=<업무_비평가>[,<업무_비평가>]...>`

`--report` 옵션과 함께 사용할 수 있습니다.

프로젝트 비즈니스 중요도 프로젝트 속성을 하나 이상의 값으로 설정합니다(쉼표로 구분). 프로젝트 비즈니스 중요도를 지우려면`--project-business-criticality=`로 설정하세요.

허용된 값: `critical, high, medium, low`

자세한 정보는 프로젝트 속성을 참조하세요.

IaC+에서는 이 옵션을 지원하지 않습니다.

### `--project-environment=<환경>[,<환경>]...>`

`--report` 옵션과 함께 사용할 수 있습니다.

프로젝트 환경 프로젝트 속성을 하나 이상의 값으로 설정합니다(쉼표로 구분). 프로젝트 환경을 지우려면`--project-environment=`로 설정하세요.

허용된 값: `frontend`, `backend`, `internal`, `external`, `mobile`, `saas`, `onprem`, `hosted`, `distributed`

자세한 정보는 [프로젝트 속성](https://docs.snyk.io/manage-issues/introduction-to-snyk-projects/project-attributes)을 참조하세요.

IaC+에서는 이 옵션을 지원하지 않습니다.

### `--project-lifecycle=<생명주기>[,<생명주기>]...>`

`--report` 옵션과 함께 사용할 수 있습니다.

프로젝트 라이프사이클 프로젝트 속성을 하나 이상의 값으로 설정합니다(쉼표로 구분). 프로젝트 라이프사이클을 지우려면 `--project-lifecycle=`로 설정하세요.

허용된 값: `production`, `development`, `sandbox`

자세한 정보는 [프로젝트 속성](https://docs.snyk.io/manage-issues/introduction-to-snyk-projects/project-attributes)을 참조하세요.

IaC+에서는 이 옵션을 지원하지 않습니다.

### `--project-tags=<태그>[,<태그>]...>`

`--report` 옵션과 함께 사용할 수 있습니다.

프로젝트에 태그를 하나 이상의 값으로 설정합니다(키 및 값이 "="로 구분된 쉼표로 구분된 키 값 쌍). 예: `--project-tags=department=finance,team=alpha`

프로젝트 태그를 지우려면 `--project-tags=`로 설정하세요.

IaC+에서는 이 옵션을 지원하지 않습니다.

허용되는 문자 포함하여 자세한 정보는 [프로젝트 태그](https://docs.snyk.io/manage-issues/introduction-to-snyk-projects/project-tags)를 참조하세요.

### `--remote-repo-url=<URL>`

저장소를 위한 원격 URL을 설정하거나 재정의합니다.

단일 대상에 속한 모든 프로젝트를 그룹화합니다.

`--report` 옵션과 함께 사용할 수 있습니다.

### `--report`

**새로운** 옵션: 결과를 Snyk 웹 UI와 공유합니다.

현재 구성 문제 또는 응용 프로그램의 스냅숏을 만들거나 기존 프로젝트에 스냅숏을 추가하여 Snyk 계정에서 프로젝트를 만듭니다.

이 옵션을 사용한 후에는 Snyk 웹 사이트에 로그인하여 스냅숏을 확인할 수 있습니다.

예: `$ snyk iac test --report`

참고: 이 옵션은 `--rules` 옵션과 함께 사용할 수 없습니다.

### `--rules=<사용자_지정_규칙_번들_파일_경로>`

IaC 스캔에서 사용자 지정 규칙 스캔을 활성화하기 위해 `snyk-iac-rules` SDK로 생성된 사용자 지정 규칙 번들을 사용하도록 하는 전용 옵션입니다. [`snyk-iac-rules` SDK](https://github.com/snyk/snyk-iac-rules#readme)를 참조하세요.

이 옵션은 사용자 지정 규칙 설정이 Snyk UI로 구성된 경우 사용할 수 없습니다. 기본값: `--rules` 옵션이 지정되지 않은 경우 내부 Snyk 규칙만 사용하여 구성 파일을 스캔합니다.

예: 내부 Snyk 규칙 및 사용자 지정 규칙을 사용하여 구성 파일을 스캔합니다.

`--rules=bundle.tar.gz`

참고: 이 옵션은 `--report` 옵션과 함께 사용할 수 없습니다.

IaC+에서는 이 옵션을 지원하지 않습니다.

### `--severity-threshold=<낮음|중간|높음|심각>`

지정된 수준 또는 그 이상의 취약점만 보고합니다.

### `--scan=<TERRAFORM_PLAN_SCAN_MODE>`

Terraform 계획 스캔 모드를 사용하여 전체 최종 상태(예: `planned-values`) 또는 제안된 변경 사항만 분석하도록 스캔을 제어합니다.

기본값: `--scan` 옵션이 지정되지 않은 경우 기본적으로 제안된 변경 사항만 스캔합니다. 예 1: `--scan=planned-values` (전체 상태 스캔)\
예 2: `--scan=resource-changes` (제안된 변경 사항 스캔)

### `--target-name=<대상_이름>`

`--report` 옵션과 함께 사용할 수 있습니다.

저장소를 위한 프로젝트 이름을 설정하거나 재정의합니다.

참고: 이 옵션은 `--remote-repo-url`보다 우선합니다. 두 옵션을 함께 사용하는 경우 이 옵션이 우선합니다.

### `--target-reference=<대상_참조>`

`--report` 옵션과 함께 사용할 수 있습니다.

예를 들어 브랜치 이름이나 버전과 같이이 프로젝트를 구분하는 참조를 지정합니다. 동일한 참조를 가진 프로젝트는 해당 참조를 기반으로 그룹화될 수 있습니다.

예: 현재 Git 브랜치 설정하기:

`snyk iac test myproject/ --report --target-reference="$(git branch --show-current)"`

예: 최신 Git 태그로 설정하기:

`snyk iac test myproject/ --report --target-reference="$(git describe --tags --abbrev=0)"`

### `--var-file=<변수_파일_경로>`

스캔된 디렉토리와 다른 디렉토리에 있는 테라폼 변수 정의 파일을 불러옵니다.

예:

`$ snyk iac test myproject/staging/networking --var-file=myproject/vars.tf`

### `--snyk-cloud-environment=<환경_ID>`

Snyk 클라우드 환경에서 마지막 스캔을 사용하여 문제를 억제합니다. 자세한 정보는 [IaC 테스트에 클라우드 구성 추가](https://docs.snyk.io/scan-cloud-deployment/snyk-infrastructure-as-code/integrated-infrastructure-as-code/adding-cloud-context-to-your-iac-test)을 참조하세요.

이 옵션은 IaC+에서만 지원됩니다.

예:

`$ snyk iac test --snyk-cloud-environment=0d19dc1a-c2aa-4719-89ee-5f281dd92a20`

## snyk iac test 명령어에 대한 예시

자세한 정보는 [Snyk CLI for Infrastructure as Code](https://docs.snyk.io/scan-cloud-deployment/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code)를 참조하세요.

### CloudFormation 파일 테스트

```
$ snyk iac test /path/to/cloudformation_file.yaml
```

### Kubernetes 파일 테스트

```
$ snyk iac test /path/to/kubernetes_file.yaml
```

### Terraform 파일 테스트

```
$ snyk iac test /path/to/terraform_file.tf
```

### Terraform 계획 파일 테스트

```
$ terraform plan -out=tfplan.binary
$ terraform show -json tfplan.binary > tf-plan.json
$ snyk iac test tf-plan.json
```

### ARM 파일 테스트

```
$ snyk iac test /path/to/arm_file.json
```

### 디렉토리에서 일치하는 파일 테스트

```
$ snyk iac test /path/to/directory
```

### 로컬 사용자 지정 규칙 번들을 사용하여 디렉토리의 일치하는 파일 테스트

```
$ snyk iac test /path/to/directory --rules=bundle.tar.gz
```  
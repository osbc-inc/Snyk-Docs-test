# 무시

## 사용법

`snyk ignore --id=<ISSUE_ID> [--expiry=] [--reason=] [--policy-path=<POLICY_FILE_PATH>] [--path=<RESOURCE_PATH>] [옵션]`

또는

`snyk ignore [--expiry=] [--reason=] [--policy-path=<POLICY_FILE_PATH>] [--file-path=<RESOURCE_PATH>] [옵션]`

## 설명

`snyk ignore` 명령은 `.snyk` 정책 파일을 수정하여 지정된 문제를 해당 Snyk ID에 따라 모든 발생을 무시하거나 만료 날짜, 이유 또는 파일 시스템의 경로에 따라 무시합니다.

로컬 `.snyk` 파일을 다음과 비슷한 블록으로 업데이트합니다.

```yaml
ignore:
  '<ISSUE_ID>':
    - '*':
        reason: <이유>
        expires: <만료일>
```

`--path` 옵션을 사용할 때 블록은 다음과 같습니다.

```yaml
ignore:
  '<ISSUE_ID>':
    - '<RESOURCE_PATH>':
        reason: <이유>
        expires: <만료일>
```

`--file-path` 옵션을 사용하면 블록은 다음과 같습니다.

```yaml
exclude:
  '<그룹>':
    - <일치 패턴 파일>
    - <일치 패턴 파일>:
      reason: <이유>
      expires: <만료일>
      created: <생성 시간>
```

**참고**: `--file-path` \[제외] 옵션은 Snyk 코드(SAST) 테스트 또는 오픈 소스 `--unmanaged` 테스트에서만 사용할 수 있으며 다른 테스트 유형에는 작동하지 않습니다.

`.snyk` 파일을 사용하여 문제나 취약점을 무시하는 것은 {{Snyk 코드}}에서 지원되지 않습니다.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.

## 옵션

### `--id=<ISSUE_ID>`

무시할 문제에 대한 Snyk ID. `--file-path`와 함께 사용할 경우 생략 가능; 다른 용례에서 필요함.

### `--expiry=<만료일>`

`YYYY-MM-DD` 형식의 만료 날짜.

지원되는 형식:

[ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html)

[RFC 2822](https://tools.ietf.org/html/rfc2822)

기본: 30일 또는 `--file-path`와 함께 사용할 경우 없음

### `--reason=<이유>`

문제를 무시하는 데 사용되는 인간이 읽을 수 있는 `<이유>`.

기본: 없음

**참고**: {{Snyk 코드}}에서 지원되지 않음.

### `--policy-path=<POLICY_FILE_PATH>`

수동으로 전달할 `.snyk` 정책 파일의 경로.

기본: 없음

### `--path=<RESOURCE_PATH>`

이슈를 무시하기 위해 의존 그래프 내의 리소스 경로.

무시 규칙의 범위를 좁히기 위해 사용합니다. 리소스 경로를 지정하지 않으면 모든 리소스가 무시됩니다.

[https://github.com/npm/node-semver#versions](https://github.com/npm/node-semver#versions)를 사용하여 경로에 구성 요소 버전을 지정할 수 있습니다.

사용하면 `--policy-path` 옵션을 따릅니다.

기본: 모두

### `--file-path=<RESOURCE_PATH>`

이슈를 무시할 파일 시스템. `snyk code` 및 `snyk test --unmanaged`에서 사용됨

기본: 없음

### `--file-path-group=[global|code|iac-drift]`

`--file-path`와 함께 사용되는 그루핑 방법, 그렇지 않으면 생략됨.

기본: global

## snyk ignore 명령어 예시

### 특정 취약점 무시

```
$ snyk ignore --id='npm:qs:20170213' --expiry='2021-01-10' --reason='모듈에는 이 취약점과 관련이 없음'
```

### 리소스 경로가 지정된 특정 취약점 무시

```
$ snyk ignore --id='SNYK-JS-PATHPARSE-1077067' --expiry='2021-01-10' --path='nyc@11.9.0 > istanbul-lib-report@1.1.3 > path-parse@1.0.5' --reason='모듈에는 이 취약점과 관련이 없음'$ snyk ignore --id='SNYK-JS-PATHPARSE-1077067' --expiry='2021-01-10' --path='nyc@11.9.0
```

### 리소스 경로가 지정된 특정 취약점 무시 (Windows)&#x20;

이 예제에서는 Windows에서 `snyk iac test`를 실행하면 작은 따옴표를 포함하거나 백 슬래시를 포함하는 파일 명세를 반환합니다:

규칙: [https://security.snyk.io/rules/cloud/SNYK-CC-TF-118](https://security.snyk.io/rules/cloud/SNYK-CC-TF-118)\
경로: resource > aws\_iam\_role\[OrganizationAccountAccessRole] > assume\_role\_policy\['Statement']\[0]\
파일: terraform\environment\com\iam.tf\


해당하는 `snyk ignore` 명령은 다음과 같습니다.

`snyk ignore --id=SNYK-CC-TF-118 --path="terraform\environment\com/iam.tf > resource > aws_iam_role[OrganizationAccountAccessRole] > assume_role_policy['Statement'][0]"`

### 리소스 경로가 지정된 특정 취약점 무시 (Linux, Mac OS)&#x20;

이 예에서 Linux 또는 Mac OS에서 `snyk iac test`를 실행하면 작은 따옴표와 슬래시를 포함하는 경로를 반환합니다:

규칙: [https://security.snyk.io/rules/cloud/SNYK-CC-TF-118](https://security.snyk.io/rules/cloud/SNYK-CC-TF-118)\
경로: resource > aws\_iam\_role\[OrganizationAccountAccessRole] > assume\_role\_policy\['Statement']\[0]\
파일: terraform/environment/com/iam.tf

해당하는 `snyk ignore` 명령은 다음과 같습니다.

`snyk ignore --id=SNYK-CC-TF-118 --path="terraform/environment/com/iam.tf > resource > aws_iam_role[OrganizationAccountAccessRole] > assume_role_policy['Statement'][0]"`

### 30일 동안 특정 취약점 무시

```
$ snyk ignore --id=npm:tough-cookie:20160722
```

### 2031-01-20까지 특정 파일 무시

특정 파일을 2031-01-20까지 `snyk test --unmanaged`에서 사용하도록 설정하고 미래를 위해 설명을 함께 제공합니다.

```
$ snyk ignore --file-path='./deps/curl-7.58.0/src/tool_msgs.c' --expiry='2031-01-20' --reason='패치된 파일'
```

### 글로브 표현식을 사용하여 파일 또는 폴더 무시 - {{Snyk 코드}} 및 `unmanaged` 전용

글로브 표현식과 일치하는 파일을 무시하려면 특정 그룹에 추가하십시오.

이것은 {{Snyk 코드}}에 적용됩니다. {{Snyk Open Source}} 및 컨테이너 또는 IaC에는 적용되지 않습니다.

```
$ snyk ignore --file-path='./**/vendor/**/*.cpp' --file-path-group='global'
```

## `snyk ignore` 명령에 대한 자세한 정보

자세한 정보는 다음을 참조하십시오:

* [Snyk CLI를 사용하여 취약점 무시](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/ignore-vulnerabilities-using-the-snyk-cli)
* [.snyk 정책 파일을 사용한 IaC 무시](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/iac-ignores-using-the-.snyk-policy-file)
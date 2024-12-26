# 사용자 정의 규칙 만들기

Snyk CLI를 사용하여 `main.rego` 라는 파일을 포함하는 사용자 정의 규칙 템플릿을 만들 수 있습니다. Rego를 사용하여 규칙을 작성한 다음 해당 규칙을 Snyk 플랫폼 전반에 일관되게 적용할 수 있습니다. 규칙을 마치면 다음과 같은 비밀번호 정책 강제 규칙 예제처럼 보일 수 있습니다:

{% code title="rules/NEW_PASSWORD_POLICY/main_test.rego" %}
```
package rules.NEW_PASSWORD_POLICY

import data.snyk

input_type := "cloud_scan"

metadata := {
	"id": "NEW_PASSWORD_POLICY",
	"severity": "high",
	"title": "Increase the number of characters in our password policy",
	"description": "All of our password policies now require a minimum of 17 characters instead of the CIS recommendation of 14 characters",
	"product": ["cloud"],
}

password_pol := snyk.resources("aws_iam_account_password_policy")[_]

deny[info] {
	count(password_pol) < 1 
	info := {
		"message": "This account does not contain a password policy",
		"resource": password_pol
		}
}

deny[info] {
	password_pol.minimum_password_length < 17
	info := {"resource": password_pol}
}

resources[info] {
	info := {"resource": password_pol}
}
```
{% endcode %}

## 사용자 정의 규칙 템플릿 만들기

다음 CLI 명령을 사용하세요:

```
snyk iac rules init
```

이 명령은 일련의 대화식 프롬프트를 소개하며 프로젝트, 규칙, 규칙 스펙(테스트용), 자원 관계를 설정할 수 있습니다. `rule`을 선택하면 `main.rego` 파일로 규칙의 템플릿을 만들기 위한 대화식 프롬프트 시리즈를 받게 됩니다.

**대화식 프롬프트로부터의 예시 출력:**

```
What do you want to initialize? rule
Rule ID: NEW_PASSWORD_POLICY
Title: Increase the number of characters in our password policy
Severity: high
Description: All of our password policies now require a minimum of 17 characters instead of the CIS recommendation of 14 characters
Is this rule intended for , Snyk Cloud, or both? cloud
Input type: tf
Does this rule need more than one resource type? No
Resource type: aws_iam_account_password_policy
```

## 사용자 정의 규칙 작성

템플릿이 생성되면 규칙의 로직을 정의하고 규칙을 [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/)로 작성해야 합니다. 규칙에 추가할 수 있는 추가 메타데이터에 대한 세부 정보는 Snyk 정책 엔진 리포지토리의 [메타데이터 페이지](https://github.com/snyk/policy-engine/blob/main/docs/policy\_spec.md#metadata)를 참조하십시오.
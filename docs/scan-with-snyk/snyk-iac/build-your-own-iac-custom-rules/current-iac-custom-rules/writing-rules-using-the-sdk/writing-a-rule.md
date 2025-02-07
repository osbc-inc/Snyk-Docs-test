# 규칙 작성하기

## Rego에서의 규칙

규칙은 Rego에서 작성됩니다. Rego를 작성할 때 두 가지를 수행합니다.

1. 정책 결정을 내리는 **규칙**을 작성합니다. 규칙은 조건부 할당입니다.
2. **정책**에 규칙을 조직화합니다. 정책은 계층적 이름을 가진 규칙의 집합입니다.

Policy Language에 대해 더 알아보려면 [OPA Policy Language Documentation Page](https://www.openpolicyagent.org/docs/latest/policy-language/)를 방문하십시오.

{% hint style="info" %}
[OPA Playground](https://play.openpolicyagent.org)을 사용하여 Rego를 시도하거나 이 가이드의 예제를 실행할 수도 있습니다.
{% endhint %}

## 새 규칙 생성 방법

시작하는 데 두 가지 옵션이 있습니다.

1.  규칙을 작성하기 위해 필요한 파일을 생성하려면 `template` 명령을 사용하십시오:

    ```
    snyk-iac-rules template --rule <RULE-NAME> --format <hcl2|json|yaml|tf-plan>
    ```

    이 명령은 제공된 구성 형식을 기반으로 규칙에 대한 틀을 생성합니다. 자세한 내용은 [템플릿 명령에 대한 문서](../sdk-reference.md#template-options)를 읽어보세요.
2.  Rego 정책을 처음부터 작성하고 본인이 예상하는 파일 및 폴더 구조와 일치하도록 일치시킵니다:

    `rules`\
    `└── my_rule`\
    `├── main.rego`\
    `└── main_test.rego`

{% hint style="info" %}
`template` 명령을 사용하지 않는 경우 Rego 테스트 프레임워크를 직접 작성해야 합니다.
{% endhint %}

## 규칙 구조

Rego에서 요청을 허용하거나 거부하는 문을 작성할 수 있습니다:

`allow { input.name == "alice" }` 또는 `deny { input.name == "alice" }`

{% hint style="info" %}
**`template`** 명령을 사용하여 규칙을 생성한 경우 기본 진입점은 \*\*`rules/deny`\*\*입니다. 이를 재정의하고 `deny` 외 다른 이름을 사용하려면 [규칙 번들링](bundling-rules.md) 섹션을 참조하십시오.
{% endhint %}

다음은 `snyk-iac-rules template --rule NEW-RULE --format hcl2`를 실행했을 때 생성된 별도의 규칙 스켈레톤의 모습입니다:

{% code title="rules/NEW-RULE/main.rego" %}
```
package rules

deny[msg] {
	resource := input.resource.test[name]
	resource.todo
	msg := {
		# Mandatory fields
		"publicId": "NEW-RULE",
		"title": "Default title",
		"severity": "low",
		"msg": sprintf("input.resource.test[%s].todo", [name]),
		# Optional fields
		"issue": "",
		"impact": "",
		"remediation": "",
		"references": [],
	}
}
```
{% endcode %}

{% hint style="warning" %}
출력이 올바르게 표시되도록 **msg** 속성 형식을 따라야 합니다.
{% endhint %}

속성은 다음과 같습니다:

* **publicId:** 팀별 고유한 네이밍 규칙으로, 예를 들어 COMPANY-001입니다. **이 값은 내부 Snyk 규칙과 구분하기 위해 “SNYK-”를 포함해서는 안 됩니다.**
* **title:** 이슈를 요약하는 간단한 제목입니다.
* **severity:** **low/medium/high/critical** 중 하나일 수 있습니다.
* **msg:** Snyk는 자원 이름 및 속성을 변경하여 예를 들어 `input.aws_s3_bucket[%s].tags`와 일치하도록 권장합니다. `sprintf` 함수는 Rego에서 제공되며 Snyk가 발견한 문제의 정확한 위치를 설명하는 동적 오류 메시지를 제공합니다.

다음 속성은 선택적이지만 Snyk CLI에서 스캔 결과를 향상시키는 데 사용할 수 있습니다:

* **issue:** 정확한 문제에 대한 자세한 설명 문자열입니다.
* **impact:** 이 문제를 해결하지 않았을 때의 영향에 대한 자세한 설명 문자열입니다.
* **remediation:** 문제를 해결하는 방법에 대한 자세한 설명 문자열입니다. Snyk는 여기에 코드 스니펫을 제공하는 것을 권장합니다.
* **references:** 문서 URL 문자열 배열을 제공할 수 있습니다.

규칙의 생성된 테스트는 허용된 및 거부된 Fixture에 대해 올바른 `msg` 필드가 규칙에 의해 반환되는지 확인하기 위해 두 개의 생성된 Terraform 파일을 사용합니다.

```
package rules

import data.lib
import data.lib.testing

test_NEW_RULE {
	# 규칙이 허용되는 테스트 케이스를 포함하는 배열
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]

	# 규칙이 거부되는 케이스를 포함하는 배열
	denied_test_cases := [{
		"want_msgs": ["input.resource.test[denied].todo"], # 거부된 규칙에 의해 반환된 올바른 msg임을 확인합니다.
		"fixture": "denied.tf",
	}]

	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("NEW-RULE", "./rules/NEW-RULE/fixtures", test_cases)
}
```

## 규칙 예시

{% hint style="info" %}
더 많은 예시는 [특정 규칙 예시](examples-of-iac-custom-rules.md)를 참조하십시오.
{% endhint %}

이 예시에서 템플릿 규칙은 `owner` 태그가 없는 자원에 `msg`를 할당하도록 수정되었습니다:

{% code title="rules/MY_RULE/main.rego" %}
```
package rules

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    not resource.tags.owner
	
    msg := {
        "publicId": "MY_RULE",
        "title": "Missing an owner from tag",
        "severity": "medium",
        "msg": sprintf("input.resource.aws_redshift_cluster[%s].tags", [name]),
        "issue": "",
        "impact": "",
        "remediation": "",
        "references": [],
    }
}
```
{% endcode %}

## 제한 사항 및 참고 사항

* Snyk은 Rego 정책을 Wasm 모듈로 컴파일하기 때문에 Wasm을 지원하는 내장 함수만 사용할 수 있습니다. 이러한 함수를 확인하는 표가 [Policy Reference Documentation](https://www.openpolicyagent.org/docs/latest/policy-reference/) 맨 아래에 있습니다.
* 규칙은 동일한 이름으로 여러 번 정의될 수 있습니다. 파일이나 동일한 패키지 내의 별도 파일에서 정의할 수 있습니다. 예를 들어,

```
packages rules

deny[msg] {
    resource.this
}
...

deny[msg] {
    resource.that
}
...
```

이러한 규칙은 `증분적(incremental)` 규칙이라고 하며 각 정의가 누적적입니다. [Policy Reference Documentation](https://www.openpolicyagent.org/docs/latest/policy-language/#incremental-definitions)에서 증분적 정의에 대해 더 읽어볼 수 있습니다. 동일한 이름을 가진 규칙은 다른 값을 반환해야 하며 그렇지 않으면 OPA에서 오류를 반환합니다. [Policy Reference Documentation](https://www.openpolicyagent.org/docs/latest/policy-language/#complete-definitions)에서 완전한 정의에 대해 더 읽어볼 수 있습니다.

* 더 복잡한 주제에 대해서는 [OPA Conflict Resolution 해결 방법](https://www.openpolicyagent.org/docs/latest/faq/#conflict-resolution)을 확인하십시오.

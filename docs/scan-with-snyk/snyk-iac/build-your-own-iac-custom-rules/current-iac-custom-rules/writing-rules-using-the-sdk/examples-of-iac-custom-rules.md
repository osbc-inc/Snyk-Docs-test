# IaC 사용자 지정 규칙의 예

## 간단한 부울 규칙 예시

{% hint style="info" %}
[이 OPA Playground](https://play.openpolicyagent.org/p/SCYndBjWxh)에서 이 가이드의 전체 예제를 찾을 수 있으며 [snyk/custom-rules-example](https://github.com/snyk/custom-rules-example) 저장소에서도 찾아볼 수 있습니다.
{% endhint %}

SDK를 사용하여 새 규칙 `CUSTOM-RULE-1`을 생성하고, 적용된 간단한 Terraform 리소스가 포함된 픽스처 파일이 있는 것으로 가정합니다:

{% code title="rules/CUSTOM-RULE-1/fixtures/denied.tf" %}
```
resource "aws_redshift_cluster" "denied" {
  cluster_identifier = "tf-redshift-cluster"
  node_type          = "dc1.large"
  tags = {
  }
}
```
{% endcode %}

이제, 생성된 Rego를 수정하여 소유자가 태그된 리소스를 강제로 적용하십시오:

1. `aws_redshift_cluster` 리소스 전체를 나열하기 위한 `[name]` 변수를 만듭니다. 이 변수는 `i`, `j`, `name` 등 원하는 이름으로 지정할 수 있습니다.
2. 월러스 연산자 `:=`를 사용하여 리소스 변수에 해당 값을 할당하여 이를 리소스에 저장하세요. 예를 들어`resource := input.resource.aws_redshift_cluster[name]`
3. 각 리소스에 대해 소유자 태그가 있는지 확인합니다. 이를 수행하기 위해 `resource.tags.owner` 경로가 정의되어 있는 지 확인하십시오. 정의되어 있지 않으면 정의되지 않은 것으로 판단됩니다. 따라서 `NOT` 키워드를 사용하여 이 키워드의 앞에 사용하여 `TRUE`로 평가되도록 합니다. 예를 들어`not resource.tags.owner`

수정된 Rego는 다음과 같습니다:

{% code title="rules/CUSTOM-RULE-1/main.rego" %}
```
package rules

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    not resource.tags.owner

    msg := {
        "publicId": "CUSTOM-RULE-1",
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

{% hint style="info" %}
이전에 제공된 Terraform 파일을 Rego 코드로 어떻게 평가하는지 알아보시려면 [픽스처 파일 구문 분석](parsing-an-input-file.md)을 참조하십시오.
{% endhint %}

{% hint style="info" %}
Snyk은 늘 규칙이 올바른지 확인하기 위해 항상 [규칙 업데이트 및 단위 테스트 실행](testing-a-rule.md)을 권장합니다.
{% endhint %}

이 규칙에 대한 테스트는 이 가이드의 시작점 격자가 부적합하다는 것을 Rego 규칙이 식별할 수 있는지 확인합니다:

{% code title="rules/CUSTOM-RULE-1/main_test.rego" %}
```
package rules

import data.lib
import data.lib.testing

test_CUSTOM_RULE_1 {
	# 규칙이 허용된 테스트 케이스가 포함된 배열
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]

	# 규칙이 거부된 경우를 포함하는 배열
	denied_test_cases := [{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied].tags"],
		"fixture": "denied.tf",
	}]

	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("CUSTOM-RULE-1", "./rules/CUSTOM-RULE-1/fixtures", test_cases)
}
```
{% endcode %}

## 논리 AND 포함한 예시

이전 예제를 확장하고 아래 두 가지 조건을 충족하는 모든 경우를 허용하는 규칙을 업데이트하십시오:

1. 리소스에 “소유자” 태그가 있음\
   **AND**
2. 리소스에 “설명” 태그가 있음

이 새 조건을 테스트하려면 `template` 명령을 사용하여 새 규칙 `CUSTOM-RULE-2`를 생성하고 아래 픽스처 파일을 작성하십시오:

{% code title="rules/CUSTOM-RULE-2/fixtures/denied.tf" %}
```
resource "aws_redshift_cluster" "denied" {
  cluster_identifier = "tf-redshift-cluster"
  node_type          = "dc1.large"
  tags = {
    owner = "team-123"
  }
}
```
{% endcode %}

여러 표현식을 결합하여 논리 `AND`를 나타낼 수 있습니다.

* `;` 연산자를 사용할 수 있습니다.
* 또는 표현식을 여러 줄에 나누어 적용함으로써 `;`(`AND`) 연산자를 생략할 수 있습니다.

{% hint style="info" %}
논리 AND는 [OPA 문서](https://www.openpolicyagent.org/docs/latest/#expressions-logical-and)에서도 다루고 있습니다.
{% endhint %}

{% code title="rules/CUSTOM-RULE-2/main.rego" %}
```
package rules

aws_redshift_cluster_tags_present(resource) {
    resource.tags.owner
    resource.tags.description
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    not aws_redshift_cluster_tags_present(resource)
    msg := {
        "publicId": "CUSTOM-RULE-2",
        "title": "Missing a description and an owner from the tag",
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

{% hint style="info" %}
Snyk은 늘 규칙이 올바른지 확인하기 위해 [규칙 업데이트 및 단위 테스트 실행](testing-a-rule.md)을 권장합니다.
{% endhint %}

이 규칙에 대한 테스트는 `CUSTOM-RULE-1`에 대한 테스트와 동일하지만, 테스트 이름 및 `testing.evaluate_test_cases` 함수에 전달되는 첫 번째 두 인수가 다릅니다.

{% code title="rules/CUSTOM-RULE-2/main_test.rego" %}
```
package rules

import data.lib
import data.lib.testing

test_CUSTOM_RULE_2 {
	# 규칙이 허용된 테스트 케이스가 포함된 배열
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]
	# 규칙이 거부된 경우를 포함하는 배열
	denied_test_cases := [{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied].tags"],
		"fixture": "denied.tf",
	}]
	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("CUSTOM-RULE-2", "./rules/CUSTOM-RULE-2/fixtures", test_cases)
}
```
{% endcode %}

## 논리 OR를 포함한 예시

위에서 작성한 규칙을 `NOT` 연산자와 `OR` 기능을 결합하여 다시 작성할 수도 있습니다.

두 가지 조건 중 하나라도 만족하지 않는 모든 경우를 거부하기 위해 새 규칙 `CUSTOM-RULE-3`에서 예제를 업데이트하십시오. 다음 두 경우 중 하나라도 없는 `aws_redshift_cluster` 리소스를 거부합니다.

1. “소유자” 태그가 없는 경우, 또는
2. “설명” 태그가 없는 경우

이를 위해 두 가지 새로운 케이스에 대해 각각의 픽스처 파일을 사용하십시오:

{% code title="rules/CUSTOM-RULE-3/fixtures/denied1.tf" %}
```
resource "aws_redshift_cluster" "denied1" {
  cluster_identifier = "tf-redshift-cluster"
  node_type          = "dc1.large"
  tags = {
    owner = "team-123@corp-domain.com"
  }
}
```
{% endcode %}

{% code title="rules/CUSTOM-RULE-3/fixtures/denied2.tg" %}
```
resource "aws_redshift_cluster" "denied2" {
  cluster_identifier = "tf-redshift-cluster"
  node_type          = "dc1.large"
  tags = {
    description = "description",
  }
}
```
{% endcode %}

Rego에서 논리 OR를 표현하기 위해 동일한 이름을 가진 여러 규칙이나 함수를 정의할 수 있습니다. 이는 [논리 OR에 대한 OPA 문서](https://www.openpolicyagent.org/docs/latest/#logical-or)에도 설명되어 있습니다.

우선 각 태그에 대해 `NOT`를 구현하는 함수를 추가하십시오. 그런 다음 이 함수를 리소스와 호출하십시오:

{% code title="rules/CUSTOM-RULE-3/main.rego" %}
```
package rules

aws_redshift_cluster_tags_missing(resource) {
    not resource.tags.owner
}

aws_redshift_cluster_tags_missing(resource) {
    not resource.tags.description
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    aws_redshift_cluster_tags_missing(resource)
    msg := {
        "publicId": "CUSTOM-RULE-3",
        "title": "Missing a description or an owner from the tag",
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

이렇게 하면 거부하는 모든 규칙이 성공적으로 반환됩니다.

{% hint style="info" %}
Snyk은 항상 [규칙 업데이트 및 단위 테스트 실행](./#test-a-custom-rule)을 통해 규칙이 올바른지 확인하는 것을 권장합니다.
{% endhint %}

이 규칙에 대한 테스트는 논리 OR가 예상대로 작동하는지 보여주기 위해 여러 테스트 케이스가 포함될 것입니다:

{% code title="rules/CUSTOM-RULE-3/main_test.rego" %}
```
package rules

import data.lib
import data.lib.testing

test_CUSTOM_RULE_3 {
	# 규칙이 허용된 테스트 케이스가 포함된 배열
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]
	# 규칙이 거부된 경우를 포함하는 배열
	denied_test_cases := [{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied1].tags"],
		"fixture": "denied1.tf",
	},{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied2].tags"],
		"fixture": "denied2.tf",
	}]
	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("CUSTOM-RULE-3", "./rules/CUSTOM-RULE-3/fixtures", test_cases)
}
```
{% endcode %}

## 문자열과 함께한 예시

더 확장하여 세 번째 조건을 추가하십시오. 다음 중 하나라도 없는 모든 리소스를 거부하십시오:

1. “소유자” 태그 , 또는
2. “설명” 태그, 또는
3. 소유자의 이메일이 “@corp-domain.com” 도메인에 속하지 않는 경우

{% code title="rules/CUSTOM-RULE-4/main.rego" %}
```
package rules

aws_redshift_cluster_tags_missing(resource) {
    not resource.tags.owner
}

aws_redshift_cluster_tags_missing(resource) {
    not resource.tags.description
}

aws_redshift_cluster_tags_missing(resource) {
    not endswith(resource.tags.owner, "@corp-domain.com")
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    aws_redshift_cluster_tags_missing(resource)
    msg := {
        "publicId": "CUSTOM-RULE-4",
        "title": "Missing a description and an owner from tag, or owner tag does not comply with email requirements",
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

{% hint style="info" %}
Snyk는 늘 규칙이 올바른지 확인하기 위해 [규칙 업데이트 및 단위 테스트 실행](testing-a-rule.md)을 권장합니다.
{% endhint %}

이 규칙에 대한 테스트는 이전 예제와 매우 유사할 것이며 자체 픽스처 파일이 필요합니다.

## XOR을 사용한 예시

여러 가지 복잡도를 추가하고 다음을 확인하려고 가정해보십시오:

* 태그 유형이 “사용자”이면 이메일 태그도 존재합니다.
* 그렇지 않을 경우, 다른 유형은 “서비스”라고 가정하며 이 서비스 설명이 있어야 합니다.
* 이 두 가지는 상호 배타적입니다; 첫 번째 조건이 해당되면 두 번째 조건은 해당되지 않으며 그 반대도 마찬가지입니다.

| 유형  | 이메일 | 서비스 설명 |
| --- | --- | ------ |
| 사용자 | YES | NO     |
| 서비스 | NO  | YES    |

이를 위해 코드를 리팩터링하여 `checkTags` 도우미 함수를 사용하십시오. 이 함수는 태그가 있고 두 조건 모두에 대해 OR를 사용하여 검사할 수 있습니다.

{% code title="rules/CUSTOM-RULE-5/main.rego" %}
```
package rules

checkTags(resource){
    resource.tags.type == "user"
    not resource.tags.email
}

checkTags(resource){
    resource.tags.type == "service"
    not resource.tags.serviceDescription
}

checkTags(resource){
    count(resource.tags) == 0
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
    checkTags(resource)   

    msg := {
        "publicId": "CUSTOM-RULE-5",
        "title": "Complex rule",
        "severity": "medium",
        "msg": sprintf("input.resource.aws_redshift_cluster[%v].tags", [name]),
        "issue": "",
        "impact": "",
        "remediation": "",
        "references": [],
    }
}
```
{% endcode %}

이를 XOR로 변환하려면 `else` 규칙을 사용할 수 있습니다:

{% code title="rules/CUSTOM-RULE-5/main.rego" %}
```
package rules

checkUserTag(resource){
    not resource.tags.email
}

checkUserTag(resource){
    resource.tags.serviceDescription
}

checkServiceTag(resource){
    not resource.tags.serviceDescription
}

checkServiceTag(resource){
    resource.tags.email
}

checkTags(resource){
	count(resource.tags) == 0
}

checkTags(resource) {
    resource.tags.type == "user"
    checkUserTag(resource)
} else {
    resource.tags.type == "service"
    checkServiceTag(resource)
}

deny[msg] {
    resource := input.resource.aws_redshift_cluster[name]
	checkTags(resource)
    msg := {
        "publicId": "CUSTOM-RULE-5",
        "title": "Missing the right tags from for a resource of type user or service",
        "severity": "medium",
        "msg": sprintf("input.resource.aws_redshift_cluster[%v].tags", [name]),
        "issue": "",
        "impact": "",
        "remediation": "",
        "references": [],
    }
}
```
{% endcode %}

이를 XOR로 변환하려면 `else`규칙을 사용하면 됩니다:

````
package rules
checkUserTag(resource){ 
    not resource.tags.email 
}
checkUserTag(resource){ 
    resource.tags.serviceDescription
}
checkServiceTag(resource){ 
    not resource.tags.serviceDescription 
}
checkServiceTag(resource){ 
    resource.tags.email 
}
checkTags(resource){ 
    count(resource.tags) == 0 
}
checkTags(resource) { 
    resource.tags.type == "user" 
    checkUserTag(resource) 
} else { 
    resource.tags.type == "service" 
    checkServiceTag(resource) 
}

deny[msg] { 
    resource := input.resource.aws_redshift_cluster[name] 
    checkTags(resource) 
    msg := { 
    "publicId": "CUSTOM-RULE-5", 
    "```korean "토큰", ]
    check_sensitive(keys, denylist) { 
    _ = keys[key] contains(key, denylist[_]) 
    }

deny[msg] { 
    input.kind == "ConfigMap" 
    input.data = keys 
    check_sensitive(keys, sensitive_denylist) 
    msg := { 
        "publicId": "CUSTOM-RULE-7", 
        "title": "ConfigMap exposes sensitive data", 
        "severity": "high", 
        "msg": "input.data", 
        "issue": "", 
        "impact": "", 
        "remediation": "", 
        "references": [], 
        } 
}

````

"pass", "secret", "key", "token"과 같은 부분 문자열이 포함된 키는 민감한 정보로 간주되므로 ConfigMap에 포함해서는 안 됩니다.

# 규칙 테스트

만약 [규칙 작성](writing-a-rule.md)에서 보여진대로 `template` 명령을 사용하여 규칙을 생성했다면, SDK 및 생성된 규칙과 함께 제공되는 테스팅 기능도 활용할 수 있습니다.

{% hint style="info" %}
SDK에서 생성된 테스트 기능을 수정하거나 직접 작성할 수도 있습니다. 하지만, 이 페이지의 지침은 해당되지 않을 것입니다.
{% endhint %}

이전 페이지를 기반으로 규칙을 작성했다면, SDK의 템플릿 기능이 생성한 `main_test.rego` 파일을 열고, `fixture` 필드를 `rules/MY_RULE/fixtures/` 폴더 내 파일 이름으로 구성하세요. 템플릿 기능은 지원하는 각 형식에 대해 한 파일씩 생성하고 모든 형식에 대해 테스트를 실행하도록 구성했지만, 필요에 따라 fixture 파일을 제거할 수 있습니다.

규칙에 대한 자원을 저장하기 위해 `rules/MY_RULE/fixtures` 아래에 fixture 파일을 생성하거나 수정하세요. 이 파일들은 아무 이름이나 가질 수 있으며, `denied.tf`와 `allowed.tf`와 같은 파일을 예로 들 수 있습니다:

{% hint style="info" %}
파일은 아무 이름이어도 되지만, 파일 확장자에 주의해야 합니다. 특히, Terraform Plan JSON 출력을 포함하는 fixture 파일에 대한 테스트를 작성하려는 경우에 주의해야 합니다. 파일의 이름이 `.json.tfplan` 확장자를 포함하도록해야 Snyk 테스팅 라이브러리가 일반 JSON과 Terraform Plan JSON 출력을 구별할 수 있습니다.
{% endhint %}

{% code title="rules/MY_RULE/fixtures/denied.tf" %}
```
resource "aws_redshift_cluster" "denied" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
}
```
{% endcode %}

{% code title="rules/MY_RULE/fixtures/allowed.tf" %}
```
resource "aws_redshift_cluster" "allowed" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  tags {
    owner = "snyk"
  }
}
```
{% endcode %}

테스트 케이스의 `want_msgs` 필드에는, 거부 규칙이 평가하고 반환할 것으로 예상되는 자원의 `msg` 필드를 추가해야 합니다. 예를 들어 `["input.resource.aws_redshift_cluster[denied].tags"]`와 같습니다.

{% hint style="info" %}
`want_msgs` 필드는 적절한 Rego 규칙의 계산된 `msg` 필드에 대응하는 하드코딩된 값 배열이어야 합니다.
{% endhint %}

{% code title="rules/MY_RULE/main_test.rego" %}
```
package rules

import data.lib
import data.lib.testing

test_MY_RULE {
	# 규칙이 허용되는 경우의 테스트 케이스를 포함하는 배열
	allowed_test_cases := [{
		"want_msgs": [],
		"fixture": "allowed.tf",
	}]
	# 규칙이 거부되는 경우의 테스트 케이스를 포함하는 배열
	denied_test_cases := [{
		"want_msgs": ["input.resource.aws_redshift_cluster[denied].tags"],
		"fixture": "denied.tf",
	}]
	test_cases := array.concat(allowed_test_cases, denied_test_cases)
	testing.evaluate_test_cases("MY_RULE", "./rules/MY_RULE/fixtures", test_cases)
}
```
{% endcode %}

모든 테스트를 실행하려면 다음 명령어를 실행하세요:

```
 snyk-iac-rules test
```

테스트가 성공하면, `rules/` 폴더에 세 가지 다른 규칙이 있다고 가정하면 다음과 같은 출력을 볼 수 있습니다:

```
성공: 3/3
```

그러나 이 중 하나라도 실패하면 다음과 같은 출력을 확인할 수 있습니다:

```
data.rules.test_MY_RULE: 실패 (1.12234ms)
실패: 2/3
```

`rules/` 폴더에 규칙이 하나 이상 있는 경우 다음 명령어를 실행하여 특정 테스트를 대상으로 할 수 있습니다:

```
snyk-iac-rules test --run test_MY_RULE
```

다음과 같은 출력이 됩니다:

```
Rego 테스트 케이스 실행 중...
data.rules.test_MY_RULE: 실패 (1.040468ms)
--------------------------------------------------------------------------------
실패: 1/1
```

출력에 대해 더 많은 세부 정보가 필요하면 `--explain notes` 옵션을 추가하세요:

```
 snyk-iac-rules test --run test_MY_RULE --explain notes
```

이렇게 하면 실패한 테스트의 디버그에 사용할 수 있는 추가 세부 정보가 출력됩니다.

{% hint style="info" %}
현재 폴더에 생성된 규칙 이상이 있는 경우, 테스트와 관련 없는 폴더 및 파일을 제외하기 위해 `--ignore` 옵션을 사용하는 것을 고려해보세요. `template` 명령을 사용했을 경우 `lib/` 및 `rules`를 제외하지 않도록 주의하세요. 이는 테스트를 가속화시키고 Rego가 Rego가 아닌 파일을 평가하려는 문제가 발생하는 것을 피할 수 있습니다.
{% endhint %}
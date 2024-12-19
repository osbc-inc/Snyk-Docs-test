# 입력 파일 구문 분석

Regol 코드를 작성할 때 입력 파일의 내부 표현을 이해하는 것은 어려울 수 있습니다. 규칙을 작성하는 방법을 배우면 알 수 있겠지만, 입력 값은 JSON과 유사한 객체이지만 입력 파일은 YAML, Terraform 또는 [테라폼 계획 JSON 출력](https://www.terraform.io/docs/internals/json-format.html)일 수도 있습니다. 이러한 것들이 JSON으로 어떻게 변환되는지 이해하는 데 도움이 되도록 Snyk는 `parse` 명령을 제공합니다.

Regol 코드를 작성하는 데 도움이 될 JSON으로 파싱된 파일을 얻으려면 IaC 파일이 필요합니다. [규칙 테스트](testing-a-rule.md)할 때도 사용할 수 있는 이 입력 파일은 기본적으로 파일을 JSON으로 파싱합니다.

## Terraform 파일 파싱

예를 들어 다음과 같은 Terraform 파일이 있습니다:

{% code title="example.tf" %}
```
resource "aws_redshift_cluster" "example" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
}
```
{% endcode %}

동등한 JSON 형식을 얻으려면 다음과 같이 파싱 명령을 실행하십시오:

```
snyk-iac-rules parse example.tf --format hcl2
```

이렇게 하면 JSON이 출력되고, 규칙을 작성할 때 참고할 수 있습니다:

```
{
	"resource": {
		"aws_redshift_cluster": {
			"example": {
				"cluster_identifier": "tf-redshift-cluster",
				"cluster_type": "single-node",
				"database_name": "mydb",
				"master_password": "Mustbe8characters",
				"master_username": "foo",
				"node_type": "dc1.large"
			}
		}
	}
}
```

Regol에서 `node_type` 필드에 액세스하는 것은 다음과 같이 보일 것입니다:

```
input.resource.aws_redshift_cluster.example.node_type
```

## YAML 파일 파싱

다른 예로는 Kubernetes 리소스를 정의하는 다음 YAML 파일이 있습니다:

{% code title="example.yaml" %}
```
apiVersion: v1
kind: Pod
metadata:
  name: example
spec:
  containers:
    - name: example
      image: example:latest
      securityContext:
        privileged: true 
```
{% endcode %}

동등한 JSON 형식을 얻으려면 다음과 같이 파싱 명령을 실행하십시오:

```
snyk-iac-rules parse example.yaml --format=yaml
```

이렇게 하면 JSON이 출력되고, 규칙을 작성할 때 참고할 수 있습니다:

```
{
	"apiVersion": "v1",
	"kind": "Pod",
	"metadata": {
		"name": "example"
	},
	"spec": {
		"containers": [
			{
				"image": "example:latest",
				"name": "example",
				"securityContext": {
					"privileged": true
				}
			}
		]
	}
}
```

Regol에서 `privileged` 필드에 액세스하는 것은 다음과 같이 보일 것입니다:

```
input.spec.containers[0].securityContext.privileged
```

## Terraform 계획 JSON 출력 파일 파싱

다른 예로는 `terraform show -json ./plan/example.json.tfplan` 명령으로 반환된 다음 Terraform 계획 JSON 출력 파일이 있습니다:

{% code title="example.json.tfplan" %}
```
{
  "format_version": "0.2",
  "terraform_version": "1.0.11",
  "planned_values": {
    "root_module": {
      "resources": [
        {
          "address": "aws_vpc.example",
          "mode": "managed",
          "type": "aws_vpc",
          "name": "example",
          "provider_name": "registry.terraform.io/hashicorp/aws",
          "schema_version": 1,
          "values": {
            "assign_generated_ipv6_cidr_block": false,
            "cidr_block": "10.0.0.0/16",
            "enable_dns_support": true,
            "instance_tenancy": "default",
            "tags": null
          },
          "sensitive_values": {
            "tags_all": {}
          }
        }
      ]
    }
  },
  "resource_changes": [
    {
      "address": "aws_vpc.example",
      "mode": "managed",
      "type": "aws_vpc",
      "name": "example",
      "provider_name": "registry.terraform.io/hashicorp/aws",
      "change": {
        "actions": [
          "create"
        ],
        "before": null,
        "after": {
          "assign_generated_ipv6_cidr_block": false,
          "cidr_block": "10.0.0.0/16",
          "enable_dns_support": true,
          "instance_tenancy": "default",
          "tags": null
        },
        "after_unknown": {
          "arn": true,
          "default_network_acl_id": true,
          "default_route_table_id": true,
          "default_security_group_id": true,
          "dhcp_options_id": true,
          "enable_classiclink": true,
          "enable_classiclink_dns_support": true,
          "enable_dns_hostnames": true,
          "id": true,
          "ipv6_association_id": true,
          "ipv6_cidr_block": true,
          "main_route_table_id": true,
          "owner_id": true,
          "tags_all": true
        },
        "before_sensitive": false,
        "after_sensitive": {
          "tags_all": {}
        }
      }
    }
  ],
  "configuration": {
    "provider_config": {
      "aws": {
        "name": "aws",
        "expressions": {
          "region": {
            "constant_value": "us-east-1"
          }
        }
      }
    },
    "root_module": {
      "resources": [
        {
          "address": "aws_vpc.example",
          "mode": "managed",
          "type": "aws_vpc",
          "name": "example",
          "provider_config_key": "aws",
          "expressions": {
            "cidr_block": {
              "constant_value": "10.0.0.0/16"
            }
          },
          "schema_version": 1
        }
      ]
    }
  }
}
```
{% endcode %}

동등한 JSON 형식을 얻으려면 다음과 같이 파싱 명령을 실행하십시오:

```
snyk-iac-rules parse example.json.tfplan --format=tf-plan
```

이렇게 하면 JSON이 출력되고, 규칙을 작성할 때 참고할 수 있습니다:

```
{
	"data": {},
	"resource": {
		"aws_vpc": {
			"example": {
				"assign_generated_ipv6_cidr_block": false,
				"cidr_block": "10.0.0.0/16",
				"enable_dns_support": true,
				"instance_tenancy": "default",
				"tags": null
			}
		}
	}
}
```

Regol에서 `tags` 필드에 액세스하는 것은 다음과 같이 보일 것입니다:

```
input.resource.aws_vpc.example.tags
```
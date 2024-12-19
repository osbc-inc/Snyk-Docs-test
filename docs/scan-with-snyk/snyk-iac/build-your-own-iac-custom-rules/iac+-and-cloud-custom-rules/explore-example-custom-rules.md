# 예제 사용자 지정 규칙 살펴보기

{% hint style="info" %}
Rego로 작성된 IaC to Cloud 사용자 정의 규칙의 여러 예시를 [snyk-labs/iac-to-cloud-example-custom-rules](https://github.com/snyk-labs/iac-to-cloud-example-custom-rules) 리포지토리에서 찾을 수 있습니다.
{% endhint %}

## 허용 목록의 예시

이 예시는 승인된 Amazon Machine Images (AMIs)의 목록을 생성하는 것을 보여줍니다.

{% code title="rules/APPROVED_AMIS/main.rego" %}
```
package rules.APPROVED_AMIS

import data.snyk

input_type := "tf"

metadata := {
	"id": "APPROVED-AMIS",
	"title": "EC2 instance is using an unapproved AMI",
	"severity": "high",
	"description": "We maintain a list of approved AMIs that fit our security and compliance needs. All DemoCorp EC2 instances must use one of these AMIs.",
	"product": ["iac", "cloud"],
	"platform": ["aws"],
}

approved_amis := {
	# us-east-1
	"ami-00c39f71452c08778",
	"ami-02f97949d306b597a",
	# us-east-2
	"ami-04581fbf744a7d11f",
	"ami-0533def491c57d991",
}

test_approved_amis {
	approved_amis["ami-0533def491c57d991"]
}

instances := snyk.resources("aws_instance")

deny[info] {
	instance := instances[_]
	not approved_amis[instance.ami]
	info := {"resource": instance}
}

resources[info] {
	instance := instances[_]
	info := {"resource": instance}
}
```
{% endcode %}

## 올바른 IAM 구성을 확인하는 규칙 예시

{% code title="rules/S3_ACL/main.rego" %}
```
package rules.S3_ACL

import data.snyk

input_type := "tf"

metadata := {
	"id": "S3_BUCKET_ACL",
	"severity": "critical",
	"title": "All ACLs should be private",
	"description": "Checking S3 Buckets for Private ACLs using the new terraform format.",
	"product": [
		"iac",
		"cloud",
	],
}

buckets := snyk.resources("aws_s3_bucket")

deny[info] {
	bucket := buckets[_]
	acls := snyk.relates(bucket, "aws_s3_bucket.aws_s3_bucket_acl")
	acl := acls[_]
	acl.acl != "private"
	info := {"primary_resource": bucket}
}

resources[info] {
	bucket := buckets[_]
	info := {"primary_resource": bucket}
}

resources[info] {
	bucket := buckets[_]
	acls := snyk.relates(bucket, "aws_s3_bucket.aws_s3_bucket_acl")
	info := {
		"primary_resource": bucket,
		"resource": acls[_],
	}
}
```
{% endcode %}

## GitHub Terraform Provider를 기반으로 한 규칙 예시

GitHub나 Snowflake와 같은 어떤 유형의 terraform 공급자나 리소스에 대한 규칙을 작성할 수 있습니다.

{% code title="rules/GITHUB_DEFAULT_BRANCH_DELETION_PROTECTION/main.rego" %}
```
package rules.GITHUB_DEFAULT_BRANCH_DELETION_PROTECTION

import data.snyk

resource_type := "MULTIPLE"

input_type := "tf"

metadata := {
	"id": "GITHUB-DEFAULT-BRANCH-DELETION-PROTECTION",
	"title": "Default branch deletion protection not enabled",
	"severity": "high",
	"description": "The history of the default branch is not protected against deletion for this repository.",
	"product": ["iac"],
}

repos := snyk.resources("github_repository")

is_valid(repo) {
	branch_protection := snyk.relates(repo, "github_repository.branch_protection")[_]
	not branch_protection.allows_deletions
}

deny[info] {
	repo := repos[_]
	not is_valid(repo)
	info := {"resource": repo}
}

resources[info] {
	repo := repos[_]
	info := {"resource": repo}
}

resources[info] {
	repo := repos[_]
	branch_protection := snyk.relates(repo, "github_repository.branch_protection")[_]
	info := {"primary_resource": repo, "resource": branch_protection}
}
```
{% endcode %}

## 리소스 태그 강제 예시

{% code title="rules/REQUIRED_S3_BUCKET_TAGS/main.rego" %}
```
package rules.REQUIRED_S3_BUCKET_TAGS

import data.snyk

input_type := "tf"

metadata := {
	"id": "REQUIRED_S3_BUCKET_TAGS",
	"severity": "high",
	"title": "S3 Bucket Tags",
	"description": "All S3 Buckets must be tagged properly - they need to have an owner tag, and a classification tag with the proper values.",
	"product": [
		"iac",
		"cloud",
	],
}

buckets := snyk.resources("aws_s3_bucket")

owners := {
	"devteam1",
	"devteam2",
	"devteam3",
	"devteam4"
}

classifications := {
	"public",
	"internal",
	"confidential"
}

properly_tagged(bucket) {
	owners[bucket.tags.owner]
	classifications[bucket.tags.classification]
}

deny[info] {
	bucket := buckets[_]
	not properly_tagged(bucket)
	info := {"resource": bucket}
}

resources[info] {
	bucket := buckets[_]
	info := {"resource": bucket}
}
```
{% endcode %}
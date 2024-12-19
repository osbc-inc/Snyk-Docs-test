# AWS 제공자 설정

## AWS 제공자의 인증

`iac describe`를 사용하려면 AWS에 인증된 요청을 생성하기 위한 자격 증명을 설정해야 합니다. AWS CLI에서와 마찬가지로, 사용자 환경 변수로 선언되거나 로컬 AWS 구성 파일에 정의된 [자격 증명 및 구성](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) 설정을 사용합니다.

`iac describe` 명령은 [named profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html)을 지원합니다. CLI는 기본적으로 `default`라는 이름의 프로필에서 찾은 설정을 사용합니다. `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_PROFILE` 등과 같은 환경 변수를 선언하여 개별 설정을 재정의할 수 있습니다.

권한 부여 도구로 [IAM 역할](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html)을 사용하는 경우, 이는 좋은 실천 방법으로 간주됩니다. 이 경우 `~/.aws/config` 파일에서 역할을 위한 프로필을 정의함으로써 `iac describe`를 여전히 사용할 수 있습니다.

```
[profile snykrole]
role_arn = arn:aws:iam::123456789012:role/<NAMEOFTHEROLE>
source_profile = user # 역할에 가정할 프로필
region = eu-west-3
```

이제 프로필 설정을 재정의하여 `iac describe`를 사용할 수 있습니다.

```
$ AWS_PROFILE=snykrole snyk iac describe
```

## S3 백엔드에서 상태를 읽기 위한 사용자 지정 자격 증명

S3에서 상태를 읽는 데 다른 AWS 자격 증명 세트를 사용하려는 경우, 각 구체적인 AWS 환경 변수를 `DCTL_S3_` 접두사로 재정의할 수 있습니다. 이를 통해 인프라와 다른 지역에서 상태를 읽을 수 있도록 선택할 수 있습니다. 실제 인프라의 리소스를 읽기 위해서는 보통 AWS 자격 증명을 사용해야 합니다.

```
# S3 백엔드에서 상태를 읽기 위해 전용 AWS named profile(또는 다른 AWS 환경 변수) 내보내기
$ export DCTL_S3_PROFILE="s3reader"
# 보통의 AWS named profile 내보내기
$ export AWS_PROFILE="snykrole"
$ snyk iac describe --from="tfstate+s3://mybucket/terraform.tfstate"

# S3 버킷에 인증하는 데 특정 지역을 사용할 수도 있습니다
$ DCTL_S3_REGION=us-east-1 snyk iac describe --from="tfstate+s3://mybucket/terraform.tfstate"
```

## Terraform 사용자 정의 역할

다음 코드는 HCL로 작성된 `iac describe`을 실행할 수 있는 사용자 정의 역할을 나타냅니다.

```
data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    effect = "Allow"
    actions   = ["sts:AssumeRole"]
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }
  }
}

data "aws_iam_policy_document" "policy" {
  statement {
    effect = "Allow"
    actions   = [
      "apigateway:GET",
      "cloudformation:DescribeStacks",
      ...
      "autoscaling:DescribeLaunchConfigurations"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_role" "snyk_assume_role" {
  name = "snyk_assume_role"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

resource "aws_iam_role_policy" "snyk_policy" {
  name = "snyk_policy"
  role = aws_iam_role.snyk_assume_role.id
  policy = data.aws_iam_policy_document.policy.json
}
```

## CloudFormation 템플릿

이 CloudFormation 템플릿을 배포하여 이 페이지 이전 섹션의 인증 가이드에 따라 사용할 수 있는 제한된 권한 역할을 생성합니다.

[![Launch Stack](https://cdn.rawgit.com/buildkite/cloudformation-launch-stack-button-svg/master/launch-stack.svg)](https://console.aws.amazon.com/cloudformation/home?#/stacks/quickcreate?stackName=driftctl-stack&templateURL=https://driftctl-cfn-templates.s3.eu-west-3.amazonaws.com/driftctl-role.yml)

스택을 배포하면 다음 정책을 IAM 사용자에 연결하십시오. 이렇게 하면 사용자가 지정된 역할만 가정할 수 있습니다. 사용자에게 역할 가정 권한을 부여하는 자세한 내용은 [AWS Identity and Access Management 사용자 가이드](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_permissions-to-switch.html)를 참조하십시오.

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Resource": "arn:aws:iam::<IDOFYOURACCOUNT>:role/snyk_assume_role"
    }
  ]
}
```

이 템플릿을 AWS 계정에서 시작했기 때문에 Snyk 측에서 CloudFormation 템플릿을 자동으로 업데이트할 방법은 없습니다. 따라서 최신 Snyk 역할을 사용하기 위해 템플릿을 직접 업데이트해야 합니다.

### AWS 콘솔을 사용하여 CloudFormation 템플릿 업데이트

* [AWS CloudFormation 콘솔](https://console.aws.amazon.com/cloudformation)에서 스택 목록에서 **snyk 스택**을 선택합니다.
* 스택 상세 정보 패널에서 **업데이트**를 선택합니다.
* **현재 템플릿 교체**를 선택하고 Snyk **Amazon S3 URL** `https://driftctl-cfn-templates.s3.eu-west-3.amazonaws.com/driftctl-role.yml`을 지정한 후 **다음**을 클릭합니다.
* **스택 세부정보 지정** 및 **스택 옵션 구성** 페이지에서 **다음**을 클릭합니다.
* **변경 세트 미리 보기** 섹션에서 AWS CloudFormation이 변경 사항을 적용할 것인지 확인합니다.
* Snyk 템플릿에 IAM 리소스가 하나 포함되어 있으므로 **이 템플릿이 IAM 리소스를 만들 수 있다는 것을 인증합니다**를 선택합니다.
* 완료하려면 **스택 업데이트**를 클릭합니다.

### AWS CLI를 사용하여 CloudFormation 템플릿 업데이트

다음 명령을 사용합니다:

```
$ aws cloudformation update-stack --stack-name SNYK_STACK_NAME --template-url https://driftctl-cfn-templates.s3.eu-west-3.amazonaws.com/snyk-role.yml --capabilities CAPABILITY_NAMED_IAM
```

## 최소 권한 정책

`iac describe` 명령은 클라우드 제공자 계정에 액세스해야 하므로 사용자 대신 리소스를 나열할 수 있어야 합니다.

AWS 설명서에서 권장하는 대로, 다음 정책은 필요한 권한만 부여합니다.

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Resource": "*",
            "Action": [
                "apigateway:GET",
                "cloudformation:DescribeStacks",
                "cloudformation:GetTemplate",
                "cloudfront:GetDistribution",
                ...
                "autoscaling:DescribeLaunchConfigurations"
            ]
        }
    ]
}
```
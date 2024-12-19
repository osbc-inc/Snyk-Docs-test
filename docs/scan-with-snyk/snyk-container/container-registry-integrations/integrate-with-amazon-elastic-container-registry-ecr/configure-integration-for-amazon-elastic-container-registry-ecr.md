# Amazon Elastic Container Registry (ECR) 통합 구성

{% hint style="warning" %}
ECR 통합을 설정할 때는 us-east-2 지역이 활성화되어 있는지 확인하세요. 이는 STS(보안 토큰 서비스)가 올바르게 작동하기 위해 필요합니다. 자세한 내용은 [관련 지원 문서](https://support.snyk.io/s/article/Connecting-to-ECR-Integration-gives-error-Could-not-connect-to-ECR-Please-ensure-your-credentials-are-correctly-configured)를 참조하십시오.
{% endhint %}

이 페이지에서는 한 Amazon ECR 레지스트리와 Snyk 조직 간의 통합을 활성화하고 이미지 보안을 관리하는 방법을 설명합니다. 여러 레지스트리와 통합하려면 각각에 대해 고유한 조직을 만들어야 합니다.

## **ECR를 위한 자동화된 통합 프로세스**

[AWS Quick Start](https://github.com/aws-quickstart/quickstart-snyk-security)를 사용하여 Snyk의 Amazon ECR 통합을 원 클릭 배포로 활성화하는 방식으로 크로스 계정 액세스를 수립할 수 있습니다. 이렇게 하면 수동 구성이 필요 없이 통합할 수 있습니다.

<figure><img src="../../../../.gitbook/assets/configure_integration_Amazon_ecr.png" alt=""><figcaption><p>AWS ECR 및 Snyk 통합 크로스 계정 IAM 역할</p></figcaption></figure>

통합을 완료하려면 Snyk **조직 ID** 및 AWS IAM [역할 ARN](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns)이 있어야 합니다. 역할 ARN은 AWS CloudFormation 콘솔의 출력 탭에 제공됩니다.

## **ECR를 위한 수동 통합 프로세스**

통합을 활성화하려면 먼저 읽기 전용 AWS Identity and Access Management (IAM) 역할을 생성해야 합니다. 이 역할은 Snyk의 각 조직에 대해 레지스트리의 모든 리포지토리에 대한 읽기 전용 액세스를 위임하며 허용된 Snyk가 할당한 조직 ID 목록을 지정합니다.

IAM 역할을 생성한 후 추가 조직을 통합하려면 필요한 추가 조직 ID를 추가할 수 있습니다.
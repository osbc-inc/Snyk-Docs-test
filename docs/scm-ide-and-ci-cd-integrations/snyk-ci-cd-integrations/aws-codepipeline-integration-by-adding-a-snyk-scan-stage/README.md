# AWS CodePipeline 통합을 위한 Snyk 스캔 단계 추가

{% hint style="warning" %}
**AWS CodePipeline을 위한 Snyk 통합이 중단될 예정입니다.**

**조치가 필요함**

Snyk 서비스 및 고객의 보안을 보호하기 위해 Snyk는 **AWS CodePipeline**과의 통합을 폐기하기 시작했습니다. 중단을 최소화하기 위해 Snyk는 동일한 사용 사례와 기능을 지원하는 대체로 **AWS CodeBuild**와 Snyk CLI를 사용하도록 전환할 것을 권장합니다.&#x20;

**이관 일정**

2024년 **10월 30일**부터는 새로운 또는 기존의 파이프라인에 대해 Snyk 플러그인을 추가하거나 수정할 수 없게 됩니다. 기존의 파이프라인은 6개월 동안 그대로 작동하게 되지만, Snyk는 가능한 빨리 새로운 프로세스로 이관할 것을 권장합니다. CI/CD 작업이 중단되지 않도록하려면 **2025년 4월 30일** 이전에 Snyk CLI로 이전해야 합니다. Snyk CLI를 AWS CodeBuild와 함께 사용하기 위한 단계에 대해서는 이 [마이그레이션 가이드](migrating-to-aws-codebuild.md)를 참조하십시오.

Snyk는 AWS CodeBuild와 Snyk CLI가 요구 사항을 충족할 것이라고 확신합니다.
{% endhint %}

Snyk는 AWS CodePipeline과 신속한 연동을 통해 애플리케이션을 오픈 소스 보안 취약점을 스캔하고 지속적인 전달 서비스로 안전한 애플리케이션을 제공하는 데 도움을 줍니다. 이 통합을 통해 CodePipeline 사용자는 보안을 빌드, 테스트 및 배포 단계의 자동화된 부분으로 만들 수 있습니다.

{% hint style="info" %}
Snyk 통합은 AWS `sa-east-1` | `ca-central-1` | `ap-southeast-1` | `ap-southeast-2` | `ap-south-1` | `ap-northeast-2` | `ap-northeast-1` | `eu-west-3` | `eu-west-1` | `eu-north-1` | `us-east-1` | `us-east-2` | `us-west-1` | `us-west-2` | `eu-west-2` | `eu-central-1` 지역에서 사용할 수 있습니다. Snyk는 추가 지역으로 확대하고 있는 중입니다.
{% endhint %}

설치 및 사용 세부 정보는 다음 페이지를 참조하십시오:

* [AWS CodePipeline 언어 지원](language-support-for-aws-codepipeline.md)
* [AWS CodePipeline 설정 요구 사항](setup-requirements-for-aws-codepipeline.md)
* [AWS CodePipeline CodeBuild 단계 예시](aws-code-pipeline-codebuild-step-example.md)
* [AWS CodePipeline 설정 단계](setup-steps-for-aws-codepipeline-integration.md)
* [AWS CodePipeline 스캔 결과 보기](view-aws-codepipeline-scan-results.md)
* [AWS CodePipeline 테스트 보고서 결과](aws-codepipeline-test-report-details.md)

{% hint style="info" %}
AWS CodePipeline과의 Snyk 통합에는 설정의 일부로 UI 기반 인증 단계가 필요합니다. 이는 CloudFormation이나 Terraform과 같은 비대화식 설정 방법에 의한 자동화와 호환되지 않습니다.
{% endhint %}  
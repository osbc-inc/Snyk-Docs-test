# AWS 코드파이프라인에 대한 언어 지원

{% hint style="warning" %}
**AWS CodePipeline을 위한 Snyk 통합이 중단될 예정입니다**

\
**조치 필요**

저희의 서비스 및 고객의 보안을 보호하기 위해, Snyk은 **AWS CodePipeline**과의 통합을 폐기하기 시작했습니다. 중단을 최소화하기 위해, **AWS CodeBuild**와 Snyk CLI를 대체로 사용하고, 동일한 사용 사례와 기능을 지원할 것을 권장합니다.

**이관 일정**

**2024년 10월 30일**부로, 새로운 파이프라인에 Snyk 플러그인을 추가하거나 수정할 수 없게 될 것입니다. 기존 파이프라인은 6개월 동안 그대로 작동할 것이지만, 가능한 빨리 새 프로세스로 이관하는 것이 좋습니다. **2025년 4월 30일** 이전에 Snyk CLI로 이관하여 CI/CD 워크플로를 방해하지 않도록해야 합니다. AWS CodeBuild에서 Snyk CLI를 사용하는 방법은 이 [이관 안내서](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/aws-codepipeline-integration-by-adding-a-snyk-scan-stage/migrating-to-aws-codebuild)를 참조하세요.\\

AWS CodeBuild와 Snyk CLI가 요구사항을 충족할 것으로 확신합니다.
{% endhint %}

AWS CodePipeline에서 다음 언어를 지원하는 Snyk 통합:

* JavaScript
* Java
* .NET
* Python
* Ruby
* PHP
* Scala
* Swift/Objective-C
* Go

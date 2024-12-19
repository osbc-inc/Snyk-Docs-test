# AWS CodePipeline CodeBuild 단계 예시

{% hint style="warning" %}
**AWS CodePipeline을 위한 Snyk 통합이 중단될 예정입니다**

\
**조치가 필요합니다**

저희의 서비스와 고객의 보안을 보호하기 위해, Snyk은 **AWS CodePipeline**과의 통합을 폐기하기 시작했습니다. 중단을 최소화하기 위해, **AWS CodeBuild** 및 Snyk CLI를 사용하도록 전환하고, 동일한 사용 사례와 기능을 지원하는 대안을 권장합니다.&#x20;

**이주 일정**

2024년 10월 30일부로, 새로운 파이프라인에 Snyk 플러그인을 추가하거나 수정할 수 없게 됩니다. 기존 파이프라인은 6개월 동안 현재 상태로 계속 작동되지만, 가능한 빨리 새로운 프로세스로 전환하는 것을 권장합니다. CI/CD 워크플로우를 방해하지 않으려면 **2025년 4월 30일** 이전에 Snyk CLI로 전환해야 합니다. **AWS CodeBuild**에서 Snyk CLI를 사용하기 위해 이 [이주 가이드](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/aws-codepipeline-integration-by-adding-a-snyk-scan-stage/migr...를 참조하세요.\

AWS CodeBuild와 Snyk CLI가 귀하의 요구 사항을 충족시킬 것이라고 확신합니다.&#x20;
{% endhint %}

스캔의 입력 아티팩트는 빌드의 출력 아티팩트로 제공되어야 함을 주의하세요. 구성에서 보여진대로 진행됩니다.

## Javascript CodeBuild의 예시(`buildspec.yml`)

```yaml
version: 0.2
phases:
  build:
    commands:
      - npm install
artifacts:
  files:
    - '**/*'
```

## Maven CodeBuild의 예시(`buildspec.yml`)

```yaml
version: 0.2
phases:
  build:
    commands:
      - mvn install
artifacts:
  files:
    - '**/*'
```
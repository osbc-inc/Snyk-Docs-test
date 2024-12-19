# AWS CodePipeline을 위한 요구 사항 설정

{% hint style="warning" %}
**AWS CodePipeline을 위한 Snyk 통합이 중단될 예정입니다**

\
**동작이 필요함**

저희의 서비스와 고객들의 보안을 보호하기 위해 Snyk은 **AWS CodePipeline**과의 통합을 폐기하기 시작했습니다. 혼란을 최소화하기 위해 **AWS CodeBuild**와 Snyk CLI를 대안으로 사용하도록 전환하는 것을 권장합니다. 이는 동일한 사용 사례와 기능을 지원할 것입니다.&#x20;



**마이그레이션 타임라인**

**2024년 10월 30일**부터는 Snyk 플러그인을 새로운 또는 기존의 파이프라인에 추가 또는 수정할 수 없게 될 것입니다. 기존의 파이프라인은 6개월 동안 그대로 작동될 것이지만, 가능한 빨리 새로운 프로세스로 마이그레이션하는 것을 권장합니다. CI/CD 워크플로우가 방해되지 않도록 **2025년 4월 30일** 이전에 Snyk CLI로 전환해야 합니다. [마이그레이션 가이드](https://docs.snyk.io/scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/aws-codepipeline-integration-by-adding-a-snyk-scan-stage/migrating-to-aws-codebuild)에서 Snyk CLI를 AWS CodeBuild와 함께 사용하는 단계를 참조하십시오.\



AWS CodeBuild와 Snyk CLI가 요구 사항을 충족시킬 것으로 확신합니다.&#x20;
{% endhint %}



CodePipeline에서 스캔 전에 프로젝트를 빌드해야 하는지 확인하십시오. 프로젝트를 빌드해야 하는 경우 Snyk 단계 앞에 CodeBuild 단계를 추가해야 합니다.

|      언어     | 프로젝트 유형 | 빌드 필요 여부 |                                            노트                                           |
| :---------------: | :----------: | -------------- | :----------------------------------------------------------------------------------------: |
|     Javascript    |      npm     | 아니오\*           |   `package-lock.json` 파일이 없는 경우에만 필요함; 생성하기 위해 npm install 실행  |
|     Javascript    |     Yarn     | 아니오\*           |      `yarn.lock` 파일이 없는 경우에만 필요함; 생성하기 위해 yarn install 실행      |
|        Java       |     Maven    | 예            |                              테스트 전에 `mvn install` 실행                              |
|        Java       |    Gradle    | 아니오             |                                                                                            |
|        .NET       |     Nuget    | 아니오\*           |                  `packages.config` 파일이 없는 경우에만 필요함                  |
|       Python      |      Pip     | 아니오\*           |     언어 설정 매개변수가 있는 Snyk 설정 파일이 없는 경우에만 필요함     |
|       Python      |   Setup.py   | 예            |                            테스트 전에 `pip install -e .` 실행                           |
|       Python      |    Poetry    | 아니오\*           |     `poetry.lock` 파일이 없는 경우에만 필요함; 생성하기 위해 `poetry lock` 실행    |
|        Ruby       |    Bundler   | 아니오\*           |   `Gemfile.lock` 파일이 없는 경우에만 필요함; 생성하기 위해 `bundle install` 실행  |
|        PHP        |   Composer   | 아니오\*           | `composer.lock` 파일이 없는 경우에만 필요함; 생성하기 위해 `composer install` 실행 |
|       Scala       |      SBT     | 아니오             |                                                                                            |
|         Go        |  모듈 Go  | 아니오             |                                                                                            |
| Swift/Objective-C |   Cocoapods  | 아니오\*           |     `Podfile.lock` 파일이 없는 경우에만 필요함; 생성하기 위해 pod install 실행     |
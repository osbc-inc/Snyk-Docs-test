# Snyk 작업을 사용하여 애플리케이션을 테스트하는 간단한 예제

다음은 Snyk 작업을 사용하여 애플리케이션의 오픈 소스 종속성(SCA)을 테스트하는 간단한 예제입니다.

```yaml
- task: SnykSecurityScan@0
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'app'
    monitorWhen: 'always'
    failOnIssues: true
```
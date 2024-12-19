# Snyk 작업을 실행하는 간단한 예제

다음은 Snyk 작업을 실행하는 간단한 예제입니다.

```
- task: SnykSecurityScan@1
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'code'
    codeSeverityThreshold: 'medium'
    failOnIssues: true
```
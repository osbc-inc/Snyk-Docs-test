# Snyk 작업을 사용하여 컨테이너 이미지 테스트하는 간단한 예제

다음은 컨테이너 이미지를 테스트하는 Snyk 작업의 간단한 예제입니다.

```yaml
- task: SnykSecurityScan@1
  inputs:
    serviceConnectionEndpoint: 'snykToken'
    testType: 'container'
    dockerImageName: 'goof'
    dockerfilePath: 'Dockerfile'
    monitorWhen: 'always'
    failOnIssues: true
```
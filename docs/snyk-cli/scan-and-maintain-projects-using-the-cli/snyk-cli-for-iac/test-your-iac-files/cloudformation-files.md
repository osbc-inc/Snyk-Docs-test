# CloudFormation 파일

Snyk Infrastructure as Code를 사용하면 CLI를 통해 CloudFormation 파일을 테스트할 수 있습니다. YAML 및 JSON 형식의 파일을 모두 테스트할 수 있습니다. 또한 AWS CDK 애플리케이션을 스캔할 수도 있습니다; [AWS CDK 파일](aws-cdk-files.md)을 참조하십시오.

다음 명령어를 사용하여 지정된 파일에서 문제를 테스트할 수 있습니다:

```
snyk iac test
```

예시:

```
snyk iac test deploy.yaml
```

또한 파일 이름을 연이어 붙여 여러 파일을 지정할 수도 있습니다. 예를 들어:

```
snyk iac test file-1.yaml file-2.yaml
```

# AWS CDK 파일

Snyk Infrastructure as Code를 사용하여 CLI로 구성 파일을 테스트할 수 있습니다. CDK CLI를 사용하여 CloudFormation 파일을 생성하여 [Amazon Web Services Cloud Development Kit (AWS CDK)](https://aws.amazon.com/cdk/)를 Snyk CLI로 스캔할 수 있습니다.

다음 단계를 따라 CDK 애플리케이션을 스캔합니다:

**CDK 애플리케이션을 스캔하려는 애플리케이션 스택이 있는 디렉토리로** 이동합니다.

**CloudFormation 파일을 생성**합니다.

```
cdk synth
```

이는 터미널에서 YAML 출력으로 표시되며, `cdk.out` 디렉토리에 JSON 파일이 생성됩니다.

**다음 Snyk IaC CLI 명령을 사용하여** JSON 파일을 스캔합니다. 스캔하려는 애플리케이션의 이름으로 `cdk.out/*.json`을(를) 대체하십시오.

```
snyk iac test cdk.out/*.json
```

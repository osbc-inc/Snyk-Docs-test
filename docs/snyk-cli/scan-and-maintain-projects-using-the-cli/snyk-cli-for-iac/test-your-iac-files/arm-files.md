# ARM 파일

Snyk의 **인프라스트럭처 as 코드(Infrastructure as Code)**를 사용하면 CLI를 통해 구성 파일을 테스트할 수 있습니다.

Azure 리소스 관리자(ARM)를 위한 **인프라스트럭처 as 코드(Infrastructure as Code)**는 JSON 형식 파일을 스캔할 수 있습니다. Bicep 형식 파일을 스캔하기 위해서는 Bicep CLI를 사용하여 구성 파일을 JSON으로 변환해야 합니다.

## 특정 JSON 파일에 문제를 테스트

다음 **Snyk** CLI 명령어를 입력하세요:

```
snyk iac test deploy.json
```

또한 파일 이름을 연이어 붙여 여러 파일을 지정할 수도 있습니다. 예를 들어:

```
snyk iac test file-1.json file-2.json
```

## 특정 Bicep 파일에 문제를 테스트

[Bicep CLI를 설치](https://docs.microsoft.com/en-us/azure/azure-resource-manager/bicep/install)해야 합니다.

Bicep CLI를 설치한 후, Bicep 파일이 있는 디렉토리로 **이동**하고 해당 Bicep 파일을 JSON으로 변환해야 합니다. 다음을 입력하세요:

```
az bicep build -f deploy.bicep
```

그런 다음 새로 생성된 JSON 파일을 다른 파일과 마찬가지로 스캔할 수 있습니다. 다음 **Snyk** CLI 명령어를 사용하세요:

```
snyk iac test deploy.json
```  
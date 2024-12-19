# Serverless 파일

Snyk의 인프라스트럭처로 코드를 사용하면 CLI를 통해 구성 파일을 테스트할 수 있습니다.

서버리스 프레임워크용 Snyk 인프라스트럭처 코드는 클라우드포메이션 JSON 형식 파일의 패키지 출력을 스캔하는 것을 지원합니다.

이 페이지에서 설명된 대로 **지정된 서버리스 파일에서 문제를 테스트**할 수 있습니다.

[서버리스 CLI를 설치](https://www.serverless.com/framework/docs/getting-started)했는지 확인하세요.

서버리스 CLI를 설치한 후, Serverless 파일이 들어 있는 디렉토리로 **이동**하여 다음을 입력하여 Serverless 아티팩트를 **생성**합니다.

```
serverless package --package serverless-artifacts
```

그런 다음 취약점을 테스트하기 위해 다음 Snyk CLI 명령을 입력할 수 있습니다.

```
snyk iac test
```

또한 다음 명령을 사용하여 결과를 Snyk UI에 업로드할 수도 있습니다.

```
snyk iac test --report
```
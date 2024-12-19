# npm 패키지 사용 전에 테스트하기

`npm` 패키지를 설치하기 전에 **알려진 취약점이 있는지 확인하려면 `snyk test`를 사용할 수 있습니다**. 패키지 이름을 사용하여 패키지의 최신 버전을 테스트할 수 있습니다.

`snyk test ionic@1.6.5`

또한 `snyk test module[@semver-range]`를 사용하여 특정 버전이나 범위를 지정할 수 있습니다.

[CLI로 시작하기](../getting-started-with-the-snyk-cli.md)도 참조하세요.
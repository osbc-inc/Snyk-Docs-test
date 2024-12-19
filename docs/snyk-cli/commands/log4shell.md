# Log4shell

## 사용법

`snyk log4shell`

## 설명

`snyk log4shell` 명령어는 Log4J 라이브러리의 Log4Shell 취약점 [CVE-2021-44228](https://security.snyk.io/vuln/SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720)에 영향을 받는 흔적을 찾습니다.

이 명령은 `pom.xml` 또는 `build.gradle`과 같은 매니페스트 파일에 선언되지 않은 Log4J 라이브러리의 흔적을 찾습니다.

## 관리되는 프로젝트

Java 프로젝트에서 패키지 관리자 매니페스트 파일을 사용하여 Log4Shell 취약점을 테스트하려면 `snyk test` 명령어를 사용하세요. [test command help](test.md) (`snyk test --help`) 및 [Java와 Kotlin](https://docs.snyk.io/scan-using-snyk/supported-languages-and-frameworks/java-and-kotlin)을 참조하세요.

관리되지 않은 파일을 테스트하려면 `snyk test --scan-all-unmanaged`를 사용하세요.

[test command help](test.md)의 Maven 옵션 섹션을 참조하세요; `snyk test --help`

## 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공 (스캔 완료), Log4Shell을 찾지 못함\
**1**: 조치 필요 (스캔 완료), Log4Shell 발견\
**2**: 실패, 명령을 다시 실행하십시오. 디버그 로그를 출력하려면 `-d`를 사용하세요.

## 디버깅

디버그 로그를 출력하려면 `-d` 옵션을 사용하세요.
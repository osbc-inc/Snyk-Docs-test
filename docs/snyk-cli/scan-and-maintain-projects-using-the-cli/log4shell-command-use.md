# Log4shell 명령 사용

## 소개

`snyk log4shell`은 **Log4Shell** 취약성에 영향을 받는 **log4j** 라이브러리의 추적을 찾아주는 Snyk CLI 명령어입니다. 이 명령어는 이 라이브러리가 manifest 파일(pom.xml 또는 build.gradle과 같은)에 선언되어 있지 않아도 작동합니다.

이 명령은 빌드된 프로젝트 및 타사 응용 프로그램을 테스트하며, `snyk test` 및 `snyk test --scan-all-unmanaged` 명령에 보충적인 역할을 합니다.

{% hint style="info" %}
Snyk [VulnDB 항목](https://security.snyk.io/vuln/SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720)에서 Log4Shell 취약성에 대해 자세히 알아보세요.
{% endhint %}

자바 프로젝트를 패키지 관리자 manifest 파일을 사용하여 테스트하려면 CLI `test` 명령어의 [Maven 프로젝트 옵션](../commands/test.md#options-for-maven-projects) 및 [Gradle 프로젝트 옵션](../commands/test.md#options-for-gradle-projects)을 참조하세요.

`snyk test --scan-all-unmanaged`에 대해 자세히 알아보려면 [CLI 참조의 Maven 옵션 섹션](../cli-commands-and-options-summary.md#option-for-maven-projects)을 확인하세요.

## 사용법

`snyk log4shell`을 사용하여 Java 프로젝트를 스캔하여 프로젝트가 다음을 포함하는지 확인합니다:

* 취약한 버전의 log4j가 포함된 `.jar` 또는 `.war` 파일.
* 취약한 log4j 라이브러리의 알려진 파일이 포함된 파일. 이러한 발견은 전체 log4j가 포함될 수 있음을 나타냅니다.

## 실행 방법

1. 최신 버전의 Snyk CLI를 설치하십시오 - [Snyk CLI 설치](../install-or-update-the-snyk-cli/)참조.
2. 프로젝트를 빌드했는지 확인하세요.
3. 스캔할 프로젝트 디렉토리로 이동합니다.
4. `snyk log4shell`을 입력하세요.\
   **참고**: 이 명령은 추가 인수가 필요하지 않으며 지원하지도 않습니다.

## 스캔 결과

스캔이 완료되면 결과가 표시됩니다.

예를 들어:

```bash
$ snyk log4shell
이 명령은 이미 빌드된 애티팩트를 대상으로 합니다. 소스 코드를 테스트하려면 `snyk test`를 사용하세요.

결과:
취약한 log4j 버전이 감지되었습니다: 
         demo-0.0.1-SNAPSHOT/WEB-INF/lib/log4j-core-2.14.1.jar
         demo-0.0.1-SNAPSHOT.war/WEB-INF/lib/log4j-core-2.14.1.jar
         demo-0.0.1-SNAPSHOT.war.original/WEB-INF/lib/log4j-core-2.14.1.jar

해당 취약점을 수정하는 것을 강력히 권장합니다. 버전을 업그레이드로 수정할 수 없는 경우 이곳에서 대처 정보를 참조하세요:
        - https://security.snyk.io/vuln/SNYK-JAVA-ORGAPACHELOGGINGLOG4J-2314720
        - https://snyk.io/blog/log4shell-remediation-cheat-sheet/
```

취약한 log4j 라이브러리의 흔적이 없는 경우 결과는 다음과 같이 표시됩니다:

```
$ snyk log4shell
이 명령은 이미 빌드된 애티팩트를 대상으로 합니다. 소스 코드를 테스트하려면 `snyk test`를 사용하세요.

결과:
취약한 log4j 버전이 감지되지 않았습니다
```

## 수정 권고

영향을 받는 패키지를 수정하는 자세한 내용은 Snyk [Log4Shell 수정 치트시트](https://snyk.io/blog/log4shell-remediation-cheat-sheet)를 참조하세요.

## 제한 사항

* Snyk CLI는 알려진 파일의 데이터베이스와 파일 서명을 비교합니다. Log4Shell 취약성이 log4j 라이브러리 업데이트가 아닌 다른 방식으로 수정된 경우에도 라이브러리는 여전히 결과에 표시됩니다.
* log4j 라이브러리의 소스 코드가 수정된 경우 감지됩니다.
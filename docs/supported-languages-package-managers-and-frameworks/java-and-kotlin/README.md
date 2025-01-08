---
description: '(사용 가능한 기능 : 아래) "및 SCM 가져오기 기능" 확인 필요'
---

# Java and Kotlin

## 적용 가능성

Snyk는 [코드 분석을 위한 Java 및 Kotlin](java-and-kotlin-for-code-analysis.md) 및 [오픈 소스를 위한 Java 및 Kotlin](java-and-kotlin-for-open-source.md) 지원합니다.

{% hint style="info" %}
**지원되는 Java 버전**

Java SE 17까지의 모든 Java 버전을 사용할 수 있습니다.
{% endhint %}

[Java 및 Kotlin을 위한 Snyk CLI에 대한 특별 고려 사항](snyk-cli-for-java-and-kotlin.md) 및 [Maven 및 Gradle과의 SCM 통합](git-repositories-with-maven-and-gradle.md)이 있습니다.

[Java 및 Kotlin에 대한 가이드](guidance-for-java-and-kotlin.md)와 [Java 및 Kotlin을 사용한 Snyk 워크플로우에 대한 정보](snyk-workflow-with-java-and-kotlin.md) 및 [Java 지원에 대한 추가 정보](more-information-about-java-support.md)가 제공됩니다.

Snyk 제품을 사용하여 가져올 수 있는 언어 가용성을 확인하십시오.

사용 가능한 기능:

* 및 용 SCM 가져오기 기능
* CLI와 IDE를 통한 앱 테스트 또는 모니터링, 및 용 사용 가능
* `pkg:maven`을 사용하여 앱의 SBOM 테스트
* `pkg:maven`을 사용하여 앱의 패키지 테스트

## 패키지 관리자 및 지원되는 파일 확장자

Snyk은 Java 및 Kotlin에서 Maven 및 Gradle을 패키지 관리자로 지원하며 [maven.org](https://maven.org/)를 패키지 레지스트리로 지원합니다.

다음 버전 중 하나를 사용할 수 있습니다:

* Maven: `3.*.` 자세한 내용은 [Snyk Maven 플러그인 readme](https://github.com/snyk/snyk-mvn-plugin#support)를 참조하십시오.
* Gradle: `4.*`, `5.*`, `6.*`, `7.*, 8.*.` 자세한 내용은 [Snyk Gradle 플러그인 readme](https://github.com/snyk/snyk-gradle-plugin#support)를 참조하십시오.

Snyk은은 Java 및 Kotlin에서 다음 파일 형식을 지원합니다:

* Snyk Open:
  * Maven: `pom.xml`
  * Gradle: `build.gradle`, `build.gradle.kts`
* :
  * Java: `.java`, `.jsp`, `.jspx`
  * Kotlin: `.kt`

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리가 Java 및 Kotlin에서 Snyk에서 지원됩니다:

* Amazon AWS SDK - 종합
* Android 표준 라이브러리 - 부분
* Apache Commons - 종합
* Apache Tomcat - 부분
* Apache XML - 종합
* apache.mahou - 종합
* bouncycastle - 종합
* com.azure.ai.openai - 종합
* com.google.ai.client.generativeai - 종합
* com.google.cloud.vertexai.generativeai - 종합
* com.google.re2j - 종합
* com.google.gwt - 부분
* Dropwizard - 종합
* elasticsearch - 부분
* FasterXML Jackson - 종합
* Google Guava - 종합
* hibernate - 종합
* http4k - 종합
* io.jsonwebtoken - 종합
* Jakarta EE - 부분
* Jakarta XML Services - 부분
* Java EE - 부분
* Java Servlet - 종합
* Java Servlet (javax) - 종합
* Java Server Pages - 부분
* Java Standard Edition - 종합
* javalin - 부분
* jooq - 종합
* Kyro - 종합
* Micronaut - 종합
* mongo-java-driver - 종합
* Netty - 종합
* okhttp3 - 종합
* org.apache.hc.client5 - 종합
* org.apache.http.client - 종합
* org.apache.sling - 부분
* org.apache.tools.zip - 종합
* org.codehaus.plexus - 종합
* org.dom4j.io - 종합
* Playframework - 종합
* rxhttp - 종합
* Seam logger - 종합
* SnakeYaml - 종합
* Spongycastle - 종합
* Spring boot - 부분
* Spring Web, MVC 및 JDBC - 종합
* Struts - 부분
* Vaadin - 종합
* XStream - 종합

Kotlin 전용:

* Android 표준 라이브러리 - 부분
* com.aallam.openai - 종합
* com.expediagroup.graphql.server - 종합
* Javalin - 부분
* Ktor - 종합
* Kotlin 표준 라이브러리 - 종합
* khttp - 종합

## 기능

다음 기능은 Java 및 Kotlin의 Snyk에서 지원됩니다:

|                                                              |                                                                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| <ul><li>보고서</li><li>수정 PRs (Maven)</li><li>라이선스 스캔</li></ul> | <ul><li>보고서</li><li>사용자 정의 규칙</li><li>인터파일 분석 - Kotlin 완전 지원</li><li>인터파일 분석 - Android 일부 지원</li></ul> |

도움이 필요하면 [Snyk 지원팀에 문의](https://support.snyk.io)하십시오.

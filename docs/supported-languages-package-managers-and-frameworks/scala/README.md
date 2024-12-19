# 스칼라

## 적용 가능성

Snyk는 [코드 분석을 위한 Scala](scala-for-code-analysis.md) 및 [오픈 소스를 위한 Scala](scala-for-open-source.md)를 지원합니다.

언어 사용 가능 여부를 확인하여 Snyk 제품을 사용하여 가져오기, 테스트 또는 모니터링할 수 있습니다.&#x20;

{% hint style="info" %}
**지원되는 Scala 버전**

버전 2.x를 사용할 수 있습니다.
{% endhint %}

사용 가능한 기능:

* SCM 가져오기, Snyk Open Source 및 Snyk Code에서 사용 가능.&#x20;
* CLI 및 IDE를 통해 응용 프로그램을 테스트하거나 모니터링할 수 있습니다. Snyk Open Source 및 Snyk Code에서 사용 가능.
* `pkg:maven`을 사용하여 응용 프로그램의 SBOM 테스트
* `pkg:maven`을 사용하여 응용 프로그램의 패키지 테스트

## 패키지 관리자 및 지원되는 파일 확장자

Snyk는 Scala를 위해 sbt를 패키지 관리자로 지원하고 [maven.org](https://maven.org/)를 패키지 레지스트리로 지원하며 다음 파일 형식을 지원합니다:

* Snyk Open Source: `build.sbt`
* Snyk Code: `.scala`

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리가 Scala를 위해 Snyk에서 지원됩니다:&#x20;

* Akka - 종합적&#x20;
* HTTP4S - 종합적&#x20;
* io.cequence.openaiscala - 종합적&#x20;
* Play Framework - 종합적&#x20;
* Scala 표준 라이브러리 - 종합적
* 모든 [Java 프레임워크 및 라이브러리](../java-and-kotlin/#frameworks-and-libraries)

## 기능

다음 기능이 스칼라를 위해 Snyk에서 지원됩니다:

| Snyk Open Source                                   | Snyk Code                                           |
| --------------------------------------------------- | ---------------------------------------------------- |
| <ul><li>라이선스 스캐닝</li><li>보고서</li></ul> | <ul><li>보고서</li><li>상세 분석</li></ul> |
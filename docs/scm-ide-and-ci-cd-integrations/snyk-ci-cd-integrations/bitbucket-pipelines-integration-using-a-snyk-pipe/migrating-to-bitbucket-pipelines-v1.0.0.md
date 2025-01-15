# Bitbucket 파이프라인 v1.0.0으로 마이그레이션하기

1.0.0 이하 버전에서 1.0.0 이상 버전으로 업그레이드할 때 구성에서 다음 변경 사항을 수행하십시오:

* `LANGUAGE` 변수를 설정할 때 [Snyk 이미지](https://hub.docker.com/r/snyk/snyk/tags)에서 지원하는 태그를 사용하여 [Snyk CLI 도커 이미지](https://hub.docker.com/r/snyk/snyk-cli/tags)에서 지원하는 태그 대신 사용합니다.
* 또는 새 `SNYK_BASE_IMAGE` 변수를 사용하여 [사용자 정의 이미지](../user-defined-custom-images-for-cli.md)를 제공할 수 있습니다.

<figure><img src="../../../.gitbook/assets/Untitled.jpg" alt="사용자는 v1.0.0에서 사용자 정의 이미지를 정의할 수 있습니다"><figcaption><p>사용자는 v1.0.0에서 사용자 정의 이미지를 정의할 수 있습니다</p></figcaption></figure>

## 기본 앱 종속성 검사 예제 <a href="#markdown-header-basic-app-dependencies-scan-example" id="markdown-header-basic-app-dependencies-scan-example"></a>

이 예제는 Snyk을 사용하여 Node.js 애플리케이션을 검사하고 취약점이 발견되면 빌드를 중단하는 방법을 보여줍니다.

{% code title="snyk/snyk-scan:1.0.0" %}
```yaml
script:
  - npm install

  - npm test

  - pipe: snyk/snyk-scan:1.0.0
    variables:
      SNYK_TOKEN: $SNYK_TOKEN
      LANGUAGE: "node" # 언어 태그는 "node"입니다.

  - npm publish
```
{% endcode %}

{% code title="snyk/snyk-scan:0.7.0" %}
```yaml
script:
  - npm install

  - npm test

  - pipe: snyk/snyk-scan:0.7.0
    variables:
      SNYK_TOKEN: $SNYK_TOKEN
      LANGUAGE: "npm" # 언어 태그는 "npm"입니다.

  - npm publish
```
{% endcode %}

## 동등한 Snyk 이미지

다음 표는 Bitbucket 파이프라인 < 1.0.0에서 사용된 Snyk CLI 도커 이미지 및 Bitbucket 파이프라인 > 1.0.0에서 사용할 수 있는 동등한 Snyk 이미지를 나열합니다.

{% hint style="info" %}
CLI 설치 목적으로 모든 Snyk CLI 도커 이미지에 Node.js 14가 설치되어 있습니다. 반면에 `node` 태그가 포함된 Snyk 이미지에만 NodeJS가 설치됩니다.
{% endhint %}

**이미지가 정확히 같지는 않음에 유의하십시오**. 기본 이미지, 설치된 Snyk CLI 버전 등 속성이 다를 수 있습니다.

그러나 Bitbucket 파이프라인을 위해 기능적으로 동일합니다.

지원되는 언어나 프레임워크가 생명주기를 완료했을 경우, 여전히 공급업체에 의해 지원되는 언어 또는 프레임워크의 새 버전을 사용하기 위해 **사용자가 제공한 사용자 정의 이미지**를 사용하는 것이 **강력히 추천**됩니다.

| **언어/도구**      | **Snyk CLI 이미지**                                                                              | **Snyk 이미지**                                                                                                                                                                                                                                                                                                                                                                           |
| -------------- | --------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| PHP/Composer   | <ul><li>1*.*-composer</li><li>composer</li></ul>                                              | <ul><li>snyk/snyk:composer</li></ul>                                                                                                                                                                                                                                                                                                                                                   |
| Docker         | <ul><li>1.<em>.</em>-docker</li><li>docker</li></ul>                                          | <ul><li>snyk/snyk:docker-latest</li><li>snyk/snyk:docker</li></ul>                                                                                                                                                                                                                                                                                                                     |
| Java/Gradle    | <ul><li>1.<em>.</em>-gradle-</li><li><em>gradle-</em></li><li>gradle-<em>_java</em></li></ul> | <ul><li>snyk/snyk:gradle</li><li>snyk/snyk:gradle-jdk11</li><li>snyk/snyk:gradle-jdk12</li><li>snyk/snyk:gradle-jdk13</li><li>snyk/snyk:gradle-jdk14</li><li>snyk/snyk:gradle-jdk16</li><li>snyk/snyk:gradle-jdk17</li><li>snyk/snyk:gradle-jdk18</li><li>snyk/snyk:gradle-jdk19</li><li>snyk/snyk:gradle-jdk20</li><li>snyk/snyk:gradle-jdk21</li><li>snyk/snyk:gradle-jdk8</li></ul> |
| Java/Maven     | <ul><li>1*.<em>-maven-</em></li><li>maven-</li><li><em>maven-</em>_java*</li></ul>            | <ul><li>snyk/snyk:maven</li><li>snyk/snyk:maven-3-jdk-11</li><li>snyk/snyk:maven-3-jdk-17</li><li>snyk/snyk:maven-3-jdk-20</li><li>snyk/snyk:maven-3-jdk-21</li><li>snyk/snyk:maven-3-jdk-22</li><li>snyk/snyk:maven-3-jdk-8</li></ul>                                                                                                                                                 |
| JS(Node)/NPM   | <ul><li>1*.*-npm</li><li>npm</li></ul>                                                        | <ul><li>snyk/snyk:node</li><li>snyk/snyk:node-18</li><li>snyk/snyk:node-20</li><li>snyk/snyk:node-22</li></ul>                                                                                                                                                                                                                                                                         |
| .Net/Nuget     | <ul><li>1*.*-nuget</li><li>nuget</li></ul>                                                    | <ul><li>snyk/snyk:dotnet</li><li>snyk/snyk:dotnet-8.0</li></ul>                                                                                                                                                                                                                                                                                                                        |
| Python         | <ul><li>1*.<em>-python-</em></li><li>python-*</li></ul>                                       | <ul><li>snyk/snyk:python</li><li>snyk/snyk:python-3.8</li><li>snyk/snyk:python-3.9</li><li>snyk/snyk:python-3.10</li><li>snyk/snyk:python-3.11</li><li>snyk/snyk:python-alpine</li></ul>                                                                                                                                                                                               |
| Ruby/Ruby Gems | <ul><li>1*.*-rubygems</li><li>rubygems</li></ul>                                              | <ul><li>snyk/snyk:ruby</li><li>snyk/snyk:ruby-3.3</li><li>snyk/snyk:ruby-alpine</li></ul>                                                                                                                                                                                                                                                                                              |
| Scala/SBT      | <ul><li>1*.<em>-sbt-</em></li><li>sbt-*</li></ul>                                             | <ul><li>snyk/snyk:sbt1.10.0-scala3.4.2</li></ul>                                                                                                                                                                                                                                                                                                                                       |

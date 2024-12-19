---
description: 이 페이지는 영향을 받는 이미지들로부터의 전환 방법을 설명합니다.
---

# Snyk 이미지 이전

당신은 [사용자 정의 이미지를 생성](../../user-defined-custom-images-for-cli.md)하여 사용할 수 있습니다. 이 옵션은 이 페이지에 나열된 Snyk 이미지가 제거되어 영향을 받는 모든 사용자에게 사용 가능합니다.

사용자 정의 이미지를 생성하는 것이 권장되는 해결책이며, 이는 시스템과의 호환성을 보장해야 합니다. 그러나 사용자 정의 이미지를 생성할 수 없다면, 대체 가능한 이미지를 사용할 수 있습니다.

{% hint style="info" %}
Snyk는 명시적으로 제시되지 않는 한 소프트웨어 패키지의 고정 버전을 사용하는 것을 권장합니다.
{% endhint %}

## snyk/snyk:docker-\* <a href="#snyk-snyk-docker" id="snyk-snyk-docker"></a>

제거될 `snyk/snyk:docker-*` 이미지에 대해 아래 표를 참조하십시오.

| 이미지                  | 기반      |
| ---------------------- | ------------ |
| snyk/snyk:docker-18.09 | docker:18.09 |
| snyk/snyk:docker-19.03 | docker:19.03 |

이러한 이미지의 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:docker`
- `snyk/snyk:docker-latest`

{% hint style="info" %}
`snyk/snyk:docker`는 Docker stable 기반으로 선호됩니다.
{% endhint %}

## snyk/snyk:golang-\* <a href="#snyk-snyk-golang" id="snyk-snyk-golang"></a>

제거될 `snyk/snyk:golang-*` 이미지에 대해 아래 표를 참조하십시오.

| 이미지                 | 기반    |
| --------------------- | ------- |
| snyk/snyk:golang-1.12 | golang:1.12 |
| snyk/snyk:golang-1.13 | golang:1.13 |
| snyk/snyk:golang-1.14 | golang:1.14 |
| snyk/snyk:golang-1.15 | golang:1.15 |
| snyk/snyk:golang-1.16 | golang:1.16 |
| snyk/snyk:golang-1.17 | golang:1.17 |
| snyk/snyk:golang-1.18 | golang:1.18 |
| snyk/snyk:golang-1.19 | golang:1.19 |

이러한 이미지의 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:golang`
- `snyk/snyk:golang-1.20`
- `snyk/snyk:golang-1.21`
- `snyk/snyk:golang-1.22`

{% hint style="info" %}
`golang-1.22`와 같은 고정된 Golang 버전을 사용하는 것이 권장됩니다.
{% endhint %}

## snyk/snyk:gradle-\* <a href="#snyk-snyk-gradle" id="snyk-snyk-gradle"></a>

제거될 `snyk/snyk:gradle-*` 이미지에 대해 아래 표를 참조하십시오.

| 이미지                      | 기반            |
| -------------------------- | ---------------- |
| snyk/snyk:gradle-6.4       | gradle:6.4       |
| snyk/snyk:gradle-6.4-jdk11 | gradle:6.4-jdk11 |
| snyk/snyk:gradle-6.4-jdk14 | gradle:6.4-jdk14 |
| snyk/snyk:gradle-6.4-jdk8  | gradle:6.4-jdk8  |

이러한 이미지의 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:gradle`
- `snyk/snyk:gradle-jdk11`
- `snyk/snyk:gradle-jdk12`
- `snyk/snyk:gradle-jdk13`
- `snyk/snyk:gradle-jdk14`
- `snyk/snyk:gradle-jdk16`
- `snyk/snyk:gradle-jdk17`
- `snyk/snyk:gradle-jdk18`
- `snyk/snyk:gradle-jdk19`
- `snyk/snyk:gradle-jdk20`
- `snyk/snyk:gradle-jdk21`
- `snyk/snyk:gradle-jdk8`

{% hint style="info" %}
- `snyk/snyk:gradle`은 발표 시점의 최신 안정 버전인 Gradle 버전 8.8 및 JDK 21을 사용합니다.
- Gradle 이미지는 명시된 JDK와 함께 Gradle 8.8을 사용합니다.
{% endhint %}

## snyk/snyk:maven-3-jdk-19 <a href="#snyk-snyk-maven-3-jdk-19" id="snyk-snyk-maven-3-jdk-19"></a>

`$nyk/snyk:maven-3-jdk-19`는 Docker에서 제거될 것입니다. 이 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:maven`
- `snyk/snyk:maven-3-jdk-11`
- `snyk/snyk:maven-3-jdk-17`
- `snyk/snyk:maven-3-jdk-20`
- `snyk/snyk:maven-3-jdk-21`
- `snyk/snyk:maven-3-jdk-22`
- `snyk/snyk:maven-3-jdk-8`

{% hint style="info" %}
`snyk/snyk/maven`은 발표 시점의 최신 안정 버전인 Maven 3.9를 사용합니다.
{% endhint %}

## snyk/snyk:dotnet-\* <a href="#snyk-snyk-dotnet" id="snyk-snyk-dotnet"></a>

제거될 `snyk/snyk:dotnet-*` 이미지에 대해 아래 표를 참조하십시오.

| 이미지                | 기반                                 |
| --------------------- | ------------------------------------ |
| snyk/snyk:dotnet-2.1  | mcr.microsoft.com/dotnet/core/sdk:2.1 |
| snyk/snyk:dotnet-2.2  | mcr.microsoft.com/dotnet/core/sdk:2.2 |
| snyk/snyk:dotnet-3.0  | mcr.microsoft.com/dotnet/core/sdk:3.0 |
| snyk/snyk:dotnet-3.1  | mcr.microsoft.com/dotnet/core/sdk:3.1 |

이러한 이미지의 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:dotnet`
- `snyk/snyk:dotnet-8.0`

{% hint style="info" %}
추세 버전의 .NET인 8.0을 기반으로 하는 `snyk/snyk:dotnet-8.0`을 사용하는 것이 권장됩니다.
{% endhint %}

## snyk/snyk:node-\* <a href="#snyk-snyk-node" id="snyk-snyk-node"></a>

제거될 `snyk/snyk:node-*` 이미지에 대해 아래 표를 참조하십시오.

| 이미지             | 기반 |
| ----------------- | ---- |
| snyk/snyk:node-8  | node:8 |
| snyk/snyk:node-10 | node:10 |
| snyk/snyk:node-12 | node:12 |
| snyk/snyk:node-13 | node:13 |
| snyk/snyk:node-14 | node:14 |
| snyk/snyk:node-15 | node:15 |
| snyk/snyk:node-16 | node:16 |

이러한 이미지의 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:node`
- `snyk/snyk:node-18`
- `snyk/snyk:node-20`
- `snyk/snyk:node-22`

{% hint style="info" %}
LTS 버전 22를 사용하는 `snyk/snyk:node-22`와 같은 고정 버전을 사용하는 것이 권장됩니다.
{% endhint %}

## snyk/snyk:python-\* <a href="#snyk-snyk-python" id="snyk-snyk-python"></a>

제거될 `snyk/snyk:python-*` 이미지에 대해 아래 표를 참조하십시오.

| 이미지                | 기반 |
| -------------------- | ---- |
| snyk/snyk:python-2.7 | python:2.7 |
| snyk/snyk:python-3.6 | python:3.6 |
| snyk/snyk:python-3.7 | python:3.7 |

이러한 이미지의 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:python`
- `snyk/snyk:python-alpine`
- `snyk/snyk:python-3.8`
- `snyk/snyk:python-3.9`
- `snyk/snyk:python-3.10`
- `snyk/snyk:python-3.11`

{% hint style="info" %}
LTS 버전 3.11을 사용하는 `snyk/snyk:python-3.11`와 같은 고정 버전을 사용하는 것이 권장됩니다.
{% endhint %}

## snyk/snyk:ruby-\* <a href="#snyk-snyk-ruby" id="snyk-snyk-ruby"></a>

제거될 `snyk/snyk:ruby-*` 이미지에 대해 아래 표를 참조하십시오.

| 이미지              | 기반 |
| ------------------ | ---- |
| snyk/snyk:ruby-2.4 | ruby:2.4 |
| snyk/snyk:ruby-2.5 | ruby:2.5 |
| snyk/snyk:ruby-2.6 | ruby:2.6 |
| snyk/snyk:ruby-2.7 | ruby:2.7 |

이러한 이미지의 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:ruby`
- `snyk/snyk:ruby-alpine`
- `snyk/snyk:ruby-3.3`

{% hint style="info" %}
Ruby 3.3를 기반으로 하는 `snyk/snyk:ruby-3.3`와 같은 고정 버전을 사용하는 것이 권장됩니다.
{% endhint %}

## snyk/snyk:sbt 및 snyk/snyk:scala <a href="#snyk-snyk-sbt-and-snyk-snyk-scala" id="snyk-snyk-sbt-and-snyk-snyk-scala"></a>

제거될 `snyk/snyk:sbt` 및 `snyk/snyk:scala` 이미지에 대해 아래 표를 참조하십시오.

| 이미지           | 기반                                      |
| --------------- | ----------------------------------------- |
| snyk/snyk:sbt   | hseeberger/scala-sbt:8u212\_1.2.8\_2.13.0 |
| snyk/snyk:scala | hseeberger/scala-sbt:8u212\_1.2.8\_2.13.0 |

이러한 이미지의 제거로 영향을 받는 사용자들은 다음 이미지로 업데이트할 수 있습니다:

- `snyk/snyk:ruby`
- `snyk/snyk:ruby-alpine`
- `snyk/snyk:ruby-3.3`

{% hint style="info" %}
Scala:3.4.2-sbt:1.10.0를 기반으로 하는 `snyk/snyk:sbt1.10.0-scala3.4.2`를 사용하는 것을 권장합니다.
{% endhint %}
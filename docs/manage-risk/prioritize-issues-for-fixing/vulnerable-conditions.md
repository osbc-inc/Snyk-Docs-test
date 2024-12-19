# 취약한 조건

Exploitable하지 않은 취약점은 애플리케이션에 보안 위협을 제공하지 않을 가능성이 높기 때문에 우선순위가 낮아지고 수정하지 않아도 됩니다.

모든 취약점은 취약점이 위협이 되도록, 즉, exploit할 수 있도록 하는 조건이 있습니다.

취약한 조건의 예로는 다음 기준 중 하나 이상을 충족하는 조건이 포함됩니다:
* 사용자 상호작용 필요
* 특정한 포트 번호가 열려 있어야 함
* 특정 CPU 아키텍처가 필요함
* 특정 설정을 활성화해야 함

일부 취약점은 exploit할 수 있도록 모든 조건을 충족해야 하는 여러 다른 조건을 가지고 있습니다. 취약한 오픈 소스 패키지는 일부 응용 프로그램에서는 exploit될 수 있지만 다른 응용 프로그램에서는 그렇지 않을 수 있습니다.

Exploit 가능성은 맥락에 의존합니다. 맥락의 예로는 환경, 설정 및 개발자가 패키지를 사용하는 방법 등이 있습니다.

## Snyk Triage Assistant

{% hint style="info" %}
**기능 가용성**\
이 기능은 GitHub를 소스로 사용하고 Snyk Code가 활성화된 경우에만 Java (Gradle 및 Maven) 생태계에서 사용할 수 있습니다.
{% endhint %}

당신의 애플리케이션 맥락에서 Triage Assistant는 취약한 조건을 평가하여 애플리케이션의 exploit 가능성을 결정하는 데 도움이 됩니다.

Snyk Code ([SAST](https://snyk.io/learn/application-security/sast-vs-dast/)) 엔진은 당신의 퍼스트 파티 코드를 읽고 Snyk Open Source (SCA)에서 발견된 취약점의 조건을 확인하는 데 사용됩니다.

{% hint style="info" %}
이 기능을 제공하기 위해 Snyk는 당신의 Git 저장소 콘텐츠의 임시 복사본을 만듭니다.

자세한 정보는 [Snyk가 데이터를 처리하는 방법](../../working-with-snyk/how-snyk-handles-your-data.md)을 참조하십시오.
{% endhint %}

## Jackson 취약한 조건의 평가

Jackson 취약한 조건은 취약점이 exploit될 수 있도록 다음 조건을 모두 충족해야 합니다.

* 취약한 버전: [Jackson 패키지](https://snyk.io/vuln/maven:com.fasterxml.jackson.core%3Ajackson-databind) (**com.fasterxml.jackson.core:jackson-databind 취약점**)은 Snyk가 취약하다고 알고 있는 특정 버전이어야 합니다.
* **특정 설정:** 특정 설정 또는 기능이 활성화되어 있어야 하며, 이 경우 Polymorphic Type Handling 기능이어야 합니다.
  * 당신의 코드에서 이 설정이 활성화되었는지 확인하려면 다음과 같은 것 중 하나를 찾아보면 됩니다:
    * `@JsonSubTypes` 주석이 사용되었음.
    * 클래스에 `@JsonTypeInfo` 주석이 사용되었음.
    * `enableDefaultTyping()`이 Polymorphic Typing을 활성화하기 위해 사용되었음.
    * `enableDefaultTypingAsProperty()`가 Polymorphic Typing을 활성화하기 위해 사용되었음.
* 사용자 상호작용: 애플리케이션은 사용자로부터 JSON 입력을 수락해야 함.
* 특정 가젯: 실행 범위 내에서 사용 가능해야 하는 "가젯"인 클래스나 함수가 있어야 합니다.

## Exploit Maturity가 있는 취약점이 exploit될 수 없는 경우

취약점은 이미 악용되는 경우가 있거나 어떻게 exploit할지에 대한 자세한 설명이 있을 수 있지만, 모든 조건이 충족되지 않는 한 취약점은 여전히 exploit될 수 없습니다.
# Snyk 코드 사용자 지정 규칙

{% hint style="info" %}
**기능 가용성**

Snyk Code 사용자 정의 규칙은 조기 액세스 상태이며 엔터프라이즈 계획에서만 사용할 수 있습니다. 이 기능을 활성화하려면 [Snyk 미리보기](../../../snyk-admin/snyk-preview.md)를 참조하세요.
{% endhint %}

조사 워크플로우의 일환으로 코드 스택에 대한 쿼리를 실행하기 위해 사용자 정의 규칙을 만듭니다. 사용자 정의 규칙을 사용하는 방법은 다음과 같습니다:

* 보안 팀이 우려할 수 있는 사용자 정의 취약한 메소드(Sink)를 정의합니다.
* 코드에 포함되어서는 안 되는 비밀 및 자격 증명 사용을 확인하기 위해 정규식 스캔을 생성합니다.
* 보안 팀이 안전하지 않다고 판단하는 특정 원치 않는 방법이 코드 기반 내에서 호출되고 있는지 확인합니다.
* 쿼리를 생성한 후, 이를 이전에 Snyk로 가져온 [코드 스니펫](run-query.md#run-query-on-a-code-snippet) 또는 [저장소](run-query.md#run-query-on-a-repository)에 대해 테스트합니다. 이렇게 하면 정기 스캔을 실행하기 전에 쿼리 결과를 볼 수 있습니다. 이 기능을 사용하여 규칙을 프로덕션에 푸시하기 전에 규칙을 확인하여 예상한 결과가 제공되는지 확인할 수 있습니다.

사용자 정의 규칙은 Snyk Code에서 어떤 스캔에서든 실행할 수 있으며 다음 중 하나를 사용할 때 이러한 기능을 사용할 수 있습니다:

* [Snyk 웹 UI](../../../getting-started/snyk-web-ui.md)
* [Snyk CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)
* [IDE](../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)

## Snyk 웹 UI에서 사용자 정의 규칙 사용

Snyk 웹 UI에서 현재 Snyk 스캔과 완전히 분리된 로컬 환경에서 쿼리를 생성하고 테스트할 수 있습니다. 다음 중 하나의 작업을 Snyk 웹 UI에서 수행하세요:

* [저장소에 대한 쿼리 실행](run-query.md#run-query-on-a-repository)
* [코드 스니펫에 대한 쿼리 실행](run-query.md#run-query-on-a-code-snippet)
* [쿼리 결과 분석](run-query.md#analyze-query-results)
* [사용자 정의 규칙 저장](create-custom-rule.md)

## Snyk CLI에서 사용자 정의 규칙 사용

Snyk CLI를 사용하여 Snyk 코드 프로젝트를 정기적인 명령과 옵션을 사용하여 테스트할 수 있습니다. 이때 Snyk 웹 UI를 사용하여 생성한 사용자 정의 규칙이 포함된 [.snyk 파일](../../../manage-risk/policies/the-.snyk-file.md)이 있어야 합니다. [Snyk 코드용 Snyk CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)을 참조하세요.

## IDE에서 사용자 정의 규칙 사용

Snyk과의 IDE 통합은 [.snyk 파일](../../../manage-risk/policies/the-.snyk-file.md)을 사용하여 Snyk 웹 UI에서 생성한 사용자 정의 규칙을 사용할 수 있습니다.

## Snyk Code 사용자 정의 규칙 작동 방식

### 쿼리 언어

Snyk Code 사용자 정의 규칙은 논리 프로그래밍인 Datalog을 기반으로 하는 독점적인 선언형 쿼리 언어를 사용합니다.

목표는 유용하고 실행 가능한 결과를 제공하며 보안 팀과 개발자가 코드 내에서 가장 중요한 취약점에 초점을 맞출 수 있도록 하는 유익하고 실행 가능한 쿼리를 작성하는 것입니다.

### Snyk Code 사용자 정의 규칙 구성 요소

#### 쿼리 템플릿

이러한 템플릿은 빠르고 쉽게 쿼리를 작성하는 방법을 제공하도록 만들어진 추상적인 사전 구축된 구성물입니다. [템플릿 및 예측](templates-and-predicates.md)을 참조하세요.

<figure><img src="../../../.gitbook/assets/query templates.png" alt="쿼리 템플릿"><figcaption><p>쿼리 템플릿</p></figcaption></figure>

#### 쿼리 예측

예측은 객체나 속성 간의 관계를 상징적으로 나타내고 참/거짓을 평가하는 것입니다. Snyk은 미리 정의된 예측의 철저한 목록을 제공합니다.

예를 들어, 모든 크로스 사이트 스크립팅(XSS) Sink가 `PRED:XssSink`로 표시됩니다. 이러한 예측을 확장하거나 고유한 예측을 정의할 수 있습니다. [쿼리 언어](./#the-query-language)를 참조하세요.

<figure><img src="../../../.gitbook/assets/query predicates.png" alt="쿼리 예측 요약"><figcaption><p>쿼리 예측</p></figcaption></figure>

#### 소스

소스는 데이터 입력의 진입점이며 사용자 또는 환경에 의해 제어될 수 있습니다. 많은 경우, 소스는 오염될 수 있다고 가정해야 합니다.

#### 살균제

살균제는 사용자 또는 환경으로부터의 데이터 입력을 살균하여 데이터가 오염되지 않도록 보장합니다. 이 살균 처리를 통해 오염된 데이터가 Sink에 의해 사용되는 것을 방지합니다.

#### Sink

Sink는 데이터가 사용되는 지점입니다. 사용된 데이터가 오염되었을 경우 응용 프로그램 내에 취약점이 발생할 수 있습니다.

#### 호스팅 규칙

사용자 정의 규칙은 `.snyk` 파일을 사용합니다. 저장소가 가져올 때마다, 이 파일은 정기 캐싱 프로세스의 일부로 수집됩니다.

`.snyk` 파일에 사용자 정의 규칙이 포함되어 있으면 스캔을 실행할 때 이러한 규칙은 일반 Snyk 기본 규칙과 함께 실행되어 다른 규칙과 동일한 결과를 제공합니다.

### 사용자 정의 규칙 동작

Snyk Code 사용자 정의 규칙은 다른 규칙과 동일하게 작동합니다. Snyk Code는 코드를 분석하여 추상 구문 트리(abstract syntax tree, AST)를 생성하고 이를 이벤트 그래프를 생성하여 분석합니다.

모든 Snyk Code 규칙, 사용자 정의 규칙 포함하여, 이벤트 그래프에 대한 규칙이 실행되며 일치하는 모든 것은 취약점으로 간주되어 개발자나 보안 팀에 식별됩니다.

취약점이 해결된 경우, 보고 탭에서 해결된 문제 섹션에 제거되어 추가됩니다. [.snyk 파일](../../../manage-risk/policies/the-.snyk-file.md)을 참조하세요.

### 제안적 AI 지원

Snyk Code는 코드를 간단하게 정의하고 테스트하기 위해 AI 기술을 활용한 사용자 친화적인 개발 환경을 제공합니다. AI는 질의에 대한 도움이 되는 제안을 제공하는 직관적인 도우미로 작동합니다.

예를 들어, 데이터가 흐르는 메소드를 찾아야 하는 경우 `DataFlowsInto`[쿼리 템플릿](./#query-templates)을 사용할 수 있습니다. 그러면 AI가 이벤트 그래프를 기반으로 코드 내에서 데이터가 흐르는 메소드를 제안합니다. 이렇게 하면 규칙 생성 프로세스가 간소화되며 새로운 쿼리 아이디어가 떠오를 수도 있습니다.

<figure><img src="../../../.gitbook/assets/suggestive_ai_support (1).gif" alt="제안적 AI 지원"><figcaption><p>제안적 AI 지원</p></figcaption></figure>

### 쿼리 언어

Snyk Code는 코드 검색을 위한 도메인 특화 언어를 사용하여 사용자 정의 쿼리를 제공합니다. 이는 우리의 경우, 코드에 사용될 논리 선언형 프로그래밍 언어이며 Turing 완전하지 않습니다. 이는 쿼리 언어로 작성된 모든 쿼리가 종료되고 결과가 0개, 1개 또는 그 이상이 반환됨을 보장한다는 장점을 가지고 있습니다.

쿼리 언어는 코드에 대한 사용하는 프로그래밍 언어와 독립적이며, 규칙은 모든 Snyk 지원 언어에서 작동합니다. 코드 스니펫이 제공되는 경우, 제공된 코드 스니펫의 언어를 선택해야 합니다.

{% hint style="info" %}
쿼리 언어는 대소문자를 구분합니다.
{% endhint %}

쿼리 언어는 코드에서 일치하는 항목을 찾기 위한 언어입니다. 각 쿼리는 지정된 프로퍼티와 일치하는 쿼리에 대한 코드 내 요소를 발견합니다.

쿼리 언어의 첫 번째 기능은 값을 기반으로 프로그램 요소를 매치하는 것입니다. 값은 이중 따옴표로 값을 인용하여 수행됩니다. 프로그램 요소는 완전한 자격 이름을 사용하여 식별됩니다. 다음과 같은 Java 코드 예제를 고려하세요:

```java
import java.time.LocalDate;
class Test {
 static void test() {
   System.out.println("test" + 123);
   System.out.println(LocalDate.now());
 }
}
```

`"java.time.LocalDate.now"` 쿼리를 사용하여 현재 시간을 가져오는 메소드 호출을 일치시킬 수 있습니다.

`"test"`를 사용하여 함수 선언인 _test_ 및 문자열 '`test`'를 일치시킬 수 있습니다.

숫자 값 123은 `"123"`을 사용하여 일치시킬 수 있습니다. 요소를 일치시키기 위해 문자열, 숫자 또는 기타 값의 식별자와 유형에 관계없이 인용 부호를 사용합니다. 요소를 정규표현식을 사용하여 일치시킬 수 있습니다. 정규표현식은 인용 부호 앞에 기호 `~`를 넣어 식별됩니다. 예를 들어, 프로그램 요소 `123`은 표현식 `~"12.*"`와 일치될 수 있습니다. 프린트 문은 `"java.lang.System.out.println"` 또는 `~".*\.println"`과 같은 쿼리로 일치할 수 있습니다.

프로그램 요소의 올바른 완전한 이름을 사용하려면 프로그램 요소를 제공된 코드 스니펫이나 제공된 저장소 내에서 가지고 있는 프로그램 요소의 값을 자동 완성합니다.

#### **예측문 (Predicate /&#x20;**_**PRED**_)

예측문은 미리 정의된 조건에 따라 프로그램 요소를 일치시킵니다. 예측문의 주요 장점은 기존 Snyk Code 지식베이스를 활용할 수 있다는 것입니다. 예를 들어, HTTP 서버가 쿠키를 처리하는 프로그램 위치를 모두 찾아야 하는 경우 `PRED:SourceCookie`를 사용할 수 있습니다.

비슷하게, SQL 쿼리를 처리하는 프로그램 위치를 찾을 수 있는 `PRED:SqliSink`가 있습니다. 사용 가능한 모든 예측문을 발견하기 위해 사용자 정의 규칙은 자동 완성 기능을 제공합니다. 모든 프로그램 요소 또는 프로그램 요소가 없음을 매치하는 두 가지 특별한 예측문 `PRED:Any`와 `PRED:None`이 있습니다.

{% hint style="info" %}
여러 일치를 제공하면 모든 조합의 결과가 됩니다.

예를 들어, 쿼리 `PRED:SourceCookie ~"get.*"`와 같은 방식으로 사용하면 쿠키를 반환하며 이름이 `get`로 시작하는 방법에 매치되는 메소드만 일치합니다(요소를 일치시킬 두 조건의 논리 AND).
{% endhint %}

#### **템플릿**

템플릿은 제공된 조건을 결합하는 데 사용됩니다. 템플릿 자체는 제공된 템플릿 매개변수로 작성한 하나 이상의 조건을 기술한다. 사전 정의 템플릿은 여러 사용 가능한 경우를 위해 설계되었으며 여기에서 정의된 사용 사례가 정의되어 있습니다.

템플릿을 사용하여 매개변수의 매치를 제한할 수 있습니다. 예를 들어 `StringLiteral<"test">`는 값이 test인 모든 프로그램 요소를 가져오며 문자열 리터럴인 요소만 반환합니다.

프로그램의 다른 요소 간의 관계를 나타내기 위해 템플릿을 사용할 수 있습니다. 다음 쿼리는 첫 번째 인수로 문자열 리터럴 test를 가진 모든 프로그램 엔티티를 찾을 것입니다:\
`HasArg1<StringLiteral<"test">>`.

템플릿 `HasArg1`은 프로그램 요소 사이의 의미적 관계를 인코딩합니다. 예를 들어, 다음

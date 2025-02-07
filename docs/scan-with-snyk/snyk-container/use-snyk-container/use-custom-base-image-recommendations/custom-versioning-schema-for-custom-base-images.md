# 사용자 지정 기본 이미지에 대한 사용자 지정 버전 관리 스키마

사용자 지정 버전 스키마(CVS)는 Snyk이 회사의 컨테이너 이미지 태그 버전 스키마를 이해하는 방법입니다. 이를 통해 Snyk은 정확한 기본 이미지 업그레이드 권고를 제공할 수 있습니다.

## Snyk CVS를 위한 전제 조건

CVS는 Snyk 컨테이너의 사용자 정의 베이스 이미지 기능의 일부입니다. 자세한 내용은 [사용자 정의 베이스 이미지 권장 사항 사용](./)을 참조하십시오.

## CVS를 사용해야 하는 시기

컨테이너 이미지의 태그가 [Semantic Versioning](https://semver.org/)(SemVer)를 따르지 않는 경우에는 사용자 정의 버전 스키마를 이미지 저장소에 선택하는 것이 매우 권장됩니다.

## CVS 표현 가이드

CVS는 사실상 이미지 태그의 다른 부분을 비교 가능한 섹션으로 그룹화하는 정규 표현식입니다.

### CVS의 순서

예를 들어 다음 이미지 태그를 고려해 봅시다.

```
snyk/example:1.2_V3
snyk/example:1.2_V4
snyk/example:1.3_V1
```

이 저장소의 이미지 태그가 의미론적 버전 관리 표준을 따르지 않으므로 태그를 사용자 지정 버전 스키마로 설명해야 합니다.

`snyk/example` 태그 스키마는 다음 요소에 의해 정의됩니다.

1. 가장 중요한 영향을 미치는 값의 숫자 (주요 부분)
2. 점
3. 첫 번째 숫자와 중요도가 낮은 다른 숫자 (마이너 부분)
4. "V"로 시작하는 밑줄(버전)
5. 값이 가장 중요하지 않은 숫자.

Snyk이 서로 다른 부분과 역할을 이해하려면 스키마를 정의해야합니다. 이 정규 표현식에서 명명된 그룹은 중요한 변수를 나타냅니다.

아래의 스키마는 위 예제와 해당 요소의 번역 버전입니다.

```regex
(?<SIGNIFICANT>\d+)\.(?<LESS_SIGNIFICANT>\d+)_V(?<LEAST_SIGNIFICANT>\d+)
```

그룹을 "SIGNIFICANT"로 지정하는 대신, 이름을 숫자 뒤에 따르는 "C"로 바꾸어 사용합니다. "C"는 "비교"를 나타내며, 숫자는 해당 그룹의 중요성을 나타냅니다. 여기서 0이 가장 중요합니다.

<pre class="language-regex"><code class="lang-regex">(?&#x3C;<a data-footnote-ref href="#user-content-fn-1">C0</a>>\d+)\.(?&#x3C;<a data-footnote-ref href="#user-content-fn-2">C1</a>>\d+)_V(?&#x3C;<a data-footnote-ref href="#user-content-fn-3">C2</a>>\d+)
</code></pre>

그러면 Snyk이 다음을 수행합니다:

* 이 표현식을 사용하여 저장소의 모든 태그를 구문 분석합니다.
* 중요도 순서대로 값을 비교합니다.
* 태그별로 정렬된 이미지 세트를 생성합니다.

그런 다음 Snyk은 이 정렬된 세트를 사용하여 더 나은 권고 사항을 제공할 수 있습니다.

### CVS로 필터링

다음 예시는 이전 예시의 저장소를 확장하고 각 이미지에 슬림 버전을 추가한 것입니다.

```
snyk/example:1.2_V3-full
snyk/example:1.2_V3-slim
snyk/example:1.2_V4-full
snyk/example:1.2_V4-slim
snyk/example:1.3_V1-full
snyk/example:1.3_V1-slim
```

다음 스키마를 사용하면 현재 사용 중인 플레이버에 대해서만 권장 사항을 받게 될 것입니다.

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)_V(?&#x3C;C2>\d+)<a data-footnote-ref href="#user-content-fn-4">\-(?&#x3C;M0>.*)</a>
</code></pre>

여기에 "M"이라는 새로운 그룹 M0을 추가합니다. Snyk는 이 그룹을 사용하여 MATCH 그룹의 값이 같지 않은 이미지 권장 사항을 걸러냅니다.

{% hint style="info" %}
"M" 뒤에 오는 숫자는 태그가 가질 수 있는 여러 범주를 구별합니다. 이 경우에는 일치할 항목의 범주가 하나이므로 MATCH 그룹이 하나만 있습니다.
{% endhint %}

예를 들어 `snyk/example:1.2_V3-slim`에 대한 권장 사항에는 "slim"이 아닌 M0 그룹 값이 같지 않은 이미지는 포함되지 않습니다.

## 표현식 예시

### 표준 형식 표현식

다음 예시는 이 레지스트리의 표현식을 개발하는 방법을 보여줍니다.

```
snyk/example:1.2.5_deb9_2023061209
snyk/example:1.2.5_deb9_2023090208
snyk/example:1.2.5_deb10_2023110208
snyk/example:1.2.6_deb10_2023110508
snyk/example:1.4.1_deb10_2023110808
snyk/example:1.4.1_deb10_2023110208-slim
```

먼저 간단한 SemVer 요소를 정의합니다.

```regex
(?<C0>\d+)\.(?<C1>\d+)\.(?<C2>\d+)
```

이는 태그의 메이저, 마이너 및 패치 부분을 그룹화합니다.

다음은 기본 OS 배포를 나타내는 섹션입니다. 여기에서 두 가지 옵션이 있습니다:

* 기본 이미지에서 사용 중인 OS가 중요하지 않고 버전의 크기만 중요한 경우, `_.*_`로 그룹화하지 않고 이 섹션을 "무시합니다".
* 분포 간 불일치를 피하기 위해 MATCH 그룹을 추가합니다: `(?<M0>deb\d+)`

이제 표현식은 다음과 같습니다.

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)\.(?&#x3C;C2>\d+)<a data-footnote-ref href="#user-content-fn-5">_(?&#x3C;M0>deb\d+)_</a>
</code></pre>

다음은 날짜 요소입니다. 때로는 날짜가 추가 정보만 제공하는 경우이며 버전을 비교할 때 고려할 필요가 없을 수도 있습니다. 이 경우, 이는 건너뛸 수 있습니다.

날짜 요소가 중요한 경우, 각 날짜 요소가 SemVer에 비해 중요성의 정도를 결정해야 합니다. 예를 들어, 연도가 메이저 버전보다 중요한 경우는요?

중요도 순서를 유지하려면 다음 정규 표현식을 사용하세요:

```regex
(?<C3>\d{4})(?<C4>\d{2})(?<C5>\d{2})(?<C6>\d{2})
```

날짜를 연도, 월, 일 및 시간을 연결한 숫자를 올바르게 비교할 수 있는 방식으로 순서대로 나열하므로 위의 긴 정규 표현식 대신에 더 간단한 것을 사용할 수 있습니다:

```regex
(?<C3>\d{10})
```

이제 정규 표현식은 다음과 같습니다.

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)\.(?&#x3C;C2>\d+)_(?&#x3C;M0>deb\d+)_(?&#x3C;C3>\d{10})<a data-footnote-ref href="#user-content-fn-6">(?:\-(?&#x3C;M1>.*))?</a>
</code></pre>

마지막에는 선택적인 플레이버가 있습니다. 여기에 추가적인 MATCH 그룹을 추가하고 이를 선택 사항으로 만듭니다:

```regex
(?:\_(?<M1>.*))?
```

이를 통해 `slim`을 사용하지 않을 때는 `slim`이 권장되지 않고, 사용 중일 때는 `slim`만 권장되도록합니다.

전체 사용자 지정 버전 스키마 표현은 다음과 같습니다:

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)\.(?&#x3C;C2>\d+)_(?&#x3C;M0>deb\d+)_(?&#x3C;C3>\d{10})<a data-footnote-ref href="#user-content-fn-7">(?:\-(?&#x3C;M1>.*))?</a>
</code></pre>

### 일관성 없는 태그 형식 (선택적 그룹)

저장소가 일관되지 않은 태그 형식을 가지고 있는 경우, 캡처 그룹을 사용할 수 있습니다.

```
snyk/example:1.1
snyk/example:1.1.2
snyk/example:1.2
snyk/example:1.2.4
snyk/example:1.2.8
snyk/example:1.3
snyk/example:1.3.5
```

위의 저장소는 일관되지 않은 수의 캡처 그룹을 포함합니다. 이를 처리하려면 다음 표현식을 사용하십시오:

<pre class="language-regex"><code class="lang-regex">(?&#x3C;C0>\d+)\.(?&#x3C;C1>\d+)<a data-footnote-ref href="#user-content-fn-8">(?:\.(?&#x3C;C2>\d+))?</a>
</code></pre>

`(?:\.(?<C2>\d+))?` 부분에서는 표현식이 선택적으로 COMPARE 그룹을 포함합니다.

COMPARE 그룹의 수가 불일치하면, Snyk는 비교를 정확하게 수행할 정보가 충분하지 않은 태그를 필터링합니다. 즉, `snyk/example:1.2.4`에 대한 권고 사항을 얻기 위해 Snyk는 `snyk/example:1.2`를 더 높은 버전으로 간주하지 않습니다. 왜냐하면 `1.2`가 `1.2.0`과 동일한 것인지 또는 `1.2.8`을 가리키는 롤링 태그인지 여부를 알 수 없기 때문입니다. 그러나 `1.3` 및 `1.3.5`는 모두 `1.2.4`보다 높으며 가능한 권장 사항으로 고려됩니다.

`1.3` 및 `1.3.5`가 `1.2` 및 `1.2.4`와 같은 문제가 있는 경우, 권장 사항은 `1.3` 또는 `1.3.5`입니다. 이 시점에서 권장되는 특정 버전은 정의되지 않은 내부 요소에 따라 결정됩니다. 그러나 Snyk은 이 논리를 개선하기 위해 노력합니다.

## CVS의 기술적 규칙

### 정규 표현식 구문 스타일

* CVS는 ECMAScript regex의 하위 집합을 사용합니다. [MDN의 ECMAScript regex 구문 안내](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions)를 참조하십시오.
* Backreferences 및 lookahead assertion은 지원되지 않습니다. 내부적으로, Snyk는 RE2 라이브러리를 사용합니다. 지원되지 않는 기능 목록 전체는 [re2 GitHub의 Syntax](https://github.com/google/re2/wiki/Syntax)를 참조하십시오.
* 정규 표현식 문자열은 ECMAScript regex로 구문 분석되고 내부적으로 RE2 구문으로 변환됩니다. 예를 들어, 그룹 지정에는 `(?<name>re)` 구문을 사용해야 합니다. `(?P<name>re)`는 올바르게 파싱되지 않습니다.
* [지원되는 기능 목록](https://github.com/google/re2/wiki/Syntax)에서 구문이 아닌 기능만을 고려하여야 합니다.

### 크기 제한

표현식의 최대 길이는 1,000자이며, 최대 100개의 COMPARE 그룹 및 100개의 MATCH 그룹을 가질 수 있습니다.

### 그룹

* 모든 명명된 그룹은 "C" 또는 "M"이라는 글자로 시작해야하며, 인덱스가 뒤에 따라야합니다.
* 적어도 하나의 COMPARE 그룹이 있어야 합니다.
* 명명되지 않은 그룹이 필요하면 캡처하지 않음으로 명시적으로 표시해야 합니다.

### 비교 로직

* 비교 그룹이 부호 없는 정수인 경우, 숫자 값을 비교합니다. 그렇지 않으면 문자열로 처리됩니다.
* 숫자가 아닌 문자의 비교는 UTF-16 코드 단위 값의 순서를 비교하여 수행됩니다. 자세한 정보는 [ECMAScript® 2023 언어 사얌규](https://tc39.es/ecma262/#sec-abstract-relational-comparison)을 참조하십시오.

### Matcher 로직

* MATCH 그룹은 대소문자를 구분합니다.
* 선택 그룹을 사용하여 한 태그에는 MATCH 그룹이 있고, 다른 태그에는 그룹이 없는 경우, 선택적 그룹 사용으로 인해 태그가 일치하지 않는 것으로 간주됩니다.

## CVS의 오류

`Use of unsupported or invalid regex syntax`

### 왜 발생하는가요?

ECMAScript 구문을 준수하지 않는 표현식을 전달하는 경우이 오류가 발생할 수 있습니다.

예를 들어 Python의 명명된 캡처 그룹 구문 `(?P<C0>.*)`를 사용하는 경우입니다.

### 어떻게 해결할 수 있나요?

#### 특정 문자가 아닌 [양자](https://developer.mozilla.org/ko/docs/Web/JavaScript/Guide/Regular_Expressions/Quantifiers)와 [문자 클래스](https://developer.mozilla.org/ko/docs/Web/JavaScript/Guide/Regular_Expressions/Character_Classes)를 사용하세요.

태그를 설명하는 데 1,000자 이상이 필요하다면 CSV는 적합하지 않을 수 있습니다.

다른 옵션에 대한 정보는 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.

[^1]: 가장 중요한 것

[^2]: 덜 중요한 것

[^3]: 가장 덜 중요한 것

[^4]: 우리가 추가한 것

[^5]: 우리가 추가한 것

[^6]: 우리가 추가한 것

[^7]: 우리가 추가한 것

[^8]: 선택적인 캡처 그룹 내에서 그룹을 비교

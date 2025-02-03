# 쿼리 만들기

Snyk Code 사용자 정의 규칙을 쿼리 생성에 사용하려면 [제안하는 AI 지원](./#suggestive-ai-support)을 통해 제공된 [템플릿](./#template) 및 [predicates](./#pred)에서 선택할 수 있습니다. (./#predicate-pred) 또는 자체 프리디케이트를 만들어서 [사용자 정의 규칙으로 저장할 수 있습니다.](create-custom-rule.md).

Snyk Code 사용자 정의 규칙과 함께 사용할 쿼리 예제 및 규칙을 고려하십시오. 이 페이지에서 [CWE 312 쿼리 예제](create-query.md#cwe-312-query-example)가 제공됩니다.

## 간단한 구문 쿼리

다음 소스 코드 스니펫을 복사하여 스니펫 창으로 이동하고 언어로 \*\*C#\*\*을 선택하십시오

{% hint style="info" %}
이것은 단지 스니펫이며 완전한 프로그램이 아닙니다. 컴파일되지 않습니다.
{% endhint %}

<pre class="language-csharp"><code class="lang-csharp"><strong>// 요청 본문 읽기
</strong>string body;
using (var reader = new StreamReader(context.Request.Body))
{
   body = await reader.ReadToEndAsync();
}
// JSON 데이터 구문 분석
var form = JsonConvert.DeserializeObject&#x3C;SignupForm>(body);
var sql = String.Format("INSERT INTO submissions(email, name) VALUES('%s', '%s')", form.Email, form.Name);
form.Email = "nobody@notrealdomain.co.uk";
using var cmd = new NpgsqlCommand(sql, conn);
</code></pre>

### 쿼리 실행

다음 쿼리를 쿼리 창에 입력하고 **쿼리 실행**을 눌러 결과를 확인합니다.

1. 쿼리에 의해 `body`를 선택하세요: `“body”`

{% hint style="info" %}
이 쿼리는 Body를 대문자 'B'로 선택하지 않습니다. 쿼리 언어는 대소문자를 구분합니다.
{% endhint %}

2. 쿼리에 `Body`를 결과에 추가하여 쿼리를 `Or<”body”,”Body”>`로 만드세요.
3. 동일한 결과를 얻으려면 정규식 `~"body|Body"` 또는 `~"[Bb]ody"`를 사용할 수 있습니다.
4. 더 복잡한 정규식 및 쿼리 수행:\
   ``~"[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?"``\
   이것은 하드코딩된 이메일 주소와 일치합니다.

<figure><img src="../../../.gitbook/assets/simple syntactical query.png" alt="Syntactical query example"><figcaption><p>구문 쿼리 예제</p></figcaption></figure>

### 직접 해보세요

코드에 대해 다음 쿼리를 실행하십시오: `~"([a-zA-Z0-9+/]{40})"` 결과가 나오면 확인해 보세요. AWS 비밀번호를 유출할 수 있습니다.

특정 객체 유형에 관심이 있는 경우 [템플릿](templates-and-predicates.md)을 사용할 수 있습니다. 예를 들어, 쿼리 `CallExpression<"Format">`는 함수 호출을 일치시키고 `Literal<"nobody@notrealdomain.co.uk">`은 이메일 주소가 포함된 문자열과 일치합니다.

## 데이터 흐름 또는 오염 분석

이 예시에서 JavaScript 코드 스니펫을 사용합니다. 스니펫 창에 복사하고 언어로 **JavaScript**를 선택하십시오.

```javascript
const express = require('express');
const bodyParser = require('body-parser');
const { Client } = require('pg');
const fs = require('fs');


const app = express();
app.use(bodyParser.json());


const client = new Client({
   host: 'localhost',
   user: 'youruser',
   password: 'yourpassword',
   database: 'yourdbname'
});


async function connectDb(client) {
   await client.connect();
}


async function insertSubmission(client, email, name) {
   await client.query(`INSERT INTO submissions(email, name) VALUES(${email}, ${name})`);
}


function logSubmission(email, name) {
   const logMessage = `New submission: Email=${email}, Name=${name}\n`;
   fs.appendFileSync('myapp.log', logMessage);
}


app.post('/signup', async (req, res) => {
   try {
       const { email, name } = req.body;
       await insertSubmission(client, email, name);
       logSubmission(email, name);
       res.send({ message: 'Signup successful!' });
   } catch (err) {
       console.error(err);
       res.status(500).send({ message: 'An error occurred.' });
   }
});


connectDb(client).then(() => {
   app.listen(3000, () => console.log('Server is running on port 3000'));
});

```

Snyk Code는 `PRED: AnySource` 프리디케이트에서 알려지지 않은 외부 데이터 소스 목록을 알고 있습니다. 다음 쿼리는 `app.post()`가 식별된 것을 보여줍니다.

쿼리 `PRED: SqliSinks`는 `query()`가 SQL 인젝션 싱크 목록에 포함되어 있음을 보여줍니다. 쿼리 엔진에는 다양한 소스, 싱크 및 살균제 타입을위한 많은 다른 프리디케이트가 함께 제공됩니다. 이들을 모두 보려면 프리디케이트 목록을 확인하십시오.

데이터가 SQL 인젝션 싱크로 흘러들어가는지 확인하려면 다음을 사용하십시오. `DataFlowsInto<PRED:SqliSink>`. 이는 프로그램에서 `req` 매개변수에서 `query()`로 여러번 흐르는 데이터를 보여줍니다.

데이터 흐름이 살균제를 거치면서 흐르는 경우 특수화된 템플릿을 사용할 수 있습니다. 쿼리를 ​​`Taint<PRED:AnySource, PRED:SqliSanitizer, PRED:SqliSink>`로 변경하십시오.

{% hint style="info" %}
쿼리에 언어별 특성이 없습니다. 다른 언어에서 비슷한 코드에서도 작동합니다.
{% endhint %}

## **신규 데이터 흐름 규칙**

Snyk이 내부 빌드된 프로프리어터리 소스를 인식하지 못하여 검사 중에 발견되는 취약점을 놓치게 되므로 **신규 규칙을 만드세요.**

데이터 흐름 [템플릿](templates-and-predicates.md)인 `Taint`를 사용하여 [데이터 흐름 쿼리](run-query.md#run-query-on-a-repository)를 만들 때 쓰는 것이 좋습니다.

```javascript
Taint<PRED:"SourceFoo",PRED:XssSanitizer,PRED:XssSink>
```

다음 매개변수를 구성할 수 있습니다:

* **소스**: 첫 번째 매개변수는 데이터 흐름이 시작되는 곳을 나타냅니다.
* **살균제**: 데이터를 살균화하여 오염되지 않고 데이터를 전달하는 알려진 살균화제를 나타내는 두 번째 매개변수
* **싱크**: 데이터 흐름이 끝나는 곳을 나타내는 세 번째 매개변수

사용자 정의 [predicates](templates-and-predicates.md)는 이름을 대괄호로 감싸서 나타납니다. 이 시나리오에서 사용자 정의 메소드의 이름은 'SourceFoo'입니다.

이 쿼리를 사용하면 'SourceFoo'에서 시작하는 데이터 흐름을 찾을 수 있습니다. Snyk이알지 못하는 소스가 알려진 보안 취약한 크로스 사이트 스크립팅 (XSS) 싱크에 도달하고 알려진 크로스 사이트 스크립팅 (XSS) 살균화제를 거치지 않는 상태에서 데이터가 오염되었다는 가정을 할 수 있습니다.

## **데이터 흐름 규칙 확장**

Snyk 규칙을 재생성하고 현재 Snyk에서 스캔 중 이용되지 않아 검사에서 제외되고 있는 이미 알려진 보안 소스 목록에 소스를 추가하십시오. 결과인 취약점 놓침으로 이어지게 됩니다.

[신규 데이터 흐름 규칙](create-query.md#net-new-data-flow-rule)과 마찬가지로 `Or` 연산자를 사용하는 `Taint` 데이터 흐름 템플릿을 사용하게 됩니다. 쿼리에서 논리적 명제를 만드는데 사용할 수 있는 `Or` 또는 `And`와 같은 연산자가 제공됩니다.

이 쿼리를 참조하여 Snyk에서 알려진 소스 외에도 사용자 정의 소스를 포함하여 데이터 흐름 규칙을 실행하세요 [`SourceFoo`](create-query.md#user-content-fn-1)를 사용한 예시는 본문의 쿼리를 참고하세요.

```javascript
Taint<Or<PRED:AnySource,"SourceFoo">,PRED:XssSanitizer,PRED:XssSink>
```

이 쿼리를 사용하면 Snyk 알려진 소스 또는 "SourceFoo"에서 시작하는 데이터 흐름을 찾게 됩니다. Snyk이 알지 못하는 소스가 알려진 취약한 크로스 사이트 스크립팅 (XSS) 싱크에 도달하고 알려진 크로스 사이트 스크립팅 (XSS) 살균화제를 거치지 않는 상태에서 데이터가 오염되었다는 가정을 할 수 있습니다.

연산자를 사용하는 문은 모두 소괄호 안에 기입됩니다 _`< statement >`_.

## **데이터 흐름 규칙에 콘텍스트 추가**

Snyk 규칙을 재생성하고 현재 Snyk에서 알려진 취약한 소스 목록에서 소스를 제거합니다. 이는 응용 프로그램 내에서 취약한 소스가 실제로 취약하지 않음을 나타내므로 결함을 제거합니다.

[신규 데이터 흐름](create-query.md#net-new-data-flow-rule) 및 [데이터 흐름 확장](create-query.md#extend-a-data-flow-rule) 규칙과 같이 `Taint` 데이터 흐름 템플릿을 사용하게 됩니다. `And` 연산자로 사용하고 있으며, 가정이 맞는게 아닌 거짓인 경우에만 표현하는 명제적 부정 문(`Not`)을 사용합니다.

Snyk에서 알려진 소스를 사용하고 결과에서 `SnykSource`가 나오지 않도록하는 데이터 흐름 규칙을 실행하세요. 이 예시에서 `SnykSource`는 `AnySource` [프리디케이트](templates-and-predicates.md)를 사용합니다.

```javascript
Taint<And<PRED:AnySource,Not<PRED:”SnykSource”>>,PRED:XssSanitizer,PRED:XssSink>
```

이 쿼리를 사용하면 알려진 Snyk 소스에서 시작하는 데이터 흐름을 찾으면서 결과로 `SnykSource`에서 비롯되지 않는 것을 제거합니다. 취약한 크로스 사이트 스크립팅 (XSS) 싱크에 도달하고 알려진 크로스 사이트 스크립팅 (XSS) 살균화제를 거치지 않는 상태에서 데이터가 오염되었다는 가정을 할 수 있습니다.

## **높은 회수 모드**

소스 코드에서 모든 소스와 싱크를 확인하여 데이터가 어디서부터 어디로 흘러가는 지 이해하는 것이 중요합니다. 이 분석은 데이터 흐름이 존재하지 않더라도 데이터가 어디서 시작해서 어디로 끝나는 지 폭넓게 평가할 수 있게 합니다.

이 모드는 조사 도구로 자주 사용되며 코드 스택으로 깊이 들여다보고 응용 프로그램 내에서 데이터의 출처와 종단을 이해합니다.

이 분석을 시작하려면 모든 Snyk 알려진 소스와 싱크를 포함하는 `AnySource` 또는 `AnySink` 프리디케이트를 사용하십시오.

```javascript
PRED:AnySource
```

이 쿼리는 분석 중인 코드에서 알려진 모든 소스를 식별하고 강조합니다.

```
PRED:AnySink
```

이 쿼리는 코드에서 발견되는 알려진 싱크를 식별하고 강조합니다. 잠재적인 데이터 종단점의 완전한 전망을 제공합니다.

## 커버리지 평가 및 누락된 데이터 흐름 링크 식별

데이터 흐름 경로에서의 갭을 식별하는 것은 데이터의 목적지를 이해하는 데 매우 중요합니다. 특히 데이터베이스나 파일 시스템과 같이 예상한 위치에 도달하지 않는 사용자 데이터의 목적지는 식별을 통해 발견할 수 있습니다. 이러한 갭은 지원되지 않는 라이브러리 또는 프레임워크 또는 해당 구성 요소의 잘못(invalidated)을 드러낼 수 있으며, 잘못된 음성(potentially leading to false negatives)로 이어질 수 있습니다. 이 내용은 포괄적인 보안 평가 및 견고한 커버리지를 보장하기 위해 매우 중요합니다.

<details>

<summary><strong>자바 예시: 사용자 지정 WebServer와 WebServlet의 상호 작용</strong></summary>

이 자바 예시는 "WebServer"와 "WebServlet" 두 구성 요소를 보여줍니다.

* **WebServer**: 지원되지 않는 또는 프로프리어터리 컴포넌트를 사용하는 사용자 지정 HTTP 서버
* **WebServlet**: 웹 상호 작용을 위해 Java의 표준 서\`\`\`txt">



````

<div data-gb-custom-block data-tag="hint" data-style='info'>

**`CallExpression`** 및 **`HasArg1`**은 개별적으로도 사용할 수 있습니다. **`and`** 연산자를 사용하여 이를 연결하면 관계가 설정되어 결합된 상태에서 일치하려고 시도할 것입니다.

</div>

### 모든 파일 작성자 탐지

.NET에서는 [`WriteAllText`](https://learn.microsoft.com/en-us/dotnet/api/system.io.file.writealltext?view=net-7.0)를 사용하여 파일을 작성할 수 있습니다. `WriteAllLines` 및 `WriteAllBytes`도 사용할 수 있습니다.

다음과 같이 코드 스니펫을 확장할 수 있습니다.

```csharp
string userData = $"Username: {username}";
byte[] userBytes = Encoding.UTF8.GetBytes(userData);
string[] userLines = new string[] {userData};

File.WriteAllText("testFile.txt", userData);
File.WriteAllLines("testFile.txt", userLines);
File.WriteAllBytes("testFile.bin", userBytes);
````

정규 표현식을 사용하여 먼저 변형을 캡처해봅시다. 함수와 두 가지 파일 이름 변형(`.bin` 및 `.txt`)을 찾겠습니다.

```ada
Taint<
  "global::System.Console.ReadLine",
  PRED:None,
  CallExpression<
    ~"global::System\.IO\.File\.(Write|Append)All(Text|Lines|Bytes)(Async)?"
  > 
    and 
      HasArg1<"testFile.txt" or "testFile.bin">
>
```

`CallExpression`이 이제 **정규 표현식**을 포함하고 있고, `HasArg1`이 **`or`** 연산자를 사용하고 있다는 것에 주목하세요. 두 가지 방법 중 아무 것이나 선택할 수 있습니다.

마지막으로 .NET `Async` 변형과 `Append` 메서드에 대한 지원을 추가해봅시다:

```ada
Taint<
  "global::System.Console.ReadLine",
  PRED:None,
  CallExpression<
    ~"global::System\.IO\.File\.(Write|Append)All(Text|Lines|Bytes)(Async)?"
  >
    and 
      HasArg1<"testFile.txt" or "testFile.bin">
>
```

#### "중요 데이터" 정의

이전 예제에서 `ReadLine`을 민감한 데이터의 소스로 다루었습니다. 특정 민감한 필드만을 가진 간단한 객체를 고려해봅시다.

```csharp
public class MyUser
{
    public string EmailAddress;

    public string MembershipType;
}
```

이 클래스에서는 `EmailAddress`만을 고려해야 합니다.

따라서 다음 코드 스니펫이 주어진 경우:

```csharp
MyUser user = new MyUser();
user.EmailAddress = "support@snyk.io";
user.MembershipType = "SampleRole";

string sensitiveData = $"Username: {user.EmailAddress}";
string notSensitiveData = $"MembershipType: {user.MembershipType}";

File.WriteAllText("testFile.txt", sensitiveData);
File.WriteAllText("testFile.txt", notSensitiveData);
```

첫 번째 `WriteAllText` 호출은 방지되어야 하며 두 번째 호출은 허용되어야 합니다. 이를 수행하기 위해 쿼리에서 필드 이름을 지정할 수 있습니다:

```ada
Taint<
  "EmailAddress",
  PRED:None,
  CallExpression<
    ~"global::System\.IO\.File\.(Write|Append)All(Text|Lines|Bytes)(Async)?"
  >
    and 
      HasArg1<"testFile.txt" or "testFile.bin">
>
```

```
```

</details>

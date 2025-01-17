# Snyk 코드 CLI 결과 보기

Snyk CLI를 사용하여 `snyk code test` 명령의 결과에 대해 다음 작업을 수행할 수 있습니다:

* [{Snyk Code CLI 결과 분석](view-snyk-code-cli-results.md#analyze-snyk-code-cli-results): 테스트 결과 보기 및 취약점 분석.
* [심각도 수준으로 결과 필터링](view-snyk-code-cli-results.md#filter-results-by-severity-level): 터미널에 표시된 `snyk code test` 결과를 특정 심각도 수준 이상의 문제만 표시하도록 필터링.
* [테스트 결과 출력](view-snyk-code-cli-results.md#output-test-results): `snyk code test` 결과를 표준 CLI 형식으로 표시하는 대신 터미널에서 JSON 또는 SARIF 형식으로 출력.
* [테스트 결과 내보내기](view-snyk-code-cli-results.md#export-test-results): CLI 코드 결과를 JSON 또는 SARIF 형식 파일로 내보내기.

{% hint style="info" %}
`Snyk code test`에 대해 JSON 및 SARIF 형식은 동일합니다. 따라서 예시는 한 형식으로만 표시됩니다.

`Snyk-to-html`을 사용하여 [HTML 형식으로 CLI 결과 표시도 가능합니다](../cli-tools/snyk-to-html.md).
{% endhint %}

## CLI 결과 분석

CLI에서 `snyk code test` 명령을 실행한 후 테스트 결과가 표시됩니다:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FO6YZZfP8VS5Q1u0e95Sy%252FSnyk%2520Code%2520-%2520CLI%2520-%2520snyk%2520code%2520test%2520-%2520Results%2520Details%2520-%25202.png%3Falt%3Dmedia%26token%3D83f3734c-d68e-4b32-9f53-a288041b1592&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=98f856bb&#x26;sv=2" alt="CLI에서 나온 테스트 결과"><figcaption><p>CLI에서 나온 테스트 결과</p></figcaption></figure>

Snyk 웹 UI에서 문제를 무시했다면, 이러한 문제는 여전히 CLI 결과에 나타납니다. 이 페이지의 각 섹션은 표시된 결과의 한 부분을 설명합니다.

### 에서 감지한 취약점 문제 목록

{Snyk Code 테스트에서 발견된 문제 목록은 문제의 심각도 수준에 따라 구성되어 있습니다.

각 문제에 대해 다음 정보가 제공됩니다:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252F8bVxWHuUy6iMFBJi5d5p%252FSnyk%2520Code%2520-%2520CLI%2520-%2520snyk%2520code%2520test%2520-%2520Results%2520-%2520Issue%2520summary%2520-%25202.png%3Falt%3Dmedia%26token%3D3bcb4e4e-1836-4a4c-8dd6-5c83595f55f4&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=49d56075&#x26;sv=2" alt="문제에 대한 CLI 테스트 정보"><figcaption><p>문제에 대한 CLI 테스트 정보</p></figcaption></figure>

* 헤더: 문제의 심각도 수준 및 취약점 유형.
* 경로: 문제가 발견된 파일 이름 및 파일에서 문제가 발견된 라인. 이 위치 세부 정보는 문제의 Sink를 나타내며, 테스트된 저장소에서 취약점이 실행될 수 있는 위치를 의미합니다.
* 정보: 문제 데이터 흐름에 대한 설명.

`Info` 섹션에 나타나는 메시지는 웹 UI의 **Data flow** 섹션에 있는 메시지와 동일합니다:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FILtqK0UOp7yJgRPyXshY%252FSnyk%2520Code%2520-%2520CLI%2520-%2520snyk%2520code%2520test%2520-%2520Results%2520-%2520Issue%2520summary%2520-%2520In%2520the%2520UI%2520-%25202.png%3Falt%3Dmedia%26token%3D5498c20f-8e24-4d71-8df5-ca88af15512f&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=3dd1c6e0&#x26;sv=2" alt="문제에 대한 CLI 테스트 정보"><figcaption><p>문제에 대한 CLI 테스트 정보</p></figcaption></figure>

### 테스트 결과에 대한 일반 정보

테스트 결과에 대한 일반 정보는 다음과 같은 세부 정보를 포함합니다:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252Fh77DJ2P7FlMpVbYatTUq%252FSnyk%2520Code%2520-%2520CLI%2520-%2520snyk%2520code%2520test%2520-%2520Results%2520-%2520Test%2520summary%2520-%25202.png%3Falt%3Dmedia%26token%3D043a5d29-bf87-4ffe-aa76-9b9d4f16d387&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=3d0eb99d&#x26;sv=2" alt="CLI 테스트 결과 일반 정보 세부사항"><figcaption><p>CLI 테스트 결과 일반 정보 세부사항</p></figcaption></figure>

* 테스트 성공: 테스트가 완료되었는지 여부.
* 조직: 테스트가 실행된 조직의 Snyk ID 또는 내부 이름. 자세한 내용은 [CLI 테스트용 Snyk 조직 설정](set-the-snyk-organization-for-the-cli-tests.md)을 참조하십시오.
* 테스트 유형: 결과를 생성한 테스트 명령의 유형. {Snyk Code}의 경우 항상 `정적 코드 분석`**입니다.**
* 프로젝트 경로: 테스트된 저장소의 경로.

### 테스트 결과 요약

테스트 결과 요약에는 다음과 같은 세부 정보가 포함됩니다:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FzlkAy6wMvp4FcOEHJz1w%252FSnyk%2520Code%2520-%2520CLI%2520-%2520snyk%2520code%2520test%2520-%2520Results%2520-%2520Summary%2520-%25202.png%3Falt%3Dmedia%26token%3Dcff47c0b-b085-437a-8f35-3a6eedc5f7ce&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=3a710dca&#x26;sv=2" alt="문제에 대한 CLI 테스트 결과 요약"><figcaption><p>문제에 대한 CLI 테스트 결과 요약</p></figcaption></figure>

* Snyk Code에서 발견한 취약점 문제 수.
* 각 심각도 수준별 발견된 문제 수.

{% hint style="info" %}
`snyk code test` 명령에는 종료 코드가 있습니다. 이러한 코드의 정의는 [이 코드들에 대한 도움말](../../commands/code-test.md#exit-codes)을 참조하십시오. 종료 코드를 보려면 `snyk code test -d`를 실행하십시오.

모든 CLI 명령에 대한 종료 코드 요약은 [CLI 명령 및 옵션 요약](../../cli-commands-and-options-summary.md)을 참조하십시오.
{% endhint %}

## 심각도 수준별 결과 필터링

CLI 터미널에 표시되는 테스트 결과를 필터링하여 특정 심각도 수준 이상의 문제만 표시할 수 있습니다.

특정 심각도 수준을 초과하는 문제만 표시하려면 다음을 입력하십시오:

```
snyk code test <폴더/경로> --severity-threshold=<낮음|중간|높음>
```

결과에는 지정된 심각도 수준의 문제 및 해당 수준보다 더 높은 심각도 수준의 문제만 포함됩니다.

예를 들어, `snyk-goof-master` 폴더에서 8개의 문제가 발견되었는데, 그 중 4개는 높은 심각도 수준이었으며 4개는 중간 수준인 경우:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252F6vyDfGgr4fpKzbjjm4It%252FSnyk%2520Code%2520-%2520CLI%2520-%2520snyk%2520code%2520test%2520-%2520Results%2520-%2520Filter%2520Severity%2520-%2520Example%2520-%2520before%2520-%25202.png%3Falt%3Dmedia%26token%3D3aff742c-0b43-4df9-b0e1-965924d10799&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=b1f47ae2&#x26;sv=2" alt="의 높고 중간 심각도 문제에 대한 CLI 테스트 결과"><figcaption><p>Snyk Code의 높고 중간 심각도 문제에 대한 CLI 테스트 결과</p></figcaption></figure>

높은 심각도 수준 이상의 문제만 표시하려면 다음을 입력하십시오:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --severity-threshold=high
```

결과는 높은 심각도 수준인 4개의 문제가 모두 표시됩니다. 낮은 심각도 수준의 문제는 표시되지 않습니다:

![높은 심각도 수준의 Snyk Code CLI 결과](https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FXy4FUu2zOXs4ktKj9vGG%252FSnyk%2520Code%2520-%2520CLI%2520-%2520snyk%2520code%2520test%2520-%2520Results%2520-%2520Filter%2520Severity%2520-%2520Example%2520-%2520after%2520-%25202.png%3Falt%3Dmedia%26token%3D63042f8f-87be-4136-a02a-9dc035143338\&width=768\&dpr=4\&quality=100\&sign=787fb82d\&sv=2)

## JSON 및 SARIF 파일에서 심각도 수준

`snyk code test`를 실행하여 발견된 문제의 심각도 수준은 JSON 및 SARIF 파일에서 다르게 표시됩니다. JSON 및 SARIF 결과에서 사용되는 심각도 수준은 다음과 같습니다:

* 높음 = **에러**
* 중간 = **경고**
* 낮음 = **노트/정보**

크리티컬 표시는 {Snyk Code}에서 사용되지 않습니다.

터미널에 표시된 중간 심각도 예시는 다음과 같습니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - JSON and SARIF - Severity Level Results - in the Terminal.png" alt="JSON 또는 SARIF 출력의 중간 심각도"><figcaption><p>JSON 또는 SARIF 출력의 중간 심각도</p></figcaption></figure>

다음은 파일에서 높은 및 낮은 심각도 수준 예시를 보여줍니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - JSON and SARIF - Severity Level Results.png" alt="JSON 또는 SARIF 파일의 높은 및 낮은 심각도 수준"><figcaption><p>JSON 또는 SARIF 파일의 높은 및 낮은 심각도 수준</p></figcaption></figure>

## 테스트 결과 출력

`Snyk code test` 결과를 표준 CLI 형식으로 표시하는 대신 터미널에서 [JSON](view-snyk-code-cli-results.md#output-test-results-in-json-format) 또는 [SARIF](view-snyk-code-cli-results.md#output-test-results-in-sarif-format) 형식으로 출력할 수 있습니다.

또한 테스트 결과를 JSON 또는 SARIF 형식 파일로 [내보낼 수도 있습니다](view-snyk-code-cli-results.md#export-test-results). SARIF는 정적 분석 도구의 출력을 위한 오픈 표준입니다. 자세한 정보는 [SARIF 사이트](https://sarifweb.azurewebsites.net/)를 참조하십시오.

`snyk code test`를 실행하여 발견된 문제의 심각도 수준 및 JSON 및 SARIF 파일에서보고된 결과가 터미널 결과와 다르게 표시됩니다. 자세한 내용은 [JSON 및 SARIF 파일에서의 심각도 수준](view-snyk-code-cli-results.md#severity-levels-in-json-and-sarif-files)을 참조하십시오.

테스트 결과를 JSON 형식으로 출력하려면 다음을 입력하십시오:

```
snyk code test <폴더/경로> --json
```

테스트 결과를 SARIF 형식으로 출력하려면 다음을 입력하십시오:

```
snyk code test <폴더/경로> --sarif
```

테스트 결과는 터미널에서 JSON 또는 SARIF 형식으로 표시됩니다.

JSON 및 SARIF는 `snyk code test`에 대해 동일하므로 여기에서는 JSON 예시만 표시됩니다. `snyk-goof-master` 폴더의 테스트 결과를 JSON 형식으로 터미널에 출력하려면 다음 명령을 사용하여 리포지토리의 루트 폴더로 이동하고 다음과 같이 명령을 입력합니다:

```
snyk code test /Users/username/Documents/Repositories/snyk-goof-master --json
```

테스트 결과는 터미널에서 JSON 형식으로 표시됩니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - results - JSON output in the terminal.png" alt="JSON 형식의 snyk code test 결과"><figcaption><p><code>snyk code test</code> 결과를 JSON 형식으로 출력</p></figcaption></figure>

## 테스트 결과 내보내기

`Snyk code test` 결과를 JSON 또는 SARIF 형식 파일로 내보낼 수 있습니다. 결과를 내보낼 때 새 파일의 이름을 제공해야 합니다.

테스트 결과를 JSON 또는 SARIF 파일로 [터미널에서 출력](view-snyk-code-cli-results.md#output-test-results)하는 것도 가능합니다.

`snyk code test`를 실행하여 발견된 문제의 심각도 수준 및 JSON 및 SARIF 파일에서보고된 결과가 터미널 결과와 다르게 표시됩니다. 자세한 내용은 [JSON 및 SARIF 파일에서의 심각도 수준](view-snyk-code-cli-results.md#severity-levels-in-json-and-sarif-files)을 참조하십시오.

두 가지 방법으로 결과를 JSON 또는 SARIF 파일로 내보낼 수 있습니다. 다음 지침은 JSON 파일을 기준으로 설명하지만 SARIF 파일도 내보낼 수 있습니다.

### 표준 결과가 터미널에 표시되는 새 파일로 결과 내보내기

`Snyk code test --json-file-output=<새_파일/경로>` 명령은 Snyk CLI v. 1.910.0 이상에서 사용할 수 있습니다. Snyk CLI 버전을 업데이트하려면 [Snyk CLI 설치 또는 업데이트](../../install-or-update-the-snyk-cli/)를 참조하십시오.

새 JSON 파일로 결과를 내보내려면, 터미널에서 다음 명령을 사용합니다:

```
snyk code test --json-file-output=<새_json_파일/경로>
```

새로운 SARIF 파일로 테스트 결과를 내보내려면 다음 명령을 사용합니다:

```
snyk code test --sarif-file-output=<새_sarif_파일/경로>
```

결과가 터미널에 표시되고 지정한 경로에 JSON 또는 SARIF 파일이 만들어집니다.

JSON 및 SARIF는 `snyk code test`에 대해 동일하므로 여기에서는 JSON 예시만 표시됩니다. `snyk-goof-master` 폴더의 테스트 결과를 `json`이라는 JSON 파일로 내보내려면 디렉토리를 루트 폴더로 변경하고 다음과 같이 명령하십시오:

```
snyk code test --json-file-output=json
```

터미널에서 코드 테스트 결과가 표준 형식으로 표시됩니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - results - export to JSON - with terminal results - 2 .png" alt="터미널에서 snyk code test 결과"><figcaption><p><code>snyk code tes</code>t 결과 터미널에 표시</p></figcaption></figure>

리포지토리 폴더에 JSON 파일이 생성됩니다:

<figure><img src="../../../.gitbook/assets/snyk Code - CLI - results - export to JSON - with terminal results - JSON file.png" alt="리포지토리의 JSON 파일"><figcaption><p>리포지토리의 JSON 파일</p></figcaption></figure>

### 터미널 결과가 표시되지 않고 새 파일로 결과 내보내기

터미널에 결과를 표시하지 않고 새 JSON 파일로 결과를 내보내려면 다음

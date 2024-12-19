# snyk-to-html

CLI는 빌드를 실패로 만드는 직접 또는 자동화된 방법을 제공하며, 기본적으로는 `--json` 또는 `--sarif` 형식을 사용하지 않는 한 요약 정보만 제공합니다. 이 출력을 파일로 지정할 수 있고, 이 파일에는 발견된 문제가 포함됩니다. 이러한 형식은 사람이 읽기 어렵습니다.

`Snyk` JSON을 HTML로 매핑하는 `snyk-to-html` (Snyk JSON to HTML Mapper)를 사용하여 다음을 수행할 수 있습니다:

- 스캔 결과의 출력 부분으로 읽을 수 있는 빌드 아티팩트를 생성
- 브라우저에서 볼 수 있는 취약점의 전체 컨텍스트를 HTML 보고서로 생성

이 페이지에서는 `snyk-to-html`을 [설치하는 방법](snyk-to-html.md#install-snyk-to-html), `snyk-to-html` 도구를 사용하여 테스트 결과를 JSON 또는 SARIF 형식으로 HTML 파일로 내보내는 방법 및 [브라우저에서 테스트 결과보기](snyk-to-html.md#view-test-results-in-html-format)에 대해 설명합니다.

{% hint style="warning" %}
참고: 오픈 소스 (SCA) 스캔에 대해서는 `json` 출력만 지원됩니다. `snyk-to-html` 프로세스에서 `sarif` 출력은 오픈 소스 테스트에서 결과가 반환되지 않습니다.&#x20;
{% endhint %}

## `snyk-to-html` 설치

`snyk-to-html`을 설치하려면 컴퓨터에서 필요한 권한이 있어야합니다. 설치가 실패하면 필요한 권한을 요청하기 위해 IT 관리자에게 문의하십시오.

다음 명령을 사용하여 `snyk-to-html`을 설치할 수 있습니다:

```
npm install snyk-to-html -g
```

지역적으로 `snyk-to-html` 플러그인을 설치하려면 [snyk-to-html GitHub 저장소를 복제](https://github.com/snyk/snyk-to-html)하고 다음 스크립트를 사용하십시오:

```
npm install
npm run build
node ./dist/index.js
```

## `snyk-to-html` 도구 사용

`snyk test` 명령의 일부로 `snyk-to-html`을 실행하여 테스트의 출력 부분으로 읽을 수 있는 빌드 아티팩트를 만들 수 있습니다.

또한 `snyk test` 명령을 실행하면 결과를 JSON 파일로 내보내고 그 파일을 `snyk-to-html`을 사용하여 HTML로 변환할 수 있습니다. `Snyk` 코드 결과를 JSON 또는 SARIF 파일로 내보내고 그 파일을 HTML로 변환할 수 있습니다.

`snyk-to-html` 명령을 실행할 때 다음 명령 옵션을 사용하여 사용자 정의할 수 있습니다:

<table data-header-hidden><thead><tr><th width="105"></th><th width="134"></th><th></th><th></th></tr></thead><tbody><tr><td><strong>짧게</strong></td><td><strong>길게</strong></td><td><strong>설명</strong></td><td><strong>기본값</strong></td></tr><tr><td><code>-i</code></td><td><code>--input</code></td><td>테스트 결과를 포함하는 JSON 또는 SARIF 파일의 입력 경로. SARIF 형식은 오픈 소스 스캔 결과에 대해 지원되지 않습니다.</td><td><code>stdin</code></td></tr><tr><td><code>-o</code></td><td><code>--output</code></td><td><p>HTML 결과 파일의 이름 앞에 오는 출력 파일 이름 지정.</p><p>예:<br><code>-o results.html</code></p></td><td><code>stdout</code></td></tr><tr><td><code>-t</code></td><td><code>--template</code></td><td>HTML 생성을위한 템플릿 위치.</td><td><code>template/test-report.hbs</code></td></tr><tr><td><code>-s</code></td><td><code>--summary</code></td><td>세부 보고서가 아닌 요약 정보만 있는 HTML 파일을 생성합니다.</td><td>세부 취약점 보고서</td></tr><tr><td><code>-a</code></td><td><code>--actionalable-remediation</code></td><td>사용 가능한 치료 정보를 표시합니다.</td><td>해당 없음</td></tr><tr><td><code>-d</code></td><td><code>--debug</code></td><td>디버그 모드에서 명령을 실행합니다.</td><td>해당 없음</td></tr></tbody></table>

`snyk-to-html` 명령은 표준 종료 코드를 생성하지 않습니다.

`snyk-to-html`에 대한 도움말을 표시하려면 `snyk-to-html --help` 또는 `--h` 명령을 사용하십시오.

`CI/CD 파이프라인`에서 `snyk-to-html` 명령을 사용하려면 [Snyk CI/CD 통합 예제](https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/AzurePipelines/AzurePipelines-npm-generic-html.yml)를 참조하십시오.

자세한 내용은 [사용 `snyk-to-html` 명령 옵션](snyk-to-html.md#use-snyk-to-html-command-options)을 참조하십시오.

### 테스트의 출력 부분으로 읽을 수 있는 빌드 아티팩트 생성

`snyk test` 명령의 일부로 `snyk-to-html`을 실행하여 빌드 아티팩트를 만드는 단계는 다음과 같습니다. 이를 통해 결과가 `snyk-to-html`로 직접 스트리밍됩니다.

1. 테스트할 리포지토리의 루트 폴더로 디렉토리 변경
2. 리포지토리를 테스트하려면 결과를 JSON 형식으로 내보내고, 다음 예제 중 하나의 이름으로 HTML 파일이 리포지토리 폴더에 생성되도록 플러그인을 사용하십시오: `results-[scan type].html`.

각 Snyk 스캔 방법에 대한 사용할 명령어는 다음과 같습니다. 명령을 실행하면 이러한 예제 중 하나의 이름으로 HTML 파일이 리포지토리 폴더에 생성되어 [HTML 형식에서 테스트 결과를 볼 수 있습니다](snyk-to-html.md#view-test-results-in-html-format).

#### Snyk 오픈 소스 명령어

다음 명령을 실행하여 `results-opensource.html`이라는 파일이 생성됩니다:

`snyk test --json | snyk-to-html -o results-opensource.html`

#### Snyk 컨테이너 명령어

다음을 실행하여 `results-container.html`이라는 파일을 생성합니다:

`snyk container test [이미지] --json | snyk-to-html -o results-container.html`

#### **Snyk 코드 명령어**

다음을 실행하여 `results-code.html`이라는 파일을 생성합니다:

`snyk code test --json | snyk-to-html -o results-code.html`

#### Snyk IaC 명령어

관련된 파일이 있는 하위 폴더로 이동한 다음 다음 명령을 실행하여 `results-iac.html`이라는 파일이 생성됩니다:

`snyk iac test --json | snyk-to-html -o results-iac.html`

### JSON 또는 SARIF 파일을 HTML로 변환하여 브라우저에서 볼 수 있게 하기

자동화 목적을 위해 결과에 대한 프로그래밍적 접근을 위해 JSON 파일을 생성하거나 이전 스캔에서 이미 JSON 파일을 가지고 있는 경우가 있습니다. 이 JSON 출력을 `snyk-to-html`에 보내어 HTML 파일을 생성할 수 있습니다.

다음 단계를 따라 `snyk test`를 실행하고 그 결과 파일을 HTML로 변환하십시오.

1. 테스트하려는 리포지토리의 루트 폴더로 디렉토리 변경
2. 각 제품에 대해 아래와 같이 해당 `test` 명령을 실행하십시오:\
    `snyk test --json-file-output=results-opensource.json`

    `snyk code test --json-file-output=results-code.json`

    `snyk container test [이미지] --json-file-output=results-container.json`

    `snyk iac test  --json-file-output=results-iac.json`

    \
    출력을 도구로 파이핑하기 전에 프로세스가 중지되는 종료 코드가 있으면 단계 뒤에 따라오는 참고 사항을 참조하십시오.
3. JSON 파일을 HTML로 변환하려면 `snyk-to-html`에 전달하십시오. 입력 파일은 유효한 JSON을 사용하고 UTF-8 인코딩을 사용해야 합니다. 생성된 출력 파일 이름을 사용하십시오:\
   `snyk-to-html -i results-opensource.json -o results-opensource.html`\
   `snyk-to-html -i results-code.json -o results-code.html`\
   `snyk-to-html -i results-container.json -o results-container.html`\
   `snyk-to-html -i results-iac.json -o results-iac.html`

{% hint style="info" %}
`snyk test --json > result-opensource.json`과 같이 다단계 접근 방식을 사용하고 결과를 플러그인으로 전달하는 경우, [exit code](../../cli-commands-and-options-summary.md#exit-codes-for-cli-commands)가 프로세스를 종료하거나 막을 수 있기 때문에 빌드 시스템에서 해당 단계로 넘어가기 전에 프로세스가 중지될 수 있습니다. 빌드 도구의 기능에 따라 여러 옵션이 있습니다:\
\
1\) 반환되는 프로세스에 추가되지 않도록 매개변수에 [exit code](../../cli-commands-and-options-summary.md#exit-codes-for-cli-commands)를 캡처하여 오류 상태를 확인하십시오.\
2\) `||true` 또는 유사한 논리를 사용하여 프로세스를 종료하지 않고 [exit code](../../cli-commands-and-options-summary.md#exit-codes-for-cli-commands)를 방지하십시오. 이렇게 하면 네트워크나 Snyk 플랫폼 문제 또는 다른 비스캔 결과 문제를 나타내는 오류 코드와 같이 반환 코드가 무시되며, 다음 단계에서 JSON을 사용하는 경우 실패할 가능성이 높습니다. 스크립트의 다음 단계로 이동하기 전에 종료 코드를 검토하는 것이 좋습니다.\
3\) 오류 발생 시 계속하도록 단계를 설정하십시오, 해당 옵션이 있는 경우.
{% endhint %}

### `snyk-to-html` 명령 옵션 사용

다음 예제는 `snyk test` 명령으로 보여줍니다. 그러나 컨테이너, 코드 및 IaC를 위한 `snyk test` 명령과도 작동합니다.

#### 보고서의 단순 버전 표시

`-s` 또는 `--summary` 옵션을 사용하여 보고서의 요약만 표시할 수 있습니다.

`snyk-to-html -i results.json -o results.html -s`

#### 사용 가능한 치료법 표시

`-a` 또는 `--actionable-remediation` 옵션을 사용하여 취약점을 치료하는 조치를 표시할 수 있습니다.

`snyk-to-html -i results.json -o results.html -a`

보고서에서 치료, 업그레이드 및 패치가 취약점을 고친 수의 순서에 따라 배열됩니다. 패키지를 업그레이드하고 패치할 순서를 선택하는 경우 이를 가이드로 사용하십시오.

Snyk는 다음 패키지 관리자에 대한 치료 조언을 지원합니다:

- npm
- Yarn
- RubyGems
- Maven
- Gradle
- sbt
- Pip

## HTML 형식에서 테스트 결과 보기

HTML 파일을 보려면 리포지토리의 출력 파일을 찾아 더블 클릭하십시오. HTML 파일에 다른 이름을 사용했다면 해당 파일을 찾아 열어 보십시오.

테스트 결과 보고서가 브라우저에서 열립니다. 다음 예제는 `snyk code test` 결과를 보여줍니다. 각 문제를 클릭하여 문제별로 **데이터 플로우** 및 **고치기 분석** 정보를 볼 수 있습니다.

<figure><img src="../../../.gitbook/assets/Snyk-to-HTML - Example - HTML Report - Fix Analysis tab - 2.png" alt="Snyk Code 보고서의 데이터 플로우와 Fix Analysis 버튼 강조"><figcaption><p>Snyk Code 보고서에서 문제별로 데이터 플로우와 Fix Analysis 버튼을 강조</p></figcaption></figure>

## 라이선스

[라이선스: Apache License, Version 2.0](https://github.com/snyk/snyk-to-html/blob/master/LICENSE)
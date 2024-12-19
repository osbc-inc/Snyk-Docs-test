# SDK 참조

## 이름

`snyk-iac-rules` - {{Snyk IaC}}를 위한 사용자 지정 규칙을 작성, 디버그, 테스트 및 번들로 제작하는 SDK

## 요약

`snyk-iac-rules` \[COMMAND] \[PATH] \[OPTIONS]

## 설명

이 SDK는 Rego로 작성된 사용자 지정 규칙을 작성, 디버그, 테스트, 번들 제작 및 배포하는 데 도움이 되며, 이러한 규칙은 {{Snyk IaC}} CLI에서 IaC 구성 파일의 취약점을 찾을 수 있도록 사용할 수 있습니다.

Snyk에 대한 자세한 정보는 [Snyk 웹사이트](https://snyk.io)를 참조하십시오.

## 시작할 위치가 없는 경우?

1. 새 규칙 작성을 위해 `$ snyk-iac-rules template` 실행
2. Rego 작성을 시작하려면 `$ snyk-iac-rules parse`로 픽스처를 JSON으로 파싱하세요.
3. `$ snyk-iac-rules test`로 규칙을 테스트하세요.
4. `$ snyk-iac-rules build`로 규칙을 번들로 묶으세요.
5. `$ snyk-iac-rules push`로 규칙을 배포하세요.

## 명령어

명령에 대한 플래그 및 사용법을 보려면, 예를 들어, `snyk-iac-rules --help`를 실행하십시오.

다음 최상위 SDK 명령어가 제공됩니다.

`template` - 새 규칙을 작성하는 뼈대 생성. 자세한 지침은 `snyk-iac-rules template --help`를 참조하십시오.

`parse` - 제공된 픽스처 파일의 JSON 형식 반환. Rego는 JSON 입력이 필요하므로 이를 사용하여 Rego 규칙을 작성합니다. 자세한 지침은 `snyk-iac-rules parse --help`를 참조하십시오.

`test` - 일치하는 파일에서 발견된 모든 테스트 케이스 실행. 자세한 지침은 `snyk-iac-rules test --help`를 참조하십시오.

`build` - OPA 정책 및 데이터 파일을 번들로 묶습니다. 자세한 지침은 `snyk-iac-rules build --help`를 참조하십시오.

`push` - `build` 명령으로 생성된 번들을 지원되는 컨테이너 레지스트리 중 하나로 배포합니다. 자세한 지침은 `snyk-iac-rules push --help`를 참조하십시오.

## 경로

SDK의 모든 명령어는 폴더 경로를 선택적으로 입력받을 수 있으며, 규칙이 위치할 폴더를 지정합니다. `parse` 및 `push` 명령어를 제외한 모든 명령어에서는 이 경로가 엄격히 필요합니다.

상대 경로 또는 절대 경로를 제공할 수 있습니다.

## 옵션

명령어별 플래그 및 사용법을 보려면 `help` 명령을 실행하십시오. 예를 들어, `snyk-iac-rules template --help`를 실행하싀텨 됩니다.

### 템플릿 옵션

`--rule`=RULE

정의하려는 규칙의 필수 이름. 이는 최상위에 `rules/` 폴더를 생성하고 해당 규칙의 이름을 가진 폴더와 Rego 규칙 및 Rego 테스트 파일을 포함합니다. 동시에, `data.lib`에서 규칙 작성 및 테스트 라이브러리 확장에 액세스할 수 있는 유틸리티 함수가 포함된 `lib/` 폴더가 생성됩니다.

생성된 폴더 구조는 다음과 같습니다:

`rules`\
`└── RULE`\
`├── fixtures`\
`├── allowed.<extension>`\
`├── denied.<extension>`\
`├── main.rego`\
`├── main_test.rego`\
`lib`\
`└── testing`\
`└── main.rego`\
`└── tfplan.rego`\
`└── main.rego`

참고: 규칙 이름에는 공백이 포함되어서는 안 되며, `SNYK-`로 시작해서도 안 됩니다.

`--format`=`hcl2`|`json`|`yaml`|`tf-plan`

규칙을 정의할 필수 구성 형식. 이는 규칙이 작동을 확인하는 데 사용되는 `rules/<RULE>/fixtures` 폴더 아래에 두 개의 픽스처 파일이 생성됩니다.

`--severity`=`low`|`medium`|`high`|`critical`

규칙의 심각도는, Snyk Infrastructure as Code CLI를 실행할 때 표시됩니다.

기본값: `low`

`--title`=`TITLE`

규칙의 제목은, Snyk Infrastructure as Code CLI를 실행할 때 표시됩니다.

### 파싱 옵션

`--format`=`hcl2`|`yaml|tf-plan`

픽스처의 형식. 형식은 Snyk가 픽스처에서 JSON을 생성하는 데 사용하는 파서를 결정합니다.

기본값: `hcl2`

### 테스트 옵션

`--verbose`

추적 로그 출력합니다.

`--explain`=`full`|`notes`|`fails`

추적 로그 필터링합니다.

기본값: `fails`

`--timeout`=TIMEOUT

테스트 실행을 기다리는 시간(나노초)입니다. TIMEOUT을 초과하면 테스트가 실패합니다.

기본값: 5000000 (5초).

`--ignore`

테스트에로드할 파일 및 폴더를 무시하고 로드되지 못하게하는 데 사용할 수있는 정규 표현식을 수용합니다.

기본값: ".\*" (숨겨진 파일), "fixtures"

`--run`

테스트의 하위 집합으로 실행할 수 있는 정규 표현식을 수용합니다.

### 빌드 옵션

`--output`

결과 번들의 이름 및 위치입니다.

기본값: bundle.tar.gz

`--entrypoint`

기본적으로 Template 명령은 규칙을 위한 Rego를 `./rules/<rule>` 폴더에 배치합니다. Rego는 규칙 패키지에 속하며, 발견된 경우 잘못된 구성을 나타내는 JSON을 반환하는 `deny`라는 진입 함수를 호출합니다. 이 구조는 `rules/deny` 진입점과 함께 작동하되, 생성된 파일과 패키지 구조를 수정해야 하는 경우 사용자 정의 진입점을 제공할 수 있습니다.

기본값: "rules/deny"

`--ignore`

무시하고 번들링하지 못하도록 파일과 폴더를 로드하는 데 사용할 수 있는 정규 표현식을 수용합니다.

기본값: ".\*" (숨겨진 파일), "fixtures", "\*\_test.rego""

`--target`=`rego`|`wasm`

번들에 사용할 형식입니다. Snyk는 현재 {{Snyk IaC}} CLI에서 Rego 번들을 지원하지 않습니다.

기본값: `wasm`

### 푸시 옵션

`--registry`

번들을 푸시할 레지스트리 위치입니다. 예: `docker.io/`example`/bundle.tar.gz`

### 모든 명령어에서 사용 가능한 플래그

\[COMMAND] `--help`, `--help` \[COMMAND], `-h`

도움말 텍스트를 출력합니다. 자세한 내용을 얻으려면 COMMAND를 지정할 수 있습니다.

## SDK 예제

이름이 `CUSTOM_RULE`인 새 규칙 생성:

```
$ snyk-iac-rules template --rule CUSTOM_RULE
```

규칙 테스트:

```
$ snyk-iac-rules test --run CUSTOM_RULE
```

사용자 지정 규칙 번들 생성:

```
$ snyk-iac-rules build --output bundle-v1.0.0.tar.gz
```

## SDK 종료 코드

가능한 종료 코드 및 의미:

**0**: 성공

**1**: 실패, 사용자 지정 규칙이 잘못될 수 있음
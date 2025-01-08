# Go

## 적용 범위

Snyk의 Go 버전은 Snyk Open Source 및 Snyk Code를 지원합니다.

언어 사용 가능 여부를 확인하고 Snyk 제품을 사용하여 가져올 수 있는 애플리케이션으로 테스트하거나 모니터링할 수 있습니다.

사용 가능한 기능:

* SCM 가져 오기, Snyk Open Source 및 Snyk Code에서 사용 가능합니다.
* CLI 및 IDE를 통해 앱을 테스트하거나 모니터링합니다. 이는 Snyk Open Source 및 Snyk Code에서 사용 가능합니다.
* `pkg:golang`를 사용하여 앱의 SBOM을 테스트합니다.
* `pkg:golang`를 사용하여 앱의 패키지를 테스트합니다.

## 패키지 관리자 및 지원되는 파일 확장자

Go용 Snyk은 Go Modules 및 dep를 패키지 관리자로 지원하며, 패키지 레지스트리는 여러 원본을 사용합니다.

Snyk Go는 다음 파일 형식을 지원합니다:

* Snyk Open:
  * Go Modules: `go.mod`
  * dep: `gopkg.lock`
* Snyk Code: `.go`

## 프레임워크 및 라이브러리

다음 프레임워크와 라이브러리가 Snyk Go에서 지원됩니다:

* Azure/azure-sdk-for-go/sdk/ai/azopenai - 포괄적
* gage-technologies/mistral-go - 포괄적
* Gin - 부분적
* Go 표준 라이브러리 - 포괄적
* google/generative-ai-go/genai - 포괄적
* GORM 라이브러리 - 부분적
* labstack/echo - 부분적
* sashabaranov/go-openai - 포괄적
* spf13/pflag - 포괄적

## 기능

다음 기능은 Snyk Go에서 지원됩니다:

| Snyk Open Source                      | Snyk Code                                               |
| ------------------------------------- | ------------------------------------------------------- |
| <ul><li>라이선스 스캔</li><li>보고서</li></ul> | <ul><li>보고서</li><li>사용자 지정 규칙</li><li>파일 간 분석</li></ul> |

# 액세스 요구 사항

[CLI ](../getting-started-with-the-snyk-cli.md)또는 [IDE 통합](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)과 같은 Snyk 애플리케이션을 사용할 때는 일부 로컬 및 원격 리소스에 액세스해야 합니다. 이 문서에서는 Snyk 기능에 영향을주지 않고 시스템을 견고하게 하는 방법에 대해 설명합니다.

## 로컬 파일 시스템

필요한 파일 시스템 액세스는 제품에 따라 다를 수 있습니다.

* CACHE\_PATH (읽기,쓰기,실행)
  * Windows: `C:\\Users\\<USERNAME>\\AppData\\Local\\snyk`
  * Linux/Alpine: `/home/<USERNAME>/.cache/snyk`
  * macOS: `/Users/<USERNAME>/Library/Caches/snyk`
* CONFIG\_PATH (읽기,쓰기)
  * Windows: `%USERPROFILE%\\.config\\configstore`
  * Linux: `$HOME/.config/configstore`
  * macOs: `$HOME/.config/configstore`

## 네트워크 리소스

### 필수

* api.\<SNYK\_INSTANCE>.io:443
* app.\<SNYK\_INSTANCE>.io:443

### 선택적

* deeproxy.snyk.io:442 (snyk code를위한)
* downloads.snyk.io:443 (CLI 다운로드와 같은 기능에 따라 다를 수 있음)
* learn.snyk.io:443 (문제 세부 정보에 Snyk Learn 링크를 표시하려면)
* static.snyk.io:443 (CLI 다운로드와 같은 기능에 따라 다를 수 있음)
* snyk.io:443 (사용된 기능에 따라 다를 수 있음)
* \*_.sentry.io:443 (에러 보고)_
* \*_.amplitude.com:443 (분석)_
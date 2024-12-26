# Snyk 코드에서 누락된 또는 다른 결과

## .gitignore 및 .dcignore 파일 <a href="#gitignore-and-.dcignore-files" id="gitignore-and-.dcignore-files"></a>

* 모든 `.gitignore` 및 `.dcignore` 파일의 이름을 예를 들어 `.dcignore.bak.`로 변경합니다.
* 를 다시 실행합니다.
* 결과가 더 나아졌다면, `.gitignore` 및 `.dcignore` 파일에서 무시되는 파일을 확인하기 위해 규칙을 확인합니다.
* JetBrains에서 디버그 로깅을 활성화하여  로거에서 무시를 분석할 수 있습니다.
* 언어 서버에서 추적 로깅을 활성화하여 무시를 분석할 수 있습니다. 예를 들어, 환경 변수 `SNYK_DEBUG_LEVEL=trace`를 설정합니다.
* 올바른 결과가 나타날 때까지 무시 패턴을 단계별로 활성화하고 무시해야 할 파일만 무시되어야 합니다.
* `.dcignore`에 많은 무시 항목이 포함되어 있고 자동으로 생성되었으면 이를 비워야 합니다. IntelliJ는 2.4.63 이전에 최상위 수준에 `.gitignore` 파일이 없을 때 이 파일을 자동으로 추가했습니다. 이 파일에는 저장소 및 소스 파일과 간섭할 수 있는 많은 규칙이 포함되어 있습니다.

##  품질 <a href="#snyk-code-quality" id="snyk-code-quality"></a>

* IDE 플러그인 및 언어 서버가  품질 결과를 보고합니다.
* CLI는  품질 결과를 보고하지 않습니다.
* 이로 인해 보고된 문제가 다를 수 있습니다.
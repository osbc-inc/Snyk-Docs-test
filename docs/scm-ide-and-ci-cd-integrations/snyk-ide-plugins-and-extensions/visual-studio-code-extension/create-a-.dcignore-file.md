# .dcignore 파일 생성

특정 파일 및 디렉토리(예: **node_modules**)를 무시하려면 **.dcignore** 파일을 생성하십시오. 이 파일은 프로젝트가 위치한 디렉터리를 기준으로 어떤 수준의 디렉터리에서도 생성할 수 있습니다. 파일 구문은 `.gitignore`와 동일합니다.

* Snyk은 `.gitignore` 파일이 없을 때 파일을 추가하는 것을 권장합니다. 이 파일을 추가하면 업로드해야 하는 파일 수가 크게 줄어들어 분석 속도가 향상됩니다.
* **`.dcignore`** 파일을 빠르게 추가하려면 VS Code 및 Snyk 확장 기능에서 제공하는 명령 **Snyk create `.dcignore` file**을 사용하고 새로 생성된 `.dcignore` 파일을 저장하세요.
# Windows 시스템에서 .exe 다운로드 차단으로 인해 IDE 플러그인이 "Scan failed" 오류가 발생하는 문제

## **문제**

Windows 시스템에서 IDE에 Snyk 플러그인을 사용할 때 Snyk 스캔이 실패할 수 있습니다.&#x20;

로그에서 'Error running command snyk.start: command 'snyk.workspace.scan' not found'라는 오류 메시지가 표시될 수 있습니다. 이는 'snyk.start'를 기여하는 확장 프로그램에 의해 발생했을 가능성이 있습니다.

## **원인**

이 문제는 자동 .exe 다운로드를 방지하기 위해 방화벽이나 다른 도구에서 대화식 팝업을 표시하여 회사 정책을 가진 Windows 시스템에서 종종 발생합니다.&#x20;

IDE에서는 이 팝업이 표시되지 않으며 .exe 파일을 다운로드할 수 없습니다. 이로 인해 다운로드 위치에 0KB 파일이 생성될 수 있습니다.&#x20;

이를 확인하려면 IDE에서 플러그인 설정에 표시된 CLI 또는 Language Server 위치로 이동하세요. CLI 또는 Language Server 파일 중 하나가 0KB이면 이 문제가 발생할 수 있습니다.

## **해결책**

이 문제에 대한 기본 해결책은 모든 사용자에 대해 .exe 팝업 승인을 예외로 처리할 수 있도록 [downloads.snyk.io](https://downloads.snyk.io)와 [static.snyk.io](http://static.snyk.io/)를 .exe 파일의 기본 정책에 허용목록에 추가하는 것입니다.&#x20;

원하는 해결책이 승인되기까지 오랜 시간이 걸릴 경우 임시적으로 다음 해결책을 사용할 수 있습니다.&#x20;

* Snyk CLI 릴리스(https://github.com/snyk/cli/releases) 또는 Snyk Language Server 릴리스(https://github.com/snyk/snyk-ls/releases)에서 관련 .exe 파일을 직접 수동으로 다운로드하고 플러그인 설정에 나열된 위치에 배치합니다. 자동 의존성 관리 옵션을 꺼두는 것을 잊지 마세요.

또는

* CLI .exe 파일이 다운로드된 경우 이를 사용하여 Language Server 기능을 제공할 수 있습니다. 플러그인 설정에서 언어 서버 위치를 CLI 위치로 리디렉션하고 자동 의존성 관리 옵션을 꺼두세요. &#x20;
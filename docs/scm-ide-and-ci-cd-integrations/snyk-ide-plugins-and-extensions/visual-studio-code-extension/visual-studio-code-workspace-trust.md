# Workspace trust

취약점을 조사하기 위한 일환으로, Snyk는 분석을 위해 추가 데이터를 얻기 위해 자동으로 컴퓨터에서 코드를 실행할 수 있습니다. 이는 패키지 관리자(예: pip, gradle, maven, yarn, npm 등)를 호출하여 Snyk Open Source를 위한 의존성 정보를 가져오는 것을 포함합니다. 악의적인 구성을 가진 신뢰할 수 없는 코드에서 이러한 프로그램을 호출하는 것은 시스템을 악성 코드 실행과 악용에 노출시킬 수 있습니다.

Visual Studio Code의 내장 [Workspace Trust](https://code.visualstudio.com/docs/editor/workspace-trust) 기능 외에도, 저희의 확장 프로그램은 코드를 스캔하기 전에 폴더 신뢰를 요청할 것입니다. 의심스러운 경우, 스캔을 진행하지 마십시오.

<figure><img src="../../../.gitbook/assets/vscode-trust (1) (1).png" alt=""><figcaption><p>폴더 신뢰를 요청하는 화면</p></figcaption></figure>

한 번 프로젝트 신뢰를 부여하면, Snyk는 열린 폴더 및 하위 폴더에 대해 더 이상 신뢰를 요청하지 않습니다. 처음에 신뢰를 부여하지 않은 경우, 플러그인은 다음에 Visual Studio Code 인스턴스를 다시 시작할 때 다시 요청할 것입니다.

기존 폴더 신뢰를 취소하거나 더 많은 폴더에 신뢰를 부여하려면, Snyk 확장 프로그램 설정으로 이동하여 **Trusted Folders** (`snyk.trustedFolders`) 설정을 편집할 수 있습니다.
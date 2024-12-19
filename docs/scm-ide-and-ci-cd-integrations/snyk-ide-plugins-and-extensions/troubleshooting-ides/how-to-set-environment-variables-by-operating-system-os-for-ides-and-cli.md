# IDE 및 CLI를 위한 운영체제 OS별 환경 변수 설정 방법

## Windows <a href="#windows" id="windows"></a>

* 시스템 설정에서 중앙으로 정의할 수 있으며, 거기에서 설정해야 합니다.
* Windows WSL2 환경 변수는 Linux와 같이 설정해야 합니다. 자세한 내용은 [WSL과 Windows 간 환경 변수 공유하기](https://devblogs.microsoft.com/commandline/share-environment-vars-between-wsl-and-windows/) 참조

## Mac <a href="#mac" id="mac"></a>

환경 변수를 설정할 수 있는 여러 곳이 있습니다.

자세한 내용은 [Mac OS X에서 환경 변수 설정하는 방법 - /etc/\<wbr>launchd.conf - Dowd and Associates](http://www.dowdandassociates.com/blog/content/howto-set-an-environment-variable-in-mac-os-x-slash-etc-slash-launchd-dot-conf/)를 참조하세요.

* `~/.profile`: 터미널에서 실행된 모든 프로그램에서 설정하려는 변수에 사용합니다. Linux와 달리, Terminal.app에서 열리는 모든 쉘이 로그인 쉘입니다.
* `~/.zshrc`: 로그인 쉘이 아닌 쉘에 대해 호출됩니다. 별개의 서브 쉘에서 다시 정의해야 하는 별칭 및 기타 사항에 사용하며, 상속 받는 환경 변수에는 사용하지 마세요.
* `/etc/profile`: `~/.profile`보다 먼저 로드되지만 그 외에는 동일합니다. 시스템의 모든 사용자가 사용하는 터미널 프로그램에 변수를 적용하려면 이를 사용하세요(사용자들이 bash를 사용한다고 가정).
* 사용자의 `launchd` 인스턴스: 사용자가 시작하는 모든 프로그램(GUI 및 CLI)에 적용됩니다. `launchctl`에서 `setenv` 명령을 사용하여 언제든지 변경을 적용할 수 있습니다. 이론적으로 사용자가 로그인할 때 `launchd`가 자동으로 `~/.launchd.conf`에 `setenv` 명령을 읽어들이도록할 수 있으나, 실제로 이 파일을 지원하는 기능이 구현되지 않았습니다. 대신 로그인할 때 스크립트를 실행하여 그 스크립트가 `launchctl`을 호출하여 `launchd` 환경을 설정하도록 하는 다른 메커니즘을 사용할 수 있습니다.
* `/etc/launchd.conf:` 시스템이 시작될 때와 사용자가 로그인할 때 `launchd`가 읽습니다. `launchd`는 루트 프로세스이기 때문에 이러한 변수는 시스템의 모든 프로세스에 영향을 줍니다. 실행 중인 루트 `launchd`에 변경을 적용하려면 명령을 `sudo launchctl`에 파이핑하세요.

이해해야 할 기본적인 사항은 다음과 같습니다:

* 환경 변수는 프로세스가 포크된 시점에 해당 프로세스의 자식에게 상속됩니다.
* 루트 프로세스는 `launchd` 인스턴스이며, 사용자 세션 당 별도의 `launchd` 인스턴스도 있습니다.
* `launchd`는 `launchctl`을 사용하여 현재 환경 변수를 변경할 수 있습니다. 업데이트된 변수는 이후에 포크한 모든 새 프로세스에서 상속받습니다.

IntelliJ가 이 환경 변수를 사용하여 실행되는지 확인할 수 있습니다. `ps eww -o command <PID> | tr ' ' '\\n'` 명령을 사용하세요.

## Linux <a href="#linux" id="linux"></a>

* 자세한 내용은 [환경 변수 - Community Help Wiki](https://help.ubuntu.com/community/EnvironmentVariables#System-wide\_environment\_variables)를 참조하세요.
* UI 애플리케이션을 위해 `.profile / .bashrc`를 설정하는 것은 충분하지 않습니다. 시스템은 터미널에서 실행하지 않기 때문에 관련 변수는 앱을 시작하는 프로세스에서 사용 가능해야 합니다(예: 윈도우 매니저).

## IDEs 및 CLI를 위한 중요한 환경 변수 <a href="#important-environment-variables-for-ides-cli" id="important-environment-variables-for-ides-cli"></a>

### CLI <a href="#cli" id="cli"></a>

* `HTTP_PROXY`
* `HTTPS_PROXY`
* `NO_PROXY`
* `PATH`(Maven 및 Gradle 디렉터리를 포함해야 함)
* `JAVA_HOME`
* `...`

### Java <a href="#java" id="java"></a>

* `http.proxyHost`
* `https.proxyHost`
* `http.proxyPort`
* `https.proxyPort`

### Golang <a href="#golang" id="golang"></a>

* `GOPATH`(go 바이너리 경로)
* `GOROOT`(현재 go 설치)

### Python <a href="#python" id="python"></a>

`PYTHONPATH`

## Proxy <a href="#proxy" id="proxy"></a>

### Java에서 프록시 설정 <a href="#proxy-in-java" id="proxy-in-java"></a>

[Java에서 HTTP/HTTPS 프록시 설정 구성하기](https://memorynotfound.com/configure-http-proxy-settings-java/)를 참조하세요.

### Visual Studio Code에서 프록시 설정 <a href="#set-proxy-in-visual-studio-code" id="set-proxy-in-visual-studio-code"></a>

[JohannesHoppe/settings.json](https://gist.github.com/JohannesHoppe/23105342f6580847578701f0ced9d5b0)를 참조하세요.

### Jetbrains 앱에서 프록시 설정 <a href="#set-proxy-in-jetbrains-apps" id="set-proxy-in-jetbrains-apps"></a>

[HTTP 프록시 | IntelliJ IDEA](https://www.jetbrains.com/help/idea/settings-http-proxy.html)를 참조하세요.

Snyk Jetbrains 플러그인 **는 이 구성에서 프록시 설정을 읽지 않습니다.** **JAVA 프록시 환경 변수뿐만 아니라 CLI 프록시 환경 변수를 설정해야 합니다.**
# 최신 Language Server 사용 강제하기

최신 Snyk Language Server 사용을 강제하려면 다음 단계를 따르세요:

* Snyk 플러그인이나 익스텐션을 최신 버전으로 업데이트합니다. 플러그인의 최신 버전을 사용하면 최신 Language Server로 업데이트할 수 있습니다. 플러그인 업데이트만으로는 Language Server가 업데이트되지 않습니다.
* 플러그인의 Snyk 설정에서 Snyk Language Server 이진 경로(`snyk-ls*`) 또는 Snyk CLI 이진 경로를 찾습니다. 해당 이진 파일의 경로를 알아야 삭제할 수 있습니다. IDE에 따라 다음 지침을 따릅니다:
  * [VS Code](../visual-studio-code-extension/visual-studio-code-extension-configuration.md)
  * [JetBrains](../jetbrains-plugins/configuration-for-the-snyk-jetbrains-plugin-and-ide-proxy.md)
  * [Visual Studio](../visual-studio-extension/visual-studio-extension-configuration.md)
  * [Eclipse](../eclipse-plugin/configuration-of-the-eclipse-plugin.md)
* 해당 경로에서 이진 파일을 삭제합니다.
* Snyk 설정에서 **의존성의 자동 관리**를 활성화하거나 [GitHub](https://github.com/snyk/snyk-ls)에서 최신 Language Server를 다운로드하거나 [getLanguageServer.sh](https://github.com/snyk/snyk-ls/blob/main/getLanguageServer.sh)를 사용합니다. **의존성의 자동 관리**를 활성화하면 최신 Language Server가 자동으로 다운로드됩니다.
* 이전 단계를 완료한 후 IDE를 다시 시작합니다.

플러그인에서 경로를 찾을 수 없는 경우에는 다음과 같이 출력될 수 있습니다:\
`[Info] Snyk Language Server path: /Users/matthew/Library/Application Support/snyk-ls/snyk-ls_darwin_amd64`
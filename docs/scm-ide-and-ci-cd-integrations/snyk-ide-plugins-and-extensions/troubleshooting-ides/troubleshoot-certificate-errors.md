# 인증서 오류 해결

## 문제 <a href="#problem" id="problem"></a>

**로컬 발급자 인증서를 가져올 수 없음** 또는 인증서 경로와 관련된 메시지가 표시됩니다.

## 해결책 <a href="#solution" id="solution"></a>

* 올바른 체인에 모든 인증서를 인증서가 실행 중인 운영 체제, 환경 (JDK) 또는 둘 다에 추가합니다.
* IDE에서 **알 수 없는 인증서 기관 허용**을 설정하거나
* `--insecure` 인수를 사용하여 Snyk CLI 실행 파일에 설정을 사용하여 전달합니다. **이는 다운로드, API 액세스 및 { {Snyk Code}} 통신을 해결하지 않습니다.**

## 사용자 지정 인증서 기관 <a href="#custom-certificate-authorities" id="custom-certificate-authorities"></a>

### Java <a href="#java.1" id="java.1"></a>

CA 및 중간 인증서는 사용된 각 JDK의 인증서 저장소에 추가해야 합니다.

[신뢰할 수 있는 인증서로서 인증서 가져오기 (The Java™ Tutorials > Java SE의 보안 기능 > 코드 서명 및 권한 부여)](https://docs.oracle.com/javase/tutorial/security/toolsign/rstep2.html) 참조.



### Eclipse <a href="#eclipse" id="eclipse"></a>

<figure><img src="../../../.gitbook/assets/image (1) (16).png" alt="알 수 없는 인증서 기관 허용"><figcaption><p>알 수 없는 인증서 기관 허용</p></figcaption></figure>

* Snyk 스캔에 사용된 JDK를 업데이트하여 알 수 없는 인증서를 추가합니다.
* 최신 CLI 및 플러그인 버전으로 업데이트합니다.
* CLI 다운로드는 CLI를 삭제하여 재트리거할 수 있습니다. 경로는 Snyk 기본 설정에 표시됩니다.&#x20;

### IntelliJ <a href="#intellij" id="intellij"></a>

<figure><img src="../../../.gitbook/assets/image (1) (16) (1).png" alt="IntelliJ 설정"><figcaption><p>IntelliJ 설정</p></figcaption></figure>

* [신뢰할 수 있는 루트 인증서 | IntelliJ IDEA](https://www.jetbrains.com/help/idea/ssl-certificates.html) 참조.
* Jetbrains 인증서 처리를 업데이트하는 것만으로는 충분하지 않을 가능성이 높습니다. CLI는 Jetbrains 설정을 사용하지 않고 JAVA\_HOME 및 PATH를 사용하여 JDK를 결정합니다. 이 JDK의 인증서 저장소를 업데이트해야 합니다.&#x20;

### VSCode <a href="#vscode" id="vscode"></a>

확장 프로그램 `win-ca`를 사용하여 Trusted Root 인증서를 확장 프로그램에서 사용할 수 있도록 합니다.

[win-c](https://marketplace.visualstudio.com/items?itemName=ukoloff.win-ca) 참조.

다른 옵션은 환경 변수를 사용하는 것입니다. [노드js에 사용자 지정 인증서 기관 (CA) 추가하는 방법](https://stackoverflow.com/questions/29283040/how-to-add-custom-certificate-authority-ca-to-nodejs)를 참조하세요.

마지막 수단은 인증서 확인을 비활성화하는 것입니다. 추가 인수에 `--insecure`를 추가하여 CLI에서 확인을 비활성화하고 VSCode에서 `https` 호출의 인증서 확인을 비활성화하려면 엄격한 프록시 사용 설정 (`http.proxyStrictSSL`)의 선택을 해제하세요.

<figure><img src="../../../.gitbook/assets/image (2) (17).png" alt="--insecure 인수"><figcaption><p>--insecure 인수</p></figcaption></figure>

&#x20;

<figure><img src="../../../.gitbook/assets/image (3) (9).png" alt="프록시 설정"><figcaption><p>프록시 설정</p></figcaption></figure>
# FIPS-검증 암호화 사용하기

## Snyk에서 FIPS-검증 암호화 사용 가능 여부

FIPS-검증 암호화 지원은 Windows 및 Linux 운영 체제에서만 제한됩니다.

| 운영 체제     | FIPS 지원  |
| ------------ | ---------- |
| Windows      | ✅          |
| Linux        | ✅          |
| Alpine       | ⛔         |
| macOS        | ⛔         |

## Snyk CLI 및 Snyk 언어 서버에서의 FIPS-검증 암호화 지원 및 사용

개발자 경험을 최적화하기 위해 Snyk는 [Snyk 언어 서버](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/snyk-language-server/)와 [Snyk CLI](../getting-started-with-the-snyk-cli.md)를 결합하고 있습니다. 첫 번째 단계로, Snyk는 FIPS 이진 파일을 하나의 애플리케이션 하에 통합하고 있습니다. 나중에는 Snyk 언어 서버에 FIPS가 아닌 CLI 이진 파일도 사용될 것입니다.

Snyk 언어 서버는 이제 CLI 명령어로 실행할 수 있습니다.

```bash
> snyk language-server

# 대신에
> snyk-ls
```

따라서, FIPS-검증 암호화를 사용하는 방법은 CLI 및 언어 서버 모두 동일합니다.

### CLI 및 Snyk 언어 서버에서 FIPS-암호화 사용을 위한 전제 조건

**Linux 운영 체제**

Linux에서 Snyk는 OpenSSL 및 해당 검증된 FIPS 제공 업체를 통해 FIPS-검증 암호화를 지원합니다.

Linux 시스템이 FIPS 인증 요구 사항을 충족하도록 OpenSSL을 설치하고 구성했는지 확인하십시오. 이를 수행하는 방법에 대한 자세한 내용은 [OpenSSL 프로젝트 문서](https://www.openssl.org/docs/fips.html)를 참조하십시오.

**Windows 운영 체제**

Windows에서 Snyk는 Windows CNG API를 통해 FIPS-검증 암호화를 지원합니다.

Windows에서 FIPS를 활성화하려면 [Windows FIPS 정책을 사용](https://docs.microsoft.com/en-us/windows/security/threat-protection/fips-140-validation#step-3-enable-the-fips-security-policy)하십시오.

테스트를 위해, FIPS를 활성화하려면 다음 레지스트리 키 `HKLM\SYSTEM\CurrentControlSet\Control\Lsa\FipsAlgorithmPolicy`에서 `Enabled` 값을 1로 설정하십시오.

#### FIPS 활성화된 이진 파일 다운로드

Snyk 이진 파일은 FIPS 지원 및 미지원 모두 사용할 수 있습니다. 이 모든 파일들은 다음과 같이 Base URL로 구분되어 downloads.snyk.io에 호스팅되어 있습니다.

**FIPS Base URL:** https://downloads.snyk.io/fips/

**정규 Base URL:** https://downloads.snyk.io/

[Snyk 설치 및 사용 방법에 대한 모든 지침](../install-or-update-the-snyk-cli/)은 동일합니다. 유일한 필요한 변경 사항은 적절한 Base URL을 사용하는 것입니다.

### FIPS-활성화 Snyk CLI 다운로드 및 실행 예시

다음 예시는 FIPS-활성화 Snyk CLI를 다운로드하고 실행하는 방법을 설명합니다.

```bash
docker run -it mcr.microsoft.com/cbl-mariner/base/core:2.0 bash
> tdnf install -y ca-certificates
> curl --compressed https://downloads.snyk.io/fips/cli/latest/snyk-linux-arm64 -o snyk
> chmod +x snyk
> ./snyk -d
...
2023-09-13T11:02:49Z main - Features:
2023-09-13T11:02:49Z main -   oauth:               Disabled
2023-09-13T11:02:49Z main -   fips:                Enabled
...
```

### Snyk CLI에서 FIPS-암호화 해결

`not in FIPS mode` 오류는 기본 암호화 라이브러리가 FIPS 모드가 아님을 나타냅니다. 이러한 문제를 해결하려면 [전제 조건](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server)을 충족시키십시오.

## IDE 통합에서의 FIPS-검증 암호화 지원 및 사용

### Visual Studio Code

[FIPS-검증 암호화를 사용](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/)하려면 [Snyk Visual Studio Code 통합](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/)에서 다음을 수행하십시오:

* [전제 조건](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server)을 충족하십시오.
* [적절한 FIPS-활성화된 이진 파일 다운로드](using-fips-validated-cryptography.md#download-fips-enabled-binaries).
* Snyk 설정에서 자동 이진 관리 비활성화.
* 언어 서버 경로 및 CLI 경로를 동일한 바이너리로 설정하여 통합을 구성하십시오.

### Eclipse

[FIPS-검증 암호화를 사용](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/)하려면 [Snyk Eclipse 통합](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/)에서 다음을 수행하십시오:

* [전제 조건](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server)을 충족하십시오.
* [적절한 FIPS-활성화된 이진 파일 다운로드](using-fips-validated-cryptography.md#download-fips-enabled-binaries).
* Snyk 설정에서 자동 이진 관리 비활성화.
* 언어 서버 경로 및 CLI 경로를 동일한 실행 파일로 설정하여 통합을 구성하기.
* Java 런타임 설정을 FIPS-검증된 [JCE(Java Cryptography Extension)](https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search?SearchMode=Basic\&ModuleName=java\&CertificateStatus=Active\&ValidationYear=0.)를 사용하도록 구성하십시오.

### JetBrains

[FIPS-검증 암호화를 사용](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/)하려면 [Snyk JetBrains 통합](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/)에서 다음을 수행하십시오:

* [전제 조건](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server)을 충족하십시오.
* [적절한 FIPS-활성화된 이진 파일 다운로드](using-fips-validated-cryptography.md#download-fips-enabled-binaries).
* Snyk 설정에서 자동 이진 관리 비활성화.
* CLI 경로를 설정하여 통합을 구성하십시오.
* Java 런타임 설정을 FIPS-검증된[JCE(Java Cryptography Extension)](https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search?SearchMode=Basic\&ModuleName=java\&CertificateStatus=Active\&ValidationYear=0.)를 사용하도록 구성하십시오.

### Visual Studio

[FIPS-검증 암호화를 사용](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/)하려면 [Snyk Visual Studio 통합](../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/)에서 다음을 수행하십시오:

* [전제 조건](using-fips-validated-cryptography.md#prerequisites-for-fips-cryptography-in-the-cli-and-snyk-language-server)을 충족하십시오.
* [적절한 FIPS-활성화된 이진 파일 다운로드](using-fips-validated-cryptography.md#download-fips-enabled-binaries).
* 자동 이진 관리 비활성화.
* CLI 경로를 설정하여 통합을 구성하십시오.

## CI/CD 통합에서의 FIPS-검증 암호화 지원과 사용

[CI/CD 통합](../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)에서 FIPS를 사용하려면 FIPS-활성화된 CLI를 직접 사용해야 합니다.

## 패키지 저장소에서의 FIPS-검증 암호화 지원 및 사용

[Snyk Nexus Repository Manager Gatekeeper 플러그인](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/nexus-repository-manager-gatekeeper-plugin.md)과 [Artifactory Gatekeeper 플러그인](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/artifactory-gatekeeper-plugin.md)은 Snyk API를 사용하며 Java VM에서 실행됩니다. FIPS-검증 암호화를 사용하려면 Java 런타임을 구성하여 FIPS-검증[JCE(Java Cryptography Extension)](https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search?SearchMode=Basic\&ModuleName=java\&CertificateStatus=Active\&ValidationYear=0.)를 사용하도록 설정하십시오.
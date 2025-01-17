# Snyk CLI용 릴리스 및 채널

이 페이지에서는 Snyk CLI의 릴리스 및 지원 정책을 설명하며, 각 채널에 가입하는 방법과 각 채널의 목적도 설명합니다.

## 릴리스

v.1.1291.0 이후부터, Snyk CLI는 산업 표준 [Semantic Versioning](https://semver.org/)의 세 부분 표기법을 따릅니다.

버전 번호 MAJOR.MINOR.PATCH가 주어지면 다음과 같이 증가합니다.

1. MAJOR 버전은 \[중요한] 변경 사항이 발생할 때 증가합니다.
2. MINOR 버전은 상호 호환성을 유지하면서 기능을 추가할 때 증가합니다.
3. PATCH 버전은 상호 호환성 버그 수정 시 증가합니다.

필요에 따라 CLI 릴리스에 추가 레이블이 표준에 따라 추가됩니다.

Snyk CLI 문맥에서 Snyk는 중요 변경 사항을 자동화된 워크플로우를 망가뜨리고 기존 작업 환경에서 실패를 야기할 수 있는 변경으로 정의합니다. 중요 변경 사항은 MAJOR를 증가시키고 릴리스 노트에도 언급됩니다.

일부 중요 변경 사항의 예시:

* 출력 필드, 필드 이름 또는 환경 변수의 폐지 또는 변경
* 필수 구성 변경 사항의 도입
* 오류 또는 종료 코드의 변경

## 채널

v.1.1291.0 이후, Snyk는 고객이 필요에 맞춰 채널에 가입할 수 있도록 다른 채널을 제공합니다.

채널을 선택하면 사용할 안정성 수준인 **preview**, **rc**, 또는 **stable**을 선택하는 것입니다.

### 미리보기

{% hint style="info" %}
Snyk는 _진행 중인_ 기능을 테스트하려는 고객을 위해 미리보기 채널을 제공합니다. 그러나 이 채널에는 버그가 포함될 수 있으며 공식적인 지원을 받지 않습니다.
{% endhint %}

미리보기 버전은 프로덕션 환경에서 권장되지 않습니다. 버그가 있을 수 있으며 로컬 환경에서 가장 잘 테스트됩니다. 미리보기 버전을 설치하는 방법은 [채널로부터 독립형 실행 파일 설치](releases-and-channels-for-the-snyk-cli.md#install-standalone-executables-from-a-channel)를 참조하십시오.

* 미리보기: 미리 릴리스가 규칙적으로 배포되며 최신 변경 사항이 포함됨
* 버전 패턴: v{MAJOR}.{MINOR}.{PATCH}-preview
* 주기: 다양함
* 사용 가능:
  * Linux: [https://downloads.snyk.io/cli/preview/snyk-linux](https://downloads.snyk.io/cli/preview/snyk-linux)
  * Windows: [https://downloads.snyk.io/cli/preview/snyk-win.exe](https://downloads.snyk.io/cli/preview/snyk-win.exe)
  * MacOS: [https://downloads.snyk.io/cli/preview/snyk-macos](https://downloads.snyk.io/cli/preview/snyk-macos)
  * Alpine: [https://downloads.snyk.io/cli/preview/snyk-alpine](https://downloads.snyk.io/cli/preview/snyk-alpine)
  * MacOS arm64: [https://downloads.snyk.io/cli/preview/snyk-macos-arm64](https://downloads.snyk.io/cli/preview/snyk-macos-arm64)
  * Linux arm64: [https://downloads.snyk.io/cli/preview/snyk-linux-arm64](https://downloads.snyk.io/cli/preview/snyk-linux-arm64)
  * Alpine arm64: [https://downloads.snyk.io/cli/preview/snyk-alpine-arm64](https://downloads.snyk.io/cli/preview/snyk-alpine-arm64)
  * fips 지원 필요 시, 기본 URL에 `fips`를 추가하십시오. 예시: `https://downloads.snyk.io/fips/cli/preview/snyk-linux`

### rc

* 릴리스 후보판: 미리 릴리스는 특정 시점에 배포되며 추가 테스트를 거친 후 안정 버전으로 승격될 것으로 예상되는 CLI 버전이 포함됨
* 버전 패턴: v{MAJOR}.{MINOR}.{PATCH}-rc
* 주기: 안정 릴리스보다 2주 전인 매 8주마다 (핫픽스 릴리스 가능)
* 사용 가능:
  * Linux: [https://downloads.snyk.io/cli/rc/snyk-linux](https://downloads.snyk.io/cli/rc/snyk-linux)
  * Windows: [https://downloads.snyk.io/cli/rc/snyk-win.exe](https://downloads.snyk.io/cli/rc/snyk-win.exe)
  * MacOS: [https://downloads.snyk.io/cli/rc/snyk-macos](https://downloads.snyk.io/cli/rc/snyk-macos)
  * Alpine: [https://downloads.snyk.io/cli/rc/snyk-alpine](https://downloads.snyk.io/cli/rc/snyk-alpine)
  * MacOS arm64: [https://downloads.snyk.io/cli/rc/snyk-macos-arm64](https://downloads.snyk.io/cli/rc/snyk-macos-arm64)
  * Linux arm64: [https://downloads.snyk.io/cli/rc/snyk-linux-arm64](https://downloads.snyk.io/cli/rc/snyk-linux-arm64)
  * Alpine arm64: [https://downloads.snyk.io/cli/rc/snyk-alpine-arm64](https://downloads.snyk.io/cli/rc/snyk-alpine-arm64)
  * fips 지원 필요 시, 기본 URL에 `fips`를 추가하십시오. 예시: `https://downloads.snyk.io/fips/cli/rc/snyk-linux`

### **stable**

* 안정: 추가 테스트를 거쳐 안정 버전으로 승격된 빌드
* 버전 패턴: v{MAJOR}.{MINOR}.{PATCH}
* 주기: 매 8주마다 (핫픽스 릴리스 가능)
* 사용 가능:
  * [https://github.com/snyk/cli/releases/](https://github.com/snyk/cli/releases/)
  * Linux: [https://downloads.snyk.io/cli/stable/snyk-linux](https://downloads.snyk.io/cli/stable/snyk-linux)
  * Windows: [https://downloads.snyk.io/cli/stable/snyk-win.exe](https://downloads.snyk.io/cli/stable/snyk-win.exe)
  * MacOS: [https://downloads.snyk.io/cli/stable/snyk-macos](https://downloads.snyk.io/cli/stable/snyk-macos)
  * Alpine: [https://downloads.snyk.io/cli/stable/snyk-alpine](https://downloads.snyk.io/cli/stable/snyk-alpine)
  * MacOS arm64: [https://downloads.snyk.io/cli/stable/snyk-macos-arm64](https://downloads.snyk.io/cli/stable/snyk-macos-arm64)
  * Linux arm64: [https://downloads.snyk.io/cli/stable/snyk-linux-arm64](https://downloads.snyk.io/cli/stable/snyk-linux-arm64)
  * Alpine arm64: [https://downloads.snyk.io/cli/stable/snyk-alpine-arm64](https://downloads.snyk.io/cli/stable/snyk-alpine-arm64)
  * fips 지원 필요 시, 기본 URL에 `fips`를 추가하십시오. 예시: `https://downloads.snyk.io/fips/cli/stable/snyk-linux`
* 설치 방법:
  * [npm](install-or-update-the-snyk-cli/#install-the-snyk-cli-with-npm-or-yarn)
  * [Homebrew](install-or-update-the-snyk-cli/#install-with-homebrew-macos-linux)
  * [Scoop](install-or-update-the-snyk-cli/#install-with-scoop-windows)
  * [Snyk-images](install-or-update-the-snyk-cli/#snyk-cli-in-a-docker-image)

Snyk은 다음과 같은 이유로 안정 채널을 선택하는 것을 권장합니다:

* 안정 빌드는 Snyk 개발팀이 SDLC 프로세스에서 CLI를 사용하는 동안 8주 동안 철저히 테스트됨
* 동반되는 릴리스 노트는 어떤 버전이 사용자의 요구에 가장 적합한지 결정하는 데 도움됨

그러나 코드 변경 사항을 합병됨과 동시에 받고 싶은 고객은 미리보기 채널에 가입할 수 있습니다. Snyk은 미리보기 채널에 대해 지원을 제공하지 않으며 알려진 문제가이 채널에 존재할 것으로 예상합니다.

{% hint style="info" %}
이전에 최신 채널에 가입한 기존 Snyk 고객은 자동으로 안정 채널에 가입됩니다. Snyk은 기존 고객에게 방해가 되지 않도록 최신 채널과 안정 채널을 반영하고 있습니다. 그러나 Snyk은 위에 표시된 새로운 채널로 전환할 것을 권장합니다.
{% endhint %}

## 지원 정책

{% hint style="info" %}
본 정책은 2025년 6월 24일부터 시행됩니다.
{% endhint %}

Snyk은 안정 채널에서 가장 최근 12개월간의 CLI 버전을 지원하여 기능과 성능을 보장합니다. 오래된 버전은 지원 중단 (EOS)으로 간주되며 버그 수정이나 문제 해결 지원을 받지 못합니다.

Snyk은 새 버전에서만 수정을 제공하며 오래된 버전은 수정할 수 없습니다. 고객은 개선 사항을 받으려면 업그레이드해야 합니다.

본 정책은 혁신을 촉진하며 자원을 최적화합니다.

도움이 필요하면 [요청](https://support.snyk.io)을 Snyk 지원에 제출하십시오.

## 채널로부터 독립 실행 파일 설치

각 채널의 `release.json`을 사용하십시오. 다음은 여기에서 제공된 다운로드 링크이며, MacOS 플랫폼의 미리보기 버전 예시를 설명하겠습니다:

* [https://downloads.snyk.io/cli/preview/release.json](https://downloads.snyk.io/cli/preview/release.json)
* [https://downloads.snyk.io/cli/preview/snyk-macos](https://downloads.snyk.io/cli/preview/snyk-macos)

MacOS에 대해, `snyk-preview`라는 임시 폴더에 CLI 미리보기 버전을 다운로드하고 실행해보세요. 다음 명령어를 실행하여 진행할 수 있습니다.

```sh
mkdir snyk-preview
cd snyk-preview
curl --compressed https://downloads.snyk.io/cli/preview/snyk-macos -o snyk
chmod +x ./snyk
./snyk -version
```

## IDE에서 채널 선택

{% hint style="info" %}
이 기능은 IntelliJ IDE에서 사용 가능합니다. Snyk는 이 기능을 다른 지원되는 IDE로 확장하고 있습니다.
{% endhint %}

모든 IDE의 기본 채널은 안정 채널입니다.

IDE에서 채널을 선택하려면 스크린샷에서 보이는 대로 드롭다운을 사용하여 CLI 릴리스 채널을 선택하십시오. 사용자는 채널을 변경할 수 있으며, 예를 들어 핫픽스를 받기 위해 릴리스-후보 (**rc**)로 전환할 수 있습니다.

그러나 Snyk은또한 IDE 사용자에게 **stable** 채널을 기본으로 추천합니다.

<figure><img src="../.gitbook/assets/Screenshot 2024-09-02 at 10.32.41.png" alt="Choose a CLI release channel"><figcaption><p>CLI 릴리스 채널 선택</p></figcaption></figure>

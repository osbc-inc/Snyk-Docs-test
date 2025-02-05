# Swift 및 Objective-C

## 적용 가능성

Snyk는 [코드 분석을 위한 Swift](swift-for-code-analysis.md) 및 오픈 소스용 [Swift 및 Objective-C를 지원](swift-and-objective-c-for-open-source.md)합니다. 이에는 CocoaPods 지원도 포함됩니다.

Snyk 제품을 사용하여 애플리케이션으로 가져오거나 테스트하거나 모니터링할 수 있는 언어 가용성을 확인하세요.

{% hint style="info" %}
**지원되는 Swift 버전**

5.7.x까지의 버전을 사용할 수 있습니다.
{% endhint %}

이용 가능한 기능:

* SCM 가져오기, Snyk 오픈 소스 및 Snyk 코드에서 사용 가능합니다. Snyk 오픈 소스와 함께 사용하는 경우 SCM 가져오기는 CocoaPods에 대해 사용할 수 있습니다.
* CLI 및 IDE를 통해 앱을 테스트하거나 모니터링할 수 있습니다. Snyk 오픈 소스 및 Snyk 코드에서 사용 가능합니다.
* `pkg:swift`, `pkg:cocoapods`를 사용하여 앱의 SBOM을 테스트할 수 있습니다.
* `pkg:swift`, `pkg:cocoapods`를 사용하여 앱의 패키지를 테스트할 수 있습니다.

## 패키지 관리자 및 지원되는 파일

Swift 및 Objective-C용 Snyk은 CocoaPods, Swift 패키지 관리자 v3.0 이상을 지원하며 [cocoapods.org](https://cocoapods.org/), [maven.org](https://maven.org/)과 같은 패키지 레지스트리를 여러 소스로 사용합니다.

Swift 및 Objective-C용 Snyk은 다음 파일 형식을 지원합니다:

* Snyk 오픈 소스:
  * CocoaPods용: `podfile`, `podfile.lock`
  * Swift용: `package.swift`
* Snyk 코드: `.swift`

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리가 Swift 및 Objective-C용 Snyk에서 지원됩니다:

* Swift 표준 라이브러리 - 포괄적
* Foundation - 포괄적
* Appkit - 포괄적
* Swift UI - 포괄적
* UI Kit - 포괄적
* Asynchttpclient - 포괄적
* Commoncrypt - 포괄적
* Commoncrypto - 포괄적
* Cryptokit - 포괄적
* Cryptoswift - 포괄적
* Cryptor - 포괄적
* AlamoFire - 포괄적
* Filekit - 포괄적
* google-gemini/generative-ai-swift - 포괄적
* MacPaw/OpenAI - 포괄적
* dylanshine/openai-kit - 포괄적
* Pathos - 포괄적
* SQLite3 - 포괄적
* Webkit - 포괄적
* SwiftCLI - 포괄적
* ShellOut - 포괄적
* SwiftShell - 포괄적
* Subprocess - 포괄적
* Shout - 포괄적
* Swiftline - 포괄적
* RNCryptor - 포괄적

## 기능

Swift 및 Objective-C용 Snyk에서 지원하는 다음 기능:

| Snyk 오픈 소스                                        | Snyk 코드                               |
| ------------------------------------------------- | ------------------------------------- |
| <ul><li>라이선스 스캔 (CocoaPods)</li><li>리포트</li></ul> | <ul><li>리포트</li><li>인터파일 분석</li></ul> |

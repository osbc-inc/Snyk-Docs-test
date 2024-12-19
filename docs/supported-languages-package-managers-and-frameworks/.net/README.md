# .NET

## 적용성

Snyk는 [.NET 코드 분석](.net-for-code-analysis.md) 및 [.NET 오픈 소스](.net-for-open-source.md)을 지원합니다. [.NET을 위한 Snyk 가이드](guidance-for-snyk-for-.net.md)가 있습니다. 도움이 필요한 경우 [Snyk를 위한 문제 해결](troubleshooting-snyk-for-.net.md)을 참조하십시오.

Snyk는 NuGet 애플리케이션의 스캔 기능을 크게 향상시킨 새로운 초기 액세스 기능을 소개했습니다. 자세한 정보 및 이러한 기능에 대한 액세스는 [Improved .NET scanning](improved-.net-scanning.md) 페이지를 참조하십시오.&#x20;

Snyk 제품을 사용하여 가져올 수 있는 언어 사용 가능성을 확인하십시오.&#x20;

사용 가능한 기능:

- {{Snyk Open Source}} 및 {{Snyk Code}}의 SCM 가져오기 가능. .NET을 위해 {{Snyk Open Source}}를 사용하는 경우, SCM 가져오기는 NuGet을 사용할 때에만 가능합니다.
- CLI 및 IDE를 통해 앱을 테스트하거나 모니터링하는 기능은 {{Snyk Open Source}} 및 {{Snyk Code}} 모두에서 제공됩니다.
- `pkg:nuget`을 사용하여 앱의 SBOM을 테스트합니다.
- `pkg:nuget`을 사용하여 앱의 패키지를 테스트합니다.

## 패키지 관리자 및 지원되는 파일 확장자

Snyk for .NET은 NuGet 및 Paket을 패키지 관리자로 지원하며 [nuget.org](https://www.nuget.org/)를 패키지 레지스트리로 지원하며 다음 파일 형식을 지원합니다:

- {{Snyk Open Source}}:
  - NuGet의 경우: `project.assets.json`, `*.sln`, `packages.config,` `project.json`
  - Paket의 경우: `paket.dependencies`, `paket.lock`

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리가 Snyk for .NET에서 지원됩니다: (코드만)

- .NET 6 - 포괄적
- .NET Core - 포괄적
- .NET Framework 4.6-4.8.x - 포괄적
- Anthropic.SDK - 포괄적
- ASP.NET 6.x (C# 전용) - 포괄적
- Azure.AI.OpenAI - 포괄적
- Dapper - 포괄적
- fastJSON - 포괄적
- Google\_GenerativeAI - 포괄적
- Microsoft.CodeAnalysis.VisualBasic - 포괄적
- Mistral.SDK - 포괄적
- System.CodeDom.Compiler - 포괄적
- Windows Forms - 일부

## 기능

다음 기능이 Snyk for .NET에서 지원됩니다:

| {{Snyk Open Source}}                                              | {{Snyk Code}}                                                  |
| --------------------------------------------------------------------- | -------------------------------------------------------------------- |
| <ul><li>수정 PRs (NuGet) </li><li>라이선스 스캔 </li><li>보고서</li></ul> | <ul><li>보고서</li><li>사용자 정의 규칙 </li><li>간파 분석</li></ul> |
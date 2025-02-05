# .NET용 Snyk에 대한 지침

## 의존성 분석

.NET 생태계에서는 명시적인 의존성과 개발자에게 완전히 숨겨진 일부 의존성이 포함되어 있습니다. 주어진 .NET 응용 프로그램의 취약점을 정확하게 식별하려면 이러한 의존성을 정확하게 해결해야 합니다.

Snyk은 Snyk CLI 및 GitHub와 같은 소스 코드 관리 (SCM) 시스템에서 의존성을 다르게 해결합니다.

**CLI에서**, 프로젝트 의존성을 `PackageReference`를 사용하여 관리하는 경우 Snyk은 `obj/project.assets.json`을 스캔합니다. 의존성을 `packages.config`를 사용하여 관리하는 경우 Snyk은 `packages` 디렉토리를 스캔합니다. 이 접근 방식은 스캔 결과의 정확성에 기여합니다.

{% hint style="info" %}
런타임 의존성(런타임 환경에서 제공되는 메타 패키지)은 CLI에서 더 정확하게 해결됩니다. 호스트 머신이 앱을 실행하는 서버와 유사한 런타임 SDK를 사용하는 경우
{% endhint %}

**SCM 통합**에서는 문제가 생성된 파일이 없기 때문에 스캔이 다른 프로세스로 이루어집니다. 이를 해결하기 위해 Snyk는 NuGet 의존성 [해결 알고리즘](https://docs.microsoft.com/en-us/nuget/concepts/dependency-resolution)을 따라 의존성 트리를 구성합니다.

.NET 자동화된 수정에 대한 자세한 정보는 [Snyk 블로그](https://snyk.io/blog/automated-vulnerability-fixes-dot-net-dependencies)를 참조하십시오.

## 빌드 시간 대 런타임 의존성

* **빌드 시간 의존성**: Snyk은 빌드 시간 의존성을 빌드 시간에 해결되는 것으로 이해하며 런타임에서 변경되지 않습니다.
* **런타임 의존성**: Snyk은런타임 의존성을 설치된 런타임에 맞춰 해결되는 것으로 이해하며 예를 들어 .NET 프레임워크 (<=4) / .NET [런타임](https://docs.microsoft.com/en-us/dotnet/core/versions/selection?WT.mc_id=DOP-MVP-5001511&) (Core 및 .NET 5+)에서 `System.Net.Http`와 같은 패키지를 포함합니다. Snyk은 때때로 런타임 의존성을 메타 패키지로 지칭합니다.

**런타임 의존성에서의 취약성**에 대처하기 위해 다음 중 하나를 선택할 수 있습니다. 이는 SCM 및 CLI 사이에서 다르게 적용됩니다.

### **SCM에서의 런타임 의존성에서의 취약성**

Microsoft에서 항상 최신 패치를 설치하는 시스템에서 응용 프로그램을 실행하므로 오류가 발견된 것이 잘못된 것이라고 여길 수 있다면, 해당 취약성을 [무시하기를](../../manage-risk/prioritize-issues-for-fixing/ignore-issues/) 선택할 수 있습니다.

### **CLI에서의 런타임 의존성에서의 취약성**

응용 프로그램이 프로덕션 환경에서 항상 최신/명시적 패치를 사용하면 취약성이 더 이상 프로젝트에 영향을 주지 않을 수 있으므로 해당 취약성을 [무시](../../manage-risk/prioritize-issues-for-fixing/ignore-issues/)하고 다음을 수행할 수 있습니다:

* 프로덕션에서 응용 프로그램이 항상 최신 SDK 패치 버전에서 실행되는 경우, 프로젝트 파일에서 `TargetLatestRuntimePatch`를 `true`로 설정할 수 있습니다. 환경을 최신 런타임 버전으로 업그레이드하는 것을 잊지 말아야 합니다(예: 개발, 프로드 환경).

```
<TargetLatestRuntimePatch>true</TargetLatestRuntimePatch>
```

* 런타임을 포함하는 [독립형](https://docs.microsoft.com/en-us/dotnet/core/deploying/#publish-self-contained) 응용 프로그램을 게시할 수 있습니다. 그런 다음 프로젝트 파일에서 지정한 패치 버전에 `RuntimeFrameworkVersion`을 설정할 수 있습니다. 더 이상 관련이 없다고 여기는 취약성을 [무시](../../snyk-cli/scan-and-maintain-projects-using-the-cli/ignore-vulnerabilities-using-the-snyk-cli.md)할 수 있습니다.

```
<PropertyGroup>
  <RuntimeFrameworkVersion>5.0.7</RuntimeFrameworkVersion>
</PropertyGroup>
```

이 안내서를 사용하여 여러분의 기술 스택에서 Snyk을을 효과적으로 적용하세요.

## C\#

Snyk Code는 IDE, CLI 및 Git 통합을 사용하여 C# 코드를 분석할 수 있습니다.

프레임워크 지원은 [Snyk Code - 지원하는 언어 및 프레임워크](../)를 참조하세요.

## Nuget

**대상 프레임워크**: Snyk는 대상 프레임워크를 식별하고 git 통합을 사용하여 각 식별된 버전에 대한 결과를 제공합니다.

**개발 의존성**: Snyk는 일반적으로 프로덕션으로 푸시되지 않는 개발자 의존성을 스캔하지 않습니다. 이러한 의존성을 스캔하려면 **설정 > 언어 > .NET** 설정을 통해 Nuget git 가져오기에서 가시성을 활성화하세요 ([.NET을 위한 Git 설정](./#git-settings-for-.net)).\
Snyk는 [`*.proj`](guidance-for-snyk-for-.net.md#user-content-fn-1) 파일의 빌드 및 개발 의존성 섹션, `packages.config` 및 `project.json` 파일을 스캔하고 수정합니다.

**잠금 파일**: 현재 **packages-lock.json**은 지원되지 않습니다. Snyk는 설치된 의존성을 결정하기 위해 빌드 시스템과 상호 작용합니다.

**PackageReference:** Snyk는 현재 버전 속성이 필요합니다. 프로젝트에 이 속성이 없는 경우 Snyk는 프로젝트에 대한 PR을 열 수 없을 수 있습니다.

**Git 분석**

의존성 트리가 어떻게 생성되는지:

* .NET Core의 경우 \*.proj 파일 사용
* .NET Framework의 경우 \*.proj 파일 및 packages.config 사용

SCM 통합은 다음을 지원합니다:

* \*.csproj
* \*.fsproj
* \*.vbproj
* packages.config

수정 Pull Requests

* 현재 NuGet을 사용하여 프로젝트 의존성을 관리하고 [`PackageReference`](https://docs.microsoft.com/en-us/nuget/consume-packages/package-references-in-project-files) 또는 [`packages.config`](https://docs.microsoft.com/en-us/nuget/reference/packages-config)를 활용하는 경우 Snyk는 실제 수정이 있는 경우 자동으로 manifest 파일에서 의존성 버전을 업데이트할 수 있습니다. 그러면 수정 사항을 검토하고 병합할 수 있습니다.

**CLI 분석**

CLI는 다음 구성 파일을 지원합니다:

**project.assets.json**

Snyk는 프로젝트.assets.json을 스캔하여 의존성을 결정할 수 있지만 파일을 생성해야 합니다. 마찬가지로 솔루션 파일 (.sln)을 가리킬 때는 먼저 파일을 생성해야 합니다.

"\*\*dotnet restore"\*\*를 실행하여 "**snyk test**" 명령을 실행하기 전에 필요한 `project.assets.json`을 생성하세요.

솔루션 파일은 분석을 수행하는 데 필요한 파일을 가리킵니다. 프로젝트 자체가 스캔되려면 프로젝트에 `project.assets.json` 파일이 있어야 합니다. Snyk가 솔루션 파일을 스캔의 진입점으로 사용하도록 원한다면 Snyk CLI를 솔루션 파일을 가리키도록 `--file=<filename>.sln`을 사용합니다.

동일한 프로젝트에서 여러 대상 프레임워크를 사용하는 경우 CLI 스캔은 프로젝트에서 선언된 첫 번째 대상 프레임워크에 대한 결과를 반환합니다.

**packages.config**

**nuget install -OutputDirectory packages**를 실행한 후 **snyk test** 명령을 실행하세요.

{% hint style="info" %}
런타임 의존성(런타임 환경에서 제공되는 메타 패키지)은 CLI에서 더 정확하게 해결됩니다. 호스트 머신이 앱을 실행하는 서버와 유사한 런타임 SDK를 사용하는 경우
{% endhint %}

## Paket

Snyk은 CLI를 통해 관리되는 Paket으로 관리되는 의존성을 분석할 수 있습니다. 알고 있어야 할 점은 paket.dependencies와 paket.lock 파일이 있어야 합니다.

Paket 지원에 대한 자세한 내용은 [Snyk for .NET](./)를 참조하세요.

## 기타

Snyk는 고유한 의존성 관리 전략에 대한 사용자 정의 테스트 API를 제공합니다. 자세한 정보는 [패키지에 대한 이슈 목록](https://apidocs.snyk.io/?version=2022-11-14#get-/orgs/-org_id-/packages/-purl-/issues) 페이지를 확인하십시오.

## 닷넷 개발자를 위한 추가 보안 주제

다음은 이 생태계와 관련된 Snyk 보안 팀 및 개발자 관련 기사 모음입니다.

* [Snyk 블로그](https://snyk.io/blog/)
* [닷넷을 위한 Snyk](./)
* [.NET 애플리케이션을 컨테이너화하는 최상의 방법](https://snyk.io/blog/best-practices-for-containerizing-net-applications/)​

# .NET용 Snyk 문제 해결

## .NET에 대한 Snyk 오픈 소스 제한 사항

* [`Directory.Build.props`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets) 및 [`Directory.Build.targets`](https://docs.microsoft.com/en-us/visualstudio/msbuild/customize-your-build?view=vs-2022#directorybuildprops-and-directorybuildtargets)은 SCM 통합에 지원되지 않습니다. Snyk CLI를 사용하여 중앙 패키지 관리로 비공개 종속성을 스캔할 수 있습니다. `dotnet restore`를 실행한 다음 `--all-projects`를 사용하여 `snyk test`를 실행해야 합니다. 각 하위 폴더에는 고유한 `project.assets.json` 파일이 있습니다.
* `<ProjectReference>` 요소는 지원되지 않습니다.
* 조직 또는 그룹에 대한 개선된 .NET 스캔을 활성화하여 SCM 통합에 대한 비공개 종속성 스캔이 지원됩니다. 더 많은 세부 정보를 보려면 [Snyk Preview](https://docs.snyk.io/snyk-admin/snyk-preview) 메뉴로 이동하십시오. Snyk CLI를 사용하여 비공개 종속성을 스캔할 수도 있습니다.

도움이 필요하면 [Snyk 지원팀에 문의](https://support.snyk.io)하십시오.

## .NET에 대한 Snyk 기술 지원

다음 지원 문서를 사용할 수 있습니다:

* [프로젝트 가져오기 오류](https://support.snyk.io/s/article/Project-import-errors)
* [.NET 대상 프레임워크 변경이 Snyk 포털에 반영되지 않음](https://support.snyk.io/s/article/Changing-NET-Target-Framework-not-reflected-in-Snyk-Portal)
* [Snyk이 .NET 프로젝트를 집계하는 방법](https://support.snyk.io/s/article/How-does-Snyk-aggregate-NET-Projects)

더 많은 도움이 필요하면 [Snyk 지원팀에 문의](https://support.snyk.io)하십시오.

# Snyk IDE 플러그인 및 확장 프로그램

다음은 사용할 수 있는 Snyk 플러그인 및 확장 프로그램입니다.

- [Eclipse 플러그인](eclipse-plugin/)
- [JetBrains 플러그인](jetbrains-plugins/)
- [Visual Studio 확장 프로그램](visual-studio-extension/)
- [Visual Studio Code 확장 프로그램](visual-studio-code-extension/)
- [Language Server](snyk-language-server/)

Snyk 보안 플러그인 및 확장 프로그램은 Snyk 프로젝트의 보안 취약점 및 문제를 찾아 수정합니다. 개별 개발자, 오픈 소스 기여자 및 코드 유지 관리자가 보안 검토를 통과하고 개발 중 나중에 비용이 드는 수정 사항을 피하며, 안전한 코드를 개발하고 전달하는 데 걸리는 시간을 줄일 수 있습니다.

취약점 스캔 결과는 IDE에서 컨텍스트, 영향 및 수정 지침과 함께 표시되며, 취약점의 수정은 IDE 자체에서 직접 수행할 수 있습니다.

Snyk IDE 플러그인 및 확장 프로그램은 많은 기능을 수행하기 위해 Snyk CLI에 의존합니다. 자세한 내용은 각 IDE의 설명서를 참조하십시오. 문제 해결할 때는 항상 동일한 작업에 대해 디버그 옵션인 `-d`를 사용하여 CLI를 실행하면 도움이 됩니다.

Snyk IDE 플러그인 및 확장 프로그램은 또한 [Snyk 취약점 데이터베이스](https://security.snyk.io/)에 의존합니다. 자세한 내용은 [Snyk 취약점 데이터베이스](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/snyk-vulnerability-database.md)의 설명서를 참조하십시오.

Snyk가 기본값이 아닌 다른 지역에 데이터를 호스팅하고 있는 경우 IDE에서 `Custom endpoint`를 설정해야 합니다. 자세한 내용은 [IDEs URLS](../../working-with-snyk/regional-hosting-and-data-residency.md#ides-urls)를 확인하십시오.

교육이 가능합니다: [IDE에서 Snyk 사용 소개](https://learn.snyk.io/lesson/snyk-in-an-ide/)
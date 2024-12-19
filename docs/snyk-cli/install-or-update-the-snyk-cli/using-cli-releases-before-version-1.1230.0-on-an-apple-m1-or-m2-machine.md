# Apple M1 또는 M2 기기에서 버전 1.1230.0 이전의 CLI 릴리스 사용

**버전 1.1230.0**부터 Snyk CLI는 Apple 실리콘(M1 및 M2 기기 포함)을 지원하므로 **더 이상 이 페이지에서 제안된 Rosetta 2 소프트웨어가 필요하지 않습니다**. Snyk은 항상 CLI 설치를 최신 버전으로 유지하는 것을 권장합니다. Snyk CLI의 설치된 버전을 확인하려면 `snyk --version`을 실행하십시오.

Apple M1 기기에서 [버전 1.1230.0](https://github.com/snyk/cli/releases/tag/v1.1230.0) 이전의 Snyk CLI 릴리스를 실행하면 다음 오류가 표시될 수 있습니다:

```
$ snyk --version
bad CPU type in executable
```

M2 기기에서는 다음 오류가 표시될 수 있습니다:

`Unknown system error -86`

[버전 1.1230.0](https://github.com/snyk/cli/releases/tag/v1.1230.0) 이전의 Snyk CLI 릴리스로 이러한 오류를 피하려면 Apple [Rosetta 2](https://support.apple.com/en-gb/HT211861) 소프트웨어를 설치하십시오:

`softwareupdate --install-rosetta`
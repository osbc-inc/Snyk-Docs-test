# 프로젝트를 정기적으로 모니터링하세요

`snyk protect` 명령어와 `@snyk/protect` [패키지](https://github.com/snyk/snyk/tree/master/packages/snyk-protect)를 사용하면 알려진 취약점에 대응할 수 있습니다. 지속적으로 발표되는 새로운 취약점에 대응하기 위해 `snyk monitor`를 사용하세요.

배포 직전에 프로젝트 디렉토리에서 `snyk monitor`를 실행하세요:

`cd ~/projects/myproject/`\
`snyk monitor`

현재 종속성에 대한 스냅샷을 Snyk 대시보드로 보내어 Snyk이 해당 종속성의 새로 발표된 취약점이나 이전에 사용 불가능했던 패치 또는 업그레이드 경로가 만들어졌을 때 알릴 수 있습니다. 동일 프로젝트의 여러 스냅샷을 취할 경우, Snyk은 최신 스냅샷에 대한 새로운 정보에 대해서만 알립니다.

Snyk 웹 UI에 로그인하고 조직의 프로젝트 페이지로 이동하세요 ([https://app.snyk.io/monitor/](https://app.snyk.io/monitor/)). 여러분의 프로젝트의 최신 스냅샷과 히스토리를 확인할 수 있습니다.

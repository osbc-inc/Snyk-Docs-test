# Snyk CLI를 사용하여 Snyk 앱 만들기

{% hint style="warning" %}
릴리스 상태\
모든 `snyk apps` 하위 명령은 `--experimental` 플래그 뒤에만 사용할 수 있으며 행동은 언제든지 알림 없이 변경될 수 있습니다. 모든 명령을 신중히 사용하십시오.
{% endhint %}

`Snyk` CLI를 사용하여 `snyk apps create`를 실행하여 `Snyk` 앱을 만들 수 있습니다. 이 명령을 사용하는 두 가지 방법이 있습니다.

첫 번째는 일반 모드입니다. 예를 들면:

`snyk apps create --experimental --org=48ebb069-472f-40f4-b5bf-d2d103bc02d4 --name='내 멋진 앱' --redirect-uris=https://example1.com,https://example2.com --scopes=apps:beta`

두 번째는 대화형 모드이며, 이 모드에서는 일반 모드와 유사한 방식으로 모든 값을 입력하도록 안내합니다. 다음은 대화형 모드의 예시입니다:

`snyk apps create --experimental --interactive`

```
snyk apps create --experimental --interactive

? 사용자가 Snyk 앱을 설치할 때 사용자에게 표시되는 Snyk 앱의 이름은 무엇입니까? 내 멋진 Snyk 앱
? 앱의 리디렉션 URI(쉼표로 구분된 목록. 예: https://example1.com,https://example2.com)은 무엇입니까?  https://example1.com
? 앱의 권한 스코프(쉼표로 구분된 목록. 예: org.read)는 무엇입니까?  apps:beta
? Snyk 앱을 생성할 org ID를 입력하세요:  48ebb069-472f-40f4-b5bf-d2d103bc02d4
```

## `snyk apps create`를 위한 옵션

`--interactive`

대화형 모드에서 `snyk apps create` 명령을 사용합니다.

`--org=<ORG_ID>`

Snyk 앱을 생성할 `<ORG_ID>`를 지정합니다. `create` 명령에 필요합니다.

`--name=<SNYK_APP_NAME>`

사용자가 앱을 승인할 때 사용자에게 표시할 이름입니다. `create` 명령에 필요합니다.

`--redirect-uris=<REDIRECT_URIS>`

리디렉션 URI 목록인 쉼표로 구분된 목록입니다. 인증 후 호출하기 위한 허용되는 리디렉션 URI 목록을 형성합니다. `create` 명령에 필요합니다.

`--scopes=<SCOPES>`

Snyk 앱에서 요구하는 스코프 목록인 쉼표로 구분된 목록입니다. 승인 중에 앱이 요청할 수 있는 스코프 목록을 형성합니다. `create` 명령에 필요합니다.

`--context=<CONTEXT>`

설치된 경우 Snyk 앱이 사용할 `context`입니다.

`tenant` 또는 `user`가 될 수 있습니다. `context`가 지정되지 않은 경우 기본값은 `tenant`입니다.

`tenant` context가 있는 Snyk 앱은 봇 사용자로 작동하므로 개별 사용자와 관련이 없으며, 따라서 설치한 사용자가 Snyk 조직을 떠나더라도 지속됩니다. 반면, `user` context가 있는 Snyk 앱은 설치된 사용자처럼 작업합니다. Snyk 앱이 개별 사용자에게 특정 작업을 수행하는 경우에만 `user` context를 지정합니다. 의심이 들 경우 `tenant`를 사용하십시오.

## `snyk apps`의 서브 명령

모든 `snyk apps` 하위 명령은 `snyk apps create`와 같은 방식으로 사용됩니다.

`snyk apps` 명령 아래의 모든 사용 가능한 서브 명령에 대해 알아보려면 `--help` 옵션을 사용하십시오. `snyk apps --help`입니다.
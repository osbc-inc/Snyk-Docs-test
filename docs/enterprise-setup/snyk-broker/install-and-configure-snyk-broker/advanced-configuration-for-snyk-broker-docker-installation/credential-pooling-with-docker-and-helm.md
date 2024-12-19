# Docker와 Helm을 사용한 자격 증명 풀링

어떤 상황에서는 "풀(pool)"을 만드는 것이 바람직할 수 있습니다. 예를 들어, 속도 제한 문제를 해결하기 위해 자격 증명 풀을 만들 수 있습니다. 이를 위해 쉼표로 각 자격 증명을 구분하여 끝이 `_POOL`로 끝나는 환경 변수를 만듭니다. Broker Client는 변수 대체를 수행할 때 사용 중인 변수에 `_POOL` 접미사가 있는지 확인하고 있다면 해당 풀에서 다음 아이템을 사용합니다. 예를 들어, `GITHUB_TOKEN` 환경 변수를 설정했지만 여러 토큰을 제공하려는 경우 다음과 같이 수행할 수 있습니다:

``` 
GITHUB_TOKEN_POOL=token1, token2, token3
```

이 환경 변수와 값을 Helm Chart에 추가할 수 있습니다.

그러면 Broker Client는 `GITHUB_TOKEN`이 필요할 때마다 `GITHUB_TOKEN_POOL`에서 항목을 가져옵니다.

자격 증명은 순환 방식으로 취득되며, 처음, 두 번째, 세 번째와 같이 계속해서 Broker Client가 끝에 도달하고 다시 처음 자격 증명을 취득합니다.

`/systemcheck` 엔드포인트를 호출하면 모든 자격 증명을 확인하여 순서대로 배열을 반환하며, 첫 번째 항목부터 차례대로 반환합니다. 예를 들어, GitHub Client를 실행하고 다음과 같은 설정이 있는 경우:

```
GITHUB_TOKEN_POOL=good_token, bad_token
```

`/systemcheck` 엔드포인트는 다음과 같이 반환하며 첫 번째 객체는 `good_token`에 대한 것이고 두 번째는 `bad_token`에 대한 것입니다:

```
[
  {
    "brokerClientValidationUrl": "https://api.github.com/user",
    "brokerClientValidationMethod": "GET",
    "brokerClientValidationTimeoutMs": 5000,
    "brokerClientValidationUrlStatusCode": 200,
    "ok": true,
    "maskedCredentials": "goo***ken"
  },
  {
    "brokerClientValidationUrl": "https://api.github.com/user",
    "brokerClientValidationMethod": "GET",
    "brokerClientValidationTimeoutMs": 5000,
    "ok": false,
    "error": "401 - {\"message\":\"Bad credentials\",\"documentation_url\":\"https://docs.github.com/rest\"}",
    "maskedCredentials": "bad***ken"
  }
]
```

자격 증명은 마스킹되어 표시됩니다. 그러나 사용자 자격 증명에 문자가 6개 이하 포함된 경우 마스크로 완전히 대체됩니다.

## **자격 증명 풀링의 한계**

자격 증명의 유효성이 사용하기 전에 확인되지 않으며, 유효하지 않은 자격 증명이 풀에서 제거되지 않으므로 다른 시간에 자격 증명이 속도 제한에 도달하는 것을 피하기 위해 자격 증명이 Broker Client에 의해 독점적으로 사용되는 것이 _강력히_ 권장됩니다. 또한 사용하기 전에 `/systemcheck` 엔드포인트를 호출하는 것이 좋습니다.

GitHub와 같은 일부 제공업체는 토큰 또는 자격 증명 당이 아니라 사용자 당 속도 제한을 적용하며, 이러한 경우에는 하나의 자격 증명 당 하나의 계정을 만들어야 합니다.

## **자격 증명 매트릭스**

자격 증명 매트릭스 생성은 지원되지 않습니다.

이 경우 "매트릭스"는 길이가 'x' 및 'y' 인 두 개 이상의 `_POOL`을 가져와 최종 길이가 `x * y`인 하나의 최종 풀을 생성하는 것을 정의합니다. 예를 들어, 다음과 같은 입력이 주어진 경우:

```
USERNAME_POOL=u1, u2, u3
PASSWORD_POOL=p1, p2, p3
CREDENTIALS_POOL=$USERNAME:$PASSWORD
```

매트릭스 지원은 내부적으로 다음을 생성합니다:

```
CREDENTIALS_POOL=u1:p1,u1:p2,u1:p3,u2:p1,u2:p2,u2:p3,u3:p1,u3:p2,u3:p3
```

대신, Broker Client는 내부적으로 첫 번째로 발견된 풀만 사용하여 다음을 생성합니다:

```
CREDENTIALS_POOL=u1:$PASSWORD,u2:$PASSWORD,u3:$PASSWORD
```  
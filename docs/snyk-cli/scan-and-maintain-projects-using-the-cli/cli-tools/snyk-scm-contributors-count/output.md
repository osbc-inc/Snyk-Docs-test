---
description: SCM-Contributors-Count 도구 실행 결과 출력
---

# 출력

## 요약

요약 섹션은 출력의 시작과 끝에 모두 나타납니다. 예를 들어:

```
#### 요약
비공개 레포지토리 기여자 카운트: 1
공개 레포지토리 기여자 카운트: 1
비공개와 공개 레포지토리의 총 고유 기여자 수: 1
비공개 레포지토리 수: 1
공개 레포지토리 수: 1
총 레포지토리 수: 2
제외된 수: 1
```

* `비공개 레포지토리 기여자 카운트` - 찾거나 제공된 비공개 레포지토리의 고유 기여자 수.
* `공개 레포지토리 기여자 카운트` - 찾거나 제공된 공개 레포지토리의 고유 기여자 수.
* `비공개와 공개 레포지토리의 총 고유 기여자 수` - 조사된 모든 레포지토리 전체의 고유 기여자 수.
* `비공개 레포지토리 수` - 검색된 비공개 레포지토리 수.
* `공개 레포지토리 수` - 검색된 공개 레포지토리 수.
* `총 레포지토리 수` - 검색된 총 레포지토리 수(공개 및 비공개).
* `제외된 수` - 제공된 제외 파일에 따라 계산되지 않은 기여자 수.

## 세부 정보

```
### 세부 정보:
## 레포지토리 목록

# 비공개 레포지토리:
someOrganization/someRepository(비공개)

# 공개 레포지토리:
anotherOrganization/anotherRepository(공개)
```

* `비공개 레포지토리` - 검색된 비공개 레포지토리 목록.
* `공개 레포지토리` - 검색된 공개 레포지토리 목록.

{% hint style="info" %}
각 레포지토리 이름 옆에는 (비공개) 또는 (공개)의 가시성 표시가 있습니다.
{% endhint %}

## 기여자 세부정보

```
## 기여자 세부정보
[
    [
        "someUser",
        {
            "email": "someUser@company.io",
            "contributionsCount": 15,
            "reposContributedTo": [
                "someOrganization/someRepository(비공개)"
                "anotherOrganization/anotherRepository(공개)"
            ]
        }
    ],
    [
        "anotheruser",
        {
            "email": "anotherUser@company.io",
            "contributionsCount": 11,
            "reposContributedTo": [
                "someOrganization/someRepository(비공개)"
            ]
        }
    ],
    [
        "anotheruser(중복)",
        {
            "email": "anotherUser@otherCompany.com",
            "contributionsCount": 11,
            "reposContributedTo": [
                "someOrganization/someRepository(비공개)"
            ]
        }
    ]
]
```

* `email` - 기여자의 이메일
* `contributionsCount` - 이 기여자가 레포지토리에 기여한 횟수.
* `repoContributedTo` - 이 기여자가 기여한 레포지토리 목록.
* `(중복)` - 동일한 사용자가 다른 이메일 주소에서 감지된 것을 나타내는 표시; 다른 커미터로 **계산**될 것임을 주의하세요.

{% hint style="info" %}
출력이 길어질 수 있기 때문에 Snyk은 더 좋은 검토 및 구문 분석 옵션을 위해 출력을 파일로 보내는 것을 권장합니다.
{% endhint %}
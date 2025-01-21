# 조직의 최대 프로젝트 수

단일 Snyk 조직에서 보유할 수 있는 프로젝트 수는 Snyk [가격 요금제](https://snyk.io/plans/)에 따라 달라집니다.

| 요금제    | 프로젝트 수 |
| ------ | ------ |
| 무료     | 10,000 |
| 팀      | 25,000 |
| 엔터프라이즈 | 25,000 |

귀하의 요금제 한도에 도달하면, Snyk은 조직으로의 프로젝트 가져오기를 중단합니다.

## **한도에 도달한 경우의 알림**

한도에 도달했음을 알 수 있는 표시가 있습니다.

Snyk 웹 UI에서는 다음과 같은 경고 배너가 표시됩니다:

<figure><img src="../../.gitbook/assets/Maximum number of projects.png" alt="프로젝트 목록 상단에 나타나는 배너. 한도를 초과한 프로젝트 수가 표시됩니다"><figcaption><p>프로젝트 목록 상단에 나타나는 배너. 한도를 초과한 프로젝트 수가 표시됩니다</p></figcaption></figure>

Snyk CLI에서 `snyk monitor` 명령어를 실행하면 다음과 같은 오류가 반환됩니다:

`Maximum number of projects reached for this organization. You cannot import more projects.`

Snyk API에서는, 가져오기 요청에 대해 다음과 같은 오류가 반환됩니다:

```
"data":{
        "code":400,
        "message":"This organization has 25000 of the maximum 25000 projects.
        You will not be able to import more projects: https://docs.snyk.io/getting-started/introduction-to-snyk-projects/maximum-number-of-projects-in-an-organsation",
        "errorRef":"5bc3fb50-cbcd-4c15-81f6-b183fc95d10f"
    },
```

## 프로젝트 최대 개수에 근접했을 때 할 일

이 한도는 Snyk을 통한 경험을 보호하기 위해 존재합니다. 만들 수 있는 조직 수에는 제한이 없습니다.

이 한도에 가까워지면, 더 많은 조직을 생성하고 프로젝트를 분할하여 사용할 수 있습니다.

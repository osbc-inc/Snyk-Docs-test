# 프라이빗 컨테이너 레지스트리에 인증하기

프라이빗 컨테이너 레지스트리를 사용하는 경우, 레지스트리로의 인증 정보를 포함하는 `dockercfg.json` 파일을 생성해야 합니다. 그런 다음, `snyk-monitor`로 불려야 하는 시크릿을 생성해야 합니다.

`dockercfg.json` 파일은 Monitor가 프라이빗 레지스트리에서 이미지를 찾을 수 있도록 필요합니다. 보통 사용자 인증 정보는 `$HOME/.docker/config.json`에 저장됩니다. 그러나, 인증 정보는 `dockercfg.json` 파일에도 추가되어 있어야 합니다. 사용자 인증 정보가 `$HOME/.docker/config.json`에만 저장되어 있는 경우, Snyk Controller는 이러한 레지스트리에 액세스할 수 없습니다.

아래 단계에서는 프라이빗 컨테이너 레지스트리에 인증하는 방법을 설명합니다.

## dockercfg.json 파일 구성

`dockercfg.json` 파일을 생성하고 이 파일에 사용자 인증 정보를 저장합니다.

{% hint style="info" %}
인증 정보가 포함된 파일을 `dockercfg.json`으로 명명하는 것이 중요합니다. 이 파일명은 `snyk-monitor`에서 필요합니다.
{% endhint %}

{% hint style="info" %}
`dockercfg.json` 파일의 서식이 올바르게 유지되어야 합니다. 줄 바꿈 문자와 공백이 포함되어 있어야 합니다. 올바르지 않은 파일은 인증 실패로 이어질 수 있습니다.
{% endhint %}

클러스터가 실행되는 위치와 레지스트리가 실행되는 위치에 따라 `dockercfg.json` 파일에 입력할 내용이 달라집니다. 파일은 여러 레지스트리에 대해 인증 정보를 포함할 수 있습니다.

이미 `$HOME/.docker/config.json`에 사용자 인증 정보가 있는 경우, 이 정보를 `dockercfg.json` 파일로 복사합니다.

만약 `$HOME/.docker/config.json` 파일에 `auth` 항목이 비어있다면, 다음 명령을 실행하고 출력을 `dockercfg.json`의 `auth` 항목에 붙여넣기 하세요:

```
echo -n 'username:password' | base64
```

### dockercfg.json 파일 구성 예제

#### Nexus 이외의 프라이빗 레지스트리

만약 클러스터가 `GKE`에서 실행되지 않거나, `GKE`에서 실행 중이지만 다른 프라이빗 레지스트리에서 이미지를 가져온다면, `dockercfg.json` 파일에는 다음이 포함되어야 합니다:

```json
{  
  "auths": {
    "gcr.io": {
      "auth": "BASE64로 인코딩된 인증 정보"
    },
    // 필요한 경우 다른 레지스트리를 추가합니다. 예시:
    "<yourdomain>.azurecr.io": {
      "auth": "BASE64로 인코딩된 인증 정보"
    }
  }
}
```

#### Nexus 레지스트리

Nexus 레지스트리를 사용하는 경우, `dockercfg.json` 파일은 다음과 같아야 합니다:

```json
{
  "auths": {
    "<registry>": {
        "auth": "BASE64로 인코딩된 인증 정보"
    },
  }
}
```

#### Artifactory 컨테이너 레지스트리

여러 개인 저장소를 호스팅하는 Artifactory 컨테이너 레지스트리를 사용하는 경우, `dockercfg.json` 파일은 다음과 같아야 합니다:

```json
{
  "auths": {
    "<registry1>": {
        "auth": "BASE64로 인코딩된 인증 정보"
       },
    "<registry2>": {
       "auth": "BASE64로 인코딩된 인증 정보"
    }
  }
}
```

#### GKE에서 GCR을 사용하는 경우

클러스터가 `GKE`에서 실행되고 `GCR`를 사용하는 경우, `dockercfg.json` 파일은 다음과 같아야 합니다:

```json
{
  "credHelpers": {
    "us.gcr.io": "gcloud",
    "asia.gcr.io": "gcloud",
    "marketplace.gcr.io": "gcloud",
    "gcr.io": "gcloud",
    "eu.gcr.io": "gcloud",
    "staging-k8s.gcr.io": "gcloud"
  }
}
```

#### GKE에서 Google Artifact Registry (GAR)를 사용하는 경우

클러스터가 `GKE`에서 실행되고 `GAR`를 사용하는 경우, `dockercfg.json` 파일은 다음과 같아야 합니다:

```json
{
  "auths": {
	"northamerica-northeast2-docker.pkg.dev": {
  	"auth": "“echo -n _json_key_base64:BASE64로 인코딩된 인증 정보” 결과"
  }
}

```

이 방식은 서비스 계정을 생성하는 것에 의존합니다. [Google Cloud 서비스 계정 키](https://cloud.google.com/artifact-registry/docs/docker/authentication#json-key)를 참조하십시오. 옵션 단계를 따라 파일을 base64로 인코딩하는 것이 중요합니다.

`“auth”` 줄은 다음 명령어로 생성됩니다. 여기서 사용자 이름은 json\_key\_base64이고, 비밀번호는 base64 json 키 파일의 전체 내용입니다.

```sh
echo -n 'username:password' | base64 
```

예를 들어, 이 명령어의 출력은 `dockercfg.json` 파일의 `“auth”` 줄에 사용됩니다.

```sh
echo -n ‘json_key_base64:ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAic255ay1jeC1zZS1kZW1vIiwKICAicHJpdmF0ZV9rZXlfaWQiOiAiMDNhOTQyMWNhMDhkZTA0MzIyNTc4OTFjMDRiNGFjZmJlYTM0Y2U1ZCIsCiAgInByaXZhdGVfa2V5IjogIi0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLVxuTUlJRXZRSUJBREFOQmdrcWhraUc5dzBCQVFFRkFBU0NCS2N3Z2dTakFnRUFBb0lCQVFEQkpxT2M0NzFIMjFIOVxuSjBKMi9zRTlJb2tmY1N0SjVKbEYyYTg3bGRZTHd6QkNmcW5UUk1lSGl6RGgrTXFCVWxjdE01ZVVFdFJXRmkwMFxud2NrVTN4eTNMVlRMekhnNWVMeGZ0MHVQZWtRRVBUZ3RQMW3lsk3tzVHEydms2T1VmTHI2ZHh0TXpUQzZtMWdzQVxuYS9rTm9mMmE1TkN6aHZ5d1NqU1E0dTdMWHBhWUhNY2ZUN0lTZmZJSEJVMi9OdDlWZUpYc3F0b3l3YXNsY3l5TVxub0dEV3JTL3BqblVWZzkxOXVvaUhBaE9MR0piMFczeWQwYWR2Z3ZPMGFrbzNzTWcxbkJiSTd2S1JEaE1aNDlETVxucWZPakNOZm56UUNKREdFOHB5akhUdWNxdWx0dkxxUjFBSUh5Sy9pNG53amNNdjB0MXR0d0xaOFBxSGpHd0ZPMlxuK3FsUnlHVTdBZ01CQUFFQ3lsv0RUFEV1RCQnJHU0lBZjQxUDJpY2I3enBtb2RLUlYvWTNYYkhRbGR5ZHQzaHNSSFxudGV5em1RZFZjTFE1dFFtNy9TQzVFOVRXaDNtUXlORnIzQk10L3VrRHNuMk51ajRZL2g4OXJNTjRsVi9zbElDc1xuUXhMM3o0OHhHMkdSTDBQcEUzTDUyOVg4TWp2dnBqa0VkVWlIY2ljUC8yd3JmcTkzR1VCa0NjSDZ2aFoyaWVDaVxuR1l1QzBRVGs1eXkwZTNnV05FMkpLdkk4WkI0Y2dwUHpKMnFhRHdWSXRHV3FOY3hwbWp6WjhNSS8wTXNaQS9IeFxuNTQ0MVpISkJKZXRzWVdxV2ZkWGV0cG1EN24vRVVzTFlQTmtRTnBodTE2Z2o5LzR3Wk9RaXhoTGpTWXJ5OVdqdFxuL2JtYlkyQ2xtRmVybFB5UjdMUElQeno1TXZqVGI4dHVFRjZNWG1nc0xRS0JnUURxMkNUcmVKZ05IUDB2YVprRVxuMXU0Tk5naTJxSkxZengrZHN2SENCVHgzREpvdWtXejd3Tzh2OXFjOWg2RkptalRzTFVUNUdZUVFDTjJ4V29zb1xubEJYcE1QY1g0VmNPTGlKNFRoSGJHYmV1RFhnZXZFYnVHOU90amtxVk1xdnQvaVBJL3pzRDhGL0ZxbmxSTkdxSlxuZjFmRjcwTWIxdGMwa0EvUWhTajNQK2lxdFFLQmdRRFNqUHZ4bEtXMVhjSm1KZEN1NWxHck1JTS8vdC9INTdpYVxuc0FLa0JaektBWVZhSHh1UUdxeUZ2RjRTSlRhelNEMUdRQ2RVcExZTUJhWGZpb25DTWVIYS9CWEsrQS9kVVd5OFxuUDFFSG9MNHJTTVZsMVl6WHRFWGh5b3FDcS9lSFRrcStQa2R4Z0pJSjBpcktvV0Y3ZDk3OTVqNmVtaStEeHYyeFxuN3VQOWNiK1dMd0tCZ0Y4WG9ITjhoRTBqQk40eTZ4UUxsNTdmMTAxbkd2Y1JmMkxTdDVQeG5OY3owaWF6R2ljaVxucTNlSGI1YTVtYlI4N1pzSWhabzhHNzZHYUlaTS9IWTA2RjVoUmx4MEVWVWJsemVSblNkVDFZMXp4TVRsUmU5YVxuY3k4ZW85S2dEd0F5WFBraGFCc2pOUlNMLzgzQzVMVENUSjlJVDZzeEpqa1JjR1hsMVgyd2NoelZBb0dCQUxIbFxubHhZRi80RGZLRnFBUnZNUC9SOEVUVkVyKzAzL1ZuVzBrM2FjbTEzK3JQcDVZQ09BdGhZRkV3S0gyTkRnRDQya1xudE5hS21KcE54MW01eHkyQ1VnOWhnTlJPaGJEOGxELzF5M1FEZDhwQW9UQ3FuMmE5bFhIeVhOZU5qd1lPdTQ1RVxuTnI4SzM5bFdidnRvSVdKZDVOWm56SzdiSFp4YzdJdURpYlRoZi92WEFvR0FWNkhWMmIwS0hFbjloQURQOXZITVxuZzg1Nm0yVnZUTGhjNmZaS3lDY28zWDYwZTNHaDJNLzUyZ3YxMTlWTGFvdWlFbzc1YWVYejQxNElQWCVnbUVHMkxHNFFcblxuLzJ6VEdLejFXYzMzenhPWnBYWEc0PSIsCiAgImNsaWVudF9lbWFpbCI6ICJzcGVyY2lib2xsaS1zbnlrbW9uaXRvckBzbmtlay1jeC1zZS1kZW1vLmlhbS5nc2VydmljZWFjY291bnQuY29tIiwKICAiY2xpZW50X2lkIjogIjEwOTE4NDIxMDg1NTc0MTIzOTUxMyIsCiAgImF1dGhfdXJpIjogImh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbS9vL29hdXRoMi9hdXRoIiwKICAidG9rZW5fdXJpIjogImh0dHBzOi8vb2F1dGgyLmdvb2dsZWFwaXMuY29tL3Rva2VuIiwKICAiYXV0aF9wcm92aWRlcl94NTA5X2NlcnRfdXJsIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL29hdXRoMi92MS9jZXJ0cyIsCiAgImNsaWVudF94NTA5X2NlcnRfdXJsIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL3JvYm90L3YxL21ldGFkYXRhL3g1MDkvc3BlcmNpYmFsbGktc255a21vbml0b3IlNDBzbnlrLWN4LXNlLWRlbW8uaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLAogICJ1bml2ZXJzZV9kb21haW4iOiAiZ29vZ2xlYXBpcy5jb20iCn0K
```

#### EKS에서 ECR을 사용하는 경우

클러스터가 `EKS`에서 실행되고 `ECR`을 사용하는 경우, 다음을 추가합니다:

```json
{
  "credsStore": "ecrypted"
}
``````json
{
  "credHelpers": {
    "public.ecr.aws": "ecr-login",
    "<aws_account_id>.dkr.ecr.<region>.amazonaws.com": "ecr-login"
  }
}
```

특정 `ECR` 레지스트리에 대해 이 자격 증명 도우미를 사용하려면, ECR 레지스트리의 URI로 `credHelpers` 섹션을 생성하세요:

```json
{
  "credHelpers": {
    "public.ecr.aws": "ecr-login",
    "<aws_account_id>.dkr.ecr.<region>.amazonaws.com": "ecr-login"
  }
} 
```

#### AKS에서 ACR 사용하는 경우

클러스터가 `AKS`에서 실행되고 `ACR`을 사용하는 경우 다음을 추가하세요:

```json
{
  "credHelpers": { 
    "myregistry.azurecr.io": "acr-env"
  }
}
```

{% hint style="info" %}
또한, AKS에서 실행되고 ACR을 사용하는 클러스터의 경우, [Entra ID Workload Identity 서비스 계정](https://azure.github.io/azure-workload-identity/docs/topics/service-account-labels-and-annotations.html#service-account)을 참조하세요. `snyk-monitor` ServiceAccount에 레이블 및 어노테이션을 구성해야 할 수도 있습니다.
{% endhint %}

다른 레지스트리에 대해 서로 다른 자격 증명 도우미를 구성할 수 있습니다.

## Kubernetes 비밀 생성

다음 명령을 실행하여 Kubernetes에 비밀을 생성하세요:

{% code overflow="wrap" %}
```sh
kubectl create secret generic snyk-monitor -n snyk-monitor \ 
        --from-file=./dockercfg.json \
        --from-literal=integrationId=abcd1234-abcd-1234-abcd-1234abcd1234 \
        --from-literal=serviceAccountApiToken=aabb1212-abab-1212-dcba-4321abcd4321
```
{% endcode %}
```
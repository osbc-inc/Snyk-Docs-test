# 지원되는 워크로드, 컨테이너 레지스트리, 언어 및 운영 체제

## 지원되는 워크로드

Snyk Controller는 클러스터에서 다음과 같은 워크로드를 감지할 수 있습니다:

* 배포 (Deployment)
* 레플리카셋 (ReplicaSet)
* 복제컨트롤러 (ReplicationController)
* 데몬셋 (DaemonSet)
* 스테이트풀셋 (StatefulSet)
* 작업 (Job)
* 크론잡 (CronJob)
* 배포 설정 (DeploymentConfig) (OpenShift)
* Pod. 이 Pod가 부모 또는 소유 참조가 없는 경우.

Controller는 개별 Pod부터 최상위 워크로드에 이르기까지의 소유 참조 체인을 추적하여 이러한 워크로드를 감지합니다.

예를 들어 Pod가 주어졌을 때, Controller는 해당 Pod가 레플리카셋에 의해 소유되었음을 감지할 수 있으며, 이 레플리카셋이 다시 배포에 의해 소유되었음을 감지할 수 있습니다. 따라서 감지된 워크로드는 배포임.

워크로드가 사용자 정의 리소스에 소유되는 경우에는 `snyk monitor`가 현재 진행할 수 없으며, Controller가 처리할 수 있는 최상위 부모 워크로드로 간주해야 합니다.

## 지원되는 컨테이너 레지스트리

Snyk Controller는 다음과 같은 컨테이너 레지스트리를 지원합니다:

| 컨테이너 레지스트리                                                                                                                                         | 설치 문서                                                                                                                                                                                                     |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p>ACR<br>GCR<br>DigitalOcean</p><p>ECR<br>Github Container Registry<br>Gitlab Container Registry<br>Harbor<br>JFrog Artifactory<br>Docker Hub</p> | [install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md](../install-the-snyk-controller/install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md "mention")             |
| OCR                                                                                                                                                | [install-the-snyk-controller-with-openshift-4-and-operatorhub.md](../install-the-snyk-controller/install-the-snyk-controller-with-openshift-4-and-operatorhub.md "mention")                               |
| <p>ECR<br>JFrog Artifactory</p>                                                                                                                    | [install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md](../install-the-snyk-controller/install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md "mention") |

## 지원되는 언어

Snyk Controller는 다음 언어를 지원합니다:

* Node
* Go
* Java
* PHP
* Python

## 지원되는 운영 체제

[Snyk Container에서 지원되는 운영 체제 배포](../../how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md)를 참조하세요.

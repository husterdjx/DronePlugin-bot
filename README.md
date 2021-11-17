```
kind: pipeline
type: kubernetes
name: default
trigger:
  event:
    - push
steps:
- name: webhook
  image: registry.cn-beijing.aliyuncs.com/husterdjx/droneplugin-bot:latest
  settings:
    author: ${DRONE_COMMIT_AUTHOR}
    branch: ${DRONE_COMMIT_BRANCH}
    repourl: ${DRONE_REPO_LINK}
    message: ${DRONE_COMMIT_MESSAGE}
    githash: ${DRONE_COMMIT}
```
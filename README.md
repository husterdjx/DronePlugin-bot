kind: pipeline
type: docker
name: default

steps:
- name: webhook
  image: {}
  settings:
    author: DRONE_COMMIT_AUTHOR
    branch: DRONE_COMMIT_BRANCH
    repourl: DRONE_REPO_LINK
    message: DRONE_COMMIT_MESSAGE
    githash: DRONE_COMMIT
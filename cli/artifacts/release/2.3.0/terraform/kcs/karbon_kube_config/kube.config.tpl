# -*- mode: yaml; -*-
# vim: syntax=yaml
#
apiVersion: v1
kind: Config
clusters:
- name: ${cluster_name}
  cluster:
    server: ${host}
    certificate-authority-data: ${cluster_ca_certificate}
users:
- name: default-user-${cluster_name}
  user:
    token: ${token}
contexts:
- context:
    cluster: ${cluster_name}
    user: default-user-${cluster_name}
  name: ${cluster_name}-context
current-context: ${cluster_name}-context
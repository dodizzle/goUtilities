curl -G https://mimir-query-frontend.infra.alto.com/prometheus/api/v1/query \
  -H "Authorization: Basic $(echo -n "$MIMIR_FE_USERNAME:$MIMIR_FE_PASSWORD" | base64)" \
  -H "x-scope-orgid: stg-cluster" \
  --data-urlencode 'query=sum(rate(container_cpu_usage_seconds_total{namespace="opencost", pod=~"opencost-.*"}[5m])) by (pod)' | jq
echo -e "\n"

curl -G https://mimir-query-frontend.infra.alto.com/prometheus/api/v1/query \
  -H "Authorization: Basic $(echo -n "$MIMIR_FE_USERNAME:$MIMIR_FE_PASSWORD" | base64)" \
  -H "x-scope-orgid: stg-cluster" \
  --data-urlencode 'query=max(container_memory_usage_bytes{namespace="opencost", pod=~"opencost-.*"}) by (pod)' | jq




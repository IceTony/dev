# Получить среднее число коммитов в день за прошедший месяц (используя Gitlab API)
sudo apt update &&
sudo apt install curl jq bc -y &&
echo "scale=1; $(curl -s -X GET -H "PRIVATE-TOKEN: GITLAB-API-TOKEN" "https://gitlab.rebrainme.com/api/v4/projects/3006/repository/commits?per_page=100&since='$(date +%Y)-$(date -d "$(date +%Y-%m-1) -1 month" +%-m)-$(date +%d)'" | jq '. | length') / 30" | bc | sed 's/^\./0./'

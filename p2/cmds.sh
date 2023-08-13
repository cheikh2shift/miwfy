curl -X POST http://localhost:3000/posts \
   -H 'Content-Type: application/json' \
   -d '{"text" : "hello world" }' | json_pp

echo "\n"

curl http://localhost:3000/posts | json_pp

echo "\n"

curl -X POST http://localhost:3000/posts/comment \
   -H 'Content-Type: application/json' \
   -d '{"comment" : "hello world" , "post_id" : 1 }' | json_pp

echo "\n"

curl http://localhost:3000/posts/1 | json_pp
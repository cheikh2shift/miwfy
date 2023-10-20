export PROJECT_NAME = /air-live

run-container-air:
	docker run -it --rm \
    -w $$PROJECT_NAME \
    -e "air_wd=$$PROJECT_NAME" \
    -v $$(pwd):$$PROJECT_NAME \
    cosmtrek/air
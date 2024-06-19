.PHONY: ${SERVICE}


${SERVICE}:
	kubectl delete deployment ${SERVICE}-service --ignore-not-found
	docker rmi ${SERVICE}_image:latest || true
	docker build --build-arg SERVICE=${SERVICE} -t ${SERVICE}_image .
	minikube image load ${SERVICE}_image:latest
	cd deploy && kubectl apply -f ${SERVICE}_deployment.yaml

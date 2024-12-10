bd:
	docker build -t api-travel-destinations-backend ./backend/ && \
	docker run -p 8090:8080 api-travel-destinations-backend

fd:
	docker build -t api-travel-destinations-frontend ./frontend/ && \
	docker run -p 8090:8080 api-travel-destinations-frontend
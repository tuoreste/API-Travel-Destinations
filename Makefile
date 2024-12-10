bd: #backend
	docker build -t api-travel-destinations-backend ./backend/ && \
	docker run -p 8090:8080 api-travel-destinations-backend

fd: #frontend
	docker build -t api-travel-destinations-frontend ./frontend/ && \
	docker run -p 8090:8080 api-travel-destinations-frontend

run: #run project
	docker-compose up --build

red: #empty everything and lose all data in docker!!!!!!!!!!!!!
	docker system prune
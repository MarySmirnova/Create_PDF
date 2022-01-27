build:
	docker build -t fillpdf .

run:
	docker run -p 8000:8000 --rm -d fillpdf
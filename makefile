
make:
	go build 
	./fload

mobile:
	go build -tags mobile
	fyne package -os android -appID my.domain.fload
	fyne install -os android

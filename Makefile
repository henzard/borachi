install:
	cp build/bin/${BINNAME}_${GOOS}_${GOARCH} /usr/bin/borachi
	cp logo.png /usr/share/icons/hicolor/scalable/apps/borachi.png
	cp borachi.desktop /usr/share/applications
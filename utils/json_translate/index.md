## JSON translate file

> This is a simple tool to translate any `.json` file, without changing its keys

### Before you need to execute
```bash
docker run -ti --name local-libretranslate -p 5000:5000 libretranslate/libretranslate --load-only=en,es,pt
```
For this example only load the pairs of languages:
- en to pt
- en to pt
- pt to en
- es to en
List of languages available in https://www.argosopentech.com/argospm/index/
> More info in https://github.com/LibreTranslate/LibreTranslate

### Run
Add files to translate into dir `files`, all files need to be `.json`
```bash
go run json_translate/main.go "json_translate" "es" "en"
```
The output files are located into dir `output`
# cedictToJson
Converts the cedict Chinese/English dictionary to computer friendly JSON. Converts numbered pinyin to UTF-8 Pinyin with tone marks. for example dian4 nao3 becomes diàn nǎo. Written in Golang

Usage: go run cedicttojson.go <cedict_ts.u8 file>
       A new json file will be created in your working directory

Note: cedict file must not be compressed

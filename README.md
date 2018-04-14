# cedictToJson
Converts the cedict Chinese/English dictionary to computer friendly JSON. Converts numbered pinyin to UTF-8 Pinyin with tone marks. for example dian4 nao3 becomes diàn nǎo. Written in Golang

Usage: 
             
             go run cedicttojson.go <cedict_ts.u8 file>
      
      
A new json file will be created in your working directory


Example,
             你好 你好 [ni3 hao3] /Hello!/Hi!/How are you?/

Will be converted to
             {
                 "Traditional": "你好",
                 "Simplified": "你好",
                 "PinyinNumbered": "ni3 hao3",
                 "Pinyin": "nǐ hǎo",
                 "Definition": "Hello!;Hi!;How are you?;"
             },



Note: cedict file must not be compressed

#!/usr/bin/env bash

# API dökümanı oluştur. Hata varsa akışı kes.
echo "API dökümanı oluşturuluyor..."
OUTPUT=$(swagger generate spec -i input.yml -o images/swagger.json 2>&1)
if [ "$OUTPUT" ]; then
   echo $OUTPUT
   echo "API dökümanı oluşturulamadı, işleme devam edilemiyor :("
   exit 1
fi


# API dökümanını doğrula. Hata varsa akışı kes.
echo "Oluşturulan API dökümanı swagger 2.0 versiyonuna göre doğrulanıyor..."
OUTPUT=$(swagger validate images/swagger.json 2>&1)
if echo $OUTPUT |grep -q 'errors :'; then
   echo $OUTPUT
   echo "API dökümanı hatalı, işleme devam edilemiyor :("
   exit 1
fi

# Hata yok hayata devam et.
echo "API dökümanı geçerli :)"
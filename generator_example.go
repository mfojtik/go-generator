package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/mfojtik/go-generator/generator"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	result, err := generator.Template("[GET:http://localhost:9292/]")
	log.Printf("urlResult=%v", result)

	result, _ = generator.StrongPassword(10)
	log.Printf("password=%v", result)

	result, _ = generator.SimplePassword(12)
	log.Printf("password=%v", result)

	result = generator.RandomUUID()
	log.Printf("uuid=%v", result)

	result = generator.RandomCleanUUID()
	log.Printf("cleanUuid=%v", result)

	result = generator.RandomNumericUUID()
	log.Printf("numericUuid=%v", result)

	result = generator.RandomCleanNumbericUUID()
	log.Printf("numericCleanUuid=%v", result)

	result, _ = generator.Template("password")
	log.Printf("password=%v", result)

	result, _ = generator.Template("uuid")
	log.Printf("uuid=%v", result)

	// generator.Template 8 character long random string that contains *all* Ascii
	// characters (See Ascii const)
	result, _ = generator.Template(`foo[\w]{8}bar`)
	log.Printf("result0=%v", result)

	// generator.Template 3 character long random suffix using just lowercase alpha+numberic
	// characters
	result, _ = generator.Template(`admin[a-z0-9]{3}`)
	log.Printf("result1=%v", result)

	result, _ = generator.Template(`admin[a-z0-9]{3}something[\w]{3}`)
	log.Printf("result1=%v", result)

	// generator.Template 12 character long random password combining all alphabetic characters
	// and numeric characters
	result, _ = generator.Template(`pass[a-zA-Z0-9]{12}`)
	log.Printf("result2=%v", result)

	// generator.Template 8 character long password by using alphabetic characters
	// lowercase + uppercase
	result, _ = generator.Template(`pass[a-Z]{8}`)
	log.Printf("result3=%v", result)

	// Invalid range error
	result, err = generator.Template(`pass[z-a]{8}`)
	log.Printf("result4=%v; err=%v", result, err)

	// generator.Template 8 random numeric characters
	result, _ = generator.Template(`foo[\d]{8}bar`)
	log.Printf("result5=%v", result)
}

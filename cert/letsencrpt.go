package cert

// func acceptLetsEncryptTOS(url string) bool {
// 	return url == "https://letsencrypt.org/documents/LE-SA-v1.1.1-August-1-2016.pdf"
// }

// func decodePEMKey(data []byte) (crypto.Signer, error) {
// 	block, _ := pem.Decode(data)
// 	if block.Type == "RSA PRIVATE KEY" {
// 		return x509.ParsePKCS1PrivateKey(block.Bytes)
// 	}
// 	return nil, fmt.Errorf("Unknown private key type: ", block.Type)
// }

// type ChallengeHandler func(context.Context, *GCSUtil, *acme.Client, *acme.Challenge)

// var challengeHandler = map[string]ChallengeHandler{
// 	"http-01": startHTTP01Challenge,
// }

// func Request() {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		panic("Failed to create GCS client.")
// 	}
// 	defer client.Close()

// 	g := &GCSUtil{client: client}

// 	var signer crypto.Signer
// 	data, err := g.Load(ctx, "secret/private_key.pem")
// 	if err != nil {
// 		if err != storage.ErrObjectNotExist {
// 			log.Fatalf("Storage error: %s", err)
// 		}

// 		key, err := rsa.GenerateKey(rand.Reader, 2048)
// 		if err != nil {
// 			log.Fatalf("Failed to generate private key: %s", err)
// 		}
// 		der := x509.MarshalPKCS1PrivateKey(key)
// 		data := pem.EncodeToMemory(&pem.Block{
// 			Type:  "RSA PRIVATE KEY",
// 			Bytes: der,
// 		})
// 		err = g.Store(ctx, "secret/private_key.pem", data)
// 		if err != nil {
// 			log.Fatalf("Failed to store private key: %s", err)
// 		}

// 		signer = key
// 	} else {
// 		key, err := decodePEMKey(data)
// 		if err != nil {
// 			log.Fatalf("Failed to decode private key: %s", err)
// 		}

// 		signer = key
// 	}

// 	c := &acme.Client{
// 		Key: signer,
// 	}

// 	var acc *acme.Account
// 	data, err = g.Load(ctx, "secret/letsencrypt_id")
// 	if err != nil {
// 		acc, err = c.Register(ctx, nil, acceptLetsEncryptTOS)
// 		if err != nil {
// 			log.Fatalf("Failed to create ACME account: %s", err)
// 		}

// 		err = g.Store(ctx, "secret/letsencrypt_id", []byte(acc.URI))
// 		if err != nil {
// 			log.Fatalf("Failed to store ACME account info: %s", err)
// 		}
// 	} else {
// 		acc, err = c.GetReg(ctx, string(data))
// 		if err != nil {
// 			log.Fatalf("Failed to get ACME account: %s", err)
// 		}
// 	}

// 	if acc.Contact == nil || len(acc.Contact) == 0 {
// 		acc.Contact = []string{"mailto:mail@tzik.jp"}
// 		acc, err = c.UpdateReg(ctx, acc)
// 		if err != nil {
// 			log.Fatalf("Failed to update ACME account: %s", err)
// 		}
// 	}

// 	auth, err := c.Authorize(ctx, "tzik.jp")
// 	if err != nil {
// 		log.Fatalf("Failed to initiate ACME authorization: %s", err)
// 	}

// 	if auth.Status == acme.StatusValid {
// 		// Done.
// 		return
// 	}

// 	if auth.Status != acme.StatusPending {
// 		log.Fatalf("Unexpected state of ACME authorization: %s", auth.Status)
// 	}

// 	var combination []int = nil
// 	for _, cm := range auth.Combinations {
// 		ok := true
// 		for _, i := range cm {
// 			if _, b := challengeHandler[auth.Challenges[i].Type]; !b {
// 				ok = false
// 				break
// 			}
// 		}
// 		if ok {
// 			combination = cm
// 			break
// 		}
// 	}

// 	for _, i := range combination {
// 		ch := auth.Challenges[i]
// 		challengeHandler[ch.Type](ctx, g, c, ch)
// 	}
// 	c.WaitAuthorization(ctx, "tzik.jp")
// }

// func startHTTP01Challenge(ctx context.Context, g *GCSUtil, c *acme.Client, ch *acme.Challenge) {
// 	path := c.HTTP01ChallengePath(ch.Token)
// 	prefix := "/.well-known/acme-challenge/"
// 	if !strings.HasPrefix(path, prefix) {
// 		log.Fatalf("Unexpected ACME http-01 path: %s", path)
// 	}

// 	res, err := c.HTTP01ChallengeResponse(ch.Token)
// 	if err != nil {
// 		log.Fatalf("Failed to make ACME http-01 response: %s", err)
// 	}

// 	err = g.Store(ctx, "challenge/"+path[len(prefix):], []byte(res))
// 	if err != nil {
// 		log.Fatalf("Failed to store ACME http-01 response: %s", err)
// 	}

// 	ch, err = c.Accept(ctx, ch)
// }

// // func RenewCertificate(w http.ResponseWriter, r *http.Request) {
// // }

// func HandleHTTP01Challenge(w http.ResponseWriter, r *http.Request) {
// 	prefix := "/.well-known/acme-challenge/"
// 	path := r.URL.Path
// 	if !strings.HasPrefix(path, prefix) {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		panic("Failed to create GCS client.")
// 	}
// 	defer client.Close()

// 	g := &GCSUtil{client: client}
// 	data, err := g.Load(ctx, "challenge/"+path[len(prefix):])
// 	if err != nil {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	for len(data) != 0 {
// 		n, err := w.Write(data)
// 		if err != nil {
// 			return
// 		}
// 		data = data[n:]
// 	}
// }

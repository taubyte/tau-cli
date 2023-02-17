package domainPrompts

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	domainFlags "github.com/taubyte/tau/flags/domain"
	domainLib "github.com/taubyte/tau/lib/domain"
	"github.com/taubyte/tau/prompts"
	"github.com/urfave/cli/v2"
)

func certificate(ctx *cli.Context, domain *structureSpec.Domain, new bool) (err error) {
	defaultCertType := domainFlags.CertTypeAuto
	if new == false {
		defaultCertType = domain.CertType
	}

	domain.CertType, err = getCertType(ctx, defaultCertType)
	if err != nil {
		return
	}

	if domain.CertType == domainFlags.CertTypeInline {
		if new == true {
			domain.CertFile = GetOrRequireACertificate(ctx, CertificateFilePrompt)
			domain.KeyFile = GetOrRequireAKey(ctx, KeyFilePrompt)
		} else {
			domain.CertFile = GetOrRequireACertificate(ctx, CertificateFilePrompt, domain.CertFile)
			domain.KeyFile = GetOrRequireAKey(ctx, KeyFilePrompt, domain.KeyFile)
		}

		var (
			cert []byte
			key  []byte
		)
		cert, key, err = domainLib.ValidateCertificateKeyPairAndHostname(domain)
		if err != nil {
			// TODO verbose
			return
		}

		domain.CertFile = string(cert)
		domain.KeyFile = string(key)
	}

	return nil
}

func getCertType(ctx *cli.Context, defaultCertType string) (certType string, err error) {
	certType, isSet, err := domainFlags.GetCertType(ctx)
	if err != nil {
		return
	}

	if isSet == false {
		certType, err = prompts.SelectInterface(domainFlags.CertTypeOptions, CertificateTypePrompt, defaultCertType)
		if err != nil {
			return
		}
	}

	return
}

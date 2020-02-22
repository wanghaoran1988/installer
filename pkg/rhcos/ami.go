package rhcos

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/pkg/errors"
)

// AMI fetches the HVM AMI ID of the Red Hat Enterprise Linux CoreOS release.
func AMI(ctx context.Context, region string) (string, error) {
	if oi, ok := os.LookupEnv("OPENSHIFT_INSTALL_OS_IMAGE_OVERRIDE"); ok && oi != "" {
		logrus.Warn("Found override for OS Image. Please be warned, this is not advised")
		return oi, nil
	}
	meta, err := fetchRHCOSBuild(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to fetch RHCOS metadata")
	}

	ami, ok := meta.AMIs[region]
	if !ok {
		return "", errors.Errorf("no RHCOS AMIs found in %s", region)
	}

	return ami.HVM, nil
}

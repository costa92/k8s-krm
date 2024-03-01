
# Include

# include the common make file
ifeq ($(origin KRM_ROOT), undefined)
KRM_ROOT  :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
endif

# include the common make file
include $(KRM_ROOT)/scripts/make-rules/common-version.mk

SHELL := /usr/bin/env bash -o errexit -o pipefail +o nounset
.SHELLFLAGS = -ec

export SHELLOPTS := errexit
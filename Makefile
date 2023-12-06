CURRENTPATH :=$(shell pwd)
#static paths
ACCOUNTFILE:=UTC--2023-12-05T06-45-42.679945500Z--039cc1e608f481a8349f8a4c7b4dfcabc2b453b9

#node eth variables
DATAPATH :=$(CURRENTPATH)/l1chain/data
GENESISPATH :=$(CURRENTPATH)/genesis.json
KEYSTOREPATH:=$(DATAPATH)/keystore

IPFSNODE1PATH=$(CURRENTPATH)/ipfs/node1
STATICDATAPATH:=$(CURRENTPATH)/static_data
SWARMPATH=$(CURRENTPATH)/swarm.key

INITCMD :=docker run -v $(DATAPATH):/data -v $(GENESISPATH):/genesis.json
INITCMD +=ethereum/client-go init --datadir /data /genesis.json

ifneq ($(shell [ -d ${DATAPATH} ] && echo "true"),true)
$(shell mkdir -p $(DATAPATH))
$(shell mkdir -p $(KEYSTOREPATH))
endif
.PHONY: install
install:
		$(shell docker pull ipfs/kubo)
.PHONY: init
init:
		$(shell $(INITCMD))
.PHONY: copy
copy:
		$(shell cp $(STATICDATAPATH)/$(ACCOUNTFILE) $(KEYSTOREPATH)/$(ACCOUNTFILE))
		$(shell cp $(STATICDATAPATH)/passwordfile $(DATAPATH)/passwordfile)

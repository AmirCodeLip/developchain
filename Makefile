CURRENTPATH :=$(shell pwd)
DATAPATH :=$(CURRENTPATH)/l1chain/data
IPFSNODE1PATH=$(CURRENTPATH)/ipfs/node1
GENESISPATH :=$(CURRENTPATH)/genesis.json
SWARMPATH=$(CURRENTPATH)/swarm.key
INITCMD :=docker run -v $(DATAPATH):/data -v $(GENESISPATH):/genesis.json
INITCMD +=ethereum/client-go init --datadir /data /genesis.json
NODE1IPFSCMD :=docker run -d --name node1_ipfs -e IPFS_SWARM_KEY_FILE=$(CURRENTPATH)/swarm.key  -v $(IPFSNODE1PATH)/ipfs_staging:/export -v $(IPFSNODE1PATH)/ipfs_data:/data/ipfs
NODE1IPFSCMD+= ipfs/kubo init
ifneq ($(shell [ -d ${DATAPATH} ] && echo "true"),true)
$(shell mkdir -p $(DATAPATH))
$(shell mkdir $(DATAPATH)/l1chain)
$(shell cp $(CURRENTPATH)/l1chain/keystore $(CURRENTPATH)/l1chain/data/keystore)
endif
.PHONY: install
install:
		$(shell docker pull ipfs/kubo)
.PHONY: init
init:
		$(shell $(NODE1IPFSCMD))
		$(shell $(NODE1IPFSCMD))


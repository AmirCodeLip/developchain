CURRENTPATH :=$(shell pwd)
DATAPATH :=$(CURRENTPATH)/l1chain/data
IPFSNODE1PATH=$(CURRENTPATH)/ipfs/node1
GENESISPATH :=$(CURRENTPATH)/genesis.json
SWARMPATH=$(CURRENTPATH)/swarm.key
INITCMD :=docker run -v $(DATAPATH):/data -v $(GENESISPATH):/genesis.json
INITCMD +=ethereum/client-go init --datadir /data /genesis.json
NODE1IPFSCMD :=docker run -d --name node1_ipfs -v $(IPFSNODE1PATH)/ipfs_staging:/export -v $(IPFSNODE1PATH)/ipfs_data:/data/ipfs ipfs/kubo init
ifneq ($(shell [ -d ${DATAPATH} ] && echo "true"),true)
$(shell mkdir -p $(DATAPATH))
$(shell cp l1chain/keystore l1chain/data/keystore)
endif
.PHONY: install
install:
		$(shell docker pull ipfs/kubo)
.PHONY: init
init:
		echo $(NODE1IPFSCMD)
		#$(shell $(INITCMD))
		$(shell $(NODE1IPFSCMD))
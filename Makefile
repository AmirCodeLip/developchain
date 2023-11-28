CURRENTPATH :=$(shell pwd)
DATAPATH :=$(CURRENTPATH)/l1chain/data
IPFSNODE1PATH=$(CURRENTPATH)/ipfs/node1
GENESISPATH :=$(CURRENTPATH)/genesis.json
INITCMD :=docker run -v $(DATAPATH):/data -v $(GENESISPATH):/genesis.json
INITCMD +=ethereum/client-go init --datadir /data /genesis.json
INITIPFSCMD :=docker run -v $(IPFSNODE1PATH):/data
INITIPFSCMD+= ipfs/kubo init --repo-dir=/data --config-file=/data/config
ifneq ($(shell [ -d ${DATAPATH} ] && echo "true"),true)
$(shell mkdir -p $(DATAPATH))
$(shell cp l1chain/keystore l1chain/data/keystore)
endif
.PHONY: install
install:
		$(shell docker pull ipfs/kubo)
.PHONY: init
init:
		echo $(INITIPFSCMD)
		#$(shell $(INITCMD))
		$(shell $(INITIPFSCMD))
		#$(shell echo -e "/key/swarm/psk/1.0.0/\n/base16/\n`tr -dc 'a-f0-9' < /dev/urandom | head -c64`" > $(IPFSNODE1PATH))
		#$(shell docker run -v $(IPFSNODE1PATH):/data IPFS_PATH=/data ipfs/kubo bootstrap rm --all)

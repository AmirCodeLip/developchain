CURRENTPATH :=$(shell pwd)
DATAPATH :=$(CURRENTPATH)/l1chain/data
GENESISPATH := $(DATAPATH)/genesis.json

ifneq ($(shell [ -d ${DATAPATH} ] && echo "true"),true)
$(shell mkdir -p $(DATAPATH))
endif

.PHONY: init
init:
		$(shell cp genesis.json $(GENESISPATH))
		$(shell docker run ethereum/client-go init --datadir $(DATAPATH) genesis.json)

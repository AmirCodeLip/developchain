CURRENTPATH :=$(shell pwd)
DATAPATH :=$(CURRENTPATH)/l1chain/data
GENESISPATH :=$(CURRENTPATH)/genesis.json
INITCMD :=docker run -v $(DATAPATH):/data -v $(GENESISPATH):/genesis.json
INITCMD +=ethereum/client-go init --datadir /data /genesis.json

ifneq ($(shell [ -d ${DATAPATH} ] && echo "true"),true)
$(shell mkdir -p $(DATAPATH))
endif

.PHONY: init
init:
		$(shell $(INITCMD))


# Storage Providers

The contracts folder contains the important files required to interact with the lotus Node for the oracles and for the SP.




## Files available

Description of the important functions available in the smart contract

```bash
  dealFilter.pl - This file is used to validate if the deal is coming from one of the oracles or not. This file needs to updated in the lotus_config.toml file.
```

```bash
  index.js - This file, helps us to upload a file and generate a CID based on it for further processing
```

```bash
  getCID.js - This file interacts with the lotus lite node and generated the DataCID which is returned back to index.js
```

```bash
  storageFile.js - This file is used to create a deal between the oracle and the SP for the given DataCID.
```


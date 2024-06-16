
import { address, toNano } from "@ton/core";
import { MainContract } from "../wrappers/MainContract";
import { NetworkProvider, compile } from "@ton/blueprint";

export async function run(provider: NetworkProvider) {
    const myContract = MainContract.createFromConfig(
      {
        number: 0,
        address: address("0QDS9VGn1NrsHGzzBLjqumKzpV2d_0HsTvsathgAeUGlt8tx"),
        owner_address: address(
          "0QDS9VGn1NrsHGzzBLjqumKzpV2d_0HsTvsathgAeUGlt8tx"
        ),
      },
      await compile("MainContract")
    );
  
    const openedContract = provider.open(myContract);
  
    openedContract.sendDeploy(provider.sender(), toNano("1"));
  
    await provider.waitForDeploy(myContract.address, 6, 10000);
  }
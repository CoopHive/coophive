import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployToken: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, diamond } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  // log the admin address
  console.log(`admin: ${admin}`)
  await diamond.deploy("CoopHiveToken", {
    from: admin,
    owner: admin,
    facets: [
      {
        contract: "contracts/token/CoopHiveToken.sol:CoopHiveTokenFacet",
      },
    ],
    log: true,
  })
  return true
}

deployToken.id = 'deployToken'

export default deployToken
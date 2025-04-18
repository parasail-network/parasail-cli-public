package generate

var ContractFiles = map[string]string{
	"DepinHooks.sol": depinHooksContract,
}

const depinHooksContract = `// SPDX-License-Identifier: BUSL-1.1

interface IDePINHooks {
    function onSLAReported(address vault, uint256 dealID, bytes1 status) external;

    function onVaultSlashed(address vault, uint256 amount) external;

    function onPaymentSettled(address vault, uint256 dealID) external;

    function onPaymentTerminated(address vault, uint256 dealID) external;
}

contract DePINHooks is IDePINHooks {
    address public vaultManager;

    constructor(address _vaultManager) {
        vaultManager = _vaultManager;
    }

    /**
     * @notice Called when a SLA is reported.
     */

    function onSLAReported(address vault, uint256 dealID, bytes1 status) external onlyVaultManager {}

    /**
     * @notice Called when a vault is slashed.
     */
    function onVaultSlashed(address vault, uint256 amount) external onlyVaultManager {}

    /**
     * @notice Called when a deal is settled.
     */
    function onPaymentSettled(address vault, uint256 dealID) external onlyVaultManager {}

    /**
     * @notice Called when a deal is terminated.
     */
    function onPaymentTerminated(address vault, uint256 dealID) external onlyVaultManager {}

    modifier onlyVaultManager() {
        require(msg.sender == vaultManager, "DePINHooks: caller is not the vault manager");
        _;
    }
}`

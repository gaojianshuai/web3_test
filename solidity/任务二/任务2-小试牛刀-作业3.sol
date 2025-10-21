// 作业3：编写一个讨饭合约
// 任务目标
// 使用 Solidity 编写一个合约，允许用户向合约地址发送以太币。
// 记录每个捐赠者的地址和捐赠金额。
// 允许合约所有者提取所有捐赠的资金。

// 任务步骤
// 编写合约
// 创建一个名为 BeggingContract 的合约。
// 合约应包含以下功能：
// 一个 mapping 来记录每个捐赠者的捐赠金额。
// 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
// 一个 withdraw 函数，允许合约所有者提取所有资金。
// 一个 getDonation 函数，允许查询某个地址的捐赠金额。
// 使用 payable 修饰符和 address.transfer 实现支付和提款。
// 部署合约
// 在 Remix IDE 中编译合约。
// 部署合约到 Goerli 或 Sepolia 测试网。
// 测试合约
// 使用 MetaMask 向合约发送以太币，测试 donate 功能。
// 调用 withdraw 函数，测试合约所有者是否可以提取资金。
// 调用 getDonation 函数，查询某个地址的捐赠金额。

// 任务要求
// 合约代码：
// 使用 mapping 记录捐赠者的地址和金额。
// 使用 payable 修饰符实现 donate 和 withdraw 函数。
// 使用 onlyOwner 修饰符限制 withdraw 函数只能由合约所有者调用。
// 测试网部署：
// 合约必须部署到 Goerli 或 Sepolia 测试网。
// 功能测试：
// 确保 donate、withdraw 和 getDonation 函数正常工作
// 提交内容
// 提交完整的合约代码。 
// 提交部署合约的交易哈希。
// 提交测试过程的截图或视频，展示各功能的使用情况。
// 捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
// 捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
// 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。


// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BeggingContract {
    // 合约所有者
    address public owner;
    
    // 记录每个地址的捐赠总额
    mapping(address => uint256) public donations;
    
    // 捐赠者地址列表
    address[] public donors;
    
    // 捐赠事件
    event Donation(address indexed donor, uint256 amount, uint256 timestamp);
    
    // 提款事件
    event Withdrawal(address indexed owner, uint256 amount, uint256 timestamp);
    
    // 捐赠时间限制
    uint256 public donationStartTime;
    uint256 public donationEndTime;
    
    // 修饰器：仅所有者可调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
    
    // 修饰器：检查捐赠时间
    modifier inDonationPeriod() {
        require(block.timestamp >= donationStartTime && block.timestamp <= donationEndTime, 
                "Donations are only accepted during the specified period");
        _;
    }
    
    // 构造函数
    constructor(uint256 _donationDurationInHours) {
        owner = msg.sender;
        donationStartTime = block.timestamp;
        donationEndTime = donationStartTime + (_donationDurationInHours * 1 hours);
    }
    
    // 捐赠函数
    function donate() public payable inDonationPeriod {
        require(msg.value > 0, "Donation amount must be greater than 0");
        
        // 如果是第一次捐赠，添加到捐赠者列表
        if (donations[msg.sender] == 0) {
            donors.push(msg.sender);
        }
        
        // 更新捐赠金额
        donations[msg.sender] += msg.value;
        
        // 触发捐赠事件
        emit Donation(msg.sender, msg.value, block.timestamp);
    }
    
    // 提款函数（仅所有者）
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");
        
        payable(owner).transfer(balance);
        
        // 触发提款事件
        emit Withdrawal(owner, balance, block.timestamp);
    }
    
    // 查询捐赠金额
    function getDonation(address _donor) external view returns (uint256) {
        return donations[_donor];
    }
    
    // 获取捐赠排行榜（前3名）
    function getTopDonors() external view returns (address[3] memory topDonors, uint256[3] memory amounts) {
        // 初始化结果数组
        address[3] memory topAddresses;
        uint256[3] memory topAmounts;
        
        // 遍历所有捐赠者
        for (uint256 i = 0; i < donors.length; i++) {
            address donor = donors[i];
            uint256 amount = donations[donor];
            
            // 插入排序到前3名
            for (uint256 j = 0; j < 3; j++) {
                if (amount > topAmounts[j]) {
                    // 将当前捐赠者插入到位置j
                    for (uint256 k = 2; k > j; k--) {
                        topAddresses[k] = topAddresses[k-1];
                        topAmounts[k] = topAmounts[k-1];
                    }
                    topAddresses[j] = donor;
                    topAmounts[j] = amount;
                    break;
                }
            }
        }
        
        return (topAddresses, topAmounts);
    }
    
    // 获取捐赠者总数
    function getDonorCount() external view returns (uint256) {
        return donors.length;
    }
    
    // 获取合约余额
    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }
    
    // 获取捐赠时间段信息
    function getDonationPeriod() external view returns (uint256 start, uint256 end, bool isActive) {
        start = donationStartTime;
        end = donationEndTime;
        isActive = (block.timestamp >= start && block.timestamp <= end);
    }
    
    // 延长捐赠时间（仅所有者）
    function extendDonationPeriod(uint256 _additionalHours) external onlyOwner {
        donationEndTime += (_additionalHours * 1 hours);
    }
    
    // 接收以太币的回退函数
    receive() external payable inDonationPeriod {
        donate();
    }
}

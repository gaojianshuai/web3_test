/// @title A title that should describe the contract/interface
/// @author The name of the author
/// @notice Explain to an end user what this does
/// @dev Explain to a developer any extra details

pragma solidity ^0.8.0;


//整形加法函数
contract TestContract {
    /// @notice This function adds two numbers
    /// @dev This is a simple addition function
    /// @param a The first number
    /// @param b The second number
    /// @return sum The sum of the two numbers
    function add(uint256 a, uint256 b) public pure returns (uint256 sum) {
        sum = a + b;
    }
}

//布尔型函数
contract BoolContract {
    /// @notice This function checks if a number is even
    /// @dev This function uses the modulus operator to determine evenness
    /// @param num The number to check
    /// @return isEven True if the number is even, false otherwise
    function isEven(uint256 num) public pure returns (bool isEven) {
        isEven = (num % 2 == 0);
    }
}

//地址函数
contract AddressContract {
    /// @notice This function checks if an address is a contract
    /// @dev This function uses extcodesize to determine if the address is a contract
    /// @param addr The address to check
    /// @return isContract True if the address is a contract, false otherwise
    function isContract(address addr) public view returns (bool isContract) {
        uint32 size;
        assembly {
            size := extcodesize(addr)
        }
        isContract = (size > 0);
    }
}

//固定大小字节数组函数
contract FixedBytesContract {
    /// @notice This function converts a bytes32 to a string
    /// @dev This function iterates over the bytes32 and constructs a string
    /// @param data The bytes32 data to convert
    /// @return str The resulting string
    function bytes32ToString(bytes32 data) public pure returns (string memory str) {
        bytes memory bytesArray = new bytes(32);
        for (uint256 i = 0; i < 32; i++) {
            bytesArray[i] = data[i];
        }
        str = string(bytesArray);
    }
}

//枚举类型函数
contract EnumContract {
    /// @notice This enum represents different states
    enum State { Inactive, Active, Suspended }

    State public currentState;

    /// @notice This function sets the current state
    /// @dev This function updates the currentState variable
    /// @param newState The new state to set
    function setState(State newState) public {
        currentState = newState;
    }
}

//不需要指定 memory 等位置修饰符
contract NoMemoryContract {
    /// @notice This function returns the sum of two numbers without memory keyword
    /// @dev Solidity infers storage location for value types
    /// @param x The first number
    /// @param y The second number
    /// @return result The sum of the two numbers
    function sum(uint256 x, uint256 y) public pure returns (uint256 result) {
        result = x + y;
    }
}

//需要指定 memory 等位置修饰符
contract MemoryContract {
    /// @notice This function concatenates two strings
    /// @dev Strings are reference types and require memory keyword
    /// @param str1 The first string
    /// @param str2 The second string
    /// @return result The concatenated string
    function concatenate(string memory str1, string memory str2) public pure returns (string memory result) {
        result = string(abi.encodePacked(str1, str2));
    }
}

//大小限制
contract SizeLimitContract {
    uint8 public smallNumber; // 8 bits
    uint256 public largeNumber; // 256 bits

    /// @notice This function sets a small number
    /// @dev This function demonstrates the use of uint8
    /// @param num The small number to set
    function setSmallNumber(uint8 num) public {
        smallNumber = num;
    }

    /// @notice This function sets a large number
    /// @dev This function demonstrates the use of uint256
    /// @param num The large number to set
    function setLargeNumber(uint256 num) public {
        largeNumber = num;
    }
}

//结构体函数
contract StructContract {
    /// @notice This struct represents a person
    struct Person {
        string name;
        uint256 age;
    }

    Person public person;

    /// @notice This function sets the person's details
    /// @dev This function updates the person struct
    /// @param _name The name of the person
    /// @param _age The age of the person
    function setPerson(string memory _name, uint256 _age) public {
        person = Person({name: _name, age: _age});
    }
}
//数组函数
contract ArrayContract {
    uint256[] public numbers;

    /// @notice This function adds a number to the array
    /// @dev This function appends a number to the numbers array
    /// @param num The number to add
    function addNumber(uint256 num) public {
        numbers.push(num);
    }

    /// @notice This function gets the number at a specific index
    /// @dev This function retrieves a number from the numbers array
    /// @param index The index of the number to retrieve
    /// @return num The number at the specified index
    function getNumber(uint256 index) public view returns (uint256 num) {
        require(index < numbers.length, "Index out of bounds");
        num = numbers[index];
    }
}  

//字符串函数
contract StringContract {
    /// @notice This function returns the length of a string
    /// @dev This function uses bytes to determine the length of the string
    /// @param str The string to measure
    /// @return length The length of the string
    function getStringLength(string memory str) public pure returns (uint256 length) {
        length = bytes(str).length;
    }
}
//映射函数
contract MappingContract {
    mapping(address => uint256) public balances;

    /// @notice This function sets the balance for an address
    /// @dev This function updates the balances mapping
    /// @param user The address of the user
    /// @param amount The balance amount to set
    function setBalance(address user, uint256 amount) public {
        balances[user] = amount;
    }

    /// @notice This function gets the balance for an address
    /// @dev This function retrieves the balance from the balances mapping
    /// @param user The address of the user
    /// @return balance The balance of the user
    function getBalance(address user) public view returns (uint256 balance) {
        balance = balances[user];
    }
}

/*
创建一个名为Voting的合约，包含以下功能：
一个mapping来存储候选人的得票数
一个vote函数，允许用户投票给某个候选人
一个getVotes函数，返回某个候选人的得票数
一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    mapping(string => uint256) public votes;

    /// @notice This function allows a user to vote for a candidate
    /// @dev This function increments the vote count for the specified candidate
    /// @param candidate The name of the candidate to vote for
    function vote(string memory candidate) public {
        votes[candidate] += 1;
    }

    /// @notice This function returns the number of votes for a candidate
    /// @dev This function retrieves the vote count from the votes mapping
    /// @param candidate The name of the candidate
    /// @return voteCount The number of votes for the candidate
    function getVotes(string memory candidate) public view returns (uint256 voteCount) {
        voteCount = votes[candidate];
    }

    /// @notice This function resets all candidates' votes
    /// @dev This function iterates through known candidates and resets their vote counts
    /// Note: In a real implementation, you would need to keep track of candidates separately
    function resetVotes(string[] memory candidates) public {
        for (uint256 i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
    }
}

//题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
contract ReverseString {
    /// @notice This function reverses a given string
    /// @dev This function converts the string to bytes, reverses it, and converts it back to string
    /// @param str The string to reverse
    /// @return reversedStr The reversed string
    function reverse(string memory str) public pure returns (string memory reversedStr) {
        bytes memory strBytes = bytes(str);
        uint256 len = strBytes.length;
        bytes memory reversedBytes = new bytes(len);
        for (uint256 i = 0; i < len; i++) {
            reversedBytes[i] = strBytes[len - 1 - i];
        }
        reversedStr = string(reversedBytes);
    }
}

//减少gas消耗的合约
contract GasOptimized {
    uint256 public total;

    /// @notice This function adds a number to the total
    /// @dev This function uses unchecked to save gas on overflow checks
    /// @param num The number to add
    function add(uint256 num) public {
        unchecked {
            total += num;
        }
    }
}

//用 solidity 实现整数转罗马数字的函数
contract IntToRoman {
    /// @notice This function converts an integer to a Roman numeral
    /// @dev This function uses arrays to map integer values to Roman numeral symbols
    /// @param num The integer to convert
    /// @return roman The resulting Roman numeral as a string
    function intToRoman(uint256 num) public pure returns (string memory roman) {
        uint16[13] memory values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        
        bytes memory result;
        for (uint256 i = 0; i < values.length; i++) {
            while (num >= values[i]) {
                num -= values[i];
                result = abi.encodePacked(result, symbols[i]);
            }
        }
        roman = string(result);
    }
}
//用 solidity 实现罗马数字转数整数的函数
contract RomanToInt {
    /// @notice This function converts a Roman numeral to an integer
    /// @dev This function uses a mapping to convert Roman numeral symbols to integer values
    /// @param roman The Roman numeral string to convert
    /// @return num The resulting integer
    function romanToInt(string memory roman) public pure returns (uint256 num) {
        bytes memory romanBytes = bytes(roman);
        uint256 length = romanBytes.length;
        bytes1[7] memory romanSymbols = [bytes1("I"), bytes1("V"), bytes1("X"), bytes1("L"), bytes1("C"), bytes1("D"), bytes1("M")];
        uint256[7] memory romanValues = [1, 5, 10, 50, 100, 500, 1000];

        for (uint256 i = 0; i < length; i++) {
            uint256 currentValue = values[romanBytes[i]];
            uint256 nextValue = (i + 1 < length) ? values[romanBytes[i + 1]] : 0;

            if (currentValue < nextValue) {
                num -= currentValue;
            } else {
            uint256 currentValue = 0;
            uint256 nextValue = 0;

            for (uint256 j = 0; j < romanSymbols.length; j++) {
                if (romanBytes[i] == romanSymbols[j]) {
                    currentValue = romanValues[j];
                }
                if (i + 1 < length && romanBytes[i + 1] == romanSymbols[j]) {
                    nextValue = romanValues[j];
                }
            }
        }
    }
}

}
//将两个有序数组合并为一个有序数组
contract MergeSortedArrays {
    /// @notice This function merges two sorted arrays into one sorted array
    /// @dev This function uses a two-pointer technique to merge the arrays
    /// @param arr1 The first sorted array
    /// @param arr2 The second sorted array
    /// @return mergedArray The resulting merged sorted array
    function merge(uint256[] memory arr1, uint256[] memory arr2) public pure returns (uint256[] memory mergedArray) {
        uint256 len1 = arr1.length;
        uint256 len2 = arr2.length;
        mergedArray = new uint256[](len1 + len2);
        
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;

        while (i < len1 && j < len2) {
            if (arr1[i] <= arr2[j]) {
                mergedArray[k++] = arr1[i++];
            } else {
                mergedArray[k++] = arr2[j++];
            }
        }

        while (i < len1) {
            mergedArray[k++] = arr1[i++];
        }

        while (j < len2) {
            mergedArray[k++] = arr2[j++];
        }
    }
}

//二分查找：在一个有序数组中查找目标值
contract BinarySearch {
    /// @notice This function performs binary search on a sorted array
    /// @dev This function returns the index of the target value or -1 if not found
    /// @param arr The sorted array to search
    /// @param target The target value to find
    /// @return index The index of the target value, or -1 if not found
    function binarySearch(uint256[] memory arr, uint256 target) public pure returns (int256 index) {
        int256 left = 0;
        int256 right = int256(arr.length) - 1;

        while (left <= right) {
            int256 mid = left + (right - left) / 2;
            if (arr[uint256(mid)] == target) {
                return mid;
            } else if (arr[uint256(mid)] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return -1; // Target not found
    }
}
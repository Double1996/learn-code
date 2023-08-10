
// var solc = require('solc');

// var contract = 
// `// SPDX-License-Identifier: MIT
// pragma solidity ^0.8.10;
// contract HelloWeb3{
//     string public _string = "Hello Web3!";
// }`

// var input = {
//   language: 'Solidity',
//   sources: {
//     'HelloWeb3.sol': {
//       content: contract
//     }
//   },
//   settings: {
//     outputSelection: {
//       '*': {
//         '*': ['metadata'],
//       },
//     },
//   }
// };


// const output = JSON.parse(solc.compile(JSON.stringify(input)));

// const metadata = output.contracts['MyContract.sol']['metadata'];

// console.log(metadata);


var solc = require('solc');
var contract = `
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
contract HelloWeb3 {
    string public _string = "Hello Web3!";
}
`;
var input = {
  language: 'Solidity',
  sources: {
    'HelloWeb3.sol': {
      content: contract
    }
  },
  settings: {
    outputSelection: {
      '*': {
        '*': ['metadata'],
      },
    },
  }
};
const output = JSON.parse(solc.compile(JSON.stringify(input)));
const metadata = output.contracts['HelloWeb3.sol']['HelloWeb3'].metadata;
console.log(metadata);
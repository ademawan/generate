var Crypto = require('crypto');
var standard_input = process.stdin;
//npm install --save-dev nodemon
//npm install crypto-js
standard_input.setEncoding('utf-8');

console.log("Please input encrypted link then press enter:");
//3b4599332d162fd95a599941166249a9-11515fd0a1b8824ba58581c00637baeffc048ce3014e8e001a173259ee40ca145e72eadf35ea0d6a8da7a9a623098b97490351a5883b9419f2be11facf33c6d60ab5818cdada71f1ca2ff3ea2dab0774z
standard_input.on('data', function (data) {
  const decrypted = decryptCustParams(data.trim());
  console.log('Decrypted Cust Params: ' + decrypted);
  standard_input.setEncoding('utf-8');
  console.log("Please input encrypted link then press enter:");
});

const decryptCustParams = (custParam) => {
  const cipherAlgorithm = 'aes-256-cbc';
  // const cipherPassword = 'GDJVesGYZJNEVUU7UhgGEhwBa8fv5nmk'; //Prod
  const cipherPassword = 'GDJVesGYZJNEVUU7UhgGEhwBa8fv5nmk'; //Preprod
  const date =new Date().toString()
  console.log(date)
  const textParts = custParam.split('-');
  // console.log("ddd",textParts.shift(),textParts.join(':'))
  
  const iv = Buffer.from(textParts.shift(), 'hex');
  console.log(iv)
  console.log(textParts.join(':'))
  var cc= Buffer.from(textParts.join(':'), 'hex')
  var t = textParts.join(':')
  console.log(t)
  console.log(cc)

  const encryptedText = Buffer.from(textParts.join(':'), 'hex');
  const decipher = Crypto.createDecipheriv(
    cipherAlgorithm,
    Buffer.from(cipherPassword),
    iv,
  );
  let decrypted = decipher.update(encryptedText);
  // console.log("ddddd",decrypted.toString(),decipher.final().toString());

  decrypted = Buffer.concat([decrypted, decipher.final()]);

  return decrypted.toString();
};
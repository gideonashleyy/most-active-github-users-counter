const fs = require('fs');

// Read and parse JSON data from file
let rawdata = fs.readFileSync('./data.json');
let data = JSON.parse(rawdata);

// Sort users by contributions in descending order and assign contribution rank
data.users.sort((a, b) => b.contributions - a.contributions);
data.users.forEach((user, index) => {
    user.contributionRank = index + 1;
});

// Sort users by follower rank in ascending order
data.users.sort((a, b) => a.followerRank - b.followerRank);

// Get current date and format as DDMMYYYY
let today = new Date();
let dd = String(today.getDate()).padStart(2, '0');
let mm = String(today.getMonth() + 1).padStart(2, '0'); 
let yyyy = today.getFullYear();
let date = dd + mm + yyyy;

// Write modified data back to new JSON file with date in filename
let newData = JSON.stringify(data, null, 2);
let filename = date + '-data.json';
fs.writeFileSync(filename, newData);

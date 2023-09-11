package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "sort"
    "time"
)

// User struct to represent a user in the JSON data
type User struct {
    Username         string `json:"username"`
    Contributions    int    `json:"contributions"`
    FollowerRank     int    `json:"followerRank"`
    ContributionRank string `json:"contributionRank"`
}

// ContributionRank is a custom type for ContributionRank field
type ContributionRank int

func (c *ContributionRank) UnmarshalJSON(data []byte) error {
    var rank int
    if err := json.Unmarshal(data, &rank); err != nil {
        return err
    }
    *c = ContributionRank(rank)
    return nil
}

// Data struct to represent the JSON data
type Data struct {
    Users []User `json:"users"`
}

func main() {
    // Read and parse JSON data from file
    rawdata, err := ioutil.ReadFile("data.json")
    if err != nil {
        fmt.Println("Error reading data.json:", err)
        return
    }

    var data Data
    if err := json.Unmarshal(rawdata, &data); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    // Sort users by contributions in descending order and assign contribution rank
    sort.Slice(data.Users, func(i, j int) bool {
        return data.Users[i].Contributions > data.Users[j].Contributions
    })

    for index, user := range data.Users {
        user.ContributionRank = ContributionRank(index + 1)
        data.Users[index] = user
    }

    // Sort users by follower rank in ascending order
    sort.Slice(data.Users, func(i, j int) bool {
        return data.Users[i].FollowerRank < data.Users[j].FollowerRank
    })

    // Get current date and format as DDMMYYYY
    today := time.Now()
    date := today.Format("02012006")

    // Write modified data back to new JSON file with date in filename
    newData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }

    filename := date + "-data.json"
    err = ioutil.WriteFile(filename, newData, 0644)
    if err != nil {
        fmt.Println("Error writing JSON to file:", err)
        return
    }

    fmt.Println("Data written to", filename)
}

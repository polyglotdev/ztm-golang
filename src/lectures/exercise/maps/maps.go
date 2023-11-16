//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func onlineServersCount(servers map[string]int) int {
	count := 0
	for _, status := range servers {
		if status == Online {
			count++
		}
	}
	return count
}

func offlineServersCount(servers map[string]int) int {
	count := 0
	for _, status := range servers {
		if status == Offline {
			count++
		}
	}
	return count
}

func maintenanceServersCount(servers map[string]int) int {
	count := 0
	for _, status := range servers {
		if status == Maintenance {
			count++
		}
	}
	return count
}

func retiredServersCount(servers map[string]int) int {
	count := 0
	for _, status := range servers {
		if status == Retired {
			count++
		}
	}
	return count
}

func displayServerInfo(servers map[string]int) {
	fmt.Println("Number of servers:", len(servers))
	fmt.Println("Online servers:", onlineServersCount(servers))
	fmt.Println("Offline servers:", offlineServersCount(servers))
	fmt.Println("Maintenance servers:", maintenanceServersCount(servers))
	fmt.Println("Retired servers:", retiredServersCount(servers))
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatuses := make(map[string]int)
	for _, server := range servers {
		serverStatuses[server] = Online
	}
	displayServerInfo(serverStatuses)
	serverStatuses["darkstar"] = Retired
	serverStatuses["aiur"] = Offline
	displayServerInfo(serverStatuses)
	for server := range serverStatuses {
		serverStatuses[server] = Maintenance
	}
	displayServerInfo(serverStatuses)
}

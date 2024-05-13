package main

/**************************************************************************************************
** The version of the daemon is displayed on the server. The version is based on the latest git
** commit, trimmed to only show the first 7 characters of the commit hash.
**************************************************************************************************/
var version = ""

func GetVersion() string {
	if version == "" {
		return "dev"
	}
	return version[:7]
}

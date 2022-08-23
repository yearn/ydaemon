import os
import eth_utils
import json


def checksum_encode(addr):
    # from https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md
    hex_addr = addr.hex()
    checksummed_buffer = ""
    hashed_address = eth_utils.keccak(text=hex_addr).hex()
    for nibble_index, character in enumerate(hex_addr):
        if character in "0123456789":
            checksummed_buffer += character
        elif character in "abcdef":
            hashed_address_nibble = int(hashed_address[nibble_index], 16)
            if hashed_address_nibble > 7:
                checksummed_buffer += character.upper()
            else:
                checksummed_buffer += character
    return "0x" + checksummed_buffer


def test(addr_str):
    addr_bytes = eth_utils.to_bytes(hexstr=addr_str)
    checksum_encoded = checksum_encode(addr_bytes)
    return checksum_encoded == addr_str, addr_str, checksum_encoded


BASE = "data"
FOLDERS = [
    "vaults",  # check filenames
    "tokens",  # check filenames
    "strategies",  # check values inside files
]


def check():
    checks = []
    for folder in FOLDERS:
        for root, _, files in os.walk(f"{BASE}/{folder}"):
            for file in files:
                if file.split(".")[-1] == "json":
                    filepath = os.path.join(root, file)
                    if folder != "strategies":
                        name = file.split(".")[0]
                        if name.startswith("0x") and len(name) == 42:  # if address
                            correct, got, expected = test(name)
                            if not correct:
                                checks.append((got, expected, filepath))
                    else:
                        data = json.load(open(filepath))
                        for addr in data["addresses"]:
                            correct, got, expected = test(addr)
                            if not correct:
                                checks.append((got, expected, filepath))

    if len(checks) > 0:
        for got, expected, filepath in checks:
            print(f"Error: got {got}, expected {expected}; at {filepath}.")
        exit(1)


def main():
    check()


if __name__ == "__main__":
    main()

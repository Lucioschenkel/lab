import hashlib
import requests
import sys

def request_api_data(query_char):
    url = f'https://api.pwnedpasswords.com/range/{query_char}'
    res = requests.get(url)
    if res.status_code != 200:
        raise RuntimeError(f'Error fetching: {res.status_code}')

    return res

def get_password_leaks_count(hashes, hash_to_check) -> int:
    hashes = (line.split(':') for line in hashes.text.splitlines())
    for h, count in hashes:
        if h == hash_to_check:
            return int(count)

    return 0

def pwnd_api_check(password):
    sha1password = hashlib.sha1(password.encode('utf-8')).hexdigest()
    first5_char, tail = sha1password[:5], sha1password[5:]
    response = request_api_data(first5_char)
    return get_password_leaks_count(response, tail.upper())

def main(args):
    for password in args:
        count = pwnd_api_check(password)
        if count > 0:
            print(f'{password} was found {count} times... you should probably change your password.')
        else:
            print(f'{password} was not found. Carry on!')

if __name__ == "__main__":
    main(sys.argv[1:])

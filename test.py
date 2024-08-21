import requests

def send_request_to_server():
    url = 'http://localhost:8080/api'

    try:
        response = requests.post(url, data={'key': 'value'})

        print(f"Status Code: {response.status_code}")

        print("Headers:")
        for header, value in response.headers.items():
            print(f"{header}: {value}")

        print("\nContent:")
        print(response.text)

    except requests.RequestException as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    send_request_to_server()


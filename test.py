import requests

def send_request_to_server(route):
    url = f'http://localhost:8080/{route}'

    try:
        response = requests.get(url)

        print(f"\nContent: {response.text}")

    except requests.RequestException as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    send_request_to_server('api/users/123')
    send_request_to_server('api/unknown')



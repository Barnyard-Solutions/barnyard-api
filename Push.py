from pywebpush import webpush, WebPushException
import json

PRIVATE_KEY = 'CZtf_JUxmXkCKbzwaKedPPO9BFC99U2rk-GUYDbYAa8'

def send_web_push(subscription_information, message_body):
    message_body = json.dumps(message_body)
    return webpush(
        subscription_info=subscription_information,
        data=message_body,
        vapid_private_key=PRIVATE_KEY,
        vapid_claims= {"sub": "mailto:develop@raturi.in"},
        ttl=60*60*72
    )








if __name__ == '__main__':
    sub_info = {
    'endpoint': 'https://fcm.googleapis.com/fcm/send/f_PsZXPWnhE:APA91bHo5y4XIxPZqmYMqGwXznHdFiR14PsqecR6jSMZClW6q2_en5jK2kGBhQTaGtg7w6oUH9ua5F5-ynfCg7wNf-UQORGphhwdPGoXIhSKjH5n1HawmZKDq0MVsc1XF7mm_7iRcQnq',
    'expirationTime': None,
    'keys': {'p256dh': 'BPvayYK8h62mmllWopLVTqxA4XaPwvAITxp5teiu6ecZuB7F4MJDdqNfbpWOE3ag2wB3ZnD9Xe6KCGUCnq065kQ', 'auth': 'orwucmaF4uepp1lkiwN66Q'}
}



    send_web_push(sub_info, { "title": "hello3", "body": "BODY2" })
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
    'endpoint': 'https://fcm.googleapis.com/fcm/send/fs2hXM1Wwlk:APA91bFrbKWt3DGOjsZXRqJvorI6g3lmGpuXQlo_yhF4Zn2-sMIddVeukTSiP7gT19UVpggeYmknue7vQKdqE_MjBqIH8gNTah3NVUb3K6Z6voAQdjznVvFpwpHjLNgkhPCzgcUOpaCD',
    'expirationTime': None,
    'keys': {'p256dh': 'BDw53IaZvj1-tCG2yFrBqpWEpduOvchYu3HkCeHt4pJkrSYhIxFvChM9jeu9kI26MvVog_-jqPf4XfBgebELtho', 'auth': 'Q8gRyARlTRB66zpppNAlVw'}
}



    send_web_push(sub_info, { "title": "hello3", "body": "BODY2" })
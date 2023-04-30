import mysql.connector
import hashlib
import time
from rich import print
import os

db_host = os.getenv('DB_HOST')

print("db_host", db_host)


def log(s):
    if debug:
        print("[bold yellow][DEBUG][Account][/bold yellow]",s)
debug = False



class MyBD:
  def __init__(self):
    self.PASS ="kKQdH@qX93"
    
    sucess = False
    while (not sucess ):
      try:
        self.start()
        sucess = True
      except Exception as e:
        # to be perficetioned
        #print(Exception)
        print("wating for bd to start up ....")
        time.sleep(10)
      
    log("connection sucessfull")
    
  def start(self):
    self.mydb = mysql.connector.connect(
    host = db_host,
    user = "API",
    password = self.PASS,
    buffered=True
    )
  def cursor(self):
    try:
      return self.mydb.cursor()
    except Exception as e:
      print("[error] connection restarting",e)
      self.start()
      return self.mydb.cursor()
  def commit(self):
    self.mydb.commit()


mydb = MyBD()

def MyHash(te,sha512=None):
  encoded=te.encode()
  if sha512 == None:
  
    result = hashlib.sha256(encoded)
    return result.hexdigest()
  else:
    result = hashlib.sha512(encoded)
    return result.hexdigest()


def CreateAccount(mail,pas):
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  comm = f"INSERT INTO account  VALUES (NULL, '{mail}', '{MyHash(pas)}') "
  try:
    cursor.execute(comm)
    return True
  except Exception as e:
    print(e)
    return False
  mydb.commit()

  cursor.close()



def Login(mail,pas):
  
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  comm = f"SELECT * FROM account WHERE USER_MAIL = '{mail}'"
  print(comm, flush=True)
  cursor.execute(comm)
  res = cursor.fetchall()

  if len(res)!= 0:
    
    res = res[0]

    print(pas)
    print(MyHash(pas),res[2])

    if MyHash(pas) == res[2]:
      cursor.close()
      return True
  cursor.close()
  return False

def GenerateKey(mail,pas):
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  comm = f"SELECT * FROM account WHERE USER_MAIL = '{mail}'"
  
  cursor.execute(comm)
  res = cursor.fetchall()

  if len(res)!= 0:
    
    res = res[0]

    
    print()
    if MyHash(pas) == res[2]:
      key = MyHash(str(time.time())+'barnyard',True)
      comm = f"INSERT INTO user_key  VALUES ( '{res[0]}', '{key}') "
      print(comm, flush=True)
      cursor.execute(comm)
      mydb.commit()
      cursor.close()
      return key
    else:
      print("[error] wrong pass")
  else:
    print("[error] mail wrong")
  cursor.close()
  return False

def CreateFeed(key,name):
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")

  comm = f"SELECT * FROM user_key WHERE USER_KEY = '{key}'"
  
  cursor.execute(comm)
  res = cursor.fetchall()
  print("[info] CreateFeed user_key check",res)
  if len(res)!= 0:
    
    res = res[0]

    comm = f"INSERT INTO feed  VALUES (NULL, '{name}', '{res[0]}') "
    print("[info]",comm)
    try:
      cursor.execute(comm)
      addFeedViewer(key,name,getUserId(key))
      print("[info] feed added")
    except Exception as e:
      print("[error] name feed already existe ",e)
    mydb.commit()
    cursor.close()
    

def addFeedViewer(key,feed_name,user_id):
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")

  comm = f"SELECT * FROM user_key WHERE USER_KEY = '{key}'"
  
  cursor.execute(comm)
  res = cursor.fetchall()
  if len(res)!= 0:
    res = res[0]
    print("checking permission..",user_id,res[0])

    comm = f"SELECT * FROM feed WHERE FEED_NAME = '{feed_name}' AND OWNER_ID = '{res[0]}' "
  
    cursor.execute(comm)
    res2 = cursor.fetchall()


    if len(res2)!= 0:
      res2 = res2[0]
      comm = f"INSERT INTO view_feed  VALUES ( '{res2[0]}', '{user_id}') "
      cursor.execute(comm)
      mydb.commit()
      cursor.close()
    else:
      print("[error] feed not found",feed_id,res)
  else:
    print("[error] check checking owner",user_id,res[0])

def addFeedViewerById(key,feed_id,user_id):
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")

  comm = f"SELECT * FROM user_key WHERE USER_KEY = '{key}'"
  
  cursor.execute(comm)
  res = cursor.fetchall()
  if len(res)!= 0:
    res = res[0]
    print("checking permission..",user_id,res[0])

    comm = f"SELECT * FROM feed WHERE FEED_ID = '{feed_id}' AND OWNER_ID = '{res[0]}' "
  
    cursor.execute(comm)
    res2 = cursor.fetchall()


    if len(res2)!= 0:
      res2 = res2[0]
      comm = f"INSERT INTO view_feed  VALUES ( '{res2[0]}', '{user_id}') "
      cursor.execute(comm)
      mydb.commit()
      cursor.close()
    else:
      print("[error] feed not found",feed_id,res)
  else:
    print("[error] check checking owner",user_id,res[0])

def getUserId(key):
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")

  comm = f"SELECT * FROM user_key WHERE USER_KEY = '{key}'"
  
  cursor.execute(comm)
  res = cursor.fetchall()
  log(res)
  if len(res)!= 0:
    res = res[0]
    return res[0]

def getFeedId(key,namefeed):
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")

  comm = f"SELECT feed.FEED_ID FROM feed, view_feed WHERE feed.FEED_ID=view_feed.FEED_ID AND feed.FEED_NAME = '{namefeed}' AND view_feed.USER_ID = '{getUserId(key)}' "
  #comm = f"SELECT feed.FEED_ID, feed.FEED_NAME FROM feed, view_feed WHERE feed.FEED_ID=view_feed.FEED_ID "
  cursor.execute(comm)
  res2 = cursor.fetchall()

  if len(res2)!= 0:
    res2 = res2[0]
    return res2[0]


def ifUserViewer(key,feed_id):
  user_id = getUserId(key)

  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")

  comm = f"SELECT * FROM view_feed WHERE USER_ID = '{user_id}' AND FEED_ID = '{feed_id}' "
  
 
  cursor.execute(comm)
  res2 = cursor.fetchall()

  if len(res2)!= 0:
    cursor.close()
    return True
  cursor.close()
  return False


  pass
def addEvent(key,name1,name2,date,feed_id):
  user_id = getUserId(key)
  if ifUserViewer(key,feed_id):
    cursor = mydb.cursor()
    cursor.execute("USE barnyard;")
    comm = f"INSERT INTO event  VALUES ( NULL,'{name1}', '{name2}', '{date}', '{feed_id}') "
    print(comm)
    cursor.execute(comm)
    mydb.commit()
    cursor.close()
    return True
  print("[error] user not viewer addEvent")
  return False
  
def getEvent(key,feed_id):
  user_id = getUserId(key)
  if ifUserViewer(key,feed_id):
    cursor = mydb.cursor()
    cursor.execute("USE barnyard;")
    comm = f"SELECT * FROM event WHERE event.FEED_ID = '{feed_id}' ORDER BY date;"
  

    cursor.execute(comm)
    res2 = cursor.fetchall()
    return res2
def removeEvent(key,feed_id,event_id):
  user_id = getUserId(key)
  if ifUserViewer(key,feed_id):
    cursor = mydb.cursor()
    cursor.execute("USE barnyard;")
    comm = f"DELETE FROM event WHERE event.EVENT_ID = '{event_id}' "
  

    cursor.execute(comm)
    mydb.commit()
    cursor.close()
    
    return True

def getFeed(key):
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  user_id = getUserId(key)


  log(user_id)

  comm = f"SELECT FEED_ID from view_feed WHERE USER_ID = '{user_id}' "
  
  cursor.execute(comm)
  log(cursor.fetchall())


  comm = f"SELECT * FROM feed WHERE  FEED_ID IN (SELECT FEED_ID from view_feed WHERE USER_ID = '{user_id}') "
  
  cursor.execute(comm)
  res = cursor.fetchall()
  return res


def removeFeed(key,feed_id):
  user_id = getUserId(key)
  if ifUserViewer(key,feed_id):
    cursor = mydb.cursor()
    cursor.execute("USE barnyard;")
    comm = f"SELECT OWNER_ID FROM feed WHERE  FEED_ID  = '{feed_id}' "
    cursor.execute(comm)
    res = cursor.fetchall()

    if len(res) !=0:
      if res[0][0] == user_id:
        print("owner")
        comm = f"SELECT USER_ID FROM view_feed WHERE  FEED_ID  = '{feed_id}' AND USER_ID <> '{user_id}' "
      
        cursor.execute(comm)
        res = cursor.fetchall()   
        if len(res) != 0:
          comm = f"UPDATE  feed SET OWNER_ID = '{res[0][0]}' WHERE FEED_ID = '{feed_id}' "
          cursor.execute(comm)
          comm = f"DELETE FROM  view_feed WHERE FEED_ID = '{feed_id}' AND USER_ID = '{user_id}' "
          cursor.execute(comm)
          mydb.commit()
          return True
         
        else:
          comm = f"DELETE FROM  feed WHERE FEED_ID = '{feed_id}' "
          cursor.execute(comm)

          comm = f"DELETE FROM event WHERE  FEED_ID  = '{feed_id}' "
          cursor.execute(comm)
        
          comm = f"DELETE FROM  view_feed WHERE FEED_ID = '{feed_id}' "
          cursor.execute(comm)
          mydb.commit()
          return True

      else:
        comm = f"DELETE FROM  view_feed WHERE FEED_ID = '{feed_id}' AND USER_ID = '{user_id}' "
        cursor.execute(comm)
        mydb.commit()
        return True

    else:
      return False
 

def addClass(key,feed_id,name,jour):
 
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
 
  comm = f"select add_class('{key}',{feed_id},'{name}',{jour}) from DUAL;"
  cursor.execute(comm)
  res = cursor.fetchall()
  mydb.commit()
  cursor.close()
  if not res:
    print("[error] count not add class")
  return res
  
 

def getClass(key,feed_id):
  user_id = getUserId(key)
  if ifUserViewer(key,feed_id):
    cursor = mydb.cursor()
    cursor.execute("USE barnyard;")
    comm = f"SELECT * FROM class WHERE  FEED_ID = '{feed_id}'  ORDER BY CLASS_DAY "
    
    cursor.execute(comm)
    res = cursor.fetchall()
    return res
  print("[error] user not viewer getClass")
  return False

def removeClass(key,feed_id,class_id):

  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  comm = f"select remove_class('{key}',{feed_id},{class_id}) from DUAL;"
  cursor.execute(comm)
  res = cursor.fetchall()
  mydb.commit()
  cursor.close()
  
  if not res:
    print("[error] count not remove class")
  return res

def addEndPoint(key,feed_id,end_point):
  user_id = getUserId(key)
  
  if ifUserViewer(key,feed_id) and not isSub(key,feed_id,end_point):
    cursor = mydb.cursor()
    cursor.execute("USE barnyard;")
    
    comm = f"INSERT INTO subscription  VALUES ( NULL,'{key}', '{feed_id}', '{end_point['endpoint']}', '{list(end_point['keys'])[0]}', '{end_point['keys'][list(end_point['keys'])[0]]}', '{end_point['keys']['auth']}',  '{end_point['expirationTime']}') "
 

    print(comm,flush=True)
    cursor.execute(comm)
    mydb.commit()
    cursor.close()
    return True
  print("[error] user not viewer addEndPoint or already existe")
  return False


def removeEndPoint(key,feed_id,end_point):
  user_id = getUserId(key)
  if ifUserViewer(key,feed_id):
    cursor = mydb.cursor()
    cursor.execute("USE barnyard;")

    comm = f"DELETE FROM  subscription WHERE  FEED_ID = '{feed_id}' AND USER_KEY = '{key}' AND END_POINT_AUTH='{end_point['keys']['auth']}';  "
    print(comm,flush=True)
    
    cursor.execute(comm)
    mydb.commit()
    cursor.close()
    
    return True
  print("[error] user not viewer removeEndPoint")
  return False


def getEndPoint():
  
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  comm = f"SELECT * FROM subscription  "
    
  cursor.execute(comm)
  res = cursor.fetchall()
  return res
def getBetterEndPoint(i):
  
  tmp = {
    'endpoint': f'{i[3]}',
    'expirationTime': f'{i[7]}',
    'keys': {
        f'{i[4]}': f'{i[5]}',
        'auth': f'{i[6]}'
    }
    }
  return tmp


def isSub(key,feed_id,end_point):
  user_id = getUserId(key)
  cursor = mydb.cursor()

  cursor.execute("USE barnyard;")
  comm = f"SELECT * FROM subscription WHERE USER_KEY IN (SELECT USER_KEY FROM user_key WHERE USER_ID ='{user_id}') AND FEED_ID = '{feed_id}' AND END_POINT_AUTH='{end_point['keys']['auth']}'; "
  log(comm)
  cursor.execute(comm)
  res = cursor.fetchall()
  if len(res) != 0:
    return True
  return False


def test():
  key  = "a7fcbd8c5fb8de43a75df2faba3394082c8f4b83434ef7d2aba985f5c171590fd9fddcbb484634132513427c8698da4acaf9f2c72d03537ab3938c0561c79068"
  key2 = "ba3d2caad99701402da38b5382663d2b4cc8e7a5838508fb0ecae85f56376a78f9cc7f5a97d26b7e747873d1200a7f8616d42fbef38f75568bab68a7e5a21443"
  key3 ="0498595a3bfcd3b880602fadeee6630ce6c7c781b380ad1c3e1879942983140d5a485c99ae47cdf909b7eeca427f1a24b77dd7083103d469a2b6046de7401825"
 
  
  print(getUserId(key),getUserId(key2),getUserId(key3))

  test = "0498595a3bfcd3b880602"

  CreateFeed(key,test)

  feeds = getFeed(key)
  print("feeds",feeds)
 

  addEvent(key,"name1","name2","date",feeds[0][0])
  addEvent(key,"name1","name2","date",feeds[0][0])
  addEvent(key,"name1","name2","date",feeds[0][0])

  removeFeed(key,feeds[0][0])


  CreateFeed(key,test)

  feeds = getFeed(key)
  print("feeds",feeds)
 

  addEvent(key,"name1","name2","date",feeds[0][0])
  addEvent(key,"name1","name2","date",feeds[0][0])
  addEvent(key,"name1","name2","date",feeds[0][0])

  addFeedViewerById(key,feeds[0][0],getUserId(key2))

  removeFeed(key,feeds[0][0])
  feeds = getFeed(key2)

  addEvent(key2,"name1","name2","date",feeds[0][0])
  
  
  removeFeed(key2,feeds[0][0])



  CreateFeed(key,test)

  feeds = getFeed(key)
  print("feeds",feeds)
 

  addEvent(key,"name1","name2","date",feeds[0][0])
  addEvent(key,"name1","name2","date",feeds[0][0])
  addEvent(key,"name1","name2","date",feeds[0][0])

  addFeedViewerById(key,feeds[0][0],getUserId(key2))

  removeFeed(key,feeds[0][0])
  feeds = getFeed(key2)
  for i in getEvent(key2,feeds[0][0]):
    removeEvent(key2,feeds[0][0],i[0])

  addEvent(key2,"new","name2","date",feeds[0][0])

  addFeedViewerById(key2,feeds[0][0],getUserId(key))
  
  
  removeFeed(key2,feeds[0][0])

  removeFeed(key,feeds[0][0])

def getUserList():
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  comm = f"SELECT USER_MAIL FROM account  "
    
  cursor.execute(comm)
  res = cursor.fetchall()
  return res

if __name__ == "__main__":
  """
  key  = "a7fcbd8c5fb8de43a75df2faba3394082c8f4b83434ef7d2aba985f5c171590fd9fddcbb484634132513427c8698da4acaf9f2c72d03537ab3938c0561c79068"

  #CreateFeed(key,"test")

  feeds = getFeed(key)

  color = "#f6b73c"
  """
  end_point ="""{
    "endpoint": "https://fcm.googleapis.com/fcm/send/fKMviXEA5Wo:APA91bF6DbSxYLXs6WauEVhj6WO-B6hEQbBywAq8OsZiGSHyFiBxJ4h1iGJ-PexX3_AareYONHJH7k5l_Nju_dbj6fr8G-RbR_bZJng-vBLCRxXFnjSlcTQzoJTdf7Oywz7SDFt74bHQ",
    "expirationTime": null,
    "keys": {
        "p256dh": "BPo4NJH6nyo5uBx9n87gq6QnBipCZ00mWY38_HUq7_a-ONQWie8kShiQ1Rf1RzlQM3de1VKvWPhJqnLWGraJ3LI",
        "auth": "HXCuoaNfJ9jBbQlKsyTawQ"
    }
}"""
  
  """
  addEndPoint(key,feeds[0][0],end_point)
  print(isSub(key,feeds[0][0]))
  removeEndPoint(key,feeds[0][0])
  print(isSub(key,feeds[0][0]))

  """
  key ="a04a35acb3275bf5ddef83b1ac93a1660cafd4c64714b747113bd70428bb579577b5e2682eb9132daaf84c97c36dbd04e97aaa95c6a9a81ee10b30645f71cccc"
  print(getFeed(key))
  print(CreateFeed(key,"test"))
  print(getFeed(key))

  





  
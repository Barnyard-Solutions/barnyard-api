import mysql.connector
import hashlib
import time
from rich import print
from Push import send_web_push
import Account
from threading import Thread
import pause
from datetime import datetime, date, timedelta
def log(s):
    if debug:
        print("[bold yellow][DEBUG][Account][/bold yellow]",s)
debug =False

class MyBD:
  def __init__(self):
    self.PASS ="kKQdH@qX93"
    self.start()
    
  def start(self):
    self.mydb = mysql.connector.connect(
    host = "localhost",
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

def getEventFromFeedID(feed_id):
  
  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  comm = f"SELECT * FROM event WHERE event.FEED_ID = '{feed_id}' ORDER BY date;"
  

  cursor.execute(comm)
  res2 = cursor.fetchall()
  return res2
def getClassFromFeedID(feed_id):

  cursor = mydb.cursor()
  cursor.execute("USE barnyard;")
  comm = f"SELECT * FROM class WHERE  FEED_ID = '{feed_id}'  ORDER BY CLASS_DAY "
    
  cursor.execute(comm)
  res = cursor.fetchall()
  return res
 

def updateNotify():

  for i in Account.getEndPoint():
    
    if getClassFromFeedID(i[2]) != []:
      for z in getEventFromFeedID(i[2]):
        

        delta = (datetime.today()-datetime.strptime(z[3], '%Y-%m-%d')).total_seconds() // ( 60 * 60 * 24)
        

        for c in getClassFromFeedID(i[2]):
          
          if delta == c[3]:
            sub_info = {"endpoint":i[3],"keys": {i[4]: i[5],"auth": i[6]}}

            
            try:
              send_web_push(sub_info,{"title": c[2], "body": z[1]+" "+z[2] })
              print("[bold green][INFO][/bold green] Notification sent")
              print("[bold yellow][DEBUG][/bold yellow]",sub_info)
            except:
              print("[bold green][INFO][/bold green] End point is dead removing..")
              t1 = Thread(target=Account.removeEndPoint,args=[i[1],i[2],Account.getBetterEndPoint(i)]).start()
while True:

  if int(datetime.now().strftime('%H')) < 8:
    d = datetime.now()
    d = datetime(int("20"+d.strftime('%y')),int(d.strftime('%m')),int(d.strftime('%d')),8,0)
    p
      
  else:
    d = datetime.now()
    d = datetime(int("20"+d.strftime('%y')),int(d.strftime('%m')),int(d.strftime('%d')),8,0) + timedelta(days=1)
    

  print(f"[bold green][INFO][/bold green] tweeting start at {d}")
  pause.until(d)
  updateNotify()
  





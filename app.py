from flask import Flask
from flask_cors import CORS, cross_origin
from flask import request, jsonify, Response, json
from Account import Login, CreateAccount,GenerateKey,getFeed, getEvent,addEvent,CreateFeed,removeFeed,removeEvent,addClass, removeClass, getClass, addEndPoint, removeEndPoint,isSub,addEndPoint
from rich import print
from datetime import datetime
import time
import json

def log(s):
    if debug:
        print("[bold yellow][DEBUG][/bold yellow]",s)
debug =False


app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

@app.route('/')
def hello():
    return 'BarnYard API'

@app.route("/signup", methods=['GET', 'POST'])
def signup():
    data = request.get_json()
    
    mail = data["mail"]

    pas = data["pass"]
      
    res= CreateAccount(mail,pas)
    key = GenerateKey(mail,pas)
    res = {"login": res, "key":key}
    print("[info] key: ",res)
    response = jsonify(res)

    print("[bold green][INFO][/bold green]","Sign up",datetime.now())

    return response


@app.route("/login", methods=['GET', 'POST'])
def tryLogin():
    data = request.get_json()
    
    mail = data["mail"]

    pas = data["pass"]
      
    res= Login(mail,pas)
    key = GenerateKey(mail,pas)
    res = {"login": res, "key":key}
    response = jsonify(res)

    print("[bold green][INFO][/bold green]","Login",datetime.now())
    log(response)
    log(res)
    return response

@app.route("/addfarm", methods=['GET', 'POST'])
def addfarm():
    data = request.get_json()

    key = data["key"]

    name = data["name"]


    print(key,name, flush=True)
  
    CreateFeed(key, name)
    res = {"succes":True}
    response = jsonify(res)
    print("[bold green][INFO][/bold green]","Add Farm",datetime.now())
    log(response)
    log(res)
    return response


@app.route("/getfeed", methods=['GET', 'POST'])
def getfeed():
    data = request.get_json()

    key = data["key"]

    log("key")
    log(key)

    data = getFeed(key)
    res = {"feed":data}
    response = jsonify(res)
    print("[bold green][INFO][/bold green]","Get Feed",datetime.now())
    log(response)
    log(res)
    return response

@app.route("/getevent", methods=['GET', 'POST'])
def getevent_():
    data = request.get_json()

    key = data["key"]

    feed_id = data["feed"]



    data = getEvent(key, feed_id)
    res = {"event":data}
    response = jsonify(res)
    print("[bold green][INFO][/bold green]","Get Event",datetime.now())
    log(response)
    log(res)
    return response

@app.route("/addevent", methods=['GET','POST'])
def addevent():
    data = request.get_json()

    key = data["key"]
    feed = data["feed"]
    name1 = data["name1"]
    name2 = data["name2"]
    date = data["date"]

    data = addEvent(key,name1,name2,date,feed)
    if data:
        res = {"succes":True}
    else:
        res = {"succes":False}
    response = jsonify(res)
    print("[bold green][INFO][/bold green]","Add Event",datetime.now())
    log(response)
    log(res)
    return response

@app.route("/removefeed", methods=['GET','POST'])
def callremovefeed():
    data = request.get_json()
    key = data["key"]
    feed = data["feed"]

    res = removeFeed(key,feed)
    print("[bold green][INFO][/bold green]","Remove Feed",datetime.now())
    return jsonify({"succes":res})

@app.route("/removeevent", methods=['GET','POST'])
def callremoveevent():
    data = request.get_json()
    key = data["key"]
    event = data["event"]
    feed = data["feed"]

    res = removeEvent(key,feed,event)
    print("[bold green][INFO][/bold green]","Remove Event",datetime.now())
    return jsonify({"succes":res})

@app.route("/addclass", methods=['GET','POST'])
def calladdClass():
    data = request.get_json()
    key = data["key"]
    name = data["name"]
    feed = data["feed"]
    day = data["day"]
    

    res = addClass(key,feed,name,day)
    print("[bold green][INFO][/bold green]","Add Class",datetime.now())
    return jsonify({"succes":res})

@app.route("/removeclass", methods=['GET','POST'])
def callremoveclass():
    data = request.get_json()
    key = data["key"]
    class_id = data["class"]
    feed = data["feed"]

    res = removeClass(key,feed,class_id)
    print("[bold green][INFO][/bold green]","Remove Class",datetime.now())
    return jsonify({"succes":res})

@app.route("/getclass", methods=['GET','POST'])
def callgetclass():
    data = request.get_json()
    key = data["key"]
    feed = data["feed"]

    res = getClass(key,feed)
    print("[bold green][INFO][/bold green]","Get Class",datetime.now())
    return jsonify({"class":res})


@app.route("/addendpoint", methods=['GET','POST'])
def calladdEndPoint():
    data = request.get_json()
    key = data["key"]
    feed = data["feed"]
    end_point = data["endpoint"]

    res = addEndPoint(key,feed,end_point)
    print("[bold green][INFO][/bold green]","Add End Point",datetime.now())
    return jsonify({"succes":True})

@app.route("/removeendpoint", methods=['GET','POST'])
def callremoveEndPoint():
    data = request.get_json()
    key = data["key"]
    feed = data["feed"]
    end_point = data["endpoint"]


    res = removeEndPoint(key,feed,end_point)
    print("[bold green][INFO][/bold green]","Remove End Point",datetime.now())
    return jsonify({"succes":True})

@app.route("/issub", methods=['GET','POST'])
def callisSub():
    data = request.get_json()
    key = data["key"]
    feed = data["feed"]
    end_point = data["endpoint"]


    res = isSub(key,feed,end_point)
    print("[bold green][INFO][/bold green]","Is Sub",datetime.now())
    return jsonify({"succes":res})



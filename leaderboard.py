import json
import time
from datetime import date
from datetime import datetime
from datetime import timezone
from pathlib import Path

import requests

home = str(Path.home())
with open(home + "/.adventofcode_leaderboard_session", 'rt') as f:
    session = f.read()

session = session.strip()
  
url = 'https://adventofcode.com/2019/leaderboard/private/view/264332.json'
resp = requests.get(url, headers={'Cookie': "session=" + session})
resp.raise_for_status()
data = resp.json()

def get_num_challenges_open(event_year):
    # UTC : Dec 2nd @ 5am is Dec 1st @ 11 am CST, 11am CST is when challenges unlock for the 'following day' if using CST
    now = datetime.now(timezone.utc)
    year = now.year
    month = now.month
    day = now.day
    hour = now.hour

    # print("  now =", now)
    # print(" year =", now.year)
    # print("month =", now.month)
    # print("  day =", now.day)
    # print(" hour =", now.hour)

    if year > event_year:
        return 25
    
    if year < event_year:
        return 0

    # If we get here, we're in the correct year
    if month < 12:
        return 0
    
    if day > 25:
        return 25
    
    if hour >=5:
        return day
    else:
        return day-1

class Member:
    def __init__(self, member):
        self.num_stars = member['stars']
        self.name = member['name']
        self.last_star_ts = int(member['last_star_ts']) if 'last_star_ts' in member else None
        self.stars_by_day = member['completion_day_level']
    
    def __str__(self):
        return f"Member(name={self.name}, stars={self.num_stars})"
    
    def __repr__(self):
        return self.__str__()

    def has_completed(self, day, challenge=1):
        """
        Attributes
        -------------
        day : int
        1-25, the day of event of code
        
        part : int
        1 or 2, the number of the challenge

        Returns
        -------------
        Timestamp : if the user has completed the challnege for that day
        None : if the user has not completed the challenge for that day
        """
        day = str(day)
        challenge = str(challenge)
        if day not in self.stars_by_day:
            return False
        
        day_data = self.stars_by_day[day]

        if challenge not in day_data:
            return False

        return day_data[challenge]

# print(get_num_challenges_open(2019))
max_challenges = get_num_challenges_open(2019)

members = []
raw_members = data['members']
for member_id in raw_members:
    curmember = raw_members[member_id]
    member = Member(curmember)
    members.append(member)

members.sort(key=lambda member: member.last_star_ts)
members.sort(key=lambda member: member.num_stars, reverse=True)
# print(members) 

def print_heading():
    print(f"Num Name                              Last Star  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25")
    print(f"{'-'*3} {'-'*20} {'-'*22} ", end='') 
    for _ in range(1, 26):
        print(f"{'-'*2} ", end='')
    
    print()

  
# print(time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(1575734637)))
def print_member(index, member):
    if member.last_star_ts:
        last_star = time.strftime('%Y-%m-%d %I:%M:%S %p', time.localtime(member.last_star_ts))
    else:
        last_star = 'None'

    print(f"{index+1:3} {member.name:>20} {last_star} ", end='')

    for day in range(1, 26):
        token = ' '
        if member.has_completed(day, 2):
            token = 'X'
        elif member.has_completed(day, 1):
            token = '/'
        print(f"{token:>2} ", end='')
    
    print()


print_heading()
for i, member in enumerate(members):
    print_member(i, member)


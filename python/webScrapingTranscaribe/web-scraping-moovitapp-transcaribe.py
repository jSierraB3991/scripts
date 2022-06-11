import requests
from bs4 import BeautifulSoup

class Route:
    def __init__(self, name, link):
        self.name = name
        self.link = link
        self.isLine = self.link.startswith("https")

class Bus:
    def __init__(self, name, link, isRoute, routeFather):
        self.name = name
        self.link = link
        self.isRoute = isRoute
        self.routeFather = routeFather
        self.bus_route = []

class StopRoute:
    def __init__(self, name):
        self.name = name

url_base = "https://moovitapp.com/index/en/"
reponse = requests.get(url_base + "public_transit-Cartagena-5477")
soup = BeautifulSoup(reponse.text)
routes_all = soup.find_all("a", { "class": "agency-group" })

routes = []
for bus_for_route in routes_all:
    routes.append(Route(bus_for_route.find("h3").string, bus_for_route["href"]))

buses = []
for route in routes:
    if route.isLine:
        buses.append(Bus(route.name, route.link, True, ""))
    else:
        response = requests.get(url_base + route.link)
        soup = BeautifulSoup(response.text)
        buses_for_route = soup.find_all("li", { "class": "line-item" })
        for bus in buses_for_route:
            buses.append(Bus(bus.find("strong").string, bus.find("a")["href"], False, route.name))

for bus in buses:
    response = requests.get(bus.link)
    soup = BeautifulSoup(response.text)
    buses_for_route = soup.find_all("li", { "class": "stop-container" })
    stopRoute =  []
    for bus_for_route in buses_for_route:
        stopRoute.append(StopRoute(bus_for_route.find("h3").string))
    bus.bus_route = stopRoute

for route in routes:
    print("INSERT INTO routes(link, name, isLine) VALUES (" + route.link + ", " + route.name + ", " + str(route.isLine) + ")")

for bus in buses:
    print("INSERT INTO bus(name, link, isRoute) VALUES(" + bus.name + ", " + bus.link + ", " + str(bus.isRoute) + ")")

for bus in buses:
    for route_bus in bus.bus_route:
        print("INSERT INTO route_bus(name_bus, nameRoute) VALUES(" + bus.name + ", " + route_bus.name + ")")

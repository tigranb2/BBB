import re

def read_file(file):
    with open(file) as f:
        return f.readlines()

def calculate_metrics(lat):
  print("\n\n***** transaction report *****")
  print("\t", "latency", "average", (round(sum(lat) / len(lat), 2))/1000, "(sec)")
  print("\t", "latency", "median\t", (round(sorted(lat)[int(len(lat) * 0.5)], 2))/1000, "(sec)")
  print("\t", "latency", "95th\t", (round(sorted(lat)[int(len(lat) * 0.95)], 2))/1000, "(sec)")
  print("*****end of the report*****")
  
def main():
  lat = read_file("./data/performance.txt")
  lat = [int(i) for i in lat]
  calculate_metrics(lat)

main()

states_needed = set(["m", "w", "o", "i", "n", "u", "c", "a"])

stations = {}
stations["k1"] = set(["i", "n", "u"])
stations["k2"] = set(["w", "i", "m"])
stations["k3"] = set(["o", "n", "c"])
stations["k4"] = set(["n", "u"])
stations["k5"] = set(["c", "a"])

final_stations = set()

while states_needed:
    best_station = None
    states_covered = set()
    for station, states_for_station in stations.items():
        covered = states_needed & states_for_station
        if len(covered) > len(states_covered):
            best_station = station
            states_covered = covered
    states_needed -= states_covered
    final_stations.add(best_station)

print(final_stations)

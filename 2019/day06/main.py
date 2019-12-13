import sys
from collections import defaultdict


def construct_orbital_maps():
    satellites_to_bodies = {}
    bodies_to_satellites = defaultdict(list)

    orbits = [x.split(")") for x in sys.argv[1:]]

    for body, satellite in orbits:
        satellites_to_bodies[satellite] = body
        bodies_to_satellites[body].append(satellite)

    return satellites_to_bodies, bodies_to_satellites


def compute_orbit_counts(satellites_to_bodies, bodies_to_satellites, roots):
    orbit_counts = {}
    to_process = []

    for root in roots:
        orbit_counts[root] = 0
        to_process += bodies_to_satellites[root]

    while to_process:
        next_sat = to_process.pop()
        if next_sat in orbit_counts:
            continue
        else:
            parent = satellites_to_bodies[next_sat]
            if parent in orbit_counts:
                orbit_counts[next_sat] = orbit_counts[parent] + 1
            else:
                to_process.insert(0, parent)
                to_process.insert(1, next_sat)

        # add any satellites this one has
        to_process += bodies_to_satellites[next_sat]

    return orbit_counts


def main():
    satellites_to_bodies, bodies_to_satellites = construct_orbital_maps()
    roots = set(bodies_to_satellites.keys()) - set(satellites_to_bodies.keys())

    orbit_counts = compute_orbit_counts(
        satellites_to_bodies, bodies_to_satellites, roots,
    )

    print(sum(orbit_counts.values()))


if __name__ == "__main__":
    main()

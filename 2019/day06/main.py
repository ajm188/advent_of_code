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


def find_path_to_root(node, satellites_to_bodies):
    next_node = node
    path = []

    while next_node in satellites_to_bodies:
        path.append(next_node)
        next_node = satellites_to_bodies[next_node]

    return path[1:]


def main():
    satellites_to_bodies, bodies_to_satellites = construct_orbital_maps()
    roots = set(bodies_to_satellites.keys()) - set(satellites_to_bodies.keys())

    orbit_counts = compute_orbit_counts(
        satellites_to_bodies, bodies_to_satellites, roots,
    )

    print(sum(orbit_counts.values()))

    san_path = find_path_to_root('SAN', satellites_to_bodies)
    you_path = find_path_to_root('YOU', satellites_to_bodies)

    common = [(i, x) for i, x in enumerate(you_path) if x in san_path]
    print(common[0][0] + san_path.index(common[0][1]))


if __name__ == "__main__":
    main()

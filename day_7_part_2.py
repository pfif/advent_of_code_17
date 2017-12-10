import re
# from copy import deepcopy

input_file = open("day_7_input", "r")
regex_for_line = re.compile("([a-z]+) \(([0-9]+)\)(?: -> ((?:[a-z]+, )*[a-z]+))?")


class Node(object):
    def __init__(self, name):
        self.name = name
        self.weight = None
        self.children = []

    def actual_weight(self):
        return sum([self.weight] + [child.actual_weight() for child in self.children])


def get_or_create_node(name, node_list):
    node = {node.name: node for node in node_list}.get(name, None)

    if node is None:
        node = Node(name)
        node_list.append(node)

    return node


def parse_line(line, node_list):
    line_parsed = regex_for_line.match(line)
    parsed_node_name, parsed_node_weight, parsed_node_children = line_parsed.group(1), line_parsed.group(2), line_parsed.group(3)
    parsed_node_weight = int(parsed_node_weight)
    parsed_node_children = parsed_node_children.split(", ") if parsed_node_children else []

    node = get_or_create_node(parsed_node_name, node_list)
    node.weight = parsed_node_weight
    node.children = [
        get_or_create_node(name, node_list) for name in parsed_node_children
    ]


def get_balanced_and_unbalanced_node_of_disc(node):
    weights = [child.actual_weight() for child in node.children]
    weights_set = set(weights)
    if len(weights_set) > 1:
        unbalanced_weight = {weights.count(weight): weight for weight in weights}[1]

        nodes_by_weight = {child.actual_weight(): child for child in node.children}
        unbalanced_node = nodes_by_weight.pop(unbalanced_weight)
        balanced_node = nodes_by_weight.popitem()[1]

        return unbalanced_node, balanced_node
    return None, None


def find_unbalanced_node_and_a_sibbling_on_disc(node):
    unbalanced_node, balanced_node = get_balanced_and_unbalanced_node_of_disc(node)

    # this if statement confirms that the problematic node is not on a disc above
    if unbalanced_node:
        child_unbalanced_node, child_balanced_node = find_unbalanced_node_and_a_sibbling_on_disc(unbalanced_node)
        if child_unbalanced_node:
            return child_unbalanced_node, child_balanced_node
        else:
            return unbalanced_node, balanced_node
    return None, None


def correct_weight_of_node_that_creates_unbalance(unbalanced_node, balanced_node):
    return unbalanced_node.weight - (unbalanced_node.actual_weight() - balanced_node.actual_weight())


def find_root_node(candidate):
    for node in node_list:
        if candidate in node.children:
            return find_root_node(node)
    return candidate


node_list = []
for line in input_file:
    parse_line(line, node_list)

root_node = find_root_node(node_list[1])
print(correct_weight_of_node_that_creates_unbalance(*find_unbalanced_node_and_a_sibbling_on_disc(root_node)))

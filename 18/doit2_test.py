#!/usr/bin/env python3
# UTs written by chatgpt (!)
from doit2 import vmap
import unittest

class TestVMap(unittest.TestCase):
    def test_color(self):
        # Create a vmap object with a list of points
        points = [(1, 2, 3), (4, 5, 6), (7, 8, 9)]
        vm = vmap(points)

        # Check that the color of each point is correct
        for v, c in vm.colors.items(): 
            if v in points:            #
                self.assertEqual(c, 0) # I had to correct only these manually.
            else:                      #
                self.assertEqual(c, 2) #

        # Check that the color of a point outside the bounding box is 0
        v = (0, 0, 0)
        self.assertEqual(vm.get_color(v), 0)

        # Check that the color of a point inside the bounding box but not in the list of points is 2
        v = (2, 3, 4)
        self.assertEqual(vm.get_color(v), 2)

if __name__ == '__main__':
    unittest.main()

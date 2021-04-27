#!/usr/bin/python3
# -*- coding: utf-8 -*-
#
# (c) 2021 BridgingIT GmbH
# Author: Bastian Appenzeller
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
#
# This code generates random intervals which are used as test data.
#
# This is part of a demo for a project application for the Daimler TSS GmbH.

__shell_usage__ = """
usage: python3 create_test_data.py [options]

Options are

-h, --help            help text
-v, --verbose         verbose mode

--valid               create valid test data (default), e. g. [10000:20000] [90:121] ...
--invalid             create 90% valid nad 10% invalid test data, e. g. [100:90] [55:a] ...

--records=<number>    creates <number> records of test data (default is 200.000)
"""

import sys
import random

def create_valid_intervall_string():
    """ Creates a valid intervall string usind random numbers """
    # In case Python2 is used: max int is 2147483647
    interval_lower_value = random.randint(-2147483647, 2147483647)
    interval_higher_value = random.randint(interval_lower_value + 1,
                                                   2147483647)
    return '[' + str(interval_lower_value) + ',' + str(interval_higher_value)+']'


def create_test_data(amount_records):
    """ Creates valid or invalid intervals """
    global g_verbose_mode
    global g_valid_data
    interval_lower_value = ()
    interval_higher_value = ()

    # Open file for randeom data
    if g_valid_data == True:
        test_data_file_name = str(amount_records) + '_intervals.json'
    else:
        test_data_file_name = str(amount_records) + '_invalid_intervals.json'
    test_data_object = open(test_data_file_name , 'w')

    # Create random data
    for record_number in range(amount_records):
        if record_number == 0:
            test_data_object.write('[')
        if g_valid_data == True:
            interval = (create_valid_intervall_string() + ' ')
            if g_verbose_mode == True:
                print(str(record_number + 1) + ' ' + interval)
            test_data_object.write(create_valid_intervall_string() + ' ')
        else:
            if random.randint(0, 19) == 0:
                # 5% chance to create an invalid record consisting of
                # characters
                random_string = ''
                for _ in range(10):
                    # Create an upper case character
                    random_int = random.randint(97, 97 + 26 - 1)
                    # 50% chance to switch to loer case character
                    if random.randint(0, 1) == 1:
                       # Convert character to lowercase
                       random_int -= 32
                    # Appending random character
                    random_string += (chr(random_int))
                # 50% chance to use it as lower or higher interval limit
                if random.randint(0, 1) == 1:
                    interval_lower_value = random_string
                    interval_higher_value = random.randint(-2147483647, 2147483647)
                else:
                    interval_lower_value = random.randint(-2147483647, 2147483647)
                    interval_higher_value = random_string
                interval = '[' + str(interval_lower_value) + ',' + str(interval_higher_value)+'],'
                if g_verbose_mode == True:
                    print(str(record_number + 1) + ' ' + interval)
                test_data_object.write(interval)
            else:
                # create a valid record
                interval = (create_valid_intervall_string() + ',')
                if g_verbose_mode == True:
                    print(str(record_number + 1) + ' ' + interval)
                test_data_object.write(create_valid_intervall_string())

    test_data_object.write(']')
    # Close file for randeom data
    test_data_object.close()


if __name__ == '__main__':
    global g_verbose_mode
    global g_valid_data

    g_verbose_mode = False
    g_valid_data = True
    args_to_process = []
    amount_records = 200000

    # Collect call parameter
    for arg in sys.argv:
        if arg.endswith('create_test_data.py'):
            pass
        elif arg in ['-h', '--help']:
            print(__shell_usage__)
            sys.exit()
        elif arg in ['-v', '--verbose']:
            g_verbose_mode =  True
        elif arg == '--invalid':
            g_valid_data = False
        elif arg == '--valid':
            g_valid_data = True
        elif arg.startswith('--records='):
            record_string = arg[10:]
            if record_string.isdigit():
                amount_records = int(record_string)
            else:
                print(record_string + ' is not a number.')
                sys.exit()
        else:
            print('Parameter ' + arg + ' is not supported.')
            sys.exit()
    create_test_data(amount_records)

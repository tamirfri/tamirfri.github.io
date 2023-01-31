
# python wrapper for package github.com/tamirfri/py-run-go-run-js within overall package lodash
# This is what you import to use the package.
# File is generated by gopy. Do not edit.
# gopy build -output=py_run_go -vm=python3 github.com/tamirfri/py-run-go-run-js

# the following is required to enable dlopen to open the _go.so file
import os,sys,inspect,collections
try:
	import collections.abc as _collections_abc
except ImportError:
	_collections_abc = collections

cwd = os.getcwd()
currentdir = os.path.dirname(os.path.abspath(inspect.getfile(inspect.currentframe())))
os.chdir(currentdir)
from . import _lodash
from . import go

os.chdir(cwd)

# to use this code in your end-user python file, import it as follows:
# from lodash import lodash
# and then refer to everything using lodash. prefix
# packages imported by this package listed below:




# ---- Types ---


#---- Enums from Go (collections of consts with same type) ---


#---- Constants from Go: Python can only ask that you please don't change these! ---


# ---- Global Variables: can only use functions to access ---


# ---- Interfaces ---


# ---- Structs ---


# ---- Slices ---


# ---- Maps ---


# ---- Constructors ---


# ---- Functions ---
def Find(objects, partialObject):
	"""Find(str objects, str partialObject) str, str
	
	Find the first object in `objects` that contains `partialObject`
	"""
	return _lodash.lodash_Find(objects, partialObject)


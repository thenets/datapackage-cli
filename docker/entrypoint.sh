#!/bin/bash

# Creater virtualenv if not exist 
if [ ! -d env/bin ]; then
    virtualenv -p python3 env
fi

# Active virtualenv
source env/bin/activate

# Install dependencies
if [ -f src/pip-requirements.txt ]; then
    pip install -r src/pip-requirements.txt
fi

# Run main script
python src/*
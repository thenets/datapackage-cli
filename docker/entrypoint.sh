#!/bin/bash

# Helper: Colors
COLOR_CYAN='\033[0;36m'
COLOR_DEFAULT='\033[0m'

# Creater virtualenv if not exist 
if [ ! -f env/bin/activate ]; then 
    echo $(COLOR_CYAN)"# CRIANDO VIRTUALENV"$(COLOR_DEFAULT)
    virtualenv -p python3 env
    echo ""
fi

# Active virtualenv
source env/bin/activate

# Install dependencies
PIP_FILE=src/pip-requirements.txt
if [ -f $PIP_FILE ]; then
    md5_cache_file=env/pip_md5.cache
    touch $md5_cache_file
    md5_cached=($(cat $md5_cache_file))
    md5=($(md5sum $PIP_FILE))

    # Install dependencies if not cached
    if [ ! "$md5" = "$md5_cached"  ]; then
        echo $(COLOR_CYAN)"# BAIXANDO DEPENDÊNCIAS"$(COLOR_DEFAULT)
        pip install -r src/pip-requirements.txt
        echo ""
        # cache new hash
        echo $md5 > $md5_cache_file
    fi

fi

# Run main script
cd src/
python *
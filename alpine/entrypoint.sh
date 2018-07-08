#!/bin/bash

# Main vars
ROOT_PATH=$(pwd)
VENV_PATH=$ROOT_PATH/.brasilio/venv
SRC_PATH=src/
RUN_FILE=$ROOT_PATH/.brasilio/run.sh
PIP_FILE=$ROOT_PATH/.brasilio/requirements.txt

# Colors
COLOR_CYAN='\033[0;36m'
COLOR_DEFAULT='\033[0m'

# Creater virtualenv if not exist
if [ ! -f $VENV_PATH/bin/activate ]; then 
    echo -e "$COLOR_CYAN# CRIANDO VIRTUALENV$COLOR_DEFAULT"
    virtualenv -p python3 $VENV_PATH
    echo ""
fi

# Active virtualenv
source $VENV_PATH/bin/activate

# Install dependencies
if [ -f $PIP_FILE ]; then
    md5_cache_file=$VENV_PATH/pip_md5.cache
    touch $md5_cache_file
    md5_cached=($(cat $md5_cache_file))
    md5=($(md5sum $PIP_FILE))

    # Install dependencies if not cached
    if [ ! "$md5" = "$md5_cached"  ]; then
        echo -e "$COLOR_CYAN# BAIXANDO DEPENDÊNCIAS$COLOR_DEFAULT"
        pip install -r $PIP_FILE
        echo ""
        # cache new hash
        echo $md5 > $md5_cache_file
    fi

fi

# Run main script
if [ -f $RUN_FILE ]; then
    # Copy run file to root
    cp $RUN_FILE $ROOT_PATH/run.sh
    
    # Run script
    chmod +x $RUN_FILE
    $RUN_FILE
else
    # Run default Python file
    cd $SRC_PATH/
    python capture.py
fi

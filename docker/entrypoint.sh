#!/bin/bash

# Creater virtualenv if not exist 
if [ ! -f env/bin/activate ]; then 
    echo "# O virtualenv não foi encontrado."
    echo "# Criando novo virtualenv..."
    virtualenv -p python3 env
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
        echo "# Instalando novas bibliotecas..."
        pip install -r src/pip-requirements.txt
        echo "... instalado!"
        echo ""
        # cache new hash
        echo $md5 > $md5_cache_file
    fi

fi

# Run main script
python src/*
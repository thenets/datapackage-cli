FROM heroku/heroku:18-build

LABEL   io.brasil="Brasil IO" \
        io.brasil.cli="foo" \
        version="0.1" \
        description="Imagem do Brasil.io para executar os código de extração de dados." \
        maintainer="luiz@thenets.org"


# Copy files
ADD entrypoint.sh $HOME/
RUN set -x \
    # Add entrypoint file
    && chmod +x $HOME/entrypoint.sh \
    # Create 'app' dir and empty file
    && mkdir -p $HOME/src \
    && echo 'print("Volume não foi montado em /app/src")' > $HOME/src/capture.py \
    # HACK to fix volumes permissions
    && mkdir -p $HOME/.brasilio $HOME/src $HOME/package \
    # Set permissions
    && chown -R 1000.1000 $HOME

# TODO
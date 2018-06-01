FROM alpine

MAINTAINER luiz@thenets.org

# Main dependencies
RUN apk --no-cache add \
        python3 py3-pip git bash \
        python3-dev build-base gcc libxml2-dev libxslt-dev

# Envs
ENV HOME=/app \
    USER=brasilio

# Setup
WORKDIR $HOME
RUN set -x \
    # Add user
    && adduser -D -u 1000 -s /bin/bash $USER \
    # Setup virtualenv
    && python3 -m pip install pip --upgrade \
    && python3 -m pip install virtualenv \
    # Set permissions
    && chown -R 1000.1000 $HOME \
    && virtualenv -p python3 env

# Copy files
ADD *.sh $HOME/
RUN set -x \
    # Add entrypoint file
    && chmod +x $HOME/entrypoint.sh \
    # Create 'app' dir and empty file
    && mkdir -p $HOME/src \
    && echo 'print("Volume não foi montado em /app/src")' > $HOME/src/empty.py \
    # Set permissions
    && chown -R 1000.1000 $HOME

# Python dependencies
USER $USER

CMD ["/app/entrypoint.sh"]
VOLUME ["$HOME/env", "$HOME/src", "$HOME/data"]
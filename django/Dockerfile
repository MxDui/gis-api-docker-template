# The image you are going to inherit your Dockerfile from
FROM python:3.10-alpine
# Necessary, so Docker doesn't buffer the output and that you can see the output 
# of your application (e.g., Django logs) in real-time.
ENV PYTHONUNBUFFERED 1
# Make a directory in your Docker image, which you can use to store your source code
RUN mkdir /example
# Set the /example directory as the working directory
WORKDIR /example
# Copies from your local machine's current directory to the example folder 
# in the Docker image
COPY . .
# Copy the requirements.txt file adjacent to the Dockerfile 
# to your Docker image
COPY ./requirements.txt /requirements.txt

# Update pip 
RUN pip install --upgrade pip

# install gunicorn
RUN pip install gunicorn

RUN apk add --update postgresql-client jpeg-dev
RUN apk add --update --virtual .tmp-build-deps \
    gcc libc-dev linux-headers postgresql-dev musl-dev zlib zlib-dev
# Install postgis and gdal
RUN apk add --no-cache postgis

# install proj
RUN apk add --no-cache proj-util


ENV PROJ_DIR=/usr

RUN apk add --no-cache gcc
RUN apk add --no-cache gdal
RUN apk add --no-cache gdal-dev
RUN apk add --no-cache geos
RUN apk add --no-cache geos-dev
RUN apk add --no-cache build-base
RUN apk add --no-cache zlib

# run version of geos
RUN geos-config --version

RUN export CPLUS_INCLUDE_PATH=/usr/include/gdal
RUN export C_INCLUDE_PATH=/usr/include/gdal
RUN export LDFLAGS="-L/usr/local/opt/zlib/lib"
RUN export CPPFLAGS="-I/usr/local/opt/zlib/include"


# add gdal to path to sh 
RUN echo 'export PATH="/usr/bin/gdal:$PATH"' >> ~/.bashrc

# Install the requirements.txt file in Docker image
RUN pip install -r /requirements.txt
# Create a user that can run your container

USER root
RUN chmod +x /example/entrypoint.sh

# add gunicon to path to sh
RUN echo 'export PATH="/usr/local/bin/gunicorn:$PATH"' >> ~/.bashrc





# CMD [ "gunicorn", "api.wsgi:application", "--bind", "0.0.0.0:8000"]

ENTRYPOINT [ "/bin/sh", "/example/entrypoint.sh"]


# Use an official Python runtime as a parent image
FROM python:3.9-slim-buster

ENV DB_HOST=barnapi1

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in requirements.txt
RUN apt-get update && pip install --trusted-host pypi.python.org -r req.txt  

# Expose the port that the API will be served on
EXPOSE 5000

# Start the API using Gunicorn
CMD ["gunicorn", "--workers=1", "--bind=0.0.0.0:5000", "app:app"]
#CMD ["gunicorn", "--certfile", "fullchain.pem", "--keyfile", "privkey.pem", "--workers=1", "--bind=0.0.0.0:5000", "app:app"]
#CMD ["sh","-c","tail -f /dev/null"]

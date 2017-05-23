FROM scratch

ENV SERVICE_PORT 8000

EXPOSE $SERVICE_PORT

COPY step-by-step /

CMD ["/step-by-step"]

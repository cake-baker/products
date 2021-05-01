FROM scratch

COPY products /
ENTRYPOINT ["/products"]
EXPOSE 8080
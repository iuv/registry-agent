FROM busybox
ADD build/linux/registry-agent /app/registry-agent
ADD conf /app/conf
RUN chmod 775 /app/registry-agent
EXPOSE 80
WORKDIR /app
CMD ./registry-agent
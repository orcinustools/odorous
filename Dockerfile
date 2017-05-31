FROM pataquets/glusterfs

RUN mkdir -p /run/docker/plugins /mnt/state /mnt/volumes

COPY odorous odorous

CMD ["odorous"]
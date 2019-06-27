FROM cander56/vue
USER root
ENV HOMEDIR /home/vue
WORKDIR $HOMEDIR
COPY . .
RUN chown -R vue:vue /home/vue 
USER vue
ENTRYPOINT ["./entrypoint.sh"]
CMD ["sleep","3600"]

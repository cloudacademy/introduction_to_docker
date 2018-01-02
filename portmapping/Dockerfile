# Since we're using scratch we could omit this
# scratch is the default base.
FROM scratch

# Copy the webapp binary from the current dir to the root of the container
COPY webapp /

# Expose is more documentation than anything else.
EXPOSE 8080

# The default command to run when starting the container
CMD ["/webapp"]
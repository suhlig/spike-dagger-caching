# Spike on Dagger's Caching Behaviour

```command
$ go run .
Creating new Engine session... OK!
Establishing connection to Engine... 1: connect
1: > in init
1: starting engine
1: starting engine [1.02s]
1: starting session
1: [1.20s] OK!
1: starting session [0.18s]
1: connect DONE
OK!

8: resolve image config for docker.io/suhlig/b2:latest
8: > in from suhlig/b2
8: resolve image config for docker.io/suhlig/b2:latest DONE

15: pull docker.io/suhlig/b2:latest
15: > in from suhlig/b2
15: resolve docker.io/suhlig/b2@sha256:ba039fd53bbfbf52af1baa3a23993714063bfe9c78e84c1c4505edfbe91e2a35
15: resolve docker.io/suhlig/b2@sha256:ba039fd53bbfbf52af1baa3a23993714063bfe9c78e84c1c4505edfbe91e2a35 [0.01s]
15: pull docker.io/suhlig/b2:latest DONE

15: pull docker.io/suhlig/b2:latest CACHED
15: > in from suhlig/b2
15: pull docker.io/suhlig/b2:latest CACHED

13: exec b2 ls suhlig-transcription-test
13: [1.60s] Using https://api.backblazeb2.com
13: [1.60s] bar.txt
13: [1.60s] foo.txt
13: exec b2 ls suhlig-transcription-test DONE
Using https://api.backblazeb2.com
bar.txt
foo.txt
```

Now delete `bar.txt` from the bucket and run the same command again:

```command
$ go run .
Creating new Engine session... OK!
Establishing connection to Engine... 1: connect
1: > in init
1: starting engine
1: starting engine [0.98s]
1: starting session
1: [1.15s] OK!
1: starting session [0.16s]
1: connect DONE
OK!

8: resolve image config for docker.io/suhlig/b2:latest
8: > in from suhlig/b2
8: resolve image config for docker.io/suhlig/b2:latest DONE

15: pull docker.io/suhlig/b2:latest
15: > in from suhlig/b2
15: resolve docker.io/suhlig/b2@sha256:ba039fd53bbfbf52af1baa3a23993714063bfe9c78e84c1c4505edfbe91e2a35
15: resolve docker.io/suhlig/b2@sha256:ba039fd53bbfbf52af1baa3a23993714063bfe9c78e84c1c4505edfbe91e2a35 [0.01s]
15: pull docker.io/suhlig/b2:latest DONE

13: exec b2 ls suhlig-transcription-test CACHED
13: exec b2 ls suhlig-transcription-test CACHED
Using https://api.backblazeb2.com
bar.txt
foo.txt
```

The line saying `13: exec b2 ls suhlig-transcription-test CACHED` seems to indicate that this container was not invoked; instead a cached version of its output was probably returned.

# Reference

The registry image `suhlig/b2` was built using dagger like this:

```go
_, err := client.Container().
  From("python:alpine").
  WithExec([]string{"pip", "install", "-U", "b2"}).
  WithRegistryAuth(registryHostnme, registryUsername, registryPassword).
  Publish(ctx, fmt.Sprintf("%s/b2", registryUsername))
```

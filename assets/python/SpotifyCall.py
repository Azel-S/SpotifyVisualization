from dotenv import load_dotenv
import requests
import os

load_dotenv('.env')

# Get Access Token
token_response = requests.post("https://accounts.spotify.com/api/token", data={
                               "grant_type": "client_credentials", "client_id": os.getenv("CLIENT_ID"), "client_secret": os.getenv("CLIENT_SECRET")})

if token_response.ok:
    token = token_response.json().get('access_token')

    spotifyFile = open("input_spotify.csv", "r")

    tracksFile = open("output_tracks.csv", "w")
    artistsFile = open("output_artists.csv", "w")
    artistGenresFile = open("output_artists_genres.csv", "w")
    tracksGenresFile = open("output_tracks_genres.csv", "w")
    marketsFile = open("output_markets.csv", "w")

    for lineIndex, line in enumerate(spotifyFile):
        if lineIndex == 0:
            continue
        elif lineIndex == 5:
            break

        lineTokens = line.split(",")

        track_response = requests.get("http://api.spotify.com/v1/tracks/" + lineTokens[3], headers = {"Authorization": "Bearer " + token})
    
        if track_response.ok:
            track_json = track_response.json()

            # Artist Information
            for artist in track_json["artists"]:
                artist_response = requests.get("http://api.spotify.com/v1/artists/" + artist["id"], headers = {"Authorization": "Bearer " + token})
                if artist_response.ok:
                    artist_json = artist_response.json()
                    genres = artist_json["genres"]
                    artistsFile.write(artist_json["id"] + ", " + artist_json["name"] + ", " +
                                      str(artist_json["followers"]["total"]) + ", " +  str(artist_json["popularity"]) + "\n")
                    
                    # Genre Information
                    for genre in genres:
                        artistGenresFile.write(artist_json["id"] + ", " + genre + "\n")
                else:
                    print("Unable to access artist info, id:", artist["id"])

            # Market
            for market in track_json["available_markets"]:
                marketsFile.write(artist_json["id"] + ", " + market + "\n")

            # Explicit
            tracksFile.write(artist_json["id"] + ", " + track_json["id"] + ", " + track_json["name"] + ", " + ", " + ", ")
            print("Explicit:", track_json["explicit"])

            # Date
            print("Date:", track_json["album"]["release_date"])

            # print(track_response)

    spotifyFile.close()
    tracksFile.close()
    artistsFile.close()
    artistGenresFile.close()
else:
    print("Unable to gain token, please verify client_id and client_secret.")
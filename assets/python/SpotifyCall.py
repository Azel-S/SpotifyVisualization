from dotenv import load_dotenv
import requests
import os
import time

def getToken():
    load_dotenv('.env')

    token_response = requests.post("https://accounts.spotify.com/api/token", data={
                               "grant_type": "client_credentials", "client_id": os.getenv("CLIENT_ID"), "client_secret": os.getenv("CLIENT_SECRET")})

    if(token_response.ok):
        print(">>> Got new token: " + token_response.json().get('access_token'))
        return (token_response.json().get('access_token'), time.perf_counter())
    else:
        print("XXX Unable to gain token, please investigate.")
        return ("", time.perf_counter())

def main():
    (token, start_time) = getToken()

    if (token != ""):
        spotifyFile = open("input_spotify.csv", "r", encoding="utf8")

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

            print("> Line index " + str(lineIndex) + " Started.")
            
            lineTokens = line.split(",")

            track_response = requests.get(
                "http://api.spotify.com/v1/tracks/" + lineTokens[3], headers={"Authorization": "Bearer " + token})

            if (track_response.status_code == 429):
                print("XXX Track: Sleeping for: " + str(track_response.headers["Retry-After"]))
                time.sleep(int(track_response.headers["Retry-After"]))

            if track_response.ok:
                track_json = track_response.json()

                # Artist Information
                for artist in track_json["artists"]:
                    artist_response = requests.get(
                        "http://api.spotify.com/v1/artists/" + artist["id"], headers={"Authorization": "Bearer " + token})
                    
                    if (artist_response.status_code == 429):
                        print("XXX Artist: Sleeping for: " + str(artist_response.headers["Retry-After"]))
                        time.sleep(int(artist_response.headers["Retry-After"]))

                    if artist_response.ok:
                        artist_json = artist_response.json()
                        genres = artist_json["genres"]
                        artistsFile.write(artist_json["id"] + ", " + artist_json["name"] + ", " +
                                          str(artist_json["followers"]["total"]) + ", " + str(artist_json["popularity"]) + "\n")
                                          
                        # Genre Information
                        for genre in genres:
                            artistGenresFile.write(
                                artist_json["id"] + ", " + genre + "\n")
                    else:
                        print("Unable to access artist info, id:", artist["id"])

                # Market
                for market in track_json["available_markets"]:
                    marketsFile.write(track_json["id"] + ", " + market + "\n")

                # Track
                exp = ("true" if track_json["explicit"] else "false")
                tracksFile.write(track_json["id"] + ", " + artist_json["id"] + ", " + track_json["name"] + ", " + track_json["album"]["release_date"] + ", " +
                                 lineTokens[4] + ", " + lineTokens[6] + ", " + lineTokens[7] + ", " + lineTokens[8] + ", " + lineTokens[9] + ", " + lineTokens[10] + ", " +
                                 lineTokens[11] + ", " + lineTokens[12] + ", " + lineTokens[13] + ", " + lineTokens[14] + ", " + lineTokens[15] + ", " +
                                 lineTokens[16] + ", " + lineTokens[17] + ", " + lineTokens[18] + ", " + lineTokens[19][:len(lineTokens[19]) - 1] + ", " +
                                 ("true" if track_json["explicit"] else "false") + "\n")

                # Grab new token if enough time has passed.
                if(time.perf_counter() - start_time > 3000):
                    (token, start_time) = getToken()

                    if(token == ""):
                        break;
            
            print("> Line index " + str(lineIndex) + " Finished.")

        spotifyFile.close()
        tracksFile.close()
        artistsFile.close()
        artistGenresFile.close()

if(__name__ == "__main__"):
    main()
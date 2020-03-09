package main


import (
        "fmt"
        "log"
        "os"
        _ "net/http"
        _ "io/ioutil"
        "path/filepath"
        "strings"

        "golang.org/x/net/context"
        "google.golang.org/api/drive/v3"
)

func findNode(srv *drive.Service, parent string, name string, isFolder bool) (string) {
        query := fmt.Sprintf("\"%s\" in parents and name = \"%s\"", parent, name)
        if isFolder {
            query += "and mimeType = \"application/vnd.google-apps.folder\""
        }else {
            query += "and mimeType != \"application/vnd.google-apps.folder\""
        }
        r, err := srv.Files.List().PageSize(10).
            Fields("nextPageToken, files(id, name, parents)").Q(query).Do()
        if err != nil {
            log.Fatalf("Unable to retrieve files: %v", err)
        }
        /*
        {
            "nextPageToken": "",
            "files": []
        }
        */
        if len(r.Files) == 0 {
            return ""
        }
        if len(r.Files) > 1 {
            fmt.Println("Weird...")
            for _, i := range r.Files {
                fmt.Printf("%s (%s)\n", i.Name, i.Id)
            }
        }
        fmt.Printf("(%s, %s)\n", strings.Join(r.Files[0].Parents, " > "), r.Files[0].Id)
        return r.Files[0].Id
}

func createFolder(srv *drive.Service, parent string, name string, recursive bool) (string) {
        if recursive {
            /*
                1. a/b
                2. ./a/b (X)
                3. /a/b
                4. a/b/
                -> strip "/" first!
            */
            name = strings.Trim(name, "/")
            segments := strings.Split(name, "/")
            for _, _name := range segments {
                if _id := findNode(srv, parent, _name, true); _id != "" {
                    fmt.Printf("find Node %s -> %s\n", _name, _id)
                    parent = _id
                    continue
                }
                parent = createFolder(srv, parent, _name, false)
            }
            return parent
        }
        folderInfo := drive.File{
            Parents: []string{parent},
            Name: name,
            MimeType: "application/vnd.google-apps.folder",
        }
        driveFolder, _ := srv.Files.Create(&folderInfo).Do()
        return driveFolder.Id
}

func main() {

        if len(os.Args) > 3 || len(os.Args) == 1 {
            // go build gdrive-upload.go && GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account-key.json ./gdrive-upload
            log.Fatalf("Usage: %s <file-to-upload> [uploaded-folder-name]", os.Args[0])
        }
        targetFile := os.Args[1]
        targetFolder := ""
        if len(os.Args) == 3 {
            targetFolder = os.Args[2]
        }

        ctx := context.Background()
        driveService, err := drive.NewService(ctx)
        if err != nil {
            log.Fatalf("error %v", err)
        }

        goFile, err := os.Open(targetFile)
        if err != nil {
            log.Fatalf("error opening %q: %v", targetFile, err)
        }

        sharedFolderName := "Commandline Backup"
        r, err := driveService.Files.List().PageSize(1).
            Fields("nextPageToken, files(id, name)").Q("name = \"" + sharedFolderName + "\"").Do()
        if err != nil {
            log.Fatalf("Unable to retrieve files: %v", err)
        }

        sharedFolderId := r.Files[0].Id

        driveFolderId := ""
        if targetFolder != "" {
            driveFolderId = createFolder(driveService, sharedFolderId, targetFolder, true)
        }else {
            driveFolderId = sharedFolderId
        }

        // upload file
        fileinfo := drive.File{
            Parents: []string{driveFolderId}, // 多個parent Id => 可一次上傳到多個資料夾
            Name: filepath.Base(targetFile),
        }
        driveFile, err := driveService.Files.Create(&fileinfo).Media(goFile).Do()
        log.Printf("Got drive.File(%#v), Error(%v)", driveFile, err)

}

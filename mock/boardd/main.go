package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/ptt/pttweb/article"
	apipb "github.com/ptt/pttweb/proto/api"
	"golang.org/x/net/html"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 3000, "Port to listen on")
)

type server struct {
	apipb.UnimplementedBoardServiceServer
}

func (s *server) Hotboard(context.Context, *apipb.HotboardRequest) (*apipb.HotboardReply, error) {
	return &apipb.HotboardReply{
		Boards: []*apipb.Board{
			{
				Bid:        1,
				Name:       "Gossiping",
				Title:      "◎[老八] 共匪共諜就在本能寺",
				NumUsers:   8714,
				Bclass:     "綜合",
				Attributes: 0,
			},
			{
				Name:       "Stock",
				Title:      "◎[股票] 漲停板",
				NumUsers:   6031,
				Bclass:     "學術",
				Attributes: 0,
			},
			{
				Name:       "C_Chat",
				Title:      "◎[希洽] 發文時標題請防雷",
				NumUsers:   3792,
				Bclass:     "閒談",
				Attributes: 0,
			},
			{
				Name:       "Tech_Job",
				Title:      "◎[科技] 這裡是科技板",
				NumUsers:   378,
				Bclass:     "工作",
				Attributes: 0,
			},
			{
				Name:       "TY_Research",
				Title:      "◎強極渦劇場上映中",
				NumUsers:   27,
				Bclass:     "大氣",
				Attributes: 0,
			},
		},
	}, nil
}

func (s *server) Board(ctx context.Context, req *apipb.BoardRequest) (*apipb.BoardReply, error) {
	boards := make([]*apipb.Board, len(req.Ref))
	for i, ref := range req.Ref {
		boards[i] = &apipb.Board{
			Bid:        ref.GetBid(),
			Name:       "sysop",
			Title:      "Mock Board Title",
			NumUsers:   100,
			Bclass:     "Mock Class",
			Attributes: 0,
		}
	}
	return &apipb.BoardReply{Boards: boards}, nil
}

func (s *server) Content(ctx context.Context, req *apipb.ContentRequest) (*apipb.ContentReply, error) {
	b, err := getPttPage("/SYSOP/M.1742152964.A.660.html")
	if err != nil {
		return nil, err
	}
	return &apipb.ContentReply{
		Content: &apipb.Content{
			Content: handlePage(b),
		},
	}, nil

}
func handlePage(input []byte) []byte {
	buf := []byte{}

	doc, err := html.Parse(bytes.NewReader(input))
	if err != nil {
		slog.Error("parse error:", "err", err)
		return nil
	}
	// acceptMetaLines := true
	var mainContentNode *html.Node
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == "main-content" {
					mainContentNode = n
				}
			}
		}
	}

	if mainContentNode == nil {
		slog.Error("find main-content fail")
		return nil
	}
	state := "find-作者-line"
	for n := range mainContentNode.ChildNodes() {
		slog.Info("scan Descendants.Child", "node", n.Data, "attr", n.Attr)
		if n.Type == html.TextNode {
			buf = append(buf, []byte(n.Data)...)
			continue
		}

		if n.Type != html.ElementNode {
			slog.Error("unknown node", "node", n.Data, "type", n.Type)
			continue
		}

		if n.Data == "a" {
			buf = handleATag(buf, n)
			continue
		} else if n.Data == "span" {
			buf = handleSpanTag(buf, n)
			continue
		} else if n.Data == "div" {
			isRichcontent := false
			for _, attr := range n.Attr {
				if attr.Key == "class" {
					if attr.Val == "richcontent" {
						// skip richcontent
						isRichcontent = true
						break
					}
				}
			}
			if isRichcontent {
				// skip richcontent
				continue
			}
		}

		for _, attr := range n.Attr {

			if attr.Key == "class" && attr.Val == "push" {
				buf = handleDivPushTag(buf, n)
				break
			}

			if attr.Key == "class" && attr.Val == "article-metaline" {
				// slog.Info("found", "cn", cn)
				// buf = append(buf, []byte("state:" + state + "\n")...)
				if state == "find-作者-line" {

					// buf = append(buf, []byte("\n\n\n\n")...)
					state2 := "find-作者-tag"
					for cn2 := range n.ChildNodes() {

						for _, attr := range cn2.Attr {
							if attr.Key == "class" && attr.Val == "article-meta-tag" {
								if state2 == "find-作者-tag" {
									// 作者
									for cn3 := range cn2.ChildNodes() {
										// buf = append(buf, []byte("    " + cn3.Data + "[")...)
										// for _,attr := range cn3.Attr {
										// 	buf = append(buf, []byte(attr.Key + "=" + attr.Val)...)
										// }
										// buf = append(buf, []byte("]\n")...)
										if cn3.Data == "作者" {
											buf = append(buf, []byte("作者: ")...)
										}
									}
									state2 = "find-作者-value"
								}
							} else if attr.Key == "class" && attr.Val == "article-meta-value" {
								if state2 == "find-作者-value" {
									// 作者值
									for cn3 := range cn2.ChildNodes() {
										// buf = append(buf, []byte("    " + cn3.Data + "[")...)
										// for _,attr := range cn3.Attr {
										// 	buf = append(buf, []byte(attr.Key + "=" + attr.Val)...)
										// }

										buf = append(buf, []byte(cn3.Data+" ")...)
									}
									state2 = "find-看板-tag"
								}
							}
						}
					}

					state = "find-看板-line"
				} else if state == "find-Meta-line" {
					// 標題
					state2 := "find-tag"
					for cn2 := range n.ChildNodes() {
						for _, attr := range cn2.Attr {
							if attr.Key == "class" && attr.Val == "article-meta-tag" {
								if state2 == "find-tag" {
									// 標題
									for cn3 := range cn2.ChildNodes() {
										// if cn3.Data == "標題"{
										buf = append(buf, []byte(cn3.Data+": ")...)
										// }
									}
									state2 = "find-value"
								}
							} else if attr.Key == "class" && attr.Val == "article-meta-value" {
								if state2 == "find-value" {
									// 標題值
									for cn3 := range cn2.ChildNodes() {
										buf = append(buf, []byte(cn3.Data+"\n")...)
									}
								}
							}
						}
					}
					state = "find-Meta-line"
				}
			} else if attr.Key == "class" && attr.Val == "article-metaline-right" {
				if state == "find-看板-line" {
					// 看板
					state2 := "find-tag"
					for cn2 := range n.ChildNodes() {
						for _, attr := range cn2.Attr {
							if attr.Key == "class" && attr.Val == "article-meta-tag" {
								if state2 == "find-tag" {
									// 看板
									for cn3 := range cn2.ChildNodes() {
										if cn3.Data == "看板" {
											buf = append(buf, []byte("看板: ")...)
										}
									}
									state2 = "find-value"
								}
							} else if attr.Key == "class" && attr.Val == "article-meta-value" {
								if state2 == "find-value" {
									// 看板值
									for cn3 := range cn2.ChildNodes() {
										buf = append(buf, []byte(cn3.Data+"\n")...)
									}
								}
							}
						}
					}
					state = "find-Meta-line"
				}
			} else {
				slog.Info("unknown node", "node", n.Data)
				if state == "find-Meta-line" {
					// dont parse header any more
					// buf = append(buf, []byte(n.Data)...)
					// acceptMetaLines = false
				}
			}
		} // scan node attr
	}
	slog.Info("handleFin")

	return buf
}

func handleATag(buf []byte, n *html.Node) []byte {
	// should have more a in children
	for cn := range n.ChildNodes() {
		if cn.Type == html.TextNode {
			buf = append(buf, []byte(cn.Data)...)
		}
	}
	return buf

}

func handleSpanTag(buf []byte, n *html.Node) []byte {
	// should have more a in children
	class := ""
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			class = attr.Val
		}
	}
	classes := strings.Split(class, " ")
	t := article.TerminalState{}
	t.Reset()
	for _, c := range classes {
		fg, bg, flags := t.Fg(), t.Bg(), t.Flags()
		if strings.HasPrefix(c, article.ClassFgPrefix) { // `f`
			// foreground color
			v, err := strconv.Atoi(c[len(article.ClassBgPrefix):])
			if err != nil {
				slog.Warn("strconv.Atoi error", "err", err)
				continue
			}
			fg = 30 + v
		} else if strings.HasPrefix(c, article.ClassBgPrefix) { // `b`
			// background color
			v, err := strconv.Atoi(c[len(article.ClassBgPrefix):])
			if err != nil {
				slog.Warn("strconv.Atoi error", "err", err)
				continue
			}
			bg = 40 + v
		} else if c == article.ClassHighlight { // `hl`
			// highlighted
			flags |= article.Highlighted
		}else if c == article.ClassPushTag || 
		c == article.ClassPushUserId ||
		 c == article.ClassPushContent ||
		 c == article.ClassPushIpDatetime{
			continue
		} else { // `p`
			slog.Warn("unknown class", "class", c)
		}
		t.SetColor(fg, bg, flags)
	}
	for cn := range n.ChildNodes() {
		slog.Info("handleDivPushTag", "node", cn.Data, "attr", cn.Attr, "t", t)
		// escape := "\033"
		// if !t.IsDefaultState() {
		// slog.Info("handleDivPushTag not dfs", "node", cn.Data, "attr", cn.Attr, "t", t)
		escape := "\033["
		seg := []string{}
		if t.HasFlags(article.Highlighted) {
			seg = append(seg, "1")
		} else {
			seg = append(seg, "0")
		}
		if t.Fg() != article.DefaultFg {
			seg = append(seg, strconv.Itoa(t.Fg()))
		}
		if t.Bg() != article.DefaultBg {
			seg = append(seg, strconv.Itoa(t.Bg()))
		}
		escape += strings.Join(seg, ";") + "m"
		buf = append(buf, []byte(escape)...)
		// }
		if cn.Type == html.TextNode {
		buf = append(buf, []byte(cn.Data)...)
		}else {
			if cn.Data == "a" {
				buf = append(buf, []byte("\033[m")...)
				buf = handleATag(buf, cn)
				// buf = append(buf, []byte("\n")...)
			}else{
			slog.Warn("handleDivPushTag unknown tag in span", "node", cn.Data)
			}
		}
		if !t.IsDefaultState() {
			buf = append(buf, []byte("\033[m")...)
		}
	}
	return buf

}

func handleDivPushTag(buf []byte, n *html.Node) []byte {

	for cn := range n.ChildNodes() {
		if cn.Data == "span" {
			buf = handleSpanTag(buf, cn)
			continue
		}
		if cn.Type == html.ElementNode {
			slog.Warn("handleDivPushTag unknown node", "node", cn.Data, "type", cn.Type)
			for cn2 := range cn.ChildNodes() {

				slog.Info("handleDivPushTag", "node", cn2.Data, "attr", cn2.Attr)
				buf = append(buf, []byte(cn2.Data)...)
			}
		}
	}
	return buf
}

func handlePage2(input []byte) []byte {
	buf := []byte{}

	z := html.NewTokenizer(bytes.NewReader(input))
	for {
		tt := z.Next()
		tok := z.Token()
		slog.Info("Token", "type", tt, "all", fmt.Sprintf("%+v", tt), "tok", tok)
		switch tt {
		case html.ErrorToken:
			slog.Error("error token", "token", tt, "all", fmt.Sprintf("%+v", tt))
			return buf
		}
	}
	slog.Info("handleFin")

	return buf
}

func getPttPage(url string) ([]byte, error) {
	requestURL := fmt.Sprintf("https://www.ptt.cc/bbs/%v", url)
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get page: %s", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	apipb.RegisterBoardServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

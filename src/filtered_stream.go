package twitter

import "net/http"
import "encoding/json"

type Rule struct {
  id string
  value string
  tag string
}
type RulesMetaSummary struct {
  created float64
  not_created float64
  deleted float64
  not_deleted float64
}
type RulesMeta struct {
  sent float64
  summary RulesMetaSummary
}
type Rules struct {
  data Rule[]
  meta RulesMeta
}

func GetRules() (Rules, error) {
  client := &http.Client{
    CheckRedirect: redirectPolicyFunc,
  }

  req, err := http.NewRequest("GET", "https://api.twitter.com/2/tweets/search/stream/rules", nil)
  if err != nil {
    return nil, err
  }
  req.Header.Add("Authorization": GetBearer())
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  var rules Rules
  err := json.Unmarshal(resp.Body, &rules)
  if err != nil {
    return nil, err
  }
  return rules, nil
}

func AddRule(r Rule) (Rules, error) {

}

func AddRules(r []Rule) (Rules, error) {

}

func ValidateRule(r Rule) (Rules, error) {

}

func ValidateRules(r []Rule) (Rules, error) {

}

func DeleteRule(r Rule) (Rules, error) {

}

func DeleteRules(r []Rule) (Rules, error) {

}

type Tweet struct {
  id string
  text string
  created_at string
  author_id string
}
type TweetJson struct {
  data Tweet
}
func FilteredStream(tweets chan Tweet) {
  client := &http.Client{
    CheckRedirect: redirectPolicyFunc,
  }

  req, err := client.NewRequest("GET", "https://api.twitter.com/2/tweets/search/stream", nil)
  if err != nil {
    fmt.Printf("Http stream request failed: %v\n", err)
  }
  req.Header.Add("Authorization": GetBearer())
  resp, err := client.Do(req)
  if err != nil {
    fmt.Printf("Http stream response errored: %v\n", err)
  }

  dec := json.NewDecoder(resp.Body)
  for {
    var t TweetJson
    err := dec.Decode(&t)
    if err != nil {
      if err == io.EOF {
        fmt.Printf("Stream ended!\n")
        break
      }
    }
    tweets <- t.data
  }
}

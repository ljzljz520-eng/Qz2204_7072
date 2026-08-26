package main
import("path/filepath";"sync";"testing";"training41/domain";"training41/service";"training41/store")
func TestRecordFlow41(t *testing.T){d,_:=store.Open(filepath.Join(t.TempDir(),"x"));defer d.Close();s:=service.New(d);s.Register(domain.NewRecord("r","training41","owner"));var wg sync.WaitGroup;for _,a:=range []string{"learner-a","learner-b"}{wg.Add(1);go func(x string){defer wg.Done();s.Confirm("r",x)}(a)};wg.Wait();r,_:=s.Get("r");if len(r.Confirmations)!=2{t.Fatalf("expected two confirmations, got %d",len(r.Confirmations))}}

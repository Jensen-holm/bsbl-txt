package bsbl

//
//func Play(gs []*Game) {
//	var wg = sync.WaitGroup{}
//	claimed := make(map[*Game]bool, len(gs))
//	rs := make(map[int]string, len(gs))
//
//	for _, g := range gs {
//		wg.Add(1)
//		go func(wg *sync.WaitGroup, game *Game) {
//
//			defer wg.Done()
//			// do not think this is done right
//			if _, isIn := claimed[g]; isIn {
//				return
//			}
//
//		}(&wg, g)
//	}
//}

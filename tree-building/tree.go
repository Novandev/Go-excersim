package tree

import (
	"sort" // need the sorting algorithm
	"fmt"
)



//Record is a database item that specifies a parent record in a tree structure.
type Record struct {
	ID, Parent int
}

//Interface to turn a record into a string format
func (r *Record) String() string {
	return fmt.Sprintf("(I.D.: %d,Parent Node: %d)", r.ID, r.Parent)
}

//Interfaces for storing records in the tree.
type byID []Record

func (ids byID) Len() int           { return len(ids) }
func (ids byID) Swap(i, j int)      { ids[i], ids[j] = ids[j], ids[i] }
func (ids byID) Less(i, j int) bool { return ids[i].ID < ids[j].ID }

//A NIDE IN THE TREEE STRUCTURE if it has 0 children then it is a leaf, should pribably add this later
type Node struct {
	ID       int
	Children []*Node
}

// This should convert a list of records into a prent child structure
func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}
	sort.Sort(byID(records))
	nodes := make([]*Node, len(records))
	for r, rec := range records {
		nodes[r] = &Node{ID: rec.ID}
		if r == 0 && (rec.ID != 0 || rec.Parent != 0) { // If there isnt a record or the record id is invalie
			return nil, fmt.Errorf("Invalid root%s",rec.String())
		} else if r == 0 { // if its just r thats zero contine
			continue

			//checks to see if there is an id or a parent
		} else if r != rec.ID || rec.ID <= rec.Parent {
			return nil, fmt.Errorf("Record not valie:%s",rec.String())
		}
		// creeps down the tree
		if parent := nodes[rec.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[r])
		}
	}
	return nodes[0], nil
}
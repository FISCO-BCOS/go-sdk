package types

type Node struct {
	Agency    string   `json:"Agency"`
	IPAndPort string   `json:"IPAndPort"`
	Node      string   `json:"Node"`
	NodeID    string   `json:"NodeID"`
	Topic     []string `json:"Topic"`
}

// GetAgency returns the node agency string
func (n *Node) GetAgency() string {
	return n.Agency
}

// GetIPAndPort returns the node ip and port string
func (n *Node) GetIPAndPort() string {
	return n.IPAndPort
}

// GetNode returns the node string
func (n *Node) GetNode() string {
	return n.Node
}

// GetNodeID returns the node id string
func (n *Node) GetNodeID() string {
	return n.NodeID
}

// GetTopic returns the node attention topic infomation string
func (n *Node) GetTopic() []string {
	return n.Topic
}

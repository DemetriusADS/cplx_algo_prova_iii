package monitor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/2/machine"
)

type Monitor struct {
	machines  []*machine.Machine
	mChannel  chan machine.MetricDTO
	mqChannel chan bool
}

func NewMonitor(machines []*machine.Machine, mChannel chan machine.MetricDTO, mqChannel chan bool) *Monitor {
	return &Monitor{
		machines:  machines,
		mChannel:  mChannel,
		mqChannel: mqChannel,
	}
}

//Complexidade O(1) ou O(n^2)
// Mais uma vez, o for sem condicao de parada, torna a complexidade linear para o primeiro case.
// Para o segundo case, a complexidade é O(n^2), pois o for é executado n vezes, sendo n a quantidade de máquinas e em seguida o for é executado n vezes, sendo n a quantidade de métricas.

func (m *Monitor) Start() {
	fmt.Printf("Iniciando monitoramento\n")
	machinesOff := 0
	for {
		select {
		case msg := <-m.mChannel:
			fmt.Printf("Recebendo dados da %s\n. Dados cryptografados: %v\n\n", msg.MachineName, msg.Metrics)
			decryptedData, err := m.Decrypt(msg)
			if err != nil {
				fmt.Printf("Erro ao descriptografar dados: %s\n", err.Error())
				continue
			}
			var metric machine.Metric
			err = json.Unmarshal(decryptedData, &metric)
			if err != nil {
				fmt.Printf("Erro ao converter dados: %s\n", err.Error())
				continue
			}
			fmt.Printf("Temperatura: %f\n Volume: %f\n Leitura Estável: %t\n", metric.Temperature.Value, metric.Volume.Value, !metric.Unstable)

		case <-m.mqChannel:
			machinesOff++
			if machinesOff == len(m.machines) {
				fmt.Printf("Todas as maquinas estão desligadas\n")
				fmt.Printf("------- Realizando leitura para o desligamento do monitoramento -------\n")
				for _, machine := range m.machines {
					fmt.Printf("Lendo metricas da %s\n", machine.Name)
					metrics := machine.Read()
					for _, metric := range metrics {
						fmt.Printf("Temperatura: %f\n Volume: %f\n Leitura Estável: %t\n", metric.Temperature.Value, metric.Volume.Value, !metric.Unstable)
					}
				}
				return
			}
		}
	}
}

func (m *Monitor) Decrypt(data machine.MetricDTO) ([]byte, error) {
	fmt.Printf("Descriptografando dados\n")
	block, err := aes.NewCipher(data.Key)
	if err != nil {
		return nil, err
	}

	// Extract IV from the encrypted data
	iv := data.Metrics[:aes.BlockSize]
	encryptedData := data.Metrics[aes.BlockSize:]

	// Create a new AES cipher block mode
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the data
	decryptedData := make([]byte, len(encryptedData))
	mode.CryptBlocks(decryptedData, encryptedData)

	// Remove padding from the decrypted data
	padding := decryptedData[len(decryptedData)-1]
	decryptedData = decryptedData[:len(decryptedData)-int(padding)]

	return decryptedData, nil
}

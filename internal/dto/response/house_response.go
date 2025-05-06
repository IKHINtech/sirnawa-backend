package response

import "github.com/IKHINtech/sirnawa-backend/internal/models"

type HouseResponse struct {
	BaseResponse
	RtID          string `json:"rt_id"`
	RwID          string `json:"rw_id"`
	HousingAreaID string `json:"housing_area_id"`
	BlockID       string `json:"block_id"`
	Number        string `json:"number"`
	Status        string `json:"status"`
}

type HouseResponseDetail struct {
	HouseResponse
	Rt                    *RtResponse             `json:"rt"`
	Rw                    *RwResponse             `json:"rw"`
	Block                 *BlockResponse          `json:"block"`
	HousingArea           *HousingAreaResponse    `json:"housing_area"`
	ResidentHouseResponse []ResidentHouseResponse `json:"resident_houses"`
}

type HouseResponses []HouseResponse

func HouseModelToHouseResponse(data *models.House) *HouseResponse {
	base := BaseResponse{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &HouseResponse{
		Number:        data.Number,
		RwID:          data.RwID,
		HousingAreaID: data.HousingAreaID,
		BlockID:       data.BlockID,
		RtID:          data.RtID,
		Status:        string(data.Status),
		BaseResponse:  base,
	}
}

func HouseListToResponse(data models.Houses) HouseResponses {
	var res HouseResponses
	for _, v := range data {
		res = append(res, *HouseModelToHouseResponse(&v))
	}
	return res
}

func MapHouseDetailResponse(data *models.House) *HouseResponseDetail {
	if data == nil {
		return nil
	}

	residentHouse := make([]ResidentHouseResponse, len(data.ResidentHouses))
	for i, r := range data.ResidentHouses {
		residentHouse[i] = *ResidentHouseModelToResidentHouseResponse(&r)
	}

	var rtResponse *RtResponse
	var rwResponse *RwResponse
	var housingAreaResponse *HousingAreaResponse

	if data.Rt.ID != "" {
		rtResponse = RtModelToRtResponse(&data.Rt)
	}

	if data.Rw.ID != "" {
		rwResponse = RwModelToRwResponse(&data.Rw)
	}

	if data.HousingArea.ID != "" {
		housingAreaResponse = HousingAreaModelToHousingAreaResponse(&data.HousingArea)
	}

	return &HouseResponseDetail{
		HouseResponse:         *HouseModelToHouseResponse(data),
		Rt:                    rtResponse,
		Rw:                    rwResponse,
		Block:                 BlockModelToBlockResponse(&data.Block),
		HousingArea:           housingAreaResponse,
		ResidentHouseResponse: residentHouse,
	}
}

func HouseDetailListResponse(data models.Houses) []HouseResponseDetail {
	res := make([]HouseResponseDetail, len(data))
	for i, v := range data {
		res[i] = *MapHouseDetailResponse(&v)
	}
	return res
}

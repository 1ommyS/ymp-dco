package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/domain"
	"github.com/itpark/market/dco/internal/infrastructure/repository/segment"
	"github.com/itpark/market/dco/internal/presentation/http/segments/dto"
	"github.com/itpark/market/dco/internal/telemetry/logging"
	"math/rand"
	"slices"
)

type SegmentService struct {
	Repository *segment.SegmentRepository
}

func NewSegmentService(repository *segment.SegmentRepository) *SegmentService {
	return &SegmentService{
		Repository: repository,
	}
}

func (s SegmentService) CreateSegmentDto(ctx context.Context, dto dto.CreateSegmentDto) error {
	model := dto.ToModel()

	return s.Repository.CreateSegment(ctx, model)
}

func (s SegmentService) GetAll(ctx *gin.Context) ([]dto.GetSegmentsDto, error) {
	segments, err := s.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewGetSegmentDtoListFromModel(segments), nil
}

func (s SegmentService) GetSegmentByClientIdAndGroupTitle(ctx context.Context, clientId, groupTitle string) (domain.Segments, error) {
	segment, err := s.Repository.GetSegmentByClientIdAndGroupTitle(ctx, clientId, groupTitle)

	if err != nil {
		return domain.Segments{}, err
	}

	return segment, nil
}

// AttachUserToSegment привязывает клиента к одному из сегментов указанной группы.
// Выбор сегмента производится случайно с учётом «вероятностного веса» каждого сегмента.
//
// Алгоритм работы:
//  1. Получаем все сегменты группы по её названию.
//  2. Сортируем сегменты по возрастанию P — это не влияет на вероятность,
//     но гарантирует воспроизводимость порядка обхода.
//  3. Для каждого сегмента вычисляем вес wᵢ = 1 / Pᵢ.
//     — Чем меньше P, тем больше w и тем выше шанс выбора.
//     — Если Pᵢ ≤ 0, сегмент пропускается (его вес = 0).
//  4. Суммируем все веса: W = Σ wᵢ.
//  5. Генерируем случайное число r ∈ [0, W).
//  6. По очереди накапливаем веса: S = Σ wⱼ, пока S ≥ r — это выбранный сегмент.
//  7. Вызываем репозиторий для привязки пользователя к выбранному сегменту.
//
// Параметры:
//
//	ctx       Контекст gin, используется для передачи в репозиторий и логирования.
//	groupTitle Название группы сегментов.
//	clientId  Идентификатор клиента, которого нужно привязать.
//
// Возвращает:
//
//	(bool, error)
//	  bool  — true, если пользователь успешно привязан к сегменту.
//	  error — описание ошибки, если что-то пошло не так.
//
// Пример расчёта вероятностей:
//
//	Пусть есть три сегмента с P = {1, 5, 10}.
//	Тогда веса w = {1/1, 1/5, 1/10} = {1.0, 0.2, 0.1}.
//	Суммарный вес W = 1.3.
//	Генерируем r ∈ [0, 1.3).
//	— Если r < 1.0 → выбираем сегмент с P=1  (≈77% шанс).
//	— Если 1.0 ≤ r < 1.2 → выбираем P=5       (≈15% шанс).
//	— Если 1.2 ≤ r < 1.3 → выбираем P=10      (≈8%  шанс).
func (s SegmentService) AttachUserToSegment(ctx *gin.Context, groupTitle, clientId string) (bool, error) {
	allSegments, err := s.Repository.GetSegmentsByGroupTitle(ctx, groupTitle)
	if err != nil {
		logging.Error("Ошибка получения сегментов группы:", err)
		return false, err
	}

	slices.SortFunc(allSegments, func(a, b domain.Segments) int {
		return int(a.P - b.P)
	})

	weights := make([]float64, len(allSegments))
	var totalWeight float64
	for i, seg := range allSegments {
		if seg.P <= 0 {
			logging.Warn("Сегмент %s имеет некорректное P=%d, пропускаем", seg.ID, seg.P)
			weights[i] = 0
			continue
		}
		w := 1.0 / float64(seg.P)
		weights[i] = w
		totalWeight += w
	}

	if totalWeight == 0 {
		logging.Error("Суммарный вес всех сегментов равен нулю")
		return false, fmt.Errorf("нет доступных сегментов для привязки")
	}

	r := rand.Float64() * totalWeight

	var selected domain.Segments
	var cumulative float64
	for idx, seg := range allSegments {
		cumulative += weights[idx]
		if r <= cumulative {
			selected = seg
			break
		}
	}

	err = s.Repository.AttachUserToSegment(ctx, clientId, selected.ID)
	if err != nil {
		logging.Error("Не удалось привязать сегмент:", err)
		return false, err
	}

	return true, nil
}

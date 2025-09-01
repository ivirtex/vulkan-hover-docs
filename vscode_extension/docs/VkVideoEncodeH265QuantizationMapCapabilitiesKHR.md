# VkVideoEncodeH265QuantizationMapCapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeH265QuantizationMapCapabilitiesKHR - Structure describing H.265 encode quantization map capabilities



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) to query the capabilities of an [H.265 encode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-profile), the `VkVideoEncodeH265QuantizationMapCapabilitiesKHR` structure **can** be included in the `pNext` chain of the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure to retrieve additional video encode quantization map capabilities specific to H.265 encode profiles.

The `VkVideoEncodeH265QuantizationMapCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265 with VK_KHR_video_encode_quantization_map
typedef struct VkVideoEncodeH265QuantizationMapCapabilitiesKHR {
    VkStructureType    sType;
    void*              pNext;
    int32_t            minQpDelta;
    int32_t            maxQpDelta;
} VkVideoEncodeH265QuantizationMapCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `minQpDelta` indicates the minimum QP delta value supported for [H.265 QP delta maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-qp-delta-map).
- `maxQpDelta` indicates the maximum QP delta value supported for [H.265 QP delta maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-qp-delta-map).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265QuantizationMapCapabilitiesKHR-sType-sType)VUID-VkVideoEncodeH265QuantizationMapCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_QUANTIZATION_MAP_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265QuantizationMapCapabilitiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
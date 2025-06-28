# VkVideoEncodeH264QuantizationMapCapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeH264QuantizationMapCapabilitiesKHR - Structure describing H.264 encode quantization map capabilities



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) to query the capabilities of an [H.264 encode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-profile), the `VkVideoEncodeH264QuantizationMapCapabilitiesKHR` structure **can** be included in the `pNext` chain of the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure to retrieve additional video encode quantization map capabilities specific to H.264 encode profiles.

The `VkVideoEncodeH264QuantizationMapCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h264 with VK_KHR_video_encode_quantization_map
typedef struct VkVideoEncodeH264QuantizationMapCapabilitiesKHR {
    VkStructureType    sType;
    void*              pNext;
    int32_t            minQpDelta;
    int32_t            maxQpDelta;
} VkVideoEncodeH264QuantizationMapCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `minQpDelta` indicates the minimum QP delta value supported for [H.264 QP delta maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-qp-delta-map).
- `maxQpDelta` indicates the maximum QP delta value supported for [H.264 QP delta maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-qp-delta-map).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH264QuantizationMapCapabilitiesKHR-sType-sType)VUID-VkVideoEncodeH264QuantizationMapCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_QUANTIZATION_MAP_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html), [VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH264QuantizationMapCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
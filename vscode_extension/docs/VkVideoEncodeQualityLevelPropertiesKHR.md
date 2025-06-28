# VkVideoEncodeQualityLevelPropertiesKHR(3) Manual Page

## Name

VkVideoEncodeQualityLevelPropertiesKHR - Structure describing the video encode quality level properties



## [](#_c_specification)C Specification

The `VkVideoEncodeQualityLevelPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeQualityLevelPropertiesKHR {
    VkStructureType                            sType;
    void*                                      pNext;
    VkVideoEncodeRateControlModeFlagBitsKHR    preferredRateControlMode;
    uint32_t                                   preferredRateControlLayerCount;
} VkVideoEncodeQualityLevelPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `preferredRateControlMode` is a [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html) value indicating the preferred [rate control mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-modes) to use with the video encode quality level.
- `preferredRateControlLayerCount` indicates the preferred number of [rate control layers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-layers) to use with the video encode quality level.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-sType)VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUALITY_LEVEL_PROPERTIES_KHR`
- [](#VUID-VkVideoEncodeQualityLevelPropertiesKHR-pNext-pNext)VUID-VkVideoEncodeQualityLevelPropertiesKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoEncodeAV1QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QualityLevelPropertiesKHR.html), [VkVideoEncodeH264QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QualityLevelPropertiesKHR.html), or [VkVideoEncodeH265QualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QualityLevelPropertiesKHR.html)
- [](#VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-unique)VUID-VkVideoEncodeQualityLevelPropertiesKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html), [vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeQualityLevelPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
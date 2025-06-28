# VkVideoEncodeQualityLevelInfoKHR(3) Manual Page

## Name

VkVideoEncodeQualityLevelInfoKHR - Structure specifying used video encode quality level



## [](#_c_specification)C Specification

The `VkVideoEncodeQualityLevelInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeQualityLevelInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           qualityLevel;
} VkVideoEncodeQualityLevelInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `qualityLevel` is the used video encode quality level.

## [](#_description)Description

This structure **can** be specified in the following places:

- In the `pNext` chain of [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html) to specify the video encode quality level to use for a video session parameters object created for a video encode session. If no instance of this structure is included in the `pNext` chain of [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html), then the video session parameters object is created with a video encode quality level of zero.
- In the `pNext` chain of [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCodingControlInfoKHR.html) to change the video encode quality level state of the bound video session.

Valid Usage

- [](#VUID-VkVideoEncodeQualityLevelInfoKHR-qualityLevel-08311)VUID-VkVideoEncodeQualityLevelInfoKHR-qualityLevel-08311  
  `qualityLevel` **must** be less than [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html)::`maxQualityLevels`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used video profile

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeQualityLevelInfoKHR-sType-sType)VUID-VkVideoEncodeQualityLevelInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUALITY_LEVEL_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeQualityLevelInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
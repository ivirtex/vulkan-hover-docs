# VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR - Structure describing whether the implementation provides unified image layouts



## [](#_c_specification)C Specification

The `VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_unified_image_layouts
typedef struct VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           unifiedImageLayouts;
    VkBool32           unifiedImageLayoutsVideo;
} VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`unifiedImageLayouts` specifies whether usage of `VK_IMAGE_LAYOUT_GENERAL`, where valid, incurs no loss in efficiency. Additionally, it indicates whether it **can** be used in place of `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`.
- []()`unifiedImageLayoutsVideo` specifies whether `VK_IMAGE_LAYOUT_GENERAL` can be used in place of any of the following image layouts with no loss in efficiency.
  
  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR`
  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_QUANTIZATION_MAP_KHR`

## [](#_description)Description

If the `VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFIED_IMAGE_LAYOUTS_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_unified\_image\_layouts](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_unified_image_layouts.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
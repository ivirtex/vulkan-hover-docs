# VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR - Structure describing the video encode quantization map features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_quantization_map
typedef struct VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           videoEncodeQuantizationMap;
} VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`videoEncodeQuantizationMap` indicates that the implementation supports [video encode quantization maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map).
  
  Note
  
  Support for `videoEncodeQuantizationMap` does not indicate that all video encode profiles support quantization maps. Support for quantization maps for any specific video encode profile is subject to video-profile-specific capabilities.

## [](#_description)Description

If the `VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_ENCODE_QUANTIZATION_MAP_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
# VkPhysicalDeviceVideoDecodeVP9FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceVideoDecodeVP9FeaturesKHR - Structure describing the video decode VP9 features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVideoDecodeVP9FeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_vp9
typedef struct VkPhysicalDeviceVideoDecodeVP9FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           videoDecodeVP9;
} VkPhysicalDeviceVideoDecodeVP9FeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`videoDecodeVP9` specifies that the implementation supports [VP9 decode operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-vp9).

## [](#_description)Description

If the `VkPhysicalDeviceVideoDecodeVP9FeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceVideoDecodeVP9FeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVideoDecodeVP9FeaturesKHR-sType-sType)VUID-VkPhysicalDeviceVideoDecodeVP9FeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_DECODE_VP9_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_vp9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_vp9.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVideoDecodeVP9FeaturesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
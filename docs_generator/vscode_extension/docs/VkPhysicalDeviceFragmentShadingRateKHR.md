# VkPhysicalDeviceFragmentShadingRateKHR(3) Manual Page

## Name

VkPhysicalDeviceFragmentShadingRateKHR - Structure returning information about sample count specific additional multisampling capabilities



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFragmentShadingRateKHR` structure is defined as

```c++
// Provided by VK_KHR_fragment_shading_rate
typedef struct VkPhysicalDeviceFragmentShadingRateKHR {
    VkStructureType       sType;
    void*                 pNext;
    VkSampleCountFlags    sampleCounts;
    VkExtent2D            fragmentSize;
} VkPhysicalDeviceFragmentShadingRateKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `sampleCounts` is a bitmask of sample counts for which the shading rate described by `fragmentSize` is supported.
- `fragmentSize` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) describing the width and height of a supported shading rate.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFragmentShadingRateKHR-sType-sType)VUID-VkPhysicalDeviceFragmentShadingRateKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_KHR`
- [](#VUID-VkPhysicalDeviceFragmentShadingRateKHR-pNext-pNext)VUID-VkPhysicalDeviceFragmentShadingRateKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceFragmentShadingRatesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFragmentShadingRatesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFragmentShadingRateKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
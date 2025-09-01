# VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV - Structure describing fragment shading rate limits that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_fragment_shading_rate_enums
typedef struct VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV {
    VkStructureType          sType;
    void*                    pNext;
    VkSampleCountFlagBits    maxFragmentShadingRateInvocationCount;
} VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxFragmentShadingRateInvocationCount` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value indicating the maximum number of fragment shader invocations per fragment supported in pipeline, primitive, and attachment fragment shading rates.

## [](#_description)Description

If the `VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These properties are related to [fragment shading rates](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV-sType-sType)VUID-VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADING_RATE_ENUMS_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_fragment\_shading\_rate\_enums](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_fragment_shading_rate_enums.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
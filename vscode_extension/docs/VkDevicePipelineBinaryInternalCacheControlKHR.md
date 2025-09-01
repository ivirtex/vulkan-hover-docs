# VkDevicePipelineBinaryInternalCacheControlKHR(3) Manual Page

## Name

VkDevicePipelineBinaryInternalCacheControlKHR - Structure specifying parameter to disable the internal pipeline cache



## [](#_c_specification)C Specification

To disable the implementation’s internal pipeline cache, add a [VkDevicePipelineBinaryInternalCacheControlKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevicePipelineBinaryInternalCacheControlKHR.html) structure to the `pNext` chain of the [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) structure.

```c++
// Provided by VK_KHR_pipeline_binary
typedef struct VkDevicePipelineBinaryInternalCacheControlKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           disableInternalCache;
} VkDevicePipelineBinaryInternalCacheControlKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `disableInternalCache` specifies whether or not to disable the implementation’s internal pipeline cache.

## [](#_description)Description

If the `VkDeviceCreateInfo`::`pNext` chain does not include this structure, then `disableInternalCache` defaults to `VK_FALSE`.

Valid Usage

- [](#VUID-VkDevicePipelineBinaryInternalCacheControlKHR-disableInternalCache-09602)VUID-VkDevicePipelineBinaryInternalCacheControlKHR-disableInternalCache-09602  
  If [VkPhysicalDevicePipelineBinaryPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineBinaryPropertiesKHR.html)::`pipelineBinaryInternalCacheControl` is `VK_FALSE`, `disableInternalCache` **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkDevicePipelineBinaryInternalCacheControlKHR-sType-sType)VUID-VkDevicePipelineBinaryInternalCacheControlKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_PIPELINE_BINARY_INTERNAL_CACHE_CONTROL_KHR`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDevicePipelineBinaryInternalCacheControlKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
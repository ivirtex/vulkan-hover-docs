# VkQueueFamilyQueryResultStatusPropertiesKHR(3) Manual Page

## Name

VkQueueFamilyQueryResultStatusPropertiesKHR - Structure specifying support for result status query



## [](#_c_specification)C Specification

The [VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html) structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkQueueFamilyQueryResultStatusPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           queryResultStatusSupport;
} VkQueueFamilyQueryResultStatusPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `queryResultStatusSupport` reports `VK_TRUE` if query type `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` and use of `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` are supported.

## [](#_description)Description

If this structure is included in the `pNext` chain of the [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html) structure passed to [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html), then it is filled with information about whether [result status queries](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-result-status-only) are supported by the specified queue family.

Valid Usage (Implicit)

- [](#VUID-VkQueueFamilyQueryResultStatusPropertiesKHR-sType-sType)VUID-VkQueueFamilyQueryResultStatusPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUEUE_FAMILY_QUERY_RESULT_STATUS_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFamilyQueryResultStatusPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
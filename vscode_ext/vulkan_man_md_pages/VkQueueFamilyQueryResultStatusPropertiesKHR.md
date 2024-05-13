# VkQueueFamilyQueryResultStatusPropertiesKHR(3) Manual Page

## Name

VkQueueFamilyQueryResultStatusPropertiesKHR - Structure specifying
support for result status query



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkQueueFamilyQueryResultStatusPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyQueryResultStatusPropertiesKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkQueueFamilyQueryResultStatusPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           queryResultStatusSupport;
} VkQueueFamilyQueryResultStatusPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `queryResultStatusSupport` reports `VK_TRUE` if query type
  `VK_QUERY_TYPE_RESULT_STATUS_ONLY_KHR` and use of
  `VK_QUERY_RESULT_WITH_STATUS_BIT_KHR` are supported.

## <a href="#_description" class="anchor"></a>Description

If this structure is included in the `pNext` chain of the
[VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html) structure
passed to
[vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html),
then it is filled with information about whether <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-result-status-only"
target="_blank" rel="noopener">result status queries</a> are supported
by the specified queue family.

Valid Usage (Implicit)

- <a href="#VUID-VkQueueFamilyQueryResultStatusPropertiesKHR-sType-sType"
  id="VUID-VkQueueFamilyQueryResultStatusPropertiesKHR-sType-sType"></a>
  VUID-VkQueueFamilyQueryResultStatusPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_QUEUE_FAMILY_QUERY_RESULT_STATUS_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueueFamilyQueryResultStatusPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

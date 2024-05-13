# VkDeviceQueueGlobalPriorityCreateInfoKHR(3) Manual Page

## Name

VkDeviceQueueGlobalPriorityCreateInfoKHR - Specify a system wide
priority



## <a href="#_c_specification" class="anchor"></a>C Specification

Queues **can** be created with a system-wide priority by adding a
`VkDeviceQueueGlobalPriorityCreateInfoKHR` structure to the `pNext`
chain of [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html).

The `VkDeviceQueueGlobalPriorityCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_global_priority
typedef struct VkDeviceQueueGlobalPriorityCreateInfoKHR {
    VkStructureType             sType;
    const void*                 pNext;
    VkQueueGlobalPriorityKHR    globalPriority;
} VkDeviceQueueGlobalPriorityCreateInfoKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_global_priority
typedef VkDeviceQueueGlobalPriorityCreateInfoKHR VkDeviceQueueGlobalPriorityCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `globalPriority` is the system-wide priority associated to these
  queues as specified by
  [VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityKHR.html)

## <a href="#_description" class="anchor"></a>Description

Queues created without specifying
`VkDeviceQueueGlobalPriorityCreateInfoKHR` will default to
`VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_KHR`.

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceQueueGlobalPriorityCreateInfoKHR-sType-sType"
  id="VUID-VkDeviceQueueGlobalPriorityCreateInfoKHR-sType-sType"></a>
  VUID-VkDeviceQueueGlobalPriorityCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO_KHR`

- <a
  href="#VUID-VkDeviceQueueGlobalPriorityCreateInfoKHR-globalPriority-parameter"
  id="VUID-VkDeviceQueueGlobalPriorityCreateInfoKHR-globalPriority-parameter"></a>
  VUID-VkDeviceQueueGlobalPriorityCreateInfoKHR-globalPriority-parameter  
  `globalPriority` **must** be a valid
  [VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityKHR.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_global_priority.html),
[VK_KHR_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_global_priority.html),
[VkQueueGlobalPriorityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueGlobalPriorityKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceQueueGlobalPriorityCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

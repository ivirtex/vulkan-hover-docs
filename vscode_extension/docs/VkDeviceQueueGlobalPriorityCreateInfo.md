# VkDeviceQueueGlobalPriorityCreateInfo(3) Manual Page

## Name

VkDeviceQueueGlobalPriorityCreateInfo - Specify a system wide priority



## [](#_c_specification)C Specification

Queues **can** be created with a system-wide priority by adding a `VkDeviceQueueGlobalPriorityCreateInfo` structure to the `pNext` chain of [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html).

The `VkDeviceQueueGlobalPriorityCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkDeviceQueueGlobalPriorityCreateInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkQueueGlobalPriority    globalPriority;
} VkDeviceQueueGlobalPriorityCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_global_priority
typedef VkDeviceQueueGlobalPriorityCreateInfo VkDeviceQueueGlobalPriorityCreateInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_global_priority
typedef VkDeviceQueueGlobalPriorityCreateInfo VkDeviceQueueGlobalPriorityCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `globalPriority` is the system-wide priority associated to these queues as specified by [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html)

## [](#_description)Description

Queues created without specifying `VkDeviceQueueGlobalPriorityCreateInfo` will default to `VK_QUEUE_GLOBAL_PRIORITY_MEDIUM`.

Valid Usage (Implicit)

- [](#VUID-VkDeviceQueueGlobalPriorityCreateInfo-sType-sType)VUID-VkDeviceQueueGlobalPriorityCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_QUEUE_GLOBAL_PRIORITY_CREATE_INFO`
- [](#VUID-VkDeviceQueueGlobalPriorityCreateInfo-globalPriority-parameter)VUID-VkDeviceQueueGlobalPriorityCreateInfo-globalPriority-parameter  
  `globalPriority` **must** be a valid [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html) value

## [](#_see_also)See Also

[VK\_EXT\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_global_priority.html), [VK\_KHR\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_global_priority.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceQueueGlobalPriorityCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
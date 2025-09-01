# VkQueueFamilyGlobalPriorityProperties(3) Manual Page

## Name

VkQueueFamilyGlobalPriorityProperties - Return structure for queue family global priority information query



## [](#_c_specification)C Specification

The [VkQueueFamilyGlobalPriorityProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyGlobalPriorityProperties.html) structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkQueueFamilyGlobalPriorityProperties {
    VkStructureType          sType;
    void*                    pNext;
    uint32_t                 priorityCount;
    VkQueueGlobalPriority    priorities[VK_MAX_GLOBAL_PRIORITY_SIZE];
} VkQueueFamilyGlobalPriorityProperties;
```

or the equivalent

```c++
// Provided by VK_KHR_global_priority
typedef VkQueueFamilyGlobalPriorityProperties VkQueueFamilyGlobalPriorityPropertiesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_global_priority_query
typedef VkQueueFamilyGlobalPriorityProperties VkQueueFamilyGlobalPriorityPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `priorityCount` is the number of supported global queue priorities in this queue family, and it **must** be greater than 0.
- `priorities` is an array of `VK_MAX_GLOBAL_PRIORITY_SIZE` [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html) enums representing all supported global queue priorities in this queue family. The first `priorityCount` elements of the array will be valid.

## [](#_description)Description

If the `VkQueueFamilyGlobalPriorityProperties` structure is included in the `pNext` chain of the [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html) structure passed to [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html), it is filled in with the list of supported global queue priorities for the indicated family.

The valid elements of `priorities` **must** not contain any duplicate values.

The valid elements of `priorities` **must** be a continuous sequence of [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html) enums in the ascending order.

Note

For example, returning `priorityCount` as 3 with supported `priorities` as `VK_QUEUE_GLOBAL_PRIORITY_LOW`, `VK_QUEUE_GLOBAL_PRIORITY_MEDIUM` and `VK_QUEUE_GLOBAL_PRIORITY_REALTIME` is not allowed.

Valid Usage (Implicit)

- [](#VUID-VkQueueFamilyGlobalPriorityProperties-sType-sType)VUID-VkQueueFamilyGlobalPriorityProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUEUE_FAMILY_GLOBAL_PRIORITY_PROPERTIES`

## [](#_see_also)See Also

[VK\_EXT\_global\_priority\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_global_priority_query.html), [VK\_KHR\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_global_priority.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkQueueGlobalPriority](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueGlobalPriority.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFamilyGlobalPriorityProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
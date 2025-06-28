# VkQueueGlobalPriority(3) Manual Page

## Name

VkQueueGlobalPriority - Values specifying a system-wide queue priority



## [](#_c_specification)C Specification

Possible values of [VkDeviceQueueGlobalPriorityCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueGlobalPriorityCreateInfo.html)::`globalPriority`, specifying a system-wide priority level are:

```c++
// Provided by VK_VERSION_1_4
typedef enum VkQueueGlobalPriority {
    VK_QUEUE_GLOBAL_PRIORITY_LOW = 128,
    VK_QUEUE_GLOBAL_PRIORITY_MEDIUM = 256,
    VK_QUEUE_GLOBAL_PRIORITY_HIGH = 512,
    VK_QUEUE_GLOBAL_PRIORITY_REALTIME = 1024,
  // Provided by VK_EXT_global_priority
    VK_QUEUE_GLOBAL_PRIORITY_LOW_EXT = VK_QUEUE_GLOBAL_PRIORITY_LOW,
  // Provided by VK_EXT_global_priority
    VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT = VK_QUEUE_GLOBAL_PRIORITY_MEDIUM,
  // Provided by VK_EXT_global_priority
    VK_QUEUE_GLOBAL_PRIORITY_HIGH_EXT = VK_QUEUE_GLOBAL_PRIORITY_HIGH,
  // Provided by VK_EXT_global_priority
    VK_QUEUE_GLOBAL_PRIORITY_REALTIME_EXT = VK_QUEUE_GLOBAL_PRIORITY_REALTIME,
  // Provided by VK_KHR_global_priority
    VK_QUEUE_GLOBAL_PRIORITY_LOW_KHR = VK_QUEUE_GLOBAL_PRIORITY_LOW,
  // Provided by VK_KHR_global_priority
    VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_KHR = VK_QUEUE_GLOBAL_PRIORITY_MEDIUM,
  // Provided by VK_KHR_global_priority
    VK_QUEUE_GLOBAL_PRIORITY_HIGH_KHR = VK_QUEUE_GLOBAL_PRIORITY_HIGH,
  // Provided by VK_KHR_global_priority
    VK_QUEUE_GLOBAL_PRIORITY_REALTIME_KHR = VK_QUEUE_GLOBAL_PRIORITY_REALTIME,
} VkQueueGlobalPriority;
```

or the equivalent

```c++
// Provided by VK_KHR_global_priority
typedef VkQueueGlobalPriority VkQueueGlobalPriorityKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_global_priority
typedef VkQueueGlobalPriority VkQueueGlobalPriorityEXT;
```

## [](#_description)Description

Priority values are sorted in ascending order. A comparison operation on the enum values can be used to determine the priority order.

- `VK_QUEUE_GLOBAL_PRIORITY_LOW` is below the system default. Useful for non-interactive tasks.
- `VK_QUEUE_GLOBAL_PRIORITY_MEDIUM` is the system default priority.
- `VK_QUEUE_GLOBAL_PRIORITY_HIGH` is above the system default.
- `VK_QUEUE_GLOBAL_PRIORITY_REALTIME` is the highest priority. Useful for critical tasks.

## [](#_see_also)See Also

[VK\_EXT\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_global_priority.html), [VK\_KHR\_global\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_global_priority.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDeviceQueueGlobalPriorityCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueGlobalPriorityCreateInfo.html), [VkQueueFamilyGlobalPriorityProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyGlobalPriorityProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueGlobalPriority)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
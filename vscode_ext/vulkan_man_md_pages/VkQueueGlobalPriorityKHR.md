# VkQueueGlobalPriorityKHR(3) Manual Page

## Name

VkQueueGlobalPriorityKHR - Values specifying a system-wide queue
priority



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkDeviceQueueGlobalPriorityCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueGlobalPriorityCreateInfoKHR.html)::`globalPriority`,
specifying a system-wide priority level are:

``` c
// Provided by VK_KHR_global_priority
typedef enum VkQueueGlobalPriorityKHR {
    VK_QUEUE_GLOBAL_PRIORITY_LOW_KHR = 128,
    VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_KHR = 256,
    VK_QUEUE_GLOBAL_PRIORITY_HIGH_KHR = 512,
    VK_QUEUE_GLOBAL_PRIORITY_REALTIME_KHR = 1024,
    VK_QUEUE_GLOBAL_PRIORITY_LOW_EXT = VK_QUEUE_GLOBAL_PRIORITY_LOW_KHR,
    VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_EXT = VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_KHR,
    VK_QUEUE_GLOBAL_PRIORITY_HIGH_EXT = VK_QUEUE_GLOBAL_PRIORITY_HIGH_KHR,
    VK_QUEUE_GLOBAL_PRIORITY_REALTIME_EXT = VK_QUEUE_GLOBAL_PRIORITY_REALTIME_KHR,
} VkQueueGlobalPriorityKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_global_priority
typedef VkQueueGlobalPriorityKHR VkQueueGlobalPriorityEXT;
```

## <a href="#_description" class="anchor"></a>Description

Priority values are sorted in ascending order. A comparison operation on
the enum values can be used to determine the priority order.

- `VK_QUEUE_GLOBAL_PRIORITY_LOW_KHR` is below the system default. Useful
  for non-interactive tasks.

- `VK_QUEUE_GLOBAL_PRIORITY_MEDIUM_KHR` is the system default priority.

- `VK_QUEUE_GLOBAL_PRIORITY_HIGH_KHR` is above the system default.

- `VK_QUEUE_GLOBAL_PRIORITY_REALTIME_KHR` is the highest priority.
  Useful for critical tasks.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_global_priority.html),
[VK_KHR_global_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_global_priority.html),
[VkDeviceQueueGlobalPriorityCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueGlobalPriorityCreateInfoKHR.html),
[VkQueueFamilyGlobalPriorityPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyGlobalPriorityPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueueGlobalPriorityKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# vkGetDeviceQueue2(3) Manual Page

## Name

vkGetDeviceQueue2 - Get a queue handle from a device



## [](#_c_specification)C Specification

To retrieve a handle to a [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object with specific [VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateFlags.html) creation flags, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetDeviceQueue2(
    VkDevice                                    device,
    const VkDeviceQueueInfo2*                   pQueueInfo,
    VkQueue*                                    pQueue);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the queue.
- `pQueueInfo` is a pointer to a [VkDeviceQueueInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueInfo2.html) structure, describing parameters of the device queue to be retrieved.
- `pQueue` is a pointer to a [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object that will be filled with the handle for the requested queue.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceQueue2-device-parameter)VUID-vkGetDeviceQueue2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceQueue2-pQueueInfo-parameter)VUID-vkGetDeviceQueue2-pQueueInfo-parameter  
  `pQueueInfo` **must** be a valid pointer to a valid [VkDeviceQueueInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueInfo2.html) structure
- [](#VUID-vkGetDeviceQueue2-pQueue-parameter)VUID-vkGetDeviceQueue2-pQueue-parameter  
  `pQueue` **must** be a valid pointer to a [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceQueueInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueInfo2.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceQueue2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
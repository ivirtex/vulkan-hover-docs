# vkGetDeviceQueue(3) Manual Page

## Name

vkGetDeviceQueue - Get a queue handle from a device



## [](#_c_specification)C Specification

To retrieve a handle to a [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object, call:

```c++
// Provided by VK_VERSION_1_0
void vkGetDeviceQueue(
    VkDevice                                    device,
    uint32_t                                    queueFamilyIndex,
    uint32_t                                    queueIndex,
    VkQueue*                                    pQueue);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the queue.
- `queueFamilyIndex` is the index of the queue family to which the queue belongs.
- `queueIndex` is the index within this queue family of the queue to retrieve.
- `pQueue` is a pointer to a [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) object that will be filled with the handle for the requested queue.

## [](#_description)Description

`vkGetDeviceQueue` **must** only be used to get queues that were created with the `flags` parameter of [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html) set to zero. To get queues that were created with a non-zero `flags` parameter use [vkGetDeviceQueue2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceQueue2.html).

Valid Usage

- [](#VUID-vkGetDeviceQueue-queueFamilyIndex-00384)VUID-vkGetDeviceQueue-queueFamilyIndex-00384  
  `queueFamilyIndex` **must** be one of the queue family indices specified when `device` was created, via the [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html) structure
- [](#VUID-vkGetDeviceQueue-queueIndex-00385)VUID-vkGetDeviceQueue-queueIndex-00385  
  `queueIndex` **must** be less than the value of [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html)::`queueCount` for the queue family indicated by `queueFamilyIndex` when `device` was created
- [](#VUID-vkGetDeviceQueue-flags-01841)VUID-vkGetDeviceQueue-flags-01841  
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateInfo.html)::`flags` **must** have been zero when `device` was created

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceQueue-device-parameter)VUID-vkGetDeviceQueue-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceQueue-pQueue-parameter)VUID-vkGetDeviceQueue-pQueue-parameter  
  `pQueue` **must** be a valid pointer to a [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceQueue)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0
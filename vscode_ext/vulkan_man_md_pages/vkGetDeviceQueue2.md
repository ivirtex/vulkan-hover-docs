# vkGetDeviceQueue2(3) Manual Page

## Name

vkGetDeviceQueue2 - Get a queue handle from a device



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve a handle to a [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object with specific
[VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlags.html) creation
flags, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetDeviceQueue2(
    VkDevice                                    device,
    const VkDeviceQueueInfo2*                   pQueueInfo,
    VkQueue*                                    pQueue);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the queue.

- `pQueueInfo` is a pointer to a
  [VkDeviceQueueInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueInfo2.html) structure, describing
  parameters of the device queue to be retrieved.

- `pQueue` is a pointer to a [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object that will be
  filled with the handle for the requested queue.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceQueue2-device-parameter"
  id="VUID-vkGetDeviceQueue2-device-parameter"></a>
  VUID-vkGetDeviceQueue2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceQueue2-pQueueInfo-parameter"
  id="VUID-vkGetDeviceQueue2-pQueueInfo-parameter"></a>
  VUID-vkGetDeviceQueue2-pQueueInfo-parameter  
  `pQueueInfo` **must** be a valid pointer to a valid
  [VkDeviceQueueInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueInfo2.html) structure

- <a href="#VUID-vkGetDeviceQueue2-pQueue-parameter"
  id="VUID-vkGetDeviceQueue2-pQueue-parameter"></a>
  VUID-vkGetDeviceQueue2-pQueue-parameter  
  `pQueue` **must** be a valid pointer to a [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceQueueInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueInfo2.html), [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceQueue2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

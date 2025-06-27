# vkGetDeviceQueue(3) Manual Page

## Name

vkGetDeviceQueue - Get a queue handle from a device



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve a handle to a [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetDeviceQueue(
    VkDevice                                    device,
    uint32_t                                    queueFamilyIndex,
    uint32_t                                    queueIndex,
    VkQueue*                                    pQueue);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the queue.

- `queueFamilyIndex` is the index of the queue family to which the queue
  belongs.

- `queueIndex` is the index within this queue family of the queue to
  retrieve.

- `pQueue` is a pointer to a [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) object that will be
  filled with the handle for the requested queue.

## <a href="#_description" class="anchor"></a>Description

`vkGetDeviceQueue` **must** only be used to get queues that were created
with the `flags` parameter of
[VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html) set to zero. To
get queues that were created with a non-zero `flags` parameter use
[vkGetDeviceQueue2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceQueue2.html).

Valid Usage

- <a href="#VUID-vkGetDeviceQueue-queueFamilyIndex-00384"
  id="VUID-vkGetDeviceQueue-queueFamilyIndex-00384"></a>
  VUID-vkGetDeviceQueue-queueFamilyIndex-00384  
  `queueFamilyIndex` **must** be one of the queue family indices
  specified when `device` was created, via the
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html) structure

- <a href="#VUID-vkGetDeviceQueue-queueIndex-00385"
  id="VUID-vkGetDeviceQueue-queueIndex-00385"></a>
  VUID-vkGetDeviceQueue-queueIndex-00385  
  `queueIndex` **must** be less than the value of
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html)::`queueCount`
  for the queue family indicated by `queueFamilyIndex` when `device` was
  created

- <a href="#VUID-vkGetDeviceQueue-flags-01841"
  id="VUID-vkGetDeviceQueue-flags-01841"></a>
  VUID-vkGetDeviceQueue-flags-01841  
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html)::`flags`
  **must** have been set to zero when `device` was created

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceQueue-device-parameter"
  id="VUID-vkGetDeviceQueue-device-parameter"></a>
  VUID-vkGetDeviceQueue-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceQueue-pQueue-parameter"
  id="VUID-vkGetDeviceQueue-pQueue-parameter"></a>
  VUID-vkGetDeviceQueue-pQueue-parameter  
  `pQueue` **must** be a valid pointer to a [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceQueue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

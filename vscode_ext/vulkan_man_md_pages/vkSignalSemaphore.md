# vkSignalSemaphore(3) Manual Page

## Name

vkSignalSemaphore - Signal a timeline semaphore on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To signal a semaphore created with a
[VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`
with a particular counter value, on the host, call:

``` c
// Provided by VK_VERSION_1_2
VkResult vkSignalSemaphore(
    VkDevice                                    device,
    const VkSemaphoreSignalInfo*                pSignalInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_timeline_semaphore
VkResult vkSignalSemaphoreKHR(
    VkDevice                                    device,
    const VkSemaphoreSignalInfo*                pSignalInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the semaphore.

- `pSignalInfo` is a pointer to a
  [VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSignalInfo.html) structure
  containing information about the signal operation.

## <a href="#_description" class="anchor"></a>Description

When `vkSignalSemaphore` is executed on the host, it defines and
immediately executes a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
target="_blank" rel="noopener"><em>semaphore signal operation</em></a>
which sets the timeline semaphore to the given value.

The first synchronization scope is defined by the host execution model,
but includes execution of `vkSignalSemaphore` on the host and anything
that happened-before it.

The second synchronization scope is empty.

Valid Usage (Implicit)

- <a href="#VUID-vkSignalSemaphore-device-parameter"
  id="VUID-vkSignalSemaphore-device-parameter"></a>
  VUID-vkSignalSemaphore-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSignalSemaphore-pSignalInfo-parameter"
  id="VUID-vkSignalSemaphore-pSignalInfo-parameter"></a>
  VUID-vkSignalSemaphore-pSignalInfo-parameter  
  `pSignalInfo` **must** be a valid pointer to a valid
  [VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSignalInfo.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSignalInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSignalSemaphore"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700

# vkGetSemaphoreCounterValue(3) Manual Page

## Name

vkGetSemaphoreCounterValue - Query the current state of a timeline
semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the current counter value of a semaphore created with a
[VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`
from the host, call:

``` c
// Provided by VK_VERSION_1_2
VkResult vkGetSemaphoreCounterValue(
    VkDevice                                    device,
    VkSemaphore                                 semaphore,
    uint64_t*                                   pValue);
```

or the equivalent command

``` c
// Provided by VK_KHR_timeline_semaphore
VkResult vkGetSemaphoreCounterValueKHR(
    VkDevice                                    device,
    VkSemaphore                                 semaphore,
    uint64_t*                                   pValue);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the semaphore.

- `semaphore` is the handle of the semaphore to query.

- `pValue` is a pointer to a 64-bit integer value in which the current
  counter value of the semaphore is returned.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-submission"
target="_blank" rel="noopener">queue submission</a> command is pending
execution, then the value returned by this command <strong>may</strong>
immediately be out of date.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkGetSemaphoreCounterValue-semaphore-03255"
  id="VUID-vkGetSemaphoreCounterValue-semaphore-03255"></a>
  VUID-vkGetSemaphoreCounterValue-semaphore-03255  
  `semaphore` **must** have been created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE`

Valid Usage (Implicit)

- <a href="#VUID-vkGetSemaphoreCounterValue-device-parameter"
  id="VUID-vkGetSemaphoreCounterValue-device-parameter"></a>
  VUID-vkGetSemaphoreCounterValue-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetSemaphoreCounterValue-semaphore-parameter"
  id="VUID-vkGetSemaphoreCounterValue-semaphore-parameter"></a>
  VUID-vkGetSemaphoreCounterValue-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-vkGetSemaphoreCounterValue-pValue-parameter"
  id="VUID-vkGetSemaphoreCounterValue-pValue-parameter"></a>
  VUID-vkGetSemaphoreCounterValue-pValue-parameter  
  `pValue` **must** be a valid pointer to a `uint64_t` value

- <a href="#VUID-vkGetSemaphoreCounterValue-semaphore-parent"
  id="VUID-vkGetSemaphoreCounterValue-semaphore-parent"></a>
  VUID-vkGetSemaphoreCounterValue-semaphore-parent  
  `semaphore` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetSemaphoreCounterValue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700

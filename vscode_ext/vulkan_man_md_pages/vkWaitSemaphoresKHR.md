# vkWaitSemaphores(3) Manual Page

## Name

vkWaitSemaphores - Wait for timeline semaphores on the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To wait for a set of semaphores created with a
[VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`
to reach particular counter values on the host, call:

``` c
// Provided by VK_VERSION_1_2
VkResult vkWaitSemaphores(
    VkDevice                                    device,
    const VkSemaphoreWaitInfo*                  pWaitInfo,
    uint64_t                                    timeout);
```

or the equivalent command

``` c
// Provided by VK_KHR_timeline_semaphore
VkResult vkWaitSemaphoresKHR(
    VkDevice                                    device,
    const VkSemaphoreWaitInfo*                  pWaitInfo,
    uint64_t                                    timeout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the semaphores.

- `pWaitInfo` is a pointer to a
  [VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitInfo.html) structure containing
  information about the wait condition.

- `timeout` is the timeout period in units of nanoseconds. `timeout` is
  adjusted to the closest value allowed by the implementation-dependent
  timeout accuracy, which **may** be substantially longer than one
  nanosecond, and **may** be longer than the requested period.

## <a href="#_description" class="anchor"></a>Description

If the condition is satisfied when `vkWaitSemaphores` is called, then
`vkWaitSemaphores` returns immediately. If the condition is not
satisfied at the time `vkWaitSemaphores` is called, then
`vkWaitSemaphores` will block and wait until the condition is satisfied
or the `timeout` has expired, whichever is sooner.

If `timeout` is zero, then `vkWaitSemaphores` does not wait, but simply
returns information about the current state of the semaphores.
`VK_TIMEOUT` will be returned in this case if the condition is not
satisfied, even though no actual wait was performed.

If the condition is satisfied before the `timeout` has expired,
`vkWaitSemaphores` returns `VK_SUCCESS`. Otherwise, `vkWaitSemaphores`
returns `VK_TIMEOUT` after the `timeout` has expired.

If device loss occurs (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-lost-device"
target="_blank" rel="noopener">Lost Device</a>) before the timeout has
expired, `vkWaitSemaphores` **must** return in finite time with either
`VK_SUCCESS` or `VK_ERROR_DEVICE_LOST`.

Valid Usage (Implicit)

- <a href="#VUID-vkWaitSemaphores-device-parameter"
  id="VUID-vkWaitSemaphores-device-parameter"></a>
  VUID-vkWaitSemaphores-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkWaitSemaphores-pWaitInfo-parameter"
  id="VUID-vkWaitSemaphores-pWaitInfo-parameter"></a>
  VUID-vkWaitSemaphores-pWaitInfo-parameter  
  `pWaitInfo` **must** be a valid pointer to a valid
  [VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitInfo.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_TIMEOUT`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkWaitSemaphores"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
